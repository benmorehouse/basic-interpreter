#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <cstddef>        // std::size_t
#include "../include/directory.h"

Directory::Directory(std::string name, Directory* parent) {
	this->Name = name;
	this->IsDir = true; // postulate: assume it is a dir, prove otherwise

	if (parent == nullptr) {
		this->isHomeDir = true;
		return;	
	}
	
	this->setParent(parent);
	this->isHomeDir = false;

	std::string fileExtension = " ";
	std::string::iterator it = name.end();
	
	while(it != name.begin()) {
		if(*it == '.'){
			this->IsDir = false;
			std::string str(1, *it);
			fileExtension.insert(0, str);
			break;
		}
		std::string str(1, *it);
		fileExtension.insert(0, str);
		--it;	
	}
}	

//############################ STANDARD GETTERS ################################

std::string Directory::getName() {
	return this->Name;
}

Directory* Directory::getElementFromDirectoryMap(std::string nameOfElement) {
	if (nameOfElement == "") {
		return nullptr;
	}
	
	std::map<std::string, Directory*>::iterator exists;
	exists = this->directories.find(nameOfElement);
	if (exists == this->directories.end()) {
		return nullptr;	
	} 

	return exists->second;
}

File* Directory::getFile(std::string nameOfFile) {
	return this->file;
}

Directory* Directory::getDirectory(std::string nameOfDir) {
	Directory* result = this->getElementFromDirectoryMap(nameOfDir);
	if (result == nullptr) {
		return nullptr;	
	} else if (!result->isDirectory()) {
		return nullptr;
	} 	

	return result;
}

Directory* Directory::getParent() {
	return this->parent;
}

//############################ STANDARD SETTERS ######################################

void Directory::addDirectory(Directory* dir) {
	if (!dir->isDirectory()) {
		return;		
	} else if (dir->getName() == "") {
		return;	
	}
	
	this->directories[dir->getName()] = dir;
}

void Directory::addFile(Directory* dir) {
	if (dir->isDirectory()) {
		return;		
	} else if (dir->getName() == "") {
		return;	
	}
	
	this->directories[dir->getName()] = dir;
}

void Directory::setName(std::string name) {
	this->Name = name;
}

void Directory::setParent(Directory* parent) {
	if(parent != nullptr) {
		this->parent = parent;
	}
}

bool Directory::isHome() {
	return this->isHomeDir;
}

bool Directory::isDirectory() {
	return this->IsDir;
}

//############################ COMMAND HELPERS ######################################

// Iterates through subdirectory classes and returns array of names.
std::vector<std::string>* Directory::getAllSubAsName() {
	std::vector<std::string> names;
	std::map<std::string, Directory*>::iterator it;

	for (it = this->directories.begin(); it != this->directories.end(); ++it) {
		Directory* dir = it->second; 
		names.push_back(dir->getName());
	}

	return &names;
}


// just something to add in for right now
