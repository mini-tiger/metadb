<template>
  <div class="model-table">
    <publicTable :data="modelTable">
      <template #objIcon="{ scope }">
        <svgIcon :icon-class="`${scope.row.objIcon}`"></svgIcon>
      </template>
      <template #insCount="{ scope }">
        <span v-if="scope.row.insCount" class="is-no-zero" @click="toExample">{{
          scope.row.insCount
        }}</span>
        <span v-else>{{ scope.row.insCount }}</span>
      </template>

      <template #ispaused="{ scope }">
        <span>{{ scope.row.ispaused ? '停用' : '启用' }}</span>
      </template>
      <template #modify="{ scope }">
        <span class="modify modifyLook" @click="modifyLook(scope)">查看</span>
        <span class="modify modifyLook" @click="modify(scope)">编辑</span>
        <span class="modify modifyStop" @click="modifyStop(scope)">停用</span>
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
import { defineComponent, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import deleteDrop from '@/components/DeleteDialog/index.vue'
import publicTable from '@/components/Table/index.vue'
import { clModelTable } from '../config/cl-model'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'

export default defineComponent({
  components: {
    publicTable,
    deleteDrop,
  },
  setup() {
    const modelTable = ref({ table: clModelTable })
    let isShowDialog = ref(false)
    const title = ref('')
    const content = ref('')
    const router = useRouter()
    const modifyLook = (data) => {
      console.log('我是查看', data)
      goNext('look', data)
    }
    const modifyStop = (data) => {
      console.log('我是停用')
    }
    const modifyDelete = (data) => {
      // console.log('我是删除', data)
      isShowDialog.value = true
      title.value = '删除模型'
      content.value = `您确定删除${data.row.modelName}模型吗?`
    }
    const modify = (data) => {
      console.log('data', '我是编辑------》', data)
      goNext('edit', data)
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
        name: `modelDetail`,
        params: { data: params },
      })
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
    const store = useStore()
    const getList = async () => {
      await store.dispatch(`modelList/${types.GET_MODEL_LIST}`, {
        pageNum: modelTable.value.table.pagination.currentPage,
        pageSize: modelTable.value.table.pagination.pageSize,
        classificationId: '', //分类ID
        objName: '', //模型名称
        ispaused: '', //状态（true-停用 false-启用）
        description: '', //描述
        creator: '', //创建人
        createStartTime: '', //创建开始时间
        createEndTime: '', //创建结束时间
        label: '', //标签
      })
      modelTable.value.table.tableData = store.state.modelList.data.content;
      modelTable.value.table.pagination.total = store.state.modelList.data.totalElements;

    }
    onMounted(() => {
      getList()
    })

    const toExample = () => {
      console.log('去实例列表')
    }
    return {
      modelTable,
      isShowDialog,
      title,
      content,
      dialogCancel,
      dialogSure,
      modify,
      modifyLook,
      modifyStop,
      modifyDelete,
      toExample,
    }
  },
})
</script>

<style lang="scss" scoped>
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
.is-no-zero{
  color:#595aec ;
}
</style>
