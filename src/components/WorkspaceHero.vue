<template xmlns="http://www.w3.org/1999/html">
  <v-container style="margin-top: 8%">
    <v-responsive class="align-center text-center fill-height">
      <div class="text-h2 font-weight-bold">欢迎使用
        <!--        <text style="color: cornflowerblue">协作平台</text>-->
        <v-typical
          class="blink text-h2 font-weight-bold"
          :steps="['',500,'协', 500, '协作', 500, '协作平', 1000,'协作平台',1000]"
          :loop="Infinity"
          :wrapper="'text'"
          style="color: cornflowerblue"
        ></v-typical>
      </div>
      <div class="text-h5 font-weight-light  mt-8 typing">更高效地和你的队友/同事沟通协作</div>
      <div class="py-8"/>
      <v-row>
        <v-col>
          <v-container border="false" elevation="0" class="mx-auto" max-width="344">
            <v-card-title class="text-h5">
              文档 / 文件共享
            </v-card-title>
            <v-card-subtitle>
              支持Markdown文档在线编辑和预览
            </v-card-subtitle>
            <v-card-actions>
              <v-spacer></v-spacer>
            </v-card-actions>
            <v-expand-transition>

            </v-expand-transition>
          </v-container>
        </v-col>
        <v-col>
          <v-container class="mx-auto" max-width="344">
            <v-card-title class="text-h5">
              即时沟通
            </v-card-title>

            <v-card-subtitle>
              即时通讯助力高效工作
            </v-card-subtitle>


          </v-container>
        </v-col>
        <v-col>

          <v-container class="mx-auto" max-width="344">

            <v-card-title class="text-h5">
              丰富工具
            </v-card-title>

            <v-card-subtitle>
              提供翻译、公式渲染等趁手工具
            </v-card-subtitle>

            <v-card-actions>
              <v-spacer></v-spacer>
            </v-card-actions>

          </v-container>
        </v-col>

      </v-row>
      <v-row class="d-flex align-center justify-center">
        <!--        <v-col cols="auto" class="mr-10">-->
        <!--          <v-btn-->
        <!--            href="/signup"-->
        <!--            min-width="200"-->
        <!--            variant="flat"-->
        <!--            color="#EDEDED"-->
        <!--            size="x-large">-->
        <!--            <v-icon-->
        <!--              icon="mdi-plus"-->
        <!--              size="large"-->
        <!--              start-->
        <!--            />-->
        <!--            注册账号-->
        <!--          </v-btn>-->
        <!--        </v-col>-->

        <!--        <v-col cols="auto">-->
        <!--          <v-btn-->
        <!--            color="#07c360"-->
        <!--            href="/login"-->
        <!--            min-width="200"-->
        <!--            size="x-large"-->
        <!--            variant="flat"-->
        <!--          >-->
        <!--            <v-icon-->
        <!--              icon="mdi-login"-->
        <!--              size="large"-->
        <!--              start-->
        <!--            />-->
        <!--            登录账号-->
        <!--          </v-btn>-->
        <!--        </v-col>-->
      </v-row>
      <v-row class="d-flex align-center justify-center">

      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup>
import {VTypical} from "vue-typical";
//
</script>
<style>
.blink::after {
  content: '|';
  animation: blink 1s infinite step-start;
}
</style>
