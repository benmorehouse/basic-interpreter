#ifndef OS_H
#define OS_H

#include "directory.h"
#include "logger.h"
#include <string>
#include <map>

class OperatingSystem {
	public:
		OperatingSystem();
		~OperatingSystem();
		void Operate(char **, int);
	private:
		// Logger	
		OperatingSystemLogger *Logger;
		
		// Functions and Commands for determining the initial Command	
		void InitializeCommandMap();
		std::map<std::string, int> CommandMap;
		// a map that for each name of the directory there corresponds a directory.
		std::map<std::string, Directory*> Directories;
};

struct CommandResponse{
	bool Success;
	std::string  ErrorMessage;
	std::string	Output;
};

class Command {
	public:
		Command();
		OperatingSystemLogger *Logger;
		void HandleCommandOutput(CommandResponse*);
		virtual CommandResponse* Process(char**);
};

class ListCommand : Command {
	public:
		ListCommand();
		CommandResponse* Process(char **);
};

class ChangeDirectoryCommand : Command {
	public:
		ChangeDirectoryCommand();	
		CommandResponse* Process(char **);
};

class MakeDirectoryCommand : Command {
	public:
		MakeDirectoryCommand();
		CommandResponse* Process(char **);
};

class TouchCommand : Command {
	public:
		TouchCommand();
		CommandResponse* Process(char **);
};

class RemoveCommand : Command {
	public:
		RemoveCommand();
		CommandResponse* Process(char **);
};

class OpenCommand : Command {
	public:
		OpenCommand();
		CommandResponse* Process(char **);
};

class ProvideCommand : Command {
	public:
		ProvideCommand();
		CommandResponse* Process(char **);
	private: 
		void ProvideHelper(Directory *);	
};

class MoveCommand : Command {
	public:
		MoveCommand();
		CommandResponse* Process(char **);
};

class HelpCommand : Command {
	public:
		HelpCommand();
		CommandResponse* Process(char **);
};

#endif
