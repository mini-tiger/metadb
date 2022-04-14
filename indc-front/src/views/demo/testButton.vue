<template>
  <div class="wrapper">
    <ZHButton :options="options" @ZHBtnClick="onZHBtnClick">查询</ZHButton>
    <ZHButton :options="options1" @ZHBtnClick="onZHBtnClick1">弹框</ZHButton>
    <ZHButton :options="options2" @ZHBtnClick="onZHBtnClick2">导入</ZHButton>
    <ZHButton :options="options3" @ZHBtnClick="onZHBtnClick3">
      <svg-icon icon-class="createPlus" class="create-plus"></svg-icon
      ><span>新建模型</span>
    </ZHButton>
    <ZHButton :options="options4" @ZHBtnClick="onZHBtnClick4"
      >下载模板</ZHButton
    >
    <ZHButton :options="options5" @ZHBtnClick="onZHBtnClick5">
      <svg-icon icon-class="createPlus" class="create-plus"></svg-icon
      ><span>编辑模型关系</span>
    </ZHButton>
    <ZHButton :options="options6" @ZHBtnClick="onZHBtnClick6">
      <svg-icon icon-class="createPlus" class="create-plus"></svg-icon
      ><span>新建根目录</span>
    </ZHButton>
  </div>
</template>

<script lang="ts" setup>
import ZHButton, { IButton } from '@/components/Button/Button.vue'

let options: IButton = {
  bgColor: '#E4E4FC',
  textColor: '#595AEC',
}
let options1: IButton = {
  bgColor: '#595AEC',
  textColor: '#fff',
}
let options2: IButton = {
  bgColor: '#fff',
  textColor: '#595AEC',
  width: '60px',
  height: '26px',
}
let options3: IButton = {
  bgColor: '#595AEC',
  textColor: '#fff',
  width: '100px',
  height: '30px',
}
let options4: IButton = {
  bgColor: '#fff',
  textColor: '#595AEC',
  width: '80px',
  height: '26px',
  borderRadius: '13px',
}
let options5: IButton = {
  bgColor: '#595AEC',
  textColor: '#fff',
  width: '120px',
  height: '30px',
}
let options6: IButton = {
  bgColor: '#595AEC',
  textColor: '#fff',
  width: '110px',
  height: '30px',
}
const onZHBtnClick = () => {
  console.log('逻辑触发')
}
const onZHBtnClick1 = () => {
  console.log('逻辑触发1')
}
const onZHBtnClick2 = () => {
  console.log('逻辑触发2')
}
const onZHBtnClick3 = () => {
  console.log('逻辑触发3')
}
const onZHBtnClick4 = () => {
  console.log('逻辑触发4')
}
const onZHBtnClick5 = () => {
  console.log('逻辑触发5')
}
const onZHBtnClick6 = () => {
  console.log('逻辑触发6')
}
</script>
<style lang="scss" scoped>
.wrapper {
  display: flex;
  height: 100vh;
  justify-content: center;
  align-items: center;
}
</style>
