import { USER_SELF } from '../action-types'
import * as home from '@/api/home/home-fetch'
export default {
  state: {
    userSelf: '',
    count: 10,
  },
  mutations: {
    [USER_SELF](state, payload) {
      state.userSelf = payload
    },
  },
  actions: {
    async [USER_SELF]({ commit }) {
      try {
        const res = await home.getSelf()
        commit(USER_SELF, res)
      } catch (error) {
        return Promise.reject(error)
      }
    },
  },
}
