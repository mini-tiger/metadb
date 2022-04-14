<template>
  <div class="canvas-container">
    <div class="btn-wrapper" v-show="isShowBtn">
      <el-button @click.stop="onDeleteEdge" class="delete-btn"
        >删除边</el-button
      >
      <span title="全屏" class="canvas-btn">
        <svg-icon @click="toggle" icon-class="canvas-full"></svg-icon>
      </span>
      <span title="自适应" class="canvas-btn">
        <svg-icon @click="onAdjust" icon-class="canvas-adjust"></svg-icon>
      </span>
      <span title="缩小" class="canvas-btn">
        <svg-icon @click="onZoomOut" icon-class="canvas-minus"></svg-icon>
      </span>
      <span title="放大" class="canvas-btn">
        <svg-icon @click="onZoomIn" icon-class="canvas-add"></svg-icon>
      </span>
    </div>
    <div ref="flow"></div>
  </div>
</template>

<script lang="ts" setup>
import { v4 as uuidv4 } from 'uuid'
import {
  ref,
  onMounted,
  defineProps,
  defineEmits,
  defineExpose,
  computed,
  onUnmounted,
  watch,
  PropType,
} from 'vue'
import * as types from '@/store/action-types'
import htmlNode from '@/components/Canvas/htmlNode'
import htmlEdge from '@/components/Canvas/htmlEdge'
import htmlLineEdge from '@/components/Canvas/htmlLineEdge'

import { ContextPad } from '@/components/Canvas/contextPad.js'
import setContextPad from '@/components/Canvas/setContextPad.js'
import LogicFlow from '@logicflow/core'
import '@logicflow/core/dist/style/index.css'
import { api as fullscreen } from 'vue-fullscreen'
import { useStore } from 'vuex'

export type IUrlType = '1' | '2' | '3' | '4'
let lf: LogicFlow | null = null
let isUpdateData = ref(true)

const emit = defineEmits([
  'onEdgeClick',
  'onNodeClick',
  'update:dialog',
  'update:isAddEdge',
])
//高亮的边
const flow = ref()
const props = defineProps({
  options: {
    type: Object,
    default: () => ({}),
  },
  dialog: {
    type: Boolean,
    default: false,
  },
  isResetHighlight: {
    type: Boolean,
    default: true,
  },
  //用于接收拓扑图的数据
  pointSets: {
    type: Object,
    default: () => ({}),
  },
  //控制右上角的按钮开关
  isShowBtn: {
    type: Boolean,
    default: false,
  },
  //是否增加边
  isAddEdge: {
    type: Boolean,
    default: false,
  },
  //用于区分提交的后台
  urlType: {
    type: String as PropType<IUrlType>,
    default: '1',
  },
})
const defautOptions = {
  //图的 DOM 容器
  container: '',
  height: 640,
  //是否开启历史记录功能
  history: false,
  //是否允许拖动边的端点来调整连线
  adjustEdgeStartAndEnd: false,
  //仅浏览不可编辑模式，默认不开启
  isSilentMode: false,
  //禁止鼠标滚动移动画布
  stopScrollGraph: true,
  //禁止缩放画布
  stopZoomGraph: true,
  //是否开启文本编辑
  textEdit: false,
  //允许调整边
  adjustEdge: false,
  //网格，若设为false不开启网格，则为 1px 移动单位，不绘制网格背景，若设置为true开启则默认为 20px 点状网格
  grid: true,
  //是否隐藏节点的锚点，静默模式下默认隐藏
  hideAnchors: false,
  //鼠标hover的时候显示节点的外框
  hoverOutline: false,
  //鼠标hover的时候显示节点的外框
  nodeSelectedOutline: false,
  //是否允许拖动节点
  adjustNodePosition: true,
  //自定义键盘相关配置
  keyboard: {
    enabled: true,
  },
  plugins: [ContextPad],
}
const store = useStore()
const renderRef = ref(props.pointSets)
const currentNode = ref()
const currentEdge = ref()
// const isResetEdge = ref(props.isResetHighlight)

