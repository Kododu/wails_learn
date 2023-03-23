<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <el-input id="name" style="width: 360px" v-model="data.name" placeholder="选择配置文件" clearable @clear="clear"/>
      <el-button type="primary" class="btn" @click="connect">连接</el-button>
      <el-button type="success" class="btn" @click="create" :icon="Edit">新建</el-button>
    </div>
    <el-dialog v-model="data.dialogVisible" title="提示" :before-close="handleClose">
      <span>{{ data.msg }}</span>
      <template #footer>
        <span class="el-dialog__footer">
          <el-button type="success" @click="handleClose">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import {reactive} from 'vue'
import {ConnectList,ConnectionCreate} from '../../wailsjs/go/api/App'
import {Edit} from "@element-plus/icons";

const data = reactive({
  name: "",
  resultText: "输入点什么 👇",
  msg: "",
  dialogVisible: false,
})

function connect() {
  ConnectList().then(res => {
    if (res.code === 1000) {
      data.resultText = res.data
      data.msg = ""
    } else {
      data.msg = res.message
      data.dialogVisible = true
    }
  })
}

function create() {
  ConnectionCreate({"address":data.name}).then(res => {
    if (res.code === 1000) {
      data.msg = ""
      connect()
    } else {
      data.msg = res.message
      data.dialogVisible = true
    }
  })
}

function handleClose() {
  data.resultText = data.msg
  data.dialogVisible = false
}

</script>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .upload {
  width: 80px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
