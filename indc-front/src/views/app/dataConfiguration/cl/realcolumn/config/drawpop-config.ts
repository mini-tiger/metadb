export const table = {
  columns: [
    // { type: 'selection', minWidth: 40 },
    {
      label: 'ID',
      prop: 'id',
      minWidth: 100,
      // align: 'center',
      // isSlot: true,
    },
    {
      label: '实列名称',
      prop: 'exampleName',
      minWidth: 100,
      // align: 'center',
      // fixed: 'right',
      // sortable: true,
      // isSlot: true,
    },
    {
      label: '模型名称',
      prop: 'modeName',
      minWidth: 140,
      // align: 'left',
    },
    {
      label: '模型分类',
      prop: 'modeClassification',
      // minWidth: 200,
      // align: 'center',
    },
    {
      label: '创建人',
      prop: 'createPeople',
      minWidth: 180,
      // align: 'center',
    },
    { label: '创建时间', prop: 'createTime', minWidth: 100 },
    {
      prop: 'modify',
      label: '操作',
      width: 160,
      // align: 'center',
      // isAction: true,
      isSlot: true,
    },
  ],
  tableData: [
    {
      id: 1,
      exampleName: 'ups-1',
      modeName: 'ups模型',
      modeClassification: '设备',
      createPeople: '刘工',
      createTime: '2021-02-12',
      // align: 'left',
    },
    {
      id: 2,
      exampleName: 'ups-1',
      modeName: 'ups模型',
      modeClassification: '设备',
      createPeople: '刘工',
      createTime: '2021-02-12',
      // align: 'left',
    },
    {
      id: 3,
      exampleName: 'ups-1',
      modeName: 'ups模型',
      modeClassification: '设备',
      createPeople: '刘工',
      createTime: '2021-02-12',
      // align: 'left',
    },
    {
      id: 4,
      exampleName: 'ups-1',
      modeName: 'ups模型',
      modeClassification: '设备',
      createPeople: '刘工',
      createTime: '2021-02-12',
      // align: 'left',
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
  tableBasics: {
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
  },
}
