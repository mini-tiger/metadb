export default [
  {
    path: '/work',
    name: '主页',
    component: () => import(/* webpackChunkName:'login' */ '@/views/test.vue'),
  },
]
