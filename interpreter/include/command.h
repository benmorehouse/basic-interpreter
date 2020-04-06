#ifndef COMMAND_HPP
#define COMMAND_HPP
#include "boolean.h"
#include "arithmetic.h"

class PrintCommand{ // want this to take in a variable, and either print the variable
public:	
	PrintCommand(NumericExpression*); // will take the numericExpression and get the value
	std::string format() const;
	NumericExpression* getData(){return this->data;}
private:
	NumericExpression* data;
};

class LetVar{
public:
	LetVar(std::string,NumericExpression*);	 // 1 LET HAN 5 -> here we take in a name and what we parse as numericExpression
	std::string format() const;
	std::string getName(){return this->name;}
	NumericExpression* getNEXP(){return this->next;}

private:
	std::string name;
	NumericExpression* next;	
};

class LetVarArray{
public:
	LetVarArray(Array*, NumericExpression*, NumericExpression*);  // <LINE> LET <VAR> [<NEXP1>] <NEXP2> -> var is array
	std::string format() const;

private:
	Array* array;
	NumericExpression* index;
	NumericExpression* value;
};

class Goto{ 
public:
	Goto(NumericExpression*); // goes to the numeric expression line
	std::string format() const;
private:
	NumericExpression* line;
};

class IfGoto{ // <LINE> IF <BEXP> THEN <JLINE>
public:
	IfGoto(NumericExpression*, NumericExpression*);
	std::string format() const;
private:
	NumericExpression* booleanExp;
	NumericExpression* line;
};

class GoSub{
public:
	GoSub(NumericExpression*);
	std::string format() const;
private:
	NumericExpression* line;
};

class ReturnCommand{
public:
	ReturnCommand(){};
	std::string format() const;
};

class EndCommand{
public:
	EndCommand(){};  
	std::string format() const;
};
	
#endif
