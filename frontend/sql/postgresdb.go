package sql

import (
	"context"
	"database/sql"
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

func (a *App) createTableIfNotExists() error {
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

/*
func (d *DBcxn) PostgresSelectUser(username string)(string, error){

}
*/

/*------------------- User Auth Code -------------------------------*/
