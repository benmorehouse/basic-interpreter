#include <iostream>
#include <string>
#include <"

//######################################################
//############# Colors used for logging ################

#define RESET   "\033[0m"
#define RED     "\033[31m"      
#define YELLOW  "\033[33m"     
#define BLUE    "\033[34m"      
#define OS 	"OS::" 

//######################################################

void OperatingSystemLogger::Error(std::string err) {
	std::cout << OS << RED << "ERROR: " << err << RESET << std::endl;
}

void OperatingSystemLogger::Warning(std::string err) {
	std::cout << OS << RED << "ERROR: " << err << RESET << std::endl;
}

void OperatingSystemLogger::Info(std::string err) {
	std::cout << OS << RED << "ERROR: " << err << RESET << std::endl;
}

void OperatingSystemLogger::Debug(std::string err) {
	std::cout << OS << RED << "ERROR: " << err << RESET << std::endl;
}

