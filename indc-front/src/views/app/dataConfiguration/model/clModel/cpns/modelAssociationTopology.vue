<template>
  <div class="canvas">
    <div class="logicflow-custom-panel" id="canvas-custom-dnd">
      <el-input v-model="filterText" :suffix-icon="Search" />
      <el-tree
        ref="treeRef"
        :data="data"
        :props="defaultProps"
        @node-click="onNodeClick"
        :filter-node-method="filterNode"
      />
    </div>
    <!--    <el-button @click="onNodeClick1">高亮一个点</el-button>-->
    <MyCanvas
      ref="canvasRef"
      class="canvas-wrapper"
      :pointSets="pointSets"
      :options="{ height: 580, isSilentMode: isSilentMode }"
      @onNodeClick="handleNodeClick"
      @onEdgeClick="handleEdgeClick($event)"
    ></MyCanvas>
    <newAssociationDrawer
      :isShowDrawer="drawerData.isShowDrawer"
      :isDisabled="drawerData.isDisabled"
      drawerTitle="查看关联"
      @closeDrawer="drawerData.closeDrawer"
    />

    <newModel
      :isNewModel="isShow"
      @cancelModel="cancelModel"
      @ModelOk="ModelOk"
    />
  </div>
</template>

<script lang="ts" setup>
import MyCanvas from '@/components/Canvas/index.vue'
import { reactive, ref, defineComponent, watch, onMounted } from 'vue'
import newAssociationDrawer from './newAssociationDrawer.vue'
import { computed } from '@vue/reactivity'
import store from '@/store'
import { drawerFormData, drawerData } from '../config/model-association'
import newModel from './newModel.vue'
import type { ElTree } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'
let isSilentMode = ref(false)
const route = useRoute()
// let type = route.query.type || ''
let type =  ''

isSilentMode.value = type !== 'look' ? false : true
const canvasRef = ref()

onMounted(() => {
  isSilentMode.value = type !== 'look' ? false : true
})

const pointSets = computed(() => store.state.pointSets)

console.log('pointSets', pointSets)

const handleNodeClick = (node) => {
  if (node.data.type === 'htmlEdge') {
    return
  }
  console.log('是否触发弹框', node)
  isShow.value = true
}
const handleEdgeClick = (edge) => {
  console.log('触发边的点击', edge)
  drawerData.isShowDrawer = true
  drawerData.isDisabled = true
  drawerData.title = '查看关联'
}

const onNodeClick = (treeNode) => {
  console.log('nodeClick', treeNode)
  if (treeNode.id) {
    // canvasRef.value.onNodeTrigger(treeNode.id)
    let nodesId = [
      'fe3d016b-6d51-4c39-aff3-21e2d99810fa',
      '8a44b871-900a-4e4b-994f-87c7483e4d74',
    ]
    canvasRef.value.onNodeTrigger(nodesId)
  }
}
const onNodeClick1 = () => {
  let nodesId = [
    'fe3d016b-6d51-4c39-aff3-21e2d99810fa',
    '8a44b871-900a-4e4b-994f-87c7483e4d74',
  ]
  canvasRef.value.onNodeTrigger(nodesId)
}
interface Tree {
  id: number
  label: string
  children?: Tree[]
}
const filterText = ref('')
const treeRef = ref<InstanceType<typeof ElTree>>()

const defaultProps = {
  children: 'children',
  label: 'label',
}
watch(filterText, (val) => {
  treeRef.value!.filter(val)
})

const filterNode = (value: string, data: Tree) => {
  if (!value) return true
  return data.label.includes(value)
}

const data: Tree[] = [
  {
    label: 'Level one 1',
    id: '8a44b871-900a-4e4b-994f-87c7483e4d74',
    children: [
      {
        label: 'Level two 1-1',
        children: [
          {
            label: 'Level three 1-1-1',
            id: 'd2f43de6-0dfd-4cd5-ab6a-ea795847400a',
            isPenultimate: true,
          },
          {
            label: 'Level three 1-1-2',
            id: 'fe3d016b-6d51-4c39-aff3-21e2d99810fa',
            isPenultimate: true,
          },
        ],
      },
    ],
  },
]

const isShow = ref(false)

const cancelModel = () => {
  isShow.value = false
}
const ModelOk = () => {
  isShow.value = false
}
</script>
<style lang="scss">
.logicflow-custom-panel {
  width: 220px;
  border-right: 1px solid #e4e4fc;
  padding-right: 20px;
}
.canvas {
  display: flex;
}
.canvas-wrapper {
  flex: 1;
}
.el-tree-node__label {
  user-select: none;
}
</style>
