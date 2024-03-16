<template>
  <v-app>
    <v-main>
      <v-container class="mt-6">
        <v-responsive class="align-center text-center fill-height">
          <v-icon icon="mdi-cube" size="50" color="blue"></v-icon>

          <h2 class="font-weight-bold mt-3">登录</h2>
          <v-sheet class="mx-auto" max-width="350">
            <a-form layout="vertical" :model="userinfo">
              <a-form-item label="用户名">
                <a-input v-model:value="userinfo.username"></a-input>
              </a-form-item>
              <a-form-item label="密码">
                <a-input type="password" v-model:value="userinfo.password"></a-input>
              </a-form-item>
              <a-button size="large" html-type="submit" :disabled="infoSubmitDisable" @click="infoSubmit">
                登录
              </a-button>
            </a-form>
          </v-sheet>

        </v-responsive>
      </v-container>


    </v-main>
  </v-app>
</template>
<script>

export default {
  data() {
    return {
      userinfo: {
        username: '',
        password: '',
      }
    }
  }
}
</script>
