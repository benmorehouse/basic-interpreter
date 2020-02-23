package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConf struct {
	Port                  int    `json:ServerPort`
	DBHost                string `json:DBHost`
	DBName                string `json:DBName`
	DBUser                string `json:DBUser`
	DBPass                string `json:DBPass`
	DBPort                int    `json:DBPort`
	UserTable             string `json:UserTable`
	BasicOutFile          string `json:BasicOutFile`
	BasicInFile           string `json:BasicInFile`
	AboutPageURL          string `json:AboutPageURL`
	LoginPageURL          string `json:LoginPageURL`
	LoginAttemptedPageURL string `json:LoginAttemptedPageURL`
	TerminalPageURL       string `json:TerminalPageURL`
	GithubPageURL         string `json:GithubPageURL`
	CreateAccountURL      string `json:CreateAccountURL`
	LoginURL              string `json:LoginAttemptedURL`
	AboutPageFile         string `json:AboutPageFile`
	LoginPageFile         string `json:LoginPageFile`
	TerminalPageFile      string `json:TerminalPageFile`
	GithubPageFile        string `json:GithubPageFile`
	TerminalInputFile     string `json:TerminalInputFile`
	TerminalOutputFile    string `json:TerminalOutputFile`
}

func (a *App) LoadConfig() {

	jsonFile, err := os.Open("conf.json")
	defer jsonFile.Close()
	if err != nil {
		a.Config = getDefaultConfig()
		return
	}

	config := AppConf{}

	confData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		a.Config = getDefaultConfig()
		return
	}

	if err = json.Unmarshal(confData, &config); err != nil {
		a.Config = getDefaultConfig()
		return
	}

	a.Config = &config
}

func getDefaultConfig() *AppConf {

	log.Warning("Getting default config")
	config := &AppConf{
		Port:                  2272,
		DBHost:                "localhost",
		DBName:                "basicinterpreter",
		DBUser:                "benmorehouse",
		DBPass:                "",
		DBPort:                5432,
		UserTable:             "BasicUsers",
		BasicOutFile:          "basicOut.json",
		BasicInFile:           "basicIn.json",
		AboutPageURL:          "/about",
		LoginPageURL:          "/login",
		TerminalPageURL:       "/terminal",
		GithubPageURL:         "/github",
		CreateAccountURL:      "/createAccount",
		LoginAttemptedPageURL: "/loginAttempted",
		AboutPageFile:         "pages/about.gohtml",
		LoginPageFile:         "pages/login.gohtml",
		TerminalPageFile:      "pages/terminal.gohtml",
		GithubPageFile:        "pages/github.gohtml",
		TerminalInputFile:     "terminalInput.txt",
		TerminalOutputFile:    "terminalOutput.txt",
	}

	return config
}
