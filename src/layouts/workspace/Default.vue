<script>
import AppBar from "./AppBar.vue"
import DefaultView from "./View.vue"


import {TUILogin} from "@tencentcloud/tui-core";
import {AppState, SDKAppID} from "@/main";

export default {
  data() {
    return {}
  },
  components: {AppBar, DefaultView},
  async mounted() {
    await TUILogin.login(
      {
        SDKAppID: SDKAppID,
        userID: AppState.user_id,
        userSig: AppState.sign_key,
        useUploadPlugin: true,
        framework: "vue3"
      }
    )
    TUILogin.setLogLevel(4)
  },
  unmounted() {
    TUILogin.logout()
  },

}
</script>

<template>
  <v-app>
    <AppBar></AppBar>
    <DefaultView></DefaultView>
  </v-app>

</template>

<style scoped>

</style>
