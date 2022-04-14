import { markRaw, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { nextTick } from 'vue'

export const formData = reactive({
  searchForm: [
    {
      type: 'Input',
      label: '分类名称',
      prop: 'name',
    },
    {
      type: 'Select',
      label: '目前状态',
      prop: 'status',
      options: [
        { label: '正常', value: 1 },
        { label: '停用', value: 2 },
      ],
    },
  ],
  searchData: {
    name: '',
    status: '',
  },
  formSlot: true,
  formArray: true,
})

export const buttonData = reactive({
  buttonIn: {
    name: '导入',
    className: 'roundBackCal',
    round: true,
  },
  buttonOut: {
    name: '导出',
    className: 'roundBackCal',
    round: true,
  },
  buttonAdd: {
    name: '新建根目录',
    className: 'roundBlueIcon',
    round: true,
    icon: markRaw(Plus),
  },
  buttonClose: {
    name: '关闭',
    className: 'roundBackCal',
    round: true,
  },
  buttonCancel: {
    name: '取消',
    className: 'roundBackCal',
    round: true,
  },
  buttonSure: {
    name: '确定',
    className: 'roundBackBlue',
    round: true,
  },
})

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
export const tableData = reactive({
  table: {
    columns: [
      {
        label: 'ID',
        prop: 'id',
      },
      {
        label: '分类名称',
        prop: 'classificationName',
      },
      {
        label: '排序',
        prop: 'orderBy',
      },
      {
        label: '使用次数',
        prop: 'useCount',
        isSlot: true,
      },
      {
        label: '状态',
        prop: 'status',
        isSlot: true,
      },
      {
        prop: 'modify',
        label: '操作',
        isSlot: true,
      },
    ],
    tableData: [],
    tableBasics: {
      defaultExpandAll: true,
      lazy: false,
      rowKey: 'id',
      load: load,
      treeProps: { children: 'children', hasChildren: 'hasChildren' },
    },
    isPagination: false, // 是否展示分页
    tableWidth: '100%', // 表格的宽度
    auxiliary: false,
  },
})

export const drawerFormData = reactive({
  isRoot: true, //是否根目录
  ruleForm: {
    id: '',
    parentId: 0, //上级分类id
    parentName: '无', //上级分类名称
    classificationName: '', //分类名称
    orderBy: 1, //排序
    status: '1', //状态：1正常 2停用
    describe: '', //描述
    level: 0,
  },
  rules: {
    classificationName: [
      {
        required: true,
        message: '必输',
        trigger: 'blur',
      },
    ],
    orderBy: [
      {
        required: true,
        message: '必输',
        trigger: 'blur',
      },
    ],
    status: [
      {
        required: true,
        message: '必选',
        trigger: 'blur',
      },
    ],
  },
})

export const UploadData = reactive({
  isShowUpload: false,
  title: '导入',
  //取消
  btnCancel() {
    UploadData.isShowUpload = false
  },
  //确认
  btnSure() {
    UploadData.isShowUpload = false
  },
  //导入
  upLoad() {
    UploadData.title = '导入'
    UploadData.isShowUpload = true
  },
  //导出
  downLoad() {
    UploadData.title = '导出'
    UploadData.isShowUpload = true
  },
})
export const drawerData = reactive({
  isShowDrawer: false,
  title: '新建分类',
  isDisabled: false,
  closeDrawer(val) {
    drawerData.isShowDrawer = val
  },
})

export const dialogData = reactive({
  isShowDialog: false,
  deleteId: '',
  title: '删除分类',
  content: '确认删除该分类吗？',
  deleteDialogCancel() {
    dialogData.isShowDialog = false
  },
})
export const addRoot = () => {
  drawerData.isShowDrawer = true
  drawerFormData.ruleForm.parentId = 0
  drawerFormData.ruleForm.parentName = '无'
}
export const tableHandle = reactive({
  //新增
  add(row) {
    drawerData.title = '新建分类'
    drawerData.isShowDrawer = true
    drawerData.isDisabled = false //禁用
    drawerFormData.ruleForm.parentId = row.id
    drawerFormData.ruleForm.parentName = row.classificationName
    drawerFormData.ruleForm.level = parseInt(row.level) + 1
  },
  //查看
  look(row) {
    drawerData.title = '查看分类'
    drawerData.isShowDrawer = true //展示抽屉
    drawerData.isDisabled = true //禁用
    row.parentName = row.parentName === '-' ? '无' : row.parentName
    drawerFormData.ruleForm = { ...row }
  },
  //删除
  del(row) {
    console.log('删除', row)
    if (row.useCount === '0') {
      dialogData.isShowDialog = true
      dialogData.deleteId = row.id
    }
  },
  //编辑
  modify(row) {
    drawerData.title = '编辑分类'
    drawerData.isShowDrawer = true
    drawerData.isDisabled = false //禁用
    nextTick(() => {
      row.parentName = row.parentName === '-' ? '无' : row.parentName
      row.orderBy = parseInt(row.orderBy)
      drawerFormData.ruleForm = { ...row }
    }).then()
  },
  toModel() {
    console.log('CI模型页面')
  },
})

export const formObj = {
  classificationName: '',
  status: '',
}
