import axios from '@/utils/httpRequest'
import config from './modelClassify-config'
export const getModelClassifyList = (options) =>
  axios.post(config.getModelClassifyList, options)
export const saveModelClassify = (options) =>
  axios.post(config.saveModelClassify, options)
export const updateModelClassify = (options) =>
  axios.post(config.updateModelClassify, options)
export const delModelClassify = (options) =>
  axios.get(`${config.delModelClassify}?id=${options}`)
export const uploadModelClassify = (options) =>
  axios.post(config.uploadModelClassify, options)
export const downloadModelClassify = (options) =>
  axios.get(config.downloadModelClassify, options)
