#include "../include/logger.h"
#include "../include/os.h"

int main(int argc, char **argv) {
	Logger *log = new Logger();
	OperatingSystem *os = new OperatingSystem();
	
	if(argc < 2) {
		log->Error("Not enough arugments given for operating system");
		return 0;
	}
		
	os->Operate(argv, argc);
	delete log;
	delete os;
}




