<template>
  <div class="div-box">
    <slot />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
export default defineComponent({
  name: "divBox",
})
</script>

<style scoped>
.div-box {
  background: #F5F6FA !important;
  box-sizing: border-box;
  padding: 0 10px 30px 20px;
  height: auto;
}
</style>
