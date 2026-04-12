# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

RuoYi-Go 是一个基于 Go + Gin + Vue 的前后端分离权限管理系统，同时提供原生 Android 版本（UniApp 开发）。

## 快速开始

### 后端启动

```bash
# 安装依赖
go mod download

# 运行项目（默认使用 ./config.yaml 配置文件）
go run main.go

# 指定配置文件运行
go run main.go -f /path/to/config.yaml
```

### 前端启动

**PC 端 (ruoyi-ui)**
```bash
cd frontend/ruoyi-ui
npm install --registry=https://registry.npmmirror.com
npm run dev
# 访问 http://localhost:80
```

**移动端 (RuoYi-Mobile)**
```bash
cd frontend/RuoYi-Mobile
pnpm install
npm run dev
# 访问 http://localhost:9001
```

### Docker 部署

```bash
# 构建镜像并启动容器
docker-compose up -d

# 或使用 start.sh 脚本
./start.sh start
```

## 测试

```bash
# 运行所有测试
go test ./tests/...

# 运行单个测试
go test -run TestPingRoute ./tests/

# 带详细输出
go test -v ./tests/
```

## 代码架构

### 目录结构

```
ruoyi-go/
├── app/                    # 应用主代码
│   ├── admin/             # 管理后台模块
│   │   ├── api/           # 控制器 (API 层)
│   │   ├── model/         # 数据模型
│   │   ├── router/        # 路由配置
│   │   └── service/       # 业务逻辑层
│   ├── core/              # 核心工具
│   │   ├── routers/       # 路由初始化
│   │   └── utils/         # 工具类 (JWT, 密码，文件等)
│   └── html/              # HTML 模板路由
├── config/                # 配置加载
├── pkg/                   # 公共包
│   ├── cache/             # 缓存 (redis/ristretto/file)
│   ├── db/                # 数据库连接
│   ├── logs/              # 日志
│   ├── middleware/        # Gin 中间件
│   └── scheduler/         # 定时任务
├── frontend/              # 前端代码
│   ├── ruoyi-ui/          # PC 端 (Vue2 + ElementUI)
│   └── RuoYi-Mobile/      # 移动端 (UniApp + uview-ui)
├── tests/                 # 测试代码
├── web/                   # 静态资源和 HTML 模板
└── docs/                  # Swagger 文档
```

### 模块划分

**app/admin/api** - 按功能分组：
- `system/` - 系统管理 (用户/角色/菜单/部门/字典/岗位/配置/通知)
- `monitor/` - 监控管理 (缓存/操作日志/登录日志/在线用户/任务/服务监控)
- `tools/` - 工具 (代码生成器)

**数据模型** (`app/admin/model`):
- `system/` - 系统实体 (SysUser, SysRole, SysMenu, SysDept, SysDictData, etc.)
- `monitor/` - 监控实体
- `cache/` - 缓存相关模型

### 技术栈

**后端核心依赖**:
- Web 框架：`gin v1.11.0`
- ORM: `gorm v1.31.1`
- 数据库：`MySQL` (主), `SQLite` (开发)
- 缓存：`redis v8`, `ristretto`, 内存缓存
- JWT: `golang-jwt/jwt`
- 定时任务：`xxl-job`, `robfig/cron`
- 日志：`logrus`, `lfshook`, `file-rotatelogs`
- API 文档：`swaggo/swag`

**前端**:
- PC 端：Vue2 + ElementUI + Vuex + Vue Router
- 移动端：UniApp + uview-ui + qiun-data-charts

### 配置文件

配置文件 `config.yaml` 包含以下部分：
- `server` - 端口、运行模式、日志级别、Swagger 开关
- `database` - 主/次数据源连接
- `redis` - Redis 连接配置
- `jwt` - JWT 密钥和 TTL
- `xxl-job` - XXL-Job 分布式任务调度配置
- `log` - 日志文件路径和过滤规则
- `user-password` - 密码错误锁定策略

### 路由注册流程

