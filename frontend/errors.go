package main

import (
	"fmt"
)

//########################################################
//################### Postgres ###########################

type CreateUserErrorType struct {
	Type int
}

const (
	RequestBodyNil      = iota
	EmailAlreadyPresent = iota
)

func CreateUserError(typeof int) *CreateUserErrorType {

	c := &CreateUserErrorType{
		Type: typeof,
	}
	return c
}

func (c CreateUserErrorType) Error() string {

	errorField := []string{
		"request body is nil",
		"email already present in our database",
	}
	return fmt.Sprintf("There was an error while creating a user: %s", errorField[c.Type])
}

//########################################################
//################### Postgres ###########################

type PostgresErrorType struct {
	Type      int
	NestedErr error
}

const (
	BadConfiguration = iota
	DBHostNil        = iota
	DBNameNil        = iota
	DBPortNil        = iota
	DBUserNil        = iota
	DBUserTableNil   = iota
	NoConnection     = iota
	NoEmailRecieved  = iota
	FileIsNil        = iota
	InvalidInsert    = iota
)

func PostgresError(typeof int, err error) *PostgresErrorType {
	c := &PostgresErrorType{
		Type:      typeof,
		NestedErr: err,
	}
	return c
}

func (c PostgresErrorType) Error() string {

	errorField := []string{
		"Failed to establish database credentials: app not configured correctly",
		"DBHost unexpectedly found nil",
		"DBName unexpectedly found nil",
		"DBPort unexpectedly found nil",
		"DBUser unexpectedly found nil",
		"DBUserTable unexpectedly found nil",
		"Unable to establish connection with user database:",
		"Failed to create database connection:",
		"Didn't recieve an email in this function",
		"a file became corrupt",
		"The columns and values dont align",
	}

	if c.NestedErr != nil {
		return fmt.Sprintf("database error: %s : %s", errorField[c.Type], c.NestedErr.Error())
	}

	return fmt.Sprintf("database error: %s", errorField[c.Type])
}

//########################################################
//################### Server #############################

type ServerErrorType struct {
	Type      int
	NestedErr error
}

const (
	CreateAppFailed      = iota
	ScriptDirNotFound    = iota
	ListenAndServeFailed = iota
)

func ServerError(typeof int, err error) *ServerErrorType {
	serverError := &ServerErrorType{
		Type:      typeof,
		NestedErr: err,
	}

	return serverError
}

func (serverError ServerErrorType) Error() string {

	errorField := []string{
		"The app failed to initialize",
		"Directory for javascript files not found",
		"Listen and serve suddenly stopped",
	}

	if serverError.NestedErr != nil {
		return fmt.Sprintf("server error: %s : %s", errorField[serverError.Type], serverError.NestedErr.Error())
	}

	return fmt.Sprintf("server error: %s", errorField[serverError.Type])

}

//########################################################
//################### Operating system ###################

type OperatingSystemErrorType struct {
	Type      int
	NestedErr error
}

// constants to be used
const (
	// <ErrorName> = iota
	DirectoryIsNil       = iota
	InvalidNameError     = iota
	DirectoryIsFile      = iota
	OperatingSystemIsNil = iota
	ProvidePathError     = iota
	NoDirectoryGiven     = iota
	NoFileNameGiven      = iota
	NoFileFound          = iota
)

func OperatingSystemError(typeof int, err error) *OperatingSystemErrorType {
	operatingSystemError := &OperatingSystemErrorType{
		Type:      typeof,
		NestedErr: err,
	}

	return operatingSystemError
}

func (operatingSystemError OperatingSystemErrorType) Error() string {

	errorField := []string{
		//actual thing to display goes here.
		"The passed in directory was found as nil",
		"This is an invalid file or directory name",
		"Wrong command to delete a file.",
		"Operating system has crashed.",
		"Couldn't provide the path for the current directory",
		"Expected a directory to be given",
		"Expected a file name to be given",
	}

	if operatingSystemError.NestedErr != nil {
		return fmt.Sprintf("server error: %s : %s", errorField[operatingSystemError.Type], operatingSystemError.NestedErr.Error())
	}

	return fmt.Sprintf("server error: %s", errorField[operatingSystemError.Type])

}

// ############################################################
// ############### FileStore ##################################

type FileStoreError struct {
	Type      int
	NestedErr error
}

const (
	InvalidFileName       = iota
	InvalidFilePath       = iota
	FileTranslationFailed = iota
	FileAlreadyExists     = iota
)

func NewFileStoreError(typeof int, err error) *FileStoreError {
	fileStoreError := &FileStoreError{
		Type:      typeof,
		NestedErr: err,
	}

	return fileStoreError
}

func (fileStoreError FileStoreError) Error() string {

	errorField := []string{
		"Cannot have #, &, <space>, or : in the file name",
		"The Filepath has been found to be nil",
		"File translation failed",
		"File already exists in the fielstore",
	}

	if fileStoreError.NestedErr != nil {
		return fmt.Sprintf("filestore error: %s : %s", errorField[fileStoreError.Type], fileStoreError.NestedErr.Error())
	}

	return fmt.Sprintf("filestore error: %s", errorField[fileStoreError.Type])

}
