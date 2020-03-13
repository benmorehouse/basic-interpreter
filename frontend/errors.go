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
	}

	if c.NestedErr != nil {
		return fmt.Sprintf("database error: %s : %s", errorField[c.Type], c.NestedErr.Error())
	}

	return fmt.Sprintf("database error: %s", errorField[c.Type])
}
