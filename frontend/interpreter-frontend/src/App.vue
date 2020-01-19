<template>
	<div>
		<Root />
	</div>
</template>


<script>
import Root from  './components/Root.vue'
import MainPage from './pages/MainPage.vue'
export default {
	name: 'app',
	components: {
		Root,
	},
	beforeMount(){
		this.$router.push({path: '/about', component: MainPage})
	}
}

</script>

<style scoped>
</style>

