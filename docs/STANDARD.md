# 代码规范总则

## EditorConfig

文件：`.editorconfig`

```ini
root = true

[*]
charset = utf-8
indent_style = space
indent_size = 2
end_of_line = lf
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
insert_final_newline = false
trim_trailing_whitespace = false
```

## 命名规范

### Go 后端

| 类型 | 规范 | 示例 |
|------|------|------|
| 模型类 | 大驼峰 + `Sys` 前缀 | `SysUser`, `SysRole` |
| 参数类 | 模型名 + `Parm` | `SysUserParm` |
| Handler | 功能名 + `Handler` | `LoginHandler` |
| 路由初始化 | `Init` + 模块名 | `InitUser` |
| 文件命名 | 小驼峰 | `user.go`, `dictData.go` |

### 前端

| 类型 | 规范 | 示例 |
|------|------|------|
| Vue 组件 | PascalCase | `SysUser.vue` |
| 文件命名 | kebab-case | `user-list.vue` |
| 变量 | camelCase | `userData`, `listData` |
| 常量 | UPPER_SNAKE | `ACCESS_TOKEN` |
| 类名 | PascalCase | `UserService` |

## API 接口规范

### Go 后端

```go
// 统一响应格式
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}

// 成功
R.ReturnSuccess(data)
R.ReturnSuccessMsg("操作成功")

// 失败
R.ReturnFailMsg("错误信息")
```

### 前端

```javascript
// API 模块划分
api/system.js      // 系统管理
api/monitor.js     // 监控管理
api/tool.js        // 工具

// 请求封装
import request from '@/utils/request'

// 返回解构
const { code, msg, data } = response
```

## 样式规范

### PC 端

```scss
// 使用 ElementUI 主题变量
@import '~element-ui/packages/theme-chalk/src/index';

// BEM 命名
.block {}
.block__element {}
.block--modifier {}
```

### 移动端

```scss
// 使用 uview-ui 主题变量
@import "@/uni_modules/uview-ui/theme.scss";

// 响应式单位
.container {
  padding: 32rpx;  // 移动端用 rpx
  width: 750rpx;
}
```

## 注释规范

### Go

```go
// 函数注释：说明功能
// LoginHandler 用户登录处理
func LoginHandler(context *gin.Context) {
    // 参数绑定
    var param system.LoginParam
    if err := context.ShouldBind(&param); err != nil {
        return
    }
}
```

### 前端

```javascript
/**
 * 用户登录
 * @param {Object} params - 登录参数
 * @param {string} params.username - 用户名
 * @param {string} params.password - 密码
 * @returns {Promise}
 */
export function login(params) {
  return request.post('/login', params)
}
```

## 代码组织

### Go 包导入顺序

```go
import (
    // 1. 标准库
    "fmt"
    "net/http"
    "time"
    
    // 2. 第三方库
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    
    // 3. 项目内部包
    "haocean-health/app/admin/model/system"
    "haocean-health/pkg/db"
)
```

### 前端组件结构

```vue
<template>
  <!-- 模板内容 -->
</template>

<script>
// 1. 导入依赖
import { getUser } from '@/api/user'

// 2. 组件定义
export default {
  name: 'SysUser',
  components: {},
  props: {},
  data() {},
  computed: {},
  watch: {},
  created() {},
  methods: {}
}
</script>

<style scoped lang="scss">
// 样式
</style>
```

## 错误处理

### Go

```go
// 错误返回
if err != nil {
    context.JSON(http.StatusOK, R.ReturnFailMsg(err.Error()))
    return
}

// 恐慌恢复
defer func() {
    if r := recover(); r != nil {
        log.Printf("panic: %v", r)
    }
}()
```

### 前端

```javascript
// try-catch
try {
  await getUserList()
} catch (error) {
  console.error('获取用户列表失败:', error)
  this.$message.error('加载失败')
}

// Promise catch
getUserList().catch(error => {
  console.error(error)
})
```

## 安全规范

1. **密码加密**: 使用 bcrypt
2. **Token 验证**: 所有 API 需验证 JWT
3. **SQL 注入**: 使用 GORM 参数化查询
4. **XSS 防护**: 前端转义用户输入
5. **CORS 配置**: 限制允许的来源

## 性能优化

1. **数据库**: 使用连接池，配置合理参数
2. **缓存**: 热点数据使用 Redis
3. **前端**: 路由懒加载，组件按需引入
4. **静态资源**: 启用 Gzip 压缩

## 相关文档

- [后端开发](BACKEND.md) - Go 开发指南
- [前端开发](FRONTEND.md) - PC 端开发
- [移动端开发](MOBILE.md) - UniApp 开发
- [架构设计](ARCHITECTURE.md) - 整体架构
