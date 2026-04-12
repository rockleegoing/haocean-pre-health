# 移动端后台管理系统 RuoYi-Mobile

<br><br>

###  :watermelon:  技术栈
| 组件       | 版本     |
|----------|--------|
| uview-ui | 2.0.31 |
| qiun-data-charts | 2.4.3-20220505 |

<br><br>

###   :tw-1f348:   系统配置
1. 后端请求地址配置：

```
# 配置文件路径：/config/environment.js
# 修改baseURL属性

const environment = {
	// 开发环境配置
	development: {
		// 本地部署的后端
		baseURL: 'http://localhost:8080',
		
		// 直接使用线上后端
		// baseURL: 'http://vue.ruoyi.vip/prod-api'
	},
	// 生产环境配置
	production: {
		baseURL: 'http://vue.ruoyi.vip/prod-api' // 发布时需要修改为后端实际地址
	}
}

module.exports = {
  environment: environment[process.env.NODE_ENV]
}
```
2. H5启动端口配置：

***注意：不要在manifest.json中配置h5启动信息，可能会引发后端接口访问异常***

```
# 配置文件路径：/vue.config.js
# 修改port属性

const { environment } = require('./config/environment.js')

module.exports = {
  devServer: {
    port: 9001,
    proxy: {
      '/': {
        target: environment.baseURL,
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/': ''
        }
      }
    },
  }
}

```


###    :lemon:   系统截图
<table>
    <tr>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/login.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/index.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/work.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/my.jpg"/></td>
    </tr>
    <tr>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/notice-m.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/notice-e.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/prefile.jpg"/></td>
        <td valign="top"><img src="https://gitee.com/yinm/RuoYi-Mobile/raw/master/static/preview/pwd-reset.jpg"/></td>
    </tr>
</table>