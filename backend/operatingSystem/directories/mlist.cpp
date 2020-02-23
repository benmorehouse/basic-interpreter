#include <iostream>
#include <string>
#include <cstring>
#include "mlist.h"
using namespace std;
//------------------------------------------------------------------------------
//Constructor, construct an empty doubly linked list.
MList::MList(){
    ntop = NULL;
    nbottom = NULL;
}

//Destructor, properly clears and deallocates everything of current list.
//simply call the clear function if you have already implemented that.
MList::~MList(){
    clear();
}


//Note: RECURSION --- This function should be implemented using recursion.
//this function properly clears and deallocates everything of current list.
//there should be no loop in this function

void MList::clear(){
    if(isEmpty()) return; // this is the base case

    Data * data = ntop->nodeData;
    if (data->childList != NULL){
        data->childList->clear();
    }
    delete ntop;
    ntop = ntop->next;
    clear();
}

//returns a boolean true if the list is empty, false if otherwise.
bool MList::isEmpty(){
    return (ntop==NULL);
}

/*
    Add or insert a new node with d_item as its data at the top of the list.
    You need to dynamically allocate memory for the new node.
*/
void MList::push_top(Data* d_item){
    Node *newNode = new Node;
    newNode->nodeData = d_item;
    newNode->next = ntop;
    newNode->prev = NULL;

    if (isEmpty()) {
        nbottom = newNode;
        ntop = newNode;
        return;
    }

    ntop = newNode;
}
/*
    Add or insert a new node with d_item as its data at the bottom of the list.
    You need to dynamically allocate memory for the new node.
*/
void MList::push_bottom(Data* d_item){
    if(nbottom == NULL){
        Node *temp = new Node;
        temp->prev=NULL;
        temp->nodeData = d_item;
        temp->next=NULL;
        nbottom = temp;
    }
    else{
        Node *temp = new Node;
        temp->prev = NULL;
        temp->nodeData = d_item;
        temp->next = nbottom;
        nbottom = temp;
    }
}

/*
    Delete or remove the top node of the list.
    Once pop, you need to properly deallocate the memory for the node (not the data).
    If your list is currently empty, you don't have to do anything.
*/
void MList::pop_top(){
    if(isEmpty()) return;  // just to test to make sure that this is not empty

    Node *temp = ntop;
    ntop = temp->next;
    if(ntop!=NULL){ // if the next thing in the list was not the end of the list
        ntop->prev = NULL;
    }
    else{
        nbottom = NULL;
    }
    delete temp;
}

/*
    Delete or remove the bottom node of the list.
    Once pop, you need to properly deallocate the memory for the node (not the data).
    If your list is currently empty, you don't have to do anything.
*/
void MList::pop_bottom(){
    if(isEmpty()) return;  // just to test to make sure that this is not empty

    Node *temp = nbottom;
    nbottom = temp->prev;
    if(nbottom!=NULL){ // if the next thing in the list was not the end of the list
        nbottom->prev = NULL;
    }
    else{
        ntop = NULL;
    }
    delete temp;
}

/*
    Note: RECURSION --- This function should be implemented using recursion.
    Search a node in the list that has data with the same name as the parameter.
    If one is found, return the pointer to the node.
    If one is not found, return NULL.
*/
Node* MList::search(Node* start, string name){
    if(isEmpty()) return NULL; // if the thing is empty to begin with then return NULL
    if(start==NULL) return NULL; // if we are at the end of our ropes here then we are going to stop

    //if both of the base cases are not applicable then go into the rest of the search

    if(start->nodeData->name == name){
        return start;
    }
    else{
        start = search(start->next,name);
        return start; // when it echos all the way back out, it will return what start is
    }
}


//Swapping node a with node b.
void MList::swapNode(Node* a, Node*b){ // changes everything about the nodes. They switch from one to another
//this includes the data, previous and next pointers
    Node * temp = new Node;
    temp->nodeData = a->nodeData;
    temp->next= a->next;
    temp->prev= a->prev;

    a->nodeData = b->nodeData;
    a->next= b->next;
    a->prev= b->prev;

    b->nodeData= temp->nodeData;
    b->next = temp->next;
    b->prev= temp->prev;
}

