<template>
  <div class="realtimealarm">
    <Breadcrum></Breadcrum>
    <TestSearch :queryItems="queryItems" @onSearch="onSearch"></TestSearch>
    <publicTable
      class="realtimealarm-table"
      :data="tableInfo"
      @handleSelectionChange="handleSelectionChange"
      @colChange="colChange"
    >
      <template v-slot:btns class="realtimealarm-btns">
        <div class="realtimealarm-btns-left">
          <ZHButton
            :disabled="!tableCheckedData.length"
            :options="btnOptions1"
            @ZHBtnClick="openConfirmModal"
            >批量确认</ZHButton
          >
          <ZHButton
            :disabled="btnsDisabled"
            :options="btnOptions2"
            @ZHBtnClick="openFeedbackModal"
            >批量反馈</ZHButton
          >
          <ZHButton
            :disabled="btnsDisabled"
            :options="btnOptions2"
            @ZHBtnClick="openCreateModal"
            >批量建单</ZHButton
          >
        </div>
        <div class="realtimealarm-btns-right">
          <span
            >自动刷新：
            <el-switch
              v-model="autoRefresh"
              class="ml-2"
              active-color="#13ce66"
              @change="autoRefreshChange"
            />
          </span>
          <ZHButton
            :disabled="!tableCheckedData.length"
            :options="btnOptions3"
            @ZHBtnClick="topping"
            >置顶</ZHButton
          >
          <ZHButton
            :disabled="!tableCheckedData.length"
            :options="btnOptions3"
            @ZHBtnClick="cancelTopping"
            >取消置顶</ZHButton
          >
        </div>
      </template>
      <template #grade="{ scope }">
        <span>{{ scope.row }}</span>
      </template>
      <template #operation="{ scope }">
        <div class="realtimealarm-table-col-operation">
          <span
            v-if="scope.row.status === 1"
            @click="openConfirmModal(scope.row)"
            >确认</span
          >
          <span
            v-if="scope.row.status === 2"
            @click="openFeedbackModal(scope.row)"
            >反馈</span
          >
          <span v-if="scope.row.status === 2">建单</span>
        </div>
      </template>
    </publicTable>
    <DeleteDialog
      title="告警确认"
      content="您是否进行批量确认操作？"
      :is-show-dialog="isShowConfirm"
      @deleteDialogCancel="closeConfirmModal"
      @deleteDialogSure="confirmSubmit"
    ></DeleteDialog>
    <el-drawer
      v-model="isShowFeedback"
      title="告警反馈"
      direction="rtl"
      :size="460"
    >
      <el-form
        ref="ruleFormRef"
        class="ruleFormRef"
        :model="formData"
        :rules="formRules"
      >
        <el-form-item label="告警原因：" prop="reason">
          <el-radio-group v-model="formData.reason">
            <el-radio :label="3">Option A</el-radio>
            <el-radio :label="6">Option B</el-radio>
            <el-radio :label="9">Option C</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div class="form-line">
        <div class="form-border"></div>
        <span>添加新的告警原因</span>
      </div>
      <el-form
        ref="ruleFormAddReasonRef"
        class="ruleFormAddReasonRef"
        :model="addformData"
        :rules="addformRules"
      >
        <el-form-item label="" prop="remark">
          <el-input
            v-model="addformData.remark"
            :rows="2"
            maxlength="30"
            show-word-limit
            type="textarea"
            placeholder="请输入"
          ></el-input>
        </el-form-item>
        <el-form-item class="add-reason-btn">
          <ZHButton :options="btnOptions3" @ZHBtnClick="addReason"
            >添加</ZHButton
          >
        </el-form-item>
      </el-form>
      <div class="feedback-modal-btns">
        <ZHButton :options="btnOptions3" @ZHBtnClick="closeFeedbackModal"
          >取消</ZHButton
        >
        <ZHButton :options="btnOptions3" @ZHBtnClick="feedbackSubmit"
          >确定</ZHButton
        >
      </div>
    </el-drawer>
    <!-- <DeleteDialog
      title="告警建单"
      content="您是否进行批量建单操作？"
      :is-show-dialog="isShowCreate"
      @deleteDialogCancel="closeCreateModal"
      @deleteDialogSure="createSubmit"
    ></DeleteDialog> -->
  </div>
