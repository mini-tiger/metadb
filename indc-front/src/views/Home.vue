<template>
  <div class="home">
    <!-- <public-chart :data="stateOption"></public-chart> -->
    <public-calendar></public-calendar>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive } from 'vue'
import axios from '@/utils/httpRequest'
// import publicChart from '@/components/publicChart.vue'
import publicCalendar from '@/components/Calendar/index.vue'
export default defineComponent({
  name: 'Home',
  components: {
    // publicChart,
    publicCalendar,
  },
  setup() {
    onMounted(() => {
      axios.get('/getApi', { name: 'tyg' }).then((res) => {
        // console.log(res)
      })
    })

    return {}
  },
})
</script>

<style scoped>
.home {
  width: 100%;
  height: 100%;
  /* background: gray; */
}
</style>
