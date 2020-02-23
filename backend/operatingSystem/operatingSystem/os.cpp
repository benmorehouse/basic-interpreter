#include <iostream>
#include "os.h"
#include <vector>

using namespace std;

/*
Constructor
	Dynamically allocate root data.
	set isFolder = true.
	dynamically allocate Mlist for our root data as well (since it's a directory type).
    push our newly created data object pointer to wd and dataStack from the top.
*/
OS::OS(){
    //is folder is part of the node data
    root_data = new Data; // dynamically allocate this new data
    root_data->isFolder = true; // it is a folder
    root_data->childList = new MList; // then make its childlist new as well
    wd.push_top(root_data); // then push it to the top of the mlist that we are currently on 
    dataStack.push_top(root_data);
}

/*
Destructor to clean-up memory, i will leave this for you to figure it out.
*/
OS::~OS(){
    wd.clear();
    dataStack.clear();
    delete root_data; // not sure if i need to get rid of this as well
}

/*
Search a node in the current directory
	If one is found, return a pointer to that node
	If one is not found, return NULL
*/
Node* OS::search_item(string fname){
    if(wd.isEmpty()) return NULL; 
    Node* data = wd.top();
    return wd.search(data,fname);
}

/*
Delete a node in the current directory
	Note: this function only deletes (permanently) the node, not the Data obj the node points to
	If the item you want to delete doesn't exist in the current directly, do

		cout<<"Error: cannot find file or directory '"<<fname<<"'"<<endl;
*/
void OS::del(string fname){
    Node* data = wd.top();
    data = wd.search(data,fname);
    if(data==NULL){
        std::cout<<"Error: cannot find file or directory '"<<fname<<"'"<<std::endl;
        return;
    } 
    else{
        // at this point we know that the node exists in this directory
        wd.deleteNode(data);
    }
}

/*
Create a file or a folder, use boolean isFolder to tell (true if folder, false if file)
Things to keep in mind:
	1). Don't forget to search for a node in the current directory with the same name first.
		if the name already exists, do:

				cout<<"Directory or file name '"<<fname<<"' already exists"<<endl;

	2). If you are creating a folder, make sure to allocate a memory for its MList object
		and set the boolean isFolder appropriately.
    3). At this point you should initialize the size of file and folder to 0
	4). Once the data object is created, add the pointer to that to dataStack from the "top" 
		and add the node to the current directory list from the "bottom".
	5). Once added, sort the current directory list by name.
*/
void OS::create_item(string fname, bool isFolder){
    Node* data = wd.top();
    data = wd.search(data,fname);
    if(data == NULL){
        std::cout<<"Directory or file name '"<<fname<<"' already exists"<<endl;
        return;
    }
    else{
        if(isFolder){
            Data* newfolder = new Data; // creating new data in the form of a folder 
            newfolder->name = fname;
            newfolder->isFolder= true;
            newfolder->childList = new MList;
            newfolder->size = 0;
            wd.push_bottom(newfolder);
            wd.sortByNameInsertion();
            dataStack.push_bottom(newfolder);
        }
        else{
            Data* newfile = new Data;
            newfile->name = fname;
            newfile->isFolder = false;
            newfile->childList = NULL;
            newfile->size = 0;
            wd.push_bottom(newfile);
            wd.sortByNameInsertion();
            dataStack.push_bottom(newfile);
        }
    }
}

