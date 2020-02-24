#include <string>

class Logger {
	public:
		virtual void Error(std::string);
		virtual void Warning(std::string);
		virtual void Info(std::string);
		virtual void Debug(std::string);
};

class OperatingSystemLogger : protected Logger {
	public:
		OperatingSystemLogger();
		~OperatingSystemLogger();
		void Error(std::string) override;
		void Warning(std::string) override;
		void Info(std::string) override;
		void Debug(std::string) override;
};
