#include <iostream>
#include <string>
#include <cstddef>        // std::size_t
#include "../include/directory.h"

Directory::Directory(std::string name, Directory* parent) {
	this->Name = name;
	this->IsDir = true; // postulate: assume it is a dir, prove otherwise
	this->dir = nullptr;
	this->file = nullptr;

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

void Directory::setName(std::string name) {
	this->Name = name;
}

void Directory::setParent(Directory* parent) {
	if(parent != nullptr) {
		this->parent = parent;
	}
}

std::string Directory::getName() {
	return this->Name;
}

bool Directory::isDirectory() {
	return this->IsDir;
}

File* Directory::getFile() {
	if(!this->isDirectory()) {
		return this->file;
	}
	return nullptr;
}

Directory* Directory::getDirectory() {
	if(this->isDirectory()) {
		return this->dir;
	}
	return nullptr;
}

Directory* Directory::getParent() {
	return this->parent;
}

void Directory::setDirectory(Directory* dir) {
	if (this->isDirectory()) {
		this->dir = dir;	
	}
	return;
}

bool Directory::isHome() {
	return this->isHomeDir;
}
