package system

import (
	"encoding/json"
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysEnforcementRecord 执法记录
type SysEnforcementRecord struct {
	RecordId       int64           `gorm:"primary_key;autoIncrement" json:"recordId"`                // 记录 ID
	RecordNo       string          `gorm:"size:50" json:"recordNo"`                                  // 记录编号
	SubjectId      int64           `gorm:"index" json:"subjectId"`                                   // 单位 ID
	SubjectName    string          `gorm:"size:100" json:"subjectName"`                              // 单位名称（冗余）
	IndustryId     int64           `gorm:"index" json:"industryId"`                                  // 行业分类 ID
	CheckDate      time.Time       `gorm:"type:datetime" json:"checkDate"`                           // 检查日期
	CheckType      string          `gorm:"size:20" json:"checkType"`                                 // 检查类型（日常/专项/复查/投诉）
	Status         int             `gorm:"default:0" json:"status"`                                  // 状态（0:草稿/1:待上报/2:已上报/3:已审核/4:已归档）
	OfficialIds    json.RawMessage `gorm:"type:json" json:"officialIds"`                             // 参与执法人员 ID 列表
	OfficialNames  json.RawMessage `gorm:"type:json" json:"officialNames"`                           // 参与执法人员姓名列表
	CheckResult    string          `gorm:"type:text" json:"checkResult"`                             // 检查结果
	ProblemDesc    string          `gorm:"type:text" json:"problemDesc"`                             // 问题描述
	RectifyOpinion string          `gorm:"type:text" json:"rectifyOpinion"`                          // 整改意见
	RectifyDeadline time.Time      `gorm:"type:date" json:"rectifyDeadline"`                         // 整改期限
	Latitude       float64         `gorm:"type:decimal(10,8)" json:"latitude"`                       // 检查地点纬度
	Longitude      float64         `gorm:"type:decimal(11,8)" json:"longitude"`                      // 检查地点经度
	EvidenceCount  int             `gorm:"default:0" json:"evidenceCount"`                           // 证据数量
	DocumentCount  int             `gorm:"default:0" json:"documentCount"`                           // 文书数量
	CreateBy       string          `gorm:"size:64;default:''" json:"createBy"`                       // 创建者
	CreateTime     time.Time       `gorm:"autoCreateTime" json:"createTime"`                         // 创建时间
	UpdateBy       string          `gorm:"size:64;default:''" json:"updateBy"`                       // 更新者
	UpdateTime     time.Time       `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`          // 更新时间
	Remark         string          `gorm:"size:500" json:"remark"`                                   // 备注
	SyncStatus     int             `gorm:"default:0" json:"syncStatus"`                              // 同步状态（0:待同步/1:已同步）
	SyncTime       time.Time       `gorm:"type:datetime" json:"syncTime"`                            // 同步时间
}

func (SysEnforcementRecord) TableName() string {
	return "law_enforcement_record"
}

// FindRecordById 根据 ID 查询执法记录
func FindRecordById(id int64) SysEnforcementRecord {
	var record SysEnforcementRecord
	mysql.MysqlDb().Where("record_id = ?", id).First(&record)
	return record
}

// SaveRecord 保存执法记录
func SaveRecord(record SysEnforcementRecord) string {
	if record.RecordId == 0 {
		// 生成记录编号
		record.RecordNo = generateRecordNo()
		mysql.MysqlDb().Create(&record)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&record)
	return "修改成功"
}

// DeleteRecord 删除执法记录
func DeleteRecord(ids []int64) string {
	mysql.MysqlDb().Delete(&SysEnforcementRecord{}, "record_id IN ?", ids)
	return "删除成功"
}

// SelectRecordList 查询执法记录列表
func SelectRecordList(param SearchRecordParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysEnforcementRecord{})

	if param.SubjectId != 0 {
		db = db.Where("subject_id = ?", param.SubjectId)
	}
	if param.IndustryId != 0 {
		db = db.Where("industry_id = ?", param.IndustryId)
	}
	if param.Status != -1 {
		db = db.Where("status = ?", param.Status)
	}
	if param.CheckType != "" {
		db = db.Where("check_type = ?", param.CheckType)
	}

	var total int64
	db.Count(&total)

	var result []SysEnforcementRecord
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("check_date DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}

// SearchRecordParam 执法记录搜索参数
type SearchRecordParam struct {
	PageNum     int    `form:"pageNum"`
	PageSize    int    `form:"pageSize"`
	SubjectId   int64  `form:"subjectId"`
	IndustryId  int64  `form:"industryId"`
	Status      int    `form:"status"`
	CheckType   string `form:"checkType"`
	BeginTime   string `form:"beginTime"`
	EndTime     string `form:"endTime"`
}

// generateRecordNo 生成记录编号
func generateRecordNo() string {
	return "JL" + time.Now().Format("20060102150405")
}
