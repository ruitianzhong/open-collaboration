<template>
  <div :class="['TUIKit', isH5 && 'TUIKit-h5']">
    <div class="TUIKit-main-container">

      <TUIChat v-show="!isH5 || currentConversationID">
        <v-container class="align-center">
          <v-spacer></v-spacer>
          <a-result status="404" class="ml-8 mt-12" title="服务暂时不可用" sub-title="请刷新重试或联系网站管理员">

          </a-result>
        </v-container>
      </TUIChat>
      <TUICallKit
        class="callkit-container"
        :allowedMinimized="true"
        :allowedFullScreen="true"
      />
    </div>
  </div>

</template>
<script setup lang="ts">
import {ref} from "vue";
import {TUIStore, StoreName, TUIConversationService} from "@tencentcloud/chat-uikit-engine";
import {TUICallKit} from "@tencentcloud/call-uikit-vue";
import {genTestUserSig, TUIChat} from "@/TUIKit";
import {isH5} from "@/TUIKit/utils/env";

const userID = "root";

const currentConversationID = ref<string>("");

const groupID = "group1";
TUIStore.watch(StoreName.CONV, {
  currentConversationID: (id: string) => {
    currentConversationID.value = id;
  },
});
TUIConversationService.switchConversation(`GROUP${groupID}`);

function openChat() {
  // 切换会话进入聊天
  TUIConversationService.switchConversation(`C2C${userID}`);
}

function openGroupChat() {
  // 切换会话进入聊天
  TUIConversationService.switchConversation(`GROUP${groupID}`);
}
</script>
<style scoped lang="scss">
@import "@/TUIKit/assets/styles/common.scss";
@import "@/TUIKit/assets/styles/sample.scss";

.chat {
  flex: 0.8;
  background: white;
}
</style>
