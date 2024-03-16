/**
 * main.js
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Plugins
import {registerPlugins} from '@/plugins'

// Components
import App from './App.vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';


// Composables
import {createApp} from 'vue'
import VueMathjax from 'vue-mathjax-next'

const app = createApp(App)

registerPlugins(app)

app.use(Antd).use(VueMathjax).mount('#app')
