<template>
  <draggable
    v-model="dataList"
    group="section"
    animation="300"
    dragClass="dragClass"
    ghostClass="ghostClass"
    chosenClass="chosenClass"
    :draggable="'.collapse'"
    handle=".move-icon"
    @change="moveSection"
    :disabled="type == 'look'"
  >
    <template #item="{ element }">
      <div class="collapse">
        <div class="header">
          <div class="left">
            <span class="section-name">{{ element.bk_group_name }}</span>
            <span
              class="edit-icon icon"
              @click="editSection(element)"
              v-if="type !== 'look'"
            >
              <svg-icon icon-class="edit"></svg-icon>
            </span>
          </div>
          <div class="right">
            <span class="move-icon icon" v-if="type !== 'look'">
              <svg-icon icon-class="move"></svg-icon>
            </span>
            <span
              class="del-icon icon"
              @click="delSection(element)"
              v-if="type !== 'look'"
              ><svg-icon icon-class="del"></svg-icon
            ></span>
            <span class="down-icon icon" @click="collapse(element.bk_group_id)"
              ><svg-icon icon-class="down"></svg-icon
            ></span>
          </div>
        </div>
        <div
          class="content"
          v-show="isCollapseList.indexOf(element.bk_group_id) > -1"
        >
          <draggable
            v-model="element.list"
            :group="element.bk_group_id"
            animation="300"
            dragClass="dragClass"
            ghostClass="ghostClass"
            chosenClass="chosenClass"
            @change="moveFiled(element)"
            :draggable="'.item'"
            handle=".input,.name,.tag"
            :disabled="type == 'look'"
          >
            <template #item="{ element }">
              <div class="item" v-if="!element.is_collapse">
                <div class="filed" v-if="element.name">
                  <p class="name">
                    <span class="is-required">*</span>{{ element.name }} :
                  </p>
                  <p class="input"></p>

                  <p
                    class="icon"
                    @click="editFiled(element)"
                    v-if="element.name !== '实例名称' && type !== 'look'"
                  >
                    <svg-icon icon-class="edit"></svg-icon>
                  </p>
                  <p
                    class="icon"
                    @click="delFiled(element)"
                    v-if="element.name !== '实例名称' && type !== 'look'"
                  >
                    <svg-icon icon-class="del"></svg-icon>
                  </p>
                </div>
                <div v-else class="new-filed" @click="newFiled(element)">
                  <span class="new" v-if="type !== 'look'"
                    ><svg-icon icon-class="upload"></svg-icon
                  ></span>
                  <span v-if="type !== 'look'">{{
                    element.is_collapse ? '添加标签' : '添加字段'
                  }}</span>
                </div>
              </div>
              <div class="item item-tag" v-else>
                <div class="tag" v-if="element.name">
                  <div class="tag-name">
                    {{ element.name }}
                  </div>
                  <p
                    class="icon"
                    @click="editFiled(element)"
                    v-if="type !== 'look'"
                  >
                    <svg-icon icon-class="edit"></svg-icon>
                  </p>
                  <p
                    class="icon"
                    @click="delFiled(element)"
                    v-if="type !== 'look'"
                  >
                    <svg-icon icon-class="del"></svg-icon>
                  </p>
                </div>
                <div v-else class="new-filed" @click="newFiled(element)">
                  <span class="new" v-if="type !== 'look'"
                    ><svg-icon icon-class="upload"></svg-icon
                  ></span>
                  <span v-if="type !== 'look'">{{
                    element.is_collapse ? '添加标签' : '添加字段'
                  }}</span>
                </div>
              </div>
            </template>
          </draggable>
        </div>
      </div>
    </template>
  </draggable>
  <div class="new-section" v-if="type !== 'look'">
    <publicButton :data="buttonNew" @onButton="newSection"></publicButton>
  </div>

  <!--  新建段落-->
  <newSectionOrTagDrawer
    :is-show-drawer="drawerData.isShowDrawer"
    :title="drawerData.title"
    :is-disabled="drawerData.isDisabled"
  ></newSectionOrTagDrawer>

  <!--  新建字段-->

  <newFiled
    :IsDrawer="drawerFiledData.isShowFiled"
    :isEmit="drawerFiledData.isEditFiled"
    :pointFlag="drawerFiledData.filedType"
    :isTag="drawerFiledData.isTag"
    :title="drawerFiledData.title"
    @defineData="drawerFiledData.btnSure"
    @IsDrawer="drawerFiledData.btnCancel"
  ></newFiled>
  <!--  删除-->
  <DeleteDialog
    :is-show-dialog="isShowDelete"
    :title="title"
    :content="content"
    @deleteDialogCancel="deleteDialogCancel"
    @deleteDialogSure="deleteDialogSure"
  />
</template>

