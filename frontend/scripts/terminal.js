'use strict'

var currentTerminalInput;
var currentTerminalDirectory; // not sure if this is needed... need to see how the backend responds.

function runBasicTerminal(terminalProcessEndoint) {
	console.log("running the terminal")	
	if (term._initialized) {
	    return;
	}

	term._initialized = true;

	term.prompt = () => {
	    term.write('\r\n$ ');
	};

	term.writeln('Welcome to your Basic Interpreter Instance!');
	term.writeln('');
	prompt(term);

	term.onKey(e => {
	    const printable = !e.domEvent.altKey && !e.domEvent.altGraphKey && !e.domEvent.ctrlKey && !e.domEvent.metaKey;

	    if (e.domEvent.keyCode === 13) {
		// this is when they hit enter. we then run an endpoint which pushes to go server, which pushes to operating system.
		requestBody = JSON.stringify({
			CurrentTerminalInput: currentTerminalInput,
		})

		fetch(terminalProcessEndoint, {  
			method:"POST",
			credentials:"include",
			body: requestBody,		
		}).then(res => {
			return res.json();
		}).then(data => {
			console.log(data);	
			//console.log(e.domEvent.keyCode)
			//prompt(term); term instead will be the results from the server!
		}).catch(err => {
			console.log(err)
		});

	    } else if (e.domEvent.keyCode === 8) {
		if (term._core.buffer.x > 2) {
		    term.write('\b \b');
		}
	    } else if (printable) {
		currentTerminalInput = currentTerminalInput + e.key;
		term.write(e.key);
	    }
	});
}

function prompt(term) {
	term.write('\r\n$ ');
}