1. `main.go` → `routers.Init()` 初始化基础路由
2. `routers.InitWeb()` 注册静态资源和 Swagger
3. `admin.Routers()` 注册业务路由 (system/monitor/tools)
4. JWT 中间件保护需要认证的路由

### 定时任务

- 本地定时任务：`pkg/scheduler/scheduler.go` 基于 `robfig/cron`
- 分布式任务：`xxl-job` 集成在 `app/core/utils/job/xxljob.go`

## 开发注意事项

1. **数据库**：默认使用 SQLite (`sqlite-ry-go.db`) 进行开发，生产环境切换 MySQL
2. **跨域**：`app/core/utils/CrossCode.go` 处理 CORS
3. **统一响应**：`app/core/utils/R/R.go` 定义 API 响应格式
4. **密码加密**：使用 bcrypt 加密
5. **代码生成**：`app/admin/service/tools/gen.go` 提供代码生成功能

## 常见问题

**Node 版本要求**：前端需要 Node 8/10/12/14/16/17/19/20，如遇 OpenSSL 错误：
```bash
export NODE_OPTIONS=--openssl-legacy-provider
```

## 代码规范

### Go 后端规范

**命名约定**
- 模型类：大驼峰命名，带前缀 `Sys`，如 `SysUser`, `SysRole`, `SysMenu`
- 参数类：模型名 + `Parm`，如 `SysUserParm`, `SysUserExcel`
- Handler 函数：`功能名 + Handler`，如 `LoginHandler`, `ListUser`
- 路由初始化：`Init + 模块名`，如 `InitUser`, `InitMenu`

**API 响应格式** (`app/core/utils/R/R.go`)
```go
// 成功返回数据
R.ReturnSuccess(data)

// 成功返回消息
R.ReturnSuccessMsg("操作成功")

// 失败返回消息
R.ReturnFailMsg("错误信息")
```

**响应结构**
```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}
```

**GORM 模型定义**
- 主键：`gorm:"primaryKey"`
- 列名：`gorm:"column:field_name"`
- 时间字段：`gorm:"autoCreateTime"` 自动填充

**Handler 编写规范**
1. 参数绑定使用 `context.ShouldBind(&param)`
2. 参数校验失败返回 `gin.H{"msg": "...", "code": ...}`
3. 业务逻辑封装到 model 或 service 层
4. 统一使用 `R.Result` 包装返回值

**路由规范**
- 分组前缀：`e.Group("模块名")`
- 认证中间件：`jwt.JWTAuthMiddleware()`
- 路由函数：`Init + 模块名 (e *gin.Engine)`

### 前端规范

**EditorConfig** (`.editorconfig`)
- 字符集：`utf-8`
- 缩进：2 空格
- 换行符：`lf`
- 文件末尾新行：`true`
- 去除行尾空格：`true`

---

### PC 端规范 (`frontend/ruoyi-ui`)

**目录结构**
```
src/
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
│   ├── request.js    # Axios 封装 (拦截器/防重复提交)
│   ├── auth.js       # Token 管理 (Cookies)
│   └── permission.js # 路由守卫
└── views/            # 页面视图
    ├── dashboard/    # 仪表盘
    ├── system/       # 系统管理页面
    ├── monitor/      # 监控页面
    └── tool/         # 工具页面
```

**Vue 配置** (`vue.config.js`)
- 基础路径：`/admin/` (生产环境)
- 开发端口：80
- 后端代理：`http://localhost:8080`
- 输出目录：`dist`
- 资源目录：`static`

**API 请求规范** (`utils/request.js`)
```javascript
// Token 注入
config.headers['Authorization'] = 'Bearer ' + getToken()

// 防重复提交 (POST/PUT, 间隔 1s)
const interval = 1000

// 响应错误码处理
const code = res.data.code || 200
```

**路由守卫** (`permission.js`)
1. 检查 Token 是否存在
2. 有 Token → 调用 `GetInfo` 获取用户信息
3. 根据角色生成可访问路由 (`GenerateRoutes`)
4. 动态添加路由 (`router.addRoutes`)

