import { markRaw, reactive } from 'vue'
import { Edit } from '@element-plus/icons-vue'

export const searchData = {
  columnList: [
    {
      type: 'modelKind',
      label: '模型分类',
      value: '',
      // width: '380px',
      data: [
        {
          id: 1,
          label: 'Level one 1',
          children: [
            {
              id: 4,
              label: 'Level two 1-1',
              parentId: 1,
              children: [
                {
                  id: 9,
                  label: 'Level three 1-1-1',
                  parentId: 1,
                },
                {
                  id: 10,
                  label: 'Level three 1-1-2',
                  parentId: 1,
                },
              ],
            },
          ],
        },
        {
          id: 2,
          label: 'Level one 2',
          parentId: 2,
          children: [
            {
              id: 5,
              label: 'Level two 2-1',
            },
            {
              id: 6,
              label: 'Level two 2-2',
            },
          ],
        },
        {
          id: 3,
          label: 'Level one 3',
          children: [
            {
              id: 7,
              label: 'Level two 3-1',
            },
            {
              id: 8,
              label: 'Level two 3-2',
            },
          ],
        },
      ],
    },
    {
      type: 'select',
      label: '模型名称',
      value: '',
      options: [
        { label: '楼层', value: 1 },
        { label: '温度', value: 2 },
      ],
    },
    {
      type: 'input',
      label: '实例名称',
      value: '',
      //  width: '380px',
    },
    {
      type: 'select',
      label: '关系类型',
      value: '',
      options: [
        {
          label: '空间',
          value: '1',
        }
      ],
    },
    {
      type: 'select',
      label: '展示层级',
      value: '',
      options: [
        {
          label: '3',
          value: 3,
        },
      ],
    },
  ],
  gutter: 0,
}

export const buttonConfig = {
  buttonOk: {
    name: '编辑实例关系',
    className: 'roundBlueIcon',
    round: true,
    icon: markRaw(Edit), // 引用组件
  },
}

export const drawerData = reactive({
  isShowDrawer: false,
  title: '新建关联',
  isDisabled: false,
  closeDrawer() {
    drawerData.isShowDrawer = false
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