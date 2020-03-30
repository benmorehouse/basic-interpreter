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
		if (currentTerminalInput.trim() == "") {
			currentTerminalInput = ""
			prompt(term, "")
			return 
		}

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
			} else {
				currentTerminalDirectory = data.CurrentDirectory
				if (data.Message == "clear") {
					clear(term);
				} else {
					prompt(term, data.Message);
				}
			}

			currentTerminalInput = ""
			//console.log(e.domEvent.keyCode)
			//prompt(term); term instead will be the results from the server!
		}).catch(err => {
			console.log(err)
		});
	    } else if (e.domEvent.keyCode === 8) {
		if (term._core.buffer.x > 2) {
			term.write('\b \b');
		}
		currentTerminalInput = currentTerminalInput.slice(0, -1);

	    } else if (printable) {
		currentTerminalInput = currentTerminalInput + e.key;
		term.write(e.key);
	    }
	});
}

function prompt(term, output) {
	
	if (output != "") {
		term.write("\r\n" + output + "\r\n");
	}
	
	term.write('\r\n' + currentTerminalDirectory + '$ ');
}

function clear(term) {
	
	let s = "";
	for (let i=0; i<25; i++) {
		s += '\r\n'
	}
	term.write(s + '\r\n' + currentTerminalDirectory + '$ ')
}
