import { reactive } from 'vue'
export const tableData = reactive({
  table: {
    columns: [
      {
        label: '操作对象',
        prop: 'resourceType',
      },
      {
        label: '动作',
        prop: 'action',
      },
      {
        label: '操作描述',
        prop: 'operationDetail',
      },
      {
        label: '操作时间',
        prop: 'operationTime',
      },
      {
        label: '操作账号',
        prop: 'operationId',
      },
    ],
    tableData: [],
    isPagination: true, // 是否展示分页
    pagination: {
      total: 0,
      currentPage: 1,
      pageSize: 10,
    },
    tableWidth: '100%', // 表格的宽度
    auxiliary: false,
  },
})

export const formData = reactive({
  queryItems: {
    gutter: 0,
    columnList: [
      {
        type: 'select',
        label: '操作对象',
        value: '',
        options: [
          { label: '全部', value: '全部' },
          { label: 'CI模型', value: 'CI模型' },
          { label: 'CI模型关系', value: 'CI模型关系' },
          { label: 'CI模型分类', value: 'CI模型分类' },
          { label: 'CI实例', value: 'CI实例' },
          { label: 'CI分类', value: 'CI分类' },
          { label: '关系类型', value: '关系类型' },
        ],
      },
      {
        type: 'select',
        label: '动作',
        value: '',
        options: [
          { label: '全部', value: '全部' },
          { label: '新增', value: '新增' },
          { label: '修改', value: '修改' },
          { label: '删除', value: '删除' },
        ],
      },
      {
        type: 'datepicker',
        label: '操作时间',
        value: '',
      },
      { type: 'input', label: '操作账号', value: '' },
    ],
  },
  formObj: {
    resourceType: '',
    action: '',
    time: '',
    operationId: '',
  },
})
