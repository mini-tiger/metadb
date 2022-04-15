import * as types from '../action-types'
import * as modelList from '@/api/app/cmdb/modelList-fetch'
import { ElMessage } from 'element-plus'
export default {
  state: {
    data: {
      totalElements: 0,
      content: [],
    },
  },
  mutations: {
    [types.GET_MODEL_LIST](state, payload) {
      state.data = payload
    },
  },
  actions: {
    async [types.GET_MODEL_LIST]({ commit }, options) {
      try {
        const res = await modelList.getModelList(options)
        console.log('res1111111111', res)
        commit(types.GET_MODEL_LIST, res.data)
      } catch (error) {
        return Promise.reject(error)
      }
    },
    async [types.SAVE_MODEL]({ commit, rootState, dispatch }, options) {
      try {
        const res = await modelList.saveModel(options)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelList/${types.GET_MODEL_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_LIST, rootState.modelList.data)
        } else {
          ElMessage.error(res.msg)
        }
      } catch (e) {
        return Promise.reject(e)
      }
    },
    async [types.DEL_MODEL]({ commit, rootState, dispatch }, deleteID) {
      try {
        const res = await modelList.delModel(deleteID)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelList/${types.GET_MODEL_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_LIST, rootState.modelList.data)
        } else {
          ElMessage.error(res.msg)
        }
      } catch (e) {
        return Promise.reject(e)
      }
    },
    async [types.UPDATE_MODEL]({ commit, rootState, dispatch }, options) {
      try {
        const res = await modelList.updateModel(options)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelList/${types.GET_MODEL_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_LIST, rootState.modelList.data)
        } else {
          ElMessage.error(res.msg)
        }
      } catch (e) {
        return Promise.reject(e)
      }
    },
  },
  modules: {},
}
