<script>
export default {
  data() {
    return {
      source: "abc",
      target: "",
      value: 'en'
    }
  },
  computed: {
    language() {
      switch (this.value) {
        case "en":
          return "（英语）"
        case "zh":
          return "（中文）"
        case "yue":
          return "（粤语）"

        default:
          return "（英语）"

      }
    }
  }
}
</script>

<template>
  <v-container>
    <v-textarea label="源语言" v-model="source" variant="outlined">
    </v-textarea>

    <v-textarea :label="'目标语言'+language" :model-value="target" readonly variant="outlined">
    </v-textarea>
    <text class="text-subtitle-2">选择目标语言</text>
    <a-radio-group v-model:value="value" name="radioGroup" class="mb-3">
      <a-radio value="zh">中文</a-radio>
      <a-radio value="en">英语</a-radio>
      <a-radio value="yue">粤语</a-radio>

    </a-radio-group>
    <v-btn color="primary" text="翻译" class="align-center text-center mb-5 mx-auto" block></v-btn>
  </v-container>
</template>

<style scoped>

</style>