/*
Read or write a file according to the boolean isRead (true = read, false = write)
Things to keep in mind:
	1). make sure that a file "fname" exists in our current directly, if not

			cout<<"Error: cannot find file '"<<fname<<"'"<<endl;

	2). if it exists, make sure that it is a file type, not a folder type. If it is a folder type,

			cout<<"Error: '"<<fname<<"' is not a file"<<endl;

	3). read mode is simply just cout<<...text...<<endl;
	4). for write mode you need to allow user input, use these 3 lines below:

                cout<<"$> ";
				string input;
				getline(cin,input);

		then simply just set text to the input string.
	5). the size of the file should be based on the length of the input string.
*/
void OS::file(string fname, bool isRead){
    Node* data = wd.top();
    data = wd.search(data,fname);
    if(data==NULL){
        std::cout<<"Error: cannot find file or directory '"<<fname<<"'"<<std::endl;
        return;
    } 
    else if(data->nodeData->isFolder){
		std::cout<<"Error: '"<<fname<<"' is not a file"<<std::endl;
        return; 
    }
    else{
        if(isRead){
            std::cout<<data->nodeData->text<<std::endl;
        }
        else{
            std::cout<<"$> ";
            std::string input;
            std::getline(std::cin,input);

            data->nodeData->text = input;
            data->nodeData->size = input.length();
        }
    }
}


//Change directory
void OS::cd_dir(string fname){
	if(fname == ".."){
		//this option takes you up one directory (move the directory back one directory)
		//you should not be able to go back anymore if you are currently at root.
        if(wd.top() == dataStack.top()) return;
        // we should use the top node in wd, find it in datastack and then change wd to that in datastack
        Node* changer = wd.top();
        // now find changer in datastack
        changer = dataStack.search(dataStack.top(),changer->nodeData->name)->prev;
        wd.clear();
        // idk what to do with this 
	}
    else if(fname == "~"){
		//this option takes you directly to the home directory (i.e., root).
        Node* topOfDataStack= dataStack.top();
        wd.clear(); // does this fuck shit up bad??
        // wd.push_bottom(topOfDataStack->nodeData);
        wd = *topOfDataStack->nodeData->childList;
	}
    else{
		/*
			This option means that you are trying to access (go into) a directory
			1). check whether there exists a node with that name, if you cannot find it:

					cout<<"Error: cannot find directory '"<<fname<<"'"<<endl;

			2). if it does exist, check whether it is a folder type. If it is a file type:

					cout<<"Error: '"<<fname<<"' is not a directory"<<endl;
			
			3). after checking both things, you can safely move to another directory
		*/
        Node* data = wd.top();
        data = wd.search(data,fname);
        if(data==NULL){
            std::cout<<"Error: cannot find file or directory '"<<fname<<"'"<<std::endl;
        return;
        } 
        else if(! data->nodeData->isFolder){
            std::cout<<"Error: '"<<fname<<"' is not a directory"<<std::endl;
        }   
        else{
            Node* newlist = dataStack.search(wd.top(),fname);  // this is gonna be the node for the new list, find it in datastack
            if(newlist->prev!=NULL){
                newlist = newlist->prev;
            }
            else newlist = dataStack.top();
                // now we are to the parent of the mlist
            wd.clear();
            wd = *newlist->nodeData->childList; 
        }
	}
}

//Print list of item in the current directory, the way you print it will be according to the passed-in option
void OS::ls(string option){
	if(option == "-d"){
		//Default option - print the list of items in the current directory from top to bottom.
		//use a single space as delimiter.
        wd.printTtB(option);
	}
	else if(option == "-sort=name"){
		//Use Insertion Sort to sort items in the current directory and print from top to bottom (alphabetically A-Za-z).
		//use a single space as delimiter.
        wd.sortByNameInsertion(); 
	}
    else if(option == "-r"){
		//Reverse print the list of items in the current directory (i.e., print form bottom to top).
		//use single space as delimiter.
        wd.printBtT(option);
	}
    else if(option == "-sort=size"){
		//Sort list by size and print the list of items in the current directory from top to bottom.
		//use single space as delimiter.
        wd.sortBySizeSelection(); 
	}
	else{
		cout<<"usage: ls [optional: -r for reverse list, -sort=size to sort list by file size, -sort=name to soft list by name]";
	}
}

//Priting path from root to your current directory.
//use slash "/" as our delimiter.
//Example output: root/cs103/usc/viterbi/
void OS::present_working_dir(){
    Node* temp = wd.top();
    std::vector<std::string> directories;
    while(temp!=dataStack.top()){
        directories.push_back(temp->nodeData->name);
        
    }


}
