# PC 前端开发指南

## 目录结构

```
frontend/ruoyi-ui/src/
├── api/              # API 接口 (按模块：monitor/system/tool)
├── assets/           # 静态资源
├── components/       # 公共组件
├── directive/        # 自定义指令
├── layout/           # 布局组件
├── plugins/          # 插件
├── router/           # 路由配置 (动态权限路由)
├── store/            # Vuex 状态管理
│   └── modules/      # app/dict/user/tagsView/permission/settings
├── utils/            # 工具函数
│   ├── request.js    # Axios 封装
│   ├── auth.js       # Token 管理
│   └── permission.js # 路由守卫
└── views/            # 页面视图
    ├── dashboard/    # 仪表盘
    ├── system/       # 系统管理页面
    ├── monitor/      # 监控页面
    └── tool/         # 工具页面
```

## Vue 配置

文件：`vue.config.js`

```javascript
module.exports = {
  publicPath: '/admin/',           // 生产环境路径
  outputDir: 'dist',               // 输出目录
  assetsDir: 'static',             // 静态资源目录
  productionSourceMap: false,      // 生产环境禁用 source map
  devServer: {
    port: 80,
    proxy: {
      [process.env.VUE_APP_BASE_API]: {
        target: 'http://localhost:8080',
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_API]: ''
        }
      }
    }
  }
}
```

## API 请求规范

文件：`src/utils/request.js`

```javascript
import axios from 'axios'
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 10000
})

// Token 注入
config.headers['Authorization'] = 'Bearer ' + getToken()

// 防重复提交 (POST/PUT, 间隔 1s)
const interval = 1000

// 响应错误码处理
const code = res.data.code || 200
```

**使用示例**:
```javascript
import request from '@/utils/request'

export function getUserList(params) {
  return request({
    url: '/system/user/list',
    method: 'get',
    params
  })
}
```

## 路由守卫

文件：`src/permission.js`

```javascript
router.beforeEach((to, from, next) => {
  NProgress.start()
  if (getToken()) {
    if (store.getters.roles.length === 0) {
      // 获取用户信息
      store.dispatch('GetInfo').then(() => {
        // 生成可访问路由
        store.dispatch('GenerateRoutes').then(accessRoutes => {
          router.addRoutes(accessRoutes)
          next({ ...to, replace: true })
        })
      })
    } else {
      next()
    }
  } else {
    next(`/login?redirect=${to.fullPath}`)
  }
})
```

## State 管理

文件：`src/store/modules/`

| Module | 功能 |
|--------|------|
| `app` | 应用状态 (侧边栏、设备类型) |
| `user` | 用户信息 (token/roles/permissions) |
| `dict` | 字典数据缓存 |
| `permission` | 路由权限管理 |
| `settings` | 系统设置 |
| `tagsView` | 标签页管理 |

**使用示例**:
```javascript
// 获取用户信息
this.$store.state.user.userInfo

// 调用 action
this.$store.dispatch('user/GetInfo')

// 提交 mutation
this.$store.commit('user/SET_TOKEN', token)
```

## 路由配置规范

```javascript
{
  path: '/system/user',
  component: Layout,
  meta: { 
    title: '用户管理', 
    icon: 'user' 
  },
  children: [{
    path: 'index',
    component: () => import('@/views/system/user/index'),
    name: 'User',
    meta: { title: '用户管理' }
  }]
}
```

**路由参数**:
- `hidden: true` - 不在侧边栏显示
- `alwaysShow: true` - 始终显示根路由
- `redirect: noRedirect` - 面包屑不可点击
- `meta.roles` - 访问角色权限

## 组件开发

### 模板结构

```vue
<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm">
      <!-- 搜索表单 -->
    </el-form>

    <el-table v-loading="loading" :data="list">
      <!-- 数据表格 -->
    </el-table>

    <pagination
      v-show="total > 0"
      :total="total"
      @update:page="getList"
    />
  </div>
</template>

<script>
import { listUser, getUser, deleteUser } from '@/api/system/user'

export default {
  data() {
    return {
      loading: false,
      list: [],
      total: 0,
      queryParams: {}
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listUser(this.queryParams).then(response => {
        this.list = response.rows
        this.total = response.total
        this.loading = false
      })
    }
  }
}
</script>

<style scoped lang="scss">
.app-container {
  padding: 20px;
}
</style>
```

## 工具函数

| 文件 | 功能 |
|------|------|
| `auth.js` | Token 管理 (Cookies) |
| `validate.js` | 表单验证 |
| `ruoyi.js` | 通用工具 (参数序列化等) |
| `errorCode.js` | 错误码映射 |

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
| Vue 2 | https://v2.vuejs.org/ |
| Element UI | https://element.eleme.cn/ |
| Vue Router | https://v3.router.vuejs.org/ |
| Vuex | https://vuex.vuejs.org/ |
| Axios | https://axios-http.com/ |
| ECharts | https://echarts.apache.org/zh/index.html |

### 第三方组件文档

| 组件 | 文档链接 |
|------|----------|
| vue-treeselect (树形选择器) | https://www.vue-treeselect.info/ |
| vue-cropper (图片裁剪) | https://github.com/xyxiao001/vue-cropper |
| quill (富文本编辑器) | https://quilljs.com/ |
| vuedraggable (拖拽排序) | https://github.com/SortableJS/Vue.Draggable |
| sortablejs (列表排序) | https://sortablejs.github.io/Sortable/ |
| splitpanes (可调整面板) | https://antoniandre.github.io/splitpanes/ |
| screenfull (全屏) | https://github.com/sindresorhus/screenfull.js |
| file-saver (文件下载) | https://github.com/eligrey/FileSaver.js |
