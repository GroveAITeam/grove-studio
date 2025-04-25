import {createApp} from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import './style.css';
import router from './router/index'
import { toast } from './plugins/toast'
import { initGlobalToast } from './utils/toast'

const pinia = createPinia()
const app = createApp(App)
app.use(pinia).use(router).use(toast)

app.mount('#app')

// 初始化全局 toast
initGlobalToast()
