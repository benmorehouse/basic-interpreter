#ifndef COMPILER_H
#define COMPILER_H

#include <iostream>
#include <fstream>
#include <vector>
#include <map>
#include <stack>
#include <sstream>
#include "interpreter.h"
#include <set>

/*
 * This will analyze and do what the interpretor outfiles to us
 */

class Compiler{
public:
	Compiler(std::string iFileName, std::string oFileName, Interpreter *_interpreter){
		this->ifile.open(iFileName);
		if(this->ifile.fail()){
			return;
		}
		this->ofile.open(oFileName);
		if(this->ofile.fail()){
			return;
		}
		this->ifileName = iFileName;
		this->ofileName = oFileName;
		this->ifile.close();
		this->ofile.close();
		this->interpreter = _interpreter;
	}
	
	~Compiler(){};
	void main(); // this will be the main that will execute after we initialize

private:
	/*
	 * This is a map used to keep the 
	 *	 key: the command line number
	 *	 value: the command 
	 *
	 *	 this gives us a worst case scenario of o(log n) access of elements
	 */

	std::map<int, std::string>commandList; // used to keep map of each line of the BASIC code
	std::vector<int> lineNumbers;
	/*
	 * Things we want to check for with the map is that we access things that exist
	 */	
	std::map<std::string, int> variableValues;
	/*
	 * Things we want to check for with the gosublines is that
	 * 	we dont call to pop if the gosub is empty -> we return an error
	*/

	std::stack<int> goSubLines; // these are the three variables we use to implement the more complex algorithms

	/*
	 * These two variables are used so that we can take the data which is the interpretor's output file
	 * And then write to the new outputfile that we want
	 */
	std::ifstream ifile; // from the interpretor 
	std::ofstream ofile; // the final results of the compiler
	std::string ifileName;
	std::string ofileName;
	Interpreter *interpreter; 
	/*
	 * this function will be used to assign all of the lines in the file to the commandList map
	 * will return false if anything goes wrong otherwise will return true
	 */
	
	bool parsefile();

	/*
	 * The point of this function is to genarate an error message for the user
	 */

	std::string getErrorMessage(int line, std::string error){
		return "Error on line " + std::to_string(line) + ": " + error;	
	}

	std::string getWarningMessage(int line, std::string warning){ // also gonna have to do iterator these
		return "Warning on line " + std::to_string(line) + ": " + warning;	
	}
	/*
	 * Within main, we need a way to go get which command we are doing. We will do something similar to the strings.Field() command in Golang to get the second field element ie the command
	 * this will return a string of it given the line input that we want 
	 */
	std::string getCommand(int);
	std::string getLine(int);
	int jumpToCommand(int);
};	

#endif
