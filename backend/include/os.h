#ifndef OS_H
#define OS_H

#include "logger.h"
#include <string>
#include <map>

class OperatingSystem {
	public:
		OperatingSystem();
		~OperatingSystem();
		void Operate(std::string);
	private:
		// Logger	
		OperatingSystemLogger *Logger;
		
		// Functions and Commands for determining the initial Command	
		void InitializeCommandMap();
		std::map<std::string, int> CommandMap;
};

#endif
