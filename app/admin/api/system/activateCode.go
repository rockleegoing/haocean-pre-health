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

// ListActivateCode 查询激活码列表
func ListActivateCode(c *gin.Context) {
	param := system.SearchActivateCodeParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.BatchNo = c.Query("batchNo")
	param.Status, _ = strconv.Atoi(c.DefaultQuery("status", "-1"))

	result := system.SelectActivateCodeList(param)
	result.Code = http.StatusOK
	result.Msg = "查询成功"

	c.JSON(http.StatusOK, result)
}

// GenerateActivateCode 生成激活码
func GenerateActivateCode(c *gin.Context) {
	var param struct {
		BatchNo   string `json:"batchNo"`
		Count     int    `json:"count"`
		ExpireDay int    `json:"expireDay"`
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if param.Count <= 0 || param.Count > 1000 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("生成数量必须在 1-1000 之间"))
		return
	}

	if param.ExpireDay <= 0 {
		param.ExpireDay = 30 // 默认 30 天
	}

	codes := system.GenerateActivateCode(param.BatchNo, param.Count, param.ExpireDay)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"count": len(codes),
		"codes": codes,
	}))
}

// GetActivateCode 获取激活码详情
func GetActivateCode(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	code := system.FindActivateCodeById(id)
	if code.CodeId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(code))
}

// DeleteActivateCode 删除激活码
func DeleteActivateCode(c *gin.Context) {
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

	msg := system.DeleteActivateCode(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DisableActivateCode 禁用激活码
func DisableActivateCode(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	code := system.FindActivateCodeById(id)
	if code.CodeId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("激活码不存在"))
		return
	}

	code.Status = 3 // 已禁用
	code.UpdateTime = time.Now()
	msg := system.SaveActivateCode(code)

	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}
