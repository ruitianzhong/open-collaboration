<script>
import AppBar from "./AppBar.vue"
import DefaultView from "./View.vue"

export default {
  components: {AppBar, DefaultView},
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
