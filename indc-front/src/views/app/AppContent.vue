<template>
  <div class="container">
    <NavMenu
      class="nav-menu"
      :rootMenu="rootMenu"
      v-if="rootMenu.length > 0"
    ></NavMenu>

    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import NavMenu from '@/components/NavMenu.vue'
import { useStore } from 'vuex'
export default defineComponent({
  name: 'AppContent',
  components: {
    NavMenu,
  },
  setup() {
    const router = useRouter()
    const store = useStore()
    // const leveList = ref([])
    let rootMenu = computed(() => store.state.app.rootMenu)
    const getBreadcrumb = () => {
      console.log(router, '我是路由')
    }
    watch(
      () => router.currentRoute.value.path,
      () => {
        getBreadcrumb()
      },
      { immediate: true },
    )
    return {
      rootMenu,
    }
  },
})
</script>
<style scoped lang="scss">
.container {
  display: flex;
  background-color: #f5f6fa;
  .nav-menu {
    display: inline-flex;
  }
  .content {
    padding-left: 20px;
    padding-right: 10px;
    flex: 1;
    background-color: #f5f6fa;
  }
}
</style>
