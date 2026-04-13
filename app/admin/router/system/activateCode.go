package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitActivateCode 初始化激活码管理路由
func InitActivateCode(e *gin.Engine) {
	v := e.Group("system")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/activate-code/list", system.ListActivateCode)
			auth.GET("/activate-code/:id", system.GetActivateCode)
			auth.POST("/activate-code/generate", system.GenerateActivateCode)
			auth.DELETE("/activate-code/:ids", system.DeleteActivateCode)
			auth.PUT("/activate-code/disable/:id", system.DisableActivateCode)
		}
	}
}
