import * as types from '../action-types'
import * as operate from '@/api/app/cmdb/operate-fetch'
export default {
  state: {
    auditData: {
      totalElements: '',
      content: [],
    },
  },
  mutations: {
    [types.GET_AUDIT_LIST](state, payload) {
      state.auditData = payload
    },
  },
  actions: {
    async [types.GET_AUDIT_LIST]({ commit }, options:any) {
      console.log('options', options)
      try {
        const res = await operate.getAuditList(options)
        console.log('auditRes', res)

        commit(types.GET_AUDIT_LIST, res.data)
      } catch (error) {
        console.log('error', error)
        return Promise.reject(error)
      }
    },
  },
}
