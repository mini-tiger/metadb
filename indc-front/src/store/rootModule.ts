import getters from './getters'
import * as types from './action-types'

export default {
  state: {
    pointSets: {
      edges: [
        {
          id: '26e1efe2-8fb1-4a6a-b5dc-33d78377be68',
          type: 'htmlLineEdge',
          sourceNodeId: 'cb7f2a40-f80f-4bf0-ad73-d937d8b13fa1',
          targetNodeId: '2f4145dd-faf0-4c6f-a5b0-2648fc048989',
          startPoint: { x: 140, y: 98 },
          endPoint: { x: 60, y: 222 },
          properties: {},
          pointsList: [
            { x: 140, y: 98 },
            { x: 140, y: 192 },
            { x: 60, y: 192 },
            { x: 60, y: 222 },
          ],
          text: { value: '父子' },
        },
        {
          id: '7d9ba0e2-9a7d-4195-9b42-d41f1cd2d2c6',
          type: 'htmlLineEdge',
          sourceNodeId: 'cb7f2a40-f80f-4bf0-ad73-d937d8b13fa1',
          targetNodeId: '13c465e4-1ad7-4cd4-9741-37b61003b451',
          startPoint: { x: 140, y: 98 },
          endPoint: { x: 220, y: 222 },
          properties: {},
          pointsList: [
            { x: 140, y: 98 },
            { x: 140, y: 192 },
            { x: 220, y: 192 },
            { x: 220, y: 222 },
          ],
          text: { value: '父子' },
        },
      ],
      nodes: [
        {
          id: 'cb7f2a40-f80f-4bf0-ad73-d937d8b13fa1',
          type: 'htmlNode',
          x: 146,
          y: 75,
          properties: {
            id: 'node08bf87e8-840c-485e-bba3-986e28eadc0f',
            text: 'Level one 1节点',
          },
        },
        {
          id: '2f4145dd-faf0-4c6f-a5b0-2648fc048989',
          type: 'htmlNode',
          x: 64,
          y: 244,
          properties: {
            id: 'nodef764e687-8678-427b-9575-18714a58cd47',
            text: 'Level one 1节点',
          },
        },
        {
          id: '13c465e4-1ad7-4cd4-9741-37b61003b451',
          type: 'htmlNode',
          x: 221,
          y: 246,
          properties: {
            id: 'node06e32d19-75c8-4478-8a0c-50b182344b49',
            text: 'Level one 1节点',
          },
        },
      ],
    },
  },
  getters,
  mutations: {
    [types.CHANGE_POSITION](state, payload) {
      state.pointSets = payload
    },
    [types.ADD_POINT](state, payload) {
      state.pointSets.nodes = [...state.pointSets.nodes, payload]
    },
    [types.ADD_EDGE](state, edge) {
      state.pointSets.edges = [...state.pointSets.edges, edge]
    },
    // [types.GET_RELATIVEEDGE](state, id) {
    //   const edges = state.pointSets.edges
    //   edges.forEach((edge) => {
    //     if (edge.sourceNodeId === id || edge.targetNodeId === id) {
    //       state.filterEdges.push(edge)
    //     }
    //   })
    // },
    // [types.CLEAR_FILTERS_EDGE](state, result = []) {
    //   state.filterEdges = result
    // },
    [types.CLEAR_EDGE](state) {
      state.pointSets.edges = []
    },
    [types.DELETE_EDGE](state, edgeId) {
      //  进行删除边的操作
      const edges = state.pointSets.edges
      const index = edges.findIndex((edge) => edge.id === edgeId)
      edges.splice(index, 1)
    },
    [types.DELETE_POINT](state, nodeId) {
      const nodes = state.pointSets.nodes
      const index = nodes.findIndex((node) => node.id === nodeId)
      nodes.splice(index, 1)
    },
  },
  actions: {
    [types.CHANGE_POSITION]({ commit }, { allGraphData, urlType }) {
      //根据urlType来判断请求后台的接口，成功后，存储数据
      console.log('urlType', urlType)

      commit(types.CHANGE_POSITION, allGraphData)
    },
    [types.DELETE_EDGE]({ commit }, edgeId) {
      commit(types.DELETE_EDGE, edgeId)
    },
    [types.DELETE_POINT]({ commit }, nodeId) {
      commit(types.DELETE_POINT, nodeId)
    },
    [types.CLEAR_EDGE]({ commit }) {
      commit(types.CLEAR_EDGE)
    },
  },

  modules: {},
}
