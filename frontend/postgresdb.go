package main

import(
	"fmt"
	"database/sql"
	"strconv"
	"errors"
	"time"
	"context"
	"math/rand"

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

/* 

	during testing we will leave this.

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

	tableQuery := "create table " + a.Config.UserTable + "(userid varchar(80), password varchar(80), email varchar(80), firstname varchar(80), lastname varchar(80));"
	if _, err = d.cxn.ExecContext(*d.context, tableQuery); err == nil {
		return errors.New("User table didn't exist before... restart basic") // we need to add this to the CLI
	}

	return nil
}

/*
func (d *DBcxn) PostgresSelectUser(username string)(string, error){

}
*/

/*------------------- User Auth Code -------------------------------*/

//returns boolean for if email exists in database
func (d *DBcxn) PostgresEmailExists(email string)(bool, error){
	if email == "" {
		log.Error("Didn't recieve an email in this function")
		return false, errors.New("Didn't recieve an email in this function")
	}

	if err := d.cxn.PingContext(*d.context); err != nil {
		return false, err
	}

	emailQuery := "select count(email) from " + d.UserTable
	emailQuery += " where email=\"" + email + "\";"
	log.Info("Email query:",emailQuery)

	result, err := d.cxn.ExecContext(*d.context, emailQuery)
	if err != nil{
		log.Error(err)
		return false, err
	}

	count, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return false, err
	}

	return count == 1, nil
}

// returns the hashed password in the database given the email
func (d *DBcxn) PostgresGetPassword(email string)([]byte, error){
	// ping database connection for activity firstly
	if err := d.cxn.PingContext(*d.context); err != nil {
		return nil, err
	}

	passwordQuery := "select password from " + d.UserTable
	passwordQuery = " where user=" + email + ";"

	result := d.cxn.QueryRowContext(*d.context, passwordQuery)

	i := struct{
		password string
	}{}

	if err := result.Scan(i); err != nil { //scans and puts row result to password interface
		return nil, err
	}

	return []byte(i.password), nil
}

func (a *App) ValidateUserLogin(email, password string)(bool, error){
	// this password needs to go through a hashing first

	_, err := a.connection.PostgresEmailExists(email)
	if err != nil {
		return false, err
	}

	hash, err := a.connection.PostgresGetPassword(email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *App) CreateUser(requestBody *RequestBody) (bool, error){
	if &requestBody == nil {
		return false, errors.New("Request body is nil")
	}

	exists, err := a.connection.PostgresEmailExists(requestBody.Email)
	if err != nil {
		return false, err
	} else if exists == true{
		return false, errors.New("Email already present in our database!")
	}

	if err := a.connection.PostgresCreateUser(requestBody); err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (d *DBcxn) PostgresCreateUser(requestBody *RequestBody) error {
	if err := d.cxn.PingContext(*d.context); err != nil {
		return err
	}

	if err := PasswordStrength(requestBody.ConfirmPassword); err != nil{
		return err
	}

	userId := generateUserId()
	s := "Inset into %s values(\"%s\",\"%s\",\"%s\",\"%s\")"

	insertQuery := fmt.Sprintf(s, d.UserTable, userId, requestBody.FirstName, requestBody.LastName, requestBody.Email)

	_, err := d.cxn.ExecContext(*d.context, insertQuery)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func PasswordStrength(password string) error{
	if password == ""{
		return errors.New("Password has been left empty")
	} else if len(password) < 8 {
		return errors.New("Password must be 8 digits long")
	}

	return nil
}

func generateUserId()string{
	rand.Seed(time.Now().UTC().UnixNano())
	userid := ""

	for i:=0; i<20; i++{
		userid += strconv.Itoa(rand.Intn(10))
	}

	return userid
}
