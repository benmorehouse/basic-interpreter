package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq" // driver for postgresql database
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// SQLSaver is an interface a struct can implement to allow for the save method.
type SQLSaver interface {
	TableName() string
	TableColumns() []string
	TableValues() []string
}

// Save will save a SQLSaver into the database instance.
func (a *App) Save(sqlType SQLSaver) error {

	tableName := sqlType.TableName()
	tableColumns := sqlType.TableColumns()
	tableValues := sqlType.TableValues()

	if tableName == "" || tableColumns == nil || tableValues == nil {
		err := PostgresError(BadInterface, nil)
		log.Error(err)
		return err
	} else if len(tableColumns) != len(tableValues) {
		err := PostgresError(BadInterface, nil)
		log.Error(err)
		return err
	}

	query := "insert into ? ("
	q := ""

	for _, column := range tableColumns {
		q += column + ","
	}

	q = strings.TrimRight(q, ",")
	query += q + ") VALUES ("
	q = ""

	for _, _ = range tableValues {
		q += "?, "
	}

	q = strings.TrimRight(q, ",")
	query += q + ");"

	// this SUCKS.... is there another way?
	var values []interface{} = make([]interface{}, len(tableValues))
	for i, value := range tableValues {
		values[i] = value
	}

	if _, err := a.cxn.ExecContext(*a.cxn.context, query, values...); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
