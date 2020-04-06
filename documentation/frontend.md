# the frontend

I had a huge focus on making the frontend look great, but also ensure that i didnt spend a lot of time doing so. 

After spending a weekend look into a few frameworks, I realized how limited I planned on the pages being: there we going to 
be at most 5 pages, all of which resembled a very similar stucture.

At this point, I was much more eager to get to the backend code, so a considerably short amount of time was put into making simple html 
files that all used materialize css. 

I also knew that I was going to take advantage of Go's templating engine, which is indicated by the files having the suffix .gohtml.

After some quick html editing, the next step was writing out the javascript functionality needed to make the frontend actually 
speak with the server. 

To be short: the fetch api for handling AJAX requests was used all throughout. 

Another thing I did was create my own sort of state variables within my javascript applications.
This would then call a componentDidLoad() method I created (though, it's not a component :P), which 
will use the templating engine to get the endpoint names from the configuration file.

# Terminal and TextEditor endpoints

These two endpoints rendered the terminal and the text editor respectively. These pages specifically used the
cdn for xtermjs. I had a considerable amount of issues with using node modules to import xtermjs locally, 
and with a time frame in mind, decided to instead take a small hit in page loading in favor of 
getting the last bits of the frontend finished.

[Click here to read more of the code that created the frontend.](../server/scripts)
