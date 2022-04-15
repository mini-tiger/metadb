<template>
  <Breadcrum></Breadcrum>
  <div class="div-box-form">
    <!--        <publicForm :data="formData" class="formData">-->
    <!--          <template v-slot="data">-->
    <!--            <el-button @click="onSubmit(data.data.formRef, data.data.formData)"-->
    <!--              >搜索</el-button-->
    <!--            >-->
    <!--            <el-button @click="onReset(data.data.formRef, data.data.formData)"-->
    <!--              >重置</el-button-->
    <!--            >-->
    <!--          </template>-->
    <!--        </publicForm>-->
    <SearchForm
      :queryItems="formData.queryItems"
      @onSearch="onSearch"
    ></SearchForm>
  </div>
  <div class="div-box-table">
    <publicTable
      :data="tableData"
      @handleCurrentChange="handleCurrentChange"
      @handleSizeChange="handleSizeChange"
    ></publicTable>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import publicTable from '@/components/Table/index.vue'
// import publicForm from '@/components/Form/index.vue'
import { tableData, formData } from './config/auditConfig'
import SearchForm from '@/components/Search/index.vue'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'
export default defineComponent({
  name: 'OperationAudit',
  components: {
    publicTable,
    // publicForm,
    SearchForm,
  },
  setup() {
    // const onSubmit = (type, data) => {
    //   console.log(type, data)
    //   console.log('data', formData.searchData)
    // }
    // const onReset = (formRef) => {
    //   console.log('重置')
    //   if (!formRef) return
    //   formRef.resetFields()
    // }

    const form = {
      操作对象: 'resourceType',
      动作: 'action',
      操作时间: 'time',
      操作账号: 'operationId',
    }
    const onSearch = () => {
      const list = formData.queryItems.columnList
      list.forEach((item) => {
        if (Object.keys(form).includes(item.label)) {
          formData.formObj[form[item.label]] = item.value
        }
      })
      getList()
    }

    const store = useStore()

    const getList = async () => {
      await store.dispatch(`operate/${types.GET_AUDIT_LIST}`, {
        page: tableData.table.pagination.currentPage,
        size: tableData.table.pagination.pageSize,
        resourceType:
          formData.formObj.resourceType == '全部'
            ? ''
            : formData.formObj.resourceType, //操作对象
        action:
          formData.formObj.action == '全部' ? '' : formData.formObj.action, //动作
        operationStartTime: formData.formObj.time[0] || '', //开始时间
        operationEndTime: formData.formObj.time[1] || '', //结束时间
        operationId: formData.formObj.operationId || '', //操作账号
      })

      tableData.table.tableData = store.state.operate.auditData.content
      tableData.table.pagination.total =
        store.state.operate.auditData.totalElements
    }

    //改变页码
    const handleCurrentChange = (val) => {
      tableData.table.pagination.currentPage = val
      getList()
    }
    //改变每页条数
    const handleSizeChange = (val) => {
      tableData.table.pagination.pageSize = val
      getList()
    }

    onMounted(() => {
      getList()
    })

    return {
      tableData,
      handleCurrentChange,
      handleSizeChange,
      formData,
      onSearch,
    }
  },
})
</script>

<style scoped lang="scss">
@import '../../../../assets/style/overallStyle.scss';
</style>
