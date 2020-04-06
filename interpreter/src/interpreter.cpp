#include "interpreter.h"

Interpreter::Interpreter(std::ifstream& in){
	this->ofileName = "interpreterOutput.txt";
	this->parse(in);
}

Interpreter::~Interpreter(){}

/*
 * used to load the string so we can push to parse numeric Expression
 *
*/

NumericExpression* Interpreter::parseNumericExpression(std::string input, int &position){
	std::cout<<"test"<<std::endl;
	trimWhiteSpace(input, position);
	if(input[position] >= '0' && input[position] <= '9'){
		return parseConstant(input, position);
	}else if(input[position] == '('){ // must be a binary operator at this point 
		position++;;
		trimWhiteSpace(input, position);
		NumericExpression *left = parseNumericExpression(input,position);
		char operand = input[position];
		position++;
		trimWhiteSpace(input, position);
		NumericExpression *right = parseNumericExpression(input,position);
		if(operand=='+'){
			AdditionExpression *addVal = new AdditionExpression(left,right);
			return addVal;
		}else if(operand == '-'){
			SubtractionExpression *subVal = new SubtractionExpression(left,right);
			return subVal;
		}else if(operand == '/'){
			DivisionExpression *divVal = new DivisionExpression(left,right);
			return divVal;
		}else if(operand == '*'){
			MultiplicationExpression *multVal = new MultiplicationExpression(left,right);
			return multVal;
		}else if(operand == '='){
			EqualityExpression *equal = new EqualityExpression(left, right);
			return equal;
		}else if(operand == '<'){
			LessExpression *less = new LessExpression(left, right);
			return less;
		}else if(operand == '>'){
			GreaterExpression *great = new GreaterExpression(left, right);
			return great;
		}else{
			return NULL;
		}
	}else{ // we parse a variable name
		std::string name = parseVariableName(input,position);
		trimWhiteSpace(input, position);
		if(position >= input.length()){ // we are at end of line ->prevents seg faultj
			Variable *var = new Variable(0,name); 	
			return var;
		}else if(input[position] == '['){
			while(input[position] == ' ' && position < input.length()){position++;}
			NumericExpression *index = parseNumericExpression(input,position);
			trimWhiteSpace(input, position);
			Array *array = new Array(name, index);
			return array;
		}else{
			Variable *var = new Variable(0,name); 	
			return var;
		}
	}
}


std::string Interpreter::loadString(std::stringstream& ss){
	std::string output ="";
	std::string buffer ="";
	while(ss >> buffer){
		if(ss.fail()){
			return "error in loadstring";
		}
		output += buffer + ' ';
	}
	return output;
}

