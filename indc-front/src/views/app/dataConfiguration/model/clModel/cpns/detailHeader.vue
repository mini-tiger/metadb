<template>
  <div class="div-box">
    <el-row class="model-row">
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">模型分类 :</div>
          <div class="model-content">设备/强电/UPS</div>
          <div class="model-icon" @click="editModel" v-if='type!=="look"'>
            <svg-icon icon-class="edit"></svg-icon>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">模型名称 :</div>
          <div class="model-content">UPS设备模型</div>
          <div class="model-icon" @click="editModel" v-if='type!=="look"'>
            <svg-icon icon-class="edit"></svg-icon>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">实例数量 :</div>
          <div class="model-content no-zero">1</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">选择图标 :</div>
          <div class="model-icon">
            <svg-icon icon-class="nav1"></svg-icon>
          </div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">ID :</div>
          <div class="model-content">MX00001</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">创建人 :</div>
          <div class="model-content">刘工</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">创建时间 :</div>
          <div class="model-content">2022/03/09 21:12:36</div>
        </div>
      </el-col>
      <el-col :span="6">
        <div class="model-box">
          <div class="model-label">模型描述 :</div>
          <div class="model-content">
            内容描述内容描述内容描述述内容描述内容描述内容描述内容描述内容描述述内容描述内容描述
          </div>
        </div>
      </el-col>
    </el-row>

    <newModel
      :isNewModel="isShowModel"
      @cancelModel="cancelModel"
      @ModelOk="ModelOk"
    ></newModel>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import newModel from './newModel.vue'
import {useRoute} from 'vue-router'

export default defineComponent({
  name: 'detailHeader',
  components: {
    newModel,
  },
  setup() {
    const isShowModel = ref(false)
    const editModel = () => {
      console.log('编辑')
      isShowModel.value = true
    }

    const cancelModel = () => {
      isShowModel.value = false
    }

    const ModelOk = () => {
      isShowModel.value = false
    }
    const route = useRoute()

    if (route.params && !localStorage.getItem('rowData')) {
      const data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    const type = JSON.parse(localStorage.getItem('rowData')).type || ''
    return {
      editModel,
      isShowModel,
      cancelModel,
      ModelOk,
      type
    }
  },
})
</script>

<style scoped lang="scss">
.div-box {
  padding: 25px 20px 0 20px;
  box-sizing: border-box;
  background: #ffffff;
  border-radius: 8px;
  margin-bottom: 10px;
}
.model-row {
  display: flex;
  flex-wrap: wrap;
}
.model-box {
  display: flex;
  align-items: center;
  margin-bottom: 30px;
}
.model-label {
  width: 80px;
  font-size: 14px;
  color: #302e6d;
  text-align: right;
}
.model-content {
  font-size: 14px;
  color: #302e6d;
  //flex: 1;
  margin-left: 10px;
  max-width: 300px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.model-content.no-zero {
  color: #595aec;
}
.model-icon {
  font-size: 14px;
  margin-left: 12px;
}
</style>
