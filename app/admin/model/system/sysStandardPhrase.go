package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysStandardPhraseSupervisionType 监管类型
type SysStandardPhraseSupervisionType struct {
	Id          int64     `gorm:"primary_key;autoIncrement" json:"id"`          // 监管类型 ID
	Name        string    `gorm:"size:100;not null" json:"name"`                // 监管类型名称
	Code        string    `gorm:"size:50" json:"code"`                          // 监管类型代码
	Icon        string    `gorm:"size:100" json:"icon"`                         // 图标
	Description string    `gorm:"size:500" json:"description"`                  // 类型描述
	SortOrder   int       `gorm:"default:0" json:"sortOrder"`                   // 排序
	IsEnabled   int       `gorm:"default:1" json:"isEnabled"`                   // 是否启用（0:禁用/1:启用）
	CreateBy    string    `gorm:"size:64;default:''" json:"createBy"`           // 创建者
	CreateTime  time.Time `gorm:"autoCreateTime" json:"createTime"`             // 创建时间
	UpdateBy    string    `gorm:"size:64;default:''" json:"updateBy"`           // 更新者
	UpdateTime  time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark      string    `gorm:"size:500" json:"remark"`                       // 备注

	Categories []SysStandardPhraseCategory `gorm:"-" json:"categories"` // 关联的规范类别
}

// SysStandardPhraseCategory 规范类别
type SysStandardPhraseCategory struct {
	Id                int64     `gorm:"primary_key;autoIncrement" json:"id"`                // 规范类别 ID
	SupervisionTypeId int64     `gorm:"index" json:"supervisionTypeId"`                     // 监管类型 ID
	Name              string    `gorm:"size:100;not null" json:"name"`                      // 规范类别名称
	Code              string    `gorm:"size:50" json:"code"`                                // 规范类别代码
	SortOrder         int       `gorm:"default:0" json:"sortOrder"`                         // 排序
	IsEnabled         int       `gorm:"default:1" json:"isEnabled"`                         // 是否启用
	CreateBy          string    `gorm:"size:64;default:''" json:"createBy"`                 // 创建者
	CreateTime        time.Time `gorm:"autoCreateTime" json:"createTime"`                   // 创建时间
	UpdateBy          string    `gorm:"size:64;default:''" json:"updateBy"`                 // 更新者
	UpdateTime        time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`    // 更新时间
	Remark            string    `gorm:"size:500" json:"remark"`                             // 备注

	SupervisionType SysStandardPhraseSupervisionType `gorm:"-" json:"supervisionType"` // 关联的监管类型
	Items           []SysStandardPhraseItem          `gorm:"-" json:"items"`           // 关联的规范条目
}

// SysStandardPhraseItem 规范条目
type SysStandardPhraseItem struct {
	Id         int64     `gorm:"primary_key;autoIncrement" json:"id"`         // 规范条目 ID
	CategoryId int64     `gorm:"index" json:"categoryId"`                     // 规范类别 ID
	Title      string    `gorm:"size:200;not null" json:"title"`              // 条目标题
	Scene      string    `gorm:"size:200" json:"scene"`                       // 适用场景
	SortOrder  int       `gorm:"default:0" json:"sortOrder"`                  // 排序
	IsEnabled  int       `gorm:"default:1" json:"isEnabled"`                  // 是否启用
	CreateBy   string    `gorm:"size:64;default:''" json:"createBy"`          // 创建者
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`            // 创建时间
	UpdateBy   string    `gorm:"size:64;default:''" json:"updateBy"`          // 更新者
	UpdateTime time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark     string    `gorm:"size:500" json:"remark"`                      // 备注

	Category SysStandardPhraseCategory   `gorm:"-" json:"category"` // 关联的规范类别
	Contents []SysStandardPhraseContent  `gorm:"-" json:"contents"` // 关联的规范内容
}

