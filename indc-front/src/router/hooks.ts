// import store from '../store'
// import Cookies from 'js-cookie'
// import * as types from '@/store/action-types'
// const loginPermission = async function (to, from, next) {
//   const flag = Cookies.get('token')
//   const needLogin = to.matched.some((item) => item.meta.needLogin)
//   if (!(store.state as any).user.hasPermission) {
//     if (needLogin) {
//       if (flag) {
//         next()
//       } else {
//         next('/login')
//       }
//     } else {
//       next()
//     }
//   } else {
//     if (to.path === '/login') {
//       next('/')
//     } else {
//       next()
//     }
//   }
// }

const menuPermission = async function (to, from, next) {
  // await store.dispatch(`user/${types.ADD_ROUTES}`)
  // if ((store.state as any).user.hasPermission ) {
  //   if (!(store.state as any).menuPermission) {
  //     await store.dispatch(`user/${types.ADD_ROUTES}`)
  //     return next({ ...to, replace: true })
  //   } else {
  //     next()
  //   }
  // } else {
  //   next()
  // }
  next()
}
export default {
  // loginPermission,
  menuPermission,
}
