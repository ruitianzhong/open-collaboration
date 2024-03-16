<template>
  <v-container class="mt-6">
    <v-responsive class="align-center  fill-height">

      <v-sheet max-width="350" class="mx-auto">
        <v-list-item class="" width="500" :prepend-avatar="avatar" :title="name"
                     :subtitle="'群组编号'+groupId"
                     lines="two"
                     density="comfortable">
          <template v-slot:title><span style="font-weight: 600;">{{ nickname }} <span class="text-subtitle"
                                                                                      style="font-weight: 50;">({{
              userid
            }})</span></span></template>
          <template v-slot:prepend>
            <v-avatar size="50">
              <v-img :src="avatar" alt="John"></v-img>
            </v-avatar>
          </template>
        </v-list-item>
        <h2 class="mt-5" style="font-weight: 500;">修改密码</h2>
        <v-divider class="ml-3 mt-5 mr-6"></v-divider>

        <p class="mb-2 mt-5">原密码</p>
        <v-text-field v-model="oldPassword" type="password" variant="outlined" color="blue" bg-color="white"
                      density="compact"></v-text-field>
        <p class="mb-2 ">原密码</p>
        <v-text-field v-model="newPassword" type="password" :disabled="oldPassword == ''" variant="outlined"
                      color="blue"
                      bg-color="white" density="compact"></v-text-field>
        <p class="mb-2">确认密码</p>
        <v-text-field v-model="confirmPassword" type="password" :disabled="newPassword == ''" variant="outlined"
                      color="blue"
                      density="compact"></v-text-field>
        <div class="mx-auto  text-center">
          <v-btn class="" :loading="passwordLoading" width="100" color="green" height="35" slim
                 variant="outlined" density="comfortable" text="更新密码"
                 :disabled="this.newPassword == '' || this.oldPassword == '' || this.confirmPassword == ''"
                 @click="submitNewPassword()">

          </v-btn>
          <v-alert class="mt-3 mb-3" width="400" density="compact" v-model="alert" @click:close="alert = false" closable
                   variant="outlined" :type="type"
                   :text="msg"></v-alert>
        </div>
      </v-sheet>


    </v-responsive>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      nickname: "Pony",
      userid: "12345",
      oldPassword: "",
      newPassword: "",
      confirmPassword: "",
      msg: "",
      avatar: "https://api.multiavatar.com/Starcrasher.png",
      alert: false,
      groupId: "222"


    }
  }
}
</script>
