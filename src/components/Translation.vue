<script>
export default {
  data() {
    return {
      source: "abc",
      target: "",
    }
  }
}
</script>

<template>
  <v-container>
    <v-textarea label="源语言（中文）" v-model="source" variant="outlined">
    </v-textarea>

    <v-textarea label="目标语言（英语）" :model-value="target" readonly variant="outlined">
    </v-textarea>
    <v-btn color="primary" text="翻译" class="align-center text-center mb-5 mx-auto" block></v-btn>
  </v-container>
</template>

<style scoped>

</style>
