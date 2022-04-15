import axios from '@/utils/httpRequest'
import config from './app-config'
export const getMenu = (options) => axios.get(config.getMenu, options)
