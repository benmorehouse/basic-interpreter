#ifndef OS_H
#define OS_H
#include "mlist.h"

class OS{
	public:
		//constructor
		OS();
		//destructor
		~OS();
		//read and write a file (depending on the boolean flag)
		void file(string fname,bool isRead);
		//delete file and folder:
		//Note; node that holds the data is permanently deleted, but not the data itself
		void del(string fname);
		//ls - list a file in the directory (the order of listing depends on the option)
		void ls(string option);
		//create folder or file (depending on the boolean flag)
		void create_item(string fname,bool isFolder);
		//delete directory or file
		void rm(string fname);
		//cd - changing directory
		void cd_dir(string fname);
		//pwd - present working directory (print path)
		void present_working_dir();
		//search item in dir
		Node* search_item(string s);
	private:
		Data* root_data;//data for root - technically don't really need this! but help keeping track of root
		MList wd;//keep track of the current directory, LIFO manner
		MList dataStack;//keep track of all dynamically allocate, LIFO manner
};
#endif
