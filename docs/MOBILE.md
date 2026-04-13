# 移动端开发指南

## 项目说明

本项目有两个移动端解决方案：

| 方案 | 技术栈 | 说明 |
|------|--------|------|
| Android 原生 | Kotlin | 轻量级 Android 应用，支持 Android 终端 |
| UniApp | Vue | 跨平台方案 (H5/小程序/App) |

---

## Android 原生 (Kotlin)

### 技术栈

| 组件 | 说明 |
|------|------|
| ViewBinding | 查找控件 |
| lifecycle-viewmodel | 分离数据与视图 |
| BaseRecyclerViewAdapterHelper | 快速加载列表 |
| ToastUtils | 提示工具 |
| TitleBar | 公共标头 |
| glide | 图片展示 |
| immersionbar | 沉浸式状态栏 |
| XXPermissions | 权限申请 |
| Net + okhttp | 网络请求 |
| therouter | 路由 |

### 环境要求

- **系统**: Windows 10/11 (64 位), macOS 10.14+, Linux (64 位)
- **内存**: 至少 8GB (推荐 16GB)
- **磁盘**: 20GB+ 空闲空间
- **CPU**: 支持虚拟化技术
- **最低 SDK**: 21

### 获取代码

```bash
git clone git@gitee.com:OptimisticDevelopers/Ruoyi-Android-App.git
cd Ruoyi-Android-App
```

### 常见问题

**错误 1**: `Unresolved reference: BuildConfig`
- 解决方法：重新编译项目

**错误 2**: 依赖下载失败
- 解决方法：添加阿里云镜像
```gradle
maven { url 'https://maven.aliyun.com/repository/google' }
maven { url 'https://maven.aliyun.com/repository/central' }
maven { url 'https://maven.aliyun.com/repository/public' }
```

---

## UniApp

### 技术栈

| 组件 | 版本 |
|------|------|
| uview-ui | 2.0.31 |
| qiun-data-charts | 2.4.3-20220505 |
| vue | 2 或 3 |

## 目录结构

```
frontend/RuoYi-Mobile/
├── api/              # API 接口
├── components/       # 自定义组件 (navbar/tabs/qian-tree)
├── config/
│   ├── environment.js  # 环境配置
│   ├── request.js      # 请求封装
│   └── http.interceptor.js  # HTTP 拦截器
├── pages/            # 页面 (index/login/work/center)
├── store/            # Vuex 状态管理
│   └── modules/      # user.js/system.js
├── uni_modules/      # Uni 模块插件 (uview-ui/qiun-data-charts)
└── utils/            # 工具函数
    ├── storage.js    # 缓存工具
    └── verify.js     # 表单验证
```

## 技术栈

| 用途 | 库/插件 |
|------|--------|
| 框架 | UniApp |
| UI | uview-ui 2.0 |
| 图表 | qiun-data-charts 2.4 |
| 状态管理 | Vuex |

## 环境配置

文件：`config/environment.js`

```javascript
const environment = {
  development: {
    baseURL: 'http://localhost:8080'
  },
  production: {
    baseURL: 'http://vue.ruoyi.vip/prod-api'
  }
}

module.exports = {
  environment: environment[process.env.NODE_ENV]
}
```

## 请求封装

文件：`config/request.js`

```javascript
import { http } from '@/uni_modules/uview-ui'

const request = {
  post (url, params, config = {}) {
    config = mixinCustom(config)
    return http.post(url, params, config)
  },
  get (url, params, config = {}) {
    config = mixinCustom(config)
    return http.get(url, config)
  }
}

// 默认配置
config.custom = {
  auth: true,      // 是否需要 token
  toast: true,     // 错误提示
  catch: true,     // 是否 reject
  loading: true    // 显示 loading
}
```

**使用示例**:
```javascript
import request from '@/config/request'

// 登录
request.post('/login', {
  username: 'admin',
  password: 'admin123',
  code: '1234',
  uuid: 'xxx'
})

// 获取用户信息
request.get('/getInfo')
```

## HTTP 拦截器

文件：`config/http.interceptor.js`

```javascript
// 请求拦截
uni.$u.http.interceptors.request.use(config => {
  // Token 注入
  if (config?.custom?.auth) {
    config.header.Authorization = 'Bearer ' + storage.get('AccessToken')
  }
  // 显示 loading
  if (config?.custom?.loading) {
    uni.showLoading({ title: '加载中...' })
  }
  return config
})

// 响应拦截
uni.$u.http.interceptors.response.use(response => {
  const data = response.data
  uni.hideLoading()
  
  if (data.code !== 200) {
    if (data.code === 401) {
      uni.reLaunch({ url: '/' })  // 跳转登录
      return
    }
    if (config?.custom?.toast !== false) {
      uni.$u.toast(data.msg)
    }
  }
  return data
})
```

## State 管理

### user.js

```javascript
export const state = {
  token: '',
  userInfo: null
}

export const mutations = {
  SET_TOKEN: (state, value) => {
    state.token = value
  },
  SET_USER: (state, value) => {
    state.userInfo = value
  }
}

export const actions = {
  // 登录
  Login({ commit }, data) {
    return new Promise((resolve, reject) => {
      LoginApi.login(data).then(response => {
        commit('SET_TOKEN', response.token)
        resolve(response)
      }).catch(reject)
    })
  },
  
  // 获取用户信息
  Info({ commit, state }) {
    return new Promise((resolve, reject) => {
      UserApi.getInfo().then(response => {
        commit('SET_USER', response)
        resolve(response)
      }).catch(reject)
    })
  }
}
```

