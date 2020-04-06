#ifndef DIRECTORY_H
#define DIRECTORY_H

#include <iostream>
#include <string>
#include <map>
#include <vector>

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
		// IF file does not have dot it is considered a file.
		Directory(std::string name, Directory*);
		void setName(std::string name);
		std::string getName();
		bool isHome();
	
		File* getFile(std::string);
		Directory* getDirectory(std::string);
		std::vector<std::string>* getAllSubAsName();
		
		bool isDirectory();
		void addDirectory(Directory*);
		void addFile(Directory*);

		void removeDirectory(Directory*);
		void removeFile(Directory*);

		Directory* getParent();
		void setParent(Directory*);
	private:
		std::string Name;
		// map corresponding to each subdirectory or file within the directory.
		std::map<std::string, Directory*> directories;
		bool IsDir;
		File* file;
		bool isHomeDir;
		Directory* parent;
		// Function used as helper for public getters.
		Directory* getElementFromDirectoryMap(std::string);
};

#endif
