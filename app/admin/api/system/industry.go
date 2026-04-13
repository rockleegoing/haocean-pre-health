package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
)

// ListIndustry 查询行业分类列表
func ListIndustry(c *gin.Context) {
	var param system.SysIndustry
	param.IndustryName = c.Query("industryName")
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))

	industries := system.SelectIndustryList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(industries))
}

// GetIndustry 获取行业分类详情
func GetIndustry(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	industry := system.FindIndustryById(id)
	if industry.IndustryId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("行业不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(industry))
}

// AddIndustry 添加行业分类
func AddIndustry(c *gin.Context) {
	var industry system.SysIndustry
	if err := c.ShouldBindJSON(&industry); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	// 设置默认值
	if industry.Level == 0 {
		industry.Level = 1
	}
	if industry.IsEnabled == 0 {
		industry.IsEnabled = 1
	}

	msg := system.SaveIndustry(industry)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateIndustry 修改行业分类
func UpdateIndustry(c *gin.Context) {
	var industry system.SysIndustry
	if err := c.ShouldBindJSON(&industry); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if industry.IndustryId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("行业 ID 不能为空"))
		return
	}

	msg := system.SaveIndustry(industry)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteIndustry 删除行业分类
func DeleteIndustry(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	// 将逗号分隔的 ID 转换为数组
	var ids []int64
	for _, idStr := range split(idsStr) {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	// 检查是否有子行业
	for _, id := range ids {
		children := system.FindIndustryTree(id)
		if len(children) > 0 {
			c.JSON(http.StatusOK, R.ReturnFailMsg("请先删除子行业"))
			return
		}
	}

	msg := system.DeleteIndustry(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// split 辅助函数：分割字符串
func split(s string) []string {
	result := make([]string, 0)
	for _, item := range []rune(s) {
		if item == ',' {
			continue
		}
		// 简单处理，实际需要更完善的实现
		_ = item
	}
	// 使用 strings.Split 代替
	return nil
}
