<template>
  <div class="newFiled">
    <el-drawer
      ref="newFileRef"
      :model-value="isShow"
      direction="rtl"
      :title="drawerTitle"
      :size="460"
      :show-close="false"
      :destroy-on-close="true"
      :close-on-click-modal="false"
    >
      <template #footer>
        <div class="footerButton">
          <publicButton :data="cancelButton" @onButton="onButtonCancel" />
          <publicButton :data="determineButton" @onButton="onButtonOk" />
        </div>
      </template>
      <el-form
        ref="FormRefDefine"
        :model="state.defineData"
        label-width="auto"
        :rules="state.rulesConfig"
        class="demo-ruleForm"
        v-if="!tag"
      >
        <el-form-item label="唯一标识 :" prop="id">
          <el-input
            v-model="state.defineData.id"
            style="width: 300px"
            :disabled="isDisabled"
            :maxlength="30"
          />
          <el-tooltip
            class="box-item"
            effect="light"
            content="不允许重复"
            placement="top-start"
          >
            <span><svg-icon icon-class="tips"></svg-icon> </span>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="标题 :" prop="title">
          <el-input
            v-model="state.defineData.title"
            style="width: 300px"
            :maxlength="30"
          />
          <el-tooltip
            class="box-item"
            effect="light"
            content="不允许重复"
            placement="top-start"
          >
            <span><svg-icon icon-class="tips"></svg-icon> </span>
          </el-tooltip>
        </el-form-item>
        <el-form-item label="字段类型 :" prop="filedKind">
          <el-select
            v-model="state.defineData.filedKind"
            class="m-2"
            placeholder="请输入"
            style="width: 300px"
            @change="change"
            :disabled="isDisabled"
          >
            <el-option
              v-for="item in state.filedKindList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="提示 :" prop="titTop">
          <el-input
            v-model="state.defineData.titTop"
            style="width: 300px"
            :maxlength="60"
          />
        </el-form-item>
        <el-form-item label="默认内容 :" prop="defaultContent">
          <div v-if="state.flag === '下拉列表'">
            <el-select
              v-model="state.defineData.defaultContent"
              style="width: 300px"
            >
              <el-option
                v-for="op in state.selectOption"
                :key="op.value"
                :label="op.label"
                :value="op.value"
              />
            </el-select>
          </div>
          <div
            v-if="
              state.flag === '日期' &&
              state.defineData.selectFormat === 'YYYY/MM/DD'
            "
          >
            <el-date-picker
              v-model="state.defineData.defaultContent"
              value-format="YYYY-MM-DD"
              :style="{ width: '300px' }"
            />
          </div>
          <div
            v-if="
              state.flag === '日期' &&
              state.defineData.selectFormat === 'YYYY/MM/DD hh:mm'
            "
          >
            <el-date-picker
              type="datetime"
              v-model="state.defineData.defaultContent"
              value-format="YYYY-MM-DD hh:mm"
              format="YYYY-MM-DD hh:mm"
              :style="{ width: '300px' }"
            />
          </div>
          <div
            v-if="
              state.flag === '日期' &&
              state.defineData.selectFormat === 'YYYY/MM/DD hh:mm:ss'
            "
          >
            <el-date-picker
              type="datetime"
              v-model="state.defineData.defaultContent"
              value-format="YYYY-MM-DD hh:mm:ss"
              format="YYYY-MM-DD hh:mm:ss"
              :style="{ width: '300px' }"
            />
          </div>
          <div
            v-if="state.flag === '时间' && state.defineData.Time === 'hh:mm'"
          >
            <el-time-picker
              v-model="state.defineData.defaultContent"
              arrow-control
              placeholder="请输入"
              format="hh:mm"
              style="width: 300px"
            />
          </div>
          <div
            v-if="state.flag === '时间' && state.defineData.Time === 'hh:mm:ss'"
          >
            <el-time-picker
              v-model="state.defineData.defaultContent"
              arrow-control
              placeholder="请输入"
              format="hh:mm:ss"
              value-format="hh:mm:ss"
              style="width: 300px"
            />
          </div>
          <div
            v-if="
              state.flag === '起止时间' &&
              state.defineData.selectFormat === 'YYYY/MM/DD'
            "
          >
            <el-date-picker
              v-model="state.defineData.defaultContent"
              type="daterange"
              range-separator="至"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              value-format="YYYY/MM/DD"
              style="width: 300px"
            />
          </div>
          <div
            v-if="
              state.flag === '起止时间' &&
              state.defineData.selectFormat === 'YYYY/MM/DD hh:mm'
            "
          >
            <el-date-picker
              v-model="state.defineData.defaultContent"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              format="YYYY/MM/DD hh:mm"
              value-format="YYYY/MM/DD hh:mm"
              end-placeholder="结束时间"
              :style="{ width: '300px' }"
            />
          </div>
          <div
            v-if="
              state.flag === '起止时间' &&
              state.defineData.selectFormat === 'YYYY/MM/DD hh:mm:ss'
            "
          >
            <el-date-picker
              v-model="state.defineData.defaultContent"
              type="datetimerange"
              range-separator="至"
              start-placeholder="开始时间"
              format="YYYY/MM/DD hh:mm:ss"
              value-format="YYYY/MM/DD hh:mm:ss"
              end-placeholder="结束时间"
              :style="{ width: '300px' }"
            />
          </div>
          <div
            v-if="
              ['文本域', '数字', '关联已有数据', '公式计算', ''].includes(
                state.flag,
              )
            "
          >
            <el-input
              v-model="state.defineData.defaultContent"
              style="width: 300px"
            />
          </div>
        </el-form-item>
        <el-form-item label="字段选择 :" prop="filedCheck">
          <el-checkbox-group v-model="state.defineData.filedCheck">
            <el-checkbox label="必填" name="1" />
            <el-checkbox label="可编辑" name="2" />
          </el-checkbox-group>
        </el-form-item>
        <!-- 通过下拉框展示不同的页面 -->
        <div class="defineTypeInput" v-if="state.flag">
          <!-- 文本域 -->
          <div v-if="state.flag === '文本域'">
            <el-form-item
              label="正则校验"
              prop="regularCheck"
              class="myTextArea"
            >
              <el-input
                type="textarea"
                v-model="state.defineData.regularCheck"
                style="width: 300px"
              />
            </el-form-item>
            <el-form-item
              label="限制字数"
              prop="limitNumberWords"
              class="limitNumberWords"
            >
              <el-checkbox v-model="state.defineData.limitNumberWords" />
              <el-input-number
                v-model="limitNumberWordsNum"
                v-if="state.defineData.limitNumberWords"
                style="height: 34px, margin-left: 10px"
                :min="1"
                class="mx-4"
                size="small"
              />
            </el-form-item>
            <el-form-item
              label="不允许重复"
              prop="duplicateNot"
              class="duplicateNot"
            >
              <el-checkbox v-model="state.defineData.duplicateNot" />
            </el-form-item>
          </div>
          <!-- 数字 -->
          <div v-if="state.flag === '数字'">
            <el-form-item label="单位 : " prop="Company">
              <el-input
                v-model="state.defineData.company"
                style="width: 300px"
              ></el-input>
            </el-form-item>
            <el-form-item
              label="允许小数 : "
              prop="DecimalAll"
              class="DecimalAll"
            >
              <el-checkbox v-model="state.defineData.DecimalAll" />
              <p v-if="state.defineData.DecimalAll">
                限制
                <el-input
                  v-model="state.defineData.LimitDigit"
                  style="width: 60px"
                ></el-input
                >位
              </p>
            </el-form-item>
            <el-form-item
              label="数字范围 : "
              prop="DecimalLimitRange"
              class="DecimalLimitRange"
            >
              <el-checkbox v-model="state.defineData.DecimalLimitRange" />
              <p v-if="state.defineData.DecimalLimitRange">
                <el-input
                  style="width: 50px"
                  v-model="state.defineData.regularCheckMin"
                ></el-input>
                ~
                <el-input
                  style="width: 50px"
                  v-model="state.defineData.regularCheckBig"
                ></el-input>
              </p>
            </el-form-item>
          </div>
          <!-- 下拉列表 -->
          <div v-if="state.flag === '下拉列表'">
            <!-- <span> 选项内容</span> -->
            <draggable
              v-model="state.defineData.domains"
              animation="300"
              dragClass="dragClass"
              ghostClass="ghostClass"
              chosenClass="chosenClass"
              handle=".move-icon"
              @end="endMove"
            >
              <template #item="{ element }">
                <el-form-item
                  :key="element.value"
                  label="选项内容 : "
                  :prop="element.label"
                  :rules="{
                    required: true,
                    message: '请输入选项',
                    trigger: 'blur',
                  }"
                >
                  <el-input
                    v-model="element.label"
                    style="width: 275px"
                    @change="addChangeInput"
                  />
                  <div v-if="state.isShowOperation">
                    <span class="move-icon icon">
                      <svg-icon icon-class="move"></svg-icon>
                    </span>
                    <span class="del-icon icon"
                      ><svg-icon
                        icon-class="del"
                        @click.prevent="removeDomain(element)"
                      ></svg-icon
                    ></span>
                  </div>
                </el-form-item>
              </template>
            </draggable>
            <el-form-item>
              <span class="del-icon icon"
                ><svg-icon icon-class="addFrom"></svg-icon
              ></span>
              <el-button type="text" @click="addDomain">添加新选项</el-button>
              <!-- <el-button @click="addDomain"></el-button> -->
            </el-form-item>
          </div>
          <!-- 日期 起止时间-->
          <div v-if="['日期', '起止时间'].includes(state.flag)">
            <el-form-item label="显示格式" prop="selectFormat">
              <el-select
                v-model="state.defineData.selectFormat"
                style="width: 300px"
              >
                <el-option
                  v-for="op in state.selectFormatData"
                  :key="op.value"
                  :label="op.label"
                  :value="op.value"
                />
              </el-select>
            </el-form-item>
          </div>
          <!-- 时间 -->
          <div v-if="state.flag === '时间'">
            <el-form-item label="显示格式" prop="Time">
              <el-select v-model="state.defineData.Time">
                <el-option label="hh:mm" value="hh:mm"></el-option>
                <el-option label="hh:mm:ss" value="hh:mm:ss"></el-option>
              </el-select>
            </el-form-item>
          </div>
          <!-- 关联已有数据 -->
          <div v-if="state.flag === '关联已有数据'">
            <el-form-item label="关联已有数据" prop="aeData">
              <el-select v-model="state.defineData.aeData">
                <el-option
                  v-for="item in state.AssociateData"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item prop="modelFiled">
              <el-select v-model="state.defineData.aeData">
                <el-option
                  v-for="item in state.AssociateData"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
          </div>
          <!-- 公式 -->
          <div v-if="state.flag === '公式计算'">
            <el-form-item
              label="字段设置"
              prop="AssociateModel"
              class="AssociateModel"
            >
              <el-select
                v-model="state.defineData.AssociateModel"
                style="width: 300px"
                placeholder="选择关联模型"
              >
                <el-option
                  v-for="item in state.AssociateData"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item prop="AssociatedFiled">
              <el-select
                v-model="state.defineData.AssociatedFiled"
                style="width: 300px"
                placeholder="选择关联字段"
              >
                <el-option
                  v-for="item in state.AssociateData"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                ></el-option>
              </el-select>
            </el-form-item>
            <el-form-item prop="AssociatedSum">
              <el-select
                v-model="state.defineData.AssociatedSum"
                style="width: 300px"
              >
                <el-option label="SUM" value="SUM"></el-option>
              </el-select>
            </el-form-item>
            <el-form-item
              label="限制实例范围"
              prop="AssociatedInstance"
              class="AssociatedInstance"
            >
              <el-radio v-model="state.defineData.AssociatedInstance" label="1"
                >是</el-radio
              >
              <el-radio v-model="state.defineData.AssociatedInstance" label="2"
                >否</el-radio
              >
            </el-form-item>
          </div>
        </div>
      </el-form>
    </el-drawer>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  reactive,
  ref,
  toRefs,
  onUpdated,
  nextTick,
} from 'vue'
import type { ElForm } from 'element-plus'
import draggable from 'vuedraggable'
type FormInstance = InstanceType<typeof ElForm>

