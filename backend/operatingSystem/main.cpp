#include "logger.h"

int main() {
	
	OperatingSystemLogger *logger = new OperatingSystemLogger();
	logger->Error("This should output an error message");
	logger->Error("This should output an error message");
	logger->Error("This should output an error message");
	logger->Error("This should output an error message");
	delete logger;
}
