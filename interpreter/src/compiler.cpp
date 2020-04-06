#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <fstream>
#include <map>
#include "compiler.h"
#include "interpreter.h"

/*
 * The purpose of this function is to get all text after the line number and the command
*/

void Compiler::main(){
	if (!this->parsefile()){
		return; // this means that we couldnt access ifile
	}
	//at this point we want to go through and start to act the file out 
	int cursor = 0;
	bool errorFound = false; // errorfound will be true if an error is written
	
	this->ofile.open(this->ofileName);
	this->ofile.clear();
	
	while(!errorFound && cursor != -1){ 
		if(cursor >= this->lineNumbers.size()){
			errorFound = true;
			continue;
		}else if(cursor == this->lineNumbers.size() - 1){
			if(this->getCommand(cursor) == "ERROR"){
				errorFound = true;
				continue;
			}else if(this->getCommand(cursor) == "END"){
				break;	
			}else{
				std::cout << this->getErrorMessage(this->lineNumbers[cursor],"No end command at the end of the program")<<std::endl;
				errorFound = true;
				continue;
			}
		}else if(this->getCommand(cursor) == "END"){
			break;
		}else if(this->getCommand(cursor) == "PRINT"){
			int i = 0;
			NumericExpression* NEXP = this->interpreter->parseNumericExpression(this->getLine(cursor),i);
			if(this->variableValues.find(NEXP->format()) == this->variableValues.end()){
				std::cout << NEXP->getFinalValue()<<std::endl;	
			}else{
				std::cout<<this->variableValues[NEXP->format()]<<std::endl;
			}
		}else if(this->getCommand(cursor) == "LET"){
			std::stringstream ss;
			std::string temp = this->getLine(cursor);
			ss << temp; // pushes everything into ss
			std::string varName;
			ss >> varName; // variable name now
			temp = "";
			std::string buffer;
			while(ss >> buffer){
				temp += buffer + ' ';
			}
			if(this->variableValues.find(temp) == this->variableValues.end()){
				// this means we are assigning variable to something that doesnt already exist
				// ie LET HAN 5
				int i=0;
				this->variableValues[varName] = this->interpreter->parseNumericExpression(temp,i)->getFinalValue();
			}else{
				// ie LET HAN  [EXISTING VARIABLE]
				this->variableValues[varName] = this->variableValues[temp];
			}
		}else if(this->getCommand(cursor) == "GOTO"){ // goto NEXP
			int i =0;
			std::string line = this->getLine(cursor);
			NumericExpression* nexp = this->interpreter->parseNumericExpression(line,i);
			std::string temp = nexp->format();
			if(this->variableValues.find(temp) == this->variableValues.end()){
				// this means we are assigning variable to something that doesnt already exist
				double newline = nexp->getFinalValue();
				if(newline < 1){// throw an error to the user
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
					errorFound = true;
					continue;
				}else if(newline > this->lineNumbers.back()){
					errorFound = true;
					continue;
				}else{
					int mistakenNumber = this->lineNumbers[cursor];
					cursor = this->jumpToCommand(nexp->getFinalValue()); 
					if(cursor == -1){
						std::cout << this->getErrorMessage(mistakenNumber,"GOTO jump to non existant line");
					}
					continue;
				}
			}else{
				int newline = this->variableValues.find(temp)->second;
				if(newline < 1){// throw an error to the user
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
					errorFound = true;
					continue;
				}else if(newline > this->lineNumbers.back()){
					errorFound = true;
					continue;
				}else{
					int mistakenNumber = this->lineNumbers[cursor];
					cursor = this->jumpToCommand(newline); 
					if(cursor == -1){
						std::cout << this->getErrorMessage(mistakenNumber,"GOTO jump to non existant line");
					}
					continue;
				}
			}
		}else if(this->getCommand(cursor) == "GOSUB"){
			int i = 0;		
			std::string line = this->getLine(cursor);
			NumericExpression* nexp = this->interpreter->parseNumericExpression(line,i);
			std::string temp = nexp->format();
			if(this->variableValues.find(temp) == this->variableValues.end()){
				// this means we are assigning variable to something that doesnt already exist
				double newline = nexp->getFinalValue();
				if(newline < 1){// throw an error to the user
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
					errorFound = true;
					continue;
				}else if(newline > this->lineNumbers.back()){
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is out of index")<<std::endl;
					errorFound = true;
					continue;
				}else{
					this->goSubLines.push(cursor);	
					int mistakenNumber = this->lineNumbers[cursor];
					cursor = this->jumpToCommand(nexp->getFinalValue()); 
					if(cursor == -1){
						std::cout << this->getErrorMessage(mistakenNumber,"GOSUB jump to non existant line");
					}
					continue;
				}
			}else{
				int newline = this->variableValues.find(temp)->second;
				if(newline < 1){// throw an error to the user
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
					errorFound = true;
					continue;
				}else if(newline > this->lineNumbers.back()){
					std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is out of index")<<std::endl;;
					errorFound = true;
					continue;
				}else{
					this->goSubLines.push(cursor);	
					int mistakenNumber = this->lineNumbers[cursor];
					cursor = this->jumpToCommand(newline); 
					if(cursor == -1){
						std::cout << this->getErrorMessage(mistakenNumber,"IF jump to non existant line");
					}
					continue;
				}
			}
		}else if(this->getCommand(cursor) == "RETURN"){
			if(this->goSubLines.empty()){
				std::cout << this->getErrorMessage(this->lineNumbers[cursor],"No matching GOSUB for RETURN")<<std::endl;
				errorFound = true;
			}else{
				cursor = this->jumpToCommand(this->goSubLines.top()); 
				this->goSubLines.pop();
			}
		}else if(this->getCommand(cursor) == "IF"){
			std::string line = this->getLine(cursor);
			std::string buffer;
			std::stringstream ss;
			ss << line;
			std::string boolean = "(";
			std::string newLine = "";
			while(ss >> buffer){
				if(buffer == "THEN"){
					break;
				}else{
					boolean +=buffer;
				}
			}

			while(ss >> buffer){
				newLine += buffer;
			}
		
			int i = 0;		
			int j = 0;		
			NumericExpression* bexp = this->interpreter->parseNumericExpression(boolean,i);
			if(bexp == NULL){
				std::string left, operand, right;
				int place = 1;
				std::string buffer ="";
				while(boolean[place] != '<' && boolean[place] != '>' && boolean[place] != '='){
					buffer += boolean[place];	
					place++;
				}// buffer is the variable name we have in library possibly 
				
				if(this->variableValues.find(buffer) == this->variableValues.end()){ // couldnt find it 
					int x = 0;
					left = this->interpreter->parseNumericExpression(buffer,x)->format();
				}else{
					left = std::to_string(this->variableValues[buffer]);
				}
				
				operand = boolean[place];
				place++;

				buffer = "";
				while(boolean[place] == ' '){place++;}	
				
				while(place < boolean.length() && boolean[place] != ' '){
					buffer += boolean[place];	
					place++;
				}

				if(this->variableValues.find(buffer) == this->variableValues.end()){ // couldnt find it 
					int y = 0;
					right = this->interpreter->parseNumericExpression(buffer,y)->format();
				}else{
					right = std::to_string(this->variableValues[buffer]);
				}
				i = 0;
				bexp = this->interpreter->parseNumericExpression(left+operand+right,i);
			}

			NumericExpression* nexp = this->interpreter->parseNumericExpression(newLine,j);
			if(bexp->getFinalValue() == 1){	
				if(this->variableValues.find(nexp->format()) == this->variableValues.end()){
					// this means we are assigning variable to something that doesnt already exist
					double newline = nexp->getFinalValue();
					if(newline < 1){// throw an error to the user
						std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
						errorFound = true;
						continue;
					}else if(newline > this->lineNumbers.back()){
						std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is out of index")<<std::endl;
						errorFound = true;
						continue;
					}else{
						int mistakenNumber = this->lineNumbers[cursor];
						cursor = this->jumpToCommand(newline); 
						if(cursor == -1){
							std::cout << this->getErrorMessage(mistakenNumber,"IF jump to non existant line");
						}
						continue;
					}
				} else {
					int newline = this->variableValues.find(nexp->format())->second;
					if(newline < 1){// throw an error to the user
						std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is less than zero->doesnt exist")<<std::endl;
						errorFound = true;
						continue;
					}else if(newline > this->lineNumbers.back()){
						std::cout << this->getErrorMessage(this->lineNumbers[cursor],"Line is out of index");
						errorFound = true;
						continue;
					}else{
						int mistakenNumber = this->lineNumbers[cursor];
						cursor = this->jumpToCommand(nexp->getFinalValue()); 
						if(cursor == -1){
							std::cout << this->getErrorMessage(mistakenNumber,"IF jump to non existant line");
						}
						continue;
					}
				}
			}
		}
		cursor++;
	}
	this->ofile.close();
}
	
