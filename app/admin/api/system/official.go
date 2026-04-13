package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ListOfficial 查询执法人员列表
func ListOfficial(c *gin.Context) {
	param := system.SearchOfficialParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.Realname = c.Query("realname")
	param.BadgeNo = c.Query("badgeNo")
	param.Department = c.Query("department")
	param.Status, _ = strconv.Atoi(c.DefaultQuery("status", "-1"))

	result := system.SelectOfficialList(param)

	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetOfficial 获取执法人员详情
func GetOfficial(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	official := system.FindOfficialById(id)
	if official.OfficialId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("人员不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(official))
}

// AddOfficial 添加执法人员
func AddOfficial(c *gin.Context) {
	var official system.SysOfficial
	if err := c.ShouldBindJSON(&official); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if official.BadgeNo == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("执法证号不能为空"))
		return
	}

	// 检查执法证号是否重复
	if system.IsExistBadgeNo(official.BadgeNo) {
		c.JSON(http.StatusOK, R.ReturnFailMsg("执法证号已存在"))
		return
	}

	msg := system.SaveOfficial(official)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateOfficial 修改执法人员
func UpdateOfficial(c *gin.Context) {
	var official system.SysOfficial
	if err := c.ShouldBindJSON(&official); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if official.OfficialId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("人员 ID 不能为空"))
		return
	}

	msg := system.SaveOfficial(official)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteOfficial 删除执法人员
func DeleteOfficial(c *gin.Context) {
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

	msg := system.DeleteOfficial(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// BindDevice 绑定设备
func BindDevice(c *gin.Context) {
	var param struct {
		OfficialId int64  `json:"officialId"`
		DeviceId   int64  `json:"deviceId"`
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if param.OfficialId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("人员 ID 不能为空"))
		return
	}

	if param.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备 ID 不能为空"))
		return
	}

	// 更新设备绑定关系
	device := system.FindDeviceById(param.DeviceId)
	if device.DeviceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备不存在"))
		return
	}

	device.OfficialId = param.OfficialId
	msg := system.SaveDevice(device)

	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}
