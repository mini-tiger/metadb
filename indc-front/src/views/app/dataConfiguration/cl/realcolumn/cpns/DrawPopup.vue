<template>
  <div class="Draw">
    <el-drawer
      v-model="drawer"
      title="新建关联"
      :destroy-on-close="true"
      :size="460"
    >
      <template #default>
        <div v-if="state.steps === 1">
          <!-- <publicFrom :data="state.form">
            <template v-slot="data">
              <el-button @click="canCel">取消</el-button>
              <el-button
                @click="nextStep(data.data.formRef, data.data.formData)"
                >下一步</el-button
              >
            </template>
          </publicFrom> -->
          <el-form
            ref="ruleFormRef"
            class="ruleFormRef"
            :model="state.form.searchData"
            :rules="state.form.rules"
          >
            <el-form-item label="模型分类" prop="modelKind">
              <ZHDropdownTree
                :data="data1"
                v-model="state.form.searchData.modelKind"
                width="300px"
                @update:modelValue="DropTree"
              ></ZHDropdownTree>
            </el-form-item>
            <el-form-item label="模型名称" prop="modelKind">
              <el-select
                v-model="state.form.searchData.modelName"
                style="width: 300px"
              >
                <el-option
                  v-for="item in state.form.option"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
            <el-form-item class="isSlot">
              <ZHButton :options="state.option" @ZHBtnClick="canCel"
                >取消</ZHButton
              >
              <ZHButton
                :options="state.option1"
                @ZHBtnClick="nextStep(ruleFormRef)"
                >下一步</ZHButton
              >
            </el-form-item>
          </el-form>
        </div>
        <div v-else-if="state.steps === 2" class="tabs">
          <el-tabs
            v-model="activeName"
            class="demo-tabs"
            @tab-click="handleClick"
          >
            <el-tab-pane label="模型字段" name="ModelField">
              <div class="modelField">
                <div
                  class="modelFrom"
                  v-for="(item, index) in state.formData"
                  :key="index"
                >
                  <div class="modelFromHeader">{{ item.bk_group_name }}</div>
                  <el-form
                    :model="item.list"
                    :inline="true"
                    label-position="right"
                    label-width="120px"
                  >
                    <template v-for="(itm, ind) in item.data" :key="ind">
                      <el-form-item
                        :label="itm.bk_property_name + ': '"
                        :prop="itm.bk_property_id"
                      >
                        <el-input
                          v-model="item.list[itm.bk_property_id]"
                          style="width: 300px"
                        />
                      </el-form-item>
                    </template>
                  </el-form>
                </div>
              </div>
            </el-tab-pane>
            <el-tab-pane label="测点字段" name="MeasuringField">
              <publicTable :data="state" />
            </el-tab-pane>
          </el-tabs>
        </div>
      </template>
      <template #footer>
        <div class="footerButton" v-if="state.steps === 2">
          <!-- <el-button>取消</el-button>
          <el-button @click="determine">确定</el-button> -->
          <publicButton
            :data="determineOk.buttonOk"
            class="determineOk"
            @click="buttonOk"
          />
          <publicButton
            :data="cancelButton.buttonOk"
            class="determineOk"
            @click="canCel"
          />
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import type { ElForm } from 'element-plus'
type FormInstance = InstanceType<typeof ElForm>

import publicButton from '@/components/Button/index.vue'
// import publicFrom from '@/components/Form/index.vue'
import publicTable from '@/components/Table/index.vue'
import { table, determineConfig, cancelConfig } from '../config/real-config'
import { dataProcessEdit } from '../config/data-processing'
import ZHButton from '@/components/Button/Button.vue'

