#include "../include/os.h"
#include "../include/logger.h"
#include <string>
#include <map>

enum CommandEnum {
	ls,
	cd,
	mkdir,
	rmdir,
	touch,
	rm,
	open,
	pwd,
	ERROR,
};

OperatingSystem::OperatingSystem() {
	this->Logger = new OperatingSystemLogger();
	this->InitializeCommandMap();
	
}

OperatingSystem::~OperatingSystem() {
	delete this->Logger;
}

void OperatingSystem::Operate(std::string input) {
	int command = this->DetermineCommand(input); // Need to make it all lower case and trim to only get the first word.
	this->Logger->Info(std::to_string(command));
}

void OperatingSystem::InitializeCommandMap() {
	this->CommandMap.insert(std::pair<std::string, int>("ls", ls));
	this->CommandMap.insert(std::pair<std::string, int>("cd", ls));
	this->CommandMap.insert(std::pair<std::string, int>("mkdir", ls));
	this->CommandMap.insert(std::pair<std::string, int>("rmdir", ls));
	this->CommandMap.insert(std::pair<std::string, int>("touch", ls));
	this->CommandMap.insert(std::pair<std::string, int>("rm", ls));
	this->CommandMap.insert(std::pair<std::string, int>("open", ls));
	this->CommandMap.insert(std::pair<std::string, int>("pwd", ls));
}


