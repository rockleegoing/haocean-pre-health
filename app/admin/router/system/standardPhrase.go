package system

import (
	"github.com/gin-gonic/gin"
	api "haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitStandardPhrase 初始化规范用语路由
func InitStandardPhrase(e *gin.Engine) {
	v := e.Group("system")
	{
		// 公开路由（可缓存）
		v.GET("/standard-phrase/supervision-type/list", api.ListSupervisionType)
		v.GET("/standard-phrase/supervision-type/:id", api.GetSupervisionType)

		v.GET("/standard-phrase/category/list", api.ListCategory)
		v.GET("/standard-phrase/category/:id", api.GetCategory)

		v.GET("/standard-phrase/item/list", api.ListItem)
		v.GET("/standard-phrase/item/:id", api.GetItem)

		v.GET("/standard-phrase/content/list", api.ListContent)
		v.GET("/standard-phrase/content/:id", api.GetContent)

		// 搜索和树形结构
		v.GET("/standard-phrase/search", api.SearchStandardPhrase)
		v.GET("/standard-phrase/tree", api.GetFullTree)

		// 需要认证的路由
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			// 监管类型
			auth.POST("/standard-phrase/supervision-type", api.AddSupervisionType)
			auth.PUT("/standard-phrase/supervision-type", api.UpdateSupervisionType)
			auth.DELETE("/standard-phrase/supervision-type/:ids", api.DeleteSupervisionType)

			// 规范类别
			auth.POST("/standard-phrase/category", api.AddCategory)
			auth.PUT("/standard-phrase/category", api.UpdateCategory)
			auth.DELETE("/standard-phrase/category/:ids", api.DeleteCategory)

			// 规范条目
			auth.POST("/standard-phrase/item", api.AddItem)
			auth.PUT("/standard-phrase/item", api.UpdateItem)
			auth.DELETE("/standard-phrase/item/:ids", api.DeleteItem)

			// 规范内容
			auth.POST("/standard-phrase/content", api.AddContent)
			auth.PUT("/standard-phrase/content", api.UpdateContent)
			auth.DELETE("/standard-phrase/content/:ids", api.DeleteContent)
		}
	}
}
