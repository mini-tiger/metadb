import { createI18n } from 'vue-i18n'
import store from '@/store'
const messages = {
  en: {
    msg: {
      test: 'hello world',
    },
  },
  zh: {
    msg: {
      test: '你好，世界',
    },
  },
}
function getLanguage() {
  return store?.getters?.language
}
export default createI18n({
  legacy: false,
  globalInjection: true,
  locale: getLanguage(),
  messages,
})