// This is the main function used to call the initial interpretation.
void Interpreter::parse(std::ifstream& in) {
	std::string buffer;  // we read in from the filestream to this string
	std::vector<std::string> finalOutput;
	std::stringstream stream;
	while (std::getline(in, buffer)) {
		int line_number_buffer;
		stream.clear();
		stream.str(""); // this will set stringstream to empty
		stream.str(buffer); // this will set stringstream as buffer
		stream >> line_number_buffer;
		if (stream.fail()){
			std::cout<<"Stream failed when reading line number"<<std::endl;
			continue;
		}
		std::string line_number = std::to_string(line_number_buffer);
		
		std::string command;
		stream >> command;
		
		if (stream.fail()){
			std::cout<<"Stream failed when reading initial command"<<std::endl;
			continue;
		}

		int decider = initialCheck(command);
		
		std::string output = this->loadString(stream); 
		
		switch (decider)
		{
			case -1: // case couldnt be found
			{	
				this->commands.push_back(line_number + " Command Not Found");
				break;
			}
			case 0: // print command
			{	
				int position = 0;
				PrintCommand *print = new PrintCommand(this->parseNumericExpression(output, position));
				this->commands.push_back(line_number + " " + print->format());
				break;
			}
			case 1: // let command for either let array or let var
			{	
				// we will only do let var for right now
				// comes in the form LET HAN 5
				std::string name = "";
				std::string NEXP = "";
				int i =0;
				
				while(i < output.length() && output[i] != ' '){
					name+=output[i];
					i++;
				}
				i++;
				while(i < output.length() && output[i] != ' '){
					NEXP += output[i];
					i++;
				}

				int position = 0;
				LetVar *var = new LetVar(name, this->parseNumericExpression(NEXP, position));	
				this->commands.push_back(line_number + " " + var->format());
				break;
			}
			
			case 2: // this is a goto command
			{	
				int position = 0;
				Goto *goline = new Goto(this->parseNumericExpression(output, position));
				this->commands.push_back(line_number + " " + goline->format());
				break;
			}
				//goto function
			case 3:// this is the if statement command
			{	
				std::string BEXP = "";
				std::string NEXP = "";
				int i = 0;
				this->trimWhiteSpace(output,i);
				if(output[i] != '('){
					BEXP+='(';
				}
				while(i < output.length()){
					if((output[i] == 'T' || output[i] == 't') && (i+3 < output.length())){
						if(output[i+1] != 'H' && output[i+1] != 'h') continue;
						if(output[i+2] != 'E' && output[i+2] != 'e') continue;
						if(output[i+3] != 'N' && output[i+3] != 'n') continue;
						i = i+4;	
						break;
					}
					BEXP+=output[i];
					i++;
				}
				
				while(i <= output.length()){
					NEXP += output[i];
					i++;
				}

				int position1 = 0;
				int position2 = 0;
				IfGoto *ifgo = new IfGoto(this->parseNumericExpression(BEXP,position1), this->parseNumericExpression(NEXP,position2));
				this->commands.push_back(line_number + " " + ifgo->format());
				break;
				//binary expression using the IfGoto class
			}
			case 4:
			{	
				int position = 0;
				GoSub *gosub = new GoSub(this->parseNumericExpression(output,position));
				this->commands.push_back(line_number + " " + gosub->format());
				break;
			}
				// GoSub 
			case 5:
			{	
				ReturnCommand *retCom = new ReturnCommand();
				this->commands.push_back(line_number + " " + retCom->format());
				break;
				//return using gosub
			}
			case 6:
			{
				EndCommand *endCom = new EndCommand();
				this->commands.push_back(line_number + " " + endCom->format());
				break;
			}
			default:
			{	
				std::cout<<"DEVELOPER ERROR: Decider is broken"<<std::endl;
				break;
			}
		}
		buffer = "";
		// at this point we should use this->write() to append to the new file	
	}
	this->write();
}
/* 
 * used to write to the outfile so that the compiler can take the data and run with it 
 */

void Interpreter::write(){
	this->ofile.open(this->ofileName);
	std::vector<std::string>::iterator line;
	for(line = this->commands.begin(); line != this->commands.end(); ++line){
		this->ofile  << *line + '\n';
	}
	this->ofile.close();
	return;
}

/*
 * Parse Numeric Expression will be used to give us the command on the line that we want... the interpreter will take this
 * data and then use it to push the vector of commands to the compiler
*/



std::string Interpreter::parseVariableName(std::string input, int &position){
	std::string output = "";
	this->trimWhiteSpace(input, position);
	while(input[position] != ' ' && position <= input.length()){
		output += input[position];
		position++;
	}
	this->trimWhiteSpace(input, position);
	return output;
}

Variable* Interpreter::parseConstant(std::string input, int &position){
	std::string output = "";
	while(input[position] >= '0' && input[position] <= '9'){
		output+=input[position];
		position++;
	}
	std::string::size_type sz;     // alias of size_t
	Variable *var = new Variable(std::stod(output,&sz),output);
	trimWhiteSpace(input, position);
	return var;
}
/*
 * this is used in parse numeric expression -> trims whitespace
 */
void Interpreter::trimWhiteSpace(std::string input, int &position){
	while(input[position] == ' ' && position < input.length()){position++;}
	return;
}
/*
 * this is used for the switch case in main parse
 */
int Interpreter::initialCheck(std::string input){
	std::string lib[7] = {"PRINT","LET","GOTO","IF","GOSUB","RETURN","END"};
	for(int i=0;i<7;i++){
		if (input==lib[i]){
			return i;
		}
	}
	return -1;
}

std::string Interpreter::getOutFile() {
	return this->ofileName;
}


/*
 * This function is used to take in the vector we generate in parse
 * and return the string that we will write
 */
