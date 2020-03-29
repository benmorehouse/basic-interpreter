package main

import (
	log "github.com/sirupsen/logrus"
)

// App is the main structure of the web app
type App struct {
	Pages           []Page // we might not need this
	User            *Session
	Config          *AppConf
	ConfigFile      string
	connection      *DBcxn
	operatingSystem *OperatingSystem
}

// NewApp generates a new app structure
func NewApp(configFile string, init bool) (*App, error) {

	a := App{ConfigFile: configFile}

	a.LoadConfig()

	a.User = &Session{
		LoggedIn:  true,
		FirstName: "Ben Morehouse",
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

	log.Debug("Before initializing the operating system")
	a.InitOS()
	log.Debug("after initializing the operating system")

	return &a, nil
}
