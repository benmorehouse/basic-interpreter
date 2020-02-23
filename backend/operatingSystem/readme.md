# Design tactic

commands that I want to be able to have:

- ls
- cd
- cd <directory>
- mkdir 
- rmdir
- touch <file>
- rm <file>
- rm -r <directory>
- open <file> 
- pwd

In order to do this, I am going to have a directory class, and each directory class can have
a pointer which points to another vector of directory classes.

The vector will internally be a trie, so that things are sorted correctly and efficiently.

The functionality itself will take in command line arguments. This will integrate well with the golang exec library.

# error handling

Going to have a top level class that will act as a guideline for the two subclasses for error handling.

There will then be individual classes for errorHandling for the basic interpreter and the operating system itself.

