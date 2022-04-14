import Mock from 'mockjs'
import {  userLogin,userValidate,userlogout} from './data'
export default {
  init() {
    Mock.mock(/\/api\/user\/login/, 'post', userLogin),
    Mock.mock(/\/api\/user\/validate/, 'get', userValidate),
    Mock.mock(/\/api\/user\/logout/, 'get', userlogout)
  },
}