//点击侧边栏触发画布中的节点
const onNodeTrigger = (ids) => {
  const { eventCenter } = lf?.graphModel
  //主动触发节点的点击
  resetNode()
  ids.forEach((id) => {
    let data = renderRef.value.nodes.find((node) => node.id === id)
    data && eventCenter.emit('element:click', { data, flag: 'false' })
  })
}
const resetNode = () => {
  let allDom = document.querySelectorAll('.node-content')
  allDom.forEach((dom: any) => {
    dom.classList.remove('is-active')
    dom.nextSibling.classList.remove('text-active')
  })
}
const resetEdge = (active = false) => {
  let filterEdges = computed(() => store.state.pointSets.edges)
  if (filterEdges.value?.length > 0) {
    filterEdges.value.forEach((edge) => {
      const edgeModel = lf.getEdgeModelById(edge.id)
      edgeModel?.setProperties({
        isActived: active,
      })
    })
  }
}
const onZoomIn = () => {
  isUpdateData.value = false
  lf?.zoom(true)
}
const onZoomOut = () => {
  isUpdateData.value = false
  lf?.zoom(false)
}
const onAdjust = () => {
  const { transformModel } = lf?.graphModel
  transformModel.zoom(1)
}
let isFull = ref(false) //用于切换图标
const toggle = () => {
  fullscreen.toggle(document.querySelector('.canvas-container'), {
    fullscreenClass: 'full-class',
    callback: (isFullscreen) => {
      isFull.value = isFullscreen
    },
  })
}
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const onCancelEge = () => {
  lf.graphModel.deleteEdgeById(currentEdge.value.id)
}
const onDeleteEdge = () => {
  lf.graphModel.deleteEdgeById(currentEdge.value.id)
  //同时需要更新vuex中的数据
  store.dispatch(types.DELETE_EDGE, currentEdge.value.id)
}
//只有查询的时候高亮线
const onHighLightEdge = (edges) => {
  //找到相关的线，增加属性，进行高亮
  resetNode()
  let allEdge = computed(() => store.state.pointSets.edges)
  allEdge.value.forEach((edge) => {
    if (edges.includes(edge.id)) {
      const edgeModel = lf.getEdgeModelById(edge.id)
      edgeModel.setProperties({
        isActived: true,
      })
    }
  })
}
const highlightNode = (point) => {
  let dom: any = document.querySelector('#' + point.data.properties.id)
  dom?.classList.add('is-active')
  dom?.nextSibling.classList.add('text-active')
}
let isResetNode = ref(true)
const resetActiveStatus = () => {
  resetEdge(false)
  isResetNode.value && resetNode()
}

const onElementClick = (point) => {
  if (point.flag !== 'false') {
    emit('onNodeClick', point)
    isResetNode.value = true
  } else {
    isResetNode.value = false
  }
  //需要判断是否先进行高亮的重置
  resetActiveStatus()
  //高亮点
  highlightNode(point)
}

const initCanvas = () => {
  defautOptions.container = flow.value
  let options: any = Object.assign({}, defautOptions, props.options)
  lf = new LogicFlow(options)
  lf.register(htmlNode)
  lf.register(htmlEdge)
  lf.register(htmlLineEdge)
  lf.setDefaultEdgeType('htmlLineEdge')
  setContextPad(lf)
  //设置箭头的样式
  lf.setTheme({
    arrow: {
      offset: 6,
      verticalLength: 4,
    },
  })
  lf.render(renderRef.value)
}
//进入到节点
let timer = null
const onMouseEnter = (point: any) => {
  if (timer) {
    clearTimeout(timer)
    timer = null
  }
  currentNode.value = point.data
  const dom: HTMLElement = document.querySelector(
    '#' + point.data.properties.id,
  )
  if (dom) {
    dom.style.display = 'inline-block'
  }
}
//节点的放开
const onNodeDrop = () => {
  upateGraphData()
}
const onNodeDragStart = (point: any) => {
  currentNode.value = point.data
}
const upateGraphData = ({ updateData = false, x = 0, y = 0 } = {}) => {
  let edges = lf.getGraphData().edges
  let nodes = lf.getGraphData().nodes
  let allEdges = edges.filter((edge) => store.getters.edgeIds.includes(edge.id))

  if (updateData) {
    nodes.forEach((node: any) => {
      node.x += x
      node.y += y
    })
    allEdges.forEach((edge) => {
      edge?.pointsList?.forEach((point) => {
        point.x += x
        point.y += y
      })
      //更新起点和终点坐标，以及文字坐标
      if (edge.type !== 'htmlLineEdge') {
        edge.startPoint.x = edge?.pointsList[0]?.x
        edge.startPoint.y = edge?.pointsList[0]?.y
        edge.endPoint.x = edge?.pointsList[3]?.x
        edge.endPoint.y = edge?.pointsList[3]?.y
      }
      if (edge.text && edge.text.x && edge.text.y) {
        edge.text.x += x
        edge.text.y += y
      }
    })
  }
  let allGraphData = {
    edges: allEdges,
    nodes,
  }

  store.dispatch(types.CHANGE_POSITION, {
    allGraphData,
    urlType: props.urlType,
  })
}
//节点从外部拖入时更新
const onNodeDragAndDrop = (point: any) => {
  store.commit(types.ADD_POINT, {
    id: point.data.id,
    x: point.data.x || 0,
    y: point.data.y || 0,
    type: 'htmlNode',
    properties: point.data.properties,
  })
}

