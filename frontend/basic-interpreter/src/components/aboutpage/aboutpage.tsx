import { Component, Host, h } from '@stencil/core';

@Component({
  tag: 'main-aboutpage',
  styleUrl: 'aboutpage.css',
  shadow: true
})
export class Aboutpage {

  render() {
    return (
      <Host>
        <slot>
		<p> About page </p>
	</slot>
      </Host>
    );
  }

}
