import { createApp } from 'vue'
import i18n from '@/i18n'
import App from './App.vue'
import router from './router'
import store from './store'
import { echarts } from './utils/chartConfiguration'
import './assets/style/index.scss'
import './assets/style/commStyle.scss'
import 'element-plus/theme-chalk/src/message.scss'
import './assets/style/font.scss'
// import Mock from '../mock'
// Mock.init()
import './assets/icons'
//全局组件注册
import components from '@/components/index'
//全局指令的注册
import directives from './directives'
const app = createApp(App)
app.config.globalProperties.$echarts = echarts // echarts
app
  .use(i18n)
  .use(components)
  .use(directives)
  .use(store)
  .use(router)
  .mount('#app')
