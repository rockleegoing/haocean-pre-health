package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ListSubject 查询监管单位列表
func ListSubject(c *gin.Context) {
	param := system.SearchSubjectParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.Name = c.Query("name")
	param.IndustryId, _ = strconv.ParseInt(c.DefaultQuery("industryId", "0"), 10, 64)
	param.Status, _ = strconv.Atoi(c.DefaultQuery("status", "-1"))

	result := system.SelectSubjectList(param)
	result.Code = http.StatusOK
	result.Msg = "查询成功"

	c.JSON(http.StatusOK, result)
}

// GetSubject 获取监管单位详情
func GetSubject(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	subject := system.FindSubjectById(id)
	if subject.SubjectId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("单位不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(subject))
}

// AddSubject 添加监管单位
func AddSubject(c *gin.Context) {
	var subject system.SysSubject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if subject.Name == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("单位名称不能为空"))
		return
	}

	// 检查单位名称是否重复
	if system.IsExistSubject(subject.Name) {
		c.JSON(http.StatusOK, R.ReturnFailMsg("单位名称已存在"))
		return
	}

	// 获取行业名称
	if subject.IndustryId != 0 {
		industry := system.FindIndustryById(subject.IndustryId)
		subject.IndustryName = industry.IndustryName
	}

	msg := system.SaveSubject(subject)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateSubject 修改监管单位
func UpdateSubject(c *gin.Context) {
	var subject system.SysSubject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if subject.SubjectId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("单位 ID 不能为空"))
		return
	}

	// 获取行业名称
	if subject.IndustryId != 0 {
		industry := system.FindIndustryById(subject.IndustryId)
		subject.IndustryName = industry.IndustryName
	}

	msg := system.SaveSubject(subject)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteSubject 删除监管单位
func DeleteSubject(c *gin.Context) {
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

	msg := system.DeleteSubject(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// SearchSubject 搜索监管单位
func SearchSubject(c *gin.Context) {
	keyword := c.Query("keyword")
	industryId, _ := strconv.ParseInt(c.DefaultQuery("industryId", "0"), 10, 64)

	db := system.SysSubject{}
	var subjects []system.SysSubject
	query := ""
	args := make([]interface{}, 0)

	if keyword != "" {
		query = "name LIKE ?"
		args = append(args, "%"+keyword+"%")
	}
	if industryId != 0 {
		if query != "" {
			query += " AND "
		}
		query += "industry_id = ?"
		args = append(args, industryId)
	}

	// 这里需要改进查询逻辑
	_ = db
	_ = subjects
	_ = query
	_ = args

	c.JSON(http.StatusOK, R.ReturnFailMsg("搜索功能待实现"))
}
