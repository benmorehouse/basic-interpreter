package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ##########################################################################################
// ############################## Page Endpoints ###########################################

// InitOptions stands as a way to pass the command line parameters upon starting the
// server to the actual instance of the server here below.
type InitOptions struct {
	IsVerbose bool
	IsInit    bool
	Config    string
}

func StartServer(i InitOptions) error {

	setLogger(i.IsVerbose)

	log.Info("Basic Interpreter Server started")
	a, err := NewApp(i.Config, i.IsInit)
	if err != nil {
		err := ServerError(CreateAppFailed, err)
		log.Error(err)
		log.Error("Ending Server lifespan...")
		return err
	}

	log.Info("App successfully intialized")

	router := http.NewServeMux()
	router.HandleFunc(a.Config.AboutPageURL, a.HandleAbout)
	router.HandleFunc(a.Config.TerminalPageURL, a.HandleTerminal)
	router.HandleFunc(a.Config.GithubPageURL, a.HandleGithub)

	// login and sign up handlers  http.FileServer(http.Dir("/pages/script"))
	router.HandleFunc(a.Config.LoginPageURL, a.HandleLogin)
	router.HandleFunc(a.Config.CreateAccountURL, a.HandleCreateAccount)
	router.HandleFunc(a.Config.LoginAttemptedPageURL, a.HandleLoginAttempt)
	// login and sign up handlers

	// takes in the terminal command line input.
	log.Info(a.Config.RunTerminalEndpoint)
	router.HandleFunc(a.Config.RunTerminalEndpoint, a.HandleTerminalNav)

	// this will handle the pushing of where the scripts directory is.
	// this should be called from a scripts directory constant.
	// and should check to see if this directory exists
	// and also then push that to the gohtml
	if _, err := os.Stat(a.Config.PathToScripts); os.IsNotExist(err) {
		log.Error(err)
		return err
	}

	router.Handle(a.Config.ScriptsPrefix, http.StripPrefix(a.Config.ScriptsPrefix, http.FileServer(http.Dir(a.Config.PathToScripts))))
	router.Handle(a.Config.CSSPrefix, http.StripPrefix(a.Config.CSSPrefix, http.FileServer(http.Dir(a.Config.PathToCSS))))

	port := ":" + strconv.Itoa(a.Config.Port) // port is simply used to display the logging message!!
	log.Info("Basic Interpreter Is Waiting...")
	log.Info("LOCAL: http://localhost" + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// HandleLogin renders the login page file
func (a *App) HandleLogin(w http.ResponseWriter, r *http.Request) {

	log.Info("Login Page requested")
	//fileServer :=
	basicTemplate := template.Must(template.ParseFiles(a.Config.LoginPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil {
		log.Error(err)
	}
}

// RedirectIndex is a function that should redirect users to the about endpoint.
func (a *App) RedirectIndex(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "localhost:2272/about", 301)
}

// HandleAbout renders the about page file
func (a *App) HandleAbout(w http.ResponseWriter, r *http.Request) {

	log.Info("About Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.AboutPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil {
		log.Error(err)
	}
}

// HandleGithub renders the login page file
func (a *App) HandleGithub(w http.ResponseWriter, r *http.Request) {

	log.Info("Github Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.GithubPageFile))
	err := basicTemplate.Execute(w, nil)
	if err != nil {
		log.Error(err)
	}
}

// HandleTerminal renders the terminal page file
func (a *App) HandleTerminal(w http.ResponseWriter, r *http.Request) {

	log.Info("Terminal Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.TerminalPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil {
		log.Error(err)
	}
}

//##########################################################################################
//############################## User Auth Endpoints #######################################

// AuthRequestBody is a struct which holds each user's request to authenticate
type AuthRequestBody struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	CreatePassword  string `json:"CreatePassword"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Email           string `json:"Email"`
}

// writeResponse is a function used by the endpoints below to write a
// response to the frontend ajax request
func writeResponse(w http.ResponseWriter, success bool, statusMessage string) {
	response := struct {
		Success  bool   `json:"Success"`
		Messsage string `json:"Message"`
	}{
		success,
		statusMessage,
	}

	if !success {
		log.Error(statusMessage)
	}

	writeThisResponse, err := json.Marshal(response)
	if err != nil {
		log.Error("Unable to create response for page:", err)
		return
	}

	_, err = w.Write(writeThisResponse)
	if err != nil {
		log.Error("Unable to write response for attempt page:", err)
		return
	}
	return
}

// HandleLoginAttempt is triggered when a user tries to login.
func (a *App) HandleLoginAttempt(w http.ResponseWriter, r *http.Request) {

	log.Info("Attempted Login... handling now")
	requestBody := AuthRequestBody{}

	if r.Method != "POST" {
		log.Error("request method not aligned correctly for login function. Request Method:", r.Method)
		writeResponse(w, false, "request method not aligned correctly for login function. Current Request Method:"+r.Method)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Error("Couldn't decode request body. Error thrown:", err)
		writeResponse(w, false, "Couldn't decode request body. Error thrown:"+err.Error())
		return
	}
	// at this point we need to pass it over to the database instance to validate the request

	if result, err := a.ValidateUserLogin(requestBody.Email, requestBody.CreatePassword); err == nil {
		if result {
			log.Info("User successfully authenticated. Rerouting to terminal page")
			writeResponse(w, true, "")
			return
		}
	} else {
		log.Error("User information not found")
		log.Error(err)
		writeResponse(w, false, "User information not found")
		return
	}

	return
}

func (a *App) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {

	log.Info("Attempted Login... handling now")

	requestBody := &AuthRequestBody{}

	if r.Method != "POST" {
		log.Error("request method not aligned correctly for login function. Request Method:", r.Method)
		writeResponse(w, false, "request method not aligned correctly for login function. Current Request Method:"+r.Method)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Error("Couldn't decode request body. Error thrown:", err)
		writeResponse(w, false, "Couldn't decode request body. Error thrown:"+err.Error())
		return
	}

	if err := a.CreateUser(requestBody); err != nil {
		log.Error(err)
		writeResponse(w, false, err.Error())
		return
	}

	a.SetUserSession(requestBody.FirstName)

	log.Info("User successfully Created. Rerouting to terminal page")
	writeResponse(w, true, "")
	return
}

//##########################################################################################
//############################## Terminal Endpoints ########################################

type TerminalRequestBody struct {
	Command string
	IsBasic bool
}

// Endpoint used to do directory traversal.
func (a *App) HandleTerminalNav(w http.ResponseWriter, r *http.Request) {

	log.Info("Attempted to take terminal cli input... handling now")
	requestBody := &TerminalRequestBody{}

	if r.Method != "POST" {
		log.Error("request method not aligned correctly for terminal commands function. Request Method:", r.Method)
		writeResponse(w, false, "request method not aligned correctly for terminal commands function. Request Method: "+r.Method)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Error("Couldn't decode request body. Error thrown:", err)
		writeResponse(w, false, "Couldn't decode request body. Error thrown:"+err.Error())
		return
	}

	// at the point we need to execute the c++ binary and get the string output from it.
	s, err := a.TakeTerminalCommandLineInput(requestBody.Command)
	if err != nil {
		writeResponse(w, false, "Backend couldnt process the terminal request:"+err.Error())
	}

	writeResponse(w, true, s)
}