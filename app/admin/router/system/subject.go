package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitSubject 初始化监管单位路由
func InitSubject(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/subject/list", system.ListSubject)
		v.GET("/subject/:id", system.GetSubject)
		v.GET("/subject/search", system.SearchSubject)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/subject", system.AddSubject)
			auth.PUT("/subject", system.UpdateSubject)
			auth.DELETE("/subject/:ids", system.DeleteSubject)
		}
	}
}
