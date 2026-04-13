package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitDevice 初始化设备管理路由
func InitDevice(e *gin.Engine) {
	v := e.Group("system")
	{
		v.POST("/device/activate", system.ActivateDevice)
		v.GET("/device/info", system.GetDeviceInfo)
		v.POST("/device/login", system.UpdateLastLogin)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/device/list", system.ListDevice)
			auth.PUT("/device", system.UpdateDevice)
			auth.DELETE("/device/:ids", system.DeleteDevice)
			auth.PUT("/device/disable", system.DisableDevice)
		}
	}
}
