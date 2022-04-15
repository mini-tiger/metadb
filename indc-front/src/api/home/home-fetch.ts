import axios from '@/utils/httpRequest'
import config from './home-config'
export const getSelf = () => axios.get(config.getSelf)
