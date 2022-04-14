<template>
  <div>
    <el-button
      id="button"
      :class="state.button.className"
      :type="state.button.type"
      :round="state.button.round"
      :circle="state.button.circle"
      :loading="state.button.loading"
      :icon="state.button.icon"
      :disabled="state.button.disabled"
      @click="clickButt($event)"
    >
      {{ state.button.name }}</el-button
    >
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive } from 'vue'
interface IButton {
  name: string
  // eslint-disable-next-line @typescript-eslint/ban-types
  className: string
  size?: string
  type: string
  round?: boolean
  circle?: boolean
  loading?: boolean
  loadingIcon?: string
  disabled?: boolean
  icon?: Element
}
export default defineComponent({
  props: ['data'],
  setup(props, content) {
    const button: IButton = {
      name: '',
      className: '',
      size: '',
      type: '',
      round: false,
      circle: false,
      loading: false,
      loadingIcon: '',
      disabled: false,
      icon: null,
    }
    const state = reactive({
      button: button,
    })
    onMounted(() => {
      // console.log(props, '-----><')
      state.button = props.data
    })
    const clickButt = (e) => {
      // console.log(e)
      content.emit('onButton', e)
    }
    return {
      state,
      clickButt,
    }
  },
})
</script>

<style lang="scss" scoped>
.roundBackWhite,
.roundBackBlue,
.roundBackCal,
.roundBlueIcon,
.roundWhiteIcon {
  //width: 90px;
  height: 30px;
  border: 1px solid #595aec;
}
.roundBackBlue {
  border: none !important;
  background: #595aec;
}
.roundBackCal {
  background: #e4e4fc;
  border-radius: 15px;
  border: none !important;
  width: 70px;
}
.roundBlueIcon {
  border: none !important;
  background: #595aec;
}
.roundBackBlue:hover {
  background: #8f8ff6;
}

.roundBackCal:hover {
  background: #8f8ff6;
}

.roundWhiteIcon:hover {
  background: #ededfd;
}

.roundBlueIcon:hover {
  background: #8f8ff6;
}
:deep(.roundWhiteIcon > span, .roundBackCal > span, .roundBackWhite > span) {
  color: #595aec !important;
}

:deep(.roundWhiteIcon > i) {
  width: 12px;
  height: 12px;
  color: #595aec !important;
}

:deep(.roundBackBlue > span) {
  color: #fff !important;
}
:deep(.roundBlueIcon > span) {
  color: #fff !important;
}
:deep(.roundBlueIcon > i) {
  width: 12px;
  height: 12px;
  background: #fff !important;
  color: #595aec;
  border-radius: 50%;
}
:deep(.roundBlueIcon:focus) {
  color: none !important;
  background: #595aec !important;
  border: none !important;
}
:deep(.roundBackCal.el-button:focus) {
  color: none !important;
  background: #e4e4fc !important;
  border: none !important;
}

:deep(.roundBackBlue.el-button:focus) {
  color: none !important;
  background: #595aec !important;
  border: none !important;
}
::v-deep(.el-button.roundBackCal) {
  span {
    color: #595aec !important;
  }
}
</style>
