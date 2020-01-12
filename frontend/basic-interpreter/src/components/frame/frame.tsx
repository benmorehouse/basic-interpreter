import { Component, Host, h } from '@stencil/core';

@Component({
  tag: 'main-frame',
  styleUrl: 'frame.css',
  shadow: true
})
export class Frame {

  render() {
    return (
      <Host>
        <slot> 
		<div class="top-bar">
			<div class="button-container">
				<ion-button href="/terminal"> Terminal page </ion-button>
				<ion-button href="/"> About page </ion-button>
				<ion-button href="/loginPage"> Login page </ion-button>
			</div>
		</div>
	</slot>
      </Host>
    );
  }

}
