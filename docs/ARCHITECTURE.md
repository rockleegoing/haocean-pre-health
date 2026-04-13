# 架构设计

## 项目概述

**Haocean HealthEnforcement** 是一个基于 Go 语言开发的权限管理系统，结合 Gin、Gorm、JWT 等框架，提供高效、易用的后台管理系统解决方案。

**设计理念**:
- 沿袭 RuoYi 的设计理念
- 注重代码简洁性和可维护性
- 支持代码快速生成和修改

## 主要功能模块

| 模块 | 功能说明 |
|------|----------|
| 用户管理 | 系统用户配置、权限分配 |
| 部门管理 | 组织机构配置（树形结构）、数据权限控制 |
| 岗位管理 | 配置用户所属职务 |
| 菜单管理 | 菜单、操作权限、按钮权限标识配置 |
| 角色管理 | 角色菜单权限分配、数据范围权限划分 |
| 字典管理 | 维护常用固定数据 |
| 日志管理 | 操作日志、系统异常日志、登录日志 |
| 在线用户 | 监控活跃用户状态 |
| 定时任务 | 在线增删改任务调度，含执行结果日志 |
| API 文档生成 | 根据业务代码自动生成 API 接口文档 |
| 服务监控 | 监视 CPU、内存、磁盘、堆栈等系统信息 |

## 目录结构

```
haocean-health/
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

## 模块划分

### app/admin/api

| 模块 | 路径 | 功能 |
|------|------|------|
| system | `api/system/` | 用户、角色、菜单、部门、字典、岗位、配置、通知 |
| monitor | `api/monitor/` | 缓存、操作日志、登录日志、在线用户、任务、服务监控 |
| tools | `api/tools/` | 代码生成器、通用工具 |

### app/admin/model

| 模块 | 路径 | 功能 |
|------|------|------|
| system | `model/system/` | SysUser, SysRole, SysMenu, SysDept, SysDictData 等实体 |
| monitor | `model/monitor/` | 监控实体 |
| cache | `model/cache/` | 缓存相关模型 |

## 技术架构

### 后端核心依赖

| 用途 | 库 | 版本 |
|------|-----|------|
| Web 框架 | gin | v1.11.0 |
| ORM | gorm | v1.31.1 |
| JWT | golang-jwt/jwt | - |
| 缓存 | redis v8, ristretto | - |
| 定时任务 | xxl-job, robfig/cron | - |
| 日志 | logrus, lfshook | - |
| API 文档 | swaggo/swag | v1.16.6 |

### 前端

| 端 | 技术栈 |
|----|--------|
| PC | Vue2 + ElementUI + Vuex + Vue Router |
| 移动端 | UniApp + uview-ui + qiun-data-charts |

## 核心流程

### 路由注册流程

```
main.go
  ├── routers.Init()         # 初始化基础路由
  │     ├── 日志中间件
  │     ├── Recovery 中间件
  │     ├── CORS 处理
  │     └── /ping 健康检查
  │
  ├── routers.InitWeb()      # 注册静态资源和 Swagger
  │     ├── /profile 静态资源
  │     ├── /admin 前端页面
  │     └── /swagger API 文档
  │
  └── admin.Routers()        # 注册业务路由
        ├── system (用户/角色/菜单等)
        ├── monitor (缓存/日志/任务等)
        └── tools (代码生成)
```

### 认证流程

1. 用户登录 → 生成 JWT Token
2. Token 存入 Redis（支持单点登录控制）
3. 后续请求携带 `Authorization: Bearer {token}`
4. JWT 中间件验证：
   - Token 签名
   - 有效期
   - Redis 中的用户状态

## 定时任务

### 本地 Cron (`pkg/scheduler/cron.go`)

- 基于 `robfig/cron`
- 支持方法调用：
  - `ryTask.ryNoParams` - 无参数方法
  - `ryTask.ryParams` - 单参数方法
  - `ryTask.ryMultipleParams` - 多参数方法
- 自动加载数据库中的任务配置 (`sys_job` 表)

### XXL-Job (分布式)

- 配置项：`config.XxlJob`
- 执行器端口：9999
- 集成文件：`app/core/utils/job/xxljob.go`

## 部署架构

```
┌─────────────────────────────────────────────────────────┐
│                     Nginx (可选)                         │
│                  反向代理 + 负载均衡                      │
└───────────────────────┬─────────────────────────────────┘
                        │
        ┌───────────────┼───────────────┐
        │               │               │
┌───────▼───────┐ ┌────▼───────┐ ┌────▼────────┐
│   后端 API     │ │  PC 前端   │ │  移动端 H5   │
│   :8080       │ │  /admin/   │ │  :9001      │
│   (Go+Gin)    │ │  (Vue2)    │ │  (UniApp)   │
└───────┬───────┘ └────────────┘ └─────────────┘
        │
   ┌────▼────┐  ┌────▼────┐
   │  MySQL  │  │  Redis  │
   │  数据库  │  │  缓存    │
   └─────────┘  └─────────┘
```

## 相关文档

- [后端开发](BACKEND.md) - Go API 开发指南
- [前端开发](FRONTEND.md) - PC 端 Vue 开发
- [移动端开发](MOBILE.md) - UniApp 开发
- [部署指南](DEPLOY.md) - 生产环境部署
