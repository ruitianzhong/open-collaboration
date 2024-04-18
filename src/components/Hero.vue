<template>
  <v-container style="margin-top: 10%">
    <v-responsive class="align-center text-center fill-height">
      <div class="text-h2 font-weight-bold">欢迎使用
        <text style="color: cornflowerblue">协作平台</text>
      </div>
      <div class="text-h5 font-weight-light  mt-8 typing">更高效地和你的队友/同事沟通协作</div>
      <div class="py-8"/>
      <v-row class="d-flex align-center justify-center">
        <v-col cols="auto" class="mr-10">
          <v-btn
            @click="$router.push({path:'/signup'})"
            min-width="200"
            variant="flat"
            color="#EDEDED"
            size="x-large">
            <v-icon
              icon="mdi-plus"
              size="large"
              start
            />
            注册账号
          </v-btn>
        </v-col>

        <v-col cols="auto">
          <v-btn
            color="#07c360"
            @click="$router.push({path:'/login'})"
            min-width="200"
            size="x-large"
            variant="flat"
          >
            <v-icon
              icon="mdi-login"
              size="large"
              start
            />
            登录账号
          </v-btn>
        </v-col>
      </v-row>
      <v-row class="d-flex align-center justify-center">

      </v-row>
    </v-responsive>
  </v-container>
</template>

<script setup>
//
</script>
