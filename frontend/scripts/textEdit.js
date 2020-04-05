'use strict'

let errorMessage = "";

let terminalPageRedirect = "";
let compileEndpoint = "";
let saveEndpoint = "";

let fileName = "";
let filePath = "";
let isBasic = false;
let textEditorStarted = false;
let content = "";

function componentDidLoad(event, term) {
	console.log(event)
	runTextEditor(term);
	if (event.File != undefined) {
		
	}
}

function compilePresset() {

}

function savePressed() {

}

let lineNumber = 1;

function runTextEditor(term) {
	if (term._initialized) {
	    return;
	}
	term._initialized = true;
	
	if (!textEditorStarted) {
		term.write(lineNumber.toString() + makeSpacer());
		textEditorStarted = true;
	}
	
	term.onKey(e => {
	    const printable = !e.domEvent.altKey && !e.domEvent.altGraphKey && !e.domEvent.ctrlKey && !e.domEvent.metaKey;
	    if (e.domEvent.keyCode === 13) {
	   	appendLine(term) 
	    } else if (e.domEvent.keyCode === 8) {
		if (term._core.buffer.x > makeSpacer().length + lineNumber.toString().length) {
			term.write('\b \b');
		} else {
			deleteLine(term);
		}
	    } else if (printable) {
		term.write(e.key);
	    }
	});
}

function makeSpacer() {
	let spacer = "";
	if (lineNumber < 10) {
		spacer = '     ';
	} else if (lineNumber < 100) {
		spacer = '    ';
	} else if (lineNumber < 1000) {
		spacer = '   ';
	} else if (lineNumber < 10000) {
		spacer = '  ';
	} else {
		spacer = ' ';
	}

	return spacer;
}

function appendLine(term) {
	
	lineNumber = lineNumber + 1;
	term.writeln("");
	console.log(lineNumber);
	term.write(lineNumber.toString() + makeSpacer());
}

function deleteLine(term) {

	if (lineNumber == 1) {
		return;
	}

	lineNumber = lineNumber - 1;
	term.write('\x1b[2K\x1b[1A')  
}

function fillFile(content) {
	return
}
