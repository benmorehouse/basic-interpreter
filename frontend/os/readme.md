# BOS - Basic operating system

This serves as the package which offers the operating system for the command line prompt 
served out of bserv

Here is the structure of bos: 

- OperatingSystem: The general instance that is called by the program at the start. 
	Keeps track of all possible commands, as well as the current directory.

- Commands: commands are called by the operating system. There are several commands, each of which maps to a possible user input.

- CommandResponse: the response to the command that is taken from the user. 


Commands