</template>

<script setup lang="ts">
import { defineComponent, onMounted, reactive, ref, watch } from 'vue'
import TestSearch, { ISearch } from '@/components/Search/index.vue'
import publicTable from '@/components/Table/index.vue'
import ZHButton, { IButton } from '@/components/Button/Button.vue'
import DeleteDialog from '@/components/DeleteDialog/index.vue'
import type { ElForm } from 'element-plus'
type FormInstance = InstanceType<typeof ElForm>

// 按钮
let btnOptions1: IButton = {
  bgColor: '#595AEC',
  textColor: '#fff',
}

let btnOptions2: IButton = {
  bgColor: '#E4E4FC',
  textColor: '#595AEC',
}

let btnOptions3: IButton = {
  bgColor: '#fff',
  textColor: '#595AEC',
  width: '60px',
  height: '26px',
}

// 检索项
let queryItems = ref<ISearch>({
  gutter: 0,
  columnList: [
    { type: 'input', label: '请输入告警内容', value: '' },
    {
      type: 'select',
      label: '位置',
      value: '',
      options: [{ label: 'lablel1', value: 'value1' }],
    },
    { type: 'input', label: '请输入告警名称', value: '' },
    {
      type: 'zhcheckbox',
      label: '告警等级',
      value: ['value1'],
      options: [
        { label: 'E1', value: 'value1' },
        { label: 'E2', value: 'value2' },
        { label: 'E3', value: 'value3' },
        { label: 'I1', value: 'value4' },
        { label: 'I2', value: 'value5' },
        { label: 'I3', value: 'value6' },
      ],
    },
    {
      type: 'select',
      label: '工单状态',
      value: '',
      options: [
        { label: '已建单', value: 'value1' },
        { label: '未建单', value: 'value2' },
      ],
    },
    {
      type: 'select',
      label: '反馈状态',
      value: '',
      options: [
        { label: '已建单', value: 'value1' },
        { label: '未建单', value: 'value2' },
      ],
    },
    {
      type: 'select',
      label: '状态',
      value: '',
      options: [
        { label: '已建单', value: 'value1' },
        { label: '未建单', value: 'value2' },
      ],
    },
    {
      type: 'datepicker',
      label: '工单状态',
      value: '',
    },
  ],
})

// 搜索
const paramsKeys = ['a', 'b', 'c', 'd']
const onSearch = () => {
  let params = {}
  queryItems.value.columnList.forEach((item, index) => {
    params[paramsKeys[index]] = item.value
  })
}

// 表格信息
let tableInfo = ref({
  table: {
    isShowBtns: true, // 是否展示表格上方按钮
    isSelection: true,
    tableData: [
      { id: 1, status: 1 },
      { id: 2, status: 2, operation: 1 },
    ], // 表格数据
    // 表头
    columns: [
      {
        label: '告警ID',
        prop: 'id',
      },
      {
        label: '状态',
        prop: 'id',
      },
      {
        label: '告警时间',
        prop: 'id',
      },
      {
        label: '告警级别',
        prop: 'grade',
        isSlot: true,
      },
      {
        label: '告警名称',
        prop: 'id',
      },
      {
        label: '告警内容',
        prop: 'id',
      },
      {
        label: '关联事件描述',
        prop: 'id',
      },
      {
        label: '收敛数',
        prop: 'id',
      },
      {
        label: '确认时间',
        prop: 'id',
      },
      {
        label: '操作',
        prop: 'operation',
        isSlot: true,
        fixed: 'right',
      },
    ],
  },
})

// 表格勾选
let tableCheckedData = ref([])
const handleSelectionChange = (val) => {
  console.log('val', val)
  tableCheckedData.value = val
}

// 待处理信息
let willHandleData = ref([])

