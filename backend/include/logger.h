#include <string>

class Logger {
	public:
		Logger(){};
		~Logger(){};
		virtual void Error(std::string);
		virtual void Warning(std::string);
		virtual void Info(std::string);
		virtual void Debug(std::string);
};

class OperatingSystemLogger : protected Logger {
	public:
		OperatingSystemLogger();
		~OperatingSystemLogger() ;
		void Error(std::string);
		void Warning(std::string);
		void Info(std::string);
		void Debug(std::string);
};


