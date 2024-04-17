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
              <v-alert v-model="alert" closable
                       text="密码错误" density="compact" variant="tonal"
                       type="error" class="mb-4 align-center"
              ></v-alert>
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

import {login} from "@/api/api";
import router from "@/router";

export default {
  data() {
    return {
      userinfo: {
        username: '',
        password: '',
      },
      disabled: true,
      alert: false
    }
  },
  computed: {
    infoSubmitDisable() {
      return this.userinfo.password == '' || this.userinfo.username == ''
    }
  },
  methods: {
    async infoSubmit() {
      let form = {
        userid: this.userinfo.username,
        password: this.userinfo.password,
      }
      login(form).then(response => {
        const {data} = response
        if (data.code == "200") {
          router.push({path: "/workspace"})
        } else {
          this.alert = true
        }

      }).catch(
        error => {
          console.log(error)
        }
      )


    }
  }

}
</script>
