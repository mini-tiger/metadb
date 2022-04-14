const baseURL = '/topo'
export default {
  getModelClassifyList: `${baseURL}/modelClassification/treeList`, //列表
  saveModelClassify: `${baseURL}/modelClassification/save`, //新增
  updateModelClassify: `${baseURL}/modelClassification/update`, //编辑
  delModelClassify: `${baseURL}/modelClassification/del`, //删除
  uploadModelClassify: `${baseURL}/modelClassification/importObj`, //导入
  downloadModelClassify: `${baseURL}/modelClassification/exportObj`, //导出
}
