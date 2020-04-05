package main

import (
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

// OperatingSystem is a struct which maintains the current directory
// and maps cli inputs to functions.
type OperatingSystem struct {
	CurrentDirectory *Directory
	CommandMap       map[string]Command
	CommandPipe      chan string
	ResponsePipe     chan CommandResponse
	DonePipe         chan bool // this is only to be called at server termination.
	Connection       *DBcxn    // a pointer to the app's database connection
	User             *Session
}

// InitOS is called when the server first starts.
func (a *App) InitOS() {

	home, err := NewDirectory(a.User.FirstName, nil) // this should take the username of the user
	if err != nil {
		err := OperatingSystemError(InvalidNameError, err)
		log.Error(err)
		return
	}

	pipe := make(chan string)
	doneChan := make(chan bool, 1)
	responsePipe := make(chan CommandResponse)

	os := &OperatingSystem{
		CurrentDirectory: home,
		CommandPipe:      pipe,
		DonePipe:         doneChan,
		ResponsePipe:     responsePipe,
		Connection:       a.connection,
		User:             a.User,
	}

	os.createCommandMap()
	a.operatingSystem = os
}

// The command interface has just one method: process the command.
type Command interface {
	Process([]string) *CommandResponse
}

// CommandResponse is the struct returned by every command.
type CommandResponse struct {
	Success bool
	Error   error
	Output  string
}

// HandleProcess is used by operating system to initiate all commands
func (os *OperatingSystem) RunCommand(cmd Command, input []string) {

	if cmd == nil {
		log.Error("No command found for this request")
		response := CommandResponse{
			Output:  "Command not found",
			Success: false,
		}
		os.ResponsePipe <- response
		return
	}
	response := cmd.Process(input)
	os.ResponsePipe <- *response
}

// createCommandMap will create the command map for the operating system
func (os *OperatingSystem) createCommandMap() {

	m := make(map[string]Command)
	m["ls"] = &ls{os: os}
	m["mkdir"] = &mkdir{os: os}
	m["cd"] = &chdir{os: os}
	m["pwd"] = &pwd{os: os}
	m["rm"] = &rm{os: os, recursive: false}
	m["rmdir"] = &rm{os: os, recursive: true}
	m["open"] = &rm{os: os, recursive: true}
	m["touch"] = &touch{os: os}
	m["clear"] = &clear{returnVal: "clear"}
	m["help"] = &clear{returnVal: "help"}
	os.CommandMap = m
}

// This is a function that will run continuously until the program terminates.
// It takes in the input through its channel and generates output based on it, which is then pushed through the same output.
// This is run as a goroutine in the app instance upon opening the terminal page.
func (os *OperatingSystem) RunOperatingSystem() {

	log.Info("The operating system has started")
	doneCalled := false
	for {
		if doneCalled {
			break
		}

		select {
		case <-os.DonePipe:
			doneCalled = true
			break
		case commandLine := <-os.CommandPipe:
			commands := strings.Fields(commandLine)
			command := strings.ToLower(commands[0])
			os.RunCommand(os.CommandMap[strings.ToLower(command)], commands[1:])
		}
	}
}

// ########################################################################
// ######################## Directories ###################################

// map each function to a string to be called by the terminal
// Directory is a struct that will hold all Files and subdirs in a single interface.
type Directory struct {
	Name           string
	SubDirectories map[string]*Directory
	SubFiles       map[string]*File
	Parent         *Directory
	IsRoot         bool
}

// NewDirectory will create a new directory to be used by the system.
func NewDirectory(name string, parentDir *Directory) (*Directory, error) {

	if strings.TrimSpace(name) == "" {
		return nil, OperatingSystemError(InvalidNameError, nil)
	}

	dirs, Files := make(map[string]*Directory), make(map[string]*File)
	isRoot := (parentDir == nil)
	d := &Directory{
		Name:           strings.TrimSpace(name),
		SubDirectories: dirs,
		SubFiles:       Files,
		IsRoot:         isRoot,
		Parent:         parentDir,
	}

	return d, nil
}

// List is a method which lists of contents of a Directory.
func (d *Directory) List() []string {

	var arrayOfNames []string
	for name := range d.SubDirectories {
		arrayOfNames = append(arrayOfNames, name)
	}

	for name := range d.SubFiles {
		arrayOfNames = append(arrayOfNames, name)
	}

	sort.Strings(arrayOfNames)
	return arrayOfNames
}

// AddDirectory is a directory specific command to abstract adding a directory.
func (d *Directory) AddDirectory(name string) error {

	subDir, err := NewDirectory(name, d)
	if err != nil {
		log.Error(err)
		return err
	}

	d.SubDirectories[subDir.Name] = subDir
	return nil
}

// RemoveDirectory is a directory specific command to abstract removing a directory.
func (d *Directory) RemoveDirectory(subDir string) error {

	if strings.TrimSpace(subDir) == "" {
		err := OperatingSystemError(InvalidNameError, nil)
		log.Error(err)
		return err
	}

	if d.SubDirectories[subDir] == nil {
		if d.SubFiles[subDir] != nil {
			err := OperatingSystemError(DirectoryIsFile, nil)
			log.Error(err)
			return err
		}
		err := OperatingSystemError(DirectoryIsNil, nil)
		log.Error(err)
		return err
	}

	d.SubDirectories[subDir] = nil
	return nil
}

// ProvidePath will provide the path as a string to the passed in directory
func (d *Directory) ProvidePath() string {

	if d.IsRoot {
		return "/" + d.Name + "/"
	}

	return d.Parent.ProvidePath() + d.Name + "/"
}

// GetSubDirectory will return the subdirectory of the file
func (d *Directory) GetSubDirectory(nextDirName string) *Directory {

	return d.SubDirectories[nextDirName]
}

// File is a struct which contains the name as well as the method to get the File contents from the database.
type File struct {
	Name    string
	Path    string
	Basic   bool // if it is a basic file, otherwise shrug
	content []string
}

// NewFile will return a File structure
func NewFile(name string) (*File, error) {

	if strings.TrimSpace(name) == "" {
		return nil, OperatingSystemError(InvalidNameError, nil)
	}

	f := &File{
		Name:    name,
		content: nil,
	}
	return f, nil
}

// AddFile will add a file to a given directory
// and also add this file to the fileStore
func (d *Directory) AddFile(newFileName string) error {

	file, err := NewFile(newFileName)
	if err != nil {
		log.Error(err)
		return err
	}

	file.Path = d.ProvidePath()
	d.SubFiles[file.Name] = file
	return nil
}

// RemoveFile will remove a file from a given directory
func (d *Directory) RemoveFile(fileName string) error {

	if strings.TrimSpace(fileName) == "" {
		err := OperatingSystemError(InvalidNameError, nil)
		log.Error(err)
		return err
	}

	if d.SubFiles[fileName] == nil {
		if d.SubDirectories[fileName] != nil {
			err := OperatingSystemError(DirectoryIsFile, nil)
			log.Error(err)
			return err
		}
		err := OperatingSystemError(DirectoryIsNil, nil)
		log.Error(err)
		return err
	}

	d.SubFiles[fileName] = nil
	return nil
}

// OpenFile will return the file in the ajax request.
func (d *Directory) OpenFile(fileName string) (*File, error) {

	if d == nil {
		err := OperatingSystemError(DirectoryIsNil, nil)
		log.Error(err)
		return nil, err
	}

	fileName = strings.TrimSpace(strings.ToLower(fileName))
	if file := d.SubFiles[fileName]; file != nil {
		return file, nil
	}

	err := OperatingSystemError(NoFileFound, nil)
	log.Error(err)
	return nil, err

}

// IsBasicFile will return whether or not the file is a basic
// file based on the file name.
func (f *File) CheckIfBasicFile() {
	validate := func() bool {

		if strings.TrimSpace(f.Name) == "" {
			return false
		}

		possibleFiles := make(map[string]bool)
		possibleFiles["bsc"] = true
		possibleFiles["basic"] = true

		fields := strings.SplitAfter(f.Name, ".")
		if len(fields) < 2 {
			return false
		}

		return possibleFiles[fields[0]]
	}

	f.Basic = validate()
}

// ReadFileFromFilestore will populate the contents of the file
// from the byteslice given by the database instance.
func (f *File) ReadFileFromFilestore(content []byte) {

	if content == nil {
		return
	}

	var fileContent []string
	var buffer []byte

	for _, digit := range content {
		switch digit {
		case '\n', '\r':
			fileContent = append(fileContent, string(buffer))
			buffer = nil
		default:
			buffer = append(buffer, digit)
		}
	}

	f.content = fileContent
}

// WriteFileForSaving will write contents to a byte buffer
func (f *File) WriteFileForSaving() []byte {

	var buffer []byte
	for _, row := range f.content {
		for _, value := range row {
			buffer = append(buffer, byte(value))
		}
		buffer = append(buffer, byte('\n'))
	}
	return buffer
}

// ########################################################################
// ######################## Commands ######################################

// ls is the list command for a directory
type ls struct {
	os *OperatingSystem
}

// Process implements the Command interface for ls
func (cmd *ls) Process(_ []string) *CommandResponse {
	response := &CommandResponse{
		Success: true,
		Error:   nil,
		Output:  strings.Join(cmd.os.CurrentDirectory.List(), " "),
	}

	return response
}

// chdir implements the change directory command using the Command interface
type chdir struct {
	os *OperatingSystem
}

// Process implements the Command interface
func (cmd *chdir) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if cmd.os == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if input == nil || len(input) == 0 {
		cmd.TraceToRoot()
		response.Success = true
		return response
	}

	nextDirName := strings.TrimSpace(input[0])
	switch nextDirName {
	case "..":
		if cmd.os.CurrentDirectory.Parent != nil {
			cmd.os.CurrentDirectory = cmd.os.CurrentDirectory.Parent
		}
	case "":
		cmd.TraceToRoot()
	default:
		d := cmd.os.CurrentDirectory.GetSubDirectory(nextDirName)
		if d == nil {
			response.Success = false
			response.Output = "no such file or directory: " + nextDirName
			return response
		}

		cmd.os.CurrentDirectory = d
	}

	response.Success = true
	return response
}

