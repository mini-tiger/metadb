<template>
  <div class="realTable">
    <publicTable :data="state">
      <template #modify="{ scope }">
        <span class="modify modifyLook" @click="modifyLook(scope)">查看</span>
        <span class="modify modifyStop" @click="modify(scope)">编辑</span>
        <span class="modify modifyDelete" @click="modifyDelete(scope)"
          >删除</span
        >
      </template>
    </publicTable>
    <deleteDrop
      :isShowDialog="isShowDialog"
      :title="title"
      :content="content"
      @deleteDialogCancel="dialogCancel"
      @deleteDialogSure="dialogSure"
    ></deleteDrop>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import deleteDrop from '@/components/DeleteDialog/index.vue'
// import { ElMessageBox } from 'element-plus'
// import 'element-plus/es/components/message-box/style' // 按需引入提示框样式不生效解决方案

import publicTable from '@/components/Table/index.vue'
import { table } from '../config/real-config'
export default defineComponent({
  name: 'realTable',
  components: {
    publicTable,
    deleteDrop,
  },
  setup() {
    const router = useRouter()
    const isShowDialog = ref(false)
    const title = ref('')
    const content = ref('')
    const state = reactive({
      table: {
        columns: table.columns,
        tableData: table.tableData,
        isPagination: table.tableData, // 是否展示分页
        pagination: table.pagination,
        tableWidth: '100%', // 表格的宽度
        auxiliary: false,
      },
    })
    // 删除
    const modifyDelete = (data) => {
      console.log(data, 'data--->')
      isShowDialog.value = true
      title.value = '删除实例'
      content.value = `您确定删除${data.row.modeName}实例吗?`
    }
    // 删除确定
    const dialogSure = () => {
      console.log('删除成功')
      isShowDialog.value = false
    }
    //删除取消
    const dialogCancel = () => {
      isShowDialog.value = false
    }
    // 编辑
    const modify = (data) => {
      console.log('停用')
      goNext('emit', data)
    }

    // 查看
    const modifyLook = (data) => {
      console.log('查看')
      goNext('look', data)
    }
    const goNext = (status, data) => {
      console.log(data.row)
      const data1 = {
        type: status,
        params: data.row,
      }
      const params = JSON.stringify(data1)
      console.log(data1, 'data')
      router.push({
        name: `cl-exampleLook`,
        params: { data: params },
      })
    }
    return {
      state,
      isShowDialog,
      title,
      content,
      dialogCancel,
      dialogSure,
      modifyDelete,
      modify,
      modifyLook,
    }
  },
})
</script>

<style lang="scss" scoped>
.realTable {
  margin-top: 10px;
  // height: 497px;
  background: #fff;
}

.modify {
  cursor: pointer;
  font-size: 12px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  margin-right: 15px;
}
.modifyLook {
  color: #595aec;
}
.modifyStop {
  color: #595aec;
}
.modifyDelete {
  color: #fa4d05;
}
:deep(.customClass) {
  width: 460px !important;
  height: 300px !important;
  .el-message-box__header {
    border-bottom: 1px solid #fafafc;
    font-size: 16px;
    font-family: PingFangSC-Medium, PingFang SC;
    font-weight: 500;
    color: #302e6d;
    height: 50px;
  }
}
</style>
