<template>
  <el-row :style="`width:${table.tableWidth}`">
    <el-col :span="24" class="btns-slot" v-if="table.isShowBtns">
      <slot name="btns"></slot>
      <div class="showCols">
        <ZHButton :options="btnOptions" @click="showColsList">列显示</ZHButton>
        <div class="showcols-list" v-if="isShowColsList">
          <el-checkbox
            v-model="checkAll"
            :indeterminate="isIndeterminate"
            @change="handleCheckAllChange"
            >全选</el-checkbox
          >
          <el-checkbox-group
            v-model="checkedItems"
            @change="handleCheckedItemssChange"
          >
            <el-checkbox 
              v-for="item in colsList" 
              :key="item.prop" 
              :label="item.label"
              >
              {{item.label}}
            </el-checkbox>
          </el-checkbox-group>
        </div>
      </div>
    </el-col>
    <el-col style="text-align: right" v-if="table.auxiliary">
      <el-button type="primary" @click="AddListRow('addColumn')"
        >加列</el-button
      >
      <el-button type="primary" @click="AddListRow('addLine')">加行</el-button>
      <el-button @click="AddListRow('ImportExcel')">导入</el-button>
      <el-button @click="AddListRow('exportExcel')">导出</el-button>
    </el-col>
    <el-col :span="24">
      <el-table
        class="tableBox"
        ref="multipleTableRef"
        :data="table.tableData"
        :height="table.tableBasics.height ? table.tableBasics.height : null"
        :max-height="
          table.tableBasics.maxHeight ? table.tableBasics.maxHeight : null
        "
        :stripe="table.tableBasics.stripe ? table.tableBasics.stripe : null"
        :border="table.tableBasics.border ? table.tableBasics.border : false"
        :size="table.tableBasics.size ? table.tableBasics.size : 'default'"
        :fit="table.tableBasics.fit ? table.tableBasics.fit : true"
        :show-header="
          table.tableBasics.showHeader ? table.tableBasics.showHeader : true
        "
        :highlight-current-row="
          table.tableBasics.highlightCurrentRow
            ? table.tableBasics.highlightCurrentRow
            : true
        "
        :empty-text="
          table.tableBasics.emptyText ? table.tableBasics.emptyText : null
        "
        :default-expand-all="
          table.tableBasics.defaultExpandAll
            ? table.tableBasics.defaultExpandAll
            : false
        "
        :tooltip-effect="
          table.tableBasics.tooltipEffect
            ? table.tableBasics.tooltipEffect
            : 'dark'
        "
        :lazy="table.tableBasics.lazy ? table.tableBasics.lazy : true"
        style="width: auto"
        :row-style="methods.rowStyle"
        :header-row-style="methods.headerRowStyle"
        :row-class-name="tableRowClassName"
        @selection-change="methods.handleSelectionChange"
        :select="methods.selectTableData"
        @select-all="methods.selectAllTableData"
        @row-click="methods.rowClick"
        @row-dblclick="methods.rowDblclick"
        @cell-click="methods.cellClick"
        @sort-change="methods.handleTableSort"
        :header-cell-style="methods.handleTheadStyle"
        v-if="table.columns.length !== 0"
        :tree-props="
          table.tableBasics.treeProps ? table.tableBasics.treeProps : {}
        "
        :row-key="table.tableBasics.rowKey ? table.tableBasics.rowKey : ''"
        :load="table.tableBasics.load ? table.tableBasics.load : null"
      >
        <el-table-column type="selection" width="55" v-if="table.isSelection" />
        <el-table-column
          type="index"
          width="60"
          v-if="table.isIndex"
          label="序号"
        />
        <template v-for="(item, index) in table.columns" :key="index">
          <el-table-column
            :label="item.label"
            :property="item.prop"
            :align="item.align ? item.align : 'left'"
            :header-align="item.align ? item.align : 'left'"
            :column-key="item.prop"
            :fixed="item.fixed ? item.fixed : null"
            :width="item.width || 'auto'"
            :min-width="item.minWidth"
            :sortable="item.sortable ? item.sortable : false"
            :sort-orders="['descending', 'ascending', null]"
            :type="item.type"
            :formatter="item.formatter"
            show-overflow-tooltip
          >
            <template #default="scope">
              <template v-if="item.isSlot">
                <slot
                  :name="item.prop"
                  :scope="{
                    row: scope.row,
                    column: scope.column,
                    $index: scope.$index,
                  }"
                ></slot>
              </template>
              <template v-else-if="item.formatter">
                <div v-html="item.formatter(scope.row[item.prop], scope)"></div>
              </template>
              <template v-else>
                <a v-if="item.isClick">
                  {{ scope.row[item.prop] }}
                </a>
                <span v-else>{{ scope.row[item.prop] }}</span>
              </template>
            </template>
          </el-table-column>
        </template>
      </el-table>
      <!--分页栏-->
      <el-config-provider :locale="locale">
        <div
          class="toolbar"
          style="padding: 10px; text-align: right"
          v-if="table.isPagination"
        >
          <el-pagination
            v-model:currentPage="table.pagination.currentPage"
            v-model:page-size="table.pagination.pageSize"
            background
            :page-sizes="[10, 50, 100, 200]"
            layout="total, sizes, prev, pager, next, jumper"
            :total="table.pagination.total"
            @size-change="methods.handleSizeChange"
            @current-change="methods.handleCurrentChange"
          />
        </div>
      </el-config-provider>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  ref,
  watch,
  // getCurrentInstance,
} from 'vue'
import type { ElTable } from 'element-plus'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import ZHButton from '@/components/Button/Button.vue'

