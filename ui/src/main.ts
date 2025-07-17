import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

import './assets/base.css'
import './assets/main.css'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:9123'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
