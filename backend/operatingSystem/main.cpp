#include <iostream>

class errorHandler {
	public:
		errorHandler();
		~errorHandler();
		void Err(string err);
	private: 
		void handleError(string err);
}

errorHandler::errorHandler() {}

errorHandler::~errorHandler() {}

void errorHandler::Err(string err) {
	this.handleError(err);
}

void handleError(string err) {
	std::cout<<"ERROR : "<<err<<std::endl;
	return
}


int main(int argc, char **argv) {
	
	if (argc < 2) {
		
	}
}
