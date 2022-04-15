import { createStore } from 'vuex'
import rootModule from './rootModule'

const files = require.context('./modules', false, /\.ts$/)
files.keys().forEach((key) => {
  const store = files(key).default
  const moduleName = key.replace(/^\.\//, '').replace(/\.ts$/, '')
  const modules = (rootModule.modules = rootModule.modules || {})
  modules[moduleName] = store
  modules[moduleName].namespaced = true
  // rootModule.modules = modules
})
export default createStore(rootModule)
