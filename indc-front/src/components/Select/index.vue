<template>
  <el-select
    :modelValue="modelValue"
    class="zh-select"
    :placeholder="placeholder"
    @change="onChange"
  >
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.label"
      :value="item.value"
    />
  </el-select>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, PropType } from 'vue'
export interface ISelect {
  label: string
  value: string
}
defineProps({
  modelValue: {
    type: String,
    defaut: '',
  },
  placeholder: {
    type: String,
    default: '请选择',
  },
  options: {
    type: Array as PropType<ISelect[]>,
    default: () => [],
  },
})
let emit = defineEmits(['update:modelValue'])
const onChange = (val: string) => {
  emit('update:modelValue', val)
}
</script>
<style lang="scss" scoped>
.zh-select {
  :deep(.el-input__inner) {
    width: 160px;
    height: 30px;
    border: 1px solid #b8b8d5;
    padding-left: 10px;
    &::-webkit-input-placeholder {
      color: #a1a1c3;
      font-size: 12px;
    }
    &::input-placeholder {
      color: #a1a1c3;
      font-size: 12px;
    }
  }
  :deep(.el-input__suffix) {
    right: 10px;
  }
}
</style>
