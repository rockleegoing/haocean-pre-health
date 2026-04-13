package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitRegulation 初始化法律法规库路由
func InitRegulation(e *gin.Engine) {
	// 移动端公开接口（不需要 JWT）
	v := e.Group("api")
	{
		// 法律法规首页数据
		v.GET("/regulation/home", system.GetRegulationHome)
		// 法律法规书本列表
		v.GET("/regulation/book-list", system.GetRegulationBookList)
		// 法律法规书本详情
		v.GET("/regulation/book-detail/:id", system.GetRegulationBookDetail)
		// 章节内容
		v.GET("/regulation/chapter-content/:chapterId", system.GetRegulationChapterContent)
		// 搜索法律法规
		v.GET("/regulation/search", system.SearchRegulation)
		// 法律类型列表
		v.GET("/regulation/legal-type", system.GetLegalTypeList)
		// 监管类型列表
		v.GET("/regulation/supervision-type", system.GetSupervisionTypeList)
		// 定性依据列表
		v.GET("/regulation/basis-list", system.GetBasisList)
		// 定性依据详情
		v.GET("/regulation/basis-detail/:id", system.GetBasisDetail)
	}

	// 管理端接口（需要 JWT 认证）
	admin := e.Group("api/admin")
	{
		auth := admin.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			// 法律法规管理
			auth.GET("/regulation/list", system.ListRegulation)
			auth.GET("/regulation/:id", system.GetRegulation)
			auth.POST("/regulation", system.AddRegulation)
			auth.PUT("/regulation", system.UpdateRegulation)
			auth.DELETE("/regulation/:ids", system.DeleteRegulation)
		}
	}
}
