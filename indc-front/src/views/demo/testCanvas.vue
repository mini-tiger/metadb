<template>
  <el-drawer
    ref="drawerRef"
    v-model="dialog"
    title="I have a nested form inside!"
    direction="ltr"
    custom-class="demo-drawer"
  >
    <div class="demo-drawer__content">
      <el-form :model="form">
        <el-form-item label="Name" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="Area" :label-width="formLabelWidth">
          <el-select
            v-model="form.region"
            placeholder="Please select activity area"
          >
            <el-option label="Area1" value="shanghai"></el-option>
            <el-option label="Area2" value="beijing"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div class="demo-drawer__footer">
        <el-button @click="cancelForm">Cancel</el-button>
        <el-button type="primary" @click="onClick">confirm</el-button>
      </div>
    </div>
  </el-drawer>
  <div class="canvas">
    <div class="logicflow-custom-panel" id="canvas-custom-dnd">
      <el-tree :data="data" :props="defaultProps" @node-click="onNodeClick" />
    </div>
    <el-button @click="onNodeClick1">高亮一个点</el-button>
    <el-button @click="onNodeClick2">高亮新点</el-button>
    <el-button @click="onNodeClick3">查询高亮线</el-button>
    <!-- <HintTone></HintTone> -->
    <MyCanvas
      v-model:dialog="dialog"
      ref="canvasRef"
      class="canvas-wrapper"
      :pointSets="pointSets"
      v-model:isAddEdge="isAddEdge"
      :urlType="urlType"
      @onNodeClick="handleNodeClick"
      @onEdgeClick="handleEdgeClick($event)"
    ></MyCanvas>
  </div>
</template>

<script lang="ts" setup>
import MyCanvas, { IUrlType } from '@/components/Canvas/index.vue'
import { reactive, ref } from 'vue'
import type { ElDrawer } from 'element-plus'
import { computed } from '@vue/reactivity'
import store from '@/store'
// import HintTone from '@/components/HintTone/index.vue'
const urlType = ref<IUrlType>('1')
const formLabelWidth = '80px'
const dialog = ref(false)
// const loading = ref(false)
const canvasRef = ref()

const form = reactive({
  name: '',
  region: '',
  date1: '',
  date2: '',
  delivery: false,
  type: [],
  resource: '',
  desc: '',
})
const pointSets = computed(() => store.state.pointSets)
const drawerRef = ref<InstanceType<typeof ElDrawer>>()
let isAddEdge = ref(false)
const onClick = () => {
  dialog.value = false
  isAddEdge.value = true
}
const handleEdgeClick = (edge) => {
  console.log('触发边的点击', edge)
}
//这里需要先判断点击的类型，否则点击边会触发节点的点击
const handleNodeClick = (node) => {
  if (node.data.type !== 'htmlNode') {
    return
  }
  console.log('是否触发弹框', node)
}

const cancelForm = () => {
  dialog.value = false
  canvasRef.value.onCancelEge()
}
const onNodeClick = (treeNode) => {
  if (treeNode.id) {
    // canvasRef.value.onNodeTrigger(treeNode.id)
    let nodesId = [
      'cb7f2a40-f80f-4bf0-ad73-d937d8b13fa1',
      '002cd3a6-e48f-4de8-86b4-ac662146054f',
    ]
    canvasRef.value.onNodeTrigger(nodesId)
  }
}
const onNodeClick1 = () => {
  let nodesId = ['cb7f2a40-f80f-4bf0-ad73-d937d8b13fa1']
  canvasRef.value.onNodeTrigger(nodesId)
}
const onNodeClick2 = () => {
  let nodesId = [
    '2f4145dd-faf0-4c6f-a5b0-2648fc048989',
    '13c465e4-1ad7-4cd4-9741-37b61003b451',
  ]
  canvasRef.value.onNodeTrigger(nodesId)
}
const onNodeClick3 = () => {
  let edgeIds = [
    '26e1efe2-8fb1-4a6a-b5dc-33d78377be68',
    '7d9ba0e2-9a7d-4195-9b42-d41f1cd2d2c6',
  ]
  canvasRef.value.onHighLightEdge(edgeIds)
}
const defaultProps = {
  children: 'children',
  label: 'label',
}

const data = [
  {
    label: 'Level one 1',
    children: [
      {
        label: 'Level two 1-1',
        children: [
          {
            label: 'Level three 1-1-1',
            id: '7b09dd69-2611-4f99-bf39-20c61715ec61',
            isPenultimate: true,
          },
        ],
      },
    ],
  },
]
</script>
<style lang="scss">
.logicflow-custom-panel {
  width: 200px;
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
