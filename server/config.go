package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

type AppConf struct {
	Port                   int    `json:ServerPort`
	DBHost                 string `json:DBHost`
	DBName                 string `json:DBName`
	DBUser                 string `json:DBUser`
	DBPass                 string `json:DBPass`
	DBPort                 int    `json:DBPort`
	UserTable              string `json:UserTable`
	FileTable              string `json:FileTable`
	BasicOutFile           string `json:BasicOutFile`
	BasicInFile            string `json:BasicInFile`
	AboutPageURL           string `json:AboutPageURL`
	LoginPageURL           string `json:LoginPageURL`
	LoginAttemptedPageURL  string `json:LoginAttemptedPageURL`
	TerminalPageURL        string `json:TerminalPageURL`
	TextEditorPageURL      string `json:TextEditorPageURL`
	GithubPageURL          string `json:GithubPageURL`
	CreateAccountURL       string `json:CreateAccountURL`
	LoginURL               string `json:LoginAttemptedURL`
	AboutPageFile          string `json:AboutPageFile`
	LoginPageFile          string `json:LoginPageFile`
	TerminalPageFile       string `json:TerminalPageFile`
	GithubPageFile         string `json:GithubPageFile`
	TextEditorFile         string `json:TextEditorFile`
	CompileEndpoint        string `json:CompileEndpoint`
	SaveFileEndpoint       string `json:SaveFileEndpoint`
	RunTerminalEndpoint    string `json:RunTerminalEndpoint`
	RunTextEditorEndpoint  string `json:RunTextEditorEndpoint`
	PathToOperatingSystem  string `json:PathToOperatingSystem`
	PathToBasicInterpreter string `json:PathToBasicInterpreter`
	ScriptsPrefix          string `json:ScriptsPrefix`
	CSSPrefix              string `json:CSSPrefix`
	PathToCSS              string `json:PathToCSS`
	PathToScripts          string `json:PathToScripts`
	PathToBackend          string `json:PathToBackend`
	BasicBinary            string `json:BasicBinary`
}

func (a *App) LoadConfig() {

	jsonFile, err := os.Open(a.ConfigFile)
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
		//		a.Config = getDefaultConfig()
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
		FileTable:             "Filestore",
		AboutPageURL:          "/about",
		LoginPageURL:          "/login",
		TerminalPageURL:       "/terminal",
		GithubPageURL:         "/github",
		TextEditorPageURL:     "/textEditor",
		CreateAccountURL:      "/createAccount",
		LoginAttemptedPageURL: "/loginAttempted",
		AboutPageFile:         "pages/about.gohtml",
		LoginPageFile:         "pages/login.gohtml",
		TerminalPageFile:      "pages/terminal.gohtml",
		GithubPageFile:        "pages/github.gohtml",
		TextEditorFile:        "pages/textEditor.gohtml",
		CompileEndpoint:       "/CompileBasic",
		SaveFileEndpoint:      "/SaveFile",
		RunTerminalEndpoint:   "/RunTerminalEndpoint",
		RunTextEditorEndpoint: "/RunTextEditorEndpoint",
		ScriptsPrefix:         "/scripts/",
		PathToScripts:         "/scripts",
		PathToCSS:             "css/",
		PathToBackend:         "../backend/bin",
		CSSPrefix:             "/css/",
		BasicInFile:           "basic.txt",
		BasicOutFile:          "output.txt",
		BasicBinary:           "basic",
	}

	return config
}
