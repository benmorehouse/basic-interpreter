#include <iostream>
#include "../utils/log/log.h"

class OperatingSystemLogger : protected Logger {
	public:
		OperatingSystemLogger();
		~OperatingSystemLogger();
		void Error(std::string err) override;
		void Warning(std::string err) override;
		void Info(std::string err) override;
		void Debug(std::string err) override;
};
