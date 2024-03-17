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
      <a-list :data-source="list" size="large">
        <template #renderItem="{ item }">
          <a-list-item>
            <template #actions>
              <a key="list-loadmore-edit" @click="onEdit(item.id)">编辑</a>
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