import {
  commFiled,
  determineConfig,
  cancelConfig,
  defineDataConfig,
  filedKindConfig,
  rules,
} from './config/edit-config'
import publicButton from '@/components/Button/index.vue'

export default defineComponent({
  props: {
    IsDrawer: {
      type: Boolean,
      default: false,
    },
    isEmit: {
      type: Boolean,
      default: false,
    },
    pointFlag: {
      type: String,
      default: '',
    },
    //是否标签
    isTag: {
      type: Boolean,
      default: false,
    },
    title: {
      type: String,
      default: '新建字段',
    },
  },
  components: {
    // publicForm,
    publicButton,
    draggable,
  },
  setup(props, content) {
    let {
      IsDrawer: isShow,
      isEmit: isDisabled,
      isTag: tag,
      title: drawerTitle,
    } = toRefs(props)
    const state = reactive({
      drawer: false,
      form: commFiled,
      flag: '',
      defineData: defineDataConfig,
      selectFormatData: [
        {
          label: 'YYYY/MM/DD',
          value: 'YYYY/MM/DD',
        },
        {
          label: 'YYYY/MM/DD hh:mm',
          value: 'YYYY/MM/DD hh:mm',
        },
        {
          label: 'YYYY/MM/DD hh:mm:ss',
          value: 'YYYY/MM/DD hh:mm:ss',
        },
      ],
      //关联数据源
      AssociateData: [
        {
          label: '测试1',
          value: 1,
        },
        {
          label: '测试2',
          value: 2,
        },
      ],
      selectOption: [],
      filedKindList: filedKindConfig, // 字段类型下拉数据
      rulesConfig: rules, // 点击是否必填
      selectDefaultContent: '',
      isShowOperation: false,
    })
    const limitNumberWordsNum = ref(2)
    const determineButton = ref(determineConfig)
    const cancelButton = ref(cancelConfig)
    const FormRefDefine = ref<FormInstance>()

    onUpdated(() => {
      nextTick(() => {
        ;(state as any).flag = props.pointFlag
      })
    })

    // 结束拖动
    const endMove = (e) => {
      // console.log(e)
      // console.log(state.defineData.domains[e.oldIndex])
      // console.log(state.defineData.domains[e.newIndex])
      // console.log(state.defineData.domains, '我是数据')
      // swapArray(e.oldIndex, e.newIndex)
    }
    // 切换下拉框进行数据清空
    const change = (e) => {
      state.flag = e
      Object.keys(defineDataConfig).forEach((key) => {
        if (!['id', 'title', 'titTop', 'filedKind'].includes(key)) {
          if (
            [
              'limitNumberWords',
              'duplicateNot',
              'DecimalAll',
              'DecimalLimitRange',
            ].includes(key)
          ) {
            defineDataConfig[key] = false
          } else if (key === 'limitNumberWordsNum') {
            defineDataConfig[key] = 1
          } else if (key === 'selectFormat') {
            defineDataConfig[key] = 'YYYY/MM/DD'
          } else if (key === 'Time') {
            defineDataConfig[key] = 'hh:mm'
          } else if (key === 'filedCheck') {
            defineDataConfig[key] = ['可编辑']
          } else if (key === 'domains') {
            defineDataConfig[key] = [{ label: '', value: Date.now() }]
          } else {
            defineDataConfig[key] = ''
          }
        }
      })
    }

    // 删除input
    const removeDomain = (item) => {
      const index = state.defineData.domains.indexOf(item)
      if (index !== -1) {
        state.defineData.domains.splice(index, 1)
      }
      if (state.defineData.domains && state.defineData.domains.length == 1) {
        state.isShowOperation = false
      }
      const data = JSON.parse(JSON.stringify(state.defineData.domains))
      state.selectOption = data
    }

    // 添加input
    const addDomain = () => {
      state.defineData.domains.push({
        value: Date.now(),
        label: '',
      })
      if (state.defineData.domains && state.defineData.domains.length > 1) {
        state.isShowOperation = true
      }
    }

    // 确定
    const onButtonOk = () => {
      console.log(FormRefDefine.value)
      const obj = JSON.parse(JSON.stringify(state.defineData))
      if (!FormRefDefine.value) return
      FormRefDefine.value.validate((valid) => {
        if (valid) {
          // console.log('点击了---《')
          content.emit('defineData', obj)
          FormRefDefine.value && FormRefDefine.value.resetFields()
          clearDefault()
        } else {
          return false
        }
      })
    }

    // 取消
    const onButtonCancel = () => {
      // console.log('取消了')
      content.emit('IsDrawer')
      FormRefDefine.value && FormRefDefine.value.resetFields()
      clearDefault()
      // state.flag = ''
    }
    // 清空默认值
    const clearDefault = () => {
      Object.keys(defineDataConfig).forEach((key) => {
        if (
          [
            'limitNumberWords',
            'duplicateNot',
            'DecimalAll',
            'DecimalLimitRange',
          ].includes(key)
        ) {
          defineDataConfig[key] = false
        } else if (key === 'limitNumberWordsNum') {
          defineDataConfig[key] = 1
        } else if (key === 'selectFormat') {
          defineDataConfig[key] = 'YYYY/MM/DD'
        } else if (key === 'Time') {
          defineDataConfig[key] = 'hh:mm'
        } else if (key === 'filedCheck') {
          defineDataConfig[key] = ['可编辑']
        } else if (key === 'domains') {
          defineDataConfig[key] = [{ label: '', value: Date.now() }]
        } else {
          defineDataConfig[key] = ''
        }
      })
    }
    // 输入文本框事件
    const addChangeInput = () => {
      state.selectOption = []
      const data = JSON.parse(JSON.stringify(state.defineData.domains))
      state.selectOption = data
    }

    return {
      state,
      limitNumberWordsNum,
      determineButton,
      cancelButton,
      FormRefDefine,
      isShow,
      isDisabled,
      tag,
      drawerTitle,
      endMove,
      change,
      removeDomain,
      addDomain,
      onButtonOk,
      onButtonCancel,
      addChangeInput,
    }
  },
})
</script>

