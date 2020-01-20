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

	http.HandleFunc("/", a.HandleAbout)
	http.HandleFunc("/terminal", a.HandleTerminal)
	http.HandleFunc("/login", a.HandleLogin)
	http.HandleFunc("/loginAttempt", a.HandleLoginAttempt)
	http.HandleFunc("/github", a.HandleGithub)
	port := ":" + strconv.Itoa(a.ConfigFile.Port)
	log.Info("Basic Interpreter Is Waiting...")
	log.Info("LOCAL: http://localhost" + port)
	err = http.ListenAndServe(port, nil)
	if err != nil{
		log.Error(err)
	}
}
