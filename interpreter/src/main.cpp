#include "../include/logger.h"
#include "../include/interpreter.h"
#include "../include/compiler.h"
#include <fstream>
#include <string>

int main(int argc, char **argv) {
	Logger *log = new Logger();
	if(argc < 3) {
		log->Error("Not enough arugments given for operating system");
		return 0;
	}
	
	std::string ifileName(argv[1]);
	std::string ofileName(argv[2]);

	std::ifstream ifile;
	ifile.open(ifileName);

	Interpreter *interpreter = new Interpreter(ifile);
	std::string interpretedFile = interpreter->getOutFile();
	
	Compiler *compiler = new Compiler(interpretedFile, ofileName, interpreter);
	compiler->main();
	
	ifile.close();
	delete interpreter;
	delete compiler;
	delete log;
}

