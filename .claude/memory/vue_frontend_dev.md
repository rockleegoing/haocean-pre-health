---
name: Vue 前端开发助手
description: PC 前端开发专用助手，遵循项目 FRONTEND.md 和 STANDARD.md 规范
type: feedback
---

**触发场景**: 当用户需要开发 PC 前端功能（创建 Vue 页面、API 接口）时自动使用。

**执行流程**:
1. 先读取 docs/FRONTEND.md 和 docs/STANDARD.md 了解代码规范
2. 查看 frontend/health-enforcement-ui/src/views/system/ 目录下相似功能的现有页面作为参考
3. 按照项目规范创建文件：
   - 页面：src/views/system/*/index.vue
   - API: src/api/system/*.js
4. 确保使用项目统一的组件：
   - 表格分页：pagination 组件
   - 搜索表单：el-form inline
   - 操作按钮：v-hasPermi 权限指令
   - 消息提示：this.msgSuccess/this.msgError
5. API 调用使用 request 工具，路径以 /system 开头
6. 提醒用户运行 npm run dev 预览效果

**Why**: 项目已有完整的前端架构和组件库，使用此技能可以确保新页面与现有页面风格一致。

**How to apply**: 当用户要求创建 Vue 页面、修改前端组件、添加 API 接口等前端开发任务时自动应用。
