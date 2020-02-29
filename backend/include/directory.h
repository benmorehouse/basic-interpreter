#ifndef DIRECTORY_H
#define DIRECTORY_H

#include <iostream>
#include <string>

struct File {
	std::string name;
	int lineCount;
	bool isBasic;
	std::string *lines;
};

class Directory {
	// should be an iterator for this.
	public:
		// Directory(name, <parent of this directory>)
		Directory(std::string name, Directory*);
		void setName(std::string name);
		std::string getName();
		bool isHome();
	
		File* getFile();
		
		Directory* getDirectory();
		bool isDirectory();
		void setDirectory(Directory*);
		Directory* getParent();
		void setParent(Directory*);
	private:
		std::string Name;
		bool IsDir;
		File* file;
		Directory* dir;
		bool isHomeDir;
		Directory* parent;
};

#endif
