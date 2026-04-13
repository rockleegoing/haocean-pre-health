package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitOfficial 初始化执法人员路由
func InitOfficial(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/official/list", system.ListOfficial)
		v.GET("/official/:id", system.GetOfficial)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/official", system.AddOfficial)
			auth.PUT("/official", system.UpdateOfficial)
			auth.DELETE("/official/:ids", system.DeleteOfficial)
			auth.POST("/official/bind-device", system.BindDevice)
		}
	}
}
