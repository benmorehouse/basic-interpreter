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
}

// The command interface has just one method: process the command.
type Command interface {
	Process() *CommandResponse
}

// CommandResponse is the struct returned by every command.
type CommandResponse struct {
	Success bool
	Error   error
	Output  string
}

// HandleProcess is used by operating system to initiate all commands
func (os *OperatingSystem) RunCommand(cmd Command) {

	if cmd == nil {
		log.Error("No command found for this request")
		response := CommandResponse{
			Output:  "Command not found",
			Success: false,
		}
		os.ResponsePipe <- response
		return
	}
	response := cmd.Process()
	os.ResponsePipe <- *response
}

// createCommandMap will create the command map for the operating system
func (os *OperatingSystem) createCommandMap() {

	m := make(map[string]Command)
	m["ls"] = &ls{dir: os.CurrentDirectory}
	m["mkdir"] = &mkdir{currentDir: os.CurrentDirectory}
	m["cd"] = &chdir{currentDir: os.CurrentDirectory, os: os}
	os.CommandMap = m
}

// InitOS is called when the server first starts.
func (a *App) InitOS() {

	home, err := NewDirectory(a.User.FirstName, nil) // this should take the username of the user
	log.Info(a.User.FirstName)
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
	}

	os.createCommandMap()
	a.operatingSystem = os
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
		case command := <-os.CommandPipe:
			log.Info(command)
			log.Info(command)
			log.Info(command)
			os.RunCommand(os.CommandMap[command])
		}
	}
}

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
	d := &Directory{
		Name:           name,
		SubDirectories: dirs,
		SubFiles:       Files,
	}

	return d, nil
}

// List is a method which lists of contents of a Directory.
func (d Directory) List() []string {

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

// File is a struct which contains the name as well as the method to get the File contents from the database.
type File struct {
	Name    string
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
func (d Directory) AddFile(file *File) error {

	if file == nil {
		return OperatingSystemError(DirectoryIsNil, nil)
	}

	d.SubFiles[file.Name] = file
	return nil
}

// RemoveFile will remove a file from a given directory
func (d Directory) RemoveFile(fileName string) error {

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

// ########################################################################
// ######################## Commands ######################################

// ls is the list command for a directory
type ls struct {
	dir *Directory
}

// Process implements the Command interface for ls
func (cmd ls) Process() *CommandResponse {
	response := &CommandResponse{
		Success: true,
		Error:   nil,
		Output:  strings.Join(cmd.dir.List(), " "),
	}

	return response
}

// chdir implements the change directory command using the Command interface
type chdir struct {
	os         *OperatingSystem
	currentDir *Directory
	nextDir    *Directory
}

// Process implements the Command interface
func (cmd chdir) Process() *CommandResponse {

	response := &CommandResponse{}

	if cmd.currentDir == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if cmd.os == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	cmd.os.CurrentDirectory = cmd.nextDir
	response.Success = true
	return response
}

// mkdir is the make directory command for a directory
type mkdir struct {
	currentDir  *Directory
	nextDirName string
}

// Process implements the command interface
func (cmd mkdir) Process() *CommandResponse {

	response := &CommandResponse{}
	if cmd.currentDir == nil {
		response.Success = false
		response.Error = OperatingSystemError(DirectoryIsNil, nil)
		return response
	}

	if err := cmd.currentDir.AddDirectory(cmd.nextDirName); err != nil {
		log.Error(err)
		response.Error = err
		response.Success = false
		return response
	}

	response.Success = true
	return response
}
