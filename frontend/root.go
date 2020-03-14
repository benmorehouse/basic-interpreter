package main

import (
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:   "", // will run everytime you type nothing in
	Short: "A server hosting an Interpreter of the Basic Programming Language, a Linux command line prompt, and a few other features.",
	RunE: func(cmd *cobra.Command, args []string) error {
		isVerbose := viper.GetBool("verbose")
		isInit := viper.GetBool("init")
		conf := viper.GetString("conf")
		isHelp := viper.GetBool("help")

		if isHelp {
			return cmd.Help()
		}

		setLogger(isVerbose)

		log.Info("Basic Interpreter Server started")
		a, err := NewApp(conf, isInit)
		if err != nil {
			log.Error(err)
			log.Error("Ending Server lifespan...")
			return err
		}

		//mux := http.NewServeMux() look into how a server mux works!
		log.Info("App successfully intialized")
		http.Handle("/", http.FileServer(http.Dir("/pages/scripts")))
		http.HandleFunc(a.Config.AboutPageURL, a.HandleAbout)
		http.HandleFunc(a.Config.TerminalPageURL, a.HandleTerminal)

		//mux.Handle(a.Config.TerminalPageURL, http.FileServer(http.Dir("/pages/scripts")))
		http.HandleFunc(a.Config.GithubPageURL, a.HandleGithub)

		// login and sign up handlers
		http.HandleFunc(a.Config.LoginPageURL, a.HandleLogin)
		http.HandleFunc(a.Config.CreateAccountURL, a.HandleCreateAccount)
		http.HandleFunc(a.Config.LoginAttemptedPageURL, a.HandleLoginAttempt)
		// login and sign up handlers

		port := ":" + strconv.Itoa(a.Config.Port) // port is simply used to display the logging message!!
		log.Info("Basic Interpreter Is Waiting...")
		log.Info("LOCAL: http://localhost" + port)
		err = http.ListenAndServe(port, nil)
		if err != nil {
			log.Error(err)
			return err
		}
		return nil
	},
}

func init() {

	// adding in persisten flags for the server
	cmd.PersistentFlags().Bool("init", false, "Set to initialize an empty database instance")
	cmd.PersistentFlags().String("conf", "conf.json", "Set to intialize instance with custom config file")
	cmd.PersistentFlags().Bool("verbose", false, "Set to false to turn off console logging.")
	cmd.PersistentFlags().Bool("help", false, "A help function.")

	// bind the flags to viper to be used by the cli and accessed by the cobra command
	viper.BindPFlag("init", cmd.PersistentFlags().Lookup("init"))
	viper.BindPFlag("conf", cmd.PersistentFlags().Lookup("conf"))
	viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("help", cmd.PersistentFlags().Lookup("help"))
}

func setLogger(verbose bool) {
	// could expand on this function later on.
	// can send this to ioutil.Discard when verbose is true
	if verbose {
		log.SetOutput(ioutil.Discard)
	}

	log.SetReportCaller(true)
}
