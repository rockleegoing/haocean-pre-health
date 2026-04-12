package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"ruoyi-go/app/admin/model/cache"
	"ruoyi-go/app/core/routers"
	"ruoyi-go/app/core/utils"
	"ruoyi-go/app/core/utils/job"
	"ruoyi-go/app/core/utils/shutdown"
	"ruoyi-go/config"
	"ruoyi-go/pkg/logs"
	"ruoyi-go/pkg/scheduler"
	"strconv"

	xxl_job_executor_gin "github.com/gin-middleware/xxl-job-executor"
)

var configFile = flag.String("f", "./config.yaml", "")

// @title Ruoyi-Go 接口文档
// @version v1.0.2
// @description 基于Go，Gin，JWT，vue前后端分离的权限管理系统
// @contact.name Ruoyi-Go
// @contact.url https://gitee.com/OptimisticDevelopers/Ruoyi-Go
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	fmt.Println("hello ruoyi go")
	logs.Poster()
	// 初始化配置文件
	config.InitAppConfig(*configFile)
	// 初始化 定时
	scheduler.InitCron()

	// 初始化路由
	r := routers.Init()
	routers.InitWeb(r)

	// xxl_job
	if config.XxlJob.Enabled {
		cron := job.InitXxlJobCron()
		xxl_job_executor_gin.XxlJobMux(r, cron)
	}
	// 初始化缓存
	cache.InitCache()
	//打开浏览器
	if runtime.GOOS == "windows" {
		utils.OpenWin("http://127.0.0.1:" + strconv.Itoa(config.Server.Port))
	}

	if runtime.GOOS == "darwin" {
		utils.OpenMac("http://127.0.0.1:" + strconv.Itoa(config.Server.Port))
	}

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Server.Port),
		Handler: r,
	}

	go func() {
		// 启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动失败，listen: %s\n", err)
		}
	}()

	// 优雅关闭
	shutdown.GracefullyShutdown(srv)
}
