#include <iostream>
#include <string>
#include <cstddef>        // std::size_t
#include "../include/directory.h"

Directory::Directory(std::string name) {
	this->Name = name;
	this->IsDir = true; // postulate: assume it is a dir, prove otherwise

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

	this->dir = nullptr;
	this->file = nullptr;
}	

void Directory::SetName(std::string name) {
	this->Name = name;
}

std::string Directory::GetName() {
	return this->Name;
}

bool Directory::IsDirectory() {
	return this->IsDir;
}

File* Directory::GetFile() {
	if(!this->IsDirectory()) {
		return this->file;
	}
	return nullptr;
}

Directory* Directory::GetDirectory() {
	if(this->IsDirectory()) {
		return this->dir;
	}
	return nullptr;
}

void Directory::SetDirectory(Directory* dir) {
	if (this->IsDirectory()) {
		this->dir = dir;	
	}
	return;
}


