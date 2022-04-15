import { markRaw, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'

export const buttonData = reactive({
  buttonAdd: {
    name: '新建关系类型',
    className: 'roundBlueIcon',
    round: true,
    icon: markRaw(Plus),
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
  buttonClose: {
    name: '关闭',
    className: 'roundBackCal',
    round: true,
  },
})

export const formData = reactive({
  searchForm: [
    {
      type: 'Input',
      label: '关系名称',
      prop: 'bk_asst_name',
    },
    {
      type: 'Input',
      label: '创建人',
      prop: 'founder',
    },
    {
      type: 'Daterange',
      label: '创建时间',
      prop: 'time',
    },
  ],
  searchData: {
    name: '',
    status: '',
  },
  formSlot: true,
  formArray: true,
})

export const tableData = reactive({
  table: {
    columns: [
      {
        label: 'ID',
        prop: 'id',
      },
      {
        label: '关系名称',
        prop: 'bk_asst_name',
      },
      {
        label: '源至目标描述',
        prop: 'src_des',
      },
      {
        label: '目标至源描述',
        prop: 'dest_des',
      },
      {
        label: '模型使用数量',
        prop: 'bk_model_account',
        isSlot: true,
      },
      {
        label: 'CI使用数量',
        prop: 'bk_cl_account',
        isSlot: true,
      },
      {
        label: '创建人',
        prop: 'founder',
      },
      {
        label: '创建时间',
        prop: 'time',
      },
      {
        prop: 'modify',
        label: '操作',
        isSlot: true,
      },
    ],
    tableData: [
      {
        id: 1,
        bk_asst_id: 'belong',
        bk_asst_name: '属于',
        bk_model_account: 2,
        bk_cl_account: 4,
        src_des: '属于',
        dest_des: '包含',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '1',
      },
      {
        id: 2,
        bk_asst_id: 'group',
        bk_asst_name: '组成',
        bk_model_account: 0,
        bk_cl_account: 2,
        src_des: '组成',
        dest_des: '组成于',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '2',
      },
      {
        id: 3,
        bk_asst_id: 'bk_mainline',
        bk_asst_name: '拓扑组成',
        bk_model_account: 2,
        bk_cl_account: 1,
        src_des: '组成',
        dest_des: '组成于',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '2',
      },
      {
        id: 4,
        bk_asst_id: 'run',
        bk_asst_name: '运行',
        bk_model_account: 1,
        bk_cl_account: 2,
        src_des: '运行于',
        dest_des: '运行',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '3',
      },
      {
        id: 5,
        bk_asst_id: 'connect',
        bk_asst_name: '上联',
        bk_model_account: 0,
        bk_cl_account: 0,
        src_des: '上联',
        dest_des: '下联',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '2',
      },
      {
        id: 6,
        bk_asst_id: 'default',
        bk_asst_name: '默认关联',
        bk_model_account: 2,
        bk_cl_account: 0,
        src_des: '关联',
        dest_des: '关联',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '1',
      },
      {
        id: 7,
        bk_asst_id: 'shangxiaji',
        bk_asst_name: '上下级关系',
        bk_model_account: 2,
        bk_cl_account: 1,
        src_des: '子级',
        dest_des: '父级',
        direction: 'src_to_dest',
        founder: 'admin',
        time: '2022/03/09 21:12:36',
        is_direction: '2',
      },
    ],

    pagination: {
      total: 20,
      currentPage: 1,
      pageSize: 10,
    },
    isPagination: true, // 是否展示分页
    tableWidth: '100%', // 表格的宽度
    auxiliary: false,
  },
})

export const drawerFormData = reactive({
  isRoot: true, //是否根目录
  ruleForm: {
    bk_asst_name: '',
    src_des: '',
    dest_des: '',
    is_direction: '1',
  },
  rules: {
    bk_asst_name: [
      {
        required: true,
        message: '必输',
        trigger: 'blur',
      },
    ],
    src_des: [
      {
        required: true,
        message: '必输',
        trigger: 'blur',
      },
    ],
    dest_des: [
      {
        required: true,
        message: '必输',
        trigger: 'blur',
      },
    ],
  },
})

export const drawerData = reactive({
  isShowDrawer: false,
  title: '新建关系类型',
  isDisabled: false,
  closeDrawer() {
    drawerData.isShowDrawer = false
  },
})

export const deleteDialogData = reactive({
  isShowDialog: false,
  deleteId: '',
  cancel() {
    deleteDialogData.isShowDialog = false
  },
  sure() {
    deleteDialogData.isShowDialog = false
    console.log(deleteDialogData.deleteId)
  },
})
export const tableHandle = reactive({
  look(row) {
    drawerData.isShowDrawer = true
    drawerData.isDisabled = true
    drawerFormData.ruleForm = { ...row }
  },
  modify(row) {
    drawerData.isShowDrawer = true
    drawerData.isDisabled = false
    this.$nextTick(() => {
      drawerFormData.ruleForm = { ...row }
    })
  },
  delete(row) {
    deleteDialogData.isShowDialog = true
    deleteDialogData.deleteId = row.id
  },
  toModel() {
    console.log('CI模型关系')
  },
  toCI() {
    console.log('CI关系')
  },
  //每页条数
  handleSizeChange(val) {
    console.log('size', val)
  },
  //翻页
  handleCurrentChange(val) {
    console.log('page', val)
  },
})
