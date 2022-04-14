import axios from '@/utils/httpRequest'
import config from './modelList-config'

export const getModelList = (options) => axios.post(config.getModelList, options)
export const saveModel = (options) => axios.post(config.saveModel, options)
export const updateModel = (options) => axios.post(config.updateModel, options)
export const delModel = (options) => axios.delete(config.delModel, options)
export const isPausedModel = (options) => axios.post(config.isPausedModel, options)