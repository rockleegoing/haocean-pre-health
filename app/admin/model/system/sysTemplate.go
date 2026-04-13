package system

import (
	"encoding/json"
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysDocumentTemplate 文书模板
type SysDocumentTemplate struct {
	TemplateId    int64           `gorm:"primary_key;autoIncrement" json:"templateId"`    // 模板 ID
	TemplateName  string          `gorm:"size:100;not null" json:"templateName"`          // 模板名称
	CategoryId    int64           `gorm:"index" json:"categoryId"`                        // 分类 ID
	CategoryName  string          `gorm:"size:50" json:"categoryName"`                    // 分类名称（冗余）
	IndustryId    int64           `gorm:"index" json:"industryId"`                        // 行业分类 ID
	IndustryName  string          `gorm:"size:100" json:"industryName"`                   // 行业名称（冗余）
	TemplateType  string          `gorm:"size:20" json:"templateType"`                    // 模板类型（word/pdf）
	Fields        json.RawMessage `gorm:"type:json" json:"fields"`                        // 填空项定义
	FilePath      string          `gorm:"size:255" json:"filePath"`                       // 模板文件路径
	FileName      string          `gorm:"size:255" json:"fileName"`                       // 原始文件名
	FileSize      int64           `gorm:"default:0" json:"fileSize"`                      // 文件大小（字节）
	FileContent   string          `gorm:"type:longtext" json:"fileContent"`               // 模板内容（base64）
	Version       string          `gorm:"size:20;default:'1.0'" json:"version"`           // 版本号
	IsEnabled     int             `gorm:"default:1" json:"isEnabled"`                     // 是否启用
	SortOrder     int             `gorm:"default:0" json:"sortOrder"`                     // 排序
	CreateBy      string          `gorm:"size:64;default:''" json:"createBy"`             // 创建者
	CreateTime    time.Time       `gorm:"autoCreateTime" json:"createTime"`               // 创建时间
	UpdateBy      string          `gorm:"size:64;default:''" json:"updateBy"`             // 更新者
	UpdateTime    time.Time       `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark        string          `gorm:"size:500" json:"remark"`                         // 备注
}

func (SysDocumentTemplate) TableName() string {
	return "law_document_template"
}

// TemplateField 模板填空项
type TemplateField struct {
	Name     string   `json:"name"`     // 变量名
	Label    string   `json:"label"`    // 显示标签
	Type     string   `json:"type"`     // 类型：date/text/select/person
	Required bool     `json:"required"` // 是否必填
	Options  []string `json:"options"`  // 选项（下拉时用）
}

// FindTemplateById 根据 ID 查询模板
func FindTemplateById(id int64) SysDocumentTemplate {
	var template SysDocumentTemplate
	mysql.MysqlDb().Where("template_id = ?", id).First(&template)
	return template
}

// SaveTemplate 保存文书模板
func SaveTemplate(template SysDocumentTemplate) string {
	if template.TemplateId == 0 {
		mysql.MysqlDb().Create(&template)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&template)
	return "修改成功"
}

// DeleteTemplate 删除文书模板
func DeleteTemplate(ids []int64) string {
	mysql.MysqlDb().Delete(&SysDocumentTemplate{}, "template_id IN ?", ids)
	return "删除成功"
}

// SelectTemplateList 查询文书模板列表
func SelectTemplateList(param SearchTemplateParam) TableDataInfo {
	db := mysql.MysqlDb().Model(&SysDocumentTemplate{})

	if param.TemplateName != "" {
		db = db.Where("template_name LIKE ?", "%"+param.TemplateName+"%")
	}
	if param.CategoryId != 0 {
		db = db.Where("category_id = ?", param.CategoryId)
	}
	if param.IndustryId != 0 {
		db = db.Where("industry_id = ?", param.IndustryId)
	}
	if param.IsEnabled != -1 {
		db = db.Where("is_enabled = ?", param.IsEnabled)
	}

	var total int64
	db.Count(&total)

	var result []SysDocumentTemplate
	offset := (param.PageNum - 1) * param.PageSize
	db.Order("sort_order, create_time DESC").Offset(offset).Limit(param.PageSize).Find(&result)

	return TableDataInfo{
		Total: total,
		Rows:  result,
	}
}

// SearchTemplateParam 文书模板搜索参数
type SearchTemplateParam struct {
	PageNum      int    `form:"pageNum"`
	PageSize     int    `form:"pageSize"`
	TemplateName string `form:"templateName"`
	CategoryId   int64  `form:"categoryId"`
	IndustryId   int64  `form:"industryId"`
	IsEnabled    int    `form:"isEnabled"`
}

// FindTemplatesByIndustryId 根据行业 ID 查询模板列表
func FindTemplatesByIndustryId(industryId int64) []SysDocumentTemplate {
	var templates []SysDocumentTemplate
	mysql.MysqlDb().Where("industry_id = ? AND is_enabled = 1", industryId).
		Order("sort_order, create_time DESC").
		Find(&templates)
	return templates
}
