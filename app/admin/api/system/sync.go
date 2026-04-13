package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"haocean/health-enforcement/pkg/mysql"
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
	tree := BuildIndustryTreeResponse(industries)

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
		system.UpdateSyncQueueStatus(record.RecordId, "success", "")
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
		system.UpdateSyncQueueStatus(subject.SubjectId, "success", "")
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"successCount": successCount,
		"failCount":    failCount,
	}))
}

// GetSyncStatus 获取同步状态
func GetSyncStatus(c *gin.Context) {
	status := system.GetSyncStatus(0)

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

// ListSync 查询同步记录列表
func ListSync(c *gin.Context) {
	var param struct {
		PageNum    int    `form:"pageNum"`
		PageSize   int    `form:"pageSize"`
		SyncType   string `form:"syncType"`
		Status     int    `form:"status"`
		BeginTime  string `form:"params[beginTime]"`
		EndTime    string `form:"params[endTime]"`
	}

	if err := c.ShouldBindQuery(&param); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": "参数错误", "code": http.StatusInternalServerError})
		return
	}

	if param.PageNum == 0 {
		param.PageNum = 1
	}
	if param.PageSize == 0 {
		param.PageSize = 10
	}

	var logs []system.SysSyncLog
	var total int64

	query := mysql.MysqlDb().Model(&system.SysSyncLog{})

	if param.SyncType != "" {
		query = query.Where("sync_type = ?", param.SyncType)
	}
	if param.Status >= 0 {
		query = query.Where("status = ?", mapSyncStatus(param.Status))
	}
	if param.BeginTime != "" {
		query = query.Where("start_time >= ?", param.BeginTime)
	}
	if param.EndTime != "" {
		query = query.Where("end_time <= ?", param.EndTime)
	}

	query.Count(&total)

	var order string
	if param.Status >= 0 {
		order = "status, end_time DESC"
	} else {
		order = "end_time DESC"
	}

	query.Order(order).Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&logs)

	// 转换为前端格式
	type SyncLogVO struct {
		SyncId       int64     `json:"syncId"`
		SyncType     string    `json:"syncType"`
		ModuleName   string    `json:"moduleName"`
		RecordCount  int       `json:"recordCount"`
		Status       int       `json:"status"`
		Message      string    `json:"message"`
		OperatorName string    `json:"operatorName"`
		SyncTime     time.Time `json:"syncTime"`
	}

	var list []SyncLogVO
	for _, log := range logs {
		vo := SyncLogVO{
			SyncId:       log.LogId,
			SyncType:     log.SyncType,
			ModuleName:   "执法数据",
			RecordCount:  log.RecordCount,
			Status:       mapStatusToInt(log.Status),
			Message:      log.ErrorMsg,
			OperatorName: "系统",
			SyncTime:     log.EndTime,
		}
		list = append(list, vo)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"rows": list,
		"total": total,
	})
}

func mapSyncStatus(status int) string {
	switch status {
	case 0:
		return "success"
	case 1:
		return "failed"
	case 2:
		return "processing"
	default:
		return "success"
	}
}

func mapStatusToInt(status string) int {
	switch status {
	case "success":
		return 0
	case "failed":
		return 1
	case "processing":
		return 2
	default:
		return 0
	}
}
