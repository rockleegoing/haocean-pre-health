package tests

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"haocean/health-enforcement/app/core/routers"
	"haocean/health-enforcement/config"
	systemModel "haocean/health-enforcement/app/admin/model/system"
)

var configFile = flag.String("f", "./../config.yaml", "")

// TestIndustryAPI 测试行业分类 API
func TestIndustryAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListIndustry", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/industry/list", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("AddIndustry", func(t *testing.T) {
		industry := systemModel.SysIndustry{
			IndustryCode: "TEST001",
			IndustryName: "测试行业",
			ParentId:     0,
			Level:        1,
			IsEnabled:    1,
			OrderNum:     1,
		}
		body, _ := json.Marshal(industry)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/system/industry", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestSubjectAPI 测试监管单位 API
func TestSubjectAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListSubject", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/subject/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("SearchSubject", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/subject/search?keyword=测试", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestOfficialAPI 测试执法人员 API
func TestOfficialAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListOfficial", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/official/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestDeviceAPI 测试设备管理 API
func TestDeviceAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListDevice", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/device/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestActivateCodeAPI 测试激活码 API
func TestActivateCodeAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListActivateCode", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/activate-code/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestTemplateAPI 测试文书模板 API
func TestTemplateAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListTemplate", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/template/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestRecordAPI 测试执法记录 API
func TestRecordAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("ListRecord", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/record/list?pageNum=1&pageSize=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestSyncAPI 测试数据同步 API
func TestSyncAPI(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	t.Run("CheckSync", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/sync/check?deviceId=1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

	t.Run("SyncIndustries", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/system/sync/industries", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})
}

// TestPingRoute 测试基础路由
func TestPingRoute(t *testing.T) {
	config.InitAppConfig(*configFile)
	router := routers.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
