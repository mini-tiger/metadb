<template>
  <div class="works_pace">
    <!-- <el-button type="primary">primary</el-button> -->
    <div class="works_pace_top"></div>
    <!-- <el-pagination
      :page-size="100"
      small="small"
      disabled="disabled"
      layout="total, prev, pager, next"
      :total="1000"
    >
    </el-pagination> -->
    <calendar></calendar>
    <charts :data="chartData"></charts>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive } from 'vue'
import { useStore } from 'vuex'
import * as types from '@/store/action-types'

import calendar from '@/components/Calendar/index.vue'
import charts from '@/components/Chart/index.vue'

export default defineComponent({
  name: 'WorkShop',
  components: {
    calendar,
    charts,
  },
  setup() {
    const chartData = reactive({
      option: [
        {
          title: {
            text: '工单类型统计',
            left: 'left',
            id: 1,
          },
          xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
          },
          yAxis: {
            type: 'value',
          },
          series: [
            {
              title: '工单状态统计',
              data: [120, 200, 150, 80, 70, 110, 130],
              type: 'bar',
              showBackground: true,
              backgroundStyle: {
                color: 'rgba(180, 180, 180, 0.2)',
              },
            },
          ],
        },
        {
          title: {
            text: '工单状态统计',
            left: 'left',
            id: 2,
          },
          tooltip: {
            trigger: 'item',
          },
          series: [
            {
              name: 'Access From',
              type: 'pie',
              radius: '50%',
              data: [
                { value: 1048, name: 'Search Engine' },
                { value: 735, name: 'Direct' },
                { value: 580, name: 'Email' },
                { value: 484, name: 'Union Ads' },
                { value: 300, name: 'Video Ads' },
              ],
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)',
                },
              },
            },
          ],
        },
      ],
    })

    const store = useStore()
    onMounted(async () => {
      try {
        await store.dispatch(`home/${types.USER_SELF}`)
      } catch (error) {
        console.log(error)
      }
    })
    return {
      chartData,
    }
  },
})
</script>
<style lang="scss" scoped>
.works_pace {
  background: #f5f6fa;
  display: flex;
  flex-wrap: wrap;
  .works_pace_top {
    width: 100%;
    height: 60px;
  }
}

::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