export default defineComponent({
  name: 'DrawPopup',
  props: ['data'],
  components: {
    // publicFrom,
    ZHButton,
    publicTable,
    publicButton,
  },
  setup(props, content) {
    const state = reactive({
      form: {
        searchData: {
          modelKind: '',
          modelName: '',
        },
        option: [
          {
            label: '自定义',
            value: '1',
          },
          {
            label: '是不是',
            value: '2',
          },
        ],
        rules: {
          modelKind: [
            {
              required: true,
              message: '请输入模型分类',
              trigger: 'blur',
            },
          ],
          modelName: [
            {
              required: true,
              message: '请输入模型名称',
              trigger: 'blur',
            },
          ],
        },
      },
      steps: 1,
      table: {
        columns: table.columns,
        tableData: table.tableData,
        isPagination: table.tableData, // 是否展示分页
        pagination: table.pagination,
        tableWidth: '100%', // 表格的宽度
        auxiliary: false,
      },
      formData: [],
      formList: {},
      option: {
        bgColor: '#E4E4FC',
        textColor: '#595AEC',
      },
      option1: {
        bgColor: '#595AEC',
        textColor: '#fff',
      },
    })
    const drawer = ref(false) // 控制抽屉开关
    const modelKind = ref('')
    const activeName = ref('ModelField')
    const determineOk = ref(determineConfig)
    const cancelButton = ref(cancelConfig)
    const ruleFormRef = ref<FormInstance>()
    onMounted(() => {
      drawer.value = props.data
      state.formData = dataProcessEdit()
      console.log(state.formData, '---->')
    })
    // 下一步
    const nextStep = async (fuc) => {
      console.log(fuc)
      if (!fuc) return
      await fuc.validate((valid) => {
        if (valid) {
          console.log('submit!')
          const draw = document.querySelector<HTMLElement>('.el-drawer')
          console.log(draw, '-->')
          draw.style.cssText = 'width: 940px !important'
          state.steps = 2
        } else {
          // console.log('error submit!', fields)
        }
      })
    }
    // 取消
    const canCel = () => {
      content.emit('IsDrawer', false)
    }
    // tabs切换
    const handleClick = (tab: string, event: Event) => {
      console.log(tab, event)
    }
    // 提交
    const onSubmit = async (fuc, data) => {
      if (!fuc) return
      await fuc.validate((valid, fields) => {
        if (valid) {
          console.log('submit!')
          content.emit('IsDrawer')
        } else {
          // console.log('error submit!', fields)
        }
      })
    }
    const determine = () => {
      console.log(state.formData)
    }
    const data1 = [
      {
        id: 1,
        label: 'Level one 1',
        parentId: 1,
        children: [
          {
            id: 4,
            label: 'Level two 1-1',
            parentId: 1,
            children: [
              {
                id: 9,
                label: 'Level three 1-1-1',
                parentId: 1,
              },
              {
                id: 10,
                label: 'Level three 1-1-2',
                parentId: 1,
              },
            ],
          },
        ],
      },
      {
        id: 2,
        label: 'Level one 2',
        parentId: 2,
        children: [
          {
            id: 5,
            label: 'Level two 2-1',
          },
          {
            id: 6,
            label: 'Level two 2-2',
          },
        ],
      },
      {
        id: 3,
        label: 'Level one 3',
        children: [
          {
            id: 7,
            label: 'Level two 3-1',
          },
          {
            id: 8,
            label: 'Level two 3-2',
          },
        ],
      },
    ]
    const buttonOk = () => {
      content.emit('IsDrawer', false)
    }
    // 获取数据
    const DropTree = (data) => {
      console.log(data, 'data')
    }
    return {
      drawer,
      state,
      data1,
      activeName,
      determineOk,
      cancelButton,
      modelKind,
      ruleFormRef,
      buttonOk,
      DropTree,
      nextStep,
      canCel,
      handleClick,
      onSubmit,
      determine,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/assets/style/overallStyle.scss';
:deep(.el-drawer__body) {
  padding: 0;
}
.Draw {
  position: relative;
}
:deep(.isSlot .el-form-item__content) {
  position: absolute;
  bottom: 20px;
  right: 20px;
}
.modelFrom {
  margin-top: 10px;
  border-radius: 4px;
  border: 1px solid #ededfd;
}
:deep(.el-drawer__body) {
  padding: 10px !important;
}
:deep(.el-tabs__content) {
  padding: 0 !important;
}

.modelFromHeader {
  height: 40px;
  border-bottom: 1px solid #ededfd;
  margin-bottom: 20px;
  box-sizing: border-box;
  padding-left: 20px;
  line-height: 40px;
  font-size: 14px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #302e6d;
}
.footerButton {
  display: flex;
  justify-content: flex-end;
  line-height: 60px;
}
.determineOk {
  margin-right: 20px;
}
.ruleFormRef {
  padding: 10px 0 0 10px;
}
</style>
