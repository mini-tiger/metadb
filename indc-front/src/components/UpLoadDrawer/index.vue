<template>
  <div class="upload-drawer">
    <el-drawer
      v-model="isShow"
      direction="rtl"
      :show-close="false"
      :close-on-click-modal="false"
      size="460px"
    >
      <template #title>
        <span class="title">{{ UploadTitle }}</span>
      </template>
      <template #default>
        <el-upload
          class="upload-demo"
          drag
          action="https://jsonplaceholder.typicode.com/posts/"
          accept="xls,xlsx"
          :auto-upload="false"
        >
          <svg-icon icon-class="upload" class="upload-icon"
            ><upload-filled
          /></svg-icon>
          <div class="el-upload__text">拖拽文件到此处，或点击上传文件</div>
          <template #tip>
            <div class="el-upload__tip">
              <span>注意：</span>只能上传xlsx文件，且不超过XXMB
            </div>
            <div class="el-upload__download">
              <el-button class="btn">下载模板</el-button>
              <span class="tips">只支持新增导入</span>
            </div>
          </template>
        </el-upload>
      </template>
      <template #footer>
        <div style="flex: auto" class="drawerFooter">
          <publicButton
            :data="buttonData.buttonCancel"
            @onButton="btnCancel(ruleFormRef)"
          ></publicButton>
          <publicButton
            :data="buttonData.buttonSure"
            @onButton="btnSure(ruleFormRef)"
          ></publicButton>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from 'vue'
import { ElForm } from 'element-plus'
import publicButton from '@/components/Button/index.vue'
import { UploadFilled } from '@element-plus/icons-vue'
export default defineComponent({
  name: 'UpLoadDialog',
  props: {
    isShowUpLoad: {
      type: Boolean,
      default: true,
    },
    title: {
      type: String,
      default: '导入',
    },
  },
  components: {
    publicButton,
    UploadFilled,
  },
  emits: ['btnCancel', 'btnSure'],
  setup(props, ctx) {
    const ruleFormRef = ref<InstanceType<typeof ElForm>>()
    let { isShowUpLoad: isShow, title: UploadTitle } = toRefs(props)
    const btnCancel = () => {
      ctx.emit('btnCancel')
    }
    const btnSure = () => {
      ctx.emit('btnSure')
    }
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
    return {
      isShow,
      UploadTitle,
      btnCancel,
      btnSure,
      buttonData,
      ruleFormRef,
    }
  },
})
</script>

<style lang="scss">
.upload-drawer {
  .el-upload-dragger {
    height: 320px;
    width: 420px;
    border-radius: 6px;
    border: 1px dashed #302e6d;
    .upload-icon {
      font-size: 60px;
      margin-top: 100px;
      margin-bottom: 20px;
    }
    .el-upload__text {
      font-size: 14px;
      font-family: PingFangSC-Medium, PingFang SC;
      font-weight: 500;
      color: #302e6d;
    }
    .el-upload__tip {
      font-size: 12px;
      font-family: PingFangSC-Semibold, PingFang SC;
      color: #302e6d;
      span {
        font-weight: 600;
      }
    }
  }
  .el-upload__download {
    display: flex;
    margin-top: 30px;
    height: 26px;
    align-items: center;
    .btn {
      width: 80px;
      height: 26px;
      border-radius: 13px;
      border: 1px solid #595aec;
      span {
        font-size: 14px;
        font-family: PingFangSC-Medium, PingFang SC;
        font-weight: 500;
        color: #595aec;
      }
    }
    .tips {
      font-size: 12px;
      font-family: PingFangSC-Regular, PingFang SC;
      font-weight: 400;
      color: #a1a1c3;
      margin-left: 10px;
    }
  }
}
</style>
