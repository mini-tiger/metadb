<template>
  <div class="table-config">
    <slot name="table-config-left">
      <div class="table-config-left">
        <ZHButton
          v-for="(item, idx) of leftConfig"
          :Key="idx"
          :options="item.options"
          @ZHBtnClick="item.handler"
          >{{ item.name }}</ZHButton
        >
      </div>
    </slot>

    <div class="table-config-right">
      <slot name="special-config"></slot>
      <slot name="table-config-right">
        <ZHButton
          v-for="(item, idx) of rightConfig"
          :Key="idx"
          :options="item.options"
          @ZHBtnClick="item.handler"
        >
          <svg-icon icon-class="user" v-if="item.icon"></svg-icon
          >{{ item.name }}
          <ul v-if="item.showColumn?.length" class="column-item">
            <!-- <li>下拉列表</li>
            <li>下拉列表</li>
            <li>下拉列表</li>
            <li>下拉列表</li>
            <li>下拉列表</li> -->
          </ul>
        </ZHButton>
      </slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { PropType, defineProps, ref } from 'vue'
import { IButton } from '../Button/Button.vue'
export interface ILeftConfig {
  name: string
  options: IButton
  handler: () => void
}
export interface IShowColumn {
  label: string
  value: string
}
export interface IRightConfig extends ILeftConfig {
  icon?: boolean
  showColumn?: IShowColumn[]
}
let props = defineProps({
  leftConfig: {
    type: Array as PropType<ILeftConfig[]>,
    default: () => [],
  },
  rightConfig: {
    type: Array as PropType<IRightConfig[]>,
    default: () => [],
  },
})
</script>
<style lang="scss" scoped>
.table-config {
  display: flex;
  height: 50px;
  align-items: center;
  justify-content: space-between;
  padding: 10px;
  background-color: #fff;
  &-left {
    display: flex;
    align-items: center;
    justify-content: flex-start;
  }
  &-right {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    .column-item {
      position: absolute;
      z-index: 10;
      background-color: #afa;
    }
  }
}
</style>
