<template>
  <div class="new-model" v-if="isDraw">
    <el-drawer
      v-model="isDraw"
      title="新建模型"
      :size="460"
      direction="rtl"
      :show-close="false"
      :destroy-on-close="true"
      :close-on-click-modal="false"
    >
      <el-form
        :model="state.modelForm"
        label-width="auto"
        :rules="state.rules"
        ref="ruleFormRef"
      >
        <el-form-item label="模型名称 : " prop="objName">
          <el-input
            v-model="state.modelForm.objName"
            style="width: 300px"
            :maxlength="30"
          />
          <el-tooltip
            class="box-item"
            effect="light"
            content="不允许重复"
            placement="top-start"
          >
            <span><svg-icon icon-class="tips"></svg-icon> </span>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="模型分类 : " prop="classificationName">
          <el-select
            v-model="state.modelForm.classificationName"
            placeholder="请选择"
            style="width: 300px; background: #fff"
            ref="selectRefs"
            fixed
            filterable
            :filter-method="filterFuc"
          >
            <el-option
              :value="state.modelForm.classificationId"
              :label="state.modelForm.classificationName"
              class="sel-option"
            >
              <el-tree
                ref="treeRef"
                @node-click="handleNodeClick"
                :expand-on-click-node="false"
                @node-expand="handleExpand"
                class="filter-tree"
                :data="data"
                accordion
                :props="{
                  children: 'children',
                  label: 'label',
                  value: 'id',
                }"
                :filter-node-method="filterNode"
              />
            </el-option>
          </el-select>
          <span class="el-tooltip__trigger" @click="addModelKind"
            ><svg-icon icon-class="addFrom"></svg-icon>
          </span>
        </el-form-item>
        <el-form-item label="模型描述 : " props="description">
          <el-input
            v-model="state.modelForm.description"
            type="textarea"
            placeholder="请输入"
            maxlength="30"
            :show-word-limit="true"
          />
        </el-form-item>
        <el-form-item label="选择图标 : " props="objIcon">
          <span class="defaultSvg" @click="choiceSvg">
            <svg-icon
              :icon-class="state.defaultSvg ? state.defaultSvg : 'defaultSvg'"
              class="default-icon"
            ></svg-icon>
            <div class="selectIcon" v-if="state.svgFlag">
              <span v-for="item in state.svgList" :key="item">
                <svg-icon
                  :icon-class="item"
                  @click="choice(item)"
                  class="choice"
                ></svg-icon>
              </span>
            </div>
            <span class="iconAngleMark" v-if="state.svgFlag"></span>
          </span>
        </el-form-item>
      </el-form>

      <template #footer>
        <div class="footerButton">
          <publicButton :data="cancelButtonRef" @onButton="cancelButton" />
          <publicButton
            :data="determineRef"
            @onButton="onButtonOk(ruleFormRef)"
          />
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script lang="ts" setup="props, content">
import {
  toRefs,
  ref,
  reactive,
  watch,
  defineProps,
  defineEmits,
  onMounted,
  onUnmounted,
} from 'vue'
import type { ElTree } from 'element-plus'
import type { ElForm } from 'element-plus'

import publicButton from '@/components/Button/index.vue'
import {
  determineConfig,
  cancelConfig,
  state,
} from '../hook/config/edit-config'

interface Tree {
  id: number
  label: string
  children?: Tree[]
  parentId?: number
}
const props = defineProps({
  isNewModel: {
    type: Boolean,
    default: true,
  },
})
const emit = defineEmits(['ModelOk', 'cancelModel'])
const { isNewModel: isDraw } = toRefs(props)
const determineRef = ref(determineConfig)
const cancelButtonRef = ref(cancelConfig)
const treeRef = ref(null)
type FormInstance = InstanceType<typeof ElForm>
const data: Tree[] = [
  {
    id: 1,
    label: 'Level one 1',
    parentId: 1,
    children: [
      {
        id: 4,
        label: 'Level two 1-1',
        parentId: 1,
        children: [
          {
            id: 9,
            label: 'Level three 1-1-1',
            parentId: 1,
          },
          {
            id: 10,
            label: 'Level three 1-1-2',
            parentId: 1,
          },
        ],
      },
    ],
  },
  {
    id: 2,
    label: 'Level one 2',
    parentId: 2,
    children: [
      {
        id: 5,
        label: 'Level two 2-1',
      },
      {
        id: 6,
        label: 'Level two 2-2',
      },
    ],
  },
  {
    id: 3,
    label: 'Level one 3',
    children: [
      {
        id: 7,
        label: 'Level two 3-1',
      },
      {
        id: 8,
        label: 'Level two 3-2',
      },
    ],
  },
]
const ruleFormRef = ref<FormInstance>()

