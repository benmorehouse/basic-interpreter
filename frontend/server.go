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

// StartServer will start the server
func StartServer(i InitOptions) error {

	setLogger(i.IsVerbose)

	log.Info("Basic Interpreter Server started")
	a, err := NewApp(i.Config, i.IsInit)
	if err != nil {
		err := ServerError(CreateAppFailed, nil)
		log.Error(err)
		log.Error("Ending Server lifespan...")
		return err
	}

	log.Info("App successfully intialized")

	router := http.NewServeMux()
	router.HandleFunc(a.Config.AboutPageURL, a.HandleAbout)
	router.HandleFunc(a.Config.TerminalPageURL, a.HandleTerminal)
	router.HandleFunc(a.Config.GithubPageURL, a.HandleGithub)
	router.HandleFunc(a.Config.TextEditorPageURL, a.HandleTextEditor)

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

	// server the javascript files for the frontend
	router.Handle(a.Config.ScriptsPrefix,
		http.StripPrefix(a.Config.ScriptsPrefix, http.FileServer(http.Dir(a.Config.PathToScripts))))

	// server the css files for the frontend
	router.Handle(a.Config.CSSPrefix,
		http.StripPrefix(a.Config.CSSPrefix, http.FileServer(http.Dir(a.Config.PathToCSS))))

	router.Handle("/pages/",
		http.StripPrefix("/pages/", http.FileServer(http.Dir("pages/"))))

	port := ":" + strconv.Itoa(a.Config.Port) // port is simply used to display the logging message!!

	go a.operatingSystem.RunOperatingSystem()
	defer func() {
		a.operatingSystem.DonePipe <- true
	}()

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

// HandleTextEditor renders the text editor file
func (a *App) HandleTextEditor(w http.ResponseWriter, r *http.Request) {

	log.Info("Text Editor Page requested")

	type TemplateWriter struct {
		TerminalPageRedirect string `json:"TerminalPageRedirect"`
		CompileEndpoint      string `json:"CompileEndpoint"`
		SaveEndpoint         string `json:"SaveEndpoint"`
		File                 *File  `json:"File"`
	}

	write := func(success bool, msg string, writer *TemplateWriter) {
		Template := struct {
			Message string          `json:"Message"`
			Success bool            `json:"Success"`
			Body    *TemplateWriter `json:"Body"`
		}{
			msg,
			success,
			writer,
		}

		if !success {
			log.Error(msg)
		}

		basicTemplate := template.Must(template.ParseFiles(a.Config.TextEditorFile))
		err := basicTemplate.Execute(w, Template)
		if err != nil {
			log.Error(err)
		}
	}

	Template := &TemplateWriter{
		TerminalPageRedirect: a.Config.TerminalPageURL,
		CompileEndpoint:      a.Config.CompileEndpoint,
		SaveEndpoint:         a.Config.SaveFileEndpoint,
	}

	// execute the template with something different.
	fileHashField, exists := r.URL.Query()["hash"]
	if !exists || len(fileHashField) == 0 {
		write(false, "No file found", Template)
		return
	}

	fileHash := fileHashField[0]
	file, err := a.connection.GetFileFromFilestore([]byte(fileHash), a.User.Username)
	if err != nil {
		write(false, err.Error(), Template)
		return
	}

	Template.File = file
	write(true, "", Template)
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

// HandleCreateAccount will handle creating an account
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

// TerminalRequestBody is the request body for terminal page requests
type TerminalRequestBody struct {
	Command string
	IsBasic bool
}

// Endpoint used to do directory traversal.
func (a *App) HandleTerminalNav(w http.ResponseWriter, r *http.Request) {

	respond := func(success bool, output, currentDirectory string) {
		responseBody := struct {
			Success          bool   `json:"Success"`
			Messsage         string `json:"Message"`
			CurrentDirectory string `json:"CurrentDirectory"`
		}{
			success,
			output,
			currentDirectory,
		}

		writeThisResponse, err := json.Marshal(responseBody)
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

	requestBody := &TerminalRequestBody{}

	if r.Method != "POST" {
		output := "request method not aligned correctly for terminal commands function. Request Method:" + r.Method
		log.Error(output)
		respond(false, output, "")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Error("Couldn't decode request body. Error thrown:", err)
		return
	}

	s := a.TakeTerminalCommandLineInput(requestBody.Command)
	respond(true, s, "")
}

// TakeTerminalCommandLineInput will take user input and run through c++ program
// this should be consolidated to return val: (struct{}, error)
// NOTE: this will interact with the OS by feeding data through its pipes.
func (a *App) TakeTerminalCommandLineInput(input string) string {

	a.operatingSystem.CommandPipe <- input
	output := <-a.operatingSystem.ResponsePipe
	if !output.Success {
		return "Error: " + output.Output
	}

	return output.Output
}

func (a *App) GetCurrentDirectory() string {

	return a.TakeTerminalCommandLineInput("pwd")
}
