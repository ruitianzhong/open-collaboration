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
              :max-count="1"
              :action="uploadPath()"
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
            <v-snackbar
              timeout="2000"
              color="#EDEDED"
              variant="flat"
              v-model="snackbar"
              class="align-center"
            >
              <v-icon icon="mdi-check"></v-icon>
              成功上传文件 <strong>{{ text }}</strong>
            </v-snackbar>
          </v-container>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              text="完成" color="#07c360"
              @click="onFinished(isActive)"
            ></v-btn>
          </v-card-actions>
        </v-card>
      </template>
    </v-dialog>


    <a-table :columns="columns" :data-source="files">
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
          <a :href=" '/files/download?' + qs.stringify({
    filename: record.filename,group: AppState.group_id,})">下载 </a>
          <a-divider type="vertical"/>
          <a @click="onDelete(record.filename)">删除</a>
        </span>
        </template>
      </template>
    </a-table>
  </v-container>
</template>
<script setup>
import {onMounted, ref} from 'vue';
import {AppState} from "@/main";
import {deleteFiles, listFiles} from "@/api/api";
import qs from "qs";
import dayjs from "dayjs";

const fileList = ref([]);

let files = ref([])
let snackbar = ref(false)
const handleChange = info => {
  const status = info.file.status;
  if (status !== 'uploading') {
    console.log(info.file, info.fileList);
  }
  if (status === 'done') {
    text.value = info.file.name
    snackbar.value = true
  } else if (status === 'error') {

  }
};
const text = ref('')

const onFinished = isActive => {
  refresh()
  isActive.value = false
  fileList.value = []
}

const onDelete = filename => {
  const req = {
    group: AppState.group_id,
    filename: filename
  }
  deleteFiles(req)
  refresh()

}

const uploadPath = () => {
  const req = {
    groupId: AppState.group_id,
  }

  return "/files/upload?" + qs.stringify(req)
}

onMounted(() => {
    const req = {
      group: AppState.group_id
    }
    listFiles(qs.stringify(req)).then(
      response => {
        const {data} = response
        if (data.ok) {

          for (let dataKey in data.files) {
            data.files[dataKey].lastModified = dayjs(data.files[dataKey].lastModified).format("YYYY-MM-DD HH:mm:ss")
          }
          files.value = data.files
        }
      }
    )
  }
)

function refresh() {
  const req = {
    group: AppState.group_id
  }
  listFiles(qs.stringify(req)).then(
    response => {
      const {data} = response
      if (data.ok) {

        for (let dataKey in data.files) {
          data.files[dataKey].lastModified = dayjs(data.files[dataKey].lastModified).format("YYYY-MM-DD HH:mm:ss")
        }
        files.value = data.files
      }
    }
  )
}

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
    dataIndex: 'lastModified',
    key: 'uploadedTime',
  },
  {
    title: '上传者',
    dataIndex: 'uploader',
    key: 'uploader',
  },
  {
    title: "大小（字节）",
    key: "size",
    dataIndex: 'size',
  },
  {
    key: 'action',
  },
];

</script>
