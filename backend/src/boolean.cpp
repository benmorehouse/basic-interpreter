#include "iostream"
#include "boolean.h"
#include "arithmetic.h"
#include <string>
/*
BooleanExpression::BooleanExpression(NumericExpression* _left, NumericExpression *_right){
	if ((_left == NULL) || (_right == NULL)){
		return;//what do we do if there was some sort of error in the input
	}else{
		this->left = _left;
		this->right = _right;
		this->isTrue = false;
	}	
}

std::string BooleanExpression::format() const{
	return "";
}
*/

std::string NumericExpression::BooleanExpressionHelper(std::string operation) const {
	return this->left->format() + " " + operation + " " + this->right->format();
}

EqualityExpression::EqualityExpression(NumericExpression *_left, NumericExpression *_right){
	this->left = _left;
	this->right = _right;
	this->operand = EQUAL_OPERAND;
}

std::string EqualityExpression::format() const{
	return this->BooleanExpressionHelper("=");
}

GreaterExpression::GreaterExpression(NumericExpression *_left, NumericExpression *_right){
	this->left = _left;
	this->right = _right;
	this->operand = GREAT_OPERAND;
}

std::string GreaterExpression::format() const{
	return this->BooleanExpressionHelper(">");
}

LessExpression::LessExpression(NumericExpression *_left, NumericExpression *_right){
	this->left = _left;
	this->right = _right;
	this->operand = LESS_OPERAND;
}

std::string LessExpression::format() const{
	return this->BooleanExpressionHelper("<");
}

