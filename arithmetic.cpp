#include <iostream>
#include "arithmetic.h"

using namespace std;

double NumericExpression::getFinalValue(){ // we will use this to check for dividing by zero
	switch (this->operand) {
		case NULL_OPERAND:
			return this->value;
		case ADDITION_OPERAND:
			return this->left->getFinalValue() + this->right->getFinalValue();
		case SUBTRACTION_OPERAND:
			return this->left->getFinalValue() - this->right->getFinalValue();
		case DIVISION_OPERAND:
			return this->left->getFinalValue() / this->right->getFinalValue();
		case MULTIPLICATION_OPERAND:
			return this->left->getFinalValue() * this->right->getFinalValue();
		case LESS_OPERAND:
			return this->left->getFinalValue() < this->right->getFinalValue();
		case GREAT_OPERAND:
			return this->left->getFinalValue() > this->right->getFinalValue();
		case EQUAL_OPERAND:
			return this->left->getFinalValue() == this->right->getFinalValue();
	}
	return this->value;
}

void NumericExpression::setLeft(NumericExpression* Left){
	this->left = Left;
}

void NumericExpression::setRight(NumericExpression* Right){
	this->right = Right;
}

NumericExpression* NumericExpression::getLeft(){
	if(this->left == NULL){
		return NULL;
	}
	return this->left;	
}

NumericExpression* NumericExpression::getRight(){
	if(this->right == NULL){
		return NULL;
	}
	return this->right;
}

AdditionExpression::AdditionExpression(NumericExpression* _left, NumericExpression* _right){
	this->left = _left;
	this->right = _right;
	this->operand = ADDITION_OPERAND;
}

AdditionExpression::~AdditionExpression() {
	delete this->left;
	delete this->right;
}

std::string AdditionExpression::format() const {
	return this->NumericExpressionHelper("+");
}

SubtractionExpression::SubtractionExpression(NumericExpression* _left, NumericExpression* _right){
	this->left = _left;
	this->right = _right;
	this->operand = SUBTRACTION_OPERAND;
}

SubtractionExpression::~SubtractionExpression() {
	delete this->left;
	delete this->right;
}

std::string SubtractionExpression::format() const {
	return this->NumericExpressionHelper("-");
}

MultiplicationExpression::MultiplicationExpression(NumericExpression* _left, NumericExpression* _right){
	this->left = _left;
	this->right = _right;
	this->operand = MULTIPLICATION_OPERAND;
}

MultiplicationExpression::~MultiplicationExpression() {
	delete this->left;
	delete this->right;
}

std::string MultiplicationExpression::format() const {
    return this->NumericExpressionHelper("*");
}

DivisionExpression::DivisionExpression(NumericExpression* _left, NumericExpression* _right){
    this->left = _left;
    this->right = _right;
    this->operand = DIVISION_OPERAND;
}

std::string DivisionExpression::format() const {
     return this->NumericExpressionHelper("/");
}

std::string NumericExpression::NumericExpressionHelper(std::string operation) const {
    return "(" + this->left->format() + " " +  operation + " " + this->right->format() + ")";
}

Variable::Variable(double _value, std::string _name) {
    this->left = NULL;
    this->right = NULL;
    this->operand = NULL_OPERAND;
    this->value = _value;
    this->name = _name;
}

Variable::~Variable(){}

void Variable::changeVariable(double input){
    this->value = input;
}

double Variable::getValue() const{ // this is so you can pretty print this
    return this->value;
}

std::string Variable::getName() const{ // this is so you can pretty print this
    return this->name;
}

std::string Variable::format() const{
    return this->getName();
}

/**********************************************************/

Array::Array(std::string _name, NumericExpression* NEXP){ // pass in name of array to start
	this->name = _name;
	this->pushArray(NEXP); // the initial input
}

void Array::pushArray(NumericExpression* NEXP){
	this->array.push_back(NEXP);
}

NumericExpression* Array::getValueAtIndex(int index){
	if(index >= this->array.size()){
		return NULL;
	}
	
	if(index < 0){
		return NULL;
	}
	
	return this->array[index];
}

std::string Array::formatArray(NumericExpression* INDEX) const{
    return this->name + "[" + INDEX->format() + "]";
}

/*
//std::string Array::format() const{
    //if (this->left->getName() == ""){
    //  return this->name + "[" + std::to_string(this->left->getValue()) + "]";
    //}else{
    //  return this->name + "[" + this->left->format() + "]";
    //}
//}
*/