### system.js

```javascript
export const state = {
  systemInfo: null  // 设备信息
}

export const actions = {
  SystemInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      uni.getSystemInfo({
        success(res) {
          commit('SET_SYSTEM_INFO', res)
          resolve(res)
        },
        fail(err) {
          reject(err)
        }
      })
    })
  }
}
```

## 缓存工具

文件：`utils/storage.js`

```javascript
const storage = require('@/utils/storage')

// 设置缓存 (带过期时间)
storage.set('key', value, seconds)

// 获取缓存
storage.get('key', defaultValue)

// 删除缓存
storage.remove('key')

// 清空所有缓存
storage.clear()
```

## 表单验证

文件：`utils/verify.js`

```javascript
import { isEmpty, isMobile, isEmail, isNumber } from '@/utils/verify'

// 验证
if (isEmpty(this.form.username)) {
  return uni.$u.toast('用户名不能为空')
}

if (!isMobile(this.form.phone)) {
  return uni.$u.toast('手机号格式不正确')
}

if (!isEmail(this.form.email)) {
  return uni.$u.toast('邮箱格式不正确')
}
```

## 页面结构

```
pages/
├── index/          # 首页
│   └── index.vue   # 统计图表
├── login/          # 登录页
│   └── index.vue
├── work/           # 工作台
│   ├── index.vue
│   ├── user/       # 用户管理
│   │   ├── list.vue
│   │   └── edit.vue
│   └── notice/     # 通知公告
│       ├── list.vue
│       ├── detail.vue
│       └── edit.vue
└── center/         # 个人中心
    ├── index.vue
    ├── profile.vue  # 个人资料
    └── log.vue      # 操作日志
```

## 页面配置

文件：`pages.json`

```javascript
{
  "pages": [
    {
      "path": "pages/login/index",
      "style": {
        "navigationBarTitleText": "登录",
        "navigationStyle": "custom"  // 自定义导航
      }
    }
  ],
  "tabBar": {
    "list": [
      {
        "pagePath": "pages/index/index",
        "iconPath": "static/img/tabbar/home.png",
        "selectedIconPath": "static/img/tabbar/home_art.png",
        "text": "首页"
      }
    ]
  },
  "easycom": {
    "^u-(.*)": "@/uni_modules/uview-ui/components/u-$1/u-$1.vue"
  }
}
```

## UI 组件

### uview-ui 组件

```vue
<template>
  <view>
    <u-button type="primary">按钮</u-button>
    <u-input v-model="value" placeholder="请输入"></u-input>
    <u-popup v-model="show"><view>内容</view></u-popup>
  </view>
</template>
```

### 自定义组件

| 组件 | 路径 | 功能 |
|------|------|------|
| 导航栏 | `components/navbar/` | 自定义导航 |
| 标签栏 | `components/tabs/` | 标签切换 |
| 树形选择器 | `components/qian-tree/` | 支持多选/懒加载 |

## 样式规范

```scss
// 使用 rpx 响应式单位
.container {
  padding: 32rpx;
  background-color: #fff;
}

// 引入 uview-ui 样式
@import "@/uni_modules/uview-ui/index.scss";
```

## App 配置

文件：`App.vue`

```vue
<script>
export default {
  onLaunch() {
    // 加载系统信息
    this.$store.dispatch('SystemInfo')
  },
  onShow() {},
  onHide() {}
}
</script>

<style lang="scss">
@import "@/uni_modules/uview-ui/index.scss";
@import "@/static/style.scss";
</style>
```

## 相关文档

### 项目文档

| 文档 | 说明 |
|------|------|
| [架构设计](ARCHITECTURE.md) | 整体架构 |
| [代码规范](STANDARD.md) | 前端代码规范 |
| [部署指南](DEPLOY.md) | 生产部署 |

### 官方文档

| 技术栈 | 官方文档 |
|--------|----------|
| uni-app | https://uniapp.dcloud.net.cn/ |
| uni-app 组件 | https://uniapp.dcloud.net.cn/component/ |
| uni-app API | https://uniapp.dcloud.net.cn/api/ |
| uni-app CSS | https://uniapp.dcloud.net.cn/css/ |

### uView UI 2.0

| 文档 | 链接 |
|------|------|
| uView 官网 | https://www.uviewui.com/ |
| 组件介绍 | https://www.uviewui.com/components/intro |
| 快速上手 | https://www.uviewui.com/components/quickstart |
| Form 表单 | https://www.uviewui.com/components/form |
| List 列表 | https://www.uviewui.com/components/list |
| Navbar 导航 | https://www.uviewui.com/components/navbar |
| Tabbar 底部导航 | https://www.uviewui.com/components/tabbar |

### 第三方组件

| 组件 | 文档链接 |
|------|----------|
| qiun-data-charts (图表) | https://ext.dcloud.net.cn/plugin?id=271 |
