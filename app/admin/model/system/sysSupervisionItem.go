package system

import (
	"encoding/json"
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysSupervisionItem 监管事项
type SysSupervisionItem struct {
	ItemId           int64           `gorm:"primary_key;autoIncrement" json:"itemId"`                // 事项 ID
	ItemName         string          `gorm:"size:100;not null" json:"itemName"`                      // 事项名称
	ParentId         int64           `gorm:"default:0" json:"parentId"`                              // 父级 ID
	Level            int             `gorm:"default:1" json:"level"`                                 // 层级（1:一级/2:二级）
	SupervisionType  string          `gorm:"size:50" json:"supervisionType"`                         // 监管类型
	IndustryIds      json.RawMessage `gorm:"type:json" json:"industryIds"`                           // 关联行业 ID 列表
	StandardLanguageIds json.RawMessage `gorm:"type:json" json:"standardLanguageIds"`                 // 关联规范用语 ID 列表
	CheckPoints      string          `gorm:"type:text" json:"checkPoints"`                           // 检查要点
	LegalBasis       string          `gorm:"type:text" json:"legalBasis"`                            // 法律依据
	SortOrder        int             `gorm:"default:0" json:"sortOrder"`                             // 排序
	IsEnabled        int             `gorm:"default:1" json:"isEnabled"`                             // 是否启用（0:禁用/1:启用）
	CreateBy         string          `gorm:"size:64;default:''" json:"createBy"`                     // 创建者
	CreateTime       time.Time       `gorm:"autoCreateTime" json:"createTime"`                       // 创建时间
	UpdateBy         string          `gorm:"size:64;default:''" json:"updateBy"`                     // 更新者
	UpdateTime       time.Time       `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`        // 更新时间
	Remark           string          `gorm:"size:500" json:"remark"`                                 // 备注

	Children []SysSupervisionItem `gorm:"-" json:"children"` // 子事项（树形结构）
}

func (SysSupervisionItem) TableName() string {
	return "law_supervision_item"
}

// SysSupervisionCategory 监管事项分类（用于前端展示）
type SysSupervisionCategory struct {
	CategoryId   int64     `gorm:"primary_key;autoIncrement" json:"categoryId"`   // 分类 ID
	CategoryName string    `gorm:"size:100" json:"categoryName"`                  // 分类名称
	CategoryCode string    `gorm:"size:50" json:"categoryCode"`                   // 分类代码
	Icon         string    `gorm:"size:100" json:"icon"`                          // 图标
	OrderNum     int       `gorm:"default:0" json:"orderNum"`                     // 排序
	IsEnabled    int       `gorm:"default:1" json:"isEnabled"`                    // 是否启用
	CreateTime   time.Time `gorm:"autoCreateTime" json:"createTime"`              // 创建时间
	UpdateTime   time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
}

func (SysSupervisionCategory) TableName() string {
	return "sys_supervision_category"
}

// FindSupervisionItemById 根据 ID 查询监管事项
func FindSupervisionItemById(id int64) SysSupervisionItem {
	var item SysSupervisionItem
	mysql.MysqlDb().Where("item_id = ?", id).First(&item)
	return item
}

// FindSupervisionItemByName 根据名称查询监管事项
func FindSupervisionItemByName(name string) SysSupervisionItem {
	var item SysSupervisionItem
	mysql.MysqlDb().Where("item_name = ?", name).First(&item)
	return item
}

// SaveSupervisionItem 保存监管事项
func SaveSupervisionItem(item SysSupervisionItem) string {
	if item.ItemId == 0 {
		// 新增
		mysql.MysqlDb().Create(&item)
		return "添加成功"
	}
	// 更新
	mysql.MysqlDb().Save(&item)
	return "修改成功"
}

// DeleteSupervisionItem 删除监管事项
func DeleteSupervisionItem(ids []int64) string {
	mysql.MysqlDb().Delete(&SysSupervisionItem{}, "item_id IN ?", ids)
	return "删除成功"
}

// SelectSupervisionItemList 查询监管事项列表
func SelectSupervisionItemList(param SearchSupervisionItemParam) []SysSupervisionItem {
	db := mysql.MysqlDb().Model(&SysSupervisionItem{})

	if param.ItemName != "" {
		db = db.Where("item_name LIKE ?", "%"+param.ItemName+"%")
	}
	if param.SupervisionType != "" {
		db = db.Where("supervision_type = ?", param.SupervisionType)
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}
	if param.ParentId != -1 {
		db = db.Where("parent_id = ?", param.ParentId)
	}

	var result []SysSupervisionItem
	db.Order("sort_order, create_time DESC").Find(&result)

	// 如果需要构建树形结构
	if param.Tree {
		return buildSupervisionTree(result, 0)
	}

	return result
}

// SearchSupervisionItemParam 监管事项搜索参数
type SearchSupervisionItemParam struct {
	PageNum        int    `form:"pageNum"`
	PageSize       int    `form:"pageSize"`
	ItemName       string `form:"itemName"`
	SupervisionType string `form:"supervisionType"`
	IsEnabled      int    `form:"isEnabled"`
	ParentId       int64  `form:"parentId"`
	Tree           bool   `form:"tree"`
	Params         Params `form:"Params"`
}

// buildSupervisionTree 构建树形结构
func buildSupervisionTree(all []SysSupervisionItem, parentId int64) []SysSupervisionItem {
	var tree []SysSupervisionItem
	for _, item := range all {
		if item.ParentId == parentId {
			item.Children = buildSupervisionTree(all, item.ItemId)
			tree = append(tree, item)
		}
	}
	return tree
}

// FindSupervisionChildren 查询子事项
func FindSupervisionChildren(parentId int64) []SysSupervisionItem {
	var items []SysSupervisionItem
	mysql.MysqlDb().Where("parent_id = ?", parentId).Order("sort_order").Find(&items)

	// 递归查询子事项
	for i := range items {
		items[i].Children = FindSupervisionChildren(items[i].ItemId)
	}

	return items
}

// SelectSupervisionPageList 查询监管事项分页列表
func SelectSupervisionPageList(param SearchSupervisionItemParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysSupervisionItem{})

	if param.ItemName != "" {
		db = db.Where("item_name LIKE ?", "%"+param.ItemName+"%")
	}
	if param.SupervisionType != "" {
		db = db.Where("supervision_type = ?", param.SupervisionType)
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}
	if param.ParentId != -1 {
		db = db.Where("parent_id = ?", param.ParentId)
	}

	var total int64
	db.Count(&total)

	var result []SysSupervisionItem
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("sort_order, create_time DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}
