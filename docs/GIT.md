# Git 版本控制指南

## 仓库信息

- **远程仓库**: https://github.com/rockleegoing/haocean-pre-health
- **SSH 地址**: git@github.com:rockleegoing/haocean-pre-health.git
- **分支**: main

## SSH 认证配置

### SSH 密钥

文件：`~/.ssh/id_ed25519_github`

### SSH 配置

文件：`~/.ssh/config`

```
Host github.com-rockleegoing
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_github
```

### 添加 SSH Key 到 GitHub

1. 复制公钥：`cat ~/.ssh/id_ed25519_github.pub`
2. 登录 GitHub → Settings → SSH and GPG keys
3. New SSH key → 粘贴公钥 → 保存

### SSH 密钥推送流程（推荐）

每次新终端会话需要重新添加 SSH 密钥到 agent：

```bash
# 1. 添加 SSH 密钥到 agent
ssh-add ~/.ssh/id_ed25519_github

# 2. 验证密钥已添加
ssh-add -l

# 3. 推送代码
git push origin main
```

### 常见错误及解决

**错误**: `Permission denied (publickey).`

**原因**: SSH 密钥未添加到 agent

**解决**:
```bash
ssh-add ~/.ssh/id_ed25519_github
git push origin main
```

**错误**: `Could not read from remote repository.`

**原因**: SSH 密钥不存在或权限错误

**解决**:
```bash
# 检查密钥文件
ls -la ~/.ssh/id_ed25519_github*

# 重新生成密钥（如丢失）
ssh-keygen -t ed25519 -C "your_email@example.com" -f ~/.ssh/id_ed25519_github

# 重新添加到 GitHub
cat ~/.ssh/id_ed25519_github.pub
# 复制输出内容到 GitHub Settings
```

## 常用命令

### 初始化仓库

```bash
git init
git add .
git commit -m "initial commit"
git branch -M main
git remote add origin git@github.com:rockleegoing/haocean-pre-health.git
```

### 推送代码

```bash
# 首次推送 (使用特定 SSH 密钥)
GIT_SSH_COMMAND="ssh -i ~/.ssh/id_ed25519_github" git push -u origin main

# 日常推送 (SSH 已配置时)
git push

# 强制推送 (慎用)
git push -f origin main
```

### 拉取代码

```bash
git pull origin main
```

### 查看状态

```bash
# 查看远程仓库
git remote -v

# 查看工作区状态
git status

# 查看提交历史
git log --oneline
```

### 分支操作

```bash
# 创建新分支
git checkout -b feature/new-feature

# 切换分支
git checkout main

# 合并分支
git merge feature/new-feature

# 删除分支
git branch -d feature/new-feature
```

### 暂存操作

```bash
# 暂存修改
git stash

# 查看暂存列表
git stash list

# 恢复暂存
git stash pop
```

## .gitignore

已忽略的文件：

```
# 配置文件 (含敏感信息)
config.yaml

# 依赖目录
node_modules/

# 系统文件
.DS_Store
*/.DS_Store

# 二进制文件
*.exe
*.test
*.out

# 日志目录
/log/

# IDE
/.idea/
```

## 提交规范

```bash
# 功能开发
git commit -m "feat: 添加用户管理功能"

# Bug 修复
git commit -m "fix: 修复登录验证问题"

# 文档更新
git commit -m "docs: 更新 README"

# 代码重构
git commit -m "refactor: 优化数据库连接池"

# 配置修改
git commit -m "chore: 更新依赖版本"
```

## 注意事项

1. **不要提交敏感信息**：
   - 数据库密码
   - API Key
   - JWT Secret
   - 个人凭证

2. **提交前检查**：
   ```bash
   git status
   git diff
   ```

3. **大文件处理**：
   - 使用 Git LFS 管理大文件
   - 避免提交二进制文件

## 相关文档

- [部署指南](DEPLOY.md) - 生产环境部署
- [架构设计](ARCHITECTURE.md) - 项目结构
