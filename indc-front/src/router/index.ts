import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import hooks from './hooks'
const routes: Array<RouteRecordRaw> = []
const files = require.context('./', false, /\.router\.ts$/)
files.keys().forEach((key) => {
  routes.push(...files(key).default)
})
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})
Object.values(hooks).forEach((hook) => {
  router.beforeEach(hook.bind(router))
})
export default router