// TraceToRoot will change directory back to the root directory of the program.
func (cmd *chdir) TraceToRoot() {

	d := cmd.os.CurrentDirectory
	for d != nil && d.Parent != nil {
		d = d.Parent
	}
}

// mkdir is the make directory command for a directory
type mkdir struct {
	os          *OperatingSystem
	nextDirName string
}

// Process implements the command interface
func (cmd *mkdir) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if input == nil || len(input) == 0 {
		response.Success = false
		response.Error = OperatingSystemError(NoDirectoryGiven, nil)
		return response

	}

	cmd.nextDirName = input[0]
	if cmd.os.CurrentDirectory == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if err := cmd.os.CurrentDirectory.AddDirectory(cmd.nextDirName); err != nil {
		log.Error(err)
		response.Error = err
		response.Success = false
		return response
	}

	response.Success = true
	return response
}

// pwd is the make directory command for a directory
type pwd struct {
	os *OperatingSystem
}

// Process implements the command interface
func (cmd *pwd) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if cmd.os == nil {
		response.Success = false
		response.Error = OperatingSystemError(OperatingSystemIsNil, nil)
		return response
	}

	path := cmd.os.CurrentDirectory.ProvidePath()
	response.Output = path
	response.Success = true
	return response
}

// touch will create a file within the instance.
type touch struct {
	os           *OperatingSystem
	nextFileName string
}

