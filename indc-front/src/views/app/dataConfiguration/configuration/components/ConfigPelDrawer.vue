<template>
  <el-drawer
    v-model="isShow"
    direction="rtl"
    :show-close="false"
    :close-on-click-modal="false"
    size="460px"
  >
    <template #title>
      <span class="title">{{ title }}</span>
    </template>
    <template #default>
      <el-form
        ref="ruleFormRef"
        :model="drawerFormData.ruleForm"
        :rules="drawerFormData.rules"
        class="drawer-form"
        label-position="right"
        label-width="120px"
        :disabled="isDisable"
      >
        <el-form-item label="关系名称 :" prop="bk_asst_name">
          <el-input
            v-model.trim="drawerFormData.ruleForm.bk_asst_name"
            maxlength="15"
            :show-word-limit="false"
            placeholder="请输入关系名称"
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
        <el-form-item label="源至目标描述 :" prop="src_des">
          <el-input
            v-model.trim="drawerFormData.ruleForm.src_des"
            maxlength="10"
            :show-word-limit="false"
            placeholder="请输入关联描述如：关联、属于、父级"
            style="width: 92%"
          ></el-input>
        </el-form-item>
        <el-form-item label="目标至源描述 :" prop="dest_des">
          <el-input
            v-model.trim="drawerFormData.ruleForm.dest_des"
            maxlength="10"
            :show-word-limit="false"
            placeholder="请输入关联描述如：包含、上联、子级"
            style="width: 92%"
          ></el-input>
        </el-form-item>
        <el-form-item label="是否有方向 :" prop="is_direction">
          <el-radio v-model="drawerFormData.ruleForm.is_direction" label="1"
            >源指向目标</el-radio
          >
          <el-radio v-model="drawerFormData.ruleForm.is_direction" label="2"
            >双向</el-radio
          >
          <el-radio v-model="drawerFormData.ruleForm.is_direction" label="3"
            >无方向</el-radio
          >
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto" class="drawerFooter">
        <publicButton
          v-if="isDisabled"
          :data="buttonData.buttonClose"
          @onButton="drawerData.drawerButtonCancel(ruleFormRef)"
        ></publicButton>
        <publicButton
          v-if="!isDisabled"
          :data="buttonData.buttonCancel"
          @onButton="drawerData.drawerButtonCancel(ruleFormRef)"
        ></publicButton>
        <publicButton
          v-if="!isDisabled"
          :data="buttonData.buttonSure"
          @onButton="
            drawerData.drawerButtonSure(ruleFormRef, drawerFormData.ruleForm)
          "
        ></publicButton>
      </div>
    </template>
  </el-drawer>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from 'vue'
import publicButton from '@/components/Button/index.vue'
import { drawerFormData, buttonData } from '../config/relationConfig'
import { ElForm } from 'element-plus'
export default defineComponent({
  name: 'PelDrawer',
  props: {
    isShowDrawer: {
      type: Boolean,
      default: true,
    },
    drawerTitle: {
      type: String,
      default: '新建关系类型',
    },
    isDisabled: {
      type: Boolean,
      default: false,
    },
  },
  components: {
    publicButton,
  },
  emits: ['closeDrawer'],
  setup(props, ctx) {
    let {
      isShowDrawer: isShow,
      drawerTitle: title,
      isDisabled: isDisable,
    } = toRefs(props)
    const ruleFormRef = ref<InstanceType<typeof ElForm>>()
    const drawerData = reactive({
      drawerButtonSure: async (ruleFormRef, data) => {
        if (!ruleFormRef) return
        await ruleFormRef.validate((valid, fields) => {
          if (valid) {
            console.log('submit!', data)

            ctx.emit('closeDrawer', false)
            ruleFormRef.resetFields()
          } else {
            console.log('error submit!', fields)
          }
        })
      },
      drawerButtonCancel: (ruleFormRef) => {
        if (!ruleFormRef) return
        ruleFormRef.resetFields()
        ctx.emit('closeDrawer', false)
      },
    })
    return {
      isShow,
      title,
      isDisable,
      drawerFormData,
      drawerData,
      buttonData,
      ruleFormRef,
    }
  },
})
</script>

<style lang="scss">
.drawer-form {
  .el-form-item__content {
    flex-wrap: nowrap;
    .el-tooltip__trigger {
      margin-left: 10px;
      font-size: 16px;
    }
  }
}
</style>
