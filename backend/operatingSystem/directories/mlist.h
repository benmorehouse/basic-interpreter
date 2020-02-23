#ifndef MLIST_H
#define MLIST_H
#include <string>
using namespace std;

//------------------------------------------------------------------------------
//IMPORTANT: You are not allowed to remove or add any parameters to any functions.
//------------------------------------------------------------------------------


//struct and class declarations
struct Data;
class MList;

/*
   Node struct:
   		2 pointers - prev/next are for doubly linked list
		1 pointer points to the Data object.
*/
struct Node{
	//for doubly linked list
	Node* prev;
	Node* next;
	//a pointer that points to the data object that this node holds
	Data* nodeData;
};


//Data struct i.e., data object
struct Data{
	string name;
	bool isFolder;
	string text;
	size_t size;
	MList* childList;//a list of children - a pointer to MList
};


class MList{
public:
	//constructor
	MList();
	//destructor
	~MList();
	//push item to list (from top)
	void push_top(Data* d_item);
	//delete the top Node
	void pop_top();
	//push item to list (from bottom)
	void push_bottom(Data* d_item);
	//delete the bottom Node
	void pop_bottom();
	//check if list is empty
	bool isEmpty();
	//clear/empty out the list
	void clear();
	//swap two nodes
	void swapNode(Node* a, Node*b);
	//delete node permanently
	void deleteNode(Node* a);
	//remove the node, not permantly, useful for adding it later. If really want to delete, use delete function.
	void removeNode(Node* a);
	//insert Node a after Node b
	void insertAfter(Node *a, Node* b);
	//search a node with a given name from the starting point.
	Node* search(Node* start,string name);
	//retrieve the top node
	Node* top();
	//retrieve the bottom node
	Node* bottom();
	//print from bottom to the top node
	void printBtT(string delim);
	//print from top to the bottom node
	void printTtB(string delim);
    //call to recursive sort (by name) function using Selection sort algorithm
	void sortByNameSelection();
	//call to recursive sort (by name) function using Insertion sort algorithm
	void sortByNameInsertion();
	//call to recursive sort (by size) function using Selection sort algorithm
	void sortBySizeSelection();
    
private:
	//a pointer that points to the first node (top)
	Node* ntop;
	//a pointer that points to the last node (bottom)
	Node* nbottom;
	//recursively traverse the list from n_item to the top node
	void traverseToTop(Node* n_item,string s);
	//recursively traverse the list from n_item to the bottom node
	void traverseToBottom(Node* n_item,string s);
    //recursive sort using Selection sort algorithm, mode determines if it's by size or by name
	void sortSelection(Node* start, bool mode);
	//recursive sort (by name) using Insertion sort algorithm
	void sortInsertion(Node* start);
};
#endif