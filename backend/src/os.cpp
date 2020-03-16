#include "../include/os.h"
#include "../include/logger.h"
#include "../include/interpreter.h"
#include <string>
#include <iostream>
#include <fstream>
#include <map>

//##################################################
//######## Enums representing each command #########

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
	compile = 123,
	ERROR = 999,
};

OperatingSystem::OperatingSystem() {
	this->Logger = new OperatingSystemLogger();
	this->InitializeCommandMap();
}

OperatingSystem::~OperatingSystem() {
	delete this->Logger;
}

// HandleCommandOutput is used to take a command response struct
// And force wrap it to a certain output type
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
			
			CommandResponse* response;
			if (len == 2) {
				this->Logger->Info("go get the reponse!");
				response = changeDirectoryCommand->process(this->currentDirectory, "");// hoping to pass in this address and have it change the address value.
				this->Logger->Info("end of response getter");
			} else {
				response = changeDirectoryCommand->process(this->currentDirectory, input[2]);
				this->Logger->Info("hit the else clause");
			}
			this->Logger->Info("Here we are at the end of the cd command");
			HandleCommandOutput(response);
			delete changeDirectoryCommand;
			return;
		}

		case mkdir: {
			MakeDirectoryCommand *makeDirectoryCommand = new MakeDirectoryCommand();
			CommandResponse* response;
			if (len == 2) {
				this->Logger->Error("User asked to make a directory but didn't give us the name of the directory!?");
				this->Logger->Error("mdkir [<name>] ...");
				response->success = false;
				response->errorMessage = "mdkir [<name>] ...";
				return;
			} else {
				response = makeDirectoryCommand->process(this->currentDirectory, input[2]);
				return;
			}

			delete makeDirectoryCommand;
			return;
	    	}

		case touch: {
			TouchCommand *touchCommand = new TouchCommand();
			return;
		}

		case rm: {
			RemoveCommand *removeCommand = new RemoveCommand();
			return;
		}

		case open: {
			OpenCommand *openCommand = new OpenCommand();
			return;
		}

		case pwd: {
			this->Logger->Info("Recieved a pwd");	
			ProvideCommand* provideCommand = new ProvideCommand();
			CommandResponse* response = provideCommand->process(nullptr);
			HandleCommandOutput(response);
			delete provideCommand;
			return;
		}

		case mv: {
			MoveCommand *moveCommand = new MoveCommand();
		}

		case help: {
			HelpCommand *helpCommand = new HelpCommand();
			CommandResponse* response = helpCommand->process();
			HandleCommandOutput(response);
			delete helpCommand;
		}
		
		case compile: {
			CompileCommand *compileCommand = new CompileCommand();	
			CommandResponse* response;
			if (len == 2) {
				this->Logger->Error("Wasn't given a file to compile.");
				response->success = false;
				response->errorMessage = "Wasn't given a file to compile."; 
				delete compileCommand;
			}

			response = compileCommand->process(input[2]);
			HandleCommandOutput(response);
			delete compileCommand;
	  	}
	}
}

// InitializeCommandMap is used to create a map for all command enums
void OperatingSystem::InitializeCommandMap() {
	this->CommandMap.insert(std::pair<std::string, int>("ls", ls));
	this->CommandMap.insert(std::pair<std::string, int>("cd", cd));
	this->CommandMap.insert(std::pair<std::string, int>("mkdir", mkdir));
	this->CommandMap.insert(std::pair<std::string, int>("touch", touch));
	this->CommandMap.insert(std::pair<std::string, int>("rm", rm));
	this->CommandMap.insert(std::pair<std::string, int>("open", open));
	this->CommandMap.insert(std::pair<std::string, int>("pwd", pwd));
	this->CommandMap.insert(std::pair<std::string, int>("help", help));
	this->CommandMap.insert(std::pair<std::string, int>("compile", help));
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
	CommandResponse* response = new CommandResponse();
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

CommandResponse* ChangeDirectoryCommand::process(Directory* dir, std::string newDir) {
	CommandResponse* response = new CommandResponse();
	this->Logger->Info("1");
	if (dir == nullptr) {
		this->Logger->Info("2");
		response->success = false;
		response->errorMessage = "Passed in a nil directory.";
		return response;
	} else if (!dir->isDirectory()) {
		this->Logger->Info("3");
		response->success = false;
		response->errorMessage = "This is apparently of type file?";
		this->Logger->Info("5");
		return response;
	}
	
	if (newDir == "") {
		// Then trickle up the tree to the home directory.
		this->trickleUpToHome(dir);
		response->success = true;
		response->output = "gone home!";
		return response;
	}
		
	this->Logger->Info("6");

	Directory* changeToThisDirectory = dir->getDirectory(newDir);
	if (changeToThisDirectory == nullptr) {
		response->success = false;
		response->errorMessage = "This is apparently of type file?";
		return response;
	}

	// Now we need to go in and make sure that the directory actually exists in order to switch to it!
	// Also must make sure that it is a directory before we change to it.
	dir = changeToThisDirectory;
	response->success = true;
	return response;
}

void ChangeDirectoryCommand::trickleUpToHome(Directory* dir) {
	if (dir == nullptr) {
		this->Logger->Error("The directory was found nil whilst we were going home");
		this->Logger->Error("Help i am experiencing homelessness :(");
		return;
	} else {
		while (!dir->isHome()) {
			dir = dir->getParent();
			if (dir == nullptr) {
				this->Logger->Error("The directory was found nil whilst we were going home");
				this->Logger->Error("Help i am experiencing homelessness :(");
				return;
			}
		}
	}
}

//###################### mkdir ########################

MakeDirectoryCommand::MakeDirectoryCommand() : Command() {}

CommandResponse* MakeDirectoryCommand::process(Directory* dir, std::string newDir) {
	
	CommandResponse* response = new CommandResponse();
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "this doesnt seem to be a directory.";
		return response;
		return nullptr;
	} else if (newDir == "") {
		response->success = false;
		response->errorMessage = "Passed in a new directory name that is blank!";
		return response;
	}
	
	Directory* newDirectory = new Directory(newDir, dir);
	if (!newDirectory->isDirectory()) {
		response->success = false;
		response->errorMessage = "This is a file not a directory!";
		delete newDirectory;
		return response;	
	}
	
	dir->addDirectory(newDirectory);
	response->success = true;
	return response;
}

