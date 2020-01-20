package main

import(
	log "github.com/sirupsen/logrus"
	"os"
	"errors"
	"encoding/json"
)

type App struct{
	Pages	   []*Pages
	Log        *log.Logger
	User       *Session
	ConfigFile interface{}
}

type Session struct{
	LoggedIn    bool
	Username    string
}

func NewApp() (*App, error){
	jsonFile, err := os.Open("conf.json")
	defer jsonFile.Close()
	if err != nil{
		return nil, err
	}

	config := struct{
		Port            string `json:"Port"`
		LogFileName     string `json:"LogFileName"`
		DBName		string `json:"BasicDB"`
		DBUser		string `json:"DBUser"`
		DBPass		string `json:"DBPass"`
		UserTable	string `json:"UserTable"`
		BasicOutFile	string `json:"BasicOutFile"`
		BasicInFile	string `json:"BasicInFile"`
	}{}

	confData, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		return nil, err
	}

	if err = json.Unmarshal(confData, &config); err != nil{
		return nil, err
	}

	fmt.Println(config)

	f, err := os.OpenFile("basicInterpreterLogs.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return nil, errors.New("Couldnt create log file:",err)
	}
/*
	app := App{
		Log: log.Logger{
			Out: f,
		},
		User:
		Pages:
		ConfigFile:
	}
*/
	return nil, nil

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


