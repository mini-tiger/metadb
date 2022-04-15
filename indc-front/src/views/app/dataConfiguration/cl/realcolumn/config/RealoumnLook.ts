// 给个table1的配置项
export const measuringConfig = {
  columns: [
    {
      label: '测点ID',
      prop: 'id',
      minWidth: 100,
      // align: 'center',
      // isSlot: true,
    },
    {
      label: '测点名称',
      prop: 'MeasuringName',
      // minWidth: 100,
      // align: 'center',
      // fixed: 'right',
      // sortable: true,
      // isSlot: true,
    },
    {
      label: '单位',
      prop: 'company',
      minWidth: 140,
      // align: 'left',
    },
    {
      label: '精度',
      prop: 'accuracy',
      // minWidth: 200,
      // align: 'center',
    },
    {
      label: '测点值',
      prop: 'pointValue',
      minWidth: 180,
      // align: 'center',
    },
    { label: '工作区间', prop: 'WorkingRange' },
    {
      prop: 'dataSources',
      label: '数据来源',
      width: 160,
      // align: 'center',
      // isAction: true,
      // isSlot: true,
    },
  ],
  tableData: [
    {
      id: 1,
      MeasuringName: 'AB线电压',
      company: 'V',
      accuracy: '0.1',
      pointValue: '40.11',
      WorkingRange: '(0) - (460)',
      dataSources: 'B_N3_1_97_0_186',
    },
    {
      id: 2,
      MeasuringName: 'AB线电压',
      company: 'V',
      accuracy: '0.1',
      pointValue: '40.11',
      WorkingRange: '(0) - (460)',
      dataSources: 'B_N3_1_97_0_186',
    },
    {
      id: 3,
      MeasuringName: 'AB线电压',
      company: 'V',
      accuracy: '0.1',
      pointValue: '40.11',
      WorkingRange: '(0) - (460)',
      dataSources: 'B_N3_1_97_0_186',
    },
    {
      id: 4,
      MeasuringName: 'AB线电压',
      company: 'V',
      accuracy: '0.1',
      pointValue: '40.11',
      WorkingRange: '(0) - (460)',
      dataSources: 'B_N3_1_97_0_186',
    },
  ],
  // tableType: '', // 选择框
  // isSelection: true,
  // isIndex: true,
  isPagination: true, // 是否展示分页
  pagination: {
    total: 400,
    currentPage: 1,
    pageSize: 10,
  }, //分页配置信息
  auxiliary: false, //是否展示分页
  // tableBasics: {
  //   height: 300, //高度
  //   maxHeight: '1000px', //最大高度
  //   stripe: false, // 是否为斑马纹
  //   border: true, // 是否带有纵向边框
  //   size: 'default', // Table 的尺寸 分为large / default /small
  //   fit: true, // 列的宽度是否自撑开
  //   showHeader: true, // 是否显示表头
  //   highlightCurrentRow: true, // 是否要高亮当前行
  //   emptyText: '暂无数据', // 空数据时显示的文本内容
  //   defaultExpandAll: true, // 是否默认展开所有行，当 Table 包含展开行存在或者为树形表格时有效
  //   tooltipEffect: 'dark', // tooltip effect 属性
  //   lazy: true, //是否懒加载子节点数据
  //   headerRowStyle: {}, // 表头样式
  // },
}

export const clRelationConfig = {
  columns: [
    {
      label: '序号',
      prop: 'index',
      minWidth: 50,
    },
    {
      label: '关系类型',
      prop: 'Relationship',
    },
    {
      label: '源模型',
      prop: 'SourceModel',
    },
    {
      label: '源实例',
      prop: 'SourceInstance',
    },
    {
      label: '目标模型',
      prop: 'TargetModel',
    },
    {
      label: '目标实例',
      prop: 'TargetInstance',
    },
    {
      label: '操作',
      prop: 'modify',
      isSlot: true,
    },
  ],

  tableData: [
    {
      index: 1,
      Relationship: '空间关系',
      SourceModel: '房间',
      SourceInstance: '101模块间',
      TargetModel: 'UPS设备模型',
      TargetInstance: 'UPS_1',
    },
    {
      index: 2,
      Relationship: '空间关系',
      SourceModel: '房间',
      SourceInstance: '101模块间',
      TargetModel: 'UPS设备模型',
      TargetInstance: 'UPS_1',
    },
    {
      index: 3,
      Relationship: '空间关系',
      SourceModel: '房间',
      SourceInstance: '101模块间',
      TargetModel: 'UPS设备模型',
      TargetInstance: 'UPS_1',
    },
    {
      index: 4,
      Relationship: '空间关系',
      SourceModel: '房间',
      SourceInstance: '101模块间',
      TargetModel: 'UPS设备模型',
      TargetInstance: 'UPS_1',
    },
  ],
}

export const changeRecordConfig = {
  columns: [
    {
      label: '操作账号',
      prop: 'OperationAccount',
    },
    {
      label: '操作时间',
      prop: 'OperationTime',
    },
    {
      label: '操作描述',
      prop: 'OperationDescribe',
    },
  ],
  tableData: [
    {
      OperationAccount: 'admin',
      OperationTime: '2022-3-22',
      OperationDescribe: '修改C内容',
    },
  ],
  pagination: {
    total: 400,
    currentPage: 1,
    pageSize: 10,
  }, //分页配置信息
}

