package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var root = &cobra.Command{
		Use:   "", // will run everytime you type nothing in
		Short: "A Full Stack Web App for an Interpreter of the Basic Programming Language",
		Run: func(cmd *cobra.Command, args []string) {
			basicInterpreter()
		},
	}

	var init = &cobra.Command{
		Use:   "init", // will run everytime you type nothing in
		Short: "Initialize your basic interpreter server",
		//Run:
	}

	root.AddCommand(init)
	root.Flags().String("conf", "", "configuration file")
	if err := root.Execute(); err != nil {
		log.Error(err)
	}
}
