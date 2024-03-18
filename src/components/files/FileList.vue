<template>
  <v-container>
    <v-dialog max-width="700">
      <template v-slot:activator="{ props: activatorProps }">
        <v-btn class="mb-5 mt-5" style="display: flex;justify-content: flex-end;"
               v-bind="activatorProps"
               prepend-icon="mdi-upload"
               text="上传"
               variant="flat"
               color="#07c360"
        ></v-btn>
      </template>


      <template v-slot:default="{ isActive }">
        <v-card title="上传文件" prepend-icon="mdi-upload">
          <v-container>
            <a-upload-dragger
              v-model:fileList="fileList"
              name="file"
              :multiple="false"
              action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              @change="handleChange"
              @drop="handleDrop"
            >
              <p class="ant-upload-drag-icon">
                <inbox-outlined></inbox-outlined>
              </p>
              <p class="ant-upload-text">点击或拖动文件到该区域</p>
              <p class="ant-upload-hint">
                请勿上传过大文件
              </p>
            </a-upload-dragger>
          </v-container>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              text="取消"
              @click="isActive.value = false"
            ></v-btn>
            <v-btn
              text="确认" color="#07c360"
              @click="isActive.value = false"
            ></v-btn>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>


    <a-table :columns="columns" :data-source="data">
      <template #headerCell="{ column }">
        <template v-if="column.key === 'filename'">
        <span>
          文件名
        </span>
        </template>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <a>
            {{ record.name }}
          </a>
        </template>

        <template v-else-if="column.key === 'action'">
        <span>
          <a>下载 </a>
          <a-divider type="vertical"/>
          <a>删除</a>
        </span>
        </template>
      </template>
    </a-table>


  </v-container>
</template>
<script setup>
import {ref} from 'vue';
import {message} from 'ant-design-vue';
import Translation from "@/components/Translation.vue";

const fileList = ref([]);
const handleChange = info => {
  const status = info.file.status;
  if (status !== 'uploading') {
    console.log(info.file, info.fileList);
  }
  if (status === 'done') {
    message.success(`${info.file.name} file uploaded successfully.`);
  } else if (status === 'error') {
    message.error(`${info.file.name} file upload failed.`);
  }
};

function handleDrop(e) {
  console.log(e);
}


const columns = [
  {
    name: 'Name',
    dataIndex: 'filename',
    key: 'filename',
  },
  {
    title: '上传时间',
    dataIndex: 'uploadedTime',
    key: 'uploadedTime',
  },
  {
    title: '上传者',
    dataIndex: 'uploader',
    key: 'uploader',
  },
  {

    key: 'action',
  },
];
const data = [
  {
    key: '1',
    filename: "hello.c",
    uploader: "Tim Zhang",
    uploadedTime: "2024-03-18",
    id: "xyz",
  },
  {
    key: '2',
    filename: "hello.go",
    uploader: "Allen",
    uploadedTime: "2024-03-18",
    id: "xyz123",
  },
  {
    key: '3',
    filename: "print.xlsx",
    uploader: "Tim Zhang",
    uploadedTime: "2024-03-18",
    id: "abc",
  },
];

</script>
