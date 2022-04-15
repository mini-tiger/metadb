<template>
  <el-date-picker
    v-model="newValue"
    :type="type"
    :range-separator="rangeSeparator"
    :start-placeholder="startPlaceholder"
    :end-placeholder="endPlaceholder"
    @change="onChange"
    :value-format="valueFormat"
    class="zh-date-picker"
    prefix-icon="none"
    :format="format"
  >
  </el-date-picker>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, ref, watch } from 'vue'
let props = defineProps({
  type: {
    type: String,
    default: 'daterange',
  },
  format: {
    type: String,
    default: 'YYYY/MM/DD',
  },
  rangeSeparator: {
    type: String,
    default: '至',
  },
  startPlaceholder: {
    type: String,
    default: '开始日期',
  },
  endPlaceholder: {
    type: String,
    default: '结束日期',
  },
  modelValue: {
    type: [Array, String],
    default: () => [],
  },
  valueFormat: {
    type: String,
    default: 'YYYY-MM-DD',
  },
})
let emit = defineEmits(['update:modelValue'])
let newValue = ref(props.modelValue)
const onChange = (val) => {
  emit('update:modelValue', val)
}

watch(
  () => props.modelValue,
  (val) => {
    newValue.value = val
  },
)
</script>
<style lang="scss" scoped></style>
