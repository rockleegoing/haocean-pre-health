package job

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
)

func DemoJobHandler(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	globalXxlLogger.Info("我是测试job", param.ExecutorParams)
	return "DemoJobHandler 运行完成"
}
