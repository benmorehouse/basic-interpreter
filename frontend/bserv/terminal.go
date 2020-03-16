// This should call on a package that i built earlier
package main

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// TakeTerminalCommandLineInput will take user input and run through c++ program
// this should be consolidated to return val: (struct{}, error)
func (a *App) TakeTerminalCommandLineInput(input string) (string, error) {

	// use filepath.Split
	// basicCmd.Dir = directory
	pathToBackend := a.Config.PathToBackend
	basicCmd := exec.Command("./basic", input)
	basicCmd.Dir = "../backend/bin"
	output, err := basicCmd.Output()

	log.Info(input)
	log.Info(pathToBackend)

	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Info(string(output))
	// now we need to handle the common output from the program
	return string(output), nil
}
