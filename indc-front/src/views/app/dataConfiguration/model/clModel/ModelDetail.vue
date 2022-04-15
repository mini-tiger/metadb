<template>
  <Breadcrumb />
  <div class="absolute-button" v-if="type !== 'look'">
    <publicButton :data="buttonData.buttonIn" @onButton="upLoad"></publicButton>
    <publicButton
      :data="buttonData.buttonOut"
      @onButton="downLoad"
    ></publicButton>
    <!--    <publicButton :data="buttonData.buttonSubmit"></publicButton>-->
    <!--    <publicButton :data="buttonData.buttonCancel"></publicButton>-->
  </div>
  <detailHeader />
  <detailContent />
  <UpLoadDrawer
    :isShowUpLoad="isShowUpLoad"
    @btnCancel="btnCancel"
    @btnSure="btnSure"
  />
</template>

<script lang="ts">
import { defineComponent, ref, toRefs } from 'vue'
import publicButton from '@/components/Button/index.vue'
import detailHeader from '@/views/app/dataConfiguration/model/clModel/cpns/detailHeader.vue'
import detailContent from '@/views/app/dataConfiguration/model/clModel/cpns/detailContent.vue'
import { buttonData, uploadData } from './config/model-detail'
import UpLoadDrawer from '@/components/UpLoadDrawer/index.vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'ModelDetail',
  components: {
    publicButton,
    detailHeader,
    detailContent,
    UpLoadDrawer,
  },

  setup() {
    const route = useRoute()

    if (route.params && !localStorage.getItem('rowData')) {
      const data = JSON.parse(JSON.stringify(route.params.data))
      localStorage.setItem('rowData', data)
    }
    const type = JSON.parse(localStorage.getItem('rowData')).type || ''

    return {
      buttonData,
      ...toRefs(uploadData),
      type,
    }
  },
})
</script>

<style scoped></style>