//###################### touch ########################

TouchCommand::TouchCommand() : Command() {}

CommandResponse* TouchCommand::process(Directory* dir, std::string newFile) {
	CommandResponse* response;
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "this doesnt seem to be a directory.";
		return response;
		return nullptr;
	} else if (newFile == "") {
		response->success = false;
		response->errorMessage = "Passed in a new directory name that is blank!";
		return response;
	}
	
	Directory* newDirectory = new Directory(newFile, dir);
	if (!newDirectory->isDirectory()) {
		response->success = false;
		response->errorMessage = "This is a file not a directory!";
		delete newDirectory;
		return response;	
	}

	dir->addDirectory(newDirectory);
	response->success = true;
	return response;
}

//###################### rmdir #########################

RemoveCommand::RemoveCommand() : Command() {}

CommandResponse* RemoveCommand::process(Directory* dir, std::string removeFile) {
	CommandResponse* response = new CommandResponse();
	if (dir == nullptr) {
		response->success = false;
		response->errorMessage = "this doesnt seem to be a directory.";
		return response;
		return nullptr;
	} else if (removeFile == "") {
		response->success = false;
		response->errorMessage = "Passed in a new directory name that is blank!";
		return response;
	} 

	Directory *checkIfExists = dir->getDirectory(removeFile);
	if (checkIfExists == nullptr) {
		response->success = false;
		response->errorMessage = "Tried to delete a directory that wasnt there.";
		return response;
	}
	
	if (checkIfExists->isDirectory()) {
		dir->removeDirectory(checkIfExists);
		response->success = true;
		return response;
	} else {
		dir->removeFile(checkIfExists);
		response->success = true;
		return response;
	}
	
	response->success = true;
	return response;
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
	CommandResponse* response = new CommandResponse();
	if (currentDirectory == nullptr) {
		response->success = false;
		response->errorMessage = "Directory found as nil";	
	}

	std::string pwdResult = this->ProvideHelper(currentDirectory);

	if (pwdResult == "") {
		response->success = false;
		response->errorMessage = "provide working directory command not functioning as expected...";
	} else {
		response->success = true;
		response->output = pwdResult;
	}

	return response;
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
// Should be able to move a file based on where it is.

MoveCommand::MoveCommand() : Command() {}

CommandResponse* MoveCommand::process(Directory* dir) {
	return nullptr;
}

//###################### help #########################

HelpCommand::HelpCommand() : Command() {}

CommandResponse* HelpCommand::process() {
	CommandResponse *response = new CommandResponse();	
	response->success = true;
	std::string s = "Hello and welcome to my basic interpreter!\n";
	s += "The commands are:\n\n";
	s += "ls 		list all directories and files within your current directory\n";
	s += "cd <dir>		change directories\n";
	s += "mkdir <dir>	make a directory\n";
	s += "touch <dir>	make a file\n";
	s += "rm <dir> 		remove a directory or file \n";
	s += "open <file>	open a file\n";
	s += "pwd		provide the working directory\n";
	s += "mv 		move a file or directory to some other directory\n";
	s += "help		display this menu\n";
	s += "compile <file>	compile the basic file given\n";
	response->output = s;
	return response;
}

//###################### compile #########################

CompileCommand::CompileCommand() : Command() {}

CommandResponse* CompileCommand::process(std::string filename) {
	// this needs to create an interreter instance.
	CommandResponse *response = new CommandResponse();	
	
	std::ifstream ifile;
	std::string filePath = "../fqueue/" + filename;
	ifile.open(filePath);
	if (ifile.fail()) {
		this->Logger->Error("An error occurred when we tried to open the file in the file queue.");
		response->success = false;
		response->errorMessage = "An error occurred when we tried to open the file in the file queue.";
		return response;
	}
	
	Interpreter *interpreter = new Interpreter(ifile);
	return response;
}

