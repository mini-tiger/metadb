<template>
  <div>
    <template v-if="!menu.subMenus?.length">
      <el-menu-item :index="menu?.route">
        {{ menu.name }}
      </el-menu-item>
    </template>
    <template v-else>
      <el-sub-menu :index="menu.id.toString()">
        <template #title>
          <!-- <el-icon class="secondIcon"><location /></el-icon> -->
          <svg-icon icon-class="item-icon" class="menu-icon"></svg-icon>
          {{ menu.name }}
        </template>
        <template v-for="child of menu.subMenus" :key="child.id">
          <MenuTree :menu="child" />
        </template>
      </el-sub-menu>
    </template>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
// import { Location } from '@element-plus/icons-vue'
export default defineComponent({
  name: 'MenuTree',
  props: {
    menu: Object,
  },
  components: {
    // Location,
  },
})
</script>

<style scoped lang="scss">
:deep(.el-sub-menu__title) {
  font-size: 14px !important;
  height: 50px;
  // padding: 0 18px 0 10px !important;
  // font-family: PingFangSC-Medium, PingFang SC;
  font-family: PuHuiTiRegular;
}

:deep(.el-menu-item) {
  font-size: 12px !important;
  padding-left: 60px !important;
  padding-right: 0;
  height: 40px;
  line-height: 40px;
}

:deep(.el-menu-item:hover) {
  background: #ededfd !important;
}

.secondIcon {
  width: 18px;
  height: 18px;
}
.menu-icon {
  width: 24px;
  height: 24px;
  margin-right: 10px;
}
</style>
