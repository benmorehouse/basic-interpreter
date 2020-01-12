import { Component, Host, h } from '@stencil/core';

@Component({
  tag: 'main-loginpage',
  styleUrl: 'loginpage.css',
  shadow: true
})
export class Loginpage {

  render() {
    return (
      <Host>
        <slot>
		<p> login page </p>
	</slot>
      </Host>
    );
  }

}
