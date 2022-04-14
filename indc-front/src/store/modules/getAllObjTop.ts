import { GET_MODEL_RELATION } from '../action-types'
import * as modelRelation from '@/api/app/cmdb/model-relation.fetch'
export default {
  state: {
    getAllObjTop: {
      data: {},
    },
  },
  mutations: {
    [GET_MODEL_RELATION](state, payload) {
      console.log(state, payload)
      state.getAllObjTop = payload
    },
  },
  actions: {
    async [GET_MODEL_RELATION]({ commit }) {
      // console.log(options, 'options')
      // const res = await modelRelation.getAllObjTop(options)
      // console.log(res, 'res---------->')
      try {
        const res = await modelRelation.getAllObjTop()
        console.log(res, 'res-=--->')
        commit(GET_MODEL_RELATION, res)
      } catch (e) {
        return Promise.reject(e)
      }
    },
  },
  modules: {},
}
