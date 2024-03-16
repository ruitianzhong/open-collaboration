<template>
  <div class="hello">
    <v-container>
      <v-textarea v-model="formula"></v-textarea>
    </v-container>
    <vue-mathjax :formula="formula"/>
  </div>
</template>

<script setup>
import {ref} from 'vue'

const formula = ref('$$x = {-b \\pm \\sqrt{b^2-4ac} \\over 2a}.$$')
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
textarea {
  min-height: 60px;
}
</style>
