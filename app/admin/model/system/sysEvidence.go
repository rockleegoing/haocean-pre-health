package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysEvidence 证据材料
type SysEvidence struct {
	EvidenceId    int64     `gorm:"primary_key;autoIncrement" json:"evidenceId"`    // 证据 ID
	RecordId      int64     `gorm:"index" json:"recordId"`                          // 执法记录 ID
	EvidenceNo    string    `gorm:"size:50" json:"evidenceNo"`                      // 证据编号
	Type          string    `gorm:"size:20" json:"type"`                            // 证据类型（photo/audio/video/document）
	Title         string    `gorm:"size:100" json:"title"`                          // 证据标题
	Description   string    `gorm:"size:500" json:"description"`                    // 证据描述
	FilePath      string    `gorm:"size:255" json:"filePath"`                       // 文件路径
	FileName      string    `gorm:"size:100" json:"fileName"`                       // 文件名
	FileSize      int64     `gorm:"default:0" json:"fileSize"`                      // 文件大小（字节）
	FileType      string    `gorm:"size:50" json:"fileType"`                        // 文件类型（mime type）
	Duration      int       `gorm:"default:0" json:"duration"`                      // 时长（秒，音频/视频）
	ThumbnailPath string    `gorm:"size:255" json:"thumbnailPath"`                  // 缩略图路径
	Latitude      float64   `gorm:"type:decimal(10,8)" json:"latitude"`             // 拍摄地点纬度
	Longitude     float64   `gorm:"type:decimal(11,8)" json:"longitude"`            // 拍摄地点经度
	CaptureTime   time.Time `gorm:"type:datetime" json:"captureTime"`               // 采集时间
	UploadBy      string    `gorm:"size:64" json:"uploadBy"`                        // 上传人
	CreateTime    time.Time `gorm:"autoCreateTime" json:"createTime"`               // 创建时间
	SyncStatus    int       `gorm:"default:0" json:"syncStatus"`                    // 同步状态（0:待同步/1:已同步）
}

func (SysEvidence) TableName() string {
	return "law_evidence"
}

// FindEvidenceById 根据 ID 查询证据
func FindEvidenceById(id int64) SysEvidence {
	var evidence SysEvidence
	mysql.MysqlDb().Where("evidence_id = ?", id).First(&evidence)
	return evidence
}

// FindEvidenceByRecordId 根据记录 ID 查询证据列表
func FindEvidenceByRecordId(recordId int64) []SysEvidence {
	var evidences []SysEvidence
	mysql.MysqlDb().Where("record_id = ?", recordId).Order("create_time").Find(&evidences)
	return evidences
}

// SaveEvidence 保存证据
func SaveEvidence(evidence SysEvidence) string {
	if evidence.EvidenceId == 0 {
		mysql.MysqlDb().Create(&evidence)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&evidence)
	return "修改成功"
}

// DeleteEvidence 删除证据
func DeleteEvidence(ids []int64) string {
	mysql.MysqlDb().Delete(&SysEvidence{}, "evidence_id IN ?", ids)
	return "删除成功"
}

// SelectEvidenceList 查询证据列表
func SelectEvidenceList(param SearchEvidenceParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysEvidence{})

	if param.RecordId != 0 {
		db = db.Where("record_id = ?", param.RecordId)
	}
	if param.Type != "" {
		db = db.Where("type = ?", param.Type)
	}

	var total int64
	db.Count(&total)

	var result []SysEvidence
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("create_time DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}

// SearchEvidenceParam 证据搜索参数
type SearchEvidenceParam struct {
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
	RecordId int64  `form:"recordId"`
	Type     string `form:"type"`
}
