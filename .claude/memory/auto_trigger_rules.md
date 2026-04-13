---
name: 自动触发规则
description: 定义技能和行为的自动触发条件和时机
type: feedback
---

**自动触发规则**:

## 1. Go 后端开发场景
**触发关键词**: "后端"、"API"、"接口"、"Handler"、"Controller"、"Go"、"model"、"router"、"路由"、"Gin"、"GORM"
**自动行为**:
1. 读取 docs/BACKEND.md 了解开发规范
2. 读取 docs/STANDARD.md 了解代码风格
3. 使用 Glob 查找 app/admin/model/system/*.go 和 app/admin/api/system/*.go 作为参考
4. 使用 R.ReturnSuccess/R.ReturnFailMsg 作为返回格式
5. 使用 jwt.JWTAuthMiddleware() 作为认证中间件
6. 完成后提醒运行 go build 验证

## 2. Vue 前端开发场景
**触发关键词**: "前端"、"页面"、"Vue"、"ElementUI"、"PC"、"后台"、"管理"
**自动行为**:
1. 读取 docs/FRONTEND.md 了解开发规范
2. 读取 docs/STANDARD.md 了解代码风格
3. 使用 Glob 查找 frontend/health-enforcement-ui/src/views/system/ 下的相似页面
4. 使用 pagination、right-toolbar 等统一组件
5. 添加 v-hasPermi 权限指令
6. 完成后提醒运行 npm run dev 预览

## 3. UniApp 移动端开发场景
**触发关键词**: "移动端"、"UniApp"、"小程序"、"App"、"uview"、"手机"
**自动行为**:
1. 读取 docs/MOBILE.md 了解开发规范
2. 读取 docs/STANDARD.md 了解代码风格
3. 使用 Glob 查找 frontend/Haocean-Mobile/pages/ 下的相似页面
4. 使用 uview-ui 组件库
5. 更新 pages.json 添加路由
6. 完成后提醒运行 npm run dev 预览

## 4. Git 操作场景
**触发关键词**: "提交"、"推送"、"git commit"、"git push"、"push"
**自动行为**:
1. 运行 git status 查看状态
2. 生成约定式 commit message（feat/fix/docs/refactor/chore）
3. 执行 git add . && git commit -m "xxx"
4. 执行 git push origin main
5. 提醒用户查看远程仓库

## 5. 代码审查场景
**触发关键词**: "审查"、"检查"、"review"、"质量"、"规范"、"对不对"、"有问题吗"
**自动行为**:
1. 检查代码是否符合 STANDARD.md 规范
2. 检查是否有安全风险（SQL 注入、XSS）
3. 检查命名是否规范
4. 检查错误处理是否完善
5. 提供改进建议

## 6. 数据库操作场景
**触发关键词**: "数据库"、"表结构"、"迁移"、"schema"、"SQL"、"MySQL"
**自动行为**:
1. 读取 config/sql/law_schema.sql 了解现有表结构
2. 生成迁移 SQL（包括回滚语句）
3. 提醒用户备份数据
4. 更新对应的 Go model 文件

**Why**: 通过自动触发机制，减少用户重复指令，提高开发效率，确保代码规范一致性。

**How to apply**: 当用户消息中包含上述关键词时，自动执行对应的行为和流程。
