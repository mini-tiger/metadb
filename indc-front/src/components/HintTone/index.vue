<template>
  <el-button type="primary" @click="onHint">点击提示音</el-button>
</template>

<script lang="ts" setup>
import Speech from 'speak-tts'
let speech = new Speech()

const onHint = () => {
  speech
    .speak({
      text: 'Hello, how are you today ?',
      queue: false, // current speech will be interrupted,
      listeners: {
        onstart: () => {
          console.log('Start utterance')
        },
        onend: () => {
          console.log('End utterance')
        },
        onresume: () => {
          console.log('Resume utterance')
        },
        onboundary: (event) => {
          console.log(
            event.name +
              ' boundary reached after ' +
              event.elapsedTime +
              ' milliseconds.',
          )
        },
      },
    })
    .then(() => {
      console.log('Success !')
    })
    .catch((e) => {
      console.error('An error occurred :', e)
    })
}
</script>
<style lang="scss" scoped></style>
