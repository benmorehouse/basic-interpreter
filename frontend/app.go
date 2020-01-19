package main

import(
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
)

type App struct{
	Pages	   []*Pages
	Log        *log.Logger
	User       *Session
	ConfigFile *os.File
}

type Session struct{
	LoggedIn    bool
	Username    string
}

func NewApp() *App, error{
	f, err := os.OpenFile("basicInterpreterLogs.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return nil, errors.New("Couldnt create log file:",err)
	}

	app := App{
		Log: log.Logger{
			Out: f,
		},
		User:
		Pages:
		ConfigFile:
	}

}

func NewLogger() *logrus.Logger{ // returns a logger for good debugging
	f, err := os.OpenFile("basicInterpreterLogs.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		panic(err)
	}
	log := logrus.New()
	log.Out = f
	return log
}




