import { check } from '@/utils/commUtils'

export default {
  inserted(el, binding) {
    if (!check(binding.value)) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  },
}
