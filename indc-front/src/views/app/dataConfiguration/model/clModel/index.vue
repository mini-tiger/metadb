<template>
  <div class="model">
    <Breadcrumb />
    <modelSearch style="margin-bottom: 18px" />
    <modelTable style="height: 100%; margin-bottom: 31px" />
    <publicButton
      :data="modelNewButton.buttonOk"
      class="newModel"
      @onButton="onButton"
    />
    <newModel
      v-if="isNewModel"
      :isNewModel="isNewModel"
      @ModelOk="ModelOk"
      @cancelModel="cancelModel"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

import * as types from '@/store/action-types'
import publicButton from '@/components/Button/index.vue'
import modelTable from './cpns/modelTable.vue'
import modelSearch from './cpns/modelSearch.vue'
import newModel from './cpns/newModel.vue'
import { newModelConfig } from './config/cl-model'

export default defineComponent({
  components: {
    publicButton,
    modelSearch,
    modelTable,
    newModel,
  },
  setup() {
    const store = useStore()
    const router = useRouter()
    const modelNewButton = ref(newModelConfig)
    const isDraw = ref(false)
    const isNewModel = ref(false)
    const onButton = () => {
      isNewModel.value = true
    }
    // 确定新建模型
    const ModelOk = (data) => {
      isNewModel.value = false
      const data1 = {
        type: 'add',
        params: data,
      }
      const params = JSON.stringify(data1)
      router.push({
        name: 'modelDetail',
        params: { data: params },
      })
    }

    // 取消模型
    const cancelModel = () => {
      isNewModel.value = false
    }
    // 选中获取的数据
    const selectData = (name) => {
      console.log(name)
    }

    return {
      modelNewButton,
      isDraw,
      isNewModel,
      onButton,
      ModelOk,
      cancelModel,
      selectData,
    }
  },
})
</script>

<style scoped lang="scss">
.model {
  position: relative;
}
.newModel {
  position: absolute;
  right: 0;
  top: 10px;
}
</style>
