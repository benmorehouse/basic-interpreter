package main

import (
	"encoding/json"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//##########################################################################################
//############################## Page  Endpoints ###########################################

func (a *App) RedirectIndex(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "localhost:2272/about", 301)
}

func (a *App) HandleAbout(w http.ResponseWriter, r *http.Request) {

	log.Info("About Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.AboutPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil {
		log.Error(err)
	}
}

func (a *App) HandleLogin(w http.ResponseWriter, r *http.Request) {

	log.Info("Login Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.LoginPageFile))
	err := basicTemplate.Execute(w, a)
	if err != nil {
		log.Error(err)
	}
}

func (a *App) HandleGithub(w http.ResponseWriter, r *http.Request) {

	log.Info("Github Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.GithubPageFile))
	err := basicTemplate.Execute(w, nil)
	if err != nil {
		log.Error(err)
	}
}

func (a *App) HandleTerminal(w http.ResponseWriter, r *http.Request) {

	log.Info("Terminal Page requested")
	basicTemplate := template.Must(template.ParseFiles(a.Config.TerminalPageFile))
	err := basicTemplate.Execute(w, nil)
	if err != nil {
		log.Error(err)
	}
}

//##########################################################################################
//############################## User Auth Endpoints #######################################

type AuthRequestBody struct {
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	CreatePassword  string `json:"CreatePassword"`
	ConfirmPassword string `json:"ConfirmPassword"`
	Email           string `json:"Email"`
}

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
	Command    string
	IsBasic    bool // Whether or not we are compiling basic code
	FileSelect string
}

// Endpoint used to do directory traversal.
func (a *App) HandleTerminalNav(w http.ResponseWriter, r *http.Request) {

	log.Info("Attempted Login... handling now")
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
}

// Endpoint used to call the interpretation of some code.
func (a *App) HandleBasicInterpretation(w http.ResponseWriter, r *http.Request) {

}
