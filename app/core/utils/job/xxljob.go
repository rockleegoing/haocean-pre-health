package job

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"log"
	"ruoyi-go/config"
	"strconv"
)

var globalXxlLogger = &logger{}

/*
*	https://github.com/gin-middleware/xxl-job-executor（最新版本）
*	https://github.com/PGshen/go-xxl-executor(之前版本)
 */
func InitXxlJobCron() xxl.Executor {

	//初始化执行器
	exec := xxl.NewExecutor(
		xxl.ServerAddr(config.XxlJob.AdminAddress),
		xxl.AccessToken(config.XxlJob.AccessToken),         //请求令牌(默认为空)
		xxl.ExecutorIp(config.XxlJob.Ip),                   //可自动获取
		xxl.ExecutorPort(strconv.Itoa(config.XxlJob.Port)), //默认9999（此处要与gin服务启动port必需一至）
		xxl.RegistryKey(config.XxlJob.AppName),             //执行器名称
		xxl.SetLogger(globalXxlLogger),                     //自定义日志
	)

	if config.XxlJob.Enabled {
		exec.Init()

		exec.Use(customMiddleware)
		//设置日志查看handler
		exec.LogHandler(customLogHandle)

		//注册任务handler（测试）这里是JobHandler*的名字
		exec.RegTask("task.test", task.Test)
		exec.RegTask("task.test2", task.Test2)
		exec.RegTask("task.panic", task.Panic)
		exec.RegTask("demoJobHandler", DemoJobHandler)
	}

	return exec
}

// 自定义日志处理器
func customLogHandle(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "Ruoyi-Go-xxljob日志handler",
		IsEnd:       true,
	}}
}

// xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("Ruoyi-Go-xxljob日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("Ruoyi-Go-xxljob日志 - "+format, a...))
}

// 自定义中间件
func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		log.Println("I am a middleware start")
		res := tf(cxt, param)
		log.Println("I am a middleware end")
		return res
	}
}
