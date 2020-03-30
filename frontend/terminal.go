package main

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
