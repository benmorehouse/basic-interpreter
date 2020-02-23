package main

import (
	log "github.com/sirupsen/logrus"
)

type App struct {
	Pages      []Page // we might not need this
	User       *Session
	Config     *AppConf
	connection *DBcxn
}

func NewApp() (*App, error) {

	a := App{}
	a.LoadConfig()

	a.User = &Session{
		LoggedIn:  false,
		FirstName: "",
	}

	pages := []Page{
		a.LoadAboutPage(),
		a.LoadLoginPage(),
		a.LoadGithubPage(),
		a.LoadTerminalPage(),
	}

	a.Pages = pages

	if err := a.EstablishDbcxn(); err != nil {
		log.Error(err)
		return nil, err
	}

	return &a, nil
}
