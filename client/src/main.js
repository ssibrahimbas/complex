import './plugins/axios'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from "@/plugins/axios";

createApp(App).use(axios).use(router).mount('#app')
