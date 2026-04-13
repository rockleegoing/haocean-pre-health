# 部署约定（重要）

## 部署模式默认值

**默认部署方式：开发环境部署**

在后续对话中，除非明确指定"生产部署"或"生产环境"，否则所有部署相关的指令和讨论都默认指**开发环境部署**。

---

## 部署模式对比

| 项目 | 开发环境（默认） | 生产环境 |
|------|-----------------|---------|
| 配置文件 | `runMode: debug` | `runMode: release` |
| 前端服务 | 开发服务器（热更新） | 静态文件服务 |
| 启动方式 | `./start-dev.sh` | `./start.sh start` |
| 访问地址 | http://localhost:8080 | http://your-domain.com |
| Swagger | 启用 | 建议关闭 |
| 日志级别 | debug | info/warn |

---

## 开发环境部署（默认）

### 一键启动

```bash
chmod +x start-dev.sh
./start-dev.sh
```

### 访问地址

- PC 前端：http://localhost:8080/admin
- 移动端 H5: http://localhost:8080/mobile
- 后端 API: http://localhost:8080/api
- Swagger: http://localhost:8080/swagger

### 特点

- 支持热更新
- 无需手动构建前端
- 反向代理模式
- 适合开发和调试

---

## 生产环境部署（需明确指定）

当用户说"生产部署"、"发布到生产"、"部署到线上"等时，才执行生产环境部署。

### 部署步骤

1. 前端构建（`npm run build:prod`）
2. 配置文件修改（`runMode: release`）
3. 后端编译（`go build -o main main.go`）
4. 启动服务（`./start.sh start`）

### 特点

- 静态文件服务
- 需要手动构建前端
- 适合正式发布

---

## 对话中的默认行为

| 用户指令 | 默认解释 |
|---------|---------|
| "启动项目" | 开发环境启动 |
| "部署" | 开发环境部署 |
| "重启服务" | 开发环境重启 |
| "生产部署" | 生产环境部署（明确指定） |
| "发布到线上" | 生产环境部署（明确指定） |

---

## 配置文件参考

### 开发环境（默认）

```yaml
server:
  port: 8080
  runMode: debug
  enabledSwagger: true
```

### 生产环境（需指定）

```yaml
server:
  port: 8080
  runMode: release
  enabledSwagger: false
```