std::string Compiler::getLine(int cursor){
	std::stringstream ss;
	ss << this->commandList[this->lineNumbers[cursor]];// puts everything in there
	std::string buffer;
	ss >> buffer;	
	ss >> buffer;
	buffer = "";	
	std::string output = "";
	while(ss >> buffer){
		output += buffer + ' ';
	}
	return output;
}

bool Compiler::parsefile(){
	this->ifile.open(this->ifileName);
	if(!this->ifile.is_open()){
		return false;
	}
	this->ifile.seekg(0,std::ios::beg);
	this->ifile.clear();
	std::string buffer = "";
	int linenumber = 0;	
	std::stringstream ss;
	// have to pass in the pointer of ifile to meet requirements of getline

	while(std::getline(this->ifile, buffer)){ 
		ss << buffer;
		ss >> linenumber;
		if(ss.fail()){ return false;}
		this->lineNumbers.push_back(linenumber);
		this->commandList[linenumber] = buffer;
		ss.str("");
		ss.clear();
		buffer = "";
	}
	this->ifile.close();
	return true;
}

/*
	 * Within main, we need a way to go get which command we are doing. We will do something similar to the strings.Field() command in Golang to get the second field element ie the command
	 * this will return a string of it given the line input that we want 
*/

std::string Compiler::getCommand(int cursor){
	if(cursor < 0){
		return "ERROR";	
	}else if (cursor >= this->lineNumbers.size()){
		return "ERROR";	
	}else{
		std::string buffer;
		std::stringstream ss;
		ss << this->commandList[this->lineNumbers[cursor]];
		if(ss.fail()){ return "ERROR";}
		ss >> buffer;	// gets the line number
		if(ss.fail()){ return "ERROR";}
		buffer = "";
		ss >> buffer;	// gets the actual command
		if(ss.fail()){ return "ERROR";}
		return buffer;
	}
}

/*
 * jumptocommand takes value and finds it within the lineNumbers or negative one
 */

int Compiler::jumpToCommand(int nexpVal){
	std::vector<int>::iterator itr;
	int i = 0;
	
	for(itr = this->lineNumbers.begin();itr != this->lineNumbers.end(); ++itr){
		if(*itr == nexpVal){
			return i;
		}
		i++;
	}
	return -1;
}
