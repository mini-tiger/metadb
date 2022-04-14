<template>
  <div>
    <el-row>
      <el-col style="text-align: right">
        <div class="headerMid">
          <span class="icon" @click="AddListRow('addColumn')">
            <svg-icon icon-class="canvas-add"></svg-icon> <span>加列</span>
          </span>
          <span class="icon" @click="AddListRow('addLine')">
            <svg-icon icon-class="canvas-add"></svg-icon> <span>加行</span>
          </span>
        </div>
      </el-col>
      <el-table
        :data="TableList.table.tableData"
        style="width: 100%"
        :header-cell-style="handleTheadStyle"
        height="320px"
        row-key="id"
        :header-cell-class-name="headerCell"
      >
        <template v-for="(item, index) in TableList.table.columns" :key="index">
          <el-table-column property="operation" align="right" width="138px">
            <template #header="scope">
              <span class="edit-icon icon" @click="headerClick(item.id, scope)">
                <svg-icon svg-icon icon-class="edit"></svg-icon>
              </span>
              <span class="del-icon icon" @click="delSection(item.id)"
                ><svg-icon icon-class="del"></svg-icon
              ></span>
            </template>
            <el-table-column
              :property="item.id"
              align="center"
              :label="item.title"
              width="350px"
              show-overflow-tooltip
            >
              <template #default="scope">
                <div v-if="item.filedKind === '文本域'">
                  <el-input
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    :placeholder="'请输入'"
                    style="width: 90%"
                  ></el-input>
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
                <div v-if="item.filedKind === '数字'">
                  <el-input
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    :placeholder="'请输入'"
                    style="width: 90%"
                  ></el-input>
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>

                <div
                  v-if="['下拉列表', '关联已有关系'].includes(item.filedKind)"
                >
                  <el-select
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    :placeholder="'请输入'"
                    style="width: 90%"
                  >
                    <el-option
                      v-for="op in item.domains"
                      :key="op.value"
                      :label="op.label"
                      :value="op.value"
                    >
                    </el-option>
                  </el-select>
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
                <div v-if="item.filedKind === '日期'">
                  <el-date-picker
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    :value-format="item.selectFormat"
                    :format="item.selectFormat"
                    style="width: 90%"
                  />
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
                <div v-if="item.filedKind === '起止时间'">
                  <el-date-picker
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    type="daterange"
                    range-separator="至"
                    start-placeholder="开始时间"
                    end-placeholder="结束时间"
                    :value-format="item.selectFormat"
                    :format="item.selectFormat"
                    style="width: 90%"
                  />
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
                <div v-if="item.filedKind === '时间'">
                  <el-time-picker
                    v-model="TableList.table.tableData[scope.$index][item.id]"
                    placeholder="请输入"
                    arrow-control
                    :value-format="item.Time"
                    :format="item.Time"
                    style="width: 90%"
                  />
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
                <div v-if="item.filedKind === '公式计算'">
                  <span>{{ item.defaultContent }}</span>
                  <span v-if="item.titTop">
                    <el-tooltip
                      class="box-item"
                      effect="light"
                      :content="item.titTop"
                      placement="top-start"
                    >
                      <span><svg-icon icon-class="tips"></svg-icon> </span>
                    </el-tooltip>
                  </span>
                </div>
              </template>
            </el-table-column>
          </el-table-column>
        </template>

        <el-table-column
          width="40px"
          v-if="TableList.table.columns.length !== 0"
        >
          <el-table-column
            prop="lastColumn"
            width="40px"
            class-name="delRowButton"
          >
            <template #default="scope">
              <span class="del-icon icon" @click="delRow(scope)"
                ><svg-icon icon-class="del"></svg-icon
              ></span>
            </template>
          </el-table-column>
        </el-table-column>
      </el-table>
    </el-row>

    <useEdit
      :IsDrawer="state.drawer"
      :isEmit="state.isEmit"
      :pointFlag="state.pointFlag"
      @IsDrawer="IsDrawer"
      @defineData="defineData"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref, onUnmounted } from 'vue'
import { useRoute } from 'vue-router'
import Sortable from 'sortablejs'

import { defineDataConfig } from '../hook/config/edit-config'
import { point } from '../config/point-config'
import useEdit from '../hook/use-newFiled.vue'