// SysStandardPhraseContent 规范内容
type SysStandardPhraseContent struct {
	Id         int64     `gorm:"primary_key;autoIncrement" json:"id"`         // 规范内容 ID
	ItemId     int64     `gorm:"index" json:"itemId"`                         // 规范条目 ID
	Content    string    `gorm:"type:text;not null" json:"content"`           // 规范内容
	LegalBasis string    `gorm:"type:text" json:"legalBasis"`                 // 法律依据
	Tips       string    `gorm:"size:500" json:"tips"`                        // 提示要点
	SortOrder  int       `gorm:"default:0" json:"sortOrder"`                  // 排序
	CreateBy   string    `gorm:"size:64;default:''" json:"createBy"`          // 创建者
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`            // 创建时间
	UpdateBy   string    `gorm:"size:64;default:''" json:"updateBy"`          // 更新者
	UpdateTime time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark     string    `gorm:"size:500" json:"remark"`                      // 备注

	Item SysStandardPhraseItem `gorm:"-" json:"item"` // 关联的规范条目
}

// TableName 表名
func (SysStandardPhraseSupervisionType) TableName() string {
	return "law_standard_phrase_supervision_type"
}

func (SysStandardPhraseCategory) TableName() string {
	return "law_standard_phrase_category"
}

func (SysStandardPhraseItem) TableName() string {
	return "law_standard_phrase_item"
}

func (SysStandardPhraseContent) TableName() string {
	return "law_standard_phrase_content"
}

// ============ 监管类型操作方法 ============

// FindSupervisionTypeById 根据 ID 查询监管类型
func FindSupervisionTypeById(id int64) SysStandardPhraseSupervisionType {
	var item SysStandardPhraseSupervisionType
	mysql.MysqlDb().Where("id = ?", id).First(&item)
	return item
}

// FindSupervisionTypeByCode 根据代码查询监管类型
func FindSupervisionTypeByCode(code string) SysStandardPhraseSupervisionType {
	var item SysStandardPhraseSupervisionType
	mysql.MysqlDb().Where("code = ?", code).First(&item)
	return item
}

// SelectSupervisionTypeList 查询监管类型列表
func SelectSupervisionTypeList(param SysStandardPhraseSupervisionType) []SysStandardPhraseSupervisionType {
	db := mysql.MysqlDb().Model(&SysStandardPhraseSupervisionType{})

	if param.Name != "" {
		db = db.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}

	var result []SysStandardPhraseSupervisionType
	db.Order("sort_order").Find(&result)

	// 加载关联的类别
	for i := range result {
		result[i].Categories = SelectCategoryListBySupervisionTypeId(result[i].Id)
	}

	return result
}

// SaveSupervisionType 保存监管类型
func SaveSupervisionType(item SysStandardPhraseSupervisionType) string {
	if item.Id == 0 {
		mysql.MysqlDb().Create(&item)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&item)
	return "修改成功"
}

// DeleteSupervisionType 删除监管类型
func DeleteSupervisionType(ids []int64) string {
	mysql.MysqlDb().Delete(&SysStandardPhraseSupervisionType{}, "id IN ?", ids)
	return "删除成功"
}

// ============ 规范类别操作方法 ============

// FindCategoryById 根据 ID 查询规范类别
func FindCategoryById(id int64) SysStandardPhraseCategory {
	var item SysStandardPhraseCategory
	mysql.MysqlDb().Where("id = ?", id).First(&item)
	return item
}

// SelectCategoryListBySupervisionTypeId 根据监管类型 ID 查询类别列表
func SelectCategoryListBySupervisionTypeId(supervisionTypeId int64) []SysStandardPhraseCategory {
	var result []SysStandardPhraseCategory
	mysql.MysqlDb().Where("supervision_type_id = ?", supervisionTypeId).Order("sort_order").Find(&result)

	// 加载关联的条目
	for i := range result {
		result[i].Items = SelectItemListByCategoryId(result[i].Id)
	}

	return result
}

// SelectCategoryList 查询规范类别列表
func SelectCategoryList(param SysStandardPhraseCategory) []SysStandardPhraseCategory {
	db := mysql.MysqlDb().Model(&SysStandardPhraseCategory{})

	if param.SupervisionTypeId != 0 {
		db = db.Where("supervision_type_id = ?", param.SupervisionTypeId)
	}
	if param.Name != "" {
		db = db.Where("name LIKE ?", "%"+param.Name+"%")
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}

	var result []SysStandardPhraseCategory
	db.Order("sort_order").Find(&result)

	// 加载关联的条目
	for i := range result {
		result[i].Items = SelectItemListByCategoryId(result[i].Id)
	}

	return result
}

// SaveCategory 保存规范类别
func SaveCategory(item SysStandardPhraseCategory) string {
	if item.Id == 0 {
		mysql.MysqlDb().Create(&item)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&item)
	return "修改成功"
}

// DeleteCategory 删除规范类别
func DeleteCategory(ids []int64) string {
	mysql.MysqlDb().Delete(&SysStandardPhraseCategory{}, "id IN ?", ids)
	return "删除成功"
}

// ============ 规范条目操作方法 ============

// FindItemById 根据 ID 查询规范条目
func FindItemById(id int64) SysStandardPhraseItem {
	var item SysStandardPhraseItem
	mysql.MysqlDb().Where("id = ?", id).First(&item)
	return item
}

// SelectItemListByCategoryId 根据类别 ID 查询条目列表
func SelectItemListByCategoryId(categoryId int64) []SysStandardPhraseItem {
	var result []SysStandardPhraseItem
	mysql.MysqlDb().Where("category_id = ?", categoryId).Order("sort_order").Find(&result)

	// 加载关联的内容
	for i := range result {
		result[i].Contents = SelectContentListByItemId(result[i].Id)
	}

	return result
}

// SelectItemList 查询规范条目列表
func SelectItemList(param SysStandardPhraseItem) []SysStandardPhraseItem {
	db := mysql.MysqlDb().Model(&SysStandardPhraseItem{})

	if param.CategoryId != 0 {
		db = db.Where("category_id = ?", param.CategoryId)
	}
	if param.Title != "" {
		db = db.Where("title LIKE ?", "%"+param.Title+"%")
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}

	var result []SysStandardPhraseItem
	db.Order("sort_order").Find(&result)

	// 加载关联的内容
	for i := range result {
		result[i].Contents = SelectContentListByItemId(result[i].Id)
	}

	return result
}

// SaveItem 保存规范条目
func SaveItem(item SysStandardPhraseItem) string {
	if item.Id == 0 {
		mysql.MysqlDb().Create(&item)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&item)
	return "修改成功"
}

// DeleteItem 删除规范条目
func DeleteItem(ids []int64) string {
	mysql.MysqlDb().Delete(&SysStandardPhraseItem{}, "id IN ?", ids)
	return "删除成功"
}

// ============ 规范内容操作方法 ============

// FindContentById 根据 ID 查询规范内容
func FindContentById(id int64) SysStandardPhraseContent {
	var item SysStandardPhraseContent
	mysql.MysqlDb().Where("id = ?", id).First(&item)
	return item
}

// SelectContentListByItemId 根据条目 ID 查询内容列表
func SelectContentListByItemId(itemId int64) []SysStandardPhraseContent {
	var result []SysStandardPhraseContent
	mysql.MysqlDb().Where("item_id = ?", itemId).Order("sort_order").Find(&result)
	return result
}

// SelectContentList 查询规范内容列表
func SelectContentList(param SysStandardPhraseContent) []SysStandardPhraseContent {
	db := mysql.MysqlDb().Model(&SysStandardPhraseContent{})

	if param.ItemId != 0 {
		db = db.Where("item_id = ?", param.ItemId)
	}

	var result []SysStandardPhraseContent
	db.Order("sort_order").Find(&result)

	return result
}

// SaveContent 保存规范内容
func SaveContent(item SysStandardPhraseContent) string {
	if item.Id == 0 {
		mysql.MysqlDb().Create(&item)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&item)
	return "修改成功"
}

// DeleteContent 删除规范内容
func DeleteContent(ids []int64) string {
	mysql.MysqlDb().Delete(&SysStandardPhraseContent{}, "id IN ?", ids)
	return "删除成功"
}

// ============ 搜索操作方法 ============

// SearchStandardPhrase 搜索规范用语
func SearchStandardPhrase(keyword string) []SysStandardPhraseContent {
	var result []SysStandardPhraseContent
	mysql.MysqlDb().Where("content LIKE ? OR legal_basis LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&result)

	// 加载关联的条目
	for i := range result {
		result[i].Item = FindItemById(result[i].ItemId)
		if result[i].Item.Id != 0 {
			result[i].Item.Category = FindCategoryById(result[i].Item.CategoryId)
			if result[i].Item.Category.Id != 0 {
				result[i].Item.Category.SupervisionType = FindSupervisionTypeById(result[i].Item.Category.SupervisionTypeId)
			}
		}
	}

	return result
}
