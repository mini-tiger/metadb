import * as types from '../action-types'
import * as modelClassify from '@/api/app/cmdb/modelClassify-fetch'
import { ElMessage } from 'element-plus'
import { messageConfig } from 'element-plus'
export default {
  state: {
    data: [],
  },
  mutations: {
    [types.GET_MODEL_CLASSIFY_LIST](state, payload) {
      state.data = payload
    },
  },
  actions: {
    async [types.GET_MODEL_CLASSIFY_LIST]({ commit }, options) {
      try {
        const res = await modelClassify.getModelClassifyList(options)
        commit(types.GET_MODEL_CLASSIFY_LIST, res.data)
      } catch (error) {
        return Promise.reject(error)
      }
    },
    async [types.SAVE_MODEL_CLASSIFY](
      { commit, rootState, dispatch },
      options,
    ) {
      try {
        const res = await modelClassify.saveModelClassify(options)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelClassify/${types.GET_MODEL_CLASSIFY_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_CLASSIFY_LIST, rootState.modelClassify.data)
        } else {
          ElMessage.error(res.msg)
        }
      } catch (e) {
        return Promise.reject(e)
      }
    },
    async [types.DEL_MODEL_CLASSIFY](
      { commit, rootState, dispatch },
      deleteID,
    ) {
      try {
        const res = await modelClassify.delModelClassify(deleteID)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelClassify/${types.GET_MODEL_CLASSIFY_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_CLASSIFY_LIST, rootState.modelClassify.data)
        } else {
          ElMessage.error(res.msg)
        }
      } catch (e) {
        return Promise.reject(e)
      }
    },
    async [types.UPDATE_MODEL_CLASSIFY](
      { commit, rootState, dispatch },
      options,
    ) {
      try {
        const res = await modelClassify.updateModelClassify(options)
        if (res.code === 200) {
          ElMessage.success(res.msg)
          await dispatch(
            `modelClassify/${types.GET_MODEL_CLASSIFY_LIST}`,
            {},
            { root: true },
          )
          commit(types.GET_MODEL_CLASSIFY_LIST, rootState.modelClassify.data)
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
