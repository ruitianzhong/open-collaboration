<script>
import Translation from "@/components/Translation.vue";
import Formula from "@/components/Formula.vue";
import Chat from "@/components/Chat.vue";
import router from "@/router";
import {fetchUserInfo, logout} from "@/api/api";
import {AppState} from "@/main";

export default {
  components: {Chat, Formula, Translation},
  data() {
    return {}
  },
  mounted() {
    fetchUserInfo().then(response => {
      const {data} = response
      console.log(data)
      AppState.group_id = data.groupId
      AppState.user_id = data.userId
    }).catch(error => {
      router.replace({path: "/"})
    })

  },
  methods: {
    onClick(path) {
      router.push({
        path,
        replace: false
      })
    },
    onLogout() {
      logout({}).then(
        response => {
          const {status} = response
          console.log(response)
          if (status == 200) {

            router.replace({path: "/"})
          }
        }
      )

    }
  }
}
</script>
<template>
  <v-app-bar flat elevation="1" density="comfortable">
    <v-btn icon="mdi-cube" variant="plain" size="35" class="ml-3 mr-2" @click="onClick('/workspace')"
           color="black"></v-btn>
    <v-btn variant="plain" text="文档" density="comfortable" @click="onClick('/workspace/docs')" color="black"></v-btn>
    <v-btn text="文件共享" variant="plain" density="default" color="black" @click="onClick('/workspace/files')"></v-btn>
    <template v-slot:append>
      <v-divider vertical class="mr-2"></v-divider>
      <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
          <v-btn
            v-bind="activatorProps"
            icon="mdi-translate"
            variant="plain"
            color="black"
          ></v-btn>
        </template>

        <template v-slot:default="{ isActive }">
          <v-card title="中英翻译" prepend-icon="mdi-translate">
            <Translation></Translation>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn

                text="关闭"
                @click="isActive.value = false"
              ></v-btn>
            </v-card-actions>
          </v-card>
        </template>
      </v-dialog>


      <v-dialog max-width="1000">
        <template v-slot:activator="{ props: activatorProps }">
          <v-btn
            v-bind="activatorProps"
            icon="mdi-chat"
            variant="plain"
            color="black"
          ></v-btn>
        </template>

        <template v-slot:default="{ isActive }">

          <Chat></Chat>

          <v-spacer></v-spacer>


        </template>
      </v-dialog>
      <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
          <v-btn
            v-bind="activatorProps"
            icon="mdi-math-integral-box"
            variant="plain"
            color="black"
          ></v-btn>
        </template>

        <template v-slot:default="{ isActive }">
          <v-card title="Latex公式实时预览" prepend-icon="mdi-math-integral-box">
            <Formula></Formula>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn
                text="关闭"
                @click="isActive.value = false"
              ></v-btn>
            </v-card-actions>
          </v-card>
        </template>
      </v-dialog>
      <v-btn icon="mdi-account" variant="plain" @click="onClick('/workspace/userinfo')" color="black"></v-btn>
      <v-btn icon="mdi-exit-to-app" alt="exit" variant="plain" @click="onLogout" color="primary"></v-btn>
    </template>
    <!--    </v-container>-->
  </v-app-bar>
</template>
