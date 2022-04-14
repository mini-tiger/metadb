<template>
  <TableConfig :leftConfig="leftConfig" :rightConfig="rightConfig">
    <template #special-config>
      <ZHSwitch v-model="switchValue"></ZHSwitch>
    </template>
  </TableConfig>
  <ZHTable
    :tableData="tableData"
    :headerCol="headerCol"
    v-model:currentPage="currentPage"
    v-model:pageSize="pageSize"
    @updateColumn="onUpdateColumn"
    @onSortChange="onSortChange"
  >
    <template v-slot:operation="{ scopeRow }">
      <el-button type="text" @click="onClick(scopeRow.item)">确认</el-button>
      <el-button type="text">反馈</el-button>
      <el-button type="text">建单</el-button>
    </template>
  </ZHTable>
</template>

<script lang="ts" setup>
import { ref, watchPostEffect } from 'vue'
import TableConfig, {
  ILeftConfig,
  IRightConfig,
} from '@/components/Table/TableConfig.vue'
let switchValue = ref(true)
watchPostEffect(() => {
  console.log('switch发生变化，请求后台', switchValue.value)
})
let tableData = ref([
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
])
let headerCol = ref([
  {
    label: '日期',
    prop: 'date',
    headerAlign: 'center',
    align: 'center',
    sortable: 'custom',
  },
  {
    label: '姓名',
    prop: 'name',
  },
  {
    label: '地址',
    prop: 'address',
    sortable: 'custom',
  },
])
let currentPage = ref(1)
let pageSize = ref(10)
const onUpdateColumn = (newColumn) => {
  console.log('newColumn', newColumn)
  //获取到新的列顺序，提交到后台
}
const onClick = (row) => {
  console.log('row', row)
}
watchPostEffect(() => {
  console.log('当前页码变化,向后台发送请求', currentPage.value)
})
watchPostEffect(() => {
  console.log('当前页显示数变化,向后台发送请求', pageSize.value)
})
const onSortChange = ({ order, prop }) => {
  console.log(order, prop)
}
let leftConfig = ref<ILeftConfig[]>([
  {
    name: '批量确认',
    options: {
      bgColor: '#595AEC',
      textColor: '#fff',
      width: '100px',
      height: '30px',
    },
    handler: function () {
      console.log('批量确认')
    },
  },
  {
    name: '批量反馈',
    options: {
      bgColor: '#E4E4FC',
      textColor: '#595AEC',
      width: '100px',
      height: '30px',
      border: 'none',
    },
    handler: function () {
      console.log('批量反馈')
    },
  },
  {
    name: '批量建单',
    options: {
      bgColor: '#E4E4FC',
      textColor: '#595AEC',
      width: '100px',
      height: '30px',
      border: 'none',
    },
    handler: function () {
      console.log('批量建单')
    },
  },
])
let rightConfig = ref<IRightConfig[]>([
  {
    name: '置顶',
    options: {
      bgColor: '#fff',
      textColor: '#595AEC',
      width: '60px',
      height: '26px',
    },
    handler: function () {
      console.log('置顶')
    },
    icon: false,
  },
  {
    name: '取消置顶',
    options: {
      bgColor: '#fff',
      textColor: '#595AEC',
      width: '90px',
      height: '26px',
    },
    handler: function () {
      console.log('取消置顶')
    },
    icon: false,
  },
  {
    name: '列显示',
    options: {
      bgColor: '#fff',
      textColor: '#595AEC',
      width: '85px',
      height: '26px',
    },
    handler: function () {
      console.log('批量建单')
    },
    icon: true,
    showColumn: [{}],
  },
])
</script>
<style lang="scss" scoped></style>
