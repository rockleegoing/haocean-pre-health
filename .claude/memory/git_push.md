---
name: Git 推送方式
description: 记录正确的 git push 命令格式
type: reference
---

推送代码到远程仓库的命令：

```bash
git push origin main
```

**Why:** 项目中可能存在多个分支，明确指定 `origin main` 可以确保推送到正确远程仓库的主分支。

**How to apply:** 每次提交后需要推送时使用此命令。如果遇到权限问题，检查 git remote -v 确认远程仓库地址是否正确。
