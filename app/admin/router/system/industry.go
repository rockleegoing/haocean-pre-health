package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitIndustry 初始化行业分类路由
func InitIndustry(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/industry/list", system.ListIndustry)
		v.GET("/industry/:id", system.GetIndustry)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/industry", system.AddIndustry)
			auth.PUT("/industry", system.UpdateIndustry)
			auth.DELETE("/industry/:ids", system.DeleteIndustry)
		}
	}
}
