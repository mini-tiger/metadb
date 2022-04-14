<template>
  <div class="from">
    <div v-if="form.formArray">
      <el-collapse v-model="activeName">
        <!-- <el-collapse-item title="" name="1"> -->
        <el-config-provider :locale="locale">
          <el-form
            ref="formRef"
            :model="form.searchData"
            label-position="right"
            label-width="auto"
            :rules="form.rules"
            class="formRef"
          >
            <el-row>
              <template v-for="(item, index) in form.searchForm" :key="index">
                <el-col :span="item.span ? item.span : form.span">
                  <el-form-item
                    :label="item.label + ' :'"
                    :prop="item.prop"
                    v-if="item.type !== 'SelectAndInput'"
                  >
                    <!-- 文本框 -->
                    <el-input
                      v-if="item.type === 'Input'"
                      v-model="form.searchData[item.prop]"
                      :placeholder="'请输入'"
                      :style="{ width: item.width ? item.width : '180px' }"
                    />
                    <!-- 带icon的文本框 -->
                    <el-input
                      v-if="item.type === 'InputAndIcon'"
                      v-model="form.searchData[item.prop]"
                      :placeholder="'请输入'"
                      :suffix-icon="item.icon"
                      :style="{ width: item.width ? item.width : '180px' }"
                    />
                    <!-- 自定义的icon图标 -->
                    <el-input
                      v-if="item.type === 'CustomInputAndIcon'"
                      v-model="form.searchData[item.prop]"
                      :placeholder="'请输入'"
                      class="CustomInputAndIcon"
                      :style="{ width: item.width ? item.width : '180px' }"
                    >
                      <template #suffix v-if="item.icon">
                        <svg-icon :icon-class="item.icon" />
                      </template>
                    </el-input>
                    <!--  自定义带单位文本框 -->
                    <el-input
                      v-if="item.type === 'CompanyInput'"
                      v-model="form.searchData[item.prop]"
                      placeholder="请输入"
                      class="CompanyInput"
                      :style="{ width: item.width ? item.width : '180px' }"
                    >
                      <template #append>{{ item.CompanyInput }}</template>
                    </el-input>
                    <!-- 多选 -->
                    <el-select
                      v-if="item.type === 'Select'"
                      v-model="form.searchData[item.prop]"
                      class="m-2"
                      :placeholder="'请输入'"
                      :style="{ width: item.width ? item.width : '180px' }"
                      @change="
                        item.change
                          ? item.change(form.searchData[item.prop])
                          : null
                      "
                    >
                      <el-option
                        v-for="op in item.options"
                        :key="op.value"
                        :label="op.label"
                        :value="op.value"
                      />
                    </el-select>
                    <!-- 单选 -->
                    <el-radio-group
                      v-if="item.type === 'Radio'"
                      v-model="form.searchData[item.prop]"
                      :style="{ width: item.width ? item.width : '180px' }"
                    >
                      <el-radio
                        v-for="ra in item.radios"
                        :key="ra.value"
                        :label="ra.value"
                        >{{ ra.label }}</el-radio
                      >
                    </el-radio-group>
                    <!-- 日期框 -->
                    <el-date-picker
                      v-if="item.type === 'Date'"
                      v-model="form.searchData[item.prop]"
                      value-format="YYYY-MM-DD"
                      :style="{ width: item.width ? item.width : '180px' }"
                    />
                    <!-- 日期范围-->
                    <el-date-picker
                      v-if="item.type === 'Daterange'"
                      v-model="form.searchData[item.prop]"
                      type="daterange"
                      range-separator="至"
                      start-placeholder="开始时间"
                      end-placeholder="结束时间"
                      value-format="h:mm:ss"
                      :style="{ width: item.width ? item.width : '180px' }"
                    />
                    <!-- 时间框 -->
                    <el-time-picker
                      v-if="item.type === 'Time'"
                      v-model="form.searchData[item.prop]"
                      placeholder="请输入"
                      value-format="h:mm:ss"
                      :style="{ width: item.width ? item.width : '180px' }"
                    ></el-time-picker>
                    <!-- 多选 -->
                    <el-checkbox-group
                      v-if="item.type === 'Checkbox'"
                      v-model="form.searchData[item.prop]"
                      :style="{ width: item.width ? item.width : '180px' }"
                    >
                      <el-checkbox
                        v-for="ch in item.option"
                        :key="ch.value"
                        :label="ch.label"
                        >{{ ch.label }}</el-checkbox
                      >
                    </el-checkbox-group>
                    <!-- 滑块 -->
                    <el-switch
                      v-if="item.type === 'Switch'"
                      v-model="form.searchData[item.prop]"
                    ></el-switch>
                    <!-- 文本域 -->
                    <el-input
                      v-if="item.type === 'Textarea'"
                      v-model="form.searchData[item.prop]"
                      type="textarea"
                      :style="{ width: item.width ? item.width : '180px' }"
                    ></el-input>
                    <!-- 搜索 -->
                    <el-select
                      v-if="item.type === 'Search'"
                      v-model="form.searchData[item.prop]"
                      filterable
                      clearable
                      :multiple="item.isMultiple ? item.isMultiple : false"
                      placeholder="请选择"
                      :style="{ width: item.width ? item.width : '180px' }"
                    >
                      <el-option
                        v-for="ite in item.options"
                        :key="ite.value"
                        :label="ite.label"
                        :value="ite.value"
                      >
                      </el-option>
                    </el-select>
                  </el-form-item>

                  <el-form-item
                    v-if="item.type === 'SelectAndInput'"
                    id="SelectAndInput"
                    :label-width="item.label ? '80px' : '0'"
                    :style="item.label ? '0' : '300px'"
                    :label="item.label ? item.label + ' :' : ''"
                  >
                    <el-select
                      v-model="form.searchData[item.prop]"
                      filterable
                      clearable
                      :multiple="item.isMultiple ? item.isMultiple : false"
                      placeholder="请选择"
                    >
                      <el-option
                        v-for="ite in item.options"
                        :key="ite.value"
                        :label="ite.label"
                        :value="ite.value"
                      />
                    </el-select>
                    <el-input
                      v-model="form.searchData[item.value]"
                      :placeholder="'请输入'"
                      class="myInput"
                      :style="{ width: item.width ? item.width : '180px' }"
                    />
                  </el-form-item>
                </el-col>
              </template>
              <el-col :span="5">
                <el-form-item
                  v-if="form.formSlot"
                  class="fromSlot"
                  label-width="0"
                >
                  <slot
                    :data="{ formRef: formRef, formData: form.searchData }"
                  ></slot>
                </el-form-item>
              </el-col>
            </el-row>
          </el-form>
        </el-config-provider>
        <!-- </el-collapse-item> -->
      </el-collapse>
    </div>
    <div v-else>
      <el-config-provider :locale="locale">
        <el-form
          ref="formRef"
          :model="form.searchData"
          label-position="right"
          label-width="auto"
          :rules="form.rules"
          class="formRef"
        >
          <template v-for="(item, index) in form.searchForm" :key="index">
            <el-form-item
              :label="item.label + ' :'"
              :prop="item.prop"
              v-if="item.type !== 'SelectAndInput'"
            >
              <!-- 文本框 -->
              <el-input
                v-if="item.type === 'Input'"
                v-model="form.searchData[item.prop]"
                :placeholder="'请输入'"
                :style="{ width: item.width ? item.width : '180px' }"
                :disabled="item.disabled ? item.disabled : null"
              />
              <!-- 带icon的文本框 -->
              <el-input
                v-if="item.type === 'InputAndIcon'"
                v-model="form.searchData[item.prop]"
                :placeholder="'请输入'"
                :suffix-icon="item.icon"
                :style="{ width: item.width ? item.width : '180px' }"
              />
              <!-- 自定义的icon图标 -->
              <el-input
                v-if="item.type === 'CustomInputAndIcon'"
                v-model="form.searchData[item.prop]"
                :placeholder="'请输入'"
                class="CustomInputAndIcon"
                :style="{ width: item.width ? item.width : '180px' }"
              >
                <template #suffix v-if="item.icon">
                  <svg-icon :icon-class="item.icon" />
                </template>
              </el-input>
              <!--  自定义带单位文本框 -->
              <el-input
                v-if="item.type === 'CompanyInput'"
                v-model="form.searchData[item.prop]"
                placeholder="请输入"
                class="CompanyInput"
                :style="{ width: item.width ? item.width : '180px' }"
              >
                <template #append>{{ item.CompanyInput }}</template>
              </el-input>
              <!-- 多选 -->
              <el-select
                v-if="item.type === 'Select'"
                v-model="form.searchData[item.prop]"
                class="m-2"
                :disabled="item.disabled ? item.disabled : null"
                :placeholder="'请输入'"
                :style="{ width: item.width ? item.width : '180px' }"
                @change="
                  item.change
                    ? item.change(form.searchData[item.prop], formRef)
                    : null
                "
              >
                <el-option
                  v-for="op in item.options"
                  :key="op.value"
                  :label="op.label"
                  :value="op.value"
                />
              </el-select>
              <!-- 单选 -->
              <el-radio-group
                v-if="item.type === 'Radio'"
                v-model="form.searchData[item.prop]"
                :style="{ width: item.width ? item.width : '180px' }"
              >
                <el-radio
                  v-for="ra in item.radios"
                  :key="ra.value"
                  :label="ra.value"
                  >{{ ra.label }}</el-radio
                >
              </el-radio-group>
              <!-- 日期框 -->
              <el-date-picker
                v-if="item.type === 'Date'"
                v-model="form.searchData[item.prop]"
                value-format="YYYY-MM-DD"
                :style="{ width: item.width ? item.width : '180px' }"
              />
              <!-- 时间日期框 -->
              <el-date-picker
                v-if="item.type === 'datetime'"
                type="datetime"
                v-model="form.searchData[item.prop]"
                value-format="YYYY-MM-DD"
                :format="item.formatType ? item.formatType : 'YYYY/MM/DD'"
                :style="{ width: item.width ? item.width : '180px' }"
              />
              <!-- 时间日期范围框 -->
              <el-date-picker
                v-if="item.type === 'datetimerange'"
                v-model="form.searchData[item.prop]"
                type="datetimerange"
                range-separator="至"
                start-placeholder="开始时间"
                :format="item.formatType ? item.formatType : 'YYYY/MM/DD'"
                end-placeholder="结束时间"
                :style="{ width: item.width ? item.width : '180px' }"
              />
              <!-- 日期范围-->
              <el-date-picker
                v-if="item.type === 'Daterange'"
                v-model="form.searchData[item.prop]"
                type="daterange"
                range-separator="至"
                start-placeholder="开始时间"
                end-placeholder="结束时间"
                value-format="h:mm:ss"
                :style="{ width: item.width ? item.width : '180px' }"
              />
              <!-- 时间框 -->
              <el-time-picker
                v-if="item.type === 'Time'"
                v-model="form.searchData[item.prop]"
                placeholder="请输入"
                value-format="hh:mm:ss"
                :format="item.formatType ? item.formatType : 'hh:mm'"
                :style="{ width: item.width ? item.width : '180px' }"
              ></el-time-picker>
              <!-- 多选 -->
              <el-checkbox-group
                v-if="item.type === 'Checkbox'"
                v-model="form.searchData[item.prop]"
                :style="{ width: item.width ? item.width : '180px' }"
              >
                <el-checkbox
                  v-for="ch in item.option"
                  :key="ch.value"
                  :label="ch.label"
                  >{{ ch.label }}</el-checkbox
                >
              </el-checkbox-group>
              <!-- 滑块 -->
              <el-switch
                v-if="item.type === 'Switch'"
                v-model="form.searchData[item.prop]"
              ></el-switch>
              <!-- 文本域 -->
              <el-input
                v-if="item.type === 'Textarea'"
                v-model="form.searchData[item.prop]"
                type="textarea"
                :style="{ width: item.width ? item.width : '180px' }"
              ></el-input>
              <!-- 搜索 -->
              <el-select
                v-if="item.type === 'Search'"
                v-model="form.searchData[item.prop]"
                filterable
                clearable
                :multiple="item.isMultiple ? item.isMultiple : false"
                placeholder="请选择"
                :style="{ width: item.width ? item.width : '180px' }"
              >
                <el-option
                  v-for="ite in item.options"
                  :key="ite.value"
                  :label="ite.label"
                  :value="ite.value"
                >
                </el-option>
              </el-select>
            </el-form-item>
          </template>
          <el-form-item v-if="form.formSlot" class="fromSlot">
            <slot
              :data="{ formRef: formRef, formData: form.searchData }"
            ></slot>
          </el-form-item>
        </el-form>
      </el-config-provider>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue'
