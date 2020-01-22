package main

import(
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"html/template"

	log "github.com/sirupsen/logrus"
)

type App struct{
	Pages	   []Page // we might not need this
	User       *Session
	Config	   *appConf
}

type appConf struct{
	Port			int    `json:"ServerPort"`

	DBHost			string `json:"DBHost"`
	DBName			string `json:"DBName"`
	DBUser			string `json:"DBUser"`
	DBPass			string `json:"DBPass"`
	DBPort			int    `json:"DBPort"`
	UserTable		string `json:"UserTable"`

	BasicOutFile		string `json:"BasicOutFile"`
	BasicInFile		string `json:"BasicInFile"`

	AboutPageURL		string `json:"AboutPageURL"`
	LoginPageURL		string `json:"LoginPageURL"`
	LoginAttemptedPageURL	string `json:"LoginAttemptedPageURL"`
	TerminalPageURL		string `json:"TerminalPageURL"`
	GithubPageURL		string `json:"GithubPageURL"`
	CreateAccountURL	string `json:"CreateAccountURL"`
	LoginURL		string `json:"LoginAttemptedURL"`

	AboutPageFile		string `json:"AboutPageFile"`
	LoginPageFile		string `json:"LoginPageFile"`
	TerminalPageFile	string `json:"TerminalPageFile"`
	GithubPageFile		string `json:"GithubPageFile"`

	TerminalInputFile	string `json:"TerminalInputFile"`
	TerminalOutputFile	string `json:"TerminalOutputFile"`
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
		Config: &config,
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
	basicTemplate := template.Must(template.ParseFiles(a.Config.AboutPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleLogin(w http.ResponseWriter, r *http.Request){
	log.Info("Login Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.LoginPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleLoginAttempt(w http.ResponseWriter, r *http.Request){
	log.Info("Attempted Login... handling now")
	r.ParseForm()

	email    := r.Form["login-email"]
	password := r.Form["login-password"]
	if len(password) < 8 || email == ""{
		// then handle this error
	}
	// at this point we need to pass it over to the database instance to validate the request

}

func (a *App) HandleCreateAccount(w http.ResponseWriter, r *http.Request){
	log.Info("Attempted to create an account... handling now")
	r.ParseForm()
}

func (a *App) HandleGithub(w http.ResponseWriter, r *http.Request){
	log.Info("Github Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.GithubPageFile))
	err := basicTemplate.Execute(w, nil)
	if err != nil{
		log.Error(err)
	}
}

func (a *App) HandleTerminal(w http.ResponseWriter, r *http.Request){
	log.Info("Terminal Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.TerminalPageFile))
	err := basicTemplate.Execute(w, nil)
	if err != nil{
		log.Error(err)
	}
