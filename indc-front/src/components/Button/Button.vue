<template>
  <el-button :disabled="disabled" :color="opts.bgColor" :style="opts" @click.stop="onClick">
    <slot>查询</slot>
  </el-button>
</template>

<script lang="ts" setup>
import style from '@/assets/style/var.scss'
import { defineProps, defineEmits, PropType, computed } from 'vue'
export interface IButton {
  bgColor: string
  textColor: string
  width?: string
  height?: string
  borderRadius?: string
  border?: string
}
let props = defineProps({
  options: {
    type: Object as PropType<IButton>,
    default: () => ({}),
  },
  disabled: {
    type: Boolean,
    default: false,
  }
})
let opts = computed(() => ({
  color: props.options.textColor || style.btnTextColor,
  width: props.options.width || '70px',
  height: props.options.height || '30px',
  borderRadius: props.options.borderRadius || '15px',
  border: props.options.border || `1px solid ${style.mainColor}`,
  bgColor: props.options.bgColor || style.mainColor,
}))
let emit = defineEmits(['ZHBtnClick'])
const onClick = () => {
  emit('ZHBtnClick')
}
</script>
<style lang="scss" scoped>
:deep(.create-plus) {
  margin-right: 6px;
}
</style>
