#include "../include/logger.h"
#include "../include/os.h"

int main(int argc, char **argv) {
	Logger *log = new Logger();
	OperatingSystem *os = new OperatingSystem();
	if(argc < 2) {
		log->Error("Not enough arugments given for operating system");
		return 0;
	}
	
	log->Info("Running the operating system ");	
	os->Operate(argv, argc);
	log->Info("Finished Running the operating system ");	
	delete log;
	delete os;
}




