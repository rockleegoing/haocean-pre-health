package system

import (
	"encoding/json"
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysSyncQueue 同步队列
type SysSyncQueue struct {
	QueueId      int64           `gorm:"primary_key;autoIncrement" json:"queueId"`       // 队列 ID
	TableRef     string          `gorm:"size:50;not null;index;column:table_name" json:"tableName"`        // 表名
	RecordId     int64           `gorm:"not null;index" json:"recordId"`                 // 记录 ID
	Action       string          `gorm:"size:20;not null" json:"action"`                 // 操作类型（insert/update/delete）
	SyncType     string          `gorm:"size:20;default:'app_to_server'" json:"syncType"` // 同步类型
	Data         json.RawMessage `gorm:"type:json" json:"data"`                          // 变更数据
	Priority     int             `gorm:"default:0" json:"priority"`                      // 优先级（0:普通/1:重要/2:紧急）
	Status       string          `gorm:"size:20;default:'pending'" json:"status"`        // 状态
	RetryCount   int             `gorm:"default:0" json:"retryCount"`                    // 重试次数
	ErrorMsg     string          `gorm:"size:500" json:"errorMsg"`                       // 错误信息
	ConflictInfo json.RawMessage `gorm:"type:json" json:"conflictInfo"`                  // 冲突信息
	CreateTime   time.Time       `gorm:"autoCreateTime;index" json:"createTime"`         // 创建时间
	SyncTime     time.Time       `gorm:"type:datetime" json:"syncTime"`                  // 同步时间
	CreateBy     string          `gorm:"size:64" json:"createBy"`                        // 创建者
}

// TableName 返回表名
func (SysSyncQueue) TableName() string {
	return "law_sync_queue"
}

// SysSyncLog 同步日志
type SysSyncLog struct {
	LogId         int64           `gorm:"primary_key;autoIncrement" json:"logId"`         // 日志 ID
	DeviceId      int64           `gorm:"index" json:"deviceId"`                          // 设备 ID
	OfficialId    int64           `gorm:"index" json:"officialId"`                        // 执法人员 ID
	SyncType      string          `gorm:"size:20" json:"syncType"`                        // 同步类型（full/incremental）
	SyncTables    json.RawMessage `gorm:"type:json" json:"syncTables"`                    // 同步表列表
	RecordCount   int             `gorm:"default:0" json:"recordCount"`                   // 同步记录数
	SuccessCount  int             `gorm:"default:0" json:"successCount"`                  // 成功数
	FailedCount   int             `gorm:"default:0" json:"failedCount"`                   // 失败数
	ConflictCount int             `gorm:"default:0" json:"conflictCount"`                 // 冲突数
	Duration      int             `gorm:"default:0" json:"duration"`                      // 耗时（秒）
	Status        string          `gorm:"size:20;default:'success'" json:"status"`        // 状态
	ErrorMsg      string          `gorm:"size:500" json:"errorMsg"`                       // 错误信息
	StartTime     time.Time       `gorm:"type:datetime" json:"startTime"`                 // 开始时间
	EndTime       time.Time       `gorm:"type:datetime" json:"endTime"`                   // 结束时间
	CreateTime    time.Time       `gorm:"autoCreateTime" json:"createTime"`               // 创建时间
}

func (SysSyncLog) TableName() string {
	return "law_sync_log"
}

// AddSyncQueue 添加同步队列
func AddSyncQueue(queue SysSyncQueue) {
	mysql.MysqlDb().Create(&queue)
}

// AddSyncQueues 批量添加同步队列
func AddSyncQueues(queues []SysSyncQueue) {
	if len(queues) > 0 {
		mysql.MysqlDb().CreateInBatches(queues, 100)
	}
}

// GetPendingSyncQueue 获取待同步的队列
func GetPendingSyncQueue(limit int) []SysSyncQueue {
	var queues []SysSyncQueue
	mysql.MysqlDb().Where("status = ?", "pending").
		Order("priority DESC, create_time ASC").
		Limit(limit).
		Find(&queues)
	return queues
}

// UpdateSyncQueueStatus 更新同步队列状态
func UpdateSyncQueueStatus(queueId int64, status string, errorMsg string) {
	updateMap := map[string]interface{}{
		"status": status,
	}
	if status == "success" {
		updateMap["sync_time"] = time.Now()
	}
	if errorMsg != "" {
		updateMap["error_msg"] = errorMsg
		updateMap["retry_count"] = mysql.MysqlDb().Model(&SysSyncQueue{}).
			Where("queue_id = ?", queueId).
			Select("retry_count + 1")
	}
	mysql.MysqlDb().Model(&SysSyncQueue{}).Where("queue_id = ?", queueId).Updates(updateMap)
}

// AddSyncLog 添加同步日志
func AddSyncLog(log SysSyncLog) {
	mysql.MysqlDb().Create(&log)
}

// GetSyncStatus 获取同步状态
func GetSyncStatus(deviceId int64) map[string]interface{} {
	var pendingCount, failedCount int64
	mysql.MysqlDb().Model(&SysSyncQueue{}).Where("status = ?", "pending").Count(&pendingCount)
	mysql.MysqlDb().Model(&SysSyncQueue{}).Where("status = ?", "failed").Count(&failedCount)

	var lastSync SysSyncLog
	mysql.MysqlDb().Where("device_id = ?", deviceId).Order("create_time DESC").First(&lastSync)

	return map[string]interface{}{
		"pending_count":    pendingCount,
		"failed_count":     failedCount,
		"last_sync_time":   lastSync.EndTime,
		"last_sync_status": lastSync.Status,
	}
}

// GetSyncData 获取同步数据（供移动端同步使用）
func GetSyncData(deviceId string, lastSyncTime string) []map[string]interface{} {
	// 获取待同步的数据
	queues := GetPendingSyncQueue(100)
	var syncData []map[string]interface{}

	for _, queue := range queues {
		data := map[string]interface{}{
			"queue_id":   queue.QueueId,
			"table_name": queue.TableRef,
			"record_id":  queue.RecordId,
			"action":     queue.Action,
			"data":       queue.Data,
		}
		syncData = append(syncData, data)
	}

	return syncData
}

// RetrySyncQueue 重试同步队列
func RetrySyncQueue(queueId int64) string {
	var queue SysSyncQueue
	mysql.MysqlDb().Where("queue_id = ?", queueId).First(&queue)
	if queue.QueueId == 0 {
		return "队列不存在"
	}
	mysql.MysqlDb().Model(&SysSyncQueue{}).Where("queue_id = ?", queueId).Updates(map[string]interface{}{
		"status":       "pending",
		"error_msg":    "",
		"retry_count":  queue.RetryCount + 1,
	})
	return "操作成功"
}
