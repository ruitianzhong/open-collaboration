<template>
  <v-container>
    <v-sheet>
      <v-btn @click="onBack" icon="mdi-chevron-left" class="mb-1 ml-3" size="30" color="grey"></v-btn>
      <MdPreview :model-value="text">
      </MdPreview>
    </v-sheet>

  </v-container>

</template>

<script>
import {defineComponent} from "vue";
import {MdPreview} from "md-editor-v3";
import "md-editor-v3/lib/style.css";


export default defineComponent({
    name: "App",
    components: {
      MdPreview
    },
    data() {
      return {
        id: "",
        text: "# 这是标题 \r\n # heading 2 \r\n## heading 3\r\n```bash \r\n echo hello world\r\n```",
        toolbars: ["github", "htmlPreview"],
      };
    },
    methods: {
      onBack() {
        if (window.history.length <= 1) {
          this.$router.push({path: '/workspace/docs'})
          return false
        } else {
          this.$router.go(-1)
        }

      },
      mounted() {
        this.id = this.$route.params.id
      }

    }
  }
);
</script>
