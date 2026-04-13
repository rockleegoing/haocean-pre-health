const { environment } = require('./config/environment.js')

module.exports = {
  devServer: {
    port: 9001,
    disableHostCheck: true,
    proxy: {
      '/prod-api': {
        target: environment.baseURL,
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/prod-api': ''
        }
      }
    }
  }
}