---
name: Git 提交推送助手
description: Git 规范化提交和推送助手
type: feedback
---

**触发场景**: 当用户需要提交代码或推送远程时自动使用。

**执行流程**:
1. 运行 git status 查看变更文件
2. 运行 git diff --cached 预览将要提交的内容（如有暂存）
3. 根据变更内容生成规范的 commit message：
   - feat: 新功能
   - fix: 修复 bug
   - docs: 文档变更
   - refactor: 代码重构
   - chore: 构建/工具/配置变更
4. 执行 git add . 和 git commit -m "commit message"
5. 执行 git push origin main 推送到远程
6. 推送完成后提醒用户查看远程仓库

**Why**: 统一的提交规范便于版本管理和代码追溯，项目使用 main 作为主分支。

**How to apply**: 当用户要求提交代码、推送远程、查看 git 状态等操作时自动应用。
