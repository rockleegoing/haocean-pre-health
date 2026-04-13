package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysSubject 监管单位
type SysSubject struct {
	SubjectId      int64     `gorm:"primary_key;autoIncrement" json:"subjectId"`      // 单位 ID
	Name           string    `gorm:"size:100;not null" json:"name"`                   // 单位名称
	IndustryId     int64     `gorm:"index" json:"industryId"`                         // 行业分类 ID
	IndustryName   string    `gorm:"size:100" json:"industryName"`                    // 行业名称（冗余）
	Address        string    `gorm:"size:255" json:"address"`                         // 经营地址
	ContactPerson  string    `gorm:"size:50" json:"contactPerson"`                    // 联系人
	ContactPhone   string    `gorm:"size:20" json:"contactPhone"`                     // 联系电话
	LicenseNo      string    `gorm:"size:50" json:"licenseNo"`                        // 许可证号
	LicenseDate    time.Time `gorm:"type:date" json:"licenseDate"`                    // 许可证日期
	LicenseExpiry  time.Time `gorm:"type:date" json:"licenseExpiry"`                  // 许可证有效期
	BusinessScope  string    `gorm:"size:500" json:"businessScope"`                   // 经营范围
	Lat            float64   `gorm:"type:decimal(10,8)" json:"lat"`                   // 纬度
	Lng            float64   `gorm:"type:decimal(11,8)" json:"lng"`                   // 经度
	Status         int       `gorm:"default:1" json:"status"`                         // 状态（0:停用/1:正常）
	RiskLevel      int       `gorm:"default:1" json:"riskLevel"`                      // 风险等级（1:低/2:中/3:高）
	LastCheckDate  time.Time `gorm:"type:date" json:"lastCheckDate"`                  // 最后检查日期
	NextCheckDate  time.Time `gorm:"type:date" json:"nextCheckDate"`                  // 下次检查日期
	CreateBy       string    `gorm:"size:64;default:''" json:"createBy"`              // 创建者
	CreateTime     time.Time `gorm:"autoCreateTime" json:"createTime"`                // 创建时间
	UpdateBy       string    `gorm:"size:64;default:''" json:"updateBy"`              // 更新者
	UpdateTime     time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark         string    `gorm:"size:500" json:"remark"`                          // 备注
	SyncStatus     int       `gorm:"default:1" json:"syncStatus"`                     // 同步状态（0:待同步/1:已同步）
	SyncTime       time.Time `gorm:"type:datetime" json:"syncTime"`                   // 同步时间
}

func (SysSubject) TableName() string {
	return "law_subject"
}

// FindSubjectById 根据 ID 查询单位
func FindSubjectById(id int64) SysSubject {
	var subject SysSubject
	mysql.MysqlDb().Where("subject_id = ?", id).First(&subject)
	return subject
}

// FindSubjectByName 根据名称查询单位
func FindSubjectByName(name string) SysSubject {
	var subject SysSubject
	mysql.MysqlDb().Where("name = ?", name).First(&subject)
	return subject
}

// IsExistSubject 判断单位是否存在
func IsExistSubject(name string) bool {
	var count int64
	mysql.MysqlDb().Model(&SysSubject{}).Where("name = ?", name).Count(&count)
	return count > 0
}

// SaveSubject 保存监管单位
func SaveSubject(subject SysSubject) string {
	if subject.SubjectId == 0 {
		// 新增
		mysql.MysqlDb().Create(&subject)
		return "添加成功"
	}
	// 更新
	mysql.MysqlDb().Save(&subject)
	return "修改成功"
}

// DeleteSubject 删除监管单位
func DeleteSubject(ids []int64) string {
	mysql.MysqlDb().Delete(&SysSubject{}, "subject_id IN ?", ids)
	return "删除成功"
}

// SelectSubjectList 查询监管单位列表
func SelectSubjectList(param SearchSubjectParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysSubject{})

	if param.Name != "" {
		db = db.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.IndustryId != 0 {
		db = db.Where("industry_id = ?", param.IndustryId)
	}
	if param.Status != -1 {
		db = db.Where("status = ?", param.Status)
	}

	var total int64
	db.Count(&total)

	var result []SysSubject
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("create_time DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}

// SearchSubjectParam 监管单位搜索参数
type SearchSubjectParam struct {
	PageNum    int        `form:"pageNum"`
	PageSize   int        `form:"pageSize"`
	Name       string     `form:"name"`
	IndustryId int64      `form:"industryId"`
	Status     int        `form:"status"`
	Params     Params     `form:"Params"`
}

// TableDataInfo 分页响应结构
type TableDataInfo struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
}

// Params 时间范围参数
type Params struct {
	BeginTime string `form:"beginTime"`
	EndTime   string `form:"endTime"`
}
