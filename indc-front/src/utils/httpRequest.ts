import axios from 'axios'
import { ElMessage } from 'element-plus'
// import Cookies from 'js-cookie'
class HttpRequest {
  timeout: number
  baseURL: string
  constructor() {
    this.baseURL = ''
    this.timeout = 3000
  }
  setInterceptor(instance) {
    instance.interceptors.request.use(
      (config) => {
        config.headers.Authorization =
          'Bearer ' + '27195e92-9b60-4dc3-acf5-51dd78c16034'
        return config
      },
      (error) => {
        return Promise.reject(error)
      },
    )
    instance.interceptors.response.use(
      (res) => {
        return res.data
      },
      (error) => {
        if (!error.response) return
        const { status, data } = error.response

        switch (status) {
          case 401:
            ElMessage.error('token过期,请登录重新获取')
            // store.dispatch('user/logout')
            break
          default:
            break
        }
        // ElMessage.error(requestMessage) // 提示错误信息
        return Promise.reject(error)
      },
    )
  }
  request(options) {
    const instance = (axios as any).create()
    this.setInterceptor(instance)
    const opts = this.mergeOptions(options)
    return instance(opts)
  }
  mergeOptions(options) {
    return {
      baseURL: this.baseURL,
      timeout: this.timeout,
      ...options,
    }
  }
  get(url: string, config = {}) {
    return this.request({
      method: 'get',
      url,
      params: config,
    })
  }
  post(url: string, data) {
    return this.request({
      method: 'post',
      url,
      data,
    })
  }
  put(url: string, data) {
    return this.request({
      method: 'put',
      url,
      data,
    })
  }
  delete(url:string,data){
    return this.request({
      method: 'delete',
      url,
      data,
    })
  }
}
export default new HttpRequest()
