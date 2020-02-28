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
		Directory(std::string name);
		void SetName(std::string name);
		std::string GetName();
		bool IsDirectory();
		File* GetFile();
		Directory* GetDirectory();
		void SetDirectory(Directory*);

	private:
		std::string Name;
		bool IsDir;
		File* file;
		Directory* dir;
};

#endif
