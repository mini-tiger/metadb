<template>
  <div class="home">
    <el-container class="platform-container">
      <el-header>
        <PageHeader></PageHeader>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import PageHeader from '@/components/pageHeader.vue'
export default defineComponent({
  name: 'Home',
  components: {
    PageHeader,
  },
  setup() {
    return {}
  },
})
</script>
<style lang="scss" scoped>
.el-header {
  --el-header-padding: 0;
  height: 50px;
  line-height: 50px;
}
</style>