// 新建关联弹出框的搜索表单
export const formConfig = {
  searchForm: [
    {
      type: 'Select',
      label: '关联列表',
      prop: 'AssociationList',
      width: '330px',
      options: [
        { label: '楼层', value: 1 },
        { label: '温度', value: 2 },
      ],
    },
    {
      type: 'SelectAndInput',
      prop: 'exampleName',
      label: '筛选条件',
      width: '220px',
      value: 'exampleNameInput',
      options: [
        { label: '实例名称', value: 1 },
        { label: '楼层', value: 2 },
        { label: '空调', value: 3 },
      ],
    },
  ],
  searchData: {
    AssociationList: '',
    exampleName: '',
    exampleNameInput: '',
  },
  rules: {
    AssociationList: [
      {
        required: true,
        message: '请输入关联列表',
        trigger: 'blur',
      },
    ],
  },
  span: 12,
  // formArray: false,
  formSlot: true,
}

// 新建关联弹出框Table配置项
export const associationConfig = {
  columns: [
    {
      label: '关联类型',
      prop: 'relationModel',
    },
    {
      label: '源至目标约束',
      prop: 'sttc',
    },
    {
      label: '操作',
      prop: 'modify',
      isSlot: true,
    },
  ],
  tableData: [
    {
      relationModel: 'UPS-1',
      sttc: 'UPS模型',
      relation: 1,
    },
    {
      relationModel: 'UPS-1',
      sttc: 'UPS模型',
      relation: 1,
    },
    {
      relationModel: 'UPS-1',
      sttc: 'UPS模型',
      relation: 0,
    },
    {
      relationModel: 'UPS-1',
      sttc: 'UPS模型',
      relation: 0,
    },
  ],
  tableBasics: {
    headerRowStyle: {
      background: '#F8F8FF',
      'border-radius': '3px 3px 0px 0px',
    },
  },
}

export const closeButtonConfig = {
  name: '关闭',
  round: true,
  className: 'roundBackBlue',
}

// 属性字段
export const objectattr = [
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'default',
    bk_property_group_name: 'Default',
    bk_property_id: 'bk_inst_name',
    bk_property_index: 0,
    bk_property_name: '实例名',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-22 17:20:16',
    creator: 'user',
    isrequired: true,
  },
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'default',
    bk_property_group_name: 'Default',
    bk_property_id: 'dc_address',
    bk_property_index: 9,
    bk_property_name: '数据中心地址',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-22 17:37:05',
    creator: 'admin',
    description: '',
    isrequired: false,
    editable: true,
  },
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'default',
    bk_property_group_name: 'Default',
    bk_property_id: 'dc_code',
    bk_property_index: 7,
    bk_property_name: '数据中心编码',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-22 17:50:50',
    creator: 'admin',
    description: '',
    editable: true,
    isrequired: true,
  },
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'default',
    bk_property_group_name: 'Default',
    bk_property_id: 'dc_name',
    bk_property_index: 8,
    bk_property_name: '数据中心名称',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-22 17:55:17',
    creator: 'admin',
    description: '',
    editable: true,
    isrequired: true,
  },
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'default',
    bk_property_group_name: 'Default',
    bk_property_id: 'dc_area',
    bk_property_index: 10,
    bk_property_name: '数据中心面积',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-22 17:58:37',
    creator: 'admin',
    description: '',
    editable: true,
    isrequired: true,
  },
  {
    bk_biz_id: 0,
    bk_isapi: false,
    bk_issystem: false,
    bk_obj_id: 'dcCenter',
    bk_property_group: 'baa8df96-b1c2-4d17-97bd-a4d03d993189',
    bk_property_group_name: '数据中心基础信息',
    bk_property_id: 'sadfsd',
    bk_property_index: 0,
    bk_property_name: '测试字段',
    bk_property_type: 'singlechar',
    bk_supplier_account: '0',
    create_time: '2022-03-23 09:59:30',
    creator: 'admin',
    description: '',
    editable: true,
    isrequired: false,
  },
]

export const dcCenter = [
  {
    bk_biz_id: 0,
    bk_group_id: 'default',
    bk_group_index: -1,
    bk_group_name: 'Default',
    bk_isdefault: true,
    bk_obj_id: 'dcCenter',
    bk_supplier_account: '0',
    id: 18,
    is_collapse: false,
    ispre: false,
  },
  {
    bk_biz_id: 0,
    bk_group_id: 'baa8df96-b1c2-4d17-97bd-a4d03d993189',
    bk_group_index: 0,
    bk_group_name: '数据中心基础信息',
    bk_isdefault: false,
    bk_obj_id: 'dcCenter',
    bk_supplier_account: '0',
    id: 21,
    is_collapse: false,
    ispre: false,
  },
]

export const dcCenterData = [
  {
    bk_inst_id: 5,
    bk_inst_name: '测试实例',
    bk_obj_id: 'dcCenter',
    bk_supplier_account: '0',
    create_time: '2022-03-23T14:56:14.949+08:00',
    dc_address: '',
    dc_area: '2000平方米',
    dc_code: 'dc001',
    dc_name: '测试数据中心',
    last_time: '2022-03-23T14:56:14.949+08:00',
    sadfsd: '',
  },
]
