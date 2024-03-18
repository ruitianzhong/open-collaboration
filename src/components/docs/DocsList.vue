<script xmlns="http://www.w3.org/1999/html">
export default {
  name: "DocsList",
  data() {
    return {
      list: [
        {
          title: "Attention is all you need",
          author: "Tim",
          id: "12345"
        },
        {
          title: "Xen and the art of virtualization",
          author: "Jack",
          id: "12345678"
        },
        {
          title: "How to read a book?",
          author: "Taylor",
          id: "222"
        }
      ]
    }
  },
  methods: {
    onClick(id) {
      this.$router.push({
        path: `/workspace/docs/view/${id}`
      })
    },
    onEdit(id) {
      this.$router.push({
        path: `/workspace/docs/edit/${id}`
      })
    }
  }
}
</script>

<template>


  <v-container>

    <v-sheet>
      <v-btn-group class="text-lg-right mr-8 mt-4" style="display: flex;justify-content: flex-end">
        <v-btn style="" prepend-icon="mdi-plus" href="/workspace/docs/editor" width="80" color="#07c360" height="35"
               slim
               variant="flat" density="comfortable" text="添加"></v-btn>
      </v-btn-group>

      <a-list :data-source="list" size="large">
        <template #renderItem="{ item }">
          <a-list-item>
            <template #actions>
              <a key="list-edit" @click="onEdit(item.id)">编辑</a>
              <a-popconfirm title="是否要删除该篇文档？" ok-text="确认" cancel-text="取消"
                            :ok-button-props="{ghost:true,danger:true}">

                <a key="list-delete" href="#">删除</a>
              </a-popconfirm>

            </template>
            <a-list-item-meta
              description="Last updated on 2024/03/17 Author: Taylor">
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
