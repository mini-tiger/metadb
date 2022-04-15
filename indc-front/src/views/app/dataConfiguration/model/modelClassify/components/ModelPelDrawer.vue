<template>
  <div>
    <el-drawer
      v-model="isShow"
      direction="rtl"
      :show-close="false"
      :close-on-click-modal="false"
      size="460px"
      destroy-on-close
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
          label-width="86px"
          :disabled="isDisable"
        >
          <el-form-item label="上级分类 :" prop="parentName">
            <div>{{ drawerFormData.ruleForm.parentName }}</div>
          </el-form-item>

          <el-form-item label="分类名称 :" prop="classificationName">
            <el-input
              v-model.trim="drawerFormData.ruleForm.classificationName"
              maxlength="10"
              :show-word-limit="false"
              placeholder="请输入分类名称"
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
          <el-form-item label="分类排序 :" prop="orderBy">
            <el-input-number
              v-model="drawerFormData.ruleForm.orderBy"
              :min="1"
            />
          </el-form-item>
          <el-form-item label="目前状态 :" prop="status">
            <el-radio v-model="drawerFormData.ruleForm.status" label="1"
              >启用</el-radio
            >
            <el-radio v-model="drawerFormData.ruleForm.status" label="2"
              >停用</el-radio
            >
          </el-form-item>
          <el-form-item label="描述 :" prop="describe">
            <el-input
              type="textarea"
              maxlength="30"
              show-word-limit
              v-model.trim="drawerFormData.ruleForm.describe"
              style="width: 92%"
            >
            </el-input>
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
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from 'vue'
import publicButton from '@/components/Button/index.vue'
import { drawerFormData, buttonData, tableData } from '../config/config'
import { ElForm } from 'element-plus'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'
export default defineComponent({
  name: 'PelDrawer',
  props: {
    isShowDrawer: {
      type: Boolean,
      default: true,
    },
    drawerTitle: {
      type: String,
      default: '新建分类',
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
    const store = useStore()
    const drawerData = reactive({
      drawerButtonSure: async (ruleFormRef, data) => {
        if (!ruleFormRef) return
        await ruleFormRef.validate((valid, fields) => {
          if (valid) {
            if (!data.id) {
              const params = {
                parentId: data.parentId,
                parentName: data.parentName == '无' ? '-' : data.parentName,
                classificationName: data.classificationName,
                orderBy: data.orderBy,
                status: data.status, //状态：1正常 2停用
                describe: data.describe,
                level: data.level, //等级
              }
              store
                .dispatch(`modelClassify/${types.SAVE_MODEL_CLASSIFY}`, params)
                .then(() => {
                  tableData.table.tableData = store.state.modelClassify.data
                })
            } else {
              const params = {
                id: data.id,
                parentId: data.parentId,
                parentName: data.parentName == '无' ? '-' : data.parentName,
                classificationName: data.classificationName,
                orderBy: data.orderBy,
                status: data.status, //状态：1正常 2停用
                describe: data.describe,
                level: data.level, //等级
              }
              store
                .dispatch(
                  `modelClassify/${types.UPDATE_MODEL_CLASSIFY}`,
                  params,
                )
                .then(() => {
                  tableData.table.tableData = store.state.modelClassify.data
                })
            }

            ctx.emit('closeDrawer', false)
            ruleFormRef.resetFields()
          } else {
            console.log('error submit!', fields)
          }
        })
      },
      drawerButtonCancel: (
        ruleFormRef: InstanceType<typeof ElForm> | undefined,
      ) => {
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
