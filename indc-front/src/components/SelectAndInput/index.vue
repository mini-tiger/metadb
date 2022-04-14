<template>
  <div>
    <el-select
      style="width: 90px"
      class="zh-select"
      :modelValue="selectValue"
      @change="onChange"
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-input
      style="width: 180px"
      :modelValue="ValueInput"
      :placeholder="placeholder"
      @input="input"
    />
  </div>
</template>

<script lang="ts" setup>
import { defineProps, defineEmits, PropType, ref } from 'vue'
export interface ISelect {
  label: string
  value: string
}
defineProps({
  selectValue: {
    type: Number,
    default: null,
  },
  modelValue: {
    type: String,
    default: '',
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

const ValueInput = ref('')
const ValueSelect = ref(null)
const emit = defineEmits(['update:selectValue', 'update:modelValue'])
const onChange = (e) => {
  console.log(e, '---->')
  const data = e * 1
  ValueSelect.value = data
  emit('update:selectValue', e)
}
const input = (e) => {
  ValueInput.value = e
  emit('update:modelValue', e)
  console.log(ValueInput.value)
}
</script>
