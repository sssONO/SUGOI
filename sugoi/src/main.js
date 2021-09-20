//import Vue from 'vue'
import { createApp } from 'vue'
//import Home from './views/Home.vue'
import App from './App.vue'
import './index.css'
// ルーティングのために追加
import router from './router'

const app = createApp(App)
app.config.productionTip = false
app.use(router).mount('#app')

