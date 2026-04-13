package system

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ListTemplate 查询文书模板列表
func ListTemplate(c *gin.Context) {
	param := system.SearchTemplateParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.TemplateName = c.Query("templateName")
	param.CategoryId, _ = strconv.ParseInt(c.DefaultQuery("categoryId", "0"), 10, 64)
	param.IndustryId, _ = strconv.ParseInt(c.DefaultQuery("industryId", "0"), 10, 64)
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))

	result := system.SelectTemplateList(param)
	result.Code = http.StatusOK
	result.Msg = "查询成功"

	c.JSON(http.StatusOK, result)
}

// GetTemplate 获取文书模板详情
func GetTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	template := system.FindTemplateById(id)
	if template.TemplateId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("模板不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(template))
}

// UploadTemplate 上传文书模板
func UploadTemplate(c *gin.Context) {
	var template system.SysDocumentTemplate

	// 获取表单数据
	template.TemplateName = c.PostForm("templateName")
	template.CategoryId, _ = strconv.ParseInt(c.PostForm("categoryId"), 10, 64)
	template.CategoryName = c.PostForm("categoryName")
	template.IndustryId, _ = strconv.ParseInt(c.PostForm("industryId"), 10, 64)
	template.IndustryName = c.PostForm("industryName")
	template.TemplateType = c.PostForm("templateType")
	template.Version = c.PostForm("version")
	template.IsEnabled = 1

	// 获取填空项定义
	fieldsStr := c.PostForm("fields")
	if fieldsStr != "" {
		template.Fields = []byte(fieldsStr)
	}

	// 上传文件
	file, err := c.FormFile("file")
	if err == nil {
		// 保存到临时目录
		filename := "template_" + time.Now().Format("20060102150405") + "_" + file.Filename
		filePath := "./static/template/" + filename

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusOK, R.ReturnFailMsg("文件保存失败"))
			return
		}

		template.FilePath = filePath
		template.FileName = file.Filename
		template.FileSize = file.Size
	}

	msg := system.SaveTemplate(template)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateTemplate 修改文书模板
func UpdateTemplate(c *gin.Context) {
	var template system.SysDocumentTemplate
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if template.TemplateId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("模板 ID 不能为空"))
		return
	}

	msg := system.SaveTemplate(template)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteTemplate 删除文书模板
func DeleteTemplate(c *gin.Context) {
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

	msg := system.DeleteTemplate(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// PreviewTemplate 预览文书模板
func PreviewTemplate(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	template := system.FindTemplateById(id)
	if template.TemplateId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("模板不存在"))
		return
	}

	// 如果是 base64 内容，直接返回
	if template.FileContent != "" {
		data, err := base64.StdEncoding.DecodeString(template.FileContent)
		if err != nil {
			c.JSON(http.StatusOK, R.ReturnFailMsg("模板内容解析失败"))
			return
		}

		c.Header("Content-Type", "application/msword")
		c.Header("Content-Disposition", "inline; filename="+template.TemplateName+".docx")
		c.Data(http.StatusOK, "application/msword", data)
		return
	}

	// 从文件读取
	c.File(template.FilePath)
}

// ListTemplateCategory 查询模板分类列表
func ListTemplateCategory(c *gin.Context) {
	// TODO: 实现模板分类查询
	c.JSON(http.StatusOK, R.ReturnSuccess([]interface{}{}))
}
