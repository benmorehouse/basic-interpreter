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

void HandleCommandOutput(CommandResponse* resp) {
	std::cout << "FINAL --------------------------------------" << std::endl;
	
	OperatingSystemLogger* Logger = new OperatingSystemLogger();

	if (!resp->success) {
		if (resp->errorMessage == "") {
			Logger->Error("There was an error reported but lost");
		} else {
			Logger->Error(resp->errorMessage);
		}
		return;
	}

	Logger->Info(resp->output);
	return;
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
			CommandResponse* response = listCommand->process(this->currentDirectory);
			HandleCommandOutput(response);
			delete listCommand;
			// Then show the list of things
		}

		case cd: {
			ChangeDirectoryCommand *changeDirectoryCommand = new ChangeDirectoryCommand();
			CommandResponse* response = changeDirectoryCommand->process(this->currentDirectory);// hoping to pass in this address and have it change the address value.
			HandleCommandOutput(response);
			delete changeDirectoryCommand;

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
			ProvideCommand* provideCommand = new ProvideCommand();
			CommandResponse* response = provideCommand->process(nullptr);
			HandleCommandOutput(response);
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

CommandResponse* Command::process(Directory *dir) {
	this->Logger->Error("wrong process fetched");
	return nullptr;
}

//###################### ls ########################

ListCommand::ListCommand() : Command() {}

CommandResponse* ListCommand::process(Directory* dir) {
	// Here we need to iterate through all of the directories in the operating system.
	CommandResponse* response;
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "Passed in a nil directory.";
		return response;
	} 
	
	std::vector<std::string>* listOfDirectories = dir->getAllSubAsName();	
	std::vector<std::string>::iterator it;
	std::string directories = "";

	for(it = listOfDirectories->begin(); it != listOfDirectories->end(); ++it) {
		if(*it != "") {
			directories += *it + "\n";
		}
	}
	
	response->success = true;
	response->output = directories;
	return response;
}

//###################### cd ########################

ChangeDirectoryCommand::ChangeDirectoryCommand() : Command() {}

CommandResponse* ChangeDirectoryCommand::process(Directory* dir) { // meaning we are tracing back up to home directory
	CommandResponse* response;
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "Passed in a nil directory.";
		return response;
	}

	// iterate back up to things.
	return nullptr;
}

CommandResponse* ChangeDirectoryCommand::process(Directory* dir, std::string newDir) {
	CommandResponse* response;
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "Passed in a nil directory.";
		return response;
	} else if (newDir == "") {
		response->success = false;
		response->errorMessage = "Passed in a nil directory name.";
		return response;
	}

	// now we need to go in and get the 
	return nullptr;
}

//###################### mkdir ########################

MakeDirectoryCommand::MakeDirectoryCommand() : Command() {}

CommandResponse* MakeDirectoryCommand::process(Directory* dir) {
	return nullptr;
}

//###################### touch ########################

TouchCommand::TouchCommand() : Command() {}

CommandResponse* TouchCommand::process(Directory* dir) {
	return nullptr;
}

//###################### rmdir #########################

RemoveCommand::RemoveCommand() : Command() {}

CommandResponse* RemoveCommand::process(Directory* dir) {
	return nullptr;
}

//###################### open  #########################

OpenCommand::OpenCommand() : Command() {}

CommandResponse* OpenCommand::process(Directory* dir) {
	return nullptr;
}

//###################### pwd #########################

ProvideCommand::ProvideCommand() : Command() {
	this->Logger->Info("initialized a provide working directory command command");
}

CommandResponse* ProvideCommand::process(Directory* currentDirectory) {
	this->Logger->Info("we got to the root of the pwd command");
	CommandResponse* cr = new CommandResponse();
	if (currentDirectory == nullptr) {
		cr->success = false;
		cr->errorMessage = "Directory found as nil";	
	}

	std::string pwdResult = this->ProvideHelper(currentDirectory);

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

CommandResponse* MoveCommand::process(Directory* dir) {
	return nullptr;
}

//###################### help #########################

HelpCommand::HelpCommand() : Command() {}

CommandResponse* HelpCommand::process(Directory* dir) {
	return nullptr;
}


