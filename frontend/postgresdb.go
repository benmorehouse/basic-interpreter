package main

import(
	"fmt"
	"database/sql"
	"errors"
	"context"

	"golang.org/x/crypto/bcrypt"
	log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq" // driver for postgresql database
)

type DBcxn struct{
	Host		string
	Name		string
	Password	string
	User		string
	Port		int
	UserTable	string
	cxn		*sql.Conn //unexported connection to database
	context		*context.Context
}

func (a *App) EstablishDbcxn()(error){
	if &a == nil || &a.Config == nil{
		log.Error("Failed to establish database credentials: app not configured correctly")
		return errors.New("Failed to establish database credentials: app not configured correctly")
	}

	conf := a.Config
	if conf.DBHost == ""{
		log.Error("DBHost unexpectedly found nil")
		return errors.New("DBHost unexpectedly found nil")
	}

	if conf.DBName == ""{
		log.Error("DBName unexpectedly found nil")
		return errors.New("DBName unexpectedly found nil")
	}

/* during testing we will leave this.
	if conf.DBPass == ""{
		log.Error("DBPass unexpectedly found nil")
		return errors.New("DBPass unexpectedly found nil")
	}
*/

	if &conf.DBPort == nil{
		log.Error("DBPort unexpectedly found nil")
		return errors.New("DBPort unexpectedly found nil")
	}

	if conf.DBUser == ""{
		log.Error("DBUser unexpectedly found nil")
		return errors.New("DBUser unexpectedly found nil")
	}

	if conf.UserTable == ""{
		log.Error("DBUserTable unexpectedly found nil")
		return errors.New("DBUserTable unexpectedly found nil")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUser,
		conf.DBPass,
		conf.DBName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		log.Error("Unable to establish connection with user database:", err)
		return errors.New("Unable to establish connection with user database:" + err.Error())
	}

	cont := context.Background()

	cxn, err := db.Conn(cont)
	if err != nil{
		log.Error("Failed to create database connection:",err)
		return errors.New("Failed to create database connection" + err.Error())
	}

	d := DBcxn{
		Host:		conf.DBHost,
		Name:		conf.DBName,
		Password:	conf.DBPass,
		User:		conf.DBUser,
		Port:		conf.DBPort,
		UserTable:	conf.UserTable,
		cxn:		cxn,
		context:	&cont,
	}

	a.connection = &d
	return nil
}

/*
func (d *DBcxn) PostgresSelectUser(username string)(string, error){

}
*/

/*------------------- User Auth Code -------------------------------*/

//returns boolean for if email exists in database
func (d *DBcxn) PostgresEmailExists(email string)(bool, error){
	if err := d.cxn.PingContext(*d.context); err != nil{
		return false, err
	}

	emailQuery := "select count(email) from " + d.UserTable
	emailQuery += " where email=" + email + ";"

	result, err := d.cxn.ExecContext(*d.context, emailQuery)
	if err != nil{
		log.Error(err)
		return false, err
	}

	count, err := result.LastInsertId()
	if err != nil{
		log.Error(err)
		return false, err
	}

	return count == 1, nil
}

// returns the hashed password in the database given the email
func (d *DBcxn) PostgresGetPassword(email string)([]byte, error){
	// ping database connection for activity firstly
	if err := d.cxn.PingContext(*d.context); err != nil{
		return nil, err
	}

	passwordQuery := "select password from " + d.UserTable
	passwordQuery = " where user=" + email + ";"

	result := d.cxn.QueryRowContext(*d.context, passwordQuery)

	i := struct{
		password string
	}{}

	if err := result.Scan(i); err != nil{ //scans and puts row result to password interface
		return nil, err
	}

	return []byte(i.password), nil
}

func (a *App) ValidateUserLogin(email, password string)(bool, error){
	// this password needs to go through a hashing first

	_, err := a.connection.PostgresEmailExists(email)
	if err != nil{
		return false, err
	}

	hash, err := a.connection.PostgresGetPassword(email)
	if err != nil{
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil{
		return false, err
	}

	return true, nil
}

