package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ActivateDevice 设备激活
func ActivateDevice(c *gin.Context) {
	var param struct {
		ActivateCode string `json:"activateCode"`
		DeviceModel  string `json:"deviceModel"`
		OsType       string `json:"osType"`
		OsVersion    string `json:"osVersion"`
		AppVersion   string `json:"appVersion"`
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if param.ActivateCode == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码不能为空"))
		return
	}

	// 验证激活码
	code := system.FindActivateCodeByCode(param.ActivateCode)
	if code.CodeId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码无效"))
		return
	}

	if code.Status != 1 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码已使用或已禁用"))
		return
	}

	if code.ExpireTime.Before(time.Now()) {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码已过期"))
		return
	}

	// 创建设备
	device := system.SysDevice{
		OfficialId: code.OfficialId,
		DeviceModel: param.DeviceModel,
		OsType:     param.OsType,
		OsVersion:  param.OsVersion,
		AppVersion: param.AppVersion,
		Status:     1, // 已激活
	}

	msg := system.SaveDevice(device)
	if msg != "操作成功" {
		c.JSON(http.StatusOK, R.ReturnFailMsg(msg))
		return
	}

	// 更新激活码状态
	code.Status = 2 // 已使用
	code.ActivateTime = time.Now()
	system.SaveActivateCode(code)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"deviceId": device.DeviceId,
		"msg":      msg,
	}))
}

// GetDeviceInfo 获取设备信息
func GetDeviceInfo(c *gin.Context) {
	deviceId := c.Query("deviceId")
	if deviceId == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备 ID 不能为空"))
		return
	}

	id, _ := strconv.ParseInt(deviceId, 10, 64)
	device := system.FindDeviceById(id)
	if device.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备不存在"))
		return
	}

	// 获取执法人员信息
	official := system.FindOfficialById(device.OfficialId)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"device":   device,
		"official": official,
	}))
}

// ListDevice 查询设备列表
func ListDevice(c *gin.Context) {
	param := system.SearchDeviceParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.DeviceModel = c.Query("deviceModel")
	param.OsType = c.Query("osType")
	param.Status, _ = strconv.Atoi(c.DefaultQuery("status", "-1"))

	result := system.SelectDeviceList(param)

	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// UpdateDevice 修改设备信息
func UpdateDevice(c *gin.Context) {
	var device system.SysDevice
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if device.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备 ID 不能为空"))
		return
	}

	msg := system.SaveDevice(device)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteDevice 删除设备
func DeleteDevice(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteDevice(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DisableDevice 禁用设备
func DisableDevice(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	device := system.FindDeviceById(id)
	if device.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备不存在"))
		return
	}

	device.Status = 0 // 禁用
	msg := system.SaveDevice(device)

	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateLastLogin 更新最后登录信息
func UpdateLastLogin(c *gin.Context) {
	var param struct {
		DeviceId int64  `json:"deviceId"`
		LoginIP  string `json:"loginIP"`
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if param.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备 ID 不能为空"))
		return
	}

	device := system.FindDeviceById(param.DeviceId)
	if device.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备不存在"))
		return
	}

	device.LastLoginIp = param.LoginIP
	msg := system.SaveDevice(device)

	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}
