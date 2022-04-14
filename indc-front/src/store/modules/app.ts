import {
  APP_CENTER_MENU,
  USER_SELF,
  APP_ROOT_MENU,
  LANGUAGE,
  SET_LANGUAGE,
} from '../action-types'
import * as app from '@/api/app/app-fetch'
import { getItem, setItem } from '@/utils/commUtils'
let filterArr = []
function filterRootNode(data) {
  //需要清空数据
  filterArr = []
  return data.map((item) => {
    item.subMenus.length > 0 && filterArr.push([...item.subMenus])
    return {
      id: item.id,
      name: item.name,
      icon: item.icon,
    }
  })
}
export default {
  state: {
    menu: [],
    rootMenu: [],
    language: getItem(LANGUAGE) || 'zh',
  },
  getters: {
    getMenuById: (state) => (id) => {
      return state.menu.filter((item) => item.parentId === id)
    },
  },
  mutations: {
    [APP_CENTER_MENU](state, payload) {
      state.menu = payload
    },
    [APP_ROOT_MENU](state, payload) {
      state.rootMenu = payload
    },
    [SET_LANGUAGE](state, payload) {
      state.language = payload
      setItem(LANGUAGE, payload)
    },
  },
  actions: {
    async [APP_CENTER_MENU]({ commit, rootState, dispatch }) {
      let currentRoleId = rootState.home.userSelf.currentRoleId
      try {
        if (!currentRoleId) {
          await dispatch(`home/${USER_SELF}`, {}, { root: true })
          currentRoleId = rootState.home.userSelf.currentRoleId
        }
        const res = await app.getMenu({
          lang: 'zh_CN',
          roldId: currentRoleId,
          unionLabel: false,
        })
        //需要过滤出菜单的一级数据
        commit(APP_ROOT_MENU, filterRootNode(res))
        const subMenu = filterArr.flat(1)
        commit(APP_CENTER_MENU, subMenu)
        // setLocal('rootMenu', filterRootNode(res))
        // setLocal('menuTree', sendData)
      } catch (error) {
        return Promise.reject(error)
      }
    },
  },
}
