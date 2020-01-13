import "node_modules/xterm/lib/xterm.js";

export function sayHello() {
  return Math.random() < 0.5 ? 'Hello' : 'Hola';
}

export function terminalWriter(){
	var term = new Terminal();
	term.open(document.getElementById('terminal'));
	term.write('Hello from \x1B[1;3;31mxterm.js\x1B[0m $ ')
}
   