<script lang="ts">
import { ref, defineComponent, toRefs, onMounted } from 'vue'
import draggable from 'vuedraggable'
import {
  handleMethods,
  state,
  drawerData,
  drawerFiledData,
} from '../config/model-filed'
import publicButton from '@/components/Button/index.vue'
import { ElForm } from 'element-plus'
import DeleteDialog from '@/components/DeleteDialog/index.vue'
import newSectionOrTagDrawer from './newSectionDrawer.vue'
import newFiled from '../hook/use-newFiled.vue'
import { useRoute } from 'vue-router'
export default defineComponent({
  components: {
    draggable,
    publicButton,
    DeleteDialog,
    newSectionOrTagDrawer,
    newFiled,
  },
  setup() {
    const ruleFormRef = ref<InstanceType<typeof ElForm>>()
    const isCollapseList = ref(['default', 'info', 'tag', 'operations'])
    //展开收起
    const collapse = (value) => {
      if (isCollapseList.value.includes(value)) {
        isCollapseList.value.splice(isCollapseList.value.indexOf(value), 1)
      } else {
        isCollapseList.value.push(value)
      }
    }
    const route = useRoute()

    if (route.params && !localStorage.getItem('rowData')) {
      const data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    const type = JSON.parse(localStorage.getItem('rowData')).type || ''

    return {
      isCollapseList,
      collapse,
      ...toRefs(state),
      ...toRefs(handleMethods),
      ruleFormRef,
      drawerData,
      type,
      drawerFiledData,
    }
  },
})
</script>

<style lang="scss" scoped>
.collapse {
  background: #ffffff;
  margin-top: 20px;
  .header {
    height: 40px;
    border: 1px #ededfd solid;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 10px 0 20px;
    box-sizing: border-box;
    border-radius: 4px 4px 0 0;

    .left {
      display: flex;
      align-items: center;
      justify-content: space-between;
      .section-name {
        font-size: 14px;
        font-family: PingFangSC-Medium, PingFang SC;
        font-weight: 500;
        color: #302e6d;
      }
    }
    .icon {
      font-size: 16px;
      margin-left: 10px;
      cursor: pointer;
    }
    .move-icon {
      cursor: move;
    }
  }

  .content {
    border: 1px #ededfd solid;
    padding: 10px 2px;
    box-sizing: border-box;
    border-radius: 0 0 4px 4px;
    border-top: 0;
    //height: 300px;
    //display: flex;
    div {
      display: flex;
      flex-wrap: wrap;
    }
  }
}

/*定义要拖拽元素的样式*/
.ghostClass {
  border-radius: 4px;
  border: 1px dashed #595aec !important;
}
//选中元素样式
.chosenClass {
  border-radius: 4px;
  border: 1px dashed #595aec !important;
}
.dragClass {
  border-radius: 4px;
  border: 1px dashed #595aec !important;
}

.item {
  .filed {
    width: 390px;
    height: 50px;
    //padding: 10px 30px 10px 20px;
    display: flex;
    align-items: center;
    box-sizing: border-box;
    border: 1px dashed #ffffff;
    margin-right: 20px;
    cursor: move;
  }

  .is-required {
    font-size: 12px;
    font-weight: 500;
    color: #fa4d05;
    margin-right: 5px;
  }
  .name {
    font-weight: 400;
    color: #302e6d;
    width: 120px;
    text-align: right;
  }
  .input {
    width: 200px;
    height: 30px;
    background: #efefef;
    border-radius: 4px;
    border: 1px solid #ebebf1;
    margin-left: 10px;
  }
  .icon {
    margin-left: 10px;
    cursor: pointer;
  }

  &:nth-last-child(1) {
    width: 120px;
  }
}
.item.item-tag {
  .tag {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    margin: 10px 20px;
    //width: 170px;
    width: auto;
    padding: 0 10px;
  }

  .tag-name {
    //width: 140px;
    padding: 0 20px;
    margin-right: 10px;
    box-sizing: border-box;
    display: flex;
    align-items: center;
    //justify-content: center;
    height: 30px;
    border-radius: 15px;
    border: 1px solid #302e6d;
    font-size: 12px;
    color: #302e6d;
  }
}
//.item:hover {
//  border-radius: 4px;
//  border: 1px dashed #595aec !important;
//  cursor: move;
//}

//.item.item_tag:hover {
//  width: auto !important;
//  border-radius: 4px;
//  border: 1px dashed #595aec !important;
//  cursor: move;
//}
.new-section {
  margin-top: 20px;
  margin-bottom: 50px;
}
.new-filed {
  width: 120px;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  cursor: pointer;
  .new {
    font-size: 16px;
    margin-right: 2px;
    display: flex;
    align-items: center;
  }
  span {
    font-size: 16px;
    color: #595aec;
    margin-left: 2px;
    line-height: 16px;
  }
}
</style>
