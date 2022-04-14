import axios from '@/utils/httpRequest'
import config from './modelRelation-config'

// 获取模型全局拓扑
export const getAllObjTop = () => axios.get(config.getAllObjTop)
