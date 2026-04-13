# 官方文档笔记

来源：https://haocean-health.qiqjia.com

## 项目概述

**Haocean HealthEnforcement** 是一个基于 Go 语言开发的权限管理系统，结合 Gin、Gorm 等流行框架，旨在提供高效、易用的后台管理系统解决方案。

### 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Go + Gin (Web 框架) + Gorm (ORM) + JWT (认证) |
| 前端 | Vue.js (前后端分离) |
| 移动端 | Kotlin (Android 原生) / UniApp |

### 设计理念

- 沿袭 RuoYi 的设计理念
- 注重代码简洁性和可维护性
- 支持代码快速生成和修改
- 满足不同项目需求

---

## 主要功能

| 功能模块 | 说明 |
|----------|------|
| 用户管理 | 系统用户配置、权限分配 |
| 部门管理 | 组织机构配置（公司/部门/小组），树形结构，数据权限控制 |
| 岗位管理 | 配置用户所属职务 |
| 菜单管理 | 菜单、操作权限、按钮权限标识配置 |
| 角色管理 | 角色菜单权限分配、数据范围权限划分 |
| 字典管理 | 维护常用固定数据 |
| 日志管理 | 操作日志、系统异常日志、登录日志 |
| 在线用户 | 监控活跃用户状态 |
| 定时任务 | 在线增删改任务调度，含执行结果日志 |
| API 文档生成 | 根据业务代码自动生成 API 接口文档 |
| 服务监控 | 监视 CPU、内存、磁盘、堆栈等系统信息 |

---

## 快速开始

### 获取代码

```bash
git clone https://gitee.com/OptimisticDevelopers/Ruoyi-Go.git
```

### 数据库配置

1. **导入 SQL 到 MySQL**
   - 数据库文件：`config/sql/ry-go.sql`
   - 后台前端地址：`frontend/ruoyi-ui`（已修改访问路径为 `/admin`）

2. **拉取依赖**
```bash
go mod tidy
# 如超时，设置国内代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

3. **创建配置文件**
```bash
cat config/config.yaml.example > config.yaml
# 需要配置 MySQL 和 Redis 账号
```

### 启动项目

```bash
go run main.go
```

### 访问地址

- 后台：http://127.0.0.1:8080/admin

---

## 部署指南

### 环境要求

- **操作系统**: Linux/macOS (支持 Go 的环境)
- **Go 环境**: 安装 Go 并添加到 PATH
- **数据库**: MySQL
- **缓存**: Redis

### 部署步骤

1. **获取代码**
```bash
git clone <repository>
```

2. **调整配置**
   - 修改数据库配置（地址、端口、账号）
   - 更新 Redis 配置（如需要）

3. **处理依赖**
```bash
go mod tidy
go mod vendor  # 如需复制依赖
```

4. **编译项目**
```bash
go build
# 跨平台编译示例
GOOS=linux GOARCH=arm GOARM=7 go build
```

5. **服务器部署**
   - 上传文件：`view/`、`config.yaml`、`haocean-health`
   - 设置执行权限：`chmod +x haocean-health`
   - 启动服务：`nohup ./haocean-health`

6. **验证**
   - 访问：`http://ip:8080`
   - 检查日志

**重要**: `view`、`config.yaml`、`haocean-health` 必须在同一目录下。

---

## 路由系统

### Gin 路由原理

Gin 框架的路由系统基于 **httprouter**，利用 **Radix Tree（基数树）** 实现高效路由匹配。

#### 路由匹配机制

1. **Radix Tree**: 高效的前缀树结构，用于快速查找和匹配 URL 路径
2. **路由树**: 每个请求方法（GET/POST 等）维护单独的树
3. **优先级排序**: 子节点按优先级排序（基于注册的句柄数量）

### 路由配置

#### 基础路由

```go
// 路由位置：app/admin/router.go

// 需要登录才能访问的路由
v1 := e.Group("/")
{
    auth := v1.Group("")
    auth.Use(jwt.JWTAuthMiddleware())
    {
        auth.GET("getInfo", handlerFunc, api.GetInfoHandler)
        auth.GET("getRouters", handlerFunc, api.GetRoutersHandler)
    }
}
```

#### 添加新模块

```go
func InitDemo(e *gin.Engine) {
    // 不需登录的路由
    e.GET("/index", api.IndexHandler)
    
    // 需要登录的路由
    v := e.Group("demo")
    {
        auth := v.Group("")
        auth.Use(jwt.JWTAuthMiddleware())
        {
            auth.GET("/hello", demo.Hello)
        }
    }
}

// 在 router.go 中注册
demo.InitDemo(e)
```

---

## Android 移动端 (Kotlin)

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

### 最低 SDK 版本

```
min_sdk: 21
```

### 获取代码

```bash
git clone git@gitee.com:OptimisticDevelopers/Ruoyi-Android-App.git
cd Ruoyi-Android-App
```

### 常见问题

**错误 1**: `Unresolved reference: BuildConfig`
- 解决方法：重新编译项目

**错误 2**: `BaseRecyclerViewAdapterHelper` 下载失败
- 解决方法：在 `repositories` 中添加阿里云镜像
```gradle
maven { url 'https://maven.aliyun.com/repository/google' }
maven { url 'https://maven.aliyun.com/repository/central' }
maven { url 'https://maven.aliyun.com/repository/public' }
```

---

## UniApp 移动端

### 技术栈

| 组件 | 版本 |
|------|------|
| uview-ui | 2.0.31 |
| qiun-data-charts | 2.4.3-20220505 |
| vue | 2 或 3 |

### 配置

**文件**: `config/environment.js`

```javascript
{
  development: {
    baseURL: 'http://localhost:8080'
  },
  production: {
    baseURL: 'http://vue.ruoyi.vip/prod-api'
  }
}
```

### H5 启动端口

**文件**: `vue.config.js`

```javascript
module.exports = {
  devServer: {
    port: 9001,
    proxy: {
      '/': {
        target: environment.baseURL,
        changeOrigin: true,
        pathRewrite: {
          '^/': ''
        }
      }
    }
  }
}
```

**注意**: 不要在 `manifest.json` 中配置 H5 启动信息，可能引发后端接口访问异常。

---

## 相关链接

| 资源 | 链接 |
|------|------|
| Gitee 仓库 | https://gitee.com/OptimisticDevelopers/Ruoyi-Go |
| 官方文档 | https://haocean-health.qiqjia.com |
| 微信公众号文章 | https://mp.weixin.qq.com/s/XIX0qQCm4oH5wYTIHCJHCg |

---

## 版权信息

本文档内容版权属于 Haocean HealthEnforcement 作者，保留所有权利。
Copyright © 2024-present 乐观开发者
