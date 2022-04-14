<template>
  <div class="testStyle">
    <!-- <public-calendar></public-calendar> -->
    <public-table
      :data="state"
      class="list-group-item bg-gray-300 m-1 p-3 rounded-md text-center"
      @operationClick="methods.operationClick"
      @handleSelectionChange="methods.handleSelectionChange"
      @handleSizeChange="methods.handleSizeChange"
      @handleCurrentChange="methods.handleCurrentChange"
      @selectTableData="methods.selectTableData"
      @selectAllTableData="methods.selectAllTableData"
      @rowClick="methods.rowClick"
      @rowDblclick="methods.rowDblclick"
      @cellClick="methods.cellClick"
      @handleTableSort="methods.handleTableSort"
      @AddListRow="methods.AddListRow"
    >
      <template #name="{ scope }">
        {{ scope.row.id }}
      </template>

      <template #department="{ scope }">
        <el-switch v-model="scope.row.department" />
        {{ scope.row.prop }}
      </template>

      <template #modify><el-button>点击</el-button> </template>
    </public-table>
    <public-form :data="form">
      <template v-slot="data">
        <el-button @click="onSubmit(data.data.formRef, data.data.formData)"
          >点击</el-button
        >
      </template>
    </public-form>
    <public-button :data="buttonData.buttonOk" @onButton="onButton"
      >11</public-button
    >
    <!-- <public-button :data="buttonData.buttonOk"></public-button> -->

    <div class="tabs">
      <el-tabs v-model="data.activeName" @tab-click="data.handleClick">
        <el-tab-pane label="用户管理" name="first"> 用户管理 </el-tab-pane>
        <el-tab-pane label="配置管理" name="second">配置管理</el-tab-pane>
        <el-tab-pane label="角色管理" name="third">角色管理</el-tab-pane>
        <el-tab-pane label="定时任务补偿" name="fourth"
          >定时任务补偿</el-tab-pane
        >
      </el-tabs>
      <div class="button-right">
        <publicButton :data="buttonData.buttonOk"></publicButton>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, markRaw } from 'vue'
import publicTable from '@/components/Table/index.vue'
// import publicForm from '@/components/Form/index.vue'
import publicButton from '@/components/Button/index.vue'
import { Plus } from '@element-plus/icons-vue'