const onMouseLeave = () => {
  currentNode.value = ''
  timer = setTimeout(() => {
    hideIcon()
  }, 100)
}
const hideIcon = () => {
  const doms: any = document.querySelectorAll('.delete-icon')
  doms.forEach((dom) => {
    dom.style.display = 'none'
  })
}
//边的点击
const onEdgeClick = (edge: any) => {
  currentEdge.value = edge.data
  emit('onEdgeClick', currentEdge.value)
}

//画布的移动，更新所有数据
const onGraphTransform = (data: any) => {
  if (isUpdateData.value) {
    const updateX = parseFloat(data.transform.TRANSLATE_X)
    const updateY = parseFloat(data.transform.TRANSLATE_Y)
    upateGraphData({ updateData: true, x: updateX, y: updateY })
  }
}
const onEdgeAdd = (edge) => {
  // //此处需要动态获取
  edge.data.text = {
    value: '父子',
  }

  currentEdge.value = edge.data
}
const onBlankDragstart = () => {
  isUpdateData.value = true
}
const onAddEvenEdge = (allEdge) => {
  lf.graphModel.deleteEdgeBySourceAndTarget(
    allEdge[allEdge.length - 1].sourceNodeId,
    allEdge[allEdge.length - 1].targetNodeId,
  )
  allEdge.forEach((edge, idx, arr) => {
    let n = arr.length
    let gap = 30
    edge.pointsList = []
    edge.type = 'htmlEdge'
    let second = {
      y:
        edge.startPoint.y -
        (n === 2 ? (2 * idx - 1) * gap * 2 : ((n - 1) * -gap) / 2 + gap * idx),
      x: edge.startPoint.x + gap,
    }
    let third = {
      y:
        edge.endPoint.y -
        (n === 2 ? (2 * idx - 1) * gap * 2 : ((n - 1) * -gap) / 2 + gap * idx),
      x: edge.endPoint.x - gap,
    }
    edge.pointsList = [edge.startPoint, second, third, edge.endPoint]
    store.dispatch(types.DELETE_EDGE, edge.id)
    edge.properties.isActived = false
    lf.addEdge(edge)
    store.commit(types.ADD_EDGE, edge)
  })
}

const onAddNewEdge = (edge) => {
  //先判断目前的两个节点的边的个数
  // let sourceNodeId = edge.sourceNodeId
  // let targetNodeId = edge.targetNodeId
  // let allEdge = computed(() => store.state.pointSets.edges)
  // let filterEdge = allEdge.value.filter((edge) => {
  //   return (
  //     edge.sourceNodeId === sourceNodeId && edge.targetNodeId === targetNodeId
  //   )
  // })
  // if (filterEdge.length === 0) {
  //   lf.graphModel.deleteEdgeBySourceAndTarget(sourceNodeId, targetNodeId)
  //   edge.type = 'htmlLineEdge'
  //   // edge.text.x = 0
  //   // edge.text.y = 0
  //   edge.properties.isActived = false
  //   lf.addEdge(edge)
  // } else {
  //   if ((filterEdge.length + 1) % 2 === 0) {
  //     //增加偶数边
  //     filterEdge.push(edge)
  //     onAddEvenEdge(filterEdge)
  //   } else {
  //     //增加奇数边
  //     lf.graphModel.deleteEdgeById(edge.id)
  //     edge.type = 'htmlLineEdge'
  //     edge.properties.isActived = false
  //     lf.addEdge(edge)
  //   }
  // }
  lf.addEdge(edge)
  store.commit(types.ADD_EDGE, edge)
}
//锚点拖放时触发
const onAnchorDrop = () => {
  //前台的显示
  if (!currentEdge.value) return
  //触发弹框
  emit('update:dialog', true)
}
//用于检测是否增加边
watch(
  () => props.isAddEdge,
  (isAddEdge) => {
    if (isAddEdge) {
      onAddNewEdge(currentEdge.value)
      emit('update:isAddEdge', false)
    }
  },
)

