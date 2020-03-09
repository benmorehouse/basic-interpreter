# fqueue

To be straight to the point, the reason for having this file is to appeal to speed that is offered by using the operating system functions of 
c++. I ran into the interesting problem of how I can feed the c++ backend data the best way possible. And with a little bit of research, it became 
obvious that it plays best when you are allowing for it to use its previously optimized methods such as the std::ifstream and std::ofstream.

This directory serves as a queue for all files that are submitted to be interpreted by the computer. It shouldnt ever have more than one file at a time.
The frontend to the server should catch this if there is more than one file!
