<template>
  <el-input
    :modelValue="modelValue"
    @input="onInput"
    @blur="onBlur"
    @focus="onFocus"
    @change="onChange"
    class="zh-input"
    :placeholder="placeholder"
    :show-word-limit="showWordLimit"
    :maxlength="maxlength"
    :type="type"
    :clearable="clearable"
  ></el-input>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, ref } from 'vue'
let props = defineProps({
  modelValue: {
    type: String,
    defaut: '',
  },
  placeholder: {
    type: String,
    default: '请输入',
  },
  showWordLimit: {
    type: Boolean,
    default: false,
  },
  maxlength: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'text',
  },
  clearable: {
    type: Boolean,
    default: false,
  },
  width: {
    type: String,
    default: '220px',
  },
  height: {
    type: String,
    default: '30px',
  },
})
const width = ref(props.width)
const height = ref(props.height)
let emit = defineEmits([
  'update:modelValue',
  'ZHInputBlur',
  'ZHInputFocus',
  'ZHInputChange',
])

const onInput = (value) => {
  emit('update:modelValue', value)
}
const onBlur = () => {
  emit('ZHInputBlur')
}
const onFocus = () => {
  emit('ZHInputFocus')
}
const onChange = () => {
  emit('ZHInputChange')
}
</script>
<style lang="scss" scoped>
.zh-input {
  width: v-bind(width);
  :deep(.el-input__inner) {
    height: v-bind(height);
    line-height: v-bind(height);
    padding: 0 10px;
    border: 1px solid #b8b8d5;
    &::-webkit-input-placeholder {
      color: #a1a1c3;
      font-size: 12px;
    }
    &::input-placeholder {
      color: #a1a1c3;
      font-size: 12px;
    }
  }
}
</style>