/*
    This function PERMANENTLY deletes a node from the list given a pointer to that node.
    You can assume that the node exists in the list - no need to search for the node.
*/
void MList::deleteNode(Node* a){ // check this!!
if(search(ntop,a->nodeData->name)==NULL) return; // if the node doesnt exist then dont do anything!
    if(nbottom==a) pop_bottom();
    else if(ntop==a) pop_top();
    else{
        Node* temp_1 = new Node;
        temp_1 = a->prev;
        temp_1->next = a->next;

        Node* temp_2 = new Node;
        temp_2 = a->next;
        temp_2->prev = a->prev;

        delete temp_2;
        delete temp_1;  // because we dont need these now. We have
    }
}

/*
    Unlike the delete function above, this function does NOT permanently delete the node.
    What you need to do at the last step is to set the node's prev and next pointers to NULL.
    Again, you can assume that the node exists in the list - no need to search for the node.
    Note: this function will be useful
        when we want to insert the given node later at another position after you remove it from the list.
*/
void MList::removeNode(Node* a){
    Node * prevNode = a->prev;
    Node * nextNode = a->next;

    if (prevNode == NULL) {
        ntop = nextNode;
    }
    if (nextNode == NULL) {
        nbottom = prevNode;
    }

    if (prevNode != NULL) {
        prevNode->next = nextNode;
    }
}

/*
    Insert node a after node b.
    Note: b can be NULL, Q: what does that mean? // there is nothing in the set yet.
    otherwise, you can assume that b exists in the list.
*/
void MList::insertAfter(Node *a, Node* b){ // do you delete the memory now or later
    if (b == NULL) {
        ntop = a;
        nbottom = a;
        return;
    }

    // Given a list such as ...
    // B  -->  C
    // This means that nTop == B and nBottom = C
    // To insert A after B... we need to make sure that A's "next" pointer points to C!
    // If B doesn't have a next pointer, then it means we're at the end of the list, and our new nBottom is now A!
    // Now, this means that nTop == B, b -> A , A -> C  and nBottom is still C
    //
    // Given a list B
    // this means nTop == B and nBottom == B and B->next == null
    // to insert A after B we need to now set "A" to the bottom of the list
    // which means we now have:
    // B --> A    nTop = B and nBottom = A
    if (b->next == NULL) {
        nbottom = a;
    }

    a->next = b->next;
    a->prev = b;
    b->next = a;
}


/*
    Note: RECURSION --- This function should be implemented using recursion.
    Implement a SELECTION SORT algorithm using recursion.
    The function takes in a start node.
    Things to keep in mind:
        1). sort should NOT be run if the list is empty.
        2). if mode = true, sort by name in alphabetical order
            (i.e., A-Za-z) based on ASCII number if mode = true
        3). if mode = false, sort by size in ascending order (low->high)
    You can search any online resources available to you (i.e. search 'selection sort')
    to understand the basic idea behind the selection sort algorithm.
    Note:
        1). there can only be at most one loop in this function
        2). this function is private
            see sortByNameSelection and sortBySizeSelection for the calls to this function
*/
void MList::sortSelection(Node* start, bool mode){
    if(isEmpty()) return;
    if(start == NULL) return;// if the last start was the end and now we are pointing to a NULL

    Node *nextposition = start->next;

    Node *place = start;
    Node *lowest = start;

    while(place!=NULL){ // while we are not at the end of the list of nodes
        if(mode){
            if(lowest->nodeData->name > place->nodeData->name){ // if the one we are looking at is
                lowest = place;
            }
        }
        else{
            if(lowest->nodeData->size > place->nodeData->size){
                lowest = place;
            }
        }
        place = place->next;
        //lowest will keep track of which one is lowest. Place will be what goes through. and then at the end lowest and start will switch and will run next iteration with next position
    }

    swapNode(lowest, start); // this may not work as you are intending
    sortSelection(start->next, mode);
}


