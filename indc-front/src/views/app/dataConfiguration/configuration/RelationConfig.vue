<template>
  <Breadcrumb></Breadcrumb>
  <div class="absolute-button">
    <publicButton
      :data="buttonData.buttonAdd"
      @onButton="addRelationConfig"
    ></publicButton>
  </div>
  <div class="div-box-form">
<!--    <public-form :data="formData">-->
<!--      <template v-slot="data">-->
<!--        <el-button @click="onSubmit(data.data.formRef, data.data.formData)"-->
<!--          >搜索</el-button-->
<!--        >-->
<!--        <el-button @click="onReset(data.data.formRef)">重置</el-button>-->
<!--      </template>-->
<!--    </public-form>-->
    <SearchForm :queryItems="queryItems" @onSearch="onSearch"/>
  </div>
  <div class="div-box-table">
    <publicTable
      :data="tableData"
      @handleSizeChange="handleSizeChange"
      @handleCurrentChange="handleCurrentChange"
      @columnsChange="columnsChange"
    >
      <template #bk_model_account="{ scope }">
        <span class="use-is-zero" v-if="scope.row.bk_model_account === 0">
          {{ scope.row.bk_model_account }}</span
        >
        <span class="use-is-noZero" v-else @click="toModel">{{
          scope.row.bk_model_account
        }}</span>
      </template>
      <template #bk_cl_account="{ scope }">
        <span class="use-is-zero" v-if="scope.row.bk_cl_account === 0">
          {{ scope.row.bk_cl_account }}</span
        >
        <span class="use-is-noZero" v-else @click="toCI">{{
          scope.row.bk_cl_account
        }}</span>
      </template>
      <template #modify="{ scope }">
        <span class="table-look" @click="look(scope.row)">查看</span>
        <span class="table-modify" @click="modify(scope.row)">编辑</span>
        <span class="table-delete" @click="delete scope.row">删除</span>
      </template>
    </publicTable>
  </div>
  <ConfigPelDrawer
    :is-show-drawer="drawerData.isShowDrawer"
    :drawerTitle="drawerData.title"
    :isDisabled="drawerData.isDisabled"
    isType="config"
    @closeDrawer="drawerData.closeDrawer"
  >
  </ConfigPelDrawer>
  <DeleteDialog
    content="确认删除该关系类型么？"
    title="删除关系类型"
    :is-show-dialog="deleteDialogData.isShowDialog"
    @deleteDialogCancel="deleteDialogData.cancel"
    @deleteDialogSure="deleteDialogData.sure"
  ></DeleteDialog>
</template>

<script lang="ts">
import { defineComponent, toRefs, nextTick, onUpdated, watch, ref } from 'vue'
import publicButton from '@/components/Button/index.vue'
// import publicForm from '@/components/Form/index.vue'
import publicTable from '@/components/Table/index.vue'
import ConfigPelDrawer from './components/ConfigPelDrawer.vue'
import DeleteDialog from '@/components/DeleteDialog/index.vue'
import SearchForm from '@/components/Search/index.vue'
import {
  buttonData,
  formData,
  tableData,
  drawerFormData,
  drawerData,
  deleteDialogData,
  tableHandle,
} from '@/views/app/dataConfiguration/configuration/config/relationConfig'

export default defineComponent({
  name: 'RelationConfig',
  components: {
    publicButton,
    // publicForm,
    publicTable,
    ConfigPelDrawer,
    DeleteDialog,
    SearchForm
  },
  setup() {
    const onSubmit = () => {
      console.log('搜索')
    }
    const onReset = (formRef) => {
      if (!formRef) return
      formRef.resetFields()
    }
    //新建
    const addRelationConfig = () => {
      drawerData.isShowDrawer = true
      drawerData.isDisabled = false
    }
    const columnsChange = (val) => {
      nextTick(() => {
        tableData.table.columns = val
      })
    }
    let queryItems = ref({
      gutter: 0,
      columnList: [
        { type: 'input', label: '关系名称', value: '' },
        { type: 'input', label: '创建人', value: '' },
        { type: 'datepicker', label: '创建时间', value: '' },
      ],
    })
    const onSearch = () => {
      console.log(queryItems.value.columnList)
    }
    return {
      addRelationConfig,
      buttonData,
      formData,
      onSubmit,
      onReset,
      tableData,
      ...toRefs(tableHandle),
      drawerFormData,
      drawerData,
      deleteDialogData,
      columnsChange,
      queryItems,
      onSearch
    }
  },
})
</script>

<style scoped></style>
