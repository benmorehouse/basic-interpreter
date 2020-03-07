#ifndef ARITHMETIC_H
#define ARITHMETIC_H

#include <string>
#include <vector>
#include <sstream>

struct returnVal{ // used in the array class
	bool exists;
	double val;
};

enum OPERAND: int {
	NULL_OPERAND,
	ADDITION_OPERAND,
	MULTIPLICATION_OPERAND,
	SUBTRACTION_OPERAND,
	DIVISION_OPERAND,
	EQUAL_OPERAND,
	GREAT_OPERAND,
	LESS_OPERAND
};

class NumericExpression{
	public:
		virtual std::string format() const { return ""; }
		double getFinalValue();
		std::string NumericExpressionHelper(std::string operation) const;
		std::string BooleanExpressionHelper(std::string operation) const;
		
		void setLeft(NumericExpression* Left);
		void setRight(NumericExpression* Right);
		NumericExpression* getLeft();
		NumericExpression* getRight();

	protected:
		int operand; 
		double value;

		//double finalValue; // this should get the final value of the numeric expression so that we can use in the compiler
		NumericExpression* left;
		NumericExpression* right;

    //std::string NumericExpressionHelper(std::string) const;
};

class AdditionExpression : public NumericExpression {
	public:
		AdditionExpression(NumericExpression* _left, NumericExpression* _right);
		~AdditionExpression();
		std::string format() const; // essentially just used to print back to output correct so we can run
};

class SubtractionExpression : public NumericExpression {
	public:
		SubtractionExpression(NumericExpression* left, NumericExpression* right);
		~SubtractionExpression();
		std::string format() const; // essentially just used to print back to output correct so we can run
};

class MultiplicationExpression : public NumericExpression {
	public:
		MultiplicationExpression(NumericExpression* left, NumericExpression* right);
		~MultiplicationExpression();
		std::string format() const; // essentially just used to print back to output correct so we can run
};

class DivisionExpression : public NumericExpression {
	public:
		DivisionExpression(NumericExpression* left, NumericExpression* right);
		~DivisionExpression();
		std::string format() const; // essentially just used to print back to output correct so we can run
};

class Variable : public NumericExpression{
	public:
		Variable(double, std::string); // takes in val and string
		~Variable();
		void changeVariable(double);
		std::string getName() const; // const here means we wont change anything
		double getValue() const; // gets the value of the variable
		std::string format() const;
	private:
		std::string name;
};

class Array : public NumericExpression{ // in theory this is just an array of variables... so why inherit this whole class
	public:
		Array(std::string, NumericExpression*);
		~Array(){};
		std::string formatArray(NumericExpression*) const;
		void pushArray(NumericExpression*);
		NumericExpression* getValueAtIndex(int);

	private:
		std::vector<NumericExpression*> array;
		std::string name;
};

#endif
