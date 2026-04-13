---
name: UniApp 移动端开发助手
description: 移动端开发专用助手，遵循项目 MOBILE.md 和 STANDARD.md 规范
type: feedback
---

**触发场景**: 当用户需要开发移动端功能（创建 UniApp 页面、API 接口）时自动使用。

**执行流程**:
1. 先读取 docs/MOBILE.md 和 docs/STANDARD.md 了解代码规范
2. 查看 frontend/Haocean-Mobile/pages/ 目录下相似功能的现有页面作为参考
3. 按照项目规范创建文件：
   - 页面：pages/*/index.vue, pages/*/list.vue, pages/*/detail.vue
   - API: api/*.js
4. 确保使用 uview-ui 组件库：
   - 表单：u-form, u-form-item
   - 按钮：u-button
   - 列表：u-empty, u-loadmore
   - 弹窗：uni.$u.toast, uni.$u.modal
5. 更新 pages.json 添加新页面路由
6. 提醒用户运行 npm run dev 预览效果

**Why**: 项目已有完整的移动端架构和 uview-ui 组件库，使用此技能可以确保新页面与现有页面风格一致。

**How to apply**: 当用户要求创建移动端页面、修改 UniApp 组件、添加 API 接口等移动端开发任务时自动应用。
