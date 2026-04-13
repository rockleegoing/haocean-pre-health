# Go 后端开发指南

## 路由系统

### Gin 路由原理

Gin 框架的路由基于 **Radix Tree（基数树）** 实现高效匹配：
- 每个 HTTP 方法（GET/POST 等）维护单独的树
- 子节点按优先级排序
- 低内存占用，高并发处理能力

### 路由位置

全部路由配置在 `app/admin/router.go` 文件中。

### 基础路由示例

```go
// 需要登录才能访问的路由
v1 := e.Group("/")
{
    auth := v1.Group("")
    auth.Use(jwt.JWTAuthMiddleware())
    {
        auth.GET("getInfo", handlerFunc, api.GetInfoHandler)
        auth.GET("getRouters", handlerFunc, api.GetRoutersHandler)
    }
}
```

### 添加新模块路由

```go
// 1. 在 app/admin/router/ 下创建新文件，如 demo.go
func InitDemo(e *gin.Engine) {
    // 不需登录的路由
    e.GET("/index", api.IndexHandler)
    
    // 需要登录的路由
    v := e.Group("demo")
    {
        auth := v.Group("")
        auth.Use(jwt.JWTAuthMiddleware())
        {
            auth.GET("/hello", demo.Hello)
        }
    }
}

// 2. 在 app/admin/router.go 中注册
demo.InitDemo(e)
```

## 命名约定

| 类型 | 规范 | 示例 |
|------|------|------|
| 模型类 | 大驼峰 + `Sys` 前缀 | `SysUser`, `SysRole`, `SysMenu` |
| 参数类 | 模型名 + `Parm` | `SysUserParm`, `SysUserExcel` |
| Handler | 功能名 + `Handler` | `LoginHandler`, `ListUser` |
| 路由初始化 | `Init` + 模块名 | `InitUser`, `InitMenu` |

## API 响应格式

文件：`app/core/utils/R/R.go`

```go
// 成功返回数据
R.ReturnSuccess(data)

// 成功返回消息
R.ReturnSuccessMsg("操作成功")

// 失败返回消息
R.ReturnFailMsg("错误信息")
```

**响应结构**:
```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}
```

## Handler 编写规范

```go
func XxxHandler(context *gin.Context) {
    // 1. 参数绑定
    var param system.XxxParam
    if err := context.ShouldBind(&param); err != nil {
        context.JSON(http.StatusOK, R.ReturnFailMsg("参数不能为空"))
        return
    }
    
    // 2. 业务逻辑（封装到 model 或 service）
    result := system.DoSomething(param)
    
    // 3. 统一返回
    context.JSON(http.StatusOK, R.ReturnSuccess(result))
}
```

**要点**:
1. 参数绑定使用 `context.ShouldBind(&param)`
2. 参数校验失败返回 `gin.H{"msg": "...", "code": ...}`
3. 业务逻辑封装到 model 或 service 层
4. 统一使用 `R.Result` 包装返回值

## GORM 模型定义

```go
type SysUser struct {
    UserId      int       `json:"userId" gorm:"column:user_id;primaryKey"`
    UserName    string    `json:"userName" gorm:"user_name"`
    Password    string    `json:"password" gorm:"password"`
    Status      string    `json:"status" gorm:"status"`
    CreateTime  time.Time `json:"createTime" gorm:"autoCreateTime"`
    UpdateTime  time.Time `json:"updateTime" gorm:"autoUpdateTime"`
}
```

**常用 Tag**:
- `gorm:"primaryKey"` - 主键
- `gorm:"column:field_name"` - 列名
- `gorm:"autoCreateTime"` - 自动填充创建时间
- `gorm:"autoUpdateTime"` - 自动填充更新时间

## 路由规范

```go
func InitUser(e *gin.Engine) {
    v := e.Group("system")  // 路由前缀 /system
    {
        auth := v.Group("")
        auth.Use(jwt.JWTAuthMiddleware())  // JWT 认证
        {
            auth.GET("/user/list", ListUser)
            auth.POST("/user", SaveUser)
            auth.PUT("/user", UploadUser)
            auth.DELETE("/user/:userIds", DeleteUserById)
        }
    }
}
```

## 数据库操作

```go
import "haocean-health/pkg/db"

// 查询
var list []SysUser
db.Exe().Model(SysUser{}).Find(&list)

// 分页查询
var total int64
db.Exe().Model(SysUser{}).Count(&total)
db.Exe().Limit(pageSize).Offset(offset).Find(&list)

// 单条查询
var user SysUser
db.Exe().First(&user, userId)

// 创建
db.Exe().Create(&user)

// 更新
db.Exe().Save(&user)

// 删除
db.Exe().Delete(&user, userId)
```

**主从切换**:
```go
db.Exe()            // 主库
db.Exe(true)        // 从库
```

## 缓存使用

```go
import "haocean-health/pkg/cache/redisCache"

// 设置缓存
redisCache.NewRedisCache().Set(key, value, expireSeconds)

// 获取缓存
value := redisCache.NewRedisCache().Get(key)

// 删除缓存
redisCache.NewRedisCache().Delete(key)
```

## 日志使用

```go
import "github.com/sirupsen/logrus"

// Info
logrus.Info("操作成功")

// Error
logrus.Error("操作失败", err)

// Warn
logrus.Warn("警告信息")
```

## 配置文件

详见：[DEPLOY.md](DEPLOY.md#配置文件示例)

## 定时任务

### 添加定时任务

```go
import "haocean-health/pkg/scheduler"

// 每秒执行
scheduler.AddCronFunc("*/1 * * * * *", func() {
    // 任务逻辑
})
```

### 任务方法定义

```go
// 无参数方法
func NoParamsMethod() {
    // 任务逻辑
}

// 单参数方法
func ParamsMethod(param string) {
    // 任务逻辑
}
```

## 工具类

| 文件 | 功能 |
|------|------|
| `app/core/utils/pwd.go` | 密码加密 (bcrypt) |
| `app/core/utils/jwt/jwt.go` | JWT 生成和验证 |
| `app/core/utils/captchaImageUtils.go` | 验证码 |
| `app/core/utils/rateLimiter.go` | 限流中间件 |
| `app/core/utils/repeatSubmit.go` | 防重复提交中间件 |

## 相关文档

- [架构设计](ARCHITECTURE.md) - 整体架构
- [代码规范](STANDARD.md) - Go 代码规范
- [部署指南](DEPLOY.md) - 生产部署
