<script>
import {translateBaidu} from "@/api/api";

export default {
  data() {
    return {
      source: "你好，世界",
      target: "",
      value: 'en',
      loading: false
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
    },
  },
  methods: {
    translateClick() {
      if (this.source == "") {
        return
      }
      let form = {
        target: this.value,
        source: this.source
      }
      this.loading = true
      translateBaidu(form).then(response => {
        const {data} = response
        this.target = data.dst
        this.loading = false
      }).catch(error => {
        this.loading = false
      })

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
    <text class="text-subtitle-2 mr-2">选择目标语言</text>
    <a-radio-group v-model:value="value" name="radioGroup" class="mb-3">
      <a-radio value="zh">中文</a-radio>
      <a-radio value="en">英语</a-radio>
      <a-radio value="yue">粤语</a-radio>

    </a-radio-group>
    <v-btn color="primary" text="翻译" class="align-center text-center mb-5 mx-auto" block
           :disabled="source==''" @click="translateClick" :loading="loading"></v-btn>
  </v-container>
</template>

<style scoped>

</style>
