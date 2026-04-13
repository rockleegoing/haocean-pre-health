package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"time"
)

// CheckSync 检查数据更新
func CheckSync(c *gin.Context) {
	deviceId := c.Query("deviceId")
	lastSyncTime := c.Query("lastSyncTime")

	if deviceId == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("设备 ID 不能为空"))
		return
	}

	// 获取需要同步的数据
	syncData := system.GetSyncData(deviceId, lastSyncTime)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"hasUpdate": len(syncData) > 0,
		"data":      syncData,
		"syncTime":  time.Now().Format("2006-01-02 15:04:05"),
	}))
}

// SyncIndustries 同步行业分类
func SyncIndustries(c *gin.Context) {
	industries := system.SelectIndustryList(system.SysIndustry{})

	// 构建树形结构
	tree := system.BuildIndustryTreeResponse(industries)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"industries": tree,
		"updateTime": time.Now().Format("2006-01-02 15:04:05"),
	}))
}

// SyncTemplates 同步文书模板
func SyncTemplates(c *gin.Context) {
	industryId, _ := strconv.ParseInt(c.Query("industryId"), 10, 64)

	templates := system.FindTemplatesByIndustryId(industryId)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"templates":  templates,
		"updateTime": time.Now().Format("2006-01-02 15:04:05"),
	}))
}

// SyncRecords 上报执法记录
func SyncRecords(c *gin.Context) {
	var records []system.SysEnforcementRecord
	if err := c.ShouldBindJSON(&records); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	successCount := 0
	failCount := 0

	for _, record := range records {
		// 检查是否已存在
		existing := system.FindRecordById(record.RecordId)
		if existing.RecordId != 0 {
			// 更新现有记录
			record.UpdateTime = time.Now()
			msg := system.SaveRecord(record)
			if msg == "操作成功" {
				successCount++
			} else {
				failCount++
			}
		} else {
			// 新增记录
			record.CreateTime = time.Now()
			record.UpdateTime = time.Now()
			msg := system.SaveRecord(record)
			if msg == "操作成功" {
				successCount++
			} else {
				failCount++
			}
		}

		// 更新同步队列状态
		system.UpdateSyncQueueStatus(record.RecordId, "success")
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"successCount": successCount,
		"failCount":    failCount,
	}))
}

// SyncSubjects 上报单位变更
func SyncSubjects(c *gin.Context) {
	var subjects []system.SysSubject
	if err := c.ShouldBindJSON(&subjects); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	successCount := 0
	failCount := 0

	for _, subject := range subjects {
		// 检查是否已存在
		existing := system.FindSubjectById(subject.SubjectId)
		if existing.SubjectId != 0 {
			// 更新现有单位
			subject.UpdateTime = time.Now()
			msg := system.SaveSubject(subject)
			if msg == "操作成功" {
				successCount++
			} else {
				failCount++
			}
		} else {
			// 新增单位
			subject.CreateTime = time.Now()
			subject.UpdateTime = time.Now()
			msg := system.SaveSubject(subject)
			if msg == "操作成功" {
				successCount++
			} else {
				failCount++
			}
		}

		// 更新同步队列状态
		system.UpdateSyncQueueStatus(subject.SubjectId, "success")
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"successCount": successCount,
		"failCount":    failCount,
	}))
}

// GetSyncStatus 获取同步状态
func GetSyncStatus(c *gin.Context) {
	deviceId := c.Query("deviceId")
	recordId, _ := strconv.ParseInt(c.Query("recordId"), 10, 64)

	status := system.GetSyncStatusByRecordId(recordId)

	c.JSON(http.StatusOK, R.ReturnSuccess(status))
}

// RetrySync 重试同步
func RetrySync(c *gin.Context) {
	var param struct {
		QueueId  int64 `json:"queueId"`
		RecordId int64 `json:"recordId"`
	}

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	// 重置同步状态为待处理
	system.RetrySyncQueue(param.QueueId)

	c.JSON(http.StatusOK, R.ReturnSuccess("操作成功"))
}
