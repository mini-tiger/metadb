const AutoImport = require('unplugin-auto-import/webpack')
const Components = require('unplugin-vue-components/webpack')
const path = require('path')
const { ElementPlusResolver } = require('unplugin-vue-components/resolvers')
module.exports = {
  configureWebpack: {
    plugins: [
      AutoImport({
        resolvers: [ElementPlusResolver()],
      }),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
    ],
  },
  chainWebpack: (config) => {
    config.module.rules.delete('svg') //重点:删除默认配置中处理svg,
    config.module
      .rule('svg-sprite-loader')
      .test(/\.svg$/)
      .include.add(path.resolve(__dirname, 'src/assets/icons')) //处理svg目录
      .end()
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]',
      })
  },
  devServer: {
    open: true,
    historyApiFallback: true,
    proxy: {
      [process.env.VUE_APP_URL]: {
        target: 'http://21vianet.test',
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_URL]: '',
        },
      },
      [process.env.VUE_APP_BASE_URL]: {
        target: 'http://21vianet.test',
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_URL]: '',
        },
      },
     '/topo': {
        // target: 'http://172.22.50.216:60002',
        // target: 'http://172.16.50.97:60002/', // 嘉敏
        target: 'http://172.16.50.71:60002', // 徐丽蓉
        changeOrigin: true,
        // pathRewrite: {
        //   '^/topo' : '/',
        // },
      },
    },
  },
}
