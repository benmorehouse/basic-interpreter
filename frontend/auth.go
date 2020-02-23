package main

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	LoggedIn  bool
	FirstName string
}

func (a *App) SetUserSession(Firstname string) {
	a.User.FirstName = Firstname
	a.User.LoggedIn = true
}

func (a *App) ValidateUserLogin(email, password string) (bool, error) {
	// this password needs to go through a hashing first

	_, err := a.connection.PostgresEmailExists(email)
	if err != nil {
		log.Error(err)
		return false, err
	}

	log.Info(password)
	hash, err := a.connection.PostgresGetPassword(email)
	if err != nil {
		log.Error(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func (a *App) CreateUser(requestBody *AuthRequestBody) error {
	if &requestBody == nil {
		return errors.New("Request body is nil")
	}

	exists, err := a.connection.PostgresEmailExists(requestBody.Email)
	if err != nil {
		return err
	} else if exists == true {
		return errors.New("Email already present in our database!")
	}

	if err := a.connection.PostgresCreateUser(requestBody); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func PasswordStrength(password string) error {
	if password == "" {
		return errors.New("Password has been left empty")
	} else if len(password) < 8 {
		return errors.New("Password must be 8 digits long")
	}

	return nil
}

func generateUserId() string {
	rand.Seed(time.Now().UTC().UnixNano())
	userid := ""

	for i := 0; i < 20; i++ {
		userid += strconv.Itoa(rand.Intn(10))
	}

	return userid
}