/** 批量确认 */
// 批量弹窗开关
let isShowConfirm = ref(false)
// 开启批量确认弹窗
const openConfirmModal = (row) => {
  isShowConfirm.value = true
  willHandleData = row ? row : tableCheckedData.value
}
// 关闭批量确认弹窗
const closeConfirmModal = () => {
  isShowConfirm.value = false
}
// 批量确认
const confirmSubmit = () => {
  console.log()
}

// 批量反馈、批量简单按钮是否可点击
let btnsDisabled = ref(false)
watch(
  tableCheckedData,
  (newValue, oldValue) => {
    if (!newValue.length) {
      btnsDisabled.value = true
    } else {
      btnsDisabled.value = false
      newValue.forEach((item) => {
        if (item.status === 1) {
          btnsDisabled.value = true
        }
      })
    }
  },
  { immediate: true },
)

/** 批量反馈 */
// 批量反馈弹窗开关
let isShowFeedback = ref(false)
// 开启批量反馈弹窗
const openFeedbackModal = (row) => {
  isShowFeedback.value = true
  willHandleData = row ? row : tableCheckedData.value
}
// 关闭批量反馈
const closeFeedbackModal = () => {
  isShowFeedback.value = false
}
// 批量反馈表单
const ruleFormRef = ref<FormInstance>()
const formRules = {
  reason: [
    {
      required: true,
      message: '请选择告警原因',
      trigger: 'blur',
    },
  ],
}
let formData = ref({
  reason: 3,
})
// 批量反馈添加原因表单
const ruleFormAddReasonRef = ref<FormInstance>()
const addformRules = {
  remark: [
    {
      required: true,
      message: '请填写新的告警原因',
      trigger: 'blur',
    },
  ],
}
let addformData = ref({
  remark: '',
})
// 添加原因按钮
const addReason = () => {
  console.log()
}
// 确定反馈按钮
const feedbackSubmit = () => {
  console.log()
}

// 批量建单
let isShowCreate = ref(false)
const openCreateModal = () => {
  isShowCreate.value = true
}
const closeCreateModal = () => {
  isShowCreate.value = false
}
const createSubmit = () => {
  console.log()
}

// 自动刷新开发
let autoRefresh = ref(true)
const autoRefreshChange = (val) => {
  console.log(val)
}

const colChange = (val) => {
  tableInfo.value.table.columns = val
}

// 置顶按钮
const topping = () => {
  console.log()
}

// 取消置顶
const cancelTopping = () => {
  console.log()
}
</script>

<style lang="scss" scoped>
@import '@/assets/style/var.scss';
.realtimealarm {
  .realtimealarm-table {
    .btns-slot {
      .realtimealarm-btns-left {
      }
      .realtimealarm-btns-right {
        flex: 1;
        text-align: right;
        span {
          margin-right: 20px;
        }
      }
    }
    .realtimealarm-table-col-operation {
      span {
        margin-right: 15px;
        cursor: pointer;
        color: $main-color;
      }
    }
  }
  .form-line {
    margin-top: 20px;
    .form-border {
      border-top: 1px solid #f6f6f6;
      margin-bottom: 20px;
    }
    span {
      color: $blue-text-color;
      line-height: 20px;
      font-size: 16px;
      display: block;
      padding-left: 94px;
      margin-bottom: 20px;
    }
  }
  .feedback-modal-btns {
    position: absolute;
    bottom: 0;
    width: 100%;
    height: 70px;
    border-top: 1px solid #f6f6f6;
    text-align: right;
    button {
      margin-top: 20px;
      margin-right: 15x;
    }
    button:nth-child(2) {
      margin-right: 40px;
    }
  }
}
</style>
<style lang="scss">
.realtimealarm {
  .ruleFormRef {
    .el-radio-group,
    .el-radio {
      display: block;
    }
  }
  .ruleFormAddReasonRef {
    padding-left: 94px;
    padding-right: 54px;
    .add-reason-btn {
      display: block;
      text-align: right;
    }
  }
}
</style>
