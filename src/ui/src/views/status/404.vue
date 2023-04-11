<template>
  <bk-exception class="exception-wrap-item" type="404">
    <p>{{$t('页面不存在')}}</p>
    <div class="text-wrap">
      <bk-link theme="primary" @click="handleRedirect">{{$t('返回业务')}}</bk-link>
    </div>
  </bk-exception>
</template>

<script>
  import { MENU_BUSINESS } from '@/dictionary/menu-symbol'
  export default {
    methods: {
      handleRedirect() {
        this.$routerActions.redirect({
          name: MENU_BUSINESS
        })
      }
    }
  }
</script>

<style lang="scss" scoped>
    .exception-wrap-item {
        justify-content: center;
    }
</style>
