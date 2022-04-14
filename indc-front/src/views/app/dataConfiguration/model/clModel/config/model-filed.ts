import { markRaw, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'
export const state = reactive({
  dataList: [
    {
      bk_group_name: '默认段落',
      bk_group_id: 'default',
      is_collapse: false,
      index: 0,
      list: [
        { id: 1, name: '实例名称', is_collapse: false },
        { id: 2, name: '实例名称2', is_collapse: false },
        { id: 3, name: '实例名称3', is_collapse: false },
        { id: 4, name: '实例名称4', is_collapse: false },
        { id: 4, name: '', is_collapse: false },
      ],
    },
    {
      bk_group_name: '基础信息',
      bk_group_id: 'info',
      is_collapse: false,
      index: 1,
      list: [
        { id: 1, name: '名称', is_collapse: false },
        { id: 2, name: '面积', is_collapse: false },
        { id: 3, name: '地址', is_collapse: false },
        { id: 4, name: '设备数量', is_collapse: false },
        { id: 3, name: '市电容量', is_collapse: false },
        { id: 4, name: '设备数量1', is_collapse: false },
        { id: 4, name: '', is_collapse: false },
      ],
    },
    {
      bk_group_name: '标签信息',
      bk_group_id: 'tag',
      is_collapse: true,
      index: 2,
      list: [
        { id: 1, name: '名称', is_collapse: true },
        { id: 2, name: '面积', is_collapse: true },
        { id: 3, name: '地址', is_collapse: true },
        { id: 4, name: '设备数量', is_collapse: true },
        { id: 3, name: '市电容量', is_collapse: true },
        { id: 4, name: '实例名称10', is_collapse: true },
        { id: 4, name: '', is_collapse: true },
      ],
    },
    {
      bk_group_name: '运维信息',
      bk_group_id: 'operations',
      is_collapse: false,
      index: 3,
      list: [
        { id: 1, name: '实例名称', is_collapse: false },
        { id: 2, name: '实例名称2', is_collapse: false },
        { id: 3, name: '实例名称3', is_collapse: false },
        { id: 4, name: '实例名称4', is_collapse: false },
        { id: 4, name: '', is_collapse: false },
      ],
    },
  ],

  buttonNew: {
    name: '新建段落',
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
  isShowDelete: false,
  title: '删除段落',
  content: '确认删除段落？段落下的所有字段将一起删除！',
  deleteId: '',
  deleteDialogCancel() {
    state.isShowDelete = false
  },
  deleteDialogSure() {
    state.isShowDelete = false
  },
})

export const handleMethods = reactive({
  //新建段落
  newSection() {
    console.log('新建段落')
    drawerData.isShowDrawer = true
    drawerData.isDisabled = false
  },
  //删除段落
  delSection(data) {
    state.isShowDelete = true
    state.title = '删除段落'
    state.content = '确认删除该段落吗？'
    console.log('删除段落', data)
  },
  //移动段落
  moveSection(e) {

    console.log('add', e)

    console.log('list', state.dataList)
  },
  //编辑段落
  editSection(data) {
    console.log('编辑段落', data)
    drawerData.title = '编辑段落'
    drawerData.isShowDrawer = true
    drawerData.isDisabled = true
    this.$nextTick(() => {
      drawerFormData.ruleForm.bk_group_id = data.bk_group_id
      drawerFormData.ruleForm.bk_group_name = data.bk_group_name
      drawerFormData.ruleForm.is_collapse = data.is_collapse
    })
  },

  //新建字段
  newFiled(data) {
    console.log('新建字段', data)
    if (data.is_collapse) {
      //标签
      drawerFiledData.isTag = true
      drawerFiledData.title = '新建标签'
    } else {
      console.log('新建字段')
      drawerFiledData.title = '新建字段'
      drawerFiledData.isTag = false
    }
    drawerFiledData.isShowFiled = true
    drawerFiledData.isEditFiled = false
    drawerFiledData.filedType = ''
  },

  //移动字段
  moveFiled(e, val) {
    console.log('e', e)
    console.log('val', val)
  },
  //删除字段
  delFiled(data) {
    if (data.is_collapse) {
      //标签
      state.title = '删除标签'
      state.content = '确认删除该标签吗？'
    } else {
      state.title = '删除字段'
      state.content = '确认删除该字段吗？'
    }
    state.isShowDelete = true
    console.log('删除字段', data)
  },
  //编辑字段
  editFiled(data) {
    console.log('编辑字段', data)
    if (data.is_collapse) {
      //标签
      drawerFiledData.title = '编辑标签'
      drawerFiledData.isTag = true
    } else {
      console.log('编辑字段')
      drawerFiledData.title = '编辑字段'
      drawerFiledData.isTag = false
    }
    drawerFiledData.isShowFiled = true
    drawerFiledData.isEditFiled = true
  },
})
const validateLabel = (rule: any, value: any, callback: any) => {
  if (!value) {
    callback(new Error('必输'))
  } else {
    const reg = /^[a-zA-Z]+$/
    if (!reg.test(value)) {
      callback(new Error('唯一标识仅支持英文'))
    }
    callback()
  }
}

export const drawerFormData = reactive({
  ruleForm: {
    bk_group_id: '', //唯一标识？
    bk_group_name: '', //段落名称
    is_collapse: false, //是否标签
    bk_obj_id: '', //所属模型
    bk_group_index: '', //排序
  },
  rules: {
    bk_group_id: [
      {
        required: true,
        validator: validateLabel,
        trigger: 'blur',
      },
    ],
    bk_group_name: [
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
  title: '新建段落',
  isDisabled: false,
  closeDrawer() {
    drawerData.isShowDrawer = false
  },
  drawerButtonSure: async (ruleFormRef, formData) => {
    if (!ruleFormRef) return
    await ruleFormRef.validate((valid, fields) => {
      if (valid) {
        console.log('submit!', formData)

        const data = {
          bk_group_name: formData.bk_group_name,
          bk_group_id: formData.bk_group_id,
          is_collapse: formData.is_collapse,
          index: 3,
          list: [{ id: 4, name: '', is_collapse: formData.is_collapse }],
        }
        state.dataList.push(data)

        drawerData.isShowDrawer = false
        ruleFormRef.resetFields()
      } else {
        console.log('error submit!', fields)
      }
    })
  },
  drawerButtonCancel: (ruleFormRef) => {
    if (!ruleFormRef) return
    ruleFormRef.resetFields()
    drawerData.isShowDrawer = false
  },
})

export const drawerFiledData = reactive({
  isShowFiled: false,
  isEditFiled: false,
  filedType: '',
  isTag: false,
  title: '新建字段',
  btnSure(data) {
    console.log('data', data)
    drawerFiledData.isShowFiled = false
  },
  btnCancel() {
    drawerFiledData.isShowFiled = false
  },
})
