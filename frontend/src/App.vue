<template>
  <el-card class="card-box">
    <template #header>
      <div class="card-header">
        <span>链接管理</span>
        <connection-manager
            :item="undefined" :edit-icon="Edit" edit-title="新建" edit-type="primary"
            @emit-connection-change="flushConnection"
        />
      </div>
    </template>
    <div class="card-content">
      <el-scrollbar>
        <conn-list :flush="flush.flag"/>
      </el-scrollbar>
    </div>
  </el-card>

</template>

<script setup>
import ConnList from "./components/ConnectionList.vue";
import ConnectionManager from "./components/ConnectionManager.vue";
import {reactive} from "vue";
import {Edit} from "@element-plus/icons";


let flush = reactive({
  flag: true
})

function flushConnection() {
  flush.flag = !flush.flag
}

</script>

<style scoped>

.card-box {
  width: 240px;
  height: 95vh;
  margin-left: 10px;
  margin-top: 2vh;
  margin-bottom: 2vh;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-content {
  width: 215px;
  height: 80vh;
}
</style>
