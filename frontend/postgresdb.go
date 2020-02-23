package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // driver for postgresql database
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type DBcxn struct {
	Host      string
	Name      string
	Password  string
	User      string
	Port      int
	UserTable string
	cxn       *sql.Conn //unexported connection to database
	context   *context.Context
}

func (a *App) EstablishDbcxn() error {
	if &a == nil || &a.Config == nil {
		log.Error("Failed to establish database credentials: app not configured correctly")
		return errors.New("Failed to establish database credentials: app not configured correctly")
	}

	conf := a.Config
	if conf.DBHost == "" {
		log.Error("DBHost unexpectedly found nil")
		return errors.New("DBHost unexpectedly found nil")
	}

	if conf.DBName == "" {
		log.Error("DBName unexpectedly found nil")
		return errors.New("DBName unexpectedly found nil")
	}

	if &conf.DBPort == nil {
		log.Error("DBPort unexpectedly found nil")
		return errors.New("DBPort unexpectedly found nil")
	}

	if conf.DBUser == "" {
		log.Error("DBUser unexpectedly found nil")
		return errors.New("DBUser unexpectedly found nil")
	}

	if conf.UserTable == "" {
		log.Error("DBUserTable unexpectedly found nil")
		return errors.New("DBUserTable unexpectedly found nil")
	}

	var psqlInfo string
	if conf.DBPass == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			conf.DBHost,
			conf.DBPort,
			conf.DBUser,
			conf.DBName,
		)
	} else {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			conf.DBHost,
			conf.DBPort,
			conf.DBUser,
			conf.DBPass,
			conf.DBName,
		)
	}

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error("Unable to establish connection with user database:", err)
		return errors.New("Unable to establish connection with user database:" + err.Error())
	}
	defer db.Close()

	cont := context.Background()

	cxn, err := db.Conn(cont)
	if err != nil {
		log.Error("Failed to create database connection: ", err)
		return errors.New("Failed to create database connection" + err.Error())
	}

	d := DBcxn{
		Host:      conf.DBHost,
		Name:      conf.DBName,
		Password:  conf.DBPass,
		User:      conf.DBUser,
		Port:      conf.DBPort,
		UserTable: conf.UserTable,
		cxn:       cxn,
		context:   &cont,
	}

	a.connection = &d
	if err := a.createTableIfNotExists(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *App) createTableIfNotExists() error {
	conf := a.Config
	c := a.connection
	query := "Create table if not exists " + conf.UserTable
	query += `
	(
		userid varchar(30),
		password text,
		email varchar(30),
		firstname varchar(30),
		lastname varchar(30)
	);
	`

	_, err := c.cxn.ExecContext(*c.context, query)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*
func (d *DBcxn) PostgresSelectUser(username string)(string, error){

}
*/

/*------------------- User Auth Code -------------------------------*/

//returns boolean for if email exists in database
func (d *DBcxn) PostgresEmailExists(email string) (bool, error) {
	if email == "" {
		log.Error("Didn't recieve an email in this function")
		return false, errors.New("Didn't recieve an email in this function")
	}

	if err := d.cxn.PingContext(*d.context); err != nil {
		return false, err
	}

	s := "select count(email) from %s where email='%s';"
	emailQuery := fmt.Sprintf(s, d.UserTable, email)
	log.Info("Email query:", emailQuery)

	var count int
	if err := d.cxn.QueryRowContext(*d.context, emailQuery).Scan(&count); err != nil {
		log.Error(err)
		return false, err
	}

	return count == 1, nil
}

// returns the hashed password in the database given the email
func (d *DBcxn) PostgresGetPassword(email string) ([]byte, error) {
	// ping database connection for activity firstly
	if err := d.cxn.PingContext(*d.context); err != nil {
		log.Error(err)
		return nil, err
	}

	s := "select password from %s where email='%s';"
	passwordQuery := fmt.Sprintf(s, d.UserTable, email)
	log.Info(passwordQuery)
	result := d.cxn.QueryRowContext(*d.context, passwordQuery)

	var password string
	if err := result.Scan(&password); err != nil { //scans and puts row result to password interface
		log.Error(err)
		return nil, err
	}

	return []byte(password), nil
}

func (d *DBcxn) PostgresCreateUser(requestBody *AuthRequestBody) error {
	if err := d.cxn.PingContext(*d.context); err != nil {
		return err
	}

	if err := PasswordStrength(requestBody.ConfirmPassword); err != nil {
		return err
	}

	userId := generateUserId()
	s := "Insert into %s(userid, password, email, firstname, lastname) values('%s','%s','%s','%s','%s')"

	password, err := bcrypt.GenerateFromPassword([]byte(requestBody.CreatePassword), 10)
	if err != nil {
		log.Error(err)
		return err
	}

	insertQuery := fmt.Sprintf(s, d.UserTable, userId,
		password, requestBody.Email,
		requestBody.FirstName, requestBody.LastName)

	log.Info(insertQuery)
	if _, err := d.cxn.ExecContext(*d.context, insertQuery); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*------------------- User Auth Code -------------------------------*/
