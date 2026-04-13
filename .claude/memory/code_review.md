---
name: 代码审查助手
description: 代码质量和规范性审查助手
type: feedback
---

**触发场景**: 当用户要求审查代码、检查代码质量时自动使用。

**审查要点**:

1. **Go 后端**:
   - 是否使用项目统一的 R.ReturnSuccess/R.ReturnFailMsg 返回格式
   - 是否使用 jwt.JWTAuthMiddleware() 作为认证中间件
   - 命名是否符合 Go 规范（导出类型/函数首字母大写）
   - 是否有适当的错误处理
   - 是否有 SQL 注入风险（必须使用 GORM 参数化查询）

2. **Vue 前端**:
   - 是否使用项目统一的组件（pagination, right-toolbar 等）
   - 是否使用 this.msgSuccess/this.msgError 进行消息提示
   - 按钮是否添加 v-hasPermi 权限指令
   - API 路径是否正确（以/system 开头）
   - 是否有 XSS 风险（用户输入需转义）

3. **UniApp 移动端**:
   - 是否使用 uview-ui 组件库
   - 是否使用 uni.$u.toast/uni.$u.modal 进行消息提示
   - 页面是否在 pages.json 中注册
   - API 调用是否使用 request 工具

**Why**: 确保代码符合项目规范，避免安全漏洞和风格不一致问题。

**How to apply**: 当用户要求审查代码、检查代码质量、代码提交前检查等场景时自动应用。
