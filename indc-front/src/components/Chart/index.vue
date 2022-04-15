<template>
  <div>
    <div class="chart" v-for="(item, index) in chartList.option" :key="index">
      <div class="chartTop">
        <div class="dateTab" :id="'dateTab' + index" v-if="chartList.isShow">
          <div
            v-for="(itm, ind) in chartList.selectTime"
            :key="ind"
            class="selectTime"
            @click="selectClick($event, ind, itm, item)"
          >
            {{ itm }}
          </div>
        </div>
      </div>
      <div class="wrapper">
        <div
          :id="'publicChar' + index"
          class="publicChar"
          style="height: 300px"
        ></div>
      </div>
      <el-button
        style="color: #6162ec; border: 1px solid #6162ec"
        size="small"
        round
        class="exportsData"
        >导出数据</el-button
      >
    </div>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  getCurrentInstance,
  onMounted,
  reactive,
  nextTick,
} from 'vue'
export default defineComponent({
  props: ['data'],
  setup(props, context) {
    const chartList = reactive({
      titleText: '工单类型',
      selectTime: ['日', '周', '月'],
      isShow: true,
      option: [],
      dateTab: [],
    })
    const internalInstance = getCurrentInstance()
    const echarts = internalInstance.appContext.config.globalProperties.$echarts
    onMounted(() => {
      if (props.data) {
        // chartList.selectTime = props.data.selectTime
        // chartList.isShow = props.data.isShowTabs
        chartList.option = props.data.option
      }
      // console.log(chartList.selectTime, '------->')
    })
    // 给父组件传递日期
    const selectClick = (e, index, val, option) => {
      const selectR = document.querySelectorAll<HTMLElement>(
        `#${e.path[1].id} .${e.path[0].className}`,
      )
      // console.log(selectR, '------>', val, option)
      selectR.forEach((item) => {
        item.style.background = '#fff'
      })
      selectR[index].style.background = '#eeeefd'
      context.emit('selectClick', val, option.title.id)
    }
    nextTick(() => {
      if (chartList.option && chartList.option.length !== 0) {
        chartList.option.map((item, index) => {
          drawLine(`#publicChar${index}`, item)
          // chartList.selectTime.map((itm, ind) => {
          // })
        })
      }
    })
    // 不同的echarts
    const drawLine = (name, item) => {
      // console.log(name, item)
      let myChart = echarts.init(document.querySelector(name))
      myChart.clear()
      myChart.resize()
      myChart.setOption(item)
    }

    return {
      selectClick,
      chartList,
    }
  },
})
</script>

<style scoped>
.chart {
  width: 340px;
  height: 340px;
  margin-left: 20px;
  position: relative;
  box-sizing: border-box;
  padding: 10px;
  background: #fff;
  border-radius: 5px;
  margin-bottom: 20px;
}
.chart:last-child {
  margin-bottom: 0;
}
.chartTop {
  display: flex;
  justify-content: space-between;
}
.title {
  width: 99px;
  height: 20px;
  line-height: 29px;
  color: #302e6d;
  font-size: 16px;
  font-family: PingFangSC-Medium, PingFang SC;
}
.selectTime {
  width: 49px;
  height: 26px;
  text-align: center;
  line-height: 26px;
  cursor: pointer;
  border: 1px solid #f2f2f6;
}
.selectTime:first-child {
  border-top-left-radius: 5px;
  border-bottom-left-radius: 5px;
}
.selectTime:last-child {
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
}
.dateTab {
  display: flex;
  position: absolute;
  right: 0;
  z-index: 1;
}

.dateTab :hover {
  background: #eeeefd;
}

.dateTab .active {
  background: #eeeefd;
}
.exportsData {
  position: absolute;
  right: 9%;
  bottom: 4%;
}
</style>
