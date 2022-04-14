import { reactive } from 'vue'

export const commFiled = {
  searchForm: [
    {
      type: 'Input',
      label: '唯一标识',
      prop: 'id',
      width: '300px',
      disabled: null,
    },
    {
      type: 'Input',
      label: '标题',
      prop: 'title',
      width: '300px',
    },
    {
      label: '字段类型',
      type: 'Select',
      width: '300px',
      prop: 'filedKind',
      disabled: null,
      options: [
        {
          label: '文本域',
          value: '文本域',
        },
        {
          label: '数字',
          value: '数字',
        },
        {
          label: '下拉列表',
          value: '下拉列表',
        },
        {
          label: '日期',
          value: '日期',
        },
        {
          label: '起止时间',
          value: '起止时间',
        },
        {
          label: '时间',
          value: '时间',
        },
        {
          label: '关联已有数据',
          value: '关联已有数据',
        },
        {
          label: '公式计算',
          value: '公式计算',
        },
      ],
    },
    {
      type: 'Input',
      width: '300px',
      label: '提示',
      prop: 'titTop',
    },
    {
      label: '默认内容',
      width: '300px',
      type: 'Input',
      prop: 'defaultContent',
      formatType: null,
    },
    {
      label: '字段选择',
      type: 'Checkbox',
      prop: 'filedCheck',
      option: [
        {
          label: '必填',
          value: 1,
        },
        {
          label: '可编辑',
          value: 2,
        },
      ],
    },
  ],
  searchData: {
    id: '',
    title: '',
    filedKind: '',
    titTop: '',
    defaultContent: '',
    filedCheck: ['可编辑'],
  },
  rules: {
    id: [
      {
        required: true,
        message: '请输入',
        trigger: 'blur',
      },
    ],
    title: [
      {
        required: true,
        message: '请输入',
        trigger: 'blur',
      },
    ],
    filedKind: [
      {
        required: true,
        message: '请输入',
        trigger: 'blur',
      },
    ],
  },
  formArray: false,
}

// 确定配置
export const determineConfig = {
  name: '确定',
  className: 'roundBackBlue',
  round: true,
}

// 取消配置
export const cancelConfig = {
  name: '取消',
  className: 'roundBackCal',
  round: true,
}

// formEdit
export const defineDataConfig = {
  // 顶部的数据
  id: '', //唯一标识
  title: '', //标题
  filedKind: '', //字段类型
  titTop: '', //提示
  defaultContent: '', //默认内容
  filedCheck: ['可编辑'],
  // 文本域
  regularCheck: '', //正则
  limitNumberWords: false, //是否字数限制
  limitNumber: '',
  limitNumberWordsNum: 1, //字数限制输入框值
  duplicateNot: false, //是否允许重复
  //数字
  company: '', //单位
  DecimalAll: false, //是否允许小数
  LimitDigit: '', //小数限制位数
  DecimalLimitRange: false, //是否限制数字范围
  regularCheckMin: '', //最小值限制
  regularCheckBig: '', //最大值限制
  // 下拉列表
  domains: [{ label: '', value: Date.now() }], //下来列表数据
  // 日期
  selectFormat: 'YYYY/MM/DD', //日期格式
  // 时间
  Time: 'hh:mm', //时间格式
  // 关联已有数据
  aeData: '', //关联模型
  modelFiled: '', //关联字段
  // 公式
  AssociateModel: '', // 模型
  AssociatedInstance: '', // 实例
  AssociatedFiled: '', // 字段
  AssociatedSum: '', //求和

  //标签
  tagName: '',
}

export const filedKindConfig = [
  {
    label: '文本域',
    value: '文本域',
  },
  {
    label: '数字',
    value: '数字',
  },
  {
    label: '下拉列表',
    value: '下拉列表',
  },
  {
    label: '日期',
    value: '日期',
  },
  {
    label: '起止时间',
    value: '起止时间',
  },
  {
    label: '时间',
    value: '时间',
  },
  {
    label: '关联已有数据',
    value: '关联已有数据',
  },
  {
    label: '公式计算',
    value: '公式计算',
  },
]

export const rules = {
  id: [
    {
      required: true,
      message: '请输入',
      trigger: 'blur',
    },
    {
      validator: function (rule, value, callback) {
        if (/[a-zA-z]$/.test(value) == false) {
          callback(new Error('请输入英文'))
        } else {
          //校验通过
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
  title: [
    {
      required: true,
      message: '请输入',
      trigger: 'blur',
    },
  ],
  filedKind: [
    {
      required: true,
      message: '请输入',
      trigger: 'blur',
    },
  ],
  tagName: [
    {
      required: true,
      message: '请输入',
      trigger: 'blur',
    },
  ],
}

export const state = reactive({
  modelForm: {
    objName: '', //模型名称
    ispaused: false, //状态（true-停用 false-启用）
    description: '', //描述
    classificationId: '', //模型分类ID
    classificationName: '', //模型分类名称
    objIcon: '', //模型图标
  },
  rules: {
    objName: [
      {
        required: true,
        message: '请输入模型名称',
        trigger: 'blur',
      },
    ],
    classificationName: [
      {
        required: true,
        message: '请输入模型分类',
        trigger: 'blur',
      },
    ],
    objIcon: [
      {
        required: true,
        message: '请选择图标',
        trigger: 'blur',
      },
    ],
  },
  isShow: false,
  Administration: [],
  svgFlag: false,
  svgSum: 0,
  svgList: [
    'modelBuilding',
    'modelCabinet',
    'modelCenter',
    'modelColumn',
    'modelEquipment',
    'modelFloor',
    'modelRoom',
    'modelSig',
  ],
  defaultSvg: '',
})
