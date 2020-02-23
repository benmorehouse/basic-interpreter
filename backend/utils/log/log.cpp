#include <iostream>
#include <string>
#include "log.h"

//######################################################
//############# Colors used for logging ################

#define RESET   "\033[0m"
#define RED     "\033[31m"      
#define YELLOW  "\033[33m"     
#define BLUE    "\033[34m"      

//######################################################

Logger::Logger() {}

Logger::~Logger() {}

void Logger::Error(std::string err) {
	std::cout << RED << "ERROR: " << err << RESET << std::endl;
}

void Logger::Warning(std::string err) {
	std::cout << YELLOW << "WARNING: " << err << RESET << std::endl;
}

void Logger::Info(std::string err) {
	std::cout << BLUE << "INFO: " << err << RESET << std::endl;
}

void Logger::Debug(std::string err) {
	std::cout << "DEBUG: " << err << std::endl;
}

