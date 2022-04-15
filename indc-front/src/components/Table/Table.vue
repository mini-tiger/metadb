<template>
  <el-table
    class="el-table-container"
    :data="tableData"
    :border="border"
    max-height="maxHeight"
    ref="tableRef"
    :header-row-class-name="'header-row-class-name'"
    :header-cell-class-name="'header-cell-class-name'"
    :size="size"
    @sort-change="onSortChange"
    @selection-change="onSelectionChange"
  >
    <el-table-column
      v-if="isShowIndex"
      class="filtered"
      :type="indexType"
      :width="56"
      label="序号"
      headerAlign="center"
      align="center"
    />
    <el-table-column
      v-for="(item, index) in headerCol"
      :width="item.width || 'auto'"
      :min-width="item.minWidth || 'auto'"
      :key="`col_${index}`"
      :prop="dropCol[index].prop"
      :label="item.label"
      :headerAlign="item.headerAlign || 'left'"
      :align="dropCol[index].align || 'left'"
      :sortable="item.sortable"
      show-overflow-tooltip
    >
    </el-table-column>
    <template #empty>
      <div class="empty-wrapper">没有数据</div>
    </template>
    <el-table-column fixed="right" label="操作" width="130">
      <template v-slot:default="scope">
        <slot name="operation" :scopeRow="{ item: scope.row }"></slot>
      </template>
    </el-table-column>
  </el-table>
  <div class="pagination-wrapper" v-if="isShowPagination">
    <el-pagination
      v-model:currentPage="page"
      v-model:page-size="currentSize"
      :page-sizes="pageSizes"
      layout="total, sizes, prev, pager, next, jumper"
      :total="400"
      :background="true"
    />
  </div>
</template>

<script lang="ts" setup>
import Sortable from 'sortablejs'
import { cloneDeep } from 'lodash'
import {
  ref,
  onMounted,
  defineProps,
  defineEmits,
  PropType,
  watch,
  nextTick,
} from 'vue'
export interface IColumn {
  label: string
  prop: string
  width?: string
  minWidth?: string
  headerAlign?: string
  align?: string
  sortable?: boolean | string
}
const props = defineProps({
  tableData: {
    type: Array,
    default: () => [],
  },
  headerCol: {
    type: Array as PropType<IColumn[]>,
    default: () => [],
  },
  height: {
    type: String,
    default: '650',
  },
  maxHeight: {
    type: String,
    defaut: '650',
  },
  isDraggable: {
    type: Boolean,
    default: true,
  },
  isShowIndex: {
    type: Boolean,
    default: true,
  },
  border: {
    type: Boolean,
    default: false,
  },
  size: {
    type: String,
    default: 'default',
  },
  indexType: {
    type: String,
    default: 'selection',
  },
  defaultSort: {
    type: Object,
    default: () => ({ prop: 'date' }),
  },
  isShowPagination: {
    type: Boolean,
    default: true,
  },
  currentPage: {
    type: Number,
    default: 1,
  },
  pageSize: {
    type: Number,
    default: 10,
  },
  pageSizes: {
    type: Array,
    default: () => [10, 20, 50, 100],
  },
})
let emit = defineEmits([
  'updateColumn',
  'onSortChange',
  'update:currentPage',
  'update:pageSize',
])
let tableRef = ref()
let tableData = ref(props.tableData)
let dropCol = ref(cloneDeep(props.headerCol))
let page = ref(props.currentPage)
let currentSize = ref(props.pageSize)

const initDraggable = () => {
  //列的拖拽
  new Sortable(
    tableRef.value.$el.querySelector('.el-table__header-wrapper tr'),
    {
      filter: '.filtered',
      animation: 50,
      delay: 50,
      onEnd: (evt) => {
        console.log('evt', evt)
        if (!props.isShowIndex) {
          const oldItem = dropCol.value[evt.oldIndex]
          dropCol.value.splice(evt.oldIndex, 1)
          dropCol.value.splice(evt.newIndex, 0, oldItem)
        } else {
          const oldItem = dropCol.value[evt.oldIndex - 1]
          dropCol.value.splice(evt.oldIndex - 1, 1)
          dropCol.value.splice(evt.newIndex - 1, 0, oldItem)
        }

        emit('updateColumn', dropCol.value)
      },
      draggable: '.not-filtered',
      ghostClass: 'hoverBackground',
    },
  )
}
const onSortChange = ({ order, prop }) => {
  emit('onSortChange', { order, prop })
}
const onSelectionChange = (selectData) => {
  console.log('val', selectData)
}
watch(page, (currentPage) => {
  emit('update:currentPage', currentPage)
})
watch(currentSize, (pagesize) => {
  emit('update:pageSize', pagesize)
})

onMounted(() => {
  props.isDraggable && initDraggable()
  nextTick(() => {
    //需要把索引列和操作列排除掉，不可以拖动
    let dom: HTMLElement = document.querySelector('.header-row-class-name')
    let domArray = Array.from(dom.children)
    domArray.forEach((child, idx) => {
      if (idx === 0 || idx === domArray.length - 1) {
        child.classList.add('filtered')
      } else {
        child.classList.add('not-filtered')
      }
    })
  })
})
</script>
<style lang="scss" scoped>
@import '@/assets/style/var.scss';
.el-table-container {
  border-top: 1px solid #ededfd;
  min-height: 450px;
  font-size: 12px;
}
:deep(.header-row-class-name) {
  height: 50px;
  font-weight: 500;
  font-size: 14px;
  color: #302e6d;
}
:deep(.el-button--text) {
  font-size: 12px;
}
.el-table .ascending .sort-caret.ascending {
  border-bottom-color: $main-color !important;
}
:deep(.hoverBackground) {
  background-color: $main-color !important;
  color: #fff;
}
.pagination-wrapper {
  display: flex;
  justify-content: flex-end;
  padding: 20px;
  background-color: #fff;
}
.empty-wrapper {
  display: flex;
  min-height: 400px;
  justify-content: center;
  align-items: center;
}
</style>
