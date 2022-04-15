<template>
  <div class="test-search">
    <TestSearch :queryItems="queryItems" @onSearch="onSearch"> </TestSearch>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import TestSearch, { ISearch } from '@/components/Search/index.vue'
let queryItems = ref<ISearch>({
  gutter: 0,
  columnList: [
    { type: 'input', label: '请输入告警类型', value: '' },
    {
      type: 'select',
      label: '告警级别',
      value: '',
      options: [{ label: 'lablel1', value: 'value1' }],
    },
    {
      type: 'select',
      label: '位置',
      value: '',
      options: [{ label: 'lablel1', value: 'value1' }],
    },
    {
      type: 'select',
      label: '设备类型',
      value: '',
      options: [{ label: 'lablel1', value: 'value1' }],
    },
    { type: 'input', label: '告警名称', value: '' },
    // { type: 'input', label: '告警名称', value: '' },
    // { type: 'input', label: '告警名称', value: '' },
    // { type: 'input', label: '最后一个', value: '' },
  ],
})
const onSearch = () => {
  console.log(queryItems.value.columnList)
}
</script>
<style lang="scss" scoped>
.test-search {
  height: 100vh;
}
</style>
