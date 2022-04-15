<template>
  <div :class="['search-container', cname]" ref="searchContainerRef">
    <div class="search-wrapper" ref="searchWrapperRef">
      <el-row
        :gutter="queryItems.gutter || 0"
        :justify="queryItems.justify || 'start'"
        ref="searchRef"
        style="width: 76%"
      >
        <el-col
          v-for="(item, idx) of queryItems.columnList"
          :key="idx"
          class="search-col"
          :class="[idx === queryItems.columnList.length - 1 ? 'mb0' : '']"
        >
          <ZHInput
            v-if="item.type === 'input'"
            v-model.trim="item.value"
            width="160px"
            :placeholder="item.label"
          ></ZHInput>
          <ZHSelect
            v-if="item.type === 'select'"
            v-model="item.value"
            :options="item.options"
            :placeholder="item.label"
          ></ZHSelect>
          <ZHDatePicker
            v-if="item.type === 'datepicker'"
            v-model="item.value"
          ></ZHDatePicker>
          <ZHDropdownTree
            v-if="item.type === 'modelKind'"
            ref="myTest"
            v-model="item.value"
            :placeholder="item.label"
            :data="item.data"
            width="160px"
          >
          </ZHDropdownTree>
          <ZHSelectAndInput
            v-if="item.type === 'selectAndInput'"
            v-model:selectValue="item.selectValue"
            v-model="item.value"
            :options="item.options"
            :placeholder="item.label"
          >
          </ZHSelectAndInput>
          <ZHCheckbox
            v-if="item.type === 'zhcheckbox'"
            :options="item.options"
            v-model="item.value"
          ></ZHCheckbox>
        </el-col>
        <div ref="test">123</div>
      </el-row>
    </div>
    <slot>
      <div class="btn-container">
        <div v-if="isShowAdvanced">
          <span class="advance" @click="onToggle">
            <span>高级筛选</span>
            <svg-icon
              :icon-class="isExpand ? 'expand' : 'collapse'"
              class="advance-arrow"
            ></svg-icon
          ></span>
        </div>
        <div class="reset">
          <ZHButton :options="options" @ZHBtnClick="onReset">重置</ZHButton>
        </div>
      </div>
    </slot>
  </div>
</template>

<script lang="ts" setup>
import {
  defineProps,
  defineEmits,
  PropType,
  ref,
  watch,
  onMounted,
  nextTick,
} from 'vue'
import ZHInput from '@/components/Input/index.vue'
import ZHCheckbox from '@/components/FormCheckbox/index.vue'
import { IButton } from '@/components/Button/Button.vue'
import { debounce } from 'lodash'

export interface IOption {
  label: string
  value: string
}
export interface IColumnList {
  type: string
  label: string
  value: any
  options?: IOption[]
  data?: any[]
  selectValue?: string
  inputValue?: string
}
export interface ISearch {
  gutter?: number
  span?: number
  justify?: string
  columnList: IColumnList[]
  data?: any[]
  selectValue?: string
  inputValue?: string
}
//展示和隐藏的标识
let searchRef = ref(null)
let searchContainerRef = ref(null)
let searchWrapperRef = ref(null)
let isExpand = ref(true)
let options: IButton = {
  bgColor: '#E4E4FC',
  textColor: '#595AEC',
  border: '1px solid transparent',
}
let props = defineProps({
  cname: {
    type: String,
    default: '',
  },
  queryItems: {
    type: Object as PropType<ISearch>,
    default: () => ({}),
  },
})
let emit = defineEmits(['onSearch'])

watch(
  props.queryItems,
  debounce(() => {
    emit('onSearch')
  }, 500),
)

const onReset = () => {
  props.queryItems.columnList.forEach((item) => {
    if (typeof item.value === 'string') {
      item.value = ''
    } else if (item.label === '模型分类') {
      item.value = ''
    } else {
      item.value = []
    }
  })
  console.log(props.queryItems.columnList)
}
const onToggle = () => {
  if (isExpand.value) {
    searchWrapperRef.value.style.height =
      searchRef.value.$el.offsetHeight - 20 + 'px'
    searchWrapperRef.value.style.overflow = 'inherit'
    isExpand.value = false
  } else {
    searchWrapperRef.value.style.height = '30px'
    searchWrapperRef.value.style.overflow = 'hidden'
    isExpand.value = true
  }
}
let isShowAdvanced = ref(true)
let test = ref(null)
let myTest = ref(null)
onMounted(() => {
  let height = searchWrapperRef.value.firstChild.offsetHeight
  // console.log('height', height)
  // console.log(test, '---->', myTest)
  isShowAdvanced.value = height > 50 ? true : false
  // let nextDom = searchRef.value.$el.nextSiblings
  // if (nextDom.offsetHeight <= 160) {
  //   isShowAdvanced.value = false
  // }
})
</script>
<style lang="scss" scoped>
@import '@/assets/style/index.scss';
@mixin arrowStyle($deg) {
  position: absolute;
  left: 66px;
  width: 5px;
  height: 5px;

  border-left: 1px solid #595aec;
  transform: rotate($deg);
  transform-origin: center center;
  transition: all 0.2s;
}

.search-container {
  position: relative;
  width: 100%;
  padding: 20px;
  border-radius: 8px;
  transition: all 0.1s;
  background-color: #fff;
  box-sizing: border-box;
  .search-wrapper {
    position: relative;
    height: 30px;
    overflow: hidden;
    transition: all 0.1s;
  }

  .search-col {
    display: inline-flex;
    flex: inherit;
    align-items: center;
    margin-bottom: 20px;
    margin-right: 10px;
    .mb0 {
      margin-bottom: 0;
    }
  }
  .expand {
    position: absolute;
    width: 44px;
    height: 20px;
    bottom: -10px;
    left: 0;
    right: 0;
    margin: auto;
  }
  .btn-container {
    position: absolute;
    width: 24%;
    right: 10px;
    top: 20px;
    margin-left: 20px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .advance {
      display: inline-flex;
      margin-right: 10px;
      color: #595aec;
      font-size: 14px;
      cursor: pointer;
      user-select: none;
      align-items: center;
      & .advance-arrow {
        width: 32px;
        height: 32px;
      }
    }
    .reset {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      margin: auto 0;
    }
    .down-arrow {
      @include arrowStyle(-45deg);
      border-bottom: 1px solid #595aec;
    }
    .up-arrow {
      @include arrowStyle(45deg);
      border-top: 1px solid #595aec;
    }
  }
}
</style>
