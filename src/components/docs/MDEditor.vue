<template>
  <div class="fill-height">
    <v-btn href="/workspace/docs" icon="mdi-chevron-left" class="mb-3 ml-3" size="30" color="grey"></v-btn>
    <md-editor class="" v-model="text" :toolbars-exclude="toolbars" @onSave="onSave" @onChange="onChange"
               @onUploadImg="onUploadImg" no-upload-img/>
  </div>
</template>

<script>
import {defineComponent} from "vue";
import {MdEditor} from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import {message} from 'ant-design-vue';

export default defineComponent({
  name: "App",
  components: {
    MdEditor,
  },
  data() {
    return {
      text: "# 这是标题",
      toolbars: ["github", "htmlPreview"],
      id: ""
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
      message.success({
        content: () => '保存成功',
        style: {
          marginTop: '0vh',
        },
      });
    }
  },
  mounted() {
    this.id = this.$route.params.id
  }
});
</script>
