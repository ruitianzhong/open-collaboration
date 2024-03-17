<script>
import Translation from "@/components/Translation.vue";
import Formula from "@/components/Formula.vue";
import Chat from "@/components/Chat.vue";

export default {
  components: {Chat, Formula, Translation},
  data() {
    return {}
  }
}
</script>
<template>
  <v-app-bar flat>
    <v-btn icon="mdi-cube" size="35" class="ml-3 mr-2" href="/workspace"></v-btn>
    <v-btn text="Latex" density="default"></v-btn>
    <v-btn text="文件共享" density="default"></v-btn>

    <template v-slot:append>
      <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
          <v-btn
            v-bind="activatorProps"
            icon="mdi-translate"
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
      <v-btn icon="mdi-account" href="/workspace/userinfo"></v-btn>
    </template>
    <!--    </v-container>-->
  </v-app-bar>
</template>
