package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysOfficial 执法人员
type SysOfficial struct {
	OfficialId  int64     `gorm:"primary_key;autoIncrement" json:"officialId"`  // 执法人员 ID
	UserId      int64     `gorm:"uniqueIndex" json:"userId"`                    // 关联用户 ID
	BadgeNo     string    `gorm:"size:20" json:"badgeNo"`                       // 执法证号
	Department  string    `gorm:"size:100" json:"department"`                   // 所属部门
	Position    string    `gorm:"size:50" json:"position"`                      // 职位
	LawType     string    `gorm:"size:50" json:"lawType"`                       // 执法类型
	Status      int       `gorm:"default:1" json:"status"`                      // 状态（0:禁用/1:启用）
	CreateBy    string    `gorm:"size:64;default:''" json:"createBy"`           // 创建者
	CreateTime  time.Time `gorm:"autoCreateTime" json:"createTime"`             // 创建时间
	UpdateBy    string    `gorm:"size:64;default:''" json:"updateBy"`           // 更新者
	UpdateTime  time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark      string    `gorm:"size:500" json:"remark"`                       // 备注
}

func (SysOfficial) TableName() string {
	return "law_official"
}

// FindOfficialById 根据 ID 查询执法人员
func FindOfficialById(id int64) SysOfficial {
	var official SysOfficial
	mysql.MysqlDb().Where("official_id = ?", id).First(&official)
	return official
}

// FindOfficialByUserId 根据用户 ID 查询执法人员
func FindOfficialByUserId(userId int64) SysOfficial {
	var official SysOfficial
	mysql.MysqlDb().Where("user_id = ?", userId).First(&official)
	return official
}

// SaveOfficial 保存执法人员
func SaveOfficial(official SysOfficial) string {
	if official.OfficialId == 0 {
		mysql.MysqlDb().Create(&official)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&official)
	return "修改成功"
}

// DeleteOfficial 删除执法人员
func DeleteOfficial(ids []int64) string {
	mysql.MysqlDb().Delete(&SysOfficial{}, "official_id IN ?", ids)
	return "删除成功"
}
