<script>
import Translation from "@/components/Translation.vue";
import Formula from "@/components/Formula.vue";

export default {
  components: {Formula, Translation},
  data() {
    return {}
  }
}
</script>
<template>
  <v-app-bar flat>
    <v-icon icon="mdi-cube" size="35" class="ml-3 mr-2"></v-icon>
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


      <v-dialog max-width="500">
        <template v-slot:activator="{ props: activatorProps }">
          <v-btn
            v-bind="activatorProps"
            icon="mdi-chat"
          ></v-btn>
        </template>

        <template v-slot:default="{ isActive }">
          <v-card title="Dialog">
            <v-card-text>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
              dolore magna aliqua.
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn
                text="Close Dialog"
                @click="isActive.value = false"
              ></v-btn>
            </v-card-actions>
          </v-card>
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


      <v-btn icon="mdi-account"></v-btn>
    </template>
    <!--    </v-container>-->
  </v-app-bar>
</template>