export default defineComponent({
  components: {
    useEdit,
  },
  setup() {
    let state = reactive({
      drawer: false,
      emitData: [],
      isEmit: false,
      cleanForm: '',
      cleanSelect: null,
      pointFlag: '',
      type: '',
      ModelData: {},
    })
    const TableList = ref({ table: point })
    const route = useRoute()
    let data = null

    if (route.params && !localStorage.getItem('rowData')) {
      data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    onMounted(() => {
      console.log(route, '000')
      const wrapperTr = document.querySelector('.el-table__header-wrapper tr')
      const ops = {
        animation: 1000,
        sort: true,
        filter: '.delRowButton',
        onEnd: endChange,
      }
      Sortable.create(wrapperTr, ops)
      // console.log(route.params && JSON.parse(route.params))
      console.log(localStorage.getItem('rowData'))
    })
    const endChange = (e) => {
      const oldItem = TableList.value.table.columns[e.oldIndex]
      TableList.value.table.columns.splice(e.oldIndex, 1)
      TableList.value.table.columns.splice(e.newIndex, 0, oldItem)
      console.log(TableList.value, '----------<')
    }
    const headerCell = ({ row, rowIndex, column, columnIndex }) => {
      console.log(column, columnIndex, row, rowIndex)
      // const header = document.querySelectorAll('.el-table__header-wrapper tr th')[columnIndex]
      // console.log(header, '------>')
    }
    const AddListRow = (type) => {
      console.log(type)
      if (type === 'addColumn') addColumn()
      if (type === 'addLine') addLine()
    }
    // 添加列
    const addColumn = () => {
      console.log('我是添加列')
      state.isEmit = false
      state.pointFlag = ''
      state.drawer = true
    }
    //添加行
    const addLine = () => {
      console.log('我是行')
      let obj = {}
      if (TableList.value.table.columns.length === 0) return
      for (let key in TableList.value.table.tableData[0]) {
        obj[key] = ''
      }
      console.log(obj, '----obj---')
      TableList.value.table.tableData.push(obj)
      console.log(TableList.value, 'table===>')
    }
    // 关闭按钮
    const IsDrawer = () => {
      state.drawer = false
    }
    const storeData = (data) => {
      let arr = {}
      arr = {
        AssociateModel: data.AssociateModel,
        AssociatedFiled: data.AssociatedFiled,
        AssociatedInstance: data.AssociatedInstance,
        DecimalAll: data.DecimalAll,
        DecimalLimitRange: data.DecimalLimitRange,
        LimitDigit: data.LimitDigit,
        Time: data.Time,
        aeData: data.aeData,
        company: data.company,
        defaultContent: data.defaultContent,
        domains: data.domains,
        duplicateNot: data.duplicateNot,
        filedCheck: data.filedCheck,
        filedKind: data.filedKind,
        limitNumber: data.limitNumber,
        limitNumberWords: data.limitNumberWords,
        limitNumberWordsNum: data.limitNumberWordsNum,
        modelFiled: data.modelFiled,
        regularCheck: data.regularCheck,
        regularCheckBig: data.regularCheckBig,
        regularCheckMin: data.regularCheckMin,
        selectFormat: data.selectFormat,
        titTop: data.titTop,
        id: data.id,
        title: data.title,
      }
      return arr
    }
    // 确认按钮
    const defineData = (data) => {
      // console.log(data, 'data====>', TableList.value.table.columns)
      if (TableList.value.table.columns.length > 0) {
        const result = TableList.value.table.columns.find(
          (item) => item.id === data.id,
        )
        if (!result) {
          TableList.value.table.columns.push(storeData(data))
        } else {
          for (let i = 0; i < TableList.value.table.columns.length; i++) {
            if (data.id === TableList.value.table.columns[i].id) {
              assemblyData(TableList.value.table.columns, i, data)
            }
          }
        }
      } else {
        TableList.value.table.columns.push(storeData(data))
      }
      if (state.emitData.length === 0) {
        state.emitData.push(storeData(data))
      } else {
        const resultData = state.emitData.find((item) => item.id === data.id)

        if (!resultData) {
          state.emitData.push(storeData(data))
        } else {
          for (let j = 0; j < state.emitData.length; j++) {
            if (data.id === state.emitData[j].id) {
              assemblyData(state.emitData, j, data)
            }
          }
        }
      }
      let obj = {}
      if ([0, 1].includes(TableList.value.table.tableData.length)) {
        if (data.defaultContent) {
          obj[data.id] = data.defaultContent
        } else {
          obj[data.id] = ''
        }
        if (TableList.value.table.columns.length === 1) {
          TableList.value.table.tableData[0] = Object.assign(obj)
        } else {
          TableList.value.table.tableData[0] = Object.assign(
            TableList.value.table.tableData[0],
            obj,
          )
        }
      } else {
        console.log('多行')
        if (data.defaultContent) {
          obj[data.id] = data.id
        } else {
          obj[data.id] = ''
        }
        for (let k = 0; k < TableList.value.table.tableData.length; k++) {
          TableList.value.table.tableData[k] = Object.assign(
            TableList.value.table.tableData[k],
            obj,
          )
        }
      }
      console.log(TableList.value.table.tableData, '我是数据源----《》', obj)
      // // debugger
      // for (let k = 0; k < TableList.value.table.tableData.length; k++) {
      //   TableList.value.table.tableData[k] = Object.assign(
      //     TableList.value.table.tableData[k],
      //     // TableList.value.table.tableData[k],
      //     obj,
      //   )
      // }

      // console.log(state.emitData, TableList.value)
      state.drawer = false
    }
    // 格式化数据
    const assemblyData = (type, i, data) => {
      type[i].AssociateModel = data.AssociateModel
      type[i].AssociatedFiled = data.AssociatedFiled
      type[i].AssociatedInstance = data.AssociatedInstance
      type[i].DecimalAll = data.DecimalAll
      type[i].DecimalLimitRange = data.DecimalLimitRange
      type[i].LimitDigit = data.LimitDigit
      type[i].Time = data.Time
      type[i].aeData = data.aeData
      type[i].company = data.company
      type[i].domains = data.domains
      type[i].duplicateNot = data.duplicateNot
      type[i].filedCheck = data.filedCheck
      type[i].filedKind = data.filedKind
      type[i].limitNumber = data.limitNumber
      type[i].limitNumberWords = data.limitNumberWords
      type[i].limitNumberWordsNum = data.limitNumberWordsNum
      type[i].modelFiled = data.modelFiled
      type[i].regularCheck = data.regularCheck
      type[i].regularCheckBig = data.regularCheckBig
      type[i].regularCheckMin = data.regularCheckMin
      type[i].selectFormat = data.selectFormat
      type[i].titTop = data.titTop
      type[i].title = data.title
    }
    // 点击表头
    const headerClick = (index, data) => {
      console.log(data, 'data')
      state.emitData.map((item) => {
        if (item.id === index) {
          console.log(item, '我是id')
          for (let key in item) {
            defineDataConfig[key] = item[key]
          }
          state.isEmit = true
          state.drawer = true
          state.pointFlag = defineDataConfig.filedKind
        }
      })
    }
    // 点击编辑
    const editSection = (data) => {
      console.log('我是编辑', data)
    }
    // 点击删除
    const delSection = (data) => {
      console.log(data, 'data--删除')
    }
    // 初始化表头样式
    const handleTheadStyle = () => {
      // console.log('我是初始化---》')
      const style = {
        fontSize: '14px !important',
        fontWeight: '500 !important',
        color: '#302E6D !important',
        background: '#F8F8FF',
      }
      return style
    }
    //删除行
    const delRow = (data) => {
      console.log('data', data)
    }
    onUnmounted(() => {
      localStorage.removeItem('rowData')
    })
    return {
      state,
      TableList,
      headerCell,
      AddListRow,
      IsDrawer,
      defineData,
      headerClick,
      editSection,
      delSection,
      handleTheadStyle,
      delRow,
    }
  },
})
</script>

<style lang="scss" scoped>
.icon {
  margin-right: 11px;
  .svg-icon {
    background: #f8f8ff !important;
  }
}
.el-tooltip__trigger {
  margin-left: 10px;
  margin-top: 5px;
}
:deep(.delRowButton div) {
  width: 20px !important;
}
.headerMid {
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #595aec;
  .icon {
    cursor: pointer;
    margin-right: 20px;
    span {
      vertical-align: top;
    }
  }
}
</style>
