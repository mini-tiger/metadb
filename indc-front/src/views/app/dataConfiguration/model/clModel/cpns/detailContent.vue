<template>
  <div class="tabs">
    <el-tabs v-model="data.activeName" @tab-click="data.handleClick">
      <el-tab-pane label="模型字段" name="first">
        <modelFiled />
      </el-tab-pane>
      <el-tab-pane label="测点字段" name="second">
        <MeasuringPoint />
      </el-tab-pane>
      <el-tab-pane label="模型关联" name="third">
        <modelAssociationList v-if="data.type === 'list'" />
        <modelAssociationTopology v-if="data.type === 'topo'" />
      </el-tab-pane>
    </el-tabs>
    <div class="button-right">
      <publicButton
        :data="buttonData.buttonNew"
        v-if="data.tabIndex == 2 && data.type === 'list'&&type!=='look'"
        @click="data.newAssociation"
      ></publicButton>
      <publicButton
        :data="buttonData.buttonInPart"
        @onButton="upLoad"
        v-if="type!=='look'"
      ></publicButton>
      <publicButton
        :data="buttonData.buttonOutPart"
        @onButton="downLoad"
        v-if="type!=='look'"
      ></publicButton>
      <div class="list-topology" v-if="data.tabIndex == 2">
        <div
          class="list"
          :class="[data.type === 'list' ? 'is-active' : '']"
          @click="data.ListOrTopo('list')"
        >
          列表
        </div>
        <div
          class="topology"
          :class="[data.type === 'topo' ? 'is-active' : '']"
          @click="data.ListOrTopo('topo')"
        >
          拓扑
        </div>
      </div>
    </div>
  </div>

  <UpLoadDrawer
    :isShowUpLoad="isShowUpLoad"
    @btnCancel="btnCancel"
    @btnSure="btnSure"
  />
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue'
import publicButton from '@/components/Button/index.vue'
import { buttonData, uploadData } from '../config/model-detail'
import modelAssociationList from './modelAssociationList.vue'
import modelAssociationTopology from './modelAssociationTopology.vue'
import modelFiled from './modelFiled.vue'
import MeasuringPoint from './measuringPoint.vue'
import UpLoadDrawer from '@/components/UpLoadDrawer/index.vue'

import { drawerData } from '../config/model-association'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'detailContent',
  components: {
    publicButton,
    modelAssociationList,
    modelAssociationTopology,
    modelFiled,
    MeasuringPoint,
    UpLoadDrawer,
  },
  setup() {
    const data = reactive({
      activeName: 'first',
      tabIndex: 0,
      type: 'list',
      handleClick(tabs) {
        console.log(tabs.index)
        data.tabIndex = tabs.index
      },
      ListOrTopo(type) {
        data.type = type
      },
      newAssociation() {
        drawerData.isShowDrawer = true
      },
    })
    const route = useRoute()

    if (route.params && !localStorage.getItem('rowData')) {
      const data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    const type = JSON.parse(localStorage.getItem('rowData')).type || ''
    return {
      data,
      buttonData,
      drawerData,
      ...toRefs(uploadData),
      type
    }
  },
})
</script>

<style scoped lang="scss">
@import '../../../../../../assets/style/overallStyle.scss';
.tabs {
  .list-topology {
    width: 100px;
    height: 26px;
    border-radius: 8px;
    border: 1px solid #b8b8d5;
    display: flex;
    font-size: 14px;
    color: #a1a1c3;
    margin-left: 10px;
    .list {
      width: 50px;
      //height: 102%;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 7px 0px 0px 7px;
    }
    .topology {
      width: 50px;
      //height: 102%;
      display: flex;
      align-items: center;
      justify-content: center;
      border-radius: 0 7px 7px 0;
    }
    .is-active {
      color: #595aec;
      background: #e4e4fc;
      height: 102%;
    }
  }
}
</style>
