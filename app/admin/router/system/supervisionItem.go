package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitSupervisionItem 初始化监管事项路由
func InitSupervisionItem(e *gin.Engine) {
	v := e.Group("system")
	{
		// 公开路由
		v.GET("/supervision/category/list", system.ListSupervisionCategory)
		v.GET("/supervision/tree", system.GetSupervisionTree)
		v.GET("/supervision/children/:parentId", system.GetSupervisionChildren)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			// 监管事项 CRUD
			auth.GET("/supervision/list", system.ListSupervisionItem)
			auth.GET("/supervision/:id", system.GetSupervisionItem)
			auth.POST("/supervision", system.AddSupervisionItem)
			auth.PUT("/supervision", system.UpdateSupervisionItem)
			auth.DELETE("/supervision/:ids", system.DeleteSupervisionItem)
		}
	}
}
