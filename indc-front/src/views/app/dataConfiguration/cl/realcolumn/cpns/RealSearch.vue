<template>
  <div class="formData">
    <!-- <publicForm :data="state" class="div-box-form">
      <template v-slot="data">
        <div class="searchData">
          <el-button @click="onSubmit(data.data.formRef, data.data.formData)"
            >搜索</el-button
          >
          <el-button @click="resetFrom(data.data.formRef, data.data.formData)"
            >重置</el-button
          >
        </div>
      </template>
    </publicForm> -->
    <TestSearch :queryItems="queryItems" @onSearch="onSearch"></TestSearch>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'

// import publicForm from '@/components/Form/index.vue'
import TestSearch from '@/components/Search/index.vue'
import { searchData } from '../config/real-config'
export default defineComponent({
  name: 'realSearch',
  components: {
    TestSearch,
  },
  setup() {
    // const state = reactive({
    //   queryItems: searchData,
    // })
    const queryItems = ref(searchData)
    const onSearch = () => {
      console.log(queryItems.value.columnList)
    }
    return {
      queryItems,
      onSearch,
    }
  },
})
</script>

<style scoped lang="scss">
//@import '@/assets/style/index.scss';
.formData {
  position: relative;
}
:deep(.el-button:hover) {
  background: #e4e4fc;
  color: red;
  border: 1px solid #595aec;
}
.searchData {
  position: absolute;
  right: 0;
  top: 60px;
}

:deep(.fromSlot .el-form-item__content) {
  position: static !important;
}
</style>
