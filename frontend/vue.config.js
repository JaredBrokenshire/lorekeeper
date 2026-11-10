const path = require('path')

module.exports = {
  publicPath: "/",
  devServer: {
    host: 'jb-lorekeeper.com',
    liveReload: true,
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@core': path.resolve(__dirname, 'src/@core'),
        '@axios': path.resolve(__dirname, 'src/services/api'),
      },
    },
  },
  transpileDependencies: true,
}
