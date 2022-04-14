<template>
  <el-switch
    v-model="data"
    :width="width"
    :active-color="activeColor"
    :inactive-color="inactiveColor"
    :inactive-text="inactiveText"
    class="switch-wrapper"
    @change="onChange"
  >
  </el-switch>
</template>

<script lang="ts" setup>
import { ref, defineProps, defineEmits } from 'vue'
defineProps({
  width: {
    type: Number,
    default: 50,
  },
  activeColor: {
    type: String,
    default: '#13ce66',
  },
  inactiveColor: {
    type: String,
    default: '#ff4949',
  },
  activeText: {
    type: String,
    default: '自动刷新',
  },
  inactiveText: {
    type: String,
    default: '自动刷新:',
  },
})
let emit = defineEmits(['update:modelValue'])
let data = ref(false)
const onChange = () => {
  emit('update:modelValue', data.value)
}
</script>
<style lang="scss" scoped>
.switch-wrapper {
  user-select: none;
  margin-right: 20px;
}
:deep(.el-switch__core) {
  height: 24px;
  border-radius: 12px;
  .el-switch__action {
    width: 18px;
    height: 18px;
    margin-top: 1px;
  }
}
.el-switch {
  --el-switch-button-size: 20px;
}
:deep(.el-switch__label.is-active) {
  color: #302e6d;
  font-weight: 500;
}
</style>
