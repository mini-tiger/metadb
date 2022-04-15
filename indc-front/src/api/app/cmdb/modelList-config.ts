const baseURL = '/topo'
export default {
  getModelList:`${baseURL}/obj/getObjList`,//模型列表
  saveModel:`${baseURL}/obj/saveObj`,//新增模型
  updateModel:`${baseURL}/obj/updateObj`,//修改模型
  delModel:`${baseURL}/obj/delObj`,//删除模型
  isPausedModel:`${baseURL}/obj/enableORdisableObj`,//启用停用模型
}