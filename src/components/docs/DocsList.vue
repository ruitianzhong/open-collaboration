<script>
import router from "@/router";
import {deleteDocs, listDocs} from "@/api/api";
import {AppState} from "@/main";
import qs from "qs";
import dayjs from "dayjs";

export default {
  name: "DocsList",
  mounted() {
    this.refresh()
  },
  data() {
    return {
      data: [],
      list: []
    }
  },
  methods: {

    refresh() {
      const req = {
        group: AppState.group_id,
      }
      listDocs(qs.stringify(req)).then(res => {
        const {data} = res
        this.list = data.docs
      })

    },

    confirm(id) {

      const request = {id: id, group: AppState.group_id}
      deleteDocs(request).then(
        resp => {
          const {data} = resp
          if (data.ok) {
            this.refresh()
          }
        }
      )
    },
    onClick(id) {
      this.$router.push({
        path: `/workspace/docs/view/${id}`
      })
    },
    onEdit(id) {
      this.$router.push({
        path: `/workspace/docs/edit/${id}`
      })
    },
    add() {
      router.push({path: "/workspace/docs/editor"})
    },
    metaInfo(created, lastUpdated, author) {

      let c = dayjs(created).format("YYYY-MM-DD HH:mm:ss")

      let l = dayjs(lastUpdated).format("YYYY-MM-DD HH:mm:ss")

      return "Created on: " + c + " Updated on " + l + " Author: " + author

    }
  },
  computed: {}
}
</script>

<template>


  <v-container>
    <v-sheet>
      <v-btn-group class="text-lg-right mr-8 mt-4" style="display: flex;justify-content: flex-end">
        <v-btn style="" prepend-icon="mdi-plus" width="80" color="#07c360" height="35"
               slim @click="add"
               variant="flat" density="comfortable" text="添加"></v-btn>
      </v-btn-group>

      <a-list :data-source="list" size="large">
        <template #renderItem="{ item }">
          <a-list-item>
            <template #actions>
              <a key="list-edit" @click="onEdit(item.id)">编辑</a>
              <a-popconfirm title="是否要删除该篇文档？" ok-text="确认" cancel-text="取消"
                            :ok-button-props="{ghost:true,danger:true}" @confirm="confirm(item.id)">

                <a key="list-delete" href="#">删除</a>
              </a-popconfirm>

            </template>
            <a-list-item-meta
              :description="metaInfo(item.createdTime,item.lastModified,item.creator)">
              <template #title>
                <a @click="onClick(item.id)" class="text-h5">{{ item.title }}</a><br/>
              </template>
            </a-list-item-meta>
          </a-list-item>
        </template>
      </a-list>
    </v-sheet>
  </v-container>
</template>

<style scoped lang="scss">

</style>
