const getters = {
  language: (state) => state.app.language,
  themeColor: (state) => state.theme.themeColor,
  getCurrentNode: (state) => (id: string) => {
    return state.pointSets.nodes.find((node: any) => node.id === id)
  },
  edgeIds: (state) => state.pointSets.edges.map((edge) => edge.id),
}
export default getters
