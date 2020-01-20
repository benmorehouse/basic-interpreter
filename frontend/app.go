package main

import(
	log "github.com/sirupsen/logrus"
	"os"
	"encoding/json"
	"io/ioutil"
)

type App struct{
	Pages	   []Page
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
		Port            int    `json:"Port"`
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

	session := Session{
		LoggedIn: false,
		Username: "",
	}

	pages := []Page{
		LoadAboutPage(),
		LoadLoginPage(),
		LoadGithubPage(),
		LoadTerminalPage(),
	}

	logger, err := NewLogger(config.LogFileName)
	if err != nil{
		return nil, err
	}

	a := App{
		Pages: pages,
		Log: logger,
		User: &session,
		ConfigFile: confData,
	}

	return &a, nil
}

