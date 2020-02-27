#include <iostream>
#include <string>
#include <cstddef>        // std::size_t
#include "../include/directory.h"

Directory::Directory(std::string name) {
	this->Name = name;

	std::string fileExtension = "";
	std::string::iterator it = name.end();
	
	while(it != name.begin()) {
					
		--it;	
	}

}	

