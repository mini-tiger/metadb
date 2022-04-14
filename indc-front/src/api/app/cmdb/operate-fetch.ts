import axios from '@/utils/httpRequest'
import config from './operate-config'
export const getAuditList = (options) => axios.post(config.getAuditList, options)
