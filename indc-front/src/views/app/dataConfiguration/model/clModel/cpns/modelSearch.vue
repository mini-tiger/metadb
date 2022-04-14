<template>
  <TestSearch :queryItems="queryItems" @onSearch="onSearch"></TestSearch>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

import TestSearch from '@/components/Search/index.vue'
// import publicFrom from '@/components/Form/index.vue'
import { searchData } from '../config/cl-model'
export default defineComponent({
  components: {
    // publicFrom,
    TestSearch,
  },
  setup() {

    const queryItems = ref<any>(searchData)
    const onSearch = () => {
      console.log('搜索--->')
    }
    return {
      queryItems,
      onSearch,
    }
  },
})
</script>
