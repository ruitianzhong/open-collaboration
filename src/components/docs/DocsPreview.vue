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
import router from "@/router";
import {getDocs} from "@/api/api";
import qs from "qs";
import {AppState} from "@/main";

export default defineComponent({
    name: "App",
    components: {
      MdPreview
    },
    data() {
      return {
        id: "",
        text: "",
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
    },
    mounted() {
      this.id = this.$route.params.id
      if (this.id == undefined) {
        router.replace({path: '/workspace/docs'})
      } else {
        const query = {
          id: this.id,
          group: AppState.group_id,
        }
        getDocs(qs.stringify(query)).then(req => {
          const {data} = req
          if (data.ok) {
            this.text = data.markdown
          }
        })

      }
    },
    unmounted() {
      this.text = ""
    }
  }
);
</script>