// Process implements the command interface
func (cmd *touch) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if input == nil || len(input) == 0 {
		response.Success = false
		response.Error = OperatingSystemError(NoFileNameGiven, nil)
		return response

	}

	cmd.nextFileName = input[0]
	if cmd.os.CurrentDirectory == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if err := cmd.os.CurrentDirectory.AddFile(cmd.nextFileName); err != nil {
		log.Error(err)
		response.Error = err
		response.Success = false
		return response
	}

	go func() {
		//fileHash, err := ValidateFileName(cmd.nextFileName, cmd.os.CurrentDirectory.ProvidePath())
		newFile, err := cmd.os.CurrentDirectory.OpenFile(cmd.nextFileName)
		if err != nil {
			log.Error(err)
			return
		}

		if err := cmd.os.Connection.Save(newFile); err != nil {
			log.Error(err)
			return
		}
	}()

	response.Success = true
	return response
}

// rm will delete a file or directory within the instance
type rm struct {
	os         *OperatingSystem
	recursive  bool
	deleteName string
}

// Process implements the command interface
func (cmd *rm) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if input == nil || len(input) == 0 {
		response.Success = false
		response.Error = OperatingSystemError(NoFileNameGiven, nil)
		return response
	}

	if cmd.os.CurrentDirectory == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if strings.ToLower(input[0]) == "-r" {
		cmd.recursive = true
	}

	if cmd.recursive {
		if err := cmd.os.CurrentDirectory.RemoveDirectory(cmd.deleteName); err != nil {
			response.Error = err
			response.Success = false
			response.Output = err.Error()
			return response
		}
	}

	if err := cmd.os.CurrentDirectory.RemoveFile(cmd.deleteName); err != nil {
		response.Error = err
		response.Success = false
		response.Output = err.Error()
		return response
	}

	response.Success = true
	return response

}

