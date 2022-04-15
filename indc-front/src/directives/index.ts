// 自定义指令
const files = require.context('./directive', false, /\.ts$/)
export default {
  install(app) {
    files.keys().forEach((key) => {
      const directive = files(key).default
      key = key.replace(/^\.\//, '').replace(/\.ts$/, '')
      app.directive(key, directive)
    })
  },
}
