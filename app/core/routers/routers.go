package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"net/http"
	"ruoyi-go/app/admin"
	"ruoyi-go/app/core/utils"
	error2 "ruoyi-go/app/core/utils/error"
	"ruoyi-go/app/html"
	"ruoyi-go/config"
	_ "ruoyi-go/docs"
	"ruoyi-go/pkg/logs"
	"strings"
)

// Init 初始化
func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 如果URL的路径是固定的，那么重定向到配置的固定路径
	r.RedirectFixedPath = true

	r.Use(logs.Logger())
	r.Use(gin.Recovery())
	/*自定义错误*/
	r.Use(error2.Recover)

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.Use(utils.Core())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 加载多个APP的路由配置
	html.Routers(r)
	admin.Routers(r)

	return r
}

func InitWeb(r *gin.Engine) {
	r.HTMLRender = createRender()

	r.Static("/profile", "./static/images")
	r.Static("/admin", "./web/admin")
	r.Static("/static", "./web/mobile/static")
	r.Static("/favicon.ico/", "./web/admin/favicon.ico")

	if config.Server.EnabledSwagger {
		r.GET("/swagger/*any", func(context *gin.Context) {
			ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(context)
		})
	}

	// 关键点【解决页面刷新404的问题】
	// 404 NotFound
	r.NoRoute(func(context *gin.Context) {
		accept := context.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("web/admin/index.html")
			if (err) != nil {
				context.Writer.WriteHeader(404)
				context.Writer.WriteString("Not Found")
				return
			}
			context.Writer.WriteHeader(200)
			context.Writer.Header().Add("Accept", "text/html")
			context.Writer.Write((content))
			context.Writer.Flush()
		}
	})
}

// 不同模板设置
func createRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/index.html")
	p.AddFromFiles("mobile", "web/mobile/index.html")
	p.AddFromFiles("mobile_old", "web/mobile_old/index.html")
	p.AddFromFiles("protocol", "web/template/protocol.tmpl")
	return p
}
