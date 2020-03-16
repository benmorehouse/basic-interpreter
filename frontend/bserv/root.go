package main

import (
	"io/ioutil"

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

		i := InitOptions{
			IsVerbose: isVerbose,
			IsInit:    isInit,
			Config:    conf,
		}

		if err := StartServer(i); err != nil {
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
