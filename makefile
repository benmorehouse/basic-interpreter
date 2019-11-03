all: hw4

hw4: main.cpp arithmetic.o boolean.o command.o interpreter.o compiler.o
	g++ -g -Wall main.cpp arithmetic.o boolean.o command.o interpreter.o compiler.o -o main

arithmetic.o: arithmetic.h arithmetic.cpp
	g++ -g -Wall -c arithmetic.cpp

boolean.o: boolean.h arithmetic.h boolean.cpp
	g++ -g -Wall -c boolean.cpp

command.o: command.h boolean.h arithmetic.h command.cpp
	g++ -g -Wall -c command.cpp

interpreter.o: interpreter.h command.h boolean.h arithmetic.h interpreter.cpp
	g++ -g -Wall -c interpreter.cpp

compiler.o: compiler.h interpreter.h command.h boolean.h arithmetic.h compiler.cpp
	g++ -g -Wall -c compiler.cpp
