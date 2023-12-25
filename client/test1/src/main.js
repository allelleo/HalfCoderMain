import { createApp } from 'vue'
import App from './App.vue'
import Mrouter from './router/index.js'
import store from './store/index.js'

createApp(App).use(store).use(Mrouter).mount('#app')
