#ifndef BOOLEAN_H
#define BOOLEAN_H

#include <string>
#include "arithmetic.h"
/*
class BooleanExpression{
public:	
	BooleanExpression(NumericExpression* _left, NumericExpression *_right);
	virtual std::string format() const;
	virtual ~BooleanExpression() {};

protected: 
	NumericExpression* left;
	NumericExpression* right;
	std::string BooleanExpressionHelper(std::string) const;
	bool isTrue;
};
*/

class EqualityExpression : public NumericExpression{
public: 
	EqualityExpression(NumericExpression *_left, NumericExpression *_right);
	~EqualityExpression(){};	
	std::string format() const;
};

class GreaterExpression : public NumericExpression{
public: 
	GreaterExpression(NumericExpression *_left, NumericExpression *_right);
	~GreaterExpression(){};	
	std::string format() const;
};


class LessExpression : public NumericExpression{
public: 
	LessExpression(NumericExpression *_left, NumericExpression *_right);
	~LessExpression(){};
	std::string format() const;
};

#endif