**State 管理** (`store/modules/`)
- `app` - 应用状态 (侧边栏、设备类型)
- `user` - 用户信息 (token/roles/permissions)
- `dict` - 字典数据缓存
- `permission` - 路由权限管理
- `settings` - 系统设置
- `tagsView` - 标签页管理

**路由配置规范**
```javascript
{
  path: '/system/user',
  component: Layout,
  meta: { title: '用户管理', icon: 'user' },
  children: [{
    path: 'index',
    component: () => import('@/views/system/user/index'),
    name: 'User',
    meta: { title: '用户管理' }
  }]
}
```

---

### 移动端规范 (`frontend/RuoYi-Mobile`)

**目录结构**
```
RuoYi-Mobile/
├── api/              # API 接口
├── components/       # 自定义组件 (navbar/tabs/qian-tree)
├── config/
│   └── environment.js  # 环境配置 (开发/生产)
├── pages/            # 页面 (index/login/work/center)
├── store/            # Vuex 状态管理
├── uni_modules/      # Uni 模块插件 (uview-ui/qiun-data-charts)
└── utils/            # 工具函数
```

**环境配置** (`config/environment.js`)
```javascript
{
  development: { baseURL: 'http://localhost:8080' },
  production: { baseURL: 'http://vue.ruoyi.vip/prod-api' }
}
```

**页面配置** (`pages.json`)
- 启动页：`pages/login/index`
- TabBar: 首页/工作台/我的
- 自定义导航：`navigationStyle: "custom"`
- 自动组件引入：`easycom` 配置 uview-ui

**App 配置** (`manifest.json`)
- 应用名称：若依 Go 管理系统
- 版本：1.0.2
- 支持平台：H5 / 小程序 / App

**请求封装** (`config/request.js`)
```javascript
// 默认配置
{
  auth: true,      // 是否需要 token
  toast: true,     // 错误提示
  catch: true,     // 是否 reject
  loading: true    // 显示 loading
}

// Token 注入：config.header.Authorization = 'Bearer ' + storage.get('AccessToken')
```

**响应拦截器** (`config/http.interceptor.js`)
- 自动关闭 loading
- 错误码处理:
  - `code !== 200`: toast 提示
  - `code === 401`: 跳转登录
- 支持自定义 catch 返回

**State 管理** (`store/modules/`)
- `user.js` - 用户认证 (token/userInfo/登录/登出)
- `system.js` - 系统信息 (设备信息获取与缓存)

**缓存工具** (`utils/storage.js`)
- 支持过期时间设置：`storage.set(key, value, seconds)`
- 自动清理过期数据
- 数据类型：string/json/array/boolean

**页面结构**
```
pages/
├── index/        # 首页 (统计图表)
├── login/        # 登录页
├── work/         # 工作台
│   ├── index.vue
│   ├── user/     # 用户管理
│   └── notice/   # 通知公告
└── center/       # 个人中心
    ├── index.vue
    ├── profile/  # 个人资料
    └── log/      # 操作日志
```

**工具函数** (`utils/verify.js`)
- 表单验证：手机号/邮箱/整数/浮点数/正整数

**UI 组件**
- 导航栏：`components/navbar/`
- 标签栏：`components/tabs/`
- 树形选择器：`components/qian-tree/` (支持多选/懒加载)

---

### 前端通用规范

**组件命名**
- Vue 组件：PascalCase (如 `SysUser.vue`)
- 文件命名：kebab-case (如 `user-list.vue`)

**API 接口规范**
- 按模块划分：`api/system.js`, `api/monitor/`
- 统一使用 request 封装
- 返回解构：`const { code, msg, data } = response`

**样式规范**
- PC 端：SCSS + ElementUI 主题变量
- 移动端：SCSS + uview-ui 主题变量
- 响应式单位：rpx (移动端), px (PC 端)
