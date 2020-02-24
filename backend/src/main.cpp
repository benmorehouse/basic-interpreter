#include "../include/logger.h"

int main() {
	OperatingSystemLogger *logger = new OperatingSystemLogger();
	logger->Error("This is an error");
	delete logger;
}