onMounted(() => {
  initCanvas()
  //判断当前是否为静默模式，阻止画布的拖动
  if (lf.getEditConfig().isSilentMode === true) {
    lf.updateEditConfig({
      hideAnchors: true,
    })
  }
  const { eventCenter } = lf.graphModel

  eventCenter.on('element:click', onElementClick)
  eventCenter.on('node:mouseenter', onMouseEnter)
  eventCenter.on('node:dragstart', onNodeDragStart)
  eventCenter.on('node:drop', onNodeDrop)
  eventCenter.on('node:dnd-add', onNodeDragAndDrop)
  eventCenter.on('node:mouseleave', onMouseLeave)

  eventCenter.on('edge:click', onEdgeClick)
  eventCenter.on('edge:add', onEdgeAdd)
  eventCenter.on('graph:transform', onGraphTransform)
  eventCenter.on('blank:dragstart', onBlankDragstart)
  eventCenter.on('blank:click', resetActiveStatus)

  eventCenter.on('anchor:drop', onAnchorDrop)

  //在画布中生成节点
  document
    .querySelector('#canvas-custom-dnd')
    .addEventListener('mousedown', (e) => {
      const type = (e.target as HTMLElement).nodeName
      const text = (e.target as HTMLElement).innerText
      if (type === 'SPAN') {
        lf?.dnd.startDrag({
          type: 'htmlNode',
          properties: {
            id: 'node' + uuidv4(),
            text: `${text}节点`,
          },
        })
      }
    })
})
onUnmounted(() => {
  const { eventCenter } = lf.graphModel
  eventCenter.off('element:click', onElementClick)
  eventCenter.off('node:mouseenter', onMouseEnter)
  eventCenter.off('node:dragstart', onNodeDragStart)
  eventCenter.off('node:drop', onNodeDrop)
  eventCenter.off('node:dnd-add', onNodeDragAndDrop)
  eventCenter.off('node:mouseleave', onMouseLeave)

  eventCenter.off('edge:click', onEdgeClick)
  eventCenter.off('edge:add', onEdgeAdd)
  eventCenter.off('graph:transform', onGraphTransform)
  eventCenter.off('blank:dragstart', onBlankDragstart)
  eventCenter.off('blank:click', resetActiveStatus)

  eventCenter.off('anchor:drop', onAnchorDrop)
})

defineExpose({
  onNodeTrigger,
  onDeleteEdge,
  onAnchorDrop, //连线的方法导出
  onCancelEge,
  onHighLightEdge, //高亮线
})
</script>
<style lang="scss" scoped>
@import '@/assets/style/var.scss';
.canvas-container {
  position: relative;
  user-select: none;
  &.full-class {
    position: absolute;
    width: 100%;
    height: 100%;
    background: #fff;
  }
  .btn-wrapper {
    position: absolute;
    right: 30px;
    top: 20px;
    z-index: 1;
    .canvas-btn {
      display: inline-block;
      margin-left: 20px;
    }
  }
  :deep(.node-content) {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    background-color: #5c5c8d;
  }
  :deep(.delete-icon) {
    position: absolute;
    display: none;
    width: 12px;
    height: 12px;
    margin-left: 26px;
    margin-top: -16px;
    background: url('./view.png') no-repeat left center;
    background-size: 12px 12px;
    padding: 10px;
  }
  :deep(foreignObject) {
    overflow: inherit;
  }
  :deep(.node-content) {
    font-size: 12px;
  }
  :deep(.content-text) {
    font-size: 12px;
    color: #302e6d;
    &.text-active {
      color: $main-color;
    }
  }
  :deep(.is-active) {
    background-color: $main-color;
  }
  :deep(.lf-inner-context) {
    position: absolute;
  }
  :deep(.lf-context-img) {
    width: 20px;
    height: 20px;
    cursor: pointer;
  }
  // :deep(.view-wrapper) {
  //   padding: 20px 8px;
  // }
}
</style>
