var term = new Terminal({
		theme: {
			background: '#004671'
		}
	});
// the color for the font is #54a0cb		
term.open(document.getElementById('terminal'));

function runBasicTerminal() {
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
		console.log(e.domEvent.keyCode)
		prompt(term);
	    } else if (e.domEvent.keyCode === 8) {
		// Do not delete the prompt
		if (term._core.buffer.x > 2) {
		    term.write('\b \b');
		}
	    } else if (printable) {
		term.write(e.key);
	    }
	});
}

function prompt(term) {
	term.write('\r\n$ ');
}
runBasicTerminal();
