#include "command.h"
#include <iostream>

PrintCommand::PrintCommand(NumericExpression *input){
	this->data = input;
}

std::string PrintCommand::format() const{
	return "PRINT " + this->data->format();
}

LetVar::LetVar(std::string VAR, NumericExpression* NEXP){
	this->name = VAR;
	this->next = NEXP;
}

std::string LetVar::format() const{
	return "LET " + this->name + " " + this->next->format();
}

LetVarArray::LetVarArray(Array* _array, NumericExpression *_index, NumericExpression *_value){
	this->array = _array;
	this->index = _index;
	this->value = _value;
}

std::string LetVarArray::format() const{
	return "LET " + this->array->formatArray(this->index) + " " + this->index->format() + " " + this->value->format();
}

Goto::Goto(NumericExpression* JLINE){
	this->line = JLINE;	
}

std::string Goto::format() const{
	return "GOTO " + this->line->format();
}

IfGoto::IfGoto(NumericExpression* BEXP, NumericExpression* NEXP){
	this->booleanExp = BEXP;
	this->line = NEXP;	
}

std::string IfGoto::format() const{
	return "IF " + this->booleanExp->format() + " THEN " + this->line->format();
}

GoSub::GoSub(NumericExpression* NEXP){
	this->line = NEXP;
}

std::string GoSub::format() const{
	return "GOSUB " + this->line->format();
}

std::string ReturnCommand::format() const{
	return "RETURN";
}

std::string EndCommand::format() const{
	return "END";
}

