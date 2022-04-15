import * as user from '@/api/user/user-fetch'
import * as types from '@/store/action-types'
import { setItem, getItem, removeAllItem } from '@/utils/commUtils'
//所有的静态路由
import router from '../../router'
//需要动态过滤的路由
import dynamicRouter from '@/router/dynamicRouter'
const filterRoute = (authList) => {
  //后台返回的路由菜单，需要格式成一层的数组
  // const auths = authList.map((item) => item)
  function filterRoutes(routes) {
    return routes.filter((route) => {
      if (authList.includes(route.meta.id)) {
        if (route.children) {
          route.children = filterRoutes(route.children)
        }
        return route
      }
    })
  }
  return filterRoutes(dynamicRouter)
}
export default {
  state: {
    userInfo: {},
    hasPermission: false,
    menuPermission: false,
    navList: [
      {
        path: '/',
        name: '工作台',
      },
      {
        path: '/search-center',
        name: '检索中心',
      },
      {
        path: '/report-center',
        name: '报表报告',
      },
      {
        path: '/app-center',
        name: '应用中心',
      },
    ],
  },
  mutations: {
    [types.USER_INFORMATION](state, payload) {
      state.userInfo = payload
      if (payload?.token) {
        setItem('token', payload.token)
      } else {
        removeAllItem()
      }
    },
    [types.USER_PERMISSION](state, payload) {
      state.hasPermission = payload
    },
    [types.MENU_PERMISSION](state, payload) {
      state.menuPermission = payload
    },
  },
  actions: {
    async [types.ADD_ROUTES]({ commit, rootState }) {
      const resArray = []
      //过滤为一个层级的数组
      function filterOne(menus) {
        function deepArray(menus) {
          menus.forEach((menu) => {
            if (menu?.subMenus?.length > 0) {
              deepArray(menu.subMenus)
            }
            resArray.push(menu?.id)
          })
        }
        deepArray(menus)
        return resArray
      }
      const menu = filterOne(rootState.app.menu)
      const authList = [
        ...rootState.app.rootMenu.map((menu) => menu.id),
        ...menu,
      ]
      if (authList) {
        const routes = filterRoute(authList)
        routes.forEach((route) => {
          router.addRoute('appMenu', route)
        })
        commit(types.MENU_PERMISSION, true)
      }
    },
    async [types.USER_LOGIN]({ commit }, payload) {
      commit(types.USER_PERMISSION, true)
      return true
      // try {
      //   const res = await user.login(payload)
      //   if (res.code == 200) {
      //     // commit(types.USER_INFORMATION, res.data)
      //     commit(types.USER_PERMISSION, true)
      //     return true
      //   }
      // } catch (error) {
      //   return Promise.reject(error)
      // }
    },
    async [types.USER_VALIDATE]({ commit }) {
      if (!getItem('token')) {
        return false
      }
      try {
        const res = await user.validate()
        if (res.code == 200) {
          commit(types.USER_INFORMATION, res.data)
          commit(types.USER_PERMISSION, true)
        }
        return true
      } catch (error) {
        commit(types.USER_INFORMATION, {})
        commit(types.USER_PERMISSION, false)
        return false
      }
    },
    async [types.USER_LOGOUT]({ commit }) {
      try {
        const res = await user.logout()
        if (res.code == 200) {
          commit(types.USER_INFORMATION, {})
          commit(types.USER_PERMISSION, false)
        }
      } catch (error) {
        return Promise.reject(error)
      }
    },
  },
}