// 确认
const onButtonOk = async (formEl: FormInstance | undefined) => {
  console.log('确认')
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      state.modelForm.choiceSvg = state.defaultSvg
      console.log(state.modelForm)
      emit('ModelOk', state.modelForm)
    } else {
      console.log('error submit!', fields)
    }
  })

  // emit('ModelOk')
}

// 点击icon图标
const choice = (item) => {
  console.log(item, 'item')
  state.defaultSvg = item
}

// 取消
const cancelButton = () => {
  console.log('取消')
  emit('cancelModel')
}

const filterNode = (value: string, data: Tree) => {
  if (!value) return true
  // console.log(value, data, '-----')
  return data.label.includes(value)
}

watch(state.modelForm.filter_text as any, (val) => {
  treeRef.value.filter(val)
})

onMounted(() => {
  // dom初始化
  treeRef.value = ref<InstanceType<typeof ElTree>>()
})
onUnmounted(() => {
  state.svgSum = 0
})
// 新建模型分类
const addModelKind = () => {
  console.log('新建模型分类')
}

// 选择svg
const choiceSvg = () => {
  console.log('选择svg')

  state.svgSum += 1
  if (state.svgSum % 2 === 0) {
    state.svgFlag = false
  } else {
    state.svgFlag = true
  }
}
// 搜索过滤
const filterFuc = (e) => {
  // console.log(e, '---->')
  state.modelForm.filter_text = e
}
// 点击树形图
const handleNodeClick = (node, data, e) => {
  if (data.parent.data.constructor === Array) {
    state.modelForm.filter_text = data.data.label
    state.modelForm.filter_id = data.data.id
    state.Administration = []
    console.log(state.modelForm, 'state----<>', state.Administration)
  } else if (data.parent.data.constructor === Object) {
    state.modelForm.filter_text = ''
    state.modelForm.filter_id = ''
    state.Administration = []
    processingDat(node, data.parent)
  }
}
const processingDat = (node, data) => {
  for (let key in data) {
    if (data && data !== null) {
      // debugger
      if (key === 'data' && data[key].constructor === Object) {
        if (node) {
          state.Administration.push(node.label)
        }
        state.Administration.push(data[key].label)
        state.Administration.reverse()
        state.Administration.forEach((item) => {
          state.modelForm.filter_text += item + '/'
        })
      }
      if (key === 'parent') {
        processingDat('', data[key])
      }
    }
  }
  console.log(state.modelForm, 'state.modelForm')
}
const handleExpand = (e) => {
  //  state.modelForm.filter_text = ''
  //  state.modelForm.filter_id = ''
  state.Administration = []
}
</script>

<style lang="scss" scoped>
@import '@/assets/style/index.scss';
.footerButton {
  display: flex;
  justify-content: flex-end;
  margin-top: 20px;
  div {
    margin-right: 20px;
  }
}
.el-tooltip__trigger {
  margin-left: 10px;
  margin-top: 5px;
}
.el-form-item {
  margin-right: 10px;
  margin-top: 20px;
}
:deep(.el-textarea__inner) {
  width: 300px;
  height: 90px;
}
:deep(.el-textarea .el-input__count) {
  right: 50px !important;
}
.defaultSvg {
  padding-top: 3px;
}
.default-icon {
  width: 36px;
  height: 36px;
}
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
  z-index: 1;
}
.selectIcon {
  position: relative;
  width: 300px;
  height: auto;
  border: 1px solid #e4e4fc;
  background: #fff;
}
.iconAngleMark::before {
  position: absolute;
  top: 45px;
  left: 10px;
  width: 10px;
  height: 10px;
  // z-index: -1;
  content: ' ';
  transform: rotate(45deg);
  background: #e4e4fc;
  box-sizing: border-box;
  border: 1px solid var(--el-border-color-light);
  border-bottom-color: transparent !important;
  border-right-color: transparent !important;
  border-top-left-radius: 2px;
  background: var(--el-color-white);
  right: 0;
}
.choice {
  cursor: pointer;
  width: 24px;
  height: 24px;
  margin-left: 10px;
  margin-top: 10px;
}
</style>
