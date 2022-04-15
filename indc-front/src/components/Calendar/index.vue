<template>
  <div class="calendar">
    <el-config-provider :locale="locale">
      <el-calendar ref="calendar">
        <template #header="{ date }">
          <span>我的工作日历</span>
          <span>{{ date }}</span>
          <el-button-group>
            <el-button
              style="color: #6162ec; border: 1px solid #6162ec"
              size="small"
              round
              class="exportsData"
              @click="dialogVisible = true"
              >新建任务</el-button
            >
          </el-button-group>
        </template>
        <template #dateCell="{ data }">
          <p :class="data.isSelected ? 'is-selected' : ''">
            <!-- {{ data }} -->
            <!-- {{ data.day.split('-').slice(1).join('-') }}
            {{ data.isSelected ? '✔️' : '' }} -->
          </p>
        </template>
      </el-calendar>
    </el-config-provider>

    <el-dialog
      v-model="dialogVisible"
      title="Tips"
      width="30%"
      :before-close="handleClose"
    >
      <el-form
        :label-position="labelPosition"
        label-width="100px"
        :model="formLabelAlign"
        style="max-width: 460px"
      >
        <el-form-item label="日期">
          <el-date-picker
            v-model="formLabelAlign.dateTimer"
            type="date"
            placeholder="请输入时间"
          >
          </el-date-picker>
        </el-form-item>
        <el-form-item label="时间">
          <el-time-select
            v-model="formLabelAlign.startTime"
            :max-time="formLabelAlign.endTime"
            class="mr-4"
            placeholder="Start time"
            start="9:00"
            step="00:30"
            end="18:00"
          >
          </el-time-select>
          <el-time-select
            v-model="formLabelAlign.endTime"
            :min-time="formLabelAlign.startTime"
            placeholder="End time"
            start="09:00"
            step="00:30"
            end="18:00"
          >
          </el-time-select>
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="formLabelAlign.content"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="dialogOk">确认</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, reactive } from 'vue'
import { ElConfigProvider } from 'element-plus'
import zhCn from 'element-plus/lib/locale/lang/zh-cn'
import moment from 'moment'

import { getDate } from '@/utils/commUtils'
export default defineComponent({
  components: {
    ElConfigProvider,
  },
  setup() {
    let dialogVisible = ref(false)
    const labelPosition = ref('right')
    const formLabelAlign = reactive({
      startTime: '',
      endTime: '',
      content: '',
      dateTimer: '',
    })
    const handleClose = () => {
      console.log('关闭')
    }
    const dialogOk = () => {
      // dialogVisible.value = false
      const createData = {
        startTime: formLabelAlign.startTime,
        endTime: formLabelAlign.endTime,
        content: formLabelAlign.content,
        dateTimer: moment(formLabelAlign.dateTimer).format('MM-DD'),
      }
      console.log(createData, '------>')
      const calendar = document.querySelectorAll<HTMLElement>(
        '.el-calendar-table thead th',
      )
      const logOKDate = createData.dateTimer.split('-')
      console.log(calendar, 'calendar')
      calendar.forEach((item, index) => {
        // console.log(item, 'item')
        const tableData = item.innerText
        // console.log(logOKDate, '定义的')
        // console.log(tableData, 'table的')

        if (tableData.split('/')[1]) {
          console.log(tableData.split(' ')[0].split('/')[0] == logOKDate[0])
          console.log(tableData.split(' ')[0].split('/')[1] == logOKDate[1])
          if (
            tableData.split(' ')[0].split('/')[0] == logOKDate[0] &&
            tableData.split(' ')[0].split('/')[1] == logOKDate[1]
          ) {
            console.log(item, index)
          }
        }
      })
    }
    onMounted(() => {
      const calendarTitle = document.querySelectorAll('th')
      const isWeek = '日一二三四五六'.charAt(new Date().getDay())
      calendarTitle.forEach((item) => {
        if (item.innerText == isWeek) {
          item.innerText = `${getDate()}   周${isWeek}`
        }
      })
    })

    return {
      locale: zhCn,
      dialogVisible,
      handleClose,
      dialogOk,
      labelPosition,
      formLabelAlign,
    }
  },
})
</script>

<style scoped>
.calendar {
  flex: 1;
  height: 50%;
}
:v-deep(.el-form-item__label) {
  line-height: 17px;
  margin-left: 20px;
}
</style>
