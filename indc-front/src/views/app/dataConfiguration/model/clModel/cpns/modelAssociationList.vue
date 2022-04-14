<template>
  <publicTable
    :data="tableData"
    @handleSizeChange="tableHandle.handleSizeChange"
    @handleCurrentChange="tableHandle.handleCurrentChange"
  >
    <template #modify="{ scope }">
      <span class="table-delete" @click="tableHandle.delete(scope.row)"
            v-if='type!=="look"'
        >删除</span
      >
      <span class='table-be-relevance' v-else>删除</span>

    </template>
  </publicTable>
  <DeleteDialog
    :isShowDialog="deleteDialogData.isShowDialog"
    title="删除模型关联"
    content="确认删除两个模型间的关系？"
    @deleteDialogCancel="deleteDialogData.dialogCancel"
    @deleteDialogSure="deleteDialogData.dialogSure"
  ></DeleteDialog>
  <newAssociationDrawer
    :is-show-drawer="drawerData.isShowDrawer"
    @closeDrawer="drawerData.closeDrawer"
  ></newAssociationDrawer>
</template>

<script>
import { defineComponent } from 'vue'
import publicTable from '@/components/Table/index.vue'
import DeleteDialog from '@/components/DeleteDialog/index.vue'
import {
  tableData,
  deleteDialogData,
  tableHandle,
  drawerFormData,
  drawerData,
} from '../config/model-association'
import newAssociationDrawer from './newAssociationDrawer'
import { useRoute } from 'vue-router'
export default defineComponent({
  name: 'modelAssociationList',
  components: {
    publicTable,
    DeleteDialog,
    newAssociationDrawer,
  },
  setup() {
    const route = useRoute()

    if (route.params && !localStorage.getItem('rowData')) {
      const data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    const type = JSON.parse(localStorage.getItem('rowData')).type || ''
    return {
      tableData,
      deleteDialogData,
      tableHandle,
      drawerData,
      drawerFormData,
      type
    }
  },
})
</script>

<style scoped></style>
