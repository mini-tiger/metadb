import { reactive, ref } from 'vue'

export const tableData = reactive({
  table: {
    columns: [
      {
        label: '关联类型',
        prop: 'bk_asst_id',
      },
      {
        label: '源至目标约束',
        prop: 'mapping',
      },
      {
        label: '源模型',
        prop: 'bk_obj_id',
      },
      {
        label: '目标模型',
        prop: 'bk_asst_obj_id',
      },
      {
        prop: 'modify',
        label: '操作',
        isSlot: true,
      },
    ],
    tableData: [
      {
        bk_asst_id: '属于',
        mapping: '1-N',
        bk_obj_id: '楼栋模型',
        bk_asst_obj_id: '数据中心模型',
      },
      {
        bk_asst_id: '上下级',
        mapping: 'N-N',
        bk_obj_id: '房间',
        bk_asst_obj_id: 'UPS设备模型',
      },
      {
        bk_asst_id: '属于',
        mapping: '1-N',
        bk_obj_id: '楼栋模型',
        bk_asst_obj_id: '数据中心模型',
      },
      {
        bk_asst_id: '上下级',
        mapping: 'N-N',
        bk_obj_id: '房间',
        bk_asst_obj_id: 'UPS设备模型',
      },
      {
        bk_asst_id: '属于',
        mapping: '1-N',
        bk_obj_id: '楼栋模型',
        bk_asst_obj_id: '数据中心模型',
      },
      {
        bk_asst_id: '上下级',
        mapping: 'N-N',
        bk_obj_id: '房间',
        bk_asst_obj_id: 'UPS设备模型',
      },
    ],
    tableBasics: {},
    pagination: {
      total: 20,
      currentPage: 1,
      pageSize: 10,
    },
    isIndex: true,
    isPagination: true, // 是否展示分页
    tableWidth: '100%', // 表格的宽度
    auxiliary: false,
  },
})

export const deleteDialogData = reactive({
  isShowDialog: false,
  deleteId: '',
  dialogSure() {
    deleteDialogData.isShowDialog = false
  },
  dialogCancel() {
    deleteDialogData.isShowDialog = false
  },
})

export const tableHandle = reactive({
  delete(row) {
    deleteDialogData.isShowDialog = true
    deleteDialogData.deleteId = row.id
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
export const drawerFormData = reactive({
  ruleForm: {
    src_model: '',
    dest_model: '',
    relation_type: '',
    src_to_dest: '',
  },
  rules: {
    src_model: [
      {
        required: true,
        message: '必选',
        trigger: 'blur',
      },
    ],
    dest_model: [
      {
        required: true,
        message: '必选',
        trigger: 'blur',
      },
    ],
    relation_type: [
      {
        required: true,
        message: '必选',
        trigger: 'blur',
      },
    ],
    src_to_dest: [
      {
        required: true,
        message: '必选',
        trigger: 'blur',
      },
    ],
  },
})


export const drawerData = reactive({
  isShowDrawer: false,
  title: '新建关联',
  isDisabled: false,
  closeDrawer() {
    drawerData.isShowDrawer = false;

  },
})

export const buttonData = reactive({
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
