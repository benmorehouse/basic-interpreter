import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import MainPage from  './pages/MainPage.vue'
import TerminalPage from  './pages/TerminalPage.vue'
import LoginPage from  './pages/LoginPage.vue'
import GithubPage from  './pages/GithubPage.vue'

Vue.config.productionTip = false
Vue.use(VueRouter)

const router = new VueRouter({
	mode: 'history',
	routes: [
		{path: '/about', component: MainPage},
		{path: '/terminal', component: TerminalPage},
		{path: '/login', component: LoginPage},
		{path: '/github', component: GithubPage},
	]
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
