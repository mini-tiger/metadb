import { markRaw, reactive } from 'vue'
import { Plus } from '@element-plus/icons-vue'

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
      type: 'input',
      label: '模型名称',
      value: '',
    },
    {
      type: 'input',
      label: '标签信息',
      value: '',
      //  width: '380px',
    },
    {
      type: 'select',
      label: '目前状态',
      value: '',
      options: [
        {
          label: '空间',
          value: '1',
        },
      ],
    },
    {
      type: 'input',
      label: '创建人',
      value: '',
    },
    {
      type: 'datepicker',
      label: '创建时间',
      value: '',
    },
  ],
  gutter: 0,
}

export const newModelConfig = {
  buttonOk: {
    name: '新建模型',
    className: 'roundBlueIcon',
    round: true,
    icon: markRaw(Plus), // 引用组件
  },
}

export const clModelTable = {
  columns: [
    {
      label: '图标',
      prop: 'objIcon',
      // align: 'center',
      isSlot: true,
    },
    {
      label: '模型名称',
      prop: 'objName',
      // align: 'center',
      // isSlot: true,
    },
    {
      label: '模型分类',
      prop: 'classificationName',
    },

    {
      label: '实例数',
      prop: 'insCount',
      isSlot: true,
    },
    {
      label: '创建人',
      prop: 'creator',
    },
    {
      label: '创建时间',
      prop: 'createTime',
    },
    {
      label: '状态',
      prop: 'ispaused',
      isSlot: true,
    },
    {
      label: '操作',
      prop: 'modify',
      width: '200px',
      isSlot: true,
    },
  ],
  tableData: [],
  tableBasics: {},
  isIndex: true, // 索引列
  pagination: {
    total: 0,
    currentPage: 1,
    pageSize: 10,
  },
  isPagination: true,
}

