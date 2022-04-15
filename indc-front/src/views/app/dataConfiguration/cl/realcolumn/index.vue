<template>
  <div class="clExample">
    <Breadcrum></Breadcrum>
    <div class="publicButton">
      <ZHButton :options="optionExcel" @ZHBtnClick="onZHBtnClick"
        >导入</ZHButton
      >
      <ZHButton :options="optionExcel" @ZHBtnClick="onZHBtnClick"
        >导出</ZHButton
      >
      <ZHButton :options="options3" @ZHBtnClick="onButton">
        <svg-icon icon-class="createPlus" class="create-plus"></svg-icon
        ><span>新建实例</span>
      </ZHButton>
    </div>
    <RealSearch />
    <RealTable />
    <DrawPopup :data="state.drawer" :key="state.random" @IsDrawer="IsDrawer" />
    <UpLoadDrawer
      :isShowUpLoad="isShowUpLoad"
      @btnCancel="btnCancel"
      @btnSure="btnSure"
    />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue'

import UpLoadDrawer from '@/components/UpLoadDrawer/index.vue'
import ZHButton from '@/components/Button/Button.vue'
import DrawPopup from './cpns/DrawPopup.vue'
import RealSearch from './cpns/RealSearch.vue'
import RealTable from './cpns/RealTable.vue'
import { ButtonRound } from './config/real-config'

export default defineComponent({
  name: 'realColumn',
  components: {
    RealSearch,
    RealTable,
    DrawPopup,
    ZHButton,
    UpLoadDrawer,
  },
  setup() {
    const isShowUpLoad = ref(false)
    const state = reactive({
      buttonOk: ButtonRound,
      drawer: false,
      random: Math.random() * 100,
    })
    let options3 = {
      bgColor: '#595AEC',
      textColor: '#fff',
      width: '100px',
      height: '30px',
    }
    let optionExcel = {
      bgColor: '#E4E4FC',
      textColor: '#595AEC',
    }
    const onButton = () => {
      state.random = Math.random() * 100
      state.drawer = true
    }
    const IsDrawer = () => {
      state.random = Math.random() * 100
      state.drawer = false
    }
    const onZHBtnClick = () => {
      console.log('导入')
      isShowUpLoad.value = true
    }
    const btnCancel = () => {
      console.log('取消了')
      isShowUpLoad.value = false
    }
    const btnSure = () => {
      console.log('确定了---》')
      isShowUpLoad.value = false
    }
    return {
      state,
      options3,
      optionExcel,
      isShowUpLoad,
      btnCancel,
      btnSure,
      onZHBtnClick,
      onButton,
      IsDrawer,
    }
  },
})
</script>

<style scoped>
.clExample {
  position: relative;
}
.publicButton {
  position: absolute;
  right: 0;
  top: 10px;
}
</style>
