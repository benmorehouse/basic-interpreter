package main

import(
	"fmt"
	"database/sql"
	"errors"

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

	if conf.DBPass == ""{
		log.Error("DBHost unexpectedly found nil")
		return errors.New("DBHost unexpectedly found nil")
	}

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

	}

	d := DBcxn{
		Host:		conf.DBHost,
		Name:		conf.DBName,
		Password:	conf.DBPass,
		User:		conf.DBUser,
		Port:		conf.DBPort,
		UserTable:	conf.UserTable,
	}
	a.connection = &d
	return nil
}

func (d *DBcxn) PostgresSelect(){
	
}

func (a *App) ValidateUserLogin(email, password string)(bool){
	// this password needs to go through a hashing first
	conf := a.Config
	hashedPassword, err := HashPassword(password)
	if err != nil{
		log.Error("Cant hash the given password given")
		return false
	}

	
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	log.Info("Established connection with database at port:",conf.DBPort)

	if &conf.UserTable == nil || &conf.UserTable == ""{
		log.Error("Table in configuation unexpectadly found to be blank")
		return false
	}

	emailQuery := "select email from " + conf.UserTable
	emailQuery += " where email=" + username

	result, err := db.Exec(userQuery)
	if err != nil{
		return false
	}else if results != email{
		return false
	}

	passwordQuery := "select password from " + a.Config.UserTable
	passwordQuery = " where user=" + username

	result, err = db.Exec(passwordQuery);
	if err != nil{
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(result), []byte(password))
	if err != nil{
		return false
	}

	return true
}

