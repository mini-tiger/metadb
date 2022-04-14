<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <el-menu
          style="height: 100vh"
          :default-openeds="['/demo']"
          :default-active="defaultActive"
          active-text-color="#ffd04b"
          background-color="#545c64"
          text-color="#fff"
          router
        >
          <el-sub-menu index="/demo">
            <template #title>
              <span>组件</span>
            </template>
            <el-menu-item index="/demo/search">查询</el-menu-item>
            <el-menu-item index="/demo/button">按钮</el-menu-item>
            <el-menu-item index="/demo/canvas">画布</el-menu-item>
            <el-menu-item index="/demo/table">表格</el-menu-item>
          </el-sub-menu>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
let route = useRoute()
let path = sessionStorage.getItem('menuActive') || '/demo/search'
let defaultActive = ref(path)
watch(
  () => route.path,
  (path) => {
    sessionStorage.setItem('menuActive', path)
  },
)
</script>
<style lang="scss" scoped></style>
