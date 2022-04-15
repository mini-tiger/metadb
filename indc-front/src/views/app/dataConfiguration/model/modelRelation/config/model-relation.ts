import { markRaw } from 'vue'
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
      type: 'select',
      label: '关系类型',
      value: '',
      options: [
        {
          label: '空调',
          value: 1,
        }
      ],
    },
    {
      type: 'select',
      label: '层级展示',
      options: [
        {
          label: '维度',
          value: 1,
        }
      ],
      value: '',
    },
  ],
  gutter: 0,
}

export const editTopologyConfig = {
  buttonOk: {
    name: '编辑模型关系',
    className: 'roundBlueIcon',
    round: true,
    icon: markRaw(Edit), // 引用组件
  }
}