/*
    Note: RECURSION --- This function should be implemented using recursion.
    Implement an INSERTION SORT algorithm using recursion.
    The function takes in a start node.
    Things to keep in mind:
        1). sort should NOT be run if the list is empty.
        2). we are trying to sort by name in alphabetical order
            (i.e., A-Za-z) based on ASCII number.
    You can search any online resources available to you (i.e. search 'insertion sort')
    to understand the basic idea behind the insertion sort algorithm.
    The gif from the wikipedia page of this algorithm is the easiest to understand in my opinion
    Link: https://en.wikipedia.org/wiki/Insertion_sort
    Note:
        1). there can only be at most one loop in this function
        2). this function is private, see sortByNameInsertion for the call to this function
*/
void MList::sortInsertion(Node* start){
    if(isEmpty()) return; // if it is empty
    if(start==NULL) return;  // if the last start was the end and now we are pointing to a NULL
    // when we start this we are gonna start at the header of this
    Node *nextposition = start->next; // this is gonna be used so that we go to the next one on the next iteration of recursion

    while(start->prev!=NULL){ // once it is all the way to the left
        if(start->prev->nodeData->name > start->nodeData->name) swapNode(start->prev, start);
                // if the previous node's name is high on ascii, switch the nodes

        else break; // if it isnt, then break and go to the next node and push it through the sorted part on the left
    }

    if(start->prev == NULL) sortInsertion(nextposition);
}



/*
    Note: RECURSION --- This function should be implemented using recursion.
    Traverse and print our list from node n to the top.
    we will seperate each node printed by the passed-in delimiter.
    If a node you encounter is a folder type, do

        cout<<....name.....<<delim;

    If a node you encounter is a file type, do

        cout<<....name.....<<"("<<..size..<<")"<<delim;

    Note: please do NOT include <<endl; (aka newline)
    Example output (assume that the delimiter is a single space):

        folder1 file1(0) folder2 file2(0)

    There should be no loop in this function
    See printBtT function for the call to this function.
    This function is private
*/
void MList::traverseToTop(Node* n, string delim){
    if(isEmpty() || n == NULL) return;

    Data *data = n->nodeData;
    if (data->isFolder) {
        std::cout<<data->name<<"("<<data->size<<")"<<delim;
    } else {
        std::cout<<data->name<<delim;
    }

    traverseToTop(n->prev, delim);
}

/*
    Note: RECURSION --- This function should be implemented using recursion.
    Traverse and print our list from node n to the bottom.
    we will seperate each node printed by the passed-in delimiter.
    If a node you encounter is a folder type, do

        cout<<....name.....<<delim;

    If a node you encounter is a file type, do

        cout<<....name.....<<"("<<..size..<<")"<<delim;

    Note: please do NOT include <<endl; (aka newline)
    Example output (assume that the delimiter is a single space):

        folder1 file1(0) folder2 file2(0)

    There should be no loop in this function
    See printTtB function for the call to this function.
    This function is private
*/
void MList::traverseToBottom(Node* n,string delim){
    if(isEmpty() || n->next == NULL) return;

    Data *data= n->nodeData;
    if (!data->isFolder) {
        std::cout<<data->name<<"("<<data->size<<")"<<delim;
    } else {
        std::cout<<data->name<<delim;
    }

    traverseToBottom(n->next, delim);
}

//------------------------------------------------------------------------------
//FUNCTIONS BELOW ARE COMPLETE, PLEASE DON'T CHANGE ANYTHING HERE
//------------------------------------------------------------------------------

//getting the pointer to the top node.
Node* MList::top(){
    return ntop;
}

//getting the pointer to the bottom node.
Node* MList::bottom(){
    return nbottom;
}

//call traverseToBottom to print all item in the list from bottom node to top node
//the list of items is separated by a given delimiter
void MList::printBtT(string delim){
    //create a temp pointer to hold bottom
    traverseToTop(nbottom,delim);
}

//call traverseToBottom to print all item in the list from top node to bottom node
//the list of items is separated by a given delimiter
void MList::printTtB(string delim){
    Node* tmp = ntop;
    traverseToBottom(tmp,delim);
}

//call sortSelection function, mode = true = sort by name
//public member
void MList::sortByNameSelection(){
    bool mode = true;
    Node *tmp = ntop;
    sortSelection(tmp,mode);
}

//call sortSelection function, mode = false = sort by size
//public member
void MList::sortBySizeSelection(){
    bool mode = false;
    Node *tmp = ntop;
    sortSelection(tmp,mode);
}

//call sortInsertion function
//public member
void MList::sortByNameInsertion(){
    Node *tmp = ntop;
    sortInsertion(tmp);
}