import type { ElForm } from 'element-plus'
import { ElConfigProvider } from 'element-plus'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
type FormInstance = InstanceType<typeof ElForm>

export default defineComponent({
  props: ['data'],
  components: {
    ElConfigProvider,
  },
  setup(props) {
    const formRef = ref<FormInstance>()
    const activeName = ref('1')
    const form = reactive({
      searchData: {},
      searchForm: [],
      labelWidth: '120px',
      rules: {},
      formSlot: false,
      formArray: true,
      span: 6,
    })

    onMounted(() => {
      setTimeout(() => {
        if (props.data) {
          for (let key in props.data) {
            if (key) {
              form[key] = props.data[key]
            }
          }
        }
      }, 0)
      // isIconShow()
    })

    const isIconShow = () => {
      const icon = document.querySelector<HTMLElement>(
        '.el-collapse-item__arrow',
      )
      if (
        props.data.searchForm &&
        props.data.searchForm.length > 7 &&
        form.formArray
      ) {
        icon.style.display = 'block'
      } else {
        icon.style.display = 'none'
      }
    }

    return {
      locale: zhCn,
      form,
      formRef,
      activeName,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/assets/style/index.scss';

.from {
  width: 100%;
  // height: 100%;
}
.formRef {
  margin: 0 10px;
}
.el-form-item {
  margin-right: 40px;
  margin-top: 20px;
}
//.fromSlot {
//  position: absolute;
//  top: 60px;
//  right: 20px;
//}
:deep(.el-collapse-item__arrow) {
  position: absolute;
  bottom: 0;
  left: 50%;
}
:deep(.el-collapse-item__header) {
  height: 0 !important;
}

:deep(.el-button) {
  //border: 1px solid #595aec;
  width: 70px;
  height: 30px;
  background: #e4e4fc;
  border-radius: 15px;
  font-size: 14px;
  // width: 70px;
  font-weight: 500;
  span {
    color: #595aec !important;
  }
}
#SelectAndInput {
  // width: 270px;
  display: flex;
  :deep(.el-select .el-input .el-input__inner) {
    width: 90px;
    font-size: 12px;
    color: #302e6d;
  }
  .myInput {
    width: 180px;
  }
}
.CustomInputAndIcon {
  :deep(.el-input__suffix-inner) {
    padding-top: 7px;
  }
}
.CompanyInput {
  :deep(.el-input-group__append) {
    min-width: 40px;
    text-align: center;
    background: none;
    line-height: 0 !important;
    border-top: 1px solid #b8b8d5;
    border-right: 1px solid #b8b8d5;
    border-bottom: 1px solid #b8b8d5;
    border-radius: 4px;
  }
}
:deep(#SelectAndInput .el-form-item__content) {
  flex-wrap: nowrap !important;
}
</style>
