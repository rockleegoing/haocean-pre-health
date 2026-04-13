package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysIndustry 行业分类
type SysIndustry struct {
	IndustryId   int64     `gorm:"primary_key;autoIncrement" json:"industryId"`   // 行业 ID
	IndustryCode string    `gorm:"size:50" json:"industryCode"`                   // 行业代码
	IndustryName string    `gorm:"size:100" json:"industryName"`                  // 行业名称
	ParentId     int64     `gorm:"default:0" json:"parentId"`                     // 父级 ID
	Level        int       `gorm:"default:1" json:"level"`                        // 层级（1:一级/2:二级）
	IsEnabled    int       `gorm:"default:1" json:"isEnabled"`                    // 是否启用（0:禁用/1:启用）
	OrderNum     int       `gorm:"default:0" json:"orderNum"`                     // 排序
	CreateBy     string    `gorm:"size:64;default:''" json:"createBy"`            // 创建者
	CreateTime   time.Time `gorm:"autoCreateTime" json:"createTime"`              // 创建时间
	UpdateBy     string    `gorm:"size:64;default:''" json:"updateBy"`            // 更新者
	UpdateTime   time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark       string    `gorm:"size:500" json:"remark"`                        // 备注

	Children []SysIndustry `gorm:"-" json:"children"` // 子行业（树形结构）
}

func (SysIndustry) TableName() string {
	return "law_industry"
}

// FindIndustryTree 查询行业分类树
func FindIndustryTree(parentId int64) []SysIndustry {
	var industries []SysIndustry
	mysql.MysqlDb().Where("parent_id = ?", parentId).Order("order_num").Find(&industries)

	// 递归查询子行业
	for i := range industries {
		industries[i].Children = FindIndustryTree(industries[i].IndustryId)
	}

	return industries
}

// FindIndustryById 根据 ID 查询行业
func FindIndustryById(id int64) SysIndustry {
	var industry SysIndustry
	mysql.MysqlDb().Where("industry_id = ?", id).First(&industry)
	return industry
}

// FindIndustryByCode 根据代码查询行业
func FindIndustryByCode(code string) SysIndustry {
	var industry SysIndustry
	mysql.MysqlDb().Where("industry_code = ?", code).First(&industry)
	return industry
}

// SaveIndustry 保存行业分类
func SaveIndustry(industry SysIndustry) string {
	if industry.IndustryId == 0 {
		// 新增
		mysql.MysqlDb().Create(&industry)
		return "添加成功"
	}
	// 更新
	mysql.MysqlDb().Save(&industry)
	return "修改成功"
}

// DeleteIndustry 删除行业分类
func DeleteIndustry(ids []int64) string {
	mysql.MysqlDb().Delete(&SysIndustry{}, "industry_id IN ?", ids)
	return "删除成功"
}

// SelectIndustryList 查询行业分类列表
func SelectIndustryList(param SysIndustry) []SysIndustry {
	db := mysql.MysqlDb().Model(&SysIndustry{})

	if param.IndustryName != "" {
		db = db.Where("industry_name LIKE ?", "%"+param.IndustryName+"%")
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}

	var result []SysIndustry
	db.Order("order_num").Find(&result)

	// 构建树形结构
	return buildIndustryTree(result, 0)
}

// buildIndustryTree 构建树形结构
func buildIndustryTree(all []SysIndustry, parentId int64) []SysIndustry {
	var tree []SysIndustry
	for _, item := range all {
		if item.ParentId == parentId {
			item.Children = buildIndustryTree(all, item.IndustryId)
			tree = append(tree, item)
		}
	}
	return tree
}
