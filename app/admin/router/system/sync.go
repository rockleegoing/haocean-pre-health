package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitSync 初始化数据同步路由
func InitSync(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/sync/check", system.CheckSync)
		v.GET("/sync/industries", system.SyncIndustries)
		v.GET("/sync/templates", system.SyncTemplates)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/sync/records", system.SyncRecords)
			auth.POST("/sync/subjects", system.SyncSubjects)
			auth.GET("/sync/status", system.GetSyncStatus)
			auth.POST("/sync/retry", system.RetrySync)
			auth.GET("/sync/list", system.ListSync)
		}
	}
}
