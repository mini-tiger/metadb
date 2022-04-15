<template>
  <div class="modelRelation">
    <Breadcrumb />
    <!-- <div class="div-box-form model-search">
      <publicFrom :data="modelSearch">
        <template v-slot="data">
          <el-button @click="searchFrom(data.data.formRef, data.data.formData)"
            >搜索</el-button
          >
          <el-button @click="resetFrom(data.data.formRef, data.data.formData)"
            >重置</el-button
          >
        </template>
      </publicFrom>
    </div> -->
    <div style="height: 120px; margin-bottom: 10px">
      <TestSearch :queryItems="queryItems" @onSearch="onSearch"></TestSearch>
      <publicButton
        :data="modelEdit.buttonOk"
        class="model-edit"
        @onButton="onButton"
      />
    </div>
    <div id="canvas-custom-dnd">
      <div class="canvas">
        <MyCanvas ref="canvasRef" class="canvas-wrapper"></MyCanvas>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useStore } from 'vuex'

import MyCanvas from '@/components/Canvas/index.vue'
import TestSearch from '@/components/Search/index.vue'
import { searchData, editTopologyConfig } from './config/model-relation'
// import publicFrom from '@/components/Form/index.vue'
import publicButton from '@/components/Button/index.vue'
import { GET_MODEL_RELATION } from '@/store/action-types'

export default defineComponent({
  components: {
    // publicFrom,
    TestSearch,
    publicButton,
    MyCanvas,
  },
  setup() {
    const queryItems = ref(searchData)
    const modelEdit = ref(editTopologyConfig)
    const store = useStore()
    const onSearch = () => {
      console.log(queryItems.value.columnList)
    }
    const onButton = () => {
      console.log('我是编辑')
    }
    onMounted(() => {
      // console.log(store.state.getAllObjTop)
      store.dispatch(`getAllObjTop/${GET_MODEL_RELATION}`)
    })
    return {
      queryItems,
      modelEdit,
      onSearch,
      onButton,
    }
  },
})
</script>

<style lang="scss" scoped>
.modelRelation {
  position: relative;
  height: 100%;
}
.model-edit {
  position: absolute;
  right: 0;
  top: 10px;
}
.model-search {
  height: 120px;
  position: relative;
}
:deep(.fromSlot .el-form-item__content) {
  position: absolute;
  right: 0;
  bottom: -30px;
}
.emitTopology {
  background: #fff;
  height: 100%;
  margin: 0 10px 50px 20px;
}
</style>
