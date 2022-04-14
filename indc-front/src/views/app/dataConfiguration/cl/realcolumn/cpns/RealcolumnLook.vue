<template>
  <div class="tabs">
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
      <el-tab-pane label="属性字段" name="attribute">
        <div class="attribute-collapse">
          <div
            class="collapse"
            v-for="(item, index) in state.attributeList"
            :key="index"
          >
            <el-collapse accordion>
              <el-collapse-item
                :name="item.bk_group_id"
                :title="item.bk_group_name"
              >
                <div
                  class="collapse-list"
                  v-for="(itm, index) in item.data"
                  :key="index"
                >
                  <span class="collapse-span">{{
                    itm.bk_property_name + '  :  '
                  }}</span>
                  <span style="margin-right: 10px"></span>
                  <span class="collapse-span">{{ itm.bk_property_value }}</span>
                </div>
              </el-collapse-item>
            </el-collapse>
          </div>
        </div>
      </el-tab-pane>
      <el-tab-pane label="测点字段" name="measuring">
        <publicTable :data="measuringTable" />
      </el-tab-pane>
      <el-tab-pane label="cl关联" name="clRelation">
        <publicTable
          :data="clRelationTable"
          v-if="state.clRelationTabs === 'list'"
        >
          <template #modify="{ scope }">
            <span id="clRelationModify" @click="deleteData(scope)">删除</span>
          </template>
        </publicTable>
        <div v-if="state.clRelationTabs === 'topology'">123123</div>
      </el-tab-pane>
      <el-tab-pane label="表更记录" name="changeRecord">
        <publicTable :data="changeRecordTable" />
      </el-tab-pane>
    </el-tabs>
    <Public-button
      v-if="state.isFlag"
      :data="state.buttonOk"
      class="newButton"
      @onButton="onButton"
    ></Public-button>
    <div class="clRelationList" v-if="state.isFlag">
      <div
        class="clRelationTab"
        @click="clRelationTabs('list')"
        v-bind:class="{ clRelationTabActive: 'list' === state.clRelationTabs }"
      >
        列表
      </div>
      <div
        class="clRelationTab"
        @click="clRelationTabs('topology')"
        v-bind:class="{
          clRelationTabTopology: 'topology' === state.clRelationTabs,
        }"
      >
        拓扑
      </div>
    </div>

    <el-drawer v-model="isDrawer" title="新建关联" direction="rtl" :size="940">
      <div class="drawerStyle">
        <div class="headerTop">
          <publicForm :data="associationList.form">
            <template v-slot="data">
              <el-button
                @click="searchFrom(data.data.formRef, data.data.formData)"
                >搜索</el-button
              >
              <el-button
                @click="resetFrom(data.data.formRef, data.data.formData)"
                >重置</el-button
              >
            </template>
          </publicForm>
        </div>
        <div style="margin-top: 10px" class="drawerTable">
          <publicTable :data="associationTable">
            <template #modify="{ scope }">
              <span
                @click="relation(scope)"
                v-bind:class="{ relationStyle: scope.row.relation === 0 }"
                >{{ scope.row.relation === 1 ? '已关联' : '关联' }}</span
              >
            </template>
          </publicTable>
        </div>
      </div>
      <PublicButton
        :data="close"
        @onButton="closeButton"
        class="closeButton"
      ></PublicButton>
    </el-drawer>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  reactive,
  ref,
  markRaw,
  onMounted,
  onUnmounted,
} from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'

import {
  measuringConfig,
  clRelationConfig,
  changeRecordConfig,
  formConfig,
  associationConfig,
  closeButtonConfig,
} from '../config/RealoumnLook'
import publicTable from '@/components/Table/index.vue'
import PublicButton from '@/components/Button/index.vue'
import publicForm from '@/components/Form/index.vue'
import { dataProcess } from '../config/data-processing'