// open will create a file within the instance.
type open struct {
	os *OperatingSystem
}

// Process implements the command interface
func (cmd *open) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if input == nil || len(input) == 0 {
		response.Success = false
		response.Error = OperatingSystemError(NoFileNameGiven, nil)
		return response
	}

	if cmd.os.CurrentDirectory == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	// NOTE: This will return a file that we need to eventually pipe
	_, err := cmd.os.CurrentDirectory.OpenFile(input[0])
	if err != nil {
		log.Error(err)
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		response.Output = "No file found"
		return response
	}

	response.Success = true
	return response
}

// compile will compile the given basic program. This can also be triggered by a button
// from within the xterm text editor.
type compile struct {
	os *OperatingSystem
}

// Process implements the command interface
func (cmd *compile) Process(input []string) *CommandResponse {

	response := &CommandResponse{}
	if input == nil || len(input) == 0 {
		response.Success = false
		response.Error = OperatingSystemError(NoFileNameGiven, nil)
		return response
	}

	if cmd.os.CurrentDirectory == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	// NOTE: This will return a file that we need to eventually pipe
	_, err := cmd.os.CurrentDirectory.OpenFile(input[0])
	if err != nil {
		log.Error(err)
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		response.Output = "No file found"
		return response
	}

	// NOTE: After we get the file above we need to pipe it to a filestore object.

	response.Success = true
	return response
}

// clear is the make directory command for a directory
type clear struct {
	returnVal string
}

// Process implements the command interface
func (cmd *clear) Process(_ []string) *CommandResponse {

	response := &CommandResponse{
		Output:  cmd.returnVal,
		Success: true,
	}
	return response
}

// NOTE: still a open command
// move command and a compile command
// clear is the make directory command for a directory
type help struct {
	returnVal string
}

// Process implements the command interface
func (cmd *help) Process(_ []string) *CommandResponse {

	response := &CommandResponse{
		Output:  cmd.returnVal,
		Success: true,
	}
	return response
}
