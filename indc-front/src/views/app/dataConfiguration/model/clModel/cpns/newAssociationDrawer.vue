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
        label-width="90px"
        :disabled="isDisable"
      >
        <el-form-item label="源模型 :" prop="src_model">
          <el-select
            v-model="drawerFormData.ruleForm.src_model"
            placeholder="请选择源模型"
            style="width: 90%"
          >
            <el-option label="数据中心" value="数据中心" />
            <el-option label="UPS设备模型" value="UPS设备模型" />
          </el-select>
        </el-form-item>
        <el-form-item label="目标模型 :" prop="dest_model">
          <el-select
            v-model="drawerFormData.ruleForm.dest_model"
            placeholder="请选择目标模型"
            style="width: 90%"
          >
            <el-option label="UPS设备模型" value="UPS设备模型" />
            <el-option label="数据中心" value="数据中心" />
          </el-select>
        </el-form-item>
        <el-form-item label="关系类型 :" prop="relation_type">
          <el-select
            v-model="drawerFormData.ruleForm.relation_type"
            placeholder="请选择关系类型"
            style="width: 90%"
          >
            <el-option label="空间关系" value="空间关系" />
            <el-option label="属于" value="属于" />
            <el-option label="上下级关系" value="上下级关系" />
          </el-select>
          <span class="new-icon" @click="toRelationConfig"
                v-if='!isDisable'
            ><svg-icon icon-class="upload"></svg-icon
          ></span>
        </el-form-item>
        <el-form-item label="源至目标约束 :" prop="src_to_dest">
          <el-select
            v-model="drawerFormData.ruleForm.src_to_dest"
            placeholder="请选择源至目标约束"
            style="width: 90%"
          >
            <el-option label="N-N" value="N-N" />
            <el-option label="1-N" value="1-N" />
            <el-option label="1-1" value="1-1" />
          </el-select>
        </el-form-item>
        <el-form-item
          label="效果展示 :"
          v-if="
            drawerFormData.ruleForm.src_model &&
            drawerFormData.ruleForm.relation_type &&
            drawerFormData.ruleForm.dest_model
          "
        >
          <div class="item">
            <span class="icon">
              <svg-icon icon-class="upload"></svg-icon>
            </span>
            <span class="name">{{ drawerFormData.ruleForm.src_model }}</span>
          </div>
          <div class="line">
            <span class="relation">{{
              drawerFormData.ruleForm.relation_type
            }}</span>
            <span class="arrow"> </span>
          </div>
          <div class="item">
            <span class="icon">
              <svg-icon icon-class="upload"></svg-icon>
            </span>
            <span class="name">{{ drawerFormData.ruleForm.dest_model }}</span>
          </div>
        </el-form-item>
      </el-form>
    </template>
    <template #footer>
      <div style="flex: auto" class="drawerFooter">
        <publicButton
          v-if="isDisabled"
          :data="buttonData.buttonClose"
          @onButton="drawerButtonCancel(ruleFormRef)"
        ></publicButton>
        <publicButton
          v-if="!isDisabled"
          :data="buttonData.buttonCancel"
          @onButton="drawerButtonCancel(ruleFormRef)"
        ></publicButton>
        <publicButton
          v-if="!isDisabled"
          :data="buttonData.buttonSure"
          @onButton="drawerButtonSure(ruleFormRef, drawerFormData.ruleForm)"
        ></publicButton>
      </div>
    </template>
  </el-drawer>
</template>

<script lang="ts">
import { defineComponent, ref, toRefs } from 'vue'
import publicButton from '@/components/Button/index.vue'
import { ElForm } from 'element-plus'
import { buttonData, drawerFormData } from '../config/model-association'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'newAssociationDrawer',
  props: {
    isShowDrawer: {
      type: Boolean,
      default: true,
    },
    drawerTitle: {
      type: String,
      default: '新建关联',
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
    const drawerButtonSure = async (ruleFormRef, data) => {
        if (!ruleFormRef) return
        await ruleFormRef.validate((valid, fields) => {
          if (valid) {
            console.log('submit!', data)
            ctx.emit('closeDrawer', false)
          } else {
            console.log('error submit!', fields)
          }
        })
      },
      drawerButtonCancel = (ruleFormRef) => {
        if (!ruleFormRef) return
        ruleFormRef.resetFields()
        ctx.emit('closeDrawer', false)
      }

    const router = useRouter()
    const toRelationConfig = () => {
      router.push({ path: '/app-center/app-menu/relationConfig' })
    }
    return {
      isShow,
      title,
      isDisable,
      drawerFormData,
      drawerButtonSure,
      drawerButtonCancel,
      buttonData,
      ruleFormRef,
      toRelationConfig,
    }
  },
})
</script>

<style lang="scss">
.drawer-form {
  .el-form-item--default {
    margin-bottom: 20px;
  }
  .el-form-item__label {
    line-height: 18px;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    height: auto !important;
  }
  .el-form-item__content {
    flex-wrap: nowrap;
    display: flex;
    align-items: center;
  }
  .new-icon {
    width: 10%;
    box-sizing: border-box;
    font-size: 20px;
    padding-left: 10px;
  }

  .item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    .icon {
      display: flex;
      width: 36px;
      height: 36px;
      border-radius: 36px;
      background: #595aec;
      border-right: 36px;
      align-items: center;
      justify-content: center;
    }
    .name {
      font-size: 12px;
      color: #302e6d;

      font-weight: 500;
    }
  }
  .line {
    width: 96px;
    margin: 0 20px;
    text-align: center;
    .relation {
      font-size: 12px;
      font-weight: 500;
      color: #595aec;
    }
    .arrow {
      display: inline-block;
      width: 90px;
      height: 2px;
      background: #302e6d;
      border-radius: 1px;
      margin-bottom: 18px;
    }
  }
}
</style>
