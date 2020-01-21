package main

import(
	log "github.com/sirupsen/logrus"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"html/template"
)

type App struct{
	Pages	   []Page // we might not need this
	User       *Session
	ConfigFile *appConf
}

type appConf struct{
	Port			int    `json:"Port"`
	DBName			string `json:"BasicDB"`
	DBUser			string `json:"DBUser"`
	DBPass			string `json:"DBPass"`
	UserTable		string `json:"UserTable"`
	BasicOutFile		string `json:"BasicOutFile"`
	BasicInFile		string `json:"BasicInFile"`

	AboutPageURL		string `json:"AboutPageURL"`
	LoginPageURL		string `json:"LoginPageURL"`
	LoginAttemptedPageURL	string `json:"LoginAttemptedPageURL"`
	TerminalPageURL		string `json:"TerminalPageURL"`
	GithubPageURL		string `json:"GithubPageURL"`

	AboutPageFile		string `json:"AboutPageFile"`
	LoginPageFile		string `json:"LoginPageFile"`
	LoginAttemptedPageFile	string `json:"LoginAttemptedPageFile"`
	TerminalPageFile	string `json:"TerminalPageFile"`
	GithubPageFile		string `json:"GithubPageFile"`
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

	config := appConf{}

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

	if err != nil{
		return nil, err
	}

	a := App{
		User: &session,
		ConfigFile: &config,
	}

	pages := []Page{
		a.LoadAboutPage(),
		a.LoadLoginPage(),
		a.LoadGithubPage(),
		a.LoadTerminalPage(),
	}
	a.Pages = pages

	return &a, nil
}

func (a *App) HandleAbout(w http.ResponseWriter, r *http.Request){
	log.Info("About Page requested")
	basicTemplate := template.Must(template.ParseFiles("about.gohtml"))
	err := basicTemplate.Execute(w, a)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleLogin(w http.ResponseWriter, r *http.Request){
	log.Info("Login Page requested")
	basicTemplate := template.Must(template.ParseFiles("Login.gohtml"))
	err := basicTemplate.Execute(w, nil)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleLoginAttempt(w http.ResponseWriter, r *http.Request){
	log.Info("Attempted Login... handling now")
}

func (a *App) HandleGithub(w http.ResponseWriter, r *http.Request){
	log.Info("Github Page requested")
	basicTemplate := template.Must(template.ParseFiles("github.gohtml"))
	err := basicTemplate.Execute(w, nil)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleTerminal(w http.ResponseWriter, r *http.Request){
	log.Info("Terminal Page requested")
	basicTemplate := template.Must(template.ParseFiles("index.gohtml"))
	err := basicTemplate.Execute(w, nil)
	if err != nil{
		log.Error(err)
	}
}
