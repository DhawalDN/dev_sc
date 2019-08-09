import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from '@/router/router'


Vue.config.productionTip = false

Vue.use(VueAxios, axios)

var vm = new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
window.app = vm