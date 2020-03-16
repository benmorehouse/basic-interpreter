'use strict'

var currentTerminalInput = "";
var currentTerminalDirectory = ""; // not sure if this is needed... need to see how the backend responds.

function runBasicTerminal(terminalProcessEndoint) {
	console.log("running the terminal")	
	console.log(terminalProcessEndoint)	
	if (term._initialized) {
	    return;
	}

	term._initialized = true;

	term.prompt = () => {
	    term.write('\r\n$ ');
	};

	term.writeln('Welcome to your Basic Interpreter Instance!');
	term.writeln('');
	prompt(term, "");

	term.onKey(e => {
	    const printable = !e.domEvent.altKey && !e.domEvent.altGraphKey && !e.domEvent.ctrlKey && !e.domEvent.metaKey;

	    if (e.domEvent.keyCode === 13) {
		// this is when they hit enter. we then run an endpoint which pushes to go server, which pushes to operating system.
		let requestBody = JSON.stringify({
			Command: currentTerminalInput,
		})
		
		console.log(currentTerminalInput)
		fetch(terminalProcessEndoint, {  
			method:"POST",
			credentials:"include",
			body: requestBody,		
		}).then(res => {
			return res.json();
		}).then(data => {
			console.log(data);	
			if (!data.Success) {
				console.log('there was an error with the terminal')
				prompt(term, "");
				return
			} else {
				console.log('here we are passing something back!')
				prompt(term, data.Message);	
			}
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

function prompt(term, output) {
	if (output == "") {
		term.write('\r\n' + currentTerminalDirectory + '$ ');
		return
	}
	
	term.write('\r\n' + currentTerminalDirectory + '$ ' + output);
}
