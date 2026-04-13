#!/bin/bash

echo "=========================================="
echo "     启动开发环境"
echo "=========================================="
echo ""

# 设置开发模式
export GIN_MODE=debug

# 1. 启动后端
echo "[1/3] 启动后端 (Go + Gin)..."
cd /Users/arthur/Documents/workspace/haocean/project_v3/Ruoyi-Go
go run main.go &
BACKEND_PID=$!
echo "        后端进程 ID: $BACKEND_PID"

# 2. 启动 PC 前端
echo "[2/2] 启动 PC 前端 (Vue2 + ElementUI) - 端口 1024..."
cd /Users/arthur/Documents/workspace/haocean/project_v3/Ruoyi-Go/frontend/health-enforcement-ui
npm run dev &
FRONTEND_PID=$!
echo "        PC 前端进程 ID: $FRONTEND_PID"

echo ""
echo "=========================================="
echo "     等待服务启动..."
echo "=========================================="

# 等待后端服务就绪
echo "等待后端服务..."
for i in {1..30}; do
    if curl -s http://localhost:8080/ping | grep -q "pong"; then
        echo "✅ 后端服务已就绪"
        break
    fi
    sleep 1
done

# 等待前端服务就绪
echo "等待 PC 前端服务..."
for i in {1..30}; do
    if curl -s http://localhost:1024/ | grep -q "<title>"; then
        echo "✅ PC 前端服务已就绪"
        break
    fi
    sleep 1
done

echo ""
echo "=========================================="
echo "     服务已启动"
echo "=========================================="
echo ""
echo "  PC 前端：   http://localhost:8080/admin"
echo "  后端 API:   http://localhost:8080/api"
echo "  Swagger:    http://localhost:8080/swagger"
echo "  移动端：    HBuilderX 运行 (http://localhost:9001)"
echo ""
echo "  按 Ctrl+C 停止所有服务"
echo ""

# 自动打开浏览器
sleep 2
echo "正在打开浏览器..."
open http://localhost:8080/admin

# 等待所有进程
wait
