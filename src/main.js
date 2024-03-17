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
import {TUIComponents, TUIChatKit, genTestUserSig} from "./TUIKit";
import {TUILogin} from "@tencentcloud/tui-core";


const app = createApp(App)
const SDKAppID = 1400741649;
const secretKey = "db02d702e820e031b1de37875da94f200d47b33e417dc778627741f366d49897";
const userID = "root";
registerPlugins(app)
TUIChatKit.components(TUIComponents, app);
TUIChatKit.init();
TUILogin.login(
  {
    SDKAppID,
    userID,
    userSig: genTestUserSig({
      SDKAppID,
      secretKey,
      userID
    }).userSig,
    useUploadPlugin: true,
    framework: "vue3"
  }
)

app.use(Antd).use(VueMathjax).mount('#app')
export {SDKAppID, secretKey}
