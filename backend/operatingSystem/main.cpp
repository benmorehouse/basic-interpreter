#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include "arithmetic.h"
#include "boolean.h"
#include "command.h"
#include "interpreter.h"
#include "compiler.h"

int main(int argc, char* argv[]){
	if(argc < 2){ std::cout<<"Please enter in a text file that the interpreter can operate"<<std::endl;}
	std::ifstream input;
	input.open(argv[1]);
	
	if(input.fail() == true){
		std::cout<<"Please enter in a file that the interpreter can use"<<std::endl;
	}

	Interpreter interpreter(input);
	Compiler compiler(interpreter.getOutFile(),"compilerOutput.txt",&interpreter);
	compiler.main();
}

