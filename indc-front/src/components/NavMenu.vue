<template>
  <div>
    <ul class="nav-one-icon">
      <li
        v-for="(item, index) of rootMenu"
        :key="item.id"
        @click="handleMenu(item.id, index)"
        @mouseenter="handleActive(index)"
        @mouseleave="handleLeaveActive"
        v-bind:class="{
          isActive: index === isActive,
          isHoverActive: index === isHoverActive,
        }"
      >
        <div class="icon">
          <svg-icon :icon-class="imgArray[index]" class="svg-icon"></svg-icon>
        </div>
        <span class="icon-name">{{ item.name.slice(0, 2) }}</span>
      </li>
    </ul>
    <el-menu
      active-text-color="#595aec"
      background-color="#fff"
      class="el-menu-vertical"
      :class="{ hideMenu: !isCollapse }"
      text-color="#302E6D"
      :default-openeds="defaultOpeneds"
      :router="true"
      v-show="menuList.length > 0"
    >
      <template v-for="menu of menuList" :key="menu.id">
        <MenuTree :menu="menu" :index="menu.id" />
      </template>
    </el-menu>
    <div class="trigger-icon" @click="handleToggleMenu">
      <svg-icon
        :icon-class="isCollapse ? 'left-arrow' : 'right-arrow'"
        class="toggle-icon"
      ></svg-icon>
    </div>
<!--     <router-link to="/app-center/app-menu/wh">wh</router-link>-->
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useStore } from 'vuex'
import MenuTree from '@/components/MenuTree.vue'
// import { useRouter } from 'vue-router'
// import { getLocal } from '@/utils/common'
export default defineComponent({
  name: 'NavMenu',
  props: ['rootMenu'],
  components: {
    MenuTree,
  },
  setup(props) {
    const store = useStore()
    let menuList = ref([])
    let imgArray = ref([
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
      'nav1',
      'nav2',
      'nav3',
    ]) //临时一级菜单的svg
    let isCollapse = ref(false)
    let isActive = ref(null)
    // let defaultActive = ref() //默认激活的菜单
    let defaultOpeneds = ref([])
    const handleMenu = (id, index) => {
      let list = store.getters['app/getMenuById'](id)
      menuList.value = list
      isCollapse.value = true
      isActive.value = index
    }
    onMounted(() => {
      handleMenu(props.rootMenu[0]?.id, 0)
    })
    const handleActive = (index) => {
      console.log('index', index)

      isHoverActive.value = index
    }
    const isHoverActive = ref()
    const handleToggleMenu = (id, index) => {
      isCollapse.value = !isCollapse.value
      isActive.value = index
    }
    const handleLeaveActive = () => {
      isHoverActive.value = ''
    }
    return {
      menuList,
      defaultOpeneds,
      isCollapse,
      handleMenu,
      handleToggleMenu,
      handleActive,
      isActive,
      imgArray,
      isHoverActive,
      handleLeaveActive,
    }
  },
})
</script>
<style lang="scss" scoped>
@import '@/assets/style/var.scss';

.el-menu-vertical:not(.el-menu--collapse) {
  width: 180px;
  border-right: none;
}

.el-menu-vertical {
  display: inline-block;
  // transition: all 0.5s !important;
}
.hideMenu {
  transition: all 0.5s !important;
  width: 0 !important;
  overflow: hidden !important;
}
.trigger-icon {
  cursor: pointer;
  .toggle-icon {
    color: #595aec;
    display: inline-block;
    position: absolute;
    width: 24px;
    height: 24px;
    margin-left: -12px;
    margin-top: 12px;
  }
}

.nav-one-icon {
  display: inline-block;
  width: 50px;
  padding: 20px 0;
  text-align: center;
  background-color: $main-color;
  & > li {
    font-size: 12px;
    color: #fff;
    cursor: pointer;
    box-shadow: 0px 1px 4px 0px rgba(92, 107, 192, 0.2);
    opacity: 0.6;
  }
  .icon {
    width: 32px;
    height: 32px;
    line-height: 32px;
    margin: auto;
    border-radius: 50%;
    background-color: #fff;
    text-align: center;
    .svg-icon {
      font-size: 32px;
    }
  }
  .icon-name {
    display: inline-block;
    height: 20px;
    line-height: 20px;
    margin: 5px 0 20px;
    color: #fff;
    font-weight: 600;
    text-align: center;
  }
}
.el-menu {
  font-family: PuHuiTiRegular;
}
:deep(.is-opened .el-menu div .is-active) {
  font-weight: 500;
  color: #595aec;
  background-color: #ededfd;
}

.isActive,
.isHoverActive {
  opacity: 1 !important;
  box-shadow: none !important;
  .icon {
    background-color: #fff !important;
  }
}
</style>
