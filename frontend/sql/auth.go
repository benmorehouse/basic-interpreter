package sql

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