export default defineComponent({
  components: {
    // publicCalendar,
    publicTable,
    // publicForm,
    publicButton,
  },
  setup() {
    const state = reactive({
      table: {
        columns: [
          // { type: 'selection', minWidth: 40 },
          {
            label: '出差人',
            prop: 'name',
            minWidth: 100,
            // align: 'center',
            // isSlot: true,
          },
          {
            label: '部门名称',
            prop: 'department',
            minWidth: 100,
            // align: 'center',
            // fixed: 'right',
            sortable: true,
            isSlot: true,
          },
          {
            label: '项目名称',
            prop: 'project',
            minWidth: 140,
            // align: 'left',
          },
          {
            label: '目的地',
            prop: 'destination',
            // minWidth: 200,
            // align: 'center',
          },
          {
            label: '出发时间',
            prop: 'startTime',
            minWidth: 180,
            // align: 'center',
          },
          { label: '出差天数', prop: 'day', minWidth: 100, align: 'center' },
          {
            label: '出行工具',
            prop: 'TravelTool',
            minWidth: 100,
            // align: 'center',
          },
          {
            prop: 'modify',
            label: '操作',
            width: 90,
            align: 'center',
            // isAction: true,
            isSlot: true,
          },
        ],
        tableData: [
          {
            id: 1,
            name: '张三',
            department: '123',
            project: '智航',
            destination: '北京',
            startTime: '2021-02-12',
            day: '3',
            TravelTool: '火车',
            align: 'left',
          },
          {
            id: 2,
            name: '张三',
            department: '123',
            project: '智航',
            destination: '北京',
            startTime: '2021-02-12',
            day: '3',
            TravelTool: '火车',
          },
          {
            id: 3,
            name: '张三',
            department: '123',
            project: '智航',
            destination: '北京',
            startTime: '2021-02-12',
            day: '3',
            TravelTool: '火车',
          },
          {
            id: 4,
            name: '张三',
            department: '123',
            project: '智航',
            destination: '北京',
            startTime: '2021-02-12',
            day: '3',
            TravelTool: '火车',
            align: 'left',
          },
        ],
        tableType: '', // 选择框
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
        isPagination: true,
        // isSelection: true,
        // isIndex: true,
        pagination: {
          total: 400,
          currentPage: 1,
          // pageSize: 10,
        },
        tableWidth: '100%', // 表格的宽度
        auxiliary: true,
      },
    })
    const sexes = [
      { label: '男', value: 'M' },
      { label: '女', value: 'F' },
    ]
    let interests = [
      { label: '羽毛球', value: 'interst' },
      { label: '篮球', value: 'basketball' },
    ]
    // form的数据
    const form = reactive({
      searchForm: [
        {
          type: 'Input',
          label: '姓名',
          prop: 'name',
          span: 16,
        },
        { type: 'Input', label: '年龄', prop: 'age' },
        {
          type: 'Select',
          prop: 'sex',
          label: '下拉',
          options: sexes,
          // change: this.changea,
        },
        {
          type: 'Radio',
          prop: 'interst',
          label: '单选',
          radios: interests,
        },
        {
          type: 'Date',
          prop: 'interst5',
          label: '日期-datetime',
          pickerType: 'datetime',
        },
        {
          type: 'Time',
          prop: 'interst8',
          label: '时间',
        },
        {
          type: 'Switch',
          prop: 'interst6',
          label: '开关Switch',
        },
        {
          type: 'Search',
          prop: 'interst7',
          label: '搜索下拉',
          isMultiple: false,
          option: [
            { label: '张三', value: '1' },
            { label: '李四', value: '2' },
          ],
        },
        {
          type: 'Textarea',
          prop: 'interst10',
          label: '文本框',
        },
        {
          type: 'Checkbox',
          prop: 'Checkbox',
          label: '多选',
          option: [
            { label: '打球', value: '1' },
            { label: '打游戏', value: '2' },
          ],
        },
      ],
      searchData: {
        name: '',
        age: '',
        sex: '',
        interst: '',
        interst5: '',
        interst6: false,
        interst7: [],
        interst8: '',
        interst10: '',
      },
      labelWidth: '120px',
      isSubmit: true,
      formSlot: true, // 是否开始插槽
      // 验证规则
      rules: {
        name: [
          {
            required: true,
            message: '请输入名字',
            trigger: 'blur',
          },
        ],
      },
    })
    // 表格子组件穿过来的函数
    const methods = {
      // table start
      // 操作列返回数据
      operationClick() {
        // console.log(index, row, event)
      },
      // 选择框返回的数据
      handleSelectionChange() {
        // console.log(val)
      },
      // 页数
      handleSizeChange() {
        // console.log(val, 'val')
      },
      // 在第几页
      handleCurrentChange() {
        // console.log(val, 'val')
      },
      selectTableData() {
        // console.log(select, row, 'selectTableData')
      },
      selectAllTableData() {
        // console.log(selection, 'selectAllTableData')
      },
      rowClick() {
        // console.log(row, column, cell, 'rowClick')
      },
      rowDblclick() {
        // console.log(row, column, event, 'rowDblclick')
      },
      cellClick() {
        // console.log(row, column, cell, event, 'cellClick')
      },
      handleTableSort() {
        // console.log(column, prop, order, 'handleTableSort')
      },
      // 顶部操作列
      AddListRow(name) {
        console.log(name, 'name')
      },
      // table end
    }
    //
    const onSubmit = (ref, data) => {
      console.log(ref, data)
    }
    const nameList = [
      { label: '张三', value: '1' },
      { label: '张三', value: '2' },
      { label: '张三', value: '3' },
    ]
    // button组件
    const buttonData = reactive({
      buttonOk: {
        name: '点击',
        className: 'roundBlueIcon', // 1.roundBackWhite  2.roundBackBlue 3.roundBackCal 4.roundIcon 5.roundWhiteIcon 6.roundBlueIcon
        round: true,
        // loading: true,
        icon: markRaw(Plus), // 引用组件
      },
    })

    // button的点击事件
    const onButton = (e) => {
      console.log(e, '---')
    }

    const data = reactive({
      activeName: '',
      handleClick(tab, event) {
        console.log(tab, event)
      },
    })
    return {
      state,
      methods,
      form,
      onSubmit,
      buttonData,
      nameList,
      onButton,
      data,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '../assets/style/overallStyle.scss';
</style>
