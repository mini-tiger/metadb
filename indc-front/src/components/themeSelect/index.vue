<template>
  <!-- 主题图标  
  v-bind：https://v3.cn.vuejs.org/api/instance-properties.html#attrs -->
  <el-dropdown
    v-bind="$attrs"
    trigger="click"
    class="theme"
    @command="handleSetTheme"
  >
    <div>
      <el-tooltip content="换肤">
        <svg-icon icon-class="user" />
      </el-tooltip>
    </div>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item command="color"> 换肤 </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
  <!-- 展示弹出层 -->
  <div>
    <ThemePicker v-model="selectColorVisible"></ThemePicker>
  </div>
</template>

<script lang="ts" setup>
import ThemePicker from '@/components/ThemePicker/index.vue'
import { ref } from 'vue'

const selectColorVisible = ref(false)
const handleSetTheme = () => {
  selectColorVisible.value = true
}
</script>

<style lang="scss" scoped></style>
