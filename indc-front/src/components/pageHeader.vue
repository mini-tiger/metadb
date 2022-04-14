<template>
  <div class="header-container">
    <div class="logo-container">
      <!-- <span title="logo-title">智航</span> -->
    </div>
    <!-- <ThemeSelect></ThemeSelect> -->
    <!-- <Langage></Langage> -->
    <el-menu
      class="nav-menu"
      mode="horizontal"
      :router="true"
      default-active="/"
    >
      <el-menu-item
        v-for="(nav, idx) in navList"
        :index="nav.path"
        :key="idx"
        >{{ nav.name }}</el-menu-item
      >
    </el-menu>
    <el-menu class="setting-menu" mode="horizontal">
      <svg-icon icon-class="setup" class="setup-header" @click="settings" />
      <!-- <div class="news" :data-attr="content" @click="MessageNotification"> -->

      <!-- <i :data-attr="content" class="news">
        <svg-icon
          icon-class="news"
          class-name="news-header"
          @click="MessageNotification"
        >
        </svg-icon>
      </i> -->

      <el-badge value="9" class="item">
        <svg-icon
          icon-class="news"
          class-name="news-header"
          @click="MessageNotification"
        />
      </el-badge>
      <!-- <el-avatar :src="state.urlData" class="avatar"></el-avatar> -->
      <img
        src="../assets/icons/admin.png"
        alt=""
        width="24"
        height="24"
        class="img-icon"
      />
      <el-sub-menu index="profile">
        <template v-slot:title>{{ username }}</template>
        <el-menu-item index="logout" @click="handleLogout"
          >退出登录</el-menu-item
        >
      </el-sub-menu>
    </el-menu>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'
// import ThemeSelect from '@/components/themeSelect/index.vue'
// import Langage from '@/components/Language/index.vue'
export default defineComponent({
  name: 'PageHeader',
  components: {
    // ThemeSelect,
    // Langage,
  },
  setup() {
    const content = ref('99')
    // const state = reactive({
    //   urlData: require('../assets/icons/admin'),
    // })
    const store = useStore()
    const user = store.state.user
    const bgColor = computed(() => store.state.theme.themeColor)
    const username = computed(() => user.userInfo.username || '张三')
    // const hasPermission = computed(() => user.hasPermission)
    const navList = computed(() => user.navList)
    const handleLogout = async () => {
      await store.dispatch(`user/${types.USER_LOGOUT}`)
    }
    const settings = () => {
      console.log('我是设置')
    }
    const MessageNotification = () => {
      console.log('我是消息通知')
    }
    return {
      bgColor,
      username,
      // hasPermission,
      handleLogout,
      navList,
      content,
      settings,
      MessageNotification,
    }
  },
})
</script>
<style lang="scss" scoped>
@import '@/assets/style/var.scss';
.header-container {
  display: flex;
  position: absolute;
  width: 100%;
  height: 50px;
  line-height: 50px;
  z-index: 1;
  box-shadow: 3px 3px 6px 0px rgba(0, 0, 0, 0.09);
  .logo-container {
    width: 260px;
    // height: 60px;
    // line-height: 60px;
    // padding-left: 22px;
    background: url('../assets/icons/bg.png') no-repeat;
    background-size: 160px 50px;
    & > span {
      margin-left: 42px;
      color: #302e6d;
    }
  }
  .nav-menu {
    flex: 1;
    justify-content: flex-start;
    border: none;
    box-shadow: 0px 6px 9px 0px rgba(0, 0, 0, 0.09);
    user-select: none;
  }
  .setting-menu {
    display: flex;
    // position: relative;
    width: 253px;
    justify-content: flex-end;
    align-items: center;
  }

  .el-menu-item {
    width: 120px;
    font-weight: 600;
    box-sizing: border-box;
    color: $base-color;
    position: relative;
  }
  .el-menu--horizontal > .el-menu-item.is-active {
    position: relative;
    //border-bottom: 2px solid transparent;
    border: none !important;
    color: $main-color !important;
    //background-color: transparent;
  }
  .el-menu--horizontal > .el-menu-item {
    border: none !important;
  }
  .el-menu--horizontal > .el-menu-item.is-active::after {
    content: '';
    display: block;
    position: absolute;
    width: 60px;
    height: 2px;
    background: $main-color;
    border-radius: 2px;
    bottom: 5px;
  }
}

:deep(.el-menu) {
  box-shadow: none !important;
  .el-menu-item {
    // margin-right: 15px;
    // height: 50px;
    // width: ;
  }
  .el-menu-item:hover {
    color: #595aec;
  }
}
:deep(.el-sub-menu__title) {
  padding-left: 10px;
  color: #302e6d !important;
}

.el-menu--horizontal {
  border: none;
}
.avatar {
  width: 36px;
  height: 36px;
  margin-top: 12px;
}
.setup-header {
  width: 20px;
  height: 20px;
  margin-right: 33px;
  cursor: pointer;
}

.news-header {
  position: relative;
  display: inline-block;
  // position: relative;
  width: 20px;
  height: 20px;
  margin-top: 20px;
  margin-right: 33px;
  cursor: pointer;
}
.news::after {
  content: attr(data-attr);
  // content: '123';
  display: inline-block;
  width: 16px;
  height: 16px;
  font-size: 12px;
  background: #df5347;
  color: #fff;
  text-align: center;
  line-height: 16px;
  border-radius: 50%;
  position: absolute;
  left: 109px;
  top: 13px;
}
:deep(.el-badge__content) {
  top: 20px;
  right: 41px;
}

:deep(.el-badge__content--danger) {
  background: #df5347 !important;
}
</style>
