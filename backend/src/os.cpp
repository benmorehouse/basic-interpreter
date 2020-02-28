#include "../include/os.h"
#include "../include/logger.h"
#include <string>
#include <iostream>
#include <map>

enum CommandEnum {
	ls = 0,
	cd = 1,
	mkdir = 2,
	touch = 4,
	rm = 5,
	open = 6,
	pwd = 7,
	mv = 8,
	help = 111,
	ERROR = 999,
};

OperatingSystem::OperatingSystem() {
	this->Logger = new OperatingSystemLogger();
	this->InitializeCommandMap();
}

OperatingSystem::~OperatingSystem() {
	delete this->Logger;
}

// OperatingSystem should take in the arguments count and the
// argument vector and switch based on the argument.
void OperatingSystem::Operate(char **input, int len) {
	if (len < 2) {
		this->Logger->Error("No command found");
	}

	int commandEnum = this->CommandMap[input[1]];

	switch (commandEnum) {
		case ERROR: {
			std::string error(input[1]);
			this->Logger->Error("command not found: " + error);
			return;
	    	}
		   
		case ls: {
			ListCommand *listCommand = new ListCommand();	
			// Then show the list of things
		}

		case cd: {
			ChangeDirectoryCommand *changeDirectoryCommand = new ChangeDirectoryCommand();	
			// Then parse through and get the list of things.
		}

		case mkdir: {
			MakeDirectoryCommand *makeDirectoryCommand = new MakeDirectoryCommand();
	    	}

		case touch: {
			TouchCommand *touchCommand = new TouchCommand();
		}

		case rm: {
			RemoveCommand *removeCommand = new RemoveCommand();
		}

		case open: {
			OpenCommand *openCommand = new OpenCommand();
		}

		case pwd: {
			ProvideCommand *provideCommand = new ProvideCommand();
		}

		case mv: {
			MoveCommand *moveCommand = new MoveCommand();
		}

		case help: {
			HelpCommand *helpCommand = new HelpCommand();
		}
	}
}

void OperatingSystem::InitializeCommandMap() {
	this->CommandMap.insert(std::pair<std::string, int>("ls", ls));
	this->CommandMap.insert(std::pair<std::string, int>("cd", cd));
	this->CommandMap.insert(std::pair<std::string, int>("mkdir", mkdir));
	this->CommandMap.insert(std::pair<std::string, int>("touch", touch));
	this->CommandMap.insert(std::pair<std::string, int>("rm", rm));
	this->CommandMap.insert(std::pair<std::string, int>("open", open));
	this->CommandMap.insert(std::pair<std::string, int>("pwd", pwd));
	this->CommandMap.insert(std::pair<std::string, int>("help", pwd));
}

//########################################################
//###################### Commands ########################

Command::Command() {
	// Nothing to do yet but feel like that could change.
}

CommandResponse* Command::Process(char **command) {
	return nullptr;
}
	bool Success;
	std::string  ErrorMessage;
	std::string	Output;

void Command::HandleCommandOutput(CommandResponse* resp) {
	std::cout << "FINAL --------------------------------------" << std::endl;

	if (!resp->Success) {
		if (resp->ErrorMessage == "") {
			this->Logger->Error("There was an error reported but lost");
		} else {
			this->Logger->Error(resp->ErrorMessage);
		}
		return;
	}

	this->Logger->Info(resp->Output);
	return;
}

//###################### ls ########################

ListCommand::ListCommand() : Command() {}

CommandResponse* ListCommand::Process(char **command) {
	// Here we need to iterate through all of the directories in the operating system.
	return nullptr;
}

//###################### cd ########################

ChangeDirectoryCommand::ChangeDirectoryCommand() : Command() {}

CommandResponse* ChangeDirectoryCommand::Process(char **command) {
	return nullptr;
}

//###################### mkdir ########################

MakeDirectoryCommand::MakeDirectoryCommand() : Command() {}

CommandResponse* MakeDirectoryCommand::Process(char **command) {
	return nullptr;
}

//###################### touch ########################

TouchCommand::TouchCommand() : Command() {}

CommandResponse* TouchCommand::Process(char **command) {
	return nullptr;
}

//###################### rmdir #########################

RemoveCommand::RemoveCommand() : Command() {}

CommandResponse* RemoveCommand::Process(char **command) {
	return nullptr;
}

//###################### open  #########################

OpenCommand::OpenCommand() : Command() {}

CommandResponse* OpenCommand::Process(char **command) {
	return nullptr;
}

//###################### pwd #########################

ProvideCommand::ProvideCommand() : Command() {}

CommandResponse* ProvideCommand::Process(char **command) {
	return nullptr;
}

std::string ProvideCommand::ProvideHelper(Directory* dir) {
	if (dir == nullptr) {
		this->Logger->Error("Dir unexpectedly found as nil");
		return "";
	} else if (dir->isHome()) {
		return "/" + dir->getName();
	} else {
		return this->ProvideHelper(dir->getParent()) + dir->getName() + "/";
	}
}

//###################### mv #########################

MoveCommand::MoveCommand() : Command() {}

CommandResponse* MoveCommand::Process(char **command) {
	return nullptr;
}

//###################### help #########################

HelpCommand::HelpCommand() : Command() {}

CommandResponse* HelpCommand::Process(char **command) {
	return nullptr;
}


