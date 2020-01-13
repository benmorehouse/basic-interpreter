import { Component, Host, h } from '@stencil/core';
import '../../../node_modules/xterm/lib/xterm.js';
import {Terminal} from 'xterm';

@Component({
	tag: 'main-terminal',
	styleUrl: '../../../node_modules/xterm/css/xterm.css',
	shadow: true,
})

export class Terminal {

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
