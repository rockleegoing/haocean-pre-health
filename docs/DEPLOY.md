# 部署运维指南

## 环境要求

### 后端
- Go 1.24+
- MySQL 5.7+ / SQLite
- Redis 6.0+ (可选，用于缓存和会话)

### 前端
- Node.js 16/18/20
- npm / pnpm

### Docker (可选)
- Docker 20+
- Docker Compose 1.29+

---

## ⚠️ 重要说明

**默认部署方式**：本文档默认描述的是**开发环境部署**。如需生产部署，请参考 [生产环境部署](#生产环境部署) 章节。

**开发模式核心原理**：
- 后端 (8080) 通过反向代理将前端请求转发到前端开发服务器
- PC 前端开发服务器运行在 1024 端口，移动端运行在 9001 端口
- 访问 `http://localhost:8080/admin` 和 `http://localhost:8080/mobile` 即可访问最新代码
- 前端修改后自动热更新，无需手动构建

---

## 快速开始（开发环境）

### 1. 获取代码

```bash
git clone https://gitee.com/OptimisticDevelopers/Ruoyi-Go.git
```

### 2. 数据库配置

**导入 SQL 到 MySQL**：
- 数据库文件：`config/sql/ry-go.sql`

**拉取依赖**：
```bash
go mod tidy
# 如超时，设置国内代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

**创建配置文件**：
```bash
cat config/config.yaml.example > config.yaml
```

**配置文件说明**：
```yaml
server:
  port: 8080
  runMode: debug  # 开发环境必须设为 debug
  enabledSwagger: true

database:
  primary:
    enabled: true
    type: mysql
    host: localhost
    port: "3306"
    username: root
    password: your_password
    dbname: ruoyi_go

redis:
  host: localhost
  port: "6379"
  password: ""
  db: 0

jwt:
  secret: your_secret_key  # 自定义密钥
  jwt_ttl: 86400
```

---

### 3. 一键启动（推荐）

使用一键启动脚本同时启动后端、PC 前端和移动端 H5：

```bash
chmod +x start-dev.sh
./start-dev.sh
```

**访问地址：**
- PC 前端：http://localhost:8080/admin
- 移动端 H5: http://localhost:8080/mobile
- 后端 API: http://localhost:8080/api
- Swagger: http://localhost:8080/swagger

**优势：**
- 支持热更新，修改前端代码自动刷新
- 无需手动构建前端
- 统一访问入口（8080 端口）
- 三个服务同时启动，一键管理

**停止服务**：按 `Ctrl+C` 停止所有服务

---

### 3. 手动启动（开发环境）

```bash
# 1. 启动后端（确保 config.yaml 中 runMode: debug）
cd /Users/arthur/Documents/workspace/haocean/project_v3/Ruoyi-Go
go run main.go

# 2. 启动 PC 前端（新终端）
cd frontend/ruoyi-ui
npm run dev
# 启动后访问：http://localhost:8080/admin

# 3. 启动移动端 H5（新终端）
cd frontend/RuoYi-Mobile
pnpm dev 2>/dev/null || npm run dev
# 启动后访问：http://localhost:8080/mobile
```

**移动端开发建议使用 HBuilderX**：
1. 下载并安装 [HBuilderX](https://www.dcloud.io/hbuilderx.html)
2. 打开 HBuilderX，导入项目：`frontend/RuoYi-Mobile`
3. 点击工具栏 **运行** → **运行到浏览器-H5**
4. 移动端会自动启动到 `http://localhost:9001`
5. 通过后端访问：`http://localhost:8080/mobile`

---

### 开发模式配置说明

开发模式下（`config.yaml` 中 `server.runMode: debug`），后端会将前端请求**反向代理**到前端开发服务器：

| 路径 | 代理目标 | 说明 |
|------|----------|------|
| `/admin/*` | `http://localhost:1024/admin/*` | PC 前端（Vue2 + ElementUI） |
| `/mobile/*` | `http://localhost:9001/mobile/*` | 移动端 H5（uni-app） |
| `/api/*` | 后端直接处理 | API 接口 |
| `/swagger/*` | 后端直接处理 | Swagger 文档 |

**前端 API 请求流程**：
1. PC 前端请求 `/dev-api/xxx` → 代理到 `http://localhost:8080/xxx`
2. 移动端请求 `/prod-api/xxx` → 代理到 `http://localhost:8080/xxx`
3. 后端接收请求并处理

**配置文件检查**：

1. `config.yaml` - 确保开发模式：
```yaml
server:
  port: 8080
  runMode: debug  # 开发环境必须设为 debug
```

2. `frontend/ruoyi-ui/.env.development` - PC 前端 API 配置：
```
VUE_APP_BASE_API = '/dev-api'
```

3. `frontend/ruoyi-ui/vue.config.js` - PC 前端端口和 publicPath：
```javascript
const port = process.env.port || process.env.npm_config_port || 1024
publicPath: process.env.NODE_ENV === "production" ? "/admin/" : "/"
```

4. `frontend/RuoYi-Mobile/vue.config.js` - 移动端端口和代理：
```javascript
module.exports = {
  devServer: {
    port: 9001,
    proxy: {
      '/prod-api': {
        target: 'http://localhost:8080',
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/prod-api': ''
        }
      }
    }
  }
}
```

5. `frontend/RuoYi-Mobile/config/environment.js` - 移动端 API 配置：
```javascript
module.exports = {
  development: {
    baseURL: '/prod-api'  // 使用代理路径
  },
  production: {
    baseURL: 'http://your-domain.com/prod-api'
  }
}
```

6. `frontend/RuoYi-Mobile/config/http.interceptor.js` - 移动端请求配置：
```javascript
// 开发模式
config.baseURL = '/prod-api'; /* 代理路径前缀 */
```

---

### 开发环境常见问题

**问题 1：前端页面加载失败或静态资源 404**
- 确保后端正在运行且 `config.yaml` 中 `runMode: debug`
- 确保前端开发服务器正在运行（PC 前端 1024 端口，移动端 9001 端口）
- 清除浏览器缓存或使用无痕模式

**问题 2：API 请求返回 HTML 而不是 JSON**
- 移动端检查 `http.interceptor.js` 中 `baseURL` 是否为 `/prod-api`
- PC 前端检查 `.env.development` 中 `VUE_APP_BASE_API` 是否为 `/dev-api`

**问题 3：登录后 401 认证失败**
- 检查后端 JWT 配置 `config.yaml` 中 `jwt.secret`
- 清除浏览器 localStorage 和 cookie
- 重新登录

**问题 4：修改前端代码后不更新**
- 开发模式应该自动热更新
- 如果未更新，尝试手动刷新浏览器
- 检查前端开发服务器是否正常运行

---

## 生产环境部署

> **注意**：生产部署需要手动执行，默认部署为开发环境。

### 1. 前端构建

**PC 前端：**
```bash
cd frontend/ruoyi-ui
npm run build:prod
# 输出到 dist/ 目录
# 将 dist 目录复制到 ./web/admin
cp -r dist/* ../../web/admin/
```

**移动端 H5：**

方式一：使用 HBuilderX（推荐）
1. 打开 HBuilderX，导入项目 `frontend/RuoYi-Mobile`
2. 点击 **发行** → **网站-H5 手机版**
3. 构建完成后，将输出的 `dist/build/h5/` 目录复制到 `./web/mobile/static/`

方式二：使用 CLI（需要配置）
```bash
cd frontend/RuoYi-Mobile
npm run build:h5
# 输出到 dist/build/h5/ 目录
cp -r dist/build/h5/* ../../web/mobile/static/
```

---

### 2. 后端配置

确保 `config.yaml` 配置：
```yaml
server:
  port: 8080
  runMode: release  # 生产模式必须设为 release
  logLevel: info    # 生产环境建议设为 info 或 warn
  enabledSwagger: false  # 生产环境建议关闭 Swagger
```

---

### 3. 启动方式

**方式一：直接启动**
```bash
go build -o main main.go
./start.sh start
```

**方式二：Docker 部署**

```dockerfile
FROM alpine:3.12
MAINTAINER 乐观开发者
RUN mkdir -p "/data/app" \
&& ln -sf /dev/stdout /data/app/service.log
WORKDIR "/data/app"
COPY main /data/app/main
ENTRYPOINT ["./main"]
```

```yaml
# docker-compose.yml
version: '3.3'
services:
  webapp:
    restart: always
    build:
      context: ./
      dockerfile: Dockerfile
    image: haocean-health:1.0.0
    container_name: haocean-health
    ports:
      - "8080:8080"
```

```bash
# 1. 构建 Go 二进制
go build -o main main.go

# 2. 构建 Docker 镜像并启动
docker-compose up -d

# 3. 查看日志
docker logs -f haocean-health

# 4. 停止服务
docker-compose down
```

### 访问地址

- PC 前端：http://your-domain.com/admin
- 移动端 H5: http://your-domain.com/mobile
- 后端 API: http://your-domain.com/prod-api
- Swagger: http://your-domain.com/swagger（如启用）

### 生产模式说明

生产模式下（`config.yaml` 中 `server.runMode: release`），后端直接**提供静态文件服务**：

| 路径 | 静态文件目录 | 说明 |
|------|--------------|------|
| `/admin/*` | `./web/admin/*` | PC 前端构建产物 |
| `/mobile/*` | `./web/mobile/static/*` | 移动端 H5 构建产物 |
| `/favicon.ico` | `./web/admin/favicon.ico` | 网站图标 |
| `/profile/*` | `./static/images/*` | 用户头像等静态资源 |

**注意事项：**
- 生产环境下，前端修改后必须重新构建并复制文件，否则不会生效
- 建议配置 Nginx 反向代理、Gzip 压缩、HTTPS 等

---

## 运维脚本

### start.sh（生产环境）

```bash
#!/bin/bash
APP_NAME=main

start() {
  nohup ./${APP_NAME} > ./log.txt 2>&1 &
  echo "${APP_NAME} start success"
}

stop() {
  pid=$(ps -ef|grep $APP_NAME|grep -v grep|awk '{print $2}')
  if [ -n "${pid}" ]; then
    kill -9 $pid
  fi
}

restart() {
  stop
  start
}

status() {
  pid=$(ps -ef|grep $APP_NAME|grep -v grep|awk '{print $2}')
  if [ -n "${pid}" ]; then
    echo "${APP_NAME} is running. Pid is ${pid}"
  else
    echo "${APP_NAME} is NOT running."
  fi
}

case "$1" in
  "start") start ;;
  "stop") stop ;;
  "status") status ;;
  "restart") restart ;;
  *) echo "Usage: $0 {start|stop|status|restart}" ;;
esac
```

### start-dev.sh（开发环境）

```bash
#!/bin/bash
# 一键启动开发环境：后端 + PC 前端 + 移动端
chmod +x start-dev.sh
./start-dev.sh
```

---

## Nginx 配置（生产环境）

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # PC 前端
    location /admin {
        alias /path/to/web/admin;
        try_files $uri $uri/ /admin/index.html;
    }

    # 后端 API 代理
    location /prod-api {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }

    # 移动端 H5
    location /mobile {
        alias /path/to/web/mobile/static;
        try_files $uri $uri/ /mobile/index.html;
    }

    # 静态资源
    location /profile {
        alias /path/to/static/images;
    }
}
```

---

## 移动端打包

### H5 生产构建

参考上文 [生产环境部署 - 前端构建](#1-前端构建) 章节。

### Android App 打包

**云打包方式（推荐）**：

1. 修改 `frontend/RuoYi-Mobile/config/environment.js` 中的 `production.baseURL` 为生产环境后端地址
2. 打开 HBuilderX
3. 菜单栏：**发行** → **原生 App-云打包**
4. 选择打包类型：Android 包
5. 点击打包，等待完成后下载 APK

**离线打包方式**：

需要配置 Android 开发环境，参考 [uni-app 离线打包文档](https://nativesupport.dcloud.net.cn/AppDocs/download/android)

---

## 日志系统

### 日志模式

配置项：`config.LogConfig.LogMode`

| 模式 | 说明 |
|------|------|
| `default` | 标准输出 |
| `file` | 文件日志 (按日轮转，保留 7 天) |
| `mysql` | 操作日志写入数据库 |
| `es` | Elasticsearch (待实现) |

### 日志配置

```yaml
log:
  enabled: true
  logMode: file
  filePath: ./log/app.log
  filtered:
    - /ping
    - /swagger/*
```

---

## 定时任务

### 本地 Cron

- 文件：`pkg/scheduler/cron.go`
- 基于 `robfig/cron`
- 自动加载数据库中的任务配置

### XXL-Job (分布式)

配置项：
```yaml
xxl-job:
  enabled: true
  admin_address: http://xxl-job-admin:8080
  app_name: haocean-health-executor
  port: 9999
  log_path: /data/logs/xxljob
```

---

## 数据库连接池

```yaml
database:
  max_idle_conn: 10      # 最大空闲连接
  max_open_conn: 100     # 最大打开连接
  conn_max_lifetime: 60  # 连接最大生命周期 (分钟)
```

**主从切换**:
```go
mysql.MysqlDb()      // 主库
mysql.MysqlDb(true)  // 从库
```

---

## 优雅关闭

系统支持优雅关闭，处理流程：

1. 捕获 `SIGINT` / `SIGTERM` 信号
2. 停止接收新请求
3. 等待 5 秒处理完现有请求
4. 关闭 HTTP 服务
5. 关闭数据库连接

文件：`app/core/utils/shutdown/gracefully_shutdown.go`

---

## 健康检查

```bash
curl http://localhost:8080/ping
# 返回：{"message":"pong"}
```

---

## 相关文档

- [架构设计](ARCHITECTURE.md) - 整体架构
- [后端开发](BACKEND.md) - Go 后端开发
- [前端开发](FRONTEND.md) - PC 前端开发
- [移动端开发](MOBILE.md) - UniApp 开发
- [Git 指南](GIT.md) - 版本控制
- [代码规范](STANDARD.md) - 代码规范检查
