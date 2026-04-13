---
name: 项目改名记录
description: 2026-04-13 将项目从 RuoYi-Go 改名为 Haocean-HealthEnforcement（移动卫生执法系统）
type: project
---

## 项目改名记录

**日期**: 2026-04-13

**原名**: RuoYi-Go（若依 Go）

**新名**: 
- 英文：Haocean HealthEnforcement
- 中文：移动卫生执法系统

---

## 改名内容

### 1. Go Module
- `go.mod`: `module ruoyi-go` → `module haocean/health-enforcement`
- 所有 Go 文件 import 路径：`ruoyi-go/xxx` → `haocean/health-enforcement/xxx`

### 2. 前端目录
- `frontend/ruoyi-ui/` → `frontend/health-enforcement-ui/`
- `frontend/RuoYi-Mobile/` → `frontend/Haocean-Mobile/`

### 3. 配置文件
- `config/config.yaml.example`: 数据库名 `ry-go` → `haocean_health`
- `config/sql/ry-go.sql` → `config/sql/haocean-health.sql`

### 4. 前端配置
- `frontend/health-enforcement-ui/package.json`: name 和 description 更新
- `frontend/Haocean-Mobile/manifest.json`: name 更新为"移动卫生执法系统"

### 5. 文档
- `README.md`: 项目名称和描述更新
- `docs/*.md`: 所有文档中的项目名称引用更新

### 6. 启动脚本
- `start-dev.sh`: 路径和输出信息更新

---

## 验证

- [x] `go mod tidy` 成功
- [x] `go build -o health-enforcement main.go` 编译成功
- [ ] 前端开发服务器启动测试
- [ ] 登录功能测试

---

## 注意事项

1. 如需要推送 Git 仓库，需要更新远程仓库地址
2. 数据库名已改为 `haocean_health`，需要重新导入 SQL 文件
3. JWT secret 建议更换新的密钥
