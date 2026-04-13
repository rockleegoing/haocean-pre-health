# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

RuoYi-Go 是一个基于 Go + Gin + Vue 的前后端分离权限管理系统，同时提供原生 Android 版本（UniApp 开发）。

## ⚠️ 重要：开发前必读

**在修改代码或执行任务前，必须先阅读相关文档：**

严格按照流程执行：
  1. 阅读 PRD.md 对应章节
  2. 阅读原型图 + 原型图分析.md
  3. 查阅相关技术文档
  4. 选择适合的工具/Skill
  5. 正式开发

| 任务类型 | 必读文档 |
|----------|----------|
| 修改 Go 后端代码 | [BACKEND.md](docs/BACKEND.md) + [STANDARD.md](docs/STANDARD.md) |
| 修改 PC 前端代码 | [FRONTEND.md](docs/FRONTEND.md) + [STANDARD.md](docs/STANDARD.md) |
| 修改移动端代码 | [MOBILE.md](docs/MOBILE.md) + [STANDARD.md](docs/STANDARD.md) |
| 理解项目结构 | [ARCHITECTURE.md](docs/ARCHITECTURE.md) |
| 部署/运维操作 | [DEPLOY.md](docs/DEPLOY.md) |
| Git 操作 | [GIT.md](docs/GIT.md) |

**原因**：各文档包含具体的代码规范、命名约定、工具函数和最佳实践，不阅读可能导致代码风格不一致或复用现有功能。

## 技术栈

| 层级 | 技术 |
|------|------|
| 后端 | Go 1.24 + Gin + GORM + MySQL/SQLite + Redis + JWT |
| PC 前端 | Vue2 + ElementUI + Vuex + Vue Router |
| 移动端 | UniApp + uview-ui + qiun-data-charts |

## 快速开始

```bash
# 后端
go mod download
go run main.go

# PC 前端
cd frontend/ruoyi-ui
npm install --registry=https://registry.npmmirror.com
npm run dev

# 移动端
cd frontend/RuoYi-Mobile
pnpm install
npm run dev
```

## 测试

```bash
go test ./tests/...
go test -run TestPingRoute ./tests/
```

## 文档索引

| 文档 | 说明 | 适用场景 |
|------|------|----------|
| [ARCHITECTURE.md](docs/ARCHITECTURE.md) | 架构设计、功能模块 | 理解项目结构、模块划分 |
| [BACKEND.md](docs/BACKEND.md) | Go 后端开发 | 开发 API、编写 Handler |
| [FRONTEND.md](docs/FRONTEND.md) | PC 前端开发 | Vue 页面开发 |
| [MOBILE.md](docs/MOBILE.md) | 移动端开发 | UniApp/Android 开发 |
| [DEPLOY.md](docs/DEPLOY.md) | 部署运维 | 生产环境部署 |
| [GIT.md](docs/GIT.md) | Git 版本控制 | 代码提交、推送 |
| [STANDARD.md](docs/STANDARD.md) | 代码规范 | 代码风格检查 |
| [OFFICIAL-DOCS.md](docs/OFFICIAL-DOCS.md) | 官方文档笔记 | 完整官方文档整理 |

## 常见问题

**Node OpenSSL 错误**: `export NODE_OPTIONS=--openssl-legacy-provider`
