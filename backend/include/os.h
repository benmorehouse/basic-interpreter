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
		// i think this needs to be moved to the directories for each directory
		Directory* currentDirectory;
};

struct CommandResponse{
	bool success;
	std::string errorMessage;
	std::string output;
};

class Command {
	public:
		Command();
		~Command();
		OperatingSystemLogger *Logger;
		virtual CommandResponse* process(Directory*);

};

class ListCommand : Command {
	public:
		ListCommand();
		CommandResponse* process(Directory*);
};

class ChangeDirectoryCommand : Command {
	public:
		ChangeDirectoryCommand();	
		CommandResponse* process(Directory*, std::string);
	private:
		void trickleUpToHome(Directory*);
};

class MakeDirectoryCommand : Command {
	public:
		MakeDirectoryCommand();
		CommandResponse* process(Directory*, std::string);
};

class TouchCommand : Command {
	public:
		TouchCommand();
		CommandResponse* process(Directory*, std::string);
};

class RemoveCommand : Command {
	public:
		RemoveCommand();
		CommandResponse* process(Directory*, std::string);
};

class OpenCommand : Command {
	public:
		OpenCommand();
		CommandResponse* process(Directory*);
};

class ProvideCommand : Command {
	public:
		ProvideCommand();
		CommandResponse* process(Directory*);
	private: 
		std::string ProvideHelper(Directory *);	
};

class MoveCommand : Command {
	public:
		MoveCommand();
		CommandResponse* process(Directory*);
};

class HelpCommand : Command {
	public:
		HelpCommand();
		CommandResponse* process();
};

class CompileCommand : Command {
	public:
		CompileCommand();
		CommandResponse* process(std::string);
};

#endif
