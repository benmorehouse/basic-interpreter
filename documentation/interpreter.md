# Interpreter

The interpreter was broken up into two large process:

- interpreting => parsing the data that the user inputs. 
	 The stream of data is taken from a file that is generated through the server.
	 This then generates a new file: one that has syntax and format with which the the compiler can use.
	 
	 The interpreter is built with the idea of creating a parse tree in mind. The tree is created from
	 numeric expressions. These are labeled as NEXP, and stand for variables, arrays, operand statments, and binary
	 experssions. 

	 This then implements a recursive approach to creating the parse tree, and pays special mind
	 to things such as open and close brackets, as well as other enforced sytanx rules of the programming language.

- the compiler is organized using mainly the standard libraries that implement maps, sets, etc. for C++. At this point, 
  it merely just takes care of actually generating the output for the application.
  

This ends with a final output file being generated, which is then served to the frontend.

