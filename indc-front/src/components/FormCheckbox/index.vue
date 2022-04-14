<template>
  <div class="form-checkbox">
    <div
      class="form-checkbox-item"
      v-for="item in options"
      :key="item.value"
      @click.prevent="checkItem(item)"
    >
      <div
        :class="{
          'form-checkbox-item-content': true,
          'form-checkbox-item-checked': modelValue.indexOf(item.value) > -1,
        }"
      >
        <svg-icon
          :icon-class="
            modelValue.indexOf(item.value) > -1 ? 'checked' : 'check'
          "
          class-name="check-icon"
        />
        {{ item.label }}
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { defineProps, defineEmits, PropType, onMounted } from 'vue'

interface checkboxItem {
  label: string
  value: string
}
const props = defineProps({
  options: {
    type: Array as PropType<checkboxItem[]>,
    default: () => [],
  },
  modelValue: {
    type: Array,
    default: () => [],
  },
})

let emit = defineEmits(['update:modelValue', 'update'])
const checkItem = (item) => {
  const itemIndex = props.modelValue.indexOf(item.value)
  let modelValueCopy = JSON.parse(JSON.stringify(props.modelValue))
  if (itemIndex > -1) {
    modelValueCopy.splice(itemIndex, 1)
    emit('update:modelValue', modelValueCopy)
    // emit('update', modelValueCopy)
    console.log(props.modelValue)
  } else {
    modelValueCopy.push(item.value)
    emit('update:modelValue', modelValueCopy)
    // emit('update', modelValueCopy)
    console.log(props.modelValue, modelValueCopy)
  }
}
</script>
<style lang="scss" scoped>
@import '@/assets/style/var.scss';
.form-checkbox {
  height: 30px;
  display: flex;
  .form-checkbox-item {
    height: 100%;
    cursor: pointer;
    margin: 0 5px;
    .form-checkbox-item-content {
      height: 100%;
      width: 50px;
      border-radius: 4px;
      border: 1px solid #b8b8d5;
      text-align: center;
      line-height: 30px;
      color: #b8b8d5;
      position: relative;
      .svg-icon {
        position: absolute;
        right: 0;
        bottom: 0;
      }
    }
    .form-checkbox-item-checked {}
  }
}
</style>
