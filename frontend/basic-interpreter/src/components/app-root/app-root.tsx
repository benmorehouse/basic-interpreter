import { Component, h } from '@stencil/core';

@Component({
  tag: 'app-root',
  styleUrl: 'app-root.css'
})
export class AppRoot {

  render() {
    return (	
	<div>
		<ion-app>
			<ion-router useHash={false}>
				<ion-route url="/" component="main-aboutpage" />
				<ion-route url="/terminal" component="main-terminal" />
				<ion-route url="/loginPage" component="main-loginpage" />
			</ion-router>
			<main-frame></main-frame>
			<ion-nav animated={false}/>
		</ion-app>
	</div>
    );
  }
}
