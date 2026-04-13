package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/api/system"
	"haocean/health-enforcement/app/core/utils/jwt"
)

// InitRecord 初始化执法记录路由
func InitRecord(e *gin.Engine) {
	v := e.Group("system")
	{
		v.GET("/record/list", system.ListRecord)
		v.GET("/record/:id", system.GetRecord)

		auth := v.Group("")
		auth.Use(jwt.JWTAuthMiddleware())
		{
			auth.POST("/record", system.AddRecord)
			auth.PUT("/record", system.UpdateRecord)
			auth.DELETE("/record/:ids", system.DeleteRecord)
			auth.PUT("/record/submit/:id", system.SubmitRecord)
			auth.POST("/evidence/upload", system.UploadEvidence)
			auth.DELETE("/evidence/:id", system.DeleteEvidence)
		}
	}
}
