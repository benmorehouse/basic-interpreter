package main

import(
	"net/http"
	"strconv"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.Info("Basic Interpreter Server started")
	a, err := NewApp()
	log.Info("App successfully intialized")

	http.HandleFunc(a.Config.AboutPageURL, a.HandleAbout)
	http.HandleFunc(a.Config.TerminalPageURL, a.HandleTerminal)
	http.HandleFunc(a.Config.LoginPageURL, a.HandleLogin)
	http.HandleFunc(a.Config.GithubPageURL, a.HandleGithub)

	port := ":" + strconv.Itoa(a.Config.Port)
	log.Info("Basic Interpreter Is Waiting...")
	log.Info("LOCAL: http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil{
		log.Error(err)
	}
}
