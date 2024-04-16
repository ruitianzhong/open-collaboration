<script>
import AppBar from "./AppBar.vue"
import DefaultView from "./View.vue"


import {genTestUserSig} from "@/TUIKit";
import {TUILogin} from "@tencentcloud/tui-core";

export default {
  data() {
    return {
      SDKAppID: 1400741649,
      secretKey: "db02d702e820e031b1de37875da94f200d47b33e417dc778627741f366d49897",
      userID: "root"
    }
  },
  components: {AppBar, DefaultView},
  mounted() {

    TUILogin.login(
      {
        SDKAppID: this.SDKAppID,
        userID: this.userID,
        userSig: genTestUserSig({
          SDKAppID: this.SDKAppID,
          secretKey: this.secretKey,
          userID: this.userID
        }).userSig,
        useUploadPlugin: true,
        framework: "vue3"
      }
    )
    TUILogin.setLogLevel(4)


  },
  unmounted() {
    TUILogin.logout()
  }
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
