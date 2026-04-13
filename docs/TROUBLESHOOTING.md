# 常见问题排查指南

## 移动端 API 返回 401 认证失败

### 问题现象
- 移动端访问 `http://localhost:9001/` 后正常登录
- 登录后访问 API（如 `/getInfo`）返回 `{"code":401,"msg":"认证失败"}`
- 请求头中 `Authorization: Bearer <token>` 已正确发送

### 问题原因
后端 JWT 认证机制需要同时满足以下条件：
1. Token 签名验证通过
2. Redis 中存在对应的 `go_login_tokens:{uuid}` 缓存

如果前端存储的 token 与 Redis 中的缓存不匹配（uuid 不同），则认证失败。

### 排查步骤

#### 1. 检查后端 Redis 缓存
```bash
# 检查 Redis 是否运行
redis-cli ping
# 预期输出：PONG

# 检查登录 token 缓存
redis-cli keys "go_login_tokens:*"
# 预期输出：go_login_tokens:{uuid}

# 查看具体 token 数据
redis-cli get "go_login_tokens:{uuid}"
# 预期输出：包含 userId, userName, deptId 等信息的 JSON
```

#### 2. 检查前端存储的 token
在浏览器控制台执行：
```javascript
// 检查 localStorage 中的 token
console.log('Token:', uni.getStorageSync('AccessToken'))

// 解析 token 获取 uuid（需要 jwt-decode 库）
// 或者查看 Network 面板中请求的 Authorization header
```

#### 3. 对比 uuid
- 前端 token 中的 uuid（JWT payload 中的 `uuid` 字段）
- Redis key 中的 uuid（`go_login_tokens:{uuid}`）

两者必须一致才能通过认证。

### 解决方案

#### 方案一：重新登录（推荐）
1. 清除浏览器缓存或使用无痕模式
2. 访问 `http://localhost:9001`
3. 输入账号密码登录
4. 登录成功后会自动保存新的 token

#### 方案二：后端清除旧 token
```bash
# 清除所有登录 token
redis-cli keys "go_login_tokens:*" | xargs redis-cli del

# 或者直接重启后端服务，Redis 缓存会过期后自动清除
```

### 根本原因分析

1. **Token 生成机制**：
   - 每次登录成功后，后端生成新的 UUID
   - 使用 UUID 作为 Redis key 存储用户信息
   - JWT token 的 payload 中包含该 UUID

2. **认证流程**：
   ```
   客户端请求 → 提取 Authorization header → 解析 JWT → 获取 uuid
              → 查询 Redis(go_login_tokens:{uuid}) → 验证成功 → 设置 userId 到 context
   ```

3. **常见问题场景**：
   - 多次登录后前端仍使用旧 token
   - 后端重启后 Redis 数据丢失
   - 跨浏览器/设备使用同一账号（每个设备 token 不同）
   - Token 过期后未重新登录

### 预防措施

1. **前端增加 token 过期处理**：
   ```javascript
   // http.interceptor.js 中增强 401 处理
   case 401: {
     // 清除本地 token
     storage.remove(ACCESS_TOKEN)
     // 跳转到登录页
     uni.reLaunch({ url: '/pages/login/index' })
     return Promise.reject(data)
   }
   ```

2. **登录成功后强制刷新页面**：
   ```javascript
   // pages/login/index.vue
   .then(result => {
     // 登录成功，使用 reLaunch 确保页面完全刷新
     uni.reLaunch({
       url: '/pages/index/index'
     })
   })
   ```

3. **开发环境调试技巧**：
   - 在登录成功后添加 `console.log('登录成功，token:', result.token)`
   - 在 API 请求前添加 `console.log('当前 token:', storage.get(ACCESS_TOKEN))`
   - 使用 Network 面板查看实际发送的 Authorization header

---

## 开发环境代理配置

### PC 前端
- 开发服务器端口：`1024`
- 后端代理地址：`http://localhost:8080/admin`
- API 代理：`/dev-api` → `http://localhost:8080`

### 移动端（HBuilderX）
- 开发服务器端口：`9001`
- 直接访问：`http://localhost:9001`
- API 代理：`/prod-api` → `http://localhost:8080`（通过 vue.config.js）

### 后端
- 服务端口：`8080`
- JWT Token Header：`Authorization`
- Token 前缀：`Bearer `
- Redis Key 前缀：`go_login_tokens:`

---

## 快速验证命令

```bash
# 1. 测试后端服务
curl http://localhost:8080/ping

# 2. 测试登录 API
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123","code":"","uuid":""}'

# 3. 使用返回的 token 测试 getInfo
curl http://localhost:8080/getInfo \
  -H "Authorization: Bearer <token>"

# 4. 检查 Redis 缓存
redis-cli keys "go_login_tokens:*"
redis-cli get "go_login_tokens:<uuid>"
```