<style lang="scss" scoped>
.newFiled {
  .el-form-item__content {
    flex-wrap: nowrap;
    .el-tooltip__trigger {
      // margin-left: 10px; 
      // font-size: 16px;
    }
    justify-content: center;
    span {
      color: #ffffff;
    }
  }
  .icon {
    // margin-left: 10px;
    cursor: pointer;
  }
  .footerButton {
    display: flex;
    justify-content: flex-end;
    margin-top: 20px;
    div {
      margin-right: 20px;
    }
  }
}
.el-tooltip__trigger {
  // margin-left: 10px;
  margin-top: 5px;
}
:deep(.limitNumberWords) {
  margin-left: 64px;
}
:deep(.limitNumberWords .el-form-item__content .el-checkbox) {
  margin-right: 20px;
}
:deep(.limitNumberWords .el-input__inner) {
  height: 35px;
}
:deep(.myTextArea) {
  display: block;
  margin-left: 64px;
}
:deep(.AssociateModel) {
  display: block;
  margin-left: 64px;
}
:deep(.AssociateModel .el-form-item__label-wrap) {
  height: 30px;
}
:deep(.myTextArea .el-form-item__label-wrap) {
  height: 30px;
}
:deep(.AssociateModel .el-textarea__inner) {
  margin-left: 15px;
  height: 90px;
}
:deep(.myTextArea .el-textarea__inner) {
  margin-left: 15px;
  height: 90px;
}
:deep(.duplicateNot) {
  margin-left: 70px;
}
:deep(.DecimalAll .el-checkbox) {
  margin-right: 20px;
}
:deep(.AssociateModel .el-form-item__content) {
  padding-left: 28px;
}
:deep(.DecimalLimitRange .el-checkbox) {
  margin-right: 20px;
}
:deep(.AssociatedInstance) {
  margin-left: 93px;
}
</style>
