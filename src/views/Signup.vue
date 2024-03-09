<template>
  <v-app>
    <v-main>
      <div v-if="0==step">
        <v-container class="mt-6">
          <v-responsive class="align-center text-center fill-height">
            <v-icon icon="mdi-cube" size="50" color="blue"></v-icon>

            <h2 class="font-weight-bold mt-3">请填写用户信息</h2>
            <v-sheet class="mx-auto" max-width="350">
              <a-form layout="vertical" :model="userinfo">
                <a-form-item label="用户名">
                  <a-input v-model:value="userinfo.username"></a-input>
                </a-form-item>
                <a-form-item label="密码">
                  <a-input type="password" v-model:value="userinfo.password"></a-input>
                </a-form-item>
                <a-form-item label="确认密码">
                  <a-input type="password" v-model:value="userinfo.confirm"></a-input>
                </a-form-item>
                <a-button size="large" html-type="submit" :disabled="infoSubmitDisable" @click="infoSubmit">
                  下一步
                </a-button>
              </a-form>
            </v-sheet>

          </v-responsive>
        </v-container>
      </div>


      <div v-if="step==1">
        <v-container class="mt-6">
          <v-responsive class="align-center text-center fill-height">
            <v-icon icon="mdi-cube" size="50" color="blue"></v-icon>
            <h2 class="font-weight-bold mt-5">请完成以下验证</h2>
          </v-responsive>
        </v-container>

      </div>

      <div v-if="step==2">
        <v-container class="mt-6">
          <v-responsive class="align-center text-center fill-height">
            <v-icon icon="mdi-cube" size="50" color="blue"></v-icon>
            <h2 class="font-weight-bold mt-5">请输入收到的验证码</h2>
            <v-sheet
              class="pt-8 pb-12 px-6 ma-4 mx-auto"
              max-width="350"
              width="100%"
              border
            >
              <h3 class="text-h6 mb-1"></h3>

              <div class="text-body-2 font-weight-light">
                我们已经将验证码发送给 <span class="font-weight-black text-primary">+1 408 555 1212</span>
              </div>

              <v-otp-input
                v-model="otp"
                class="mt-3 ms-n2"
                length="4"
                placeholder="0"
                variant="underlined"
              ></v-otp-input>
              <v-btn
                color="primary"
                size="large"
                text="下一步"
                variant="tonal"
                class="mt-6"
                @click="otpClick"
              ></v-btn>
            </v-sheet>

          </v-responsive>
        </v-container>
      </div>

      <div v-if="step==3">
        <v-container class="mt-6">
          <v-responsive class="align-center text-center fill-height">

            <a-result
              status="success"
              title="成功注册"
              :sub-title='"用户名："+userinfo.username+" 手机号："+userinfo.telephone'
            >
              <template #extra>
                <a-button key="console" type="success">Go Console</a-button>
                <a-button key="buy">Buy Again</a-button>
              </template>
            </a-result>
          </v-responsive>
        </v-container>

      </div>

    </v-main>
  </v-app>
</template>
<script>
export default {
  data() {
    return {
      step: 2,
      userinfo: {
        password: "",
        username: "",
        confirm: "",
        telephone: ""
      },
      otp: '',

    }
  },
  computed: {
    infoSubmitDisable() {
      return !(this.userinfo.confirm != '' && this.userinfo.password != '' && this.userinfo.username != '')
    }

  },
  methods: {
    infoSubmit() {
      console.log("submit")
      this.step = 1
    },
    otpClick() {
      this.step = 3
    }
  }

}
</script>
