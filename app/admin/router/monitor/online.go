package monitor

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/monitor"
	"haocean/health-enforcement/app/core/utils/jwt"
)

func InitOnLine(e *gin.Engine) {
	// 在线用户相关
	v := e.Group("monitor/online")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("list", monitor.ListOnLine)
			auth.DELETE("/:tokenId", monitor.DetectOnLine)
		}
	}
}