export default defineComponent({
  components: {
    ZHButton
  },
  // props: ['data'],
  props: {
    data: { type: Object },
    // columns: {
    //   type: Array,
    //   default: (): any => {
    //     return []
    //   },
    // },
  },
  setup(props, content) {
    const multipleTableRef = ref<InstanceType<typeof ElTable>>()
    const multipleSelection = ref([])
    const pageSize = ref(null)
    const currentPage = ref(null)
    const table = reactive({
      isShowBtns: false,
      tableData: [],
      columns: [],
      isSelection: false, // 选择列
      isIndex: false, // 索引列
      pagination: {
        total: 0,
        currentPage: 1,
        pageSize: 10,
      },
      tableWidth: '100%',
      tableBasics: {
        height: '420px',
        maxHeight: '1000px',
        stripe: false, // 是否为斑马纹
        border: false, // 是否带有纵向边框
        size: 'default', // Table 的尺寸 分为large / default /small
        fit: true, // 列的宽度是否自撑开
        showHeader: true, // 是否显示表头
        highlightCurrentRow: true, // 是否要高亮当前行
        emptyText: '暂无数据', // 空数据时显示的文本内容
        defaultExpandAll: true, // 是否默认展开所有行，当 Table 包含展开行存在或者为树形表格时有效
        tooltipEffect: 'dark', // tooltip effect 属性
        lazy: true, //是否懒加载子节点数据
        rowKey: '',
        load: null,
        treeProps: null,
        headerRowStyle: {},
      },
      auxiliary: false,
      isPagination: true,
      isShowColumn: false,
    })

    // let { columns: columnsList } = toRefs(props) as any
    let colsList = ref([])
    onMounted(() => {
      assignment()
      checkedItems.value = table.columns.map(item=>item.label)
      colsList.value =  [...table.columns]
    })
    watch(props.data, () => {
      assignment()
    })
    const assignment = () => {
      if (props.data.table) {
        for (let key in props.data.table) {
          if (key) {
            table[key] = props.data.table[key]
          }
        }
      }
    }
    const methods = {
      operationClick(index, row, event) {
        // console.log(index, '我是索引', row, event)
        if (event.target.id) {
          // console.log('我是id')
          content.emit('operationClick', index, row, event.target.id)
        } else {
          console.log('请传入id')
        }
      },
      handleSelectionChange(val) {
        multipleSelection.value = val
        content.emit('handleSelectionChange', val)
      },
      // 改变页数
      handleSizeChange(val) {
        content.emit('handleSizeChange', val)
      },
      // 第几条
      handleCurrentChange(val) {
        content.emit('handleCurrentChange', val)
      },
      // 当用户手动勾选数据行的 Checkbox 时触发的事件
      selectTableData(select, row) {
        // console.log(select, '------<>', row)
        content.emit('selectTableData', select, row)
      },
      // 当用户手动勾选全选 Checkbox 时触发的事件
      selectAllTableData(selection) {
        // console.log(selection, 'selection')
        content.emit('selectAllTableData', selection)
      },
      // 当某个单元格被点击时会触发该事件
      rowClick(row, column, cell) {
        // console.log(row, 'row', column, 'column', cell)
        content.emit('rowClick', row, column, cell)
      },
      // 当某一行被双击时会触发该事件
      rowDblclick(row, column, event) {
        // console.log(row, column, event)
        content.emit('rowDblclick', row, column, event)
      },
      // 点击的单元格的回调
      cellClick(row, column, cell, event) {
        // console.log(row, column, cell, event, 'cellClick')
        content.emit('cellClick', row, column, cell, event)
      },
      // 排序的回调
      handleTableSort({ column, prop, order }) {
        // console.log(column, prop, order, 'handleTableSort')
        content.emit('handleTableSort', column, prop, order)
      },
      // 表头单元格的 style 的回调方法，
      handleTheadStyle() {
        content.emit('handleTheadStyle')
        let copy
        const style = {
          height: '52px',
          fontSize: '14px',
          fontWeight: '500',
          color: '#302E6D',
        }
        if (table.tableBasics.headerRowStyle) {
          copy = Object.assign(style, table.tableBasics.headerRowStyle)
        }
        return copy
      },
      rowStyle({ row, rowIndex }) {
        let styleJson = {
          padding: '10px',
        }
        content.emit('rowStyle', row, rowIndex)
        return styleJson
      },
      // 设置表头样式
      headerRowStyle() {
        // return 'background:#F3F4F7;color:#555'
      },
    }
    const tableRowClassName = () => {
      return ''
    }

    // 添加行
    const AddListRow = (name) => {
      content.emit('AddListRow', name)
    }

    let btnOptions = {
      bgColor: '#fff',
      textColor: '#595AEC',
      width: '60px',
      height: '26px',
    }

    // 显示列
    const isShowColsList = ref(false)
    const showColsList = () => {
      isShowColsList.value = !isShowColsList.value
    }
    const isIndeterminate = ref(false)
    const checkedItems = ref([])
    const checkAll = ref(true)
    const columsData = ref([])
    const handleCheckAllChange = (val: boolean) => {
      checkedItems.value = val ? colsList.value.map(item=>item.label) : []
      columsData.value = val ? colsList.value : [];
      isIndeterminate.value = false
      content.emit('colChange', columsData.value)
    }
    const handleCheckedItemssChange = (value: string[]) => {
      const checkedCount = value.length
      checkAll.value = checkedCount === colsList.value.length
      isIndeterminate.value = checkedCount > 0 && checkedCount < colsList.value.length
      columsData.value = []
      colsList.value.forEach(item=>{
          checkedItems.value.forEach(i=>{
            if(item.label===i){
            columsData.value.push(item)
          } 
        })
      })
      content.emit('colChange', columsData.value)
    }

    return {
      locale: zhCn,
      multipleTableRef,
      table,
      currentPage,
      pageSize,
      methods,
      tableRowClassName,
      AddListRow,
      btnOptions,
      checkAll,
      isIndeterminate,
      checkedItems,
      handleCheckAllChange,
      handleCheckedItemssChange,
      colsList,
      isShowColsList,
      showColsList
      // change,
      // columnsList,
    }
  },
})
</script>

<style scoped lang="scss">
@import '@/assets/style/var.scss';

.tableOperation {
  cursor: pointer;
  color: #409ffe;
}
.toolbar {
  background: #fff;
  display: flex;
  justify-content: flex-end;
}
.dropdown-list {
  margin-left: 10px;
}
.btns-slot {
  height: 50px;
  padding: 0 20px;
  margin-top: 10px;
  background: $btn-text-color;
  border-bottom: 1px solid #ededfd;
  line-height: 50px;
  display: flex;
}
.showCols{
  vertical-align: baseline;
  margin-left: 12px;
  position: relative;
  .showcols-list{
    position: absolute;
    top: 100%;
    right: 0;
    // left: -140px;
    z-index: 10;
    width: 160px;
    background: #FFFFFF;
    box-shadow: 0px 6px 24px 0px rgba(0, 0, 0, 0.16);
    border-radius: 6px;
    border: 1px solid #FFFFFF;
    padding: 7px 10px;
  }
}
</style>

<style lang="scss">
.showCols{
  .el-checkbox-group{
    label{
      .el-checkbox__label{
        width:100px;
        white-space:nowrap;
        overflow:hidden;
        text-overflow:ellipsis; 
      }
    }
  }
}
</style>
