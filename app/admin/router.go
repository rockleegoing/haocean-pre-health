package admin

import (
	api "haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/admin/router/system"
	"haocean/health-enforcement/app/core/utils/jwt"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	handlerFunc := cache.CacheByRequestURI(memoryStore, 2*time.Second)

	// 根路径 API 路由（用于向后兼容）
	e.GET("/index", api.IndexHandler)
	e.GET("/captchaImage", api.CaptchaImageHandler)
	// 登录
	e.POST("/login", api.LoginHandler)
	// 退出
	e.POST("/logout", api.LogoutHandler)
	v1 := e.Group("/")
	{
		auth := v1.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("getInfo", handlerFunc, api.GetInfoHandler)
			/*获取用户授权菜单*/
			auth.GET("getRouters", handlerFunc, api.GetRoutersHandler)
		}
	}
	/*system*/
	system.InitProfile(e)
	system.InitDict(e)
	system.InitUser(e)
	system.InitMenu(e)
	system.InitPost(e)
	system.InitNotice(e)
	system.InitRole(e)
	system.InitConfig(e)
	system.InitDept(e)

	/*business 业务路由 - 卫生执法系统*/
	system.InitIndustry(e)          // 行业分类
	system.InitSubject(e)           // 监管单位
	system.InitOfficial(e)          // 执法人员
	system.InitDevice(e)            // 设备管理
	system.InitActivateCode(e)      // 激活码
	system.InitTemplate(e)          // 文书模板
	system.InitRecord(e)            // 执法记录
	system.InitSync(e)              // 数据同步
	system.InitSupervisionItem(e)   // 监管事项
	system.InitStandardPhrase(e)    // 规范用语
	system.InitRegulation(e)        // 法律法规库

}
