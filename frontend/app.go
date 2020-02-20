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
	Config	   *AppConf
	connection *DBcxn
}

type AppConf struct{
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

type RequestBody struct{
	FirstName	string `json:"FirstName"`
	LastName	string `json:"LastName"`
	CreatePassword	string `json:"CreatePassword"`
	ConfirmPassword	string `json:"ConfirmPassword"`
	Email		string `json:"Email"`
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

	config := AppConf{}

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
	if err = a.EstablishDbcxn(); err != nil{
		log.Error(err)
		return nil, err
	}

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
	loginResponse := func(success bool, statusMessage string){
		response := struct{
			Success bool		`json:"loginsucceeded"`
			Messsage string		`json:"message"`
		}{
			success,
			statusMessage,
		}

		writeThisResponse, err := json.Marshal(response);
		if err != nil{
			log.Error("Unable to create response for login page:",err)
		}

		_, err = w.Write(writeThisResponse);
		if err != nil{
			log.Error("Unable to write response for login attempt page:",err)
		}
	}

	requestBody := RequestBody{}

	if r.Method != "POST" {
		log.Error("request method not aligned correctly for login function. Request Method:",r.Method)
		loginResponse(false,"request method not aligned correctly for login function. Current Request Method:" + r.Method)
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Error("Couldn't decode request body. Error thrown:",err)
		loginResponse(false, "Couldn't decode request body. Error thrown:" + err.Error())
	}
	// at this point we need to pass it over to the database instance to validate the request
	if result, err := a.ValidateUserLogin(requestBody.Email, requestBody.CreatePassword); err == nil {
		if result {
			log.Info("User successfully authenticated. Rerouting to terminal page")
			loginResponse(true, "")
			return
		}
	} else {
		log.Info("User information not found")
		loginResponse(false, err.Error())
	}

	return
}

func (a *App) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Info("Attempted Login... handling now")
	signUpResponse := func(success bool, statusMessage string){
		response := struct{
			Success bool		`json:"loginsucceeded"`
			Messsage string		`json:"message"`
		}{
			success,
			statusMessage,
		}

		writeThisResponse, err := json.Marshal(response);
		if err != nil {
			log.Error("Unable to create response for login page:",err)
		}

		_, err = w.Write(writeThisResponse);
		if err != nil {
			log.Error("Unable to write response for login attempt page:",err)
		}
	}

	if r.Method != "POST"{
		log.Error("request method not aligned correctly for login function. Request Method:",r.Method)
		signUpResponse(false,"request method not aligned correctly for login function. Current Request Method:" + r.Method)
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil{
		log.Error("Couldn't decode request body. Error thrown:",err)
		signUpResponse(false, "Couldn't decode request body. Error thrown:" + err.Error())
	}

	if result, err := a.CreateUser(requestBody); err == nil{
		if result{
			log.Info("User successfully Created. Rerouting to terminal page")
			signUpResponse(true, "")
			return
		}
	} else {
		log.Info("User successfully Created. Rerouting to terminal page")
		loginResponse(false, err.Error())
	}

	return
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
}

