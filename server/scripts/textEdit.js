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

// textEditorContent holds a string array of all the lines of the text editor
let textEditorContent = [""]; 

// these cursor variables keep track of key strokes and hold where the position of the cursor is
// these are used to add and subtract text from the textEditorContent.
// these are the indexes.
let cursorX = 0;
let cursorY = 0;

function componentDidLoad(event, term) {
	runTextEditor(term);
	if (event.Body == undefined) {
		return;	
	}

	terminalPageRedirect = event.Body.TerminalPageRedirect;
	compileEndpoint = event.Body.CompileEndpoint;
	saveEndpoint = event.Body.SaveEndpoint;

	if (event.Body.File != undefined && event.Body.File.Content != undefined) {
		// then write in the content to the frontend
		textEditorContent = event.Body.File.Content;
		textEditorContent.forEach(line => {
			line.forEach(byte => {
				term.write(byte);
			})
			term.writeln();
		})
	}
}

function terminalPageRedirectPressed() {

	window.location.replace("http://localhost:2272/" + terminalPageRedirect)
}

function compilePressed() {
	// take the terminal input and send to the frontend;
	let requestBody = JSON.stringify({
		FileContent: textEditorContent,
		IsBasic: isBasic, 
	})

	console.log(requestBody);

	fetch(compileEndpoint, {  
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		console.log(data);
	}).catch(err => {
		console.log(err)
	});
}

function savePressed() {

	let requestBody = JSON.stringify({
		Filename: fileName,
		Filepath: filePath,
	})

	fetch(SaveEndpoint, {  
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		console.log(data);
	}).catch(err => {
		console.log(err)
	});
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
	    const printable = !e.domEvent.altKey && !e.domEvent.altGraphKey && !e.domEvent.ctrlKey && !e.domEvent.metaKey && handleArrowKey(e.domEvent.keyCode);
	    if (e.domEvent.keyCode === 13) {
	   	appendLine(term) 
		textEditorContent.push("");
		cursorY = textEditorContent.length - 1;
		cursorX = 0;
	    } else if (e.domEvent.keyCode === 8) {
	    	if (term._core.buffer.x > makeSpacer().length + lineNumber.toString().length) {
			term.write('\b \b');
			textEditorContent[cursorY] = textEditorContent[cursorY].slice(0, -1);
			cursorX = cursorX - 1;
		} else {
			if (lineNumber > 1) {
				deleteLine(term);
				textEditorContent = textEditorContent.splice(1, cursorY);
				cursorY = cursorY - 1;
				cursorX = 0;
			}
		}
		// this needs to also look for cursor movements and delete wherever it is supposed to be
	    } else if (printable) {
		cursorX = cursorX + 1;
		textEditorContent[cursorY] += e.key.toString();
		term.write(e.key);
	    }
	});
}

// will return whether or not the keycode would put the cursor out of range.
// If a keycode that is not one of the four is asked of, it will just return true.
function handleArrowKey(keyCode) {
	console.log(keyCode);
	console.log(keyCode);
	console.log(keyCode);
	console.log(keyCode);
	let down = 40;
	let up = 38;
	let right = 39;
	let left = 37;
	switch (keyCode) {
		case down:
			if (cursorY == lineNumber - 1) {
				return false;
			}

			cursorY = cursorY + 1;
			break;
		case up: 
			if (cursorY == 0) {
				return false;
			}

			cursorY = cursorY - 1;
			break;
		case left: 
			if (cursorX == 0) {
				return false;
			}
			
			cursorX = cursorX - 1;
			break;
		case right: 
			if (cursorX >= textEditorContent[cursorY].length - 1) {
				return false
			}

			cursorX = cursorX + 1;
			break;
		default:
			return true;
			break;
	}

	return true;
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
