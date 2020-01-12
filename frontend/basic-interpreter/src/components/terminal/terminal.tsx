import { Component, Host, h } from '@stencil/core';
import '../../../node_modules/xterm/lib/xterm.js';

@Component({
  tag: 'main-terminal',
  styleUrl: '../../../node_modules/xterm/css/xterm.css',
  shadow: true,
})
export class Terminal {
	var term = new Terminal();
	term.open(document.getElementById('terminal'));
	term.write('Hello from \x1B[1;3;31mxterm.js\x1B[0m $ ')
   
 render() {
    return (
		<Host>
			<slot>
				<div class="terminal"></div>
				<p> You are now at the terminal </p>	
			</slot>
		</Host>

    );
  }

}
