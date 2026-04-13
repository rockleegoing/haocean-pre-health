package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysActivateCode 激活码
type SysActivateCode struct {
	CodeId           int64     `gorm:"primary_key;autoIncrement" json:"codeId"`              // 激活码 ID
	ActivateCode     string    `gorm:"size:20;uniqueIndex" json:"activateCode"`              // 激活码
	OfficialId       int64     `gorm:"index" json:"officialId"`                              // 绑定执法人员 ID
	BatchNo          string    `gorm:"size:50" json:"batchNo"`                               // 批次号
	ExpireTime       time.Time `gorm:"type:datetime" json:"expireTime"`                      // 过期时间
	Status           int       `gorm:"default:0" json:"status"`                              // 状态（0:未使用/1:已激活/2:已过期/3:已禁用）
	ActivateTime     time.Time `gorm:"type:datetime" json:"activateTime"`                    // 激活时间
	ActivateDeviceId int64     `gorm:"" json:"activateDeviceId"`                             // 激活设备 ID
	CreateBy         string    `gorm:"size:64;default:''" json:"createBy"`                   // 创建者
	CreateTime       time.Time `gorm:"autoCreateTime" json:"createTime"`                     // 创建时间
	UpdateBy         string    `gorm:"size:64;default:''" json:"updateBy"`                   // 更新者
	UpdateTime       time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`      // 更新时间
	Remark           string    `gorm:"size:500" json:"remark"`                               // 备注
}

func (SysActivateCode) TableName() string {
	return "law_activate_code"
}

// FindActivateCodeByCode 根据激活码查询
func FindActivateCodeByCode(code string) SysActivateCode {
	var activateCode SysActivateCode
	mysql.MysqlDb().Where("activate_code = ?", code).First(&activateCode)
	return activateCode
}

// FindActivateCodeById 根据 ID 查询激活码
func FindActivateCodeById(id int64) SysActivateCode {
	var activateCode SysActivateCode
	mysql.MysqlDb().Where("code_id = ?", id).First(&activateCode)
	return activateCode
}

// SaveActivateCode 保存激活码
func SaveActivateCode(activateCode SysActivateCode) string {
	if activateCode.CodeId == 0 {
		mysql.MysqlDb().Create(&activateCode)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&activateCode)
	return "修改成功"
}

// DeleteActivateCode 删除激活码
func DeleteActivateCode(ids []int64) string {
	mysql.MysqlDb().Delete(&SysActivateCode{}, "code_id IN ?", ids)
	return "删除成功"
}

// SelectActivateCodeList 查询激活码列表
func SelectActivateCodeList(param SearchActivateCodeParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysActivateCode{})

	if param.BatchNo != "" {
		db = db.Where("batch_no = ?", param.BatchNo)
	}
	if param.Status != -1 {
		db = db.Where("status = ?", param.Status)
	}

	var total int64
	db.Count(&total)

	var result []SysActivateCode
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("create_time DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}

// SearchActivateCodeParam 激活码搜索参数
type SearchActivateCodeParam struct {
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
	BatchNo  string `form:"batchNo"`
	Status   int    `form:"status"`
}

// GenerateActivateCode 生成激活码
func GenerateActivateCode(batchNo string, count int, expireDays int) []SysActivateCode {
	codes := make([]SysActivateCode, 0, count)
	for i := 0; i < count; i++ {
		code := generateRandomCode(8)
		codes = append(codes, SysActivateCode{
			ActivateCode: code,
			BatchNo:      batchNo,
			ExpireTime:   time.Now().AddDate(0, 0, expireDays),
			Status:       0,
		})
	}
	if len(codes) > 0 {
		mysql.MysqlDb().CreateInBatches(codes, 100)
	}
	return codes
}

// generateRandomCode 生成随机激活码
func generateRandomCode(length int) string {
	chars := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = chars[time.Now().UnixNano()%int64(len(chars))]
	}
	return string(result)
}
