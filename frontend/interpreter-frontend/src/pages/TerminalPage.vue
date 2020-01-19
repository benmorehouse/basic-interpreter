<template>
	<div>
		<h1> Welcome to the basic interpreter </h1>
		<div id="terminal" />
	</div>
</template>

<script>
import Terminal from 'xterm'
export default {
	name: 'TerminalPage',
	props:{
		TerminalOut: String,
	},
	data(){
		return{
			terminal: Terminal,
		}
	},
	mounted(){
		this.$terminal = new Terminal();
		this.$terminal.open(document.getElementById('terminal'));
		this.$terminal.write("Hello world");
	}
}
</script>

<style scoped>
</style>
