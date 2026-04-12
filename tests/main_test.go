package tests

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"ruoyi-go/app/core/routers"
	"ruoyi-go/config"
	"testing"
)

var configFile = flag.String("f", "./../config.yaml", "")

func TestPingRoute(t *testing.T) {
	config.InitAppConfig(*configFile)
	// 获取配置好的路由器实例
	router := routers.Init()

	// 创建一个响应记录器，用于捕获接口返回的响应
	w := httptest.NewRecorder()
	// 创建一个 GET 请求，访问 /ping 路径，无请求体
	req, _ := http.NewRequest("GET", "/ping", nil)
	// 让路由器处理这个测试请求
	router.ServeHTTP(w, req)

	// 断言响应状态码是 200
	assert.Equal(t, 200, w.Code)
	// 断言响应体符合预期
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
