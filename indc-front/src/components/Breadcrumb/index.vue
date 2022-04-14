<template>
  <el-breadcrumb class="breadcrumb" separator="/">
    <el-breadcrumb-item
      v-for="(item, index) in breadcrumbData"
      :key="item.path"
    >
      <!-- 不可点击项 -->
      <span v-if="index === breadcrumbData.length - 1" class="no-redirect">{{
        item.meta.title
      }}</span>
      <!-- 可点击项 -->
      <a v-else class="redirect" @click.prevent="onLinkClick(item)">{{
        item.meta.title
      }}</a>
    </el-breadcrumb-item>
  </el-breadcrumb>
</template>

<script setup lang="ts">
import { useRouter, useRoute } from 'vue-router'
import { ref, watch } from 'vue'
const router = useRouter()
const route = useRoute()
const breadcrumbData = ref([])
const getBreadcrumbData = () => {
  breadcrumbData.value = route.matched.filter(
    (item) => item.meta && item.meta.title,
  )
}
// 监听路由变化时触发
watch(
  route,
  () => {
    getBreadcrumbData()
  },
  {
    immediate: true,
  },
)
// 处理点击事件
const onLinkClick = (item) => {
  router.push(item.path)
}
</script>

<style lang="scss" scoped>
@import '@/assets/style/var.scss';
:deep .el-breadcrumb__separator {
  font-weight: 300;
}
.breadcrumb {
  width: 100%;
  height: 50px;
  line-height: 45px;
  .redirect {
    font-size: 12px;
    color: #5c5c8d;
    font-weight: 400;
  }
  .no-redirect {
    color: $main-color;
    font-size: 12px !important;
    font-weight: 600;
  }
}
</style>
