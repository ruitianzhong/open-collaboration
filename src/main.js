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
import {TUIComponents, TUIChatKit} from "./TUIKit";


const app = createApp(App)
const SDKAppID = 1400741649;
registerPlugins(app)
TUIChatKit.components(TUIComponents, app);
TUIChatKit.init();

app.use(Antd).use(VueMathjax).mount('#app')
export {SDKAppID}
export const AppState = {
  user_id: '',
  group_id: '',
  sign_key: 'empty'
}
