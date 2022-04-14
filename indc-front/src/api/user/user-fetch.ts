import axios from '@/utils/httpRequest'
import config from './user-config'
export const validate = () => axios.get(config.validate)
export const login = (options) => axios.post(config.login, options)
export const logout = () => axios.get(config.logout)
