<template>
  <main>
    <el-button id="edit-btn" :icon="editIcon" :type="editType" @click="data.dialogVisible = true">{{editTitle}}</el-button>
    <el-dialog v-model="data.dialogVisible" :title="editTitle">
      <div>
        <el-form :label-position="'right'" label-width="100px" :model="formData" style="max-width: 460px">
          <el-input type="hidden" v-model="formData.identify"></el-input>
          <el-form-item label="别名">
            <el-input :placeholder="formData.address" v-model="formData.name"/>
          </el-form-item>
          <el-form-item label="IP">
            <el-input v-model="formData.address"/>
          </el-form-item>
          <el-form-item label="端口号">
            <el-input placeholder="6404" v-model="formData.port"/>
          </el-form-item>
          <el-form-item label="用户名">
            <el-input v-model="formData.user"/>
          </el-form-item>
          <el-form-item label="密码">
            <el-input type="password" v-model="formData.password" show-password/>
          </el-form-item>
        </el-form>
      </div>
      <template #footer>
        <span class="el-dialog__footer">
          <el-button type="primary" @click="handleClose">确认</el-button>
          <el-button  @click="data.dialogVisible = false">取消</el-button>
        </span>
      </template>
    </el-dialog>
  </main>
</template>

<script setup>
import {reactive} from "vue";
import {ConnectionCreate} from "../../wailsjs/go/api/App.js";
import {ElNotification} from "element-plus";

const data = reactive({
  dialogVisible: false,
  // identify:'',
  // name: '',
  // address: '',
  // port: '',
  // user:'',
  // password: '',
})

let formData = reactive({})

const props = defineProps(['editTitle', 'editType','editIcon','item','isLink'])
if (props.item !== undefined){
  formData = props.item
}

const emits = defineEmits(['emit-connection-change'])
function handleClose() {
  ConnectionCreate(formData).then(res => {
    if (res.code !== 1000) {
      ElNotification({
        title:"错误",
        type:"error",
        duration:60000,
        message:res.message
      })
    } else {
      emits('emit-connection-change')
      ElNotification({
        title:"成功",
        type:"success",
        duration:1200,
      })
      data.dialogVisible = false
    }
  })
}
</script>

<style scoped>
#edit-btn{
  margin-left: 12px;
}
</style>