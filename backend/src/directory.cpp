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
		if(it->second == '.'){
			this->IsDir = false;
			fileExtension.insert(0, it->second);
			break;
		}
		
		fileExtension.insert(0, it->second);
		--it;	
	}

	this->Dir = nullptr;
	this->file = nullptr;
}	

