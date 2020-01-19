<template>
	<div>
		<md-field id="login">
			<h1> Login </h1>
			<br />
			<br />     
			<md-input/>
			<br />
			<br />
			<input type="password" name="password" placeholder="Password"/>
			<br />
			<br />
			<input type="submit" name="submit-button" />
			<br />
			<br />
		</md-field>

		<p> Or if you are new: create an account! </p>

		<form id="signup">
			<input type="text" name="user-name" placeholder="Username"/>
			<input type="password" name="password" placeholder="Password"/>
			<input type="password" name="confirm-password" placeholder="Confirm Password"/>
			<input type="submit" name="submit-button" />
		</form>
	</div>
</template>

<script>
import Vue from 'vue'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'
import {MdField, MdInput} from 'vue-material/dist/components'
Vue.use(MdField, MdInput,)
// for icons we should use the fingerprint icon, code icon and, menu icon https://material.io/resources/icons/?style=baseline
export default {
	name: 'LoginPage',
}
</script>

<style scoped>
</style>
