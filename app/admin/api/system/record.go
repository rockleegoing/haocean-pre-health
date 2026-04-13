package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ListRecord 查询执法记录列表
func ListRecord(c *gin.Context) {
	param := system.SearchRecordParam{
		PageNum:  1,
		PageSize: 10,
	}

	if pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1")); pageNum > 0 {
		param.PageNum = pageNum
	}
	if pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10")); pageSize > 0 {
		param.PageSize = pageSize
	}

	param.SubjectId, _ = strconv.ParseInt(c.DefaultQuery("subjectId", "0"), 10, 64)
	param.IndustryId, _ = strconv.ParseInt(c.DefaultQuery("industryId", "0"), 10, 64)
	param.Status, _ = strconv.Atoi(c.DefaultQuery("status", "-1"))
	param.CheckType = c.Query("checkType")
	param.BeginTime = c.Query("beginTime")
	param.EndTime = c.Query("endTime")

	result := system.SelectRecordList(param)
	result.Code = http.StatusOK
	result.Msg = "查询成功"

	c.JSON(http.StatusOK, result)
}

// GetRecord 获取执法记录详情
func GetRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	record := system.FindRecordById(id)
	if record.RecordId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("记录不存在"))
		return
	}

	// 获取关联的证据列表
	evidences := system.FindEvidenceByRecordId(id)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"record":    record,
		"evidences": evidences,
	}))
}

// AddRecord 添加执法记录
func AddRecord(c *gin.Context) {
	var record system.SysEnforcementRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if record.SubjectId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("请选择监管单位"))
		return
	}

	// 获取单位信息
	subject := system.FindSubjectById(record.SubjectId)
	record.SubjectName = subject.Name
	record.IndustryId = subject.IndustryId
	record.IndustryName = subject.IndustryName

	msg := system.SaveRecord(record)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateRecord 修改执法记录
func UpdateRecord(c *gin.Context) {
	var record system.SysEnforcementRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if record.RecordId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("记录 ID 不能为空"))
		return
	}

	msg := system.SaveRecord(record)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteRecord 删除执法记录
func DeleteRecord(c *gin.Context) {
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

	msg := system.DeleteRecord(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// SubmitRecord 上报执法记录
func SubmitRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	record := system.FindRecordById(id)
	if record.RecordId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("记录不存在"))
		return
	}

	record.Status = 2 // 已上报
	msg := system.SaveRecord(record)

	// 添加到同步队列
	system.AddSyncQueue(system.SysSyncQueue{
		TableName:  "law_enforcement_record",
		RecordId:   id,
		Action:     "update",
		SyncType:   "app_to_server",
		Priority:   1,
		Status:     "pending",
		RetryCount: 0,
	})

	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UploadEvidence 上传证据
func UploadEvidence(c *gin.Context) {
	var evidence system.SysEvidence

	recordIdStr := c.PostForm("recordId")
	evidence.RecordId, _ = strconv.ParseInt(recordIdStr, 10, 64)
	evidence.Type = c.PostForm("type")
	evidence.Title = c.PostForm("title")
	evidence.Description = c.PostForm("description")

	if evidence.RecordId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("执法记录 ID 不能为空"))
		return
	}

	// 上传文件
	file, err := c.FormFile("file")
	if err == nil {
		filename := "evidence_" + time.Now().Format("20060102150405") + "_" + file.Filename
		filePath := "./static/evidence/" + filename

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusOK, R.ReturnFailMsg("文件保存失败"))
			return
		}

		evidence.FilePath = filePath
		evidence.FileName = file.Filename
		evidence.FileSize = file.Size
		evidence.FileType = file.Header.Get("Content-Type")
	}

	msg := system.SaveEvidence(evidence)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteEvidence 删除证据
func DeleteEvidence(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	evidence := system.FindEvidenceById(id)
	if evidence.EvidenceId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("证据不存在"))
		return
	}

	msg := system.DeleteEvidence([]int64{id})
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}
