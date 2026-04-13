package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitTemplate 初始化文书模板路由
func InitTemplate(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/template/list", system.ListTemplate)
		v.GET("/template/:id", system.GetTemplate)
		v.GET("/template/preview/:id", system.PreviewTemplate)
		v.GET("/template/category/list", system.ListTemplateCategory)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/template/upload", system.UploadTemplate)
			auth.PUT("/template", system.UpdateTemplate)
			auth.DELETE("/template/:ids", system.DeleteTemplate)
		}
	}
}
