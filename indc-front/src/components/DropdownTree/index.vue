<template>
  <div>
    <el-select
      :modelValue="state.modelValue1"
      :placeholder="placeholder"
      ref="selectRefs"
      fixed
      filterable
      :style="{ width: width ? width : '180px' }"
      :filter-method="filterFuc"
    >
      <el-option
        :value="state.modelValueId"
        :label="state.modelValue1"
        class="sel-option"
      >
        <el-tree
          ref="treeRef"
          style="width: 180px"
          @node-click="handleNodeClick"
          :expand-on-click-node="false"
          @node-expand="handleExpand"
          class="filter-tree"
          :data="data"
          accordion
          :props="{
            children: 'children',
            label: 'label',
          }"
          :filter-node-method="filterNode"
        />
      </el-option>
    </el-select>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref, watch, defineProps, defineEmits } from 'vue'
import type { ElTree } from 'element-plus'
const props = defineProps({
  data: {
    type: Array,
    default: null,
  },
  modelValue: {
    type: String,
    default: '',
  },
  placeholder: {
    type: String,
    default: '',
  },
  width: {
    type: String,
    default: '',
  },
})
const emit = defineEmits(['update:modelValue', 'resetSelectTree'])

let modelValue1 = ref(props.modelValue)

const state = reactive({
  modelValue1: props.modelValue,
  modelValueId: '',
  Administration: [],
})

const treeRef = ref<InstanceType<typeof ElTree>>()
const selectRefs = ref()
const filterNode = (value, data) => {
  if (!value) return true
  return data.label.includes(value)
}
// 搜索过滤
const filterFuc = (e) => {
  modelValue1.value = e
}
watch(modelValue1, (val) => {
  treeRef.value.filter(val)
})

// 点击树形图
const handleNodeClick = (node, data, e) => {
  if (data.parent.data.constructor === Array) {
    console.log(selectRefs.value, '----->')
    state.modelValue1 = data.data.label
    state.modelValueId = data.data.id
    state.Administration = []
    // const arr = []
    // arr.push({
    //   label: state.modelValue1,
    //   id: state.modelValueId,
    // })
    emit('update:modelValue', state.modelValueId)
    selectRefs.value.visible = false
  } else if (data.parent.data.constructor === Object) {
    state.modelValue1 = ''
    state.modelValueId = ''
    state.Administration = []
    processingDat(node, data.parent)
  }
  selectRefs.value.visible = false // 点击下卡框收回
  console.log(state.modelValue1, state.modelValueId)
}

const processingDat = (node, data) => {
  for (let key in data) {
    if (data && data !== null) {
      // debugger
      if (key === 'data' && data[key].constructor === Object) {
        if (node) {
          state.Administration.push({
            label: node.label,
            id: node.id,
          })
        }
        state.Administration.push({
          label: data[key].label,
          id: data[key].id,
        })
        state.Administration.reverse()
        state.Administration.forEach((item) => {
          state.modelValue1 += item.label + '/'
          state.modelValueId += item.id + '-'
        })
        // console.log(state.modelValue1, state.modelValueId, '-------->-->')
        emit('update:modelValue', state.modelValueId)
      }
      if (key === 'parent') {
        processingDat('', data[key])
      }
    }
  }
  // emit('update:modelValue', state.Administration)
}
const handleExpand = () => {
  state.Administration = []
}
</script>

<style lang="scss" scoped>
.sel-option {
  width: 300px;
  height: auto;
  max-height: 200px;
  // overflow: auto;
  background-color: #fff;
  cursor: pointer;
  font-weight: normal !important;
  padding: 0 5px;
}
.el-scrollbar
  .el-select-dropdown__wrap
  .el-select-dropdown__list
  .el-select-dropdown__item {
  height: auto;
}
.el-scrollbar .el-select-dropdown__wrap .el-select-dropdown__list .hover {
  background-color: #fff !important;
  // z-index: 1;
}
</style>
