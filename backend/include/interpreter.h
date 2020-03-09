#ifndef INTERPRETER_HPP
#define INTERPRETER_HPP

#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <fstream>
#include "arithmetic.h"
#include "boolean.h"
#include "command.h"

class Interpreter {
public:
	Interpreter(std::ifstream& in);
	~Interpreter();
	void write();
	std::string getOutFile(){
		return this->ofileName;
	}

/*	double getNumericExpressionValue(std::string line){
		int i = 0;
		NumericExpression *NEXP = this->parseNumericExpression(line, i);
		return NEXP->getFinalValue();
	}
	*/
	NumericExpression* parseNumericExpression(std::string , int &position);

private:
// functions
	std::string parseVariableName(std::string input, int &position);
	Variable* parseConstant(std::string line, int &position);
	void trimWhiteSpace(std::string input, int &position);
	void parse(std::ifstream& in);
	int initialCheck(std::string);
	std::string loadString(std::stringstream&);
// functions

// interpreter objects

	std::vector<std::string> commands; // used to get the parsed commands
	std::ofstream ofile; // this will be what interpreter writes to 
	std::string ofileName;
// interpreter objects
};

#endif
