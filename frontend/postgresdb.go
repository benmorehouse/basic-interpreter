package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // driver for postgresql database
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//returns boolean for if email exists in database
func (d *DBcxn) PostgresEmailExists(email string) (bool, error) {
	if email == "" {
		log.Error(PostgresError(NoEmailRecieved, nil))
		return false, PostgresError(NoEmailRecieved, nil)
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

// PostgresCreateUser will create the user and instantiate it into the database
func (d *DBcxn) PostgresCreateUser(requestBody *AuthRequestBody) error {
	if err := d.cxn.PingContext(*d.context); err != nil {
		return err
	}

	if err := PasswordStrength(requestBody.ConfirmPassword); err != nil {
		return err
	}

	userId := generateUserId()

	password, err := bcrypt.GenerateFromPassword([]byte(requestBody.CreatePassword), 10)
	if err != nil {
		log.Error(err)
		return err
	}

	insertQuery := "Insert into ? (userid, password, email, firstname, lastname) values('?','?','?','?','?')"

	log.Info(insertQuery)
	if _, err := d.cxn.ExecContext(*d.context, insertQuery,
		d.UserTable, userId, password,
		requestBody.Email, requestBody.FirstName, requestBody.LastName); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

/*------------------- User Auth Code -------------------------------*/

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

func (a *App) EstablishDbcxn(init bool) error {
	if &a == nil || &a.Config == nil {
		log.Error(PostgresError(BadConfiguration, nil))
		return PostgresError(BadConfiguration, nil)
	}

	conf := a.Config
	if conf.DBHost == "" {
		log.Error(PostgresError(DBHostNil, nil))
		return PostgresError(DBHostNil, nil)
	}

	if conf.DBName == "" {
		log.Error(PostgresError(DBNameNil, nil))
		return PostgresError(DBNameNil, nil)
	}

	if &conf.DBPort == nil {
		log.Error(PostgresError(DBPortNil, nil))
		return PostgresError(DBPortNil, nil)
	}

	if conf.DBUser == "" {
		log.Error(PostgresError(DBUserNil, nil))
		return PostgresError(DBUserNil, nil)
	}

	if conf.UserTable == "" {
		log.Error(PostgresError(DBUserTableNil, nil))
		return PostgresError(DBUserTableNil, nil)
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
		log.Error(PostgresError(NoConnection, err))
		return PostgresError(NoConnection, err)
	}
	defer db.Close()

	cont := context.Background()

	cxn, err := db.Conn(cont)
	if err != nil {
		log.Error(PostgresError(NoConnection, err))
		return PostgresError(NoConnection, err)
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
	if init {
		if err := a.createTableIfNotExists(); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

func (d *DBcxn) PingContext() error {

	if err := d.cxn.PingContext(*d.context); err != nil {
		err := PostgresError(NoConnection, err)
		log.Error(err)
		return err
	}
	return nil
}

func (a *App) createUserTableIfNotExists() error {
	conf := a.Config
	c := a.connection
	// this should use a ? from the database/sql library
	query := "Create table ?"
	query += `
	(
		userid varchar(30),
		password text,
		email varchar(30),
		firstname varchar(30),
		lastname varchar(30)
	);
	`

	_, err := c.cxn.ExecContext(*c.context, query, conf.UserTable)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (a *App) createFileTableIfNotExists() error {

	conf := a.Config
	c := a.connection
	// this should use a ? from the database/sql library
	query := "Create table ?"
	query += `
	(
		id varchar(30),
		userid varchar(30),
		file LONGBLOB,
	);
	`
	_, err := c.cxn.ExecContext(*c.context, query, conf.FileTable)
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
