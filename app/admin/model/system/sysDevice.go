package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysDevice 设备
type SysDevice struct {
	DeviceId       int64     `gorm:"primary_key;autoIncrement" json:"deviceId"`       // 设备 ID
	OfficialId     int64     `gorm:"index" json:"officialId"`                         // 关联执法人员 ID
	DeviceName     string    `gorm:"size:100" json:"deviceName"`                      // 设备名称
	DeviceModel    string    `gorm:"size:50" json:"deviceModel"`                      // 设备型号
	OsType         string    `gorm:"size:20" json:"osType"`                           // 操作系统类型（iOS/Android）
	OsVersion      string    `gorm:"size:20" json:"osVersion"`                        // 系统版本
	AppVersion     string    `gorm:"size:20" json:"appVersion"`                       // App 版本
	Status         int       `gorm:"default:1" json:"status"`                         // 状态（0:禁用/1:启用）
	LastLoginTime  time.Time `gorm:"type:datetime" json:"lastLoginTime"`              // 最后登录时间
	LastLoginIp    string    `gorm:"size:50" json:"lastLoginIp"`                      // 最后登录 IP
	CreateBy       string    `gorm:"size:64;default:''" json:"createBy"`              // 创建者
	CreateTime     time.Time `gorm:"autoCreateTime" json:"createTime"`                // 创建时间
	UpdateBy       string    `gorm:"size:64;default:''" json:"updateBy"`              // 更新者
	UpdateTime     time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark         string    `gorm:"size:500" json:"remark"`                          // 备注
}

func (SysDevice) TableName() string {
	return "law_device"
}

// FindDeviceById 根据 ID 查询设备
func FindDeviceById(id int64) SysDevice {
	var device SysDevice
	mysql.MysqlDb().Where("device_id = ?", id).First(&device)
	return device
}

// FindDeviceByOfficialId 根据执法人员 ID 查询设备
func FindDeviceByOfficialId(officialId int64) SysDevice {
	var device SysDevice
	mysql.MysqlDb().Where("official_id = ?", officialId).Order("create_time DESC").First(&device)
	return device
}

// SaveDevice 保存设备
func SaveDevice(device SysDevice) string {
	if device.DeviceId == 0 {
		mysql.MysqlDb().Create(&device)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&device)
	return "修改成功"
}

// DeleteDevice 删除设备
func DeleteDevice(ids []int64) string {
	mysql.MysqlDb().Delete(&SysDevice{}, "device_id IN ?", ids)
	return "删除成功"
}

// SearchDeviceParam 设备查询参数
type SearchDeviceParam struct {
	PageNum     int    `form:"pageNum"`
	PageSize    int    `form:"pageSize"`
	DeviceModel string `form:"deviceModel"`
	OsType      string `form:"osType"`
	Status      int    `form:"status"`
}

// SelectDeviceList 查询设备列表
func SelectDeviceList(param SearchDeviceParam) map[string]interface{} {
	var list []SysDevice
	var total int64
	db := mysql.MysqlDb().Model(&SysDevice{})

	if param.DeviceModel != "" {
		db = db.Where("device_model LIKE ?", "%"+param.DeviceModel+"%")
	}
	if param.OsType != "" {
		db = db.Where("os_type = ?", param.OsType)
	}
	if param.Status >= 0 {
		db = db.Where("status = ?", param.Status)
	}

	db.Count(&total)
	db.Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&list)

	return map[string]interface{}{
		"rows":  list,
		"total": total,
	}
}
