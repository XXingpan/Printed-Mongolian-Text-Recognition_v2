
module.exports = {

  devServer: {
    proxy: {
      '/api': {
        target: ' http://localhost:8089', // 这里修改为你的后端服务地址
        changeOrigin: true,
        // pathRewrite: {
        //   '^/api': '' // 如果后端接口路径中不包含/api，可以将/api替换为空字符串
        // }
      }
    }
  }
}

