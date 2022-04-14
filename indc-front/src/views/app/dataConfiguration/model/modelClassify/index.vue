<template>
  <Breadcrum></Breadcrum>
  <div class="absolute-button">
    <publicButton
      :data="buttonData.buttonIn"
      @onButton="UploadData.upLoad"
    ></publicButton>
    <publicButton
      :data="buttonData.buttonOut"
      @onButton="UploadData.downLoad"
    ></publicButton>
    <publicButton
      :data="buttonData.buttonAdd"
      @onButton="addRoot"
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
    <SearchForm :queryItems="queryItems" @onSearch="onSearch" />
  </div>
  <div class="div-box-table">
    <publicTable :data="tableData">
      <template #status="{ scope }">
        <span>{{ scope.row.status == '2' ? '停用' : '启用' }}</span>
      </template>
      <template #useCount="{ scope }">
        <span class="use-is-zero" v-if="scope.row.useCount == 0">
          {{ scope.row.useCount }}</span
        >
        <span class="use-is-noZero" v-else @click="toModel">{{
          scope.row.useCount
        }}</span>
      </template>
      <template #modify="{ scope }">
        <span class="table-look" @click="look(scope.row)">查看</span>
        <span class="table-modify" @click="modify(scope.row)">编辑</span>
        <span class="table-add" @click="add(scope.row)">新增</span>
        <span>
          <span
            v-if="scope.row.useCount === '0'"
            class="table-delete"
            @click="del(scope.row)"
            >删除</span
          >
          <span v-else class="table-default">删除</span>
        </span>
      </template>
    </publicTable>
  </div>
  <ModelPelDrawer
    :is-show-drawer="drawerData.isShowDrawer"
    :drawerTitle="drawerData.title"
    :isDisabled="drawerData.isDisabled"
    isType="model"
    @closeDrawer="drawerData.closeDrawer"
  >
  </ModelPelDrawer>
  <DeleteDialog
    :isShowDialog="dialogData.isShowDialog"
    :title="dialogData.title"
    :content="dialogData.content"
    @deleteDialogCancel="dialogData.deleteDialogCancel"
    @deleteDialogSure="deleteDialogSure"
  ></DeleteDialog>
  <UpLoadDrawer
    :is-show-up-load="UploadData.isShowUpload"
    :title="UploadData.title"
    @btnCancel="UploadData.btnCancel"
    @btnSure="UploadData.btnSure"
  ></UpLoadDrawer>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, toRefs } from 'vue'
import publicTable from '@/components/Table/index.vue'
// import publicForm from '@/components/Form/index.vue'
import publicButton from '@/components/Button/index.vue'
import ModelPelDrawer from './components/ModelPelDrawer.vue'
import SearchForm from '@/components/Search/index.vue'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'
import {
  buttonData,
  tableData,
  formData,
  drawerFormData,
  UploadData,
  tableHandle,
  drawerData,
  dialogData,
  addRoot,
  formObj,
} from './config/config'
import DeleteDialog from '@/components/DeleteDialog/index.vue'
import UpLoadDrawer from '@/components/UpLoadDrawer/index.vue'
export default defineComponent({
  name: 'modelClassify',
  components: {
    publicTable,
    // publicForm,
    publicButton,
    ModelPelDrawer,
    DeleteDialog,
    UpLoadDrawer,
    SearchForm,
  },
  setup() {
    const onSubmit = () => {
      console.log('submit')
    }
    const onReset = (formRef) => {
      console.log('reset')
      if (!formRef) return
      formRef.resetFields()
    }
    const load = (row, treeNode, resolve) => {
      setTimeout(() => {
        resolve([
          {
            id: 'MX00001',
            name: '不间断电源',
            lastName: '上级2',
            sort: '1',
            useCount: 10,
            status: '1',
          },
          {
            id: 'MX00001',
            name: '变压器',
            lastName: '上级2',
            sort: '1',
            useCount: 5,
            status: '0',
          },
        ])
      }, 1000)
    }
    let queryItems = ref({
      gutter: 0,
      columnList: [
        { type: 'input', label: '分类名称', value: '' },
        {
          type: 'select',
          label: '状态',
          value: '',
          options: [
            { label: '全部', value: '全部' },
            { label: '启用', value: '1' },
            { label: '停用', value: '2' },
          ],
        },
      ],
    })

    const store = useStore()
    const getList = async () => {
      if (formObj.classificationName || formObj.status) {
        //搜索
        tableData.table.tableBasics.defaultExpandAll = true
      } else {
        tableData.table.tableBasics.defaultExpandAll = false
      }
      await store.dispatch(
        `modelClassify/${types.GET_MODEL_CLASSIFY_LIST}`,
        formObj,
      )
      tableData.table.tableData = store.state.modelClassify.data
    }

    const onSearch = () => {
      formObj.classificationName = queryItems.value.columnList[0].value
      formObj.status =
        queryItems.value.columnList[1].value == '全部'
          ? ''
          : queryItems.value.columnList[1].value
      getList()
    }

    //删除
    const deleteDialogSure = async () => {
      dialogData.isShowDialog = false
      console.log('deleteId', dialogData.deleteId)
      await store.dispatch(
        `modelClassify/${types.DEL_MODEL_CLASSIFY}`,
        dialogData.deleteId,
      )
      tableData.table.tableData = store.state.modelClassify.data
    }

    onMounted(() => {
      getList()
    })
    return {
      onSubmit,
      onReset,
      ...toRefs(tableHandle),
      drawerData,
      buttonData,
      tableData,
      formData,
      drawerFormData,
      addRoot,
      dialogData,
      UploadData,
      load,
      queryItems,
      onSearch,
      deleteDialogSure,
    }
  },
})
</script>

<style scoped lang="scss"></style>