export default defineComponent({
  components: {
    publicTable,
    PublicButton,
    publicForm,
  },
  setup() {
    const route = useRoute()
    let activeName = ref('attribute')
    const measuringTable = ref({ table: measuringConfig })
    const clRelationTable = ref({ table: clRelationConfig })
    const changeRecordTable = ref({ table: changeRecordConfig })
    const associationList = ref({ form: formConfig })
    const associationTable = ref({ table: associationConfig })
    const close = ref(closeButtonConfig)
    const isDrawer = ref(false)
    const state = reactive({
      buttonOk: {
        name: '新建关联',
        className: 'roundWhiteIcon', // 1.roundBackWhite  2.roundBackBlue 3.roundBackCal 4.roundIcon 5.roundWhiteIcon 6.roundBlueIcon
        round: true,
        // loading: true,
        icon: markRaw(Plus), // 引用组件
      },
      // attributeAttr: objectattr,
      // attributeCenter: dcCenter,
      // attributeCenterData: dcCenterData,
      attributeList: [],
      isFlag: false,
      clRelationTabs: 'list',
    })
    let data = null
    if (route.params && !localStorage.getItem('rowData')) {
      data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    onUnmounted(() => {
      localStorage.removeItem('rowData')
    })
    onMounted(() => {
      state.attributeList = dataProcess()
    })
    const handleClick = (e: any) => {
      console.log(e, 'e')
      if (e.props.name === 'clRelation') {
        state.isFlag = true
      } else {
        state.isFlag = false
      }
    }
    const deleteData = (data) => {
      console.log('删除成功')
    }
    const clRelationTabs = (type: string) => {
      console.log(type, 'type')
      if (type === 'list') {
        state.clRelationTabs = 'list'
      } else {
        state.clRelationTabs = 'topology'
      }
    }
    const onButton = () => {
      console.log('点击了')
      isDrawer.value = true
    }
    // 搜索新建关联
    const searchFrom = (fuc, data) => {
      console.log('我是搜索')
    }
    // 重置
    const resetFrom = (fuc, data) => {
      console.log('我是重置')
    }
    // 关闭
    const closeButton = () => {
      console.log('关闭弹出框')
      isDrawer.value = false
    }
    // 是否关联
    const relation = (data) => {
      console.log(data, 'data')
    }
    return {
      state,
      measuringTable,
      clRelationTable,
      changeRecordTable,
      activeName,
      associationList,
      associationTable,
      isDrawer,
      close,
      handleClick,
      closeButton,
      deleteData,
      clRelationTabs,
      onButton,
      searchFrom,
      resetFrom,
      relation,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/assets/style/overallStyle.scss';

.demo-tabs {
  position: relative;
}
.drawerStyle {
  margin: 10px;
}
.newButton {
  position: absolute;
  right: 106px;
  top: 12px;
  margin-right: 21px;
}
.clRelationList {
  position: absolute;
  right: 3px;
  top: 13px;
  width: 100px;
  height: 26px;
  border-radius: 8px;
  border: 1px solid #b8b8d5;
  display: flex;
  cursor: pointer;
  .clRelationTab {
    width: 50px;
    font-size: 14px;
    font-family: PingFangSC-Semibold, PingFang SC;
    font-weight: 600;
    color: #595aec;
    text-align: center;
    line-height: 26px;
  }
  .clRelationTabActive {
    background: #e4e4fc;
    border-radius: 8px 0 0 8px;
  }
  .clRelationTabTopology {
    background: #e4e4fc;
    border-radius: 0px 8px 8px 0px;
  }
}
#clRelationModify {
  cursor: pointer;
  font-size: 12px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  margin-right: 15px;
  color: #fa4d05;
}

:deep(.fromSlot .el-form-item__content) {
  flex-wrap: nowrap;
  position: absolute;
  right: 0;
  top: 75px;
}
.headerTop {
  height: 120px;
  border-radius: 4px;
  border: 1px solid #ededfd;
  box-sizing: border-box;
}
:deep(.el-drawer__body) {
  padding: 0 !important;
}
.relationStyle {
  cursor: pointer;
  font-size: 12px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #595aec !important;
}
.drawerTable {
  border-radius: 4px;
  border: 1px solid #ededfd;
}
.closeButton {
  position: absolute;
  bottom: 20px;
  right: 20px;
}

.collapse {
  margin-bottom: 20px;
}
.collapse-span {
  font-size: 14px;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  color: #302e6d;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
:deep(.el-collapse-item__content) {
  display: flex;
  flex-wrap: wrap;
  padding-bottom: 0 !important;
  margin-bottom: 24px;
}
:deep(.collapse-list) {
  margin-left: 20px;
  margin-top: 20px;
  width: 31%;
}
:deep(.el-collapse-item__header) {
  margin-left: 20px;
}
:deep(.el-collapse-item) {
  border-radius: 4px;
  border: 1px solid #ededfd;
}
:deep(.el-collapse) {
  --el-collapse-border-color: none !important;
}

:deep(.el-collapse-item div .is-active) {
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #302e6d;
  border-bottom: 1px solid #ededfd !important;
}
</style>
