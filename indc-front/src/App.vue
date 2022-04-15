<template>
  <el-config-provider :locale="store.getters.language === 'en' ? en : zh">
    <router-view />
  </el-config-provider>
</template>
<script lang="ts" setup>
import zh from 'element-plus/lib/locale/lang/zh-cn'
import en from 'element-plus/lib/locale/lang/en'
import { useStore } from 'vuex'
import { useI18n } from 'vue-i18n'
const i18n = useI18n()
const message = i18n.t('msg.test')
//i18n.global.t('msg,test') 在ts中的用法
console.log('message', message)

// import { generateNewStyle, writeNewStyle } from '@/utils/theme'
// import { computed } from 'vue'
const store = useStore()
// const mainColor = computed(() => store.state.theme.themeColor)
// generateNewStyle(mainColor.value).then((newStyleText) => {
//   writeNewStyle(newStyleText)
// })
</script>
<style lang="scss">
@import '@/assets/style/var.scss';
html,
body,
#app,
.home,
.el-container,
.el-main,
.container,
.nav-menu {
  min-height: 100%;
}
.el-sub-menu .el-menu-item {
  min-width: 180px !important;
}
.zh-date-picker.el-range-editor.el-input__inner {
  width: 180px;
  height: 30px;
  line-height: 30px;
  border-radius: 4px;
  border: 1px solid #b8b8d5;
  padding: 3px 4px;
}
.el-range-editor .el-range-input,
.el-date-editor .el-range-separator {
  color: #a1a1c3 !important;
}
.el-date-editor .el-range__icon {
  // position: absolute;
  // right: 10px;
  display: inline-block;
  background: url('./assets/icons/png/calendar.png') no-repeat center;
  background-size: 14px 14px;
}
.el-main {
  background-color: #f5f6fa;
}
:root {
  --el-color-primary: #595aec;
}
</style>
