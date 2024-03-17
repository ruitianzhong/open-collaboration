<template>
  <md-editor class="fill-height" v-model="text" :toolbars-exclude="toolbars" @onSave="onSave" @onChange="onChange"
             @onUploadImg="onUploadImg" no-upload-img/>
</template>

<script>
import {defineComponent} from "vue";
import {MdEditor} from "md-editor-v3";
import "md-editor-v3/lib/style.css";

export default defineComponent({
  name: "App",
  components: {
    MdEditor,
  },
  data() {
    return {
      text: "# 这是标题",
      toolbars: ["github", "htmlPreview"],
    };
  },
  methods: {
    onChange(val) {
    },
    onUploadImg(files) {
      console.log(Array.from(files));
    },
    onSave(value, html) {
      console.log(value)
    }
  },
});
</script>
