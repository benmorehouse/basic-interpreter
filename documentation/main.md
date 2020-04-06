# The main goal

I wanted to dive into building a low project with C++, and decided that building
an interpreter for a programming language would be an awesome way to fulfill that. 

I stumbled upon BASIC, which personally reminded me a lot of python, but without nearly as much 
third party support and other importing capabilites.

I decided that the best way about solving this problem would be by having a very common syntax that the 
program would expect when it did the actual interpretation. You can read more about the documentation on that later.

The goal after building this interpreter was to have a great way for users to be able to interact with
it. At this point, I got inspired by some past work I had gotten to do and decided to build a server 
to host it in go. This explains why though some of the more complex things in this program were originally written
in c++, a bulk of it is written in go: the server was a very large task.


# Future

I am pleased with where the project is at the moment. 
The server now has the ability to: 

- authenticate and safely store user data on the server
- Hold entire directories full of .basic, .txt, and .md files for users. 
- Allow traversal all across these directories from the web page
- edit these files from the web page, while maintaining the ability to save and keep secure information


Some things i would really like to see happen in particular would be if we could:

- rewrite the entire interpreter in golang (this could be argued about whether or not this is a good spenditure of time.)
- take advantage of Amazon S3 for the file storage as well user data.
- host the website on either an nginx server or GCP
- setup a schduled script that would send out a digest to users of what files they edited that day



