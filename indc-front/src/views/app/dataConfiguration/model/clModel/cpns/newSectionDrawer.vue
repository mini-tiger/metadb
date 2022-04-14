<template>
  <el-drawer
    v-model="isShow"
    direction="rtl"
    :show-close="false"
    :close-on-click-modal="false"
    size="460px"
  >
    <template #title>
      <span class="title">{{ drawerTitle }}</span>
    </template>
    <template #default>
      <!--      新建段落-->
      <el-form
        ref="ruleFormRef"
        :model="drawerFormData.ruleForm"
        :rules="drawerFormData.rules"
        class="drawer-form"
        label-position="right"
        label-width="90px"
      >
        <el-form-item label="唯一标识 :" prop="bk_group_id">
          <el-input
            v-model.trim="drawerFormData.ruleForm.bk_group_id"
            maxlength="15"
            :show-word-limit="false"
            placeholder="请输入"
            :disabled="isDisable"
            style="width: 95%"
          ></el-input>
          <el-tooltip
            class="box-item"
            effect="light"
            content="不允许重复"
            placement="top-start"
          >
            <span><svg-icon icon-class="tips"></svg-icon> </span>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="段落名称 :" prop="bk_group_name">
          <el-input
            v-model.trim="drawerFormData.ruleForm.bk_group_name"
            maxlength="15"
            :show-word-limit="false"
            placeholder="请输入"
            style="width: 95%"
          ></el-input>
          <el-tooltip
            class="box-item"
            effect="light"
            content="不允许重复"
            placement="top-start"
          >
            <span><svg-icon icon-class="tips"></svg-icon> </span>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="是否标签 :" prop="is_collapse">
          <el-radio
            v-model="drawerFormData.ruleForm.is_collapse"
            :label="true"
            :disabled="isDisable"
            >是</el-radio
          >
          <el-radio
            v-model="drawerFormData.ruleForm.is_collapse"
            :label="false"
            :disabled="isDisable"
            >否</el-radio
          >
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto" class="drawerFooter">
        <publicButton
          :data="buttonCancel"
          @onButton="drawerData.drawerButtonCancel(ruleFormRef)"
        ></publicButton>
        <publicButton
          :data="buttonSure"
          @onButton="
            drawerData.drawerButtonSure(ruleFormRef, drawerFormData.ruleForm)
          "
        ></publicButton>
      </div>
    </template>
  </el-drawer>
</template>

<script lang="ts">
import { defineComponent, ref, toRefs } from 'vue'
import { drawerFormData, drawerData, state } from '../config/model-filed'
import publicButton from '@/components/Button/index.vue'
import { ElForm } from 'element-plus'
export default defineComponent({
  name: 'newSectionDrawer',
  components: {
    publicButton,
  },
  props: {
    isShowDrawer: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: '新建段落',
    },
    isDisabled: {
      type: Boolean,
      default: false,
    },
  },
  setup(props) {
    const ruleFormRef = ref<InstanceType<typeof ElForm>>()
    let {
      isShowDrawer: isShow,
      title: drawerTitle,
      isDisabled: isDisable,
    } = toRefs(props)
    return {
      ...toRefs(state),
      drawerData,
      drawerFormData,
      ruleFormRef,
      isShow,
      drawerTitle,
      isDisable,
    }
  },
})
</script>

<style scoped lang="scss">
.drawer-form {
  .el-form-item__content {
    flex-wrap: nowrap;
    .el-tooltip__trigger {
      margin-left: 10px;
      font-size: 16px;
    }
    justify-content: center;
    span {
      color: #ffffff;
    }
  }
}
</style>
