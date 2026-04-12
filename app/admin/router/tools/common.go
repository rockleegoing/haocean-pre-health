package tools

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/admin/api/tools"
	"ruoyi-go/app/core/utils/jwt"
)

func InitCommon(e *gin.Engine) {
	// 公共相关
	v := e.Group("common")
	{
		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.GET("/download", tools.GetDownload)
			auth.POST("/upload", tools.UploadCommon)
			auth.POST("/uploads", tools.UploadCommons)
			auth.GET("/download/resource", tools.UploadRsource)
		}
	}
}
