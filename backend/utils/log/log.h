class Logger {
	public:
		Logger();
		~Logger();
		virtual void Error(std::string err);
		virtual void Warning(std::string err);
		virtual void Info(std::string err);
		virtual void Debug(std::string err);
};


