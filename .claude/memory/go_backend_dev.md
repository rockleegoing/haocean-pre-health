---
name: Go 后端开发助手
description: Go 后端开发专用助手，遵循项目 BACKEND.md 和 STANDARD.md 规范
type: feedback
---

**触发场景**: 当用户需要开发 Go 后端功能（创建 model、api、router）时自动使用。

**执行流程**:
1. 先读取 docs/BACKEND.md 和 docs/STANDARD.md 了解代码规范
2. 使用 Glob/Grep 查找 app/admin/model/ 和 app/admin/api/ 目录下相似功能的现有代码作为参考
3. 按照项目规范创建/修改文件：
   - Model 文件：app/admin/model/system/sys*.go
   - API 文件：app/admin/api/system/*.go
   - Router 文件：app/admin/router/system/*.go
4. 确保使用项目统一的 R.ReturnSuccess/R.ReturnFailMsg 返回格式
5. 使用 jwt.JWTAuthMiddleware() 作为认证中间件
6. 检查代码是否符合 Go 语言规范和项目风格
7. 提醒用户运行 go build 验证编译

**Why**: 项目已有完整的后端架构和代码规范，使用此技能可以确保新代码与现有代码风格一致，避免重复造轮子。

**How to apply**: 当用户要求创建 API、修改 Handler、添加路由等后端开发任务时自动应用。
