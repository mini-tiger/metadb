<template>
  <div class="cl-relation">
    <Breadcrumb />
    <div style="height: 120px; margin-bottom: 10px">
      <TestSearch :queryItems="queryItems" @onSearch="onSearch"></TestSearch>
    </div>
    <publicButton
      :data="buttonRelation.buttonOk"
      class="publicButton"
      @click="emitTopology"
    />
    <div id="canvas-custom-dnd">
      <div class="canvas">
        <MyCanvas ref="canvasRef" class="canvas-wrapper"></MyCanvas>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import MyCanvas from '@/components/Canvas/index.vue'
// import publicFrom from '@/components/Form/index.vue'
import TestSearch from '@/components/Search/index.vue'
import publicButton from '@/components/Button/index.vue'
import { searchData, buttonConfig } from './config/relation-config'

export default defineComponent({
  components: {
    // publicFrom,
    publicButton,
    TestSearch,
    MyCanvas,
  },
  setup() {
    const router = useRouter()
    const buttonRelation = ref(buttonConfig)
    const dialog = ref(false)
    const canvasRef = ref()
    const queryItems = ref<any>(searchData)
    onMounted(() => {
      console.log(canvasRef, '-----><')
    })
    const onSearch = (e) => {
      console.log(e, '--->')
    }
    const emitTopology = () => {
      console.log('我是编辑')
      router.push({
        path: 'emitTopology',
      })
    }
    return {
      queryItems,
      buttonRelation,
      dialog,
      canvasRef,
      emitTopology,
      onSearch,
    }
  },
})
</script>

<style scoped lang="scss">
.cl-relation {
  position: relative;
  // height: 100%;
}
.headerH {
  height: 120px;
  margin-bottom: 10px;
}
.canvas {
  width: 100%;
  height: 100%;
}
.publicButton {
  position: absolute;
  right: 0;
  top: 10px;
}
:deep(.fromSlot .el-form-item__content) {
  position: absolute;
  right: 0;
  bottom: 0;
}
</style>
