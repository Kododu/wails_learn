<template>
  <main>
    <div style="width: 100%">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" :name="item.identify">
          <template #title>
            <div class="item">
              <span>{{ item.name }}</span>
            </div>
            <div style="display: flex">
              <span><connection-manager :item="item" @click.stop edit-title="编辑" edit-type="text"/></span>
              <el-popconfirm @confirm="deleteItem(item.identify)" width="220" confirm-button-text="是"
                             cancel-button-text="否" title="确认删除吗？">
                <template #reference>
                  <el-button @click.stop link type="danger">删除</el-button>
                </template>
              </el-popconfirm>
            </div>
          </template>
        </el-collapse-item>
      </el-collapse>
    </div>
  </main>
</template>

<script setup>
import {ref, watch} from "vue";
import {ConnectionDelete, ConnectList} from "../../wailsjs/go/api/App.js";
import {ElNotification} from "element-plus"
import ConnectionManager from "./ConnectionManager.vue";

let props = defineProps(['flush'])
watch(props, (newProps) => {
  console.log(newProps.flush)
  refreshList()
})

let list = ref([])

function refreshList() {
  ConnectList().then(result => {
    if (result.code !== 1000) {
      ElNotification({title: "错误", message: result.message, type: "error"})
    } else {
      list.value = result.data.list
    }
  })
}

function deleteItem(identify) {
  ConnectionDelete(identify).then(res => {
    if (res.code !== 1000) {
      ElNotification({title: "错误", message: res.message, type: "error"})
    } else {
      ElNotification({
        title: "成功",
        type: "success",
        duration: 1200,
      })
      refreshList()
    }
  })
}

refreshList()
</script>

<style scoped>
.item {
  margin-left: 18px;
  width: 100%;
  display: flex;
}
</style>