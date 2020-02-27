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
	public:
		Directory(std::string name);
		void SetName();
		std::string GetName();
		bool IsDirectory();
		File* GetFile();
		Directory* GetDirectory();

	private:
		std::string Name;
		bool IsDir;
		File* file;
		Directory* dir;
};

#endif
