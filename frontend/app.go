package main

import (
	log "github.com/sirupsen/logrus"
)

type App struct {
	Pages      []Page // we might not need this
	User       *Session
	Config     *AppConf
	ConfigFile string
	connection *DBcxn
}

func NewApp(configFile string, init bool) (*App, error) {

	a := App{ConfigFile: configFile}

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

	if err := a.EstablishDbcxn(init); err != nil {
		log.Error(err)
		return nil, err
	}

	return &a, nil
}
