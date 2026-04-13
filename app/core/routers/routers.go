package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"haocean/health-enforcement/app/admin"
	"haocean/health-enforcement/app/core/utils"
	error2 "haocean/health-enforcement/app/core/utils/error"
	"haocean/health-enforcement/app/html"
	"haocean/health-enforcement/config"
	_ "haocean/health-enforcement/docs"
	"haocean/health-enforcement/pkg/logs"
	"strings"
)

// Init 初始化
func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 如果 URL 的路径是固定的，那么重定向到配置的固定路径
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

	// 加载多个 APP 的路由配置
	html.Routers(r)
	admin.Routers(r)

	return r
}

// 判断是否为开发模式
func isDevelopment() bool {
	return config.Server.RunMode == "debug"
}

// 创建反向代理
func createProxy(target string) gin.HandlerFunc {
	remote, err := url.Parse(target)
	if err != nil {
		log.Printf("[ERROR] 反向代理配置错误：%v", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	return func(c *gin.Context) {
		c.Request.Host = remote.Host
		c.Request.Header.Set("X-Forwarded-For", c.ClientIP())
		c.Request.Header.Set("X-Forwarded-Host", c.Request.Host)
		c.Request.Header.Set("X-Forwarded-Proto", "http")
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func InitWeb(r *gin.Engine) {
	r.HTMLRender = createRender()

	r.Static("/profile", "./static/images")

	if isDevelopment() {
		// === 开发环境：代理到前端开发服务器 ===

		// PC 前端代理 (Vue dev server: 1024)
		r.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path

			// favicon.ico 代理到 PC 前端
			if strings.HasPrefix(path, "/favicon.ico") {
				createProxy("http://localhost:1024")(c)
				return
			}

			// PC 前端及其静态资源
			if strings.HasPrefix(path, "/admin") || strings.HasPrefix(path, "/static") {
				createProxy("http://localhost:1024")(c)
				return
			}

			// API 404
			if strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/swagger") {
				c.JSON(404, gin.H{"error": "Not found"})
				return
			}

			// 默认返回 PC 前端 (SPA fallback)
			createProxy("http://localhost:1024")(c)
		})

	} else {
		// === 生产环境：使用静态文件 ===

		r.Static("/admin", "./web/admin")
		r.Static("/static", "./web/mobile/static")

		// 404 NotFound (SPA fallback)
		r.NoRoute(func(c *gin.Context) {
			accept := c.Request.Header.Get("Accept")
			flag := strings.Contains(accept, "text/html")
			if flag {
				content, err := ioutil.ReadFile("web/admin/index.html")
				if err != nil {
					c.Writer.WriteHeader(404)
					c.Writer.WriteString("Not Found")
					return
				}
				c.Writer.WriteHeader(200)
				c.Writer.Header().Add("Accept", "text/html")
				c.Writer.Write(content)
				c.Writer.Flush()
			} else {
				c.JSON(404, gin.H{"error": "Not Found"})
			}
		})
	}

	if config.Server.EnabledSwagger {
		r.GET("/swagger/*any", func(context *gin.Context) {
			ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(context)
		})
	}
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
