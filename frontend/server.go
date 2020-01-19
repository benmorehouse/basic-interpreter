package main

import(
	"html/template"
	"net/http"
)

var log = NewLogger()

// should make an init function

type Page struct{
	Title         string
	User          *Session
}

type Session struct{
	LoggedIn    bool
	Username    string
}

func main(){
	log.Error("Basic Interpreter has started...")
	basicTemplate := template.Must(template.ParseFiles("about.gohtml"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		err := basicTemplate.Execute(w, nil)
		if err != nil{
			log.Error(err)
			panic(err)
		}
	})

	http.HandleFunc("/terminal", func(w http.ResponseWriter, r *http.Request){
		if err := basicTemplate.Execute(w, nil); err != nil{
			log.Error(err)
			panic(err)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){
		if err := basicTemplate.Execute(w, nil); err != nil{
			log.Error(err)
			panic(err)
		}
	})

	http.HandleFunc("/github", func(w http.ResponseWriter, r *http.Request){
		basicTemplate.Execute(w, nil)
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		log.Error(err)
	}

	log.Info("Basic Interpreter Is Waiting...")
}
