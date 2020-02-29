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
			this->Logger->Info("Recieved a pwd");	
			ProvideCommand *provideCommand = new ProvideCommand();
			CommandResponse *cr = provideCommand->Process(nullptr);
			delete provideCommand;
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
	this->Logger = new OperatingSystemLogger();
}

Command::~Command() {
	delete this->Logger;
}

CommandResponse* Command::Process(char **command) {
	this->Logger->Error("wrong process fetched");
	return nullptr;
}

void Command::HandleCommandOutput(CommandResponse* resp) {
	std::cout << "FINAL --------------------------------------" << std::endl;

	if (!resp->success) {
		if (resp->errorMessage == "") {
			this->Logger->Error("There was an error reported but lost");
		} else {
			this->Logger->Error(resp->errorMessage);
		}
		return;
	}

	this->Logger->Info(resp->output);
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

ProvideCommand::ProvideCommand() : Command() {
	this->Logger->Info("initialized a provide working directory command command");
}

CommandResponse* ProvideCommand::Process(Directory* currentDirectory) {
	this->Logger->Info("we got to the root of the pwd command");
	std::string pwdResult = this->ProvideHelper(currentDirectory);
	CommandResponse* cr = new CommandResponse();

	if (pwdResult == "") {
		cr->success = false;
		cr->errorMessage = "provide working directory command not functioning as expected...";
	} else {
		cr->success = true;
		cr->output = pwdResult;
	}

	return cr;
}

std::string ProvideCommand::ProvideHelper(Directory* dir) {
	if (dir == nullptr) {
		this->Logger->Error("Dir unexpectedly found as nil");
		return "";
	} else if (dir->isHome()) {
		return "/" + dir->getName() + "/";
	} else {
		std::string child = this->ProvideHelper(dir->getParent());
		if (child != "") {
			return child + dir->getName() + "/";
		} else {
			return child;
		}
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


