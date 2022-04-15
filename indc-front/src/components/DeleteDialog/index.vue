<template>
  <ElDialog
    v-model="isShowDia"
    :title="tips"
    width="460px"
    :close-on-click-modal="false"
    :show-close="false"
  >
    <span>{{ bodyContent }}</span>
    <template #footer>
      <span class="dialog-footer">
        <publicButton :data="buttonData.buttonCancel" @onButton="dialogCancel"
          >取消</publicButton
        >
        <publicButton :data="buttonData.buttonSure" @onButton="dialogSure"
          >确定</publicButton
        >
      </span>
    </template>
  </ElDialog>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from 'vue'
import { ElDialog } from 'element-plus'
import 'element-plus/es/components/dialog/style'
import publicButton from '@/components/Button/index.vue'
export default defineComponent({
  name: 'DeleteDialog',
  components: {
    publicButton,
    ElDialog,
  },
  props: {
    isShowDialog: {
      //是否展示弹窗
      type: Boolean,
      default: false,
    },
    title: {
      //标题
      type: String,
      default: '删除分类',
    },
    content: {
      //提示内容
      type: String,
      default: '',
      require: true,
    },
  },
  emits: ['deleteDialogCancel', 'deleteDialogSure'],
  setup(props, ctx) {
    let {
      isShowDialog: isShowDia,
      title: tips,
      content: bodyContent,
    } = toRefs(props)
    const buttonData = reactive({
      buttonCancel: {
        name: '取消',
        className: 'roundBackCal',
        round: true,
      },
      buttonSure: {
        name: '确定',
        className: 'roundBackBlue',
        round: true,
      },
    })
    //取消
    const dialogCancel = () => {
      ctx.emit('deleteDialogCancel')
    }
    //确认
    const dialogSure = () => {
      ctx.emit('deleteDialogSure')
    }
    return {
      dialogCancel,
      dialogSure,
      buttonData,
      isShowDia,
      tips,
      bodyContent,
    }
  },
})
</script>

<style lang="scss">
.el-dialog {
  height: 300px;
  background: #ffffff;
  box-shadow: 0px 0px 9px 0px rgba(0, 0, 0, 0.3);
  border-radius: 8px;
}

.el-dialog__header {
  height: 50px;
  line-height: 50px;
  border-bottom: 1px solid #f3f3f3;
  padding: 0 0 0 20px !important;
  font-size: 16px;
  color: #302e6d;
}
.el-dialog__body {
  height: 188px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.el-dialog__footer {
  height: 60px;
  line-height: 60px;
  border-top: 1px solid #f3f3f3;
  padding: 0 20px 0 0 !important;
  .dialog-footer {
    display: flex;
    align-items: center;
    justify-content: flex-end;
    button {
      margin-left: 20px;
    }
  }
}
</style>
