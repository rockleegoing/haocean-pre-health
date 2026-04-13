package system

import (
	"haocean/health-enforcement/pkg/mysql"
	"time"
)

// SysRegulation 法律法规表
type SysRegulation struct {
	RegulationId   int64     `gorm:"primary_key;autoIncrement" json:"regulationId"`   // 法规 ID
	Title          string    `gorm:"size:200;notNull" json:"title"`                   // 法规标题
	LegalType      string    `gorm:"size:50" json:"legalType"`                        // 法律类型（法律/行政法规/部门规章/地方法规）
	SupervisionTypes string  `gorm:"type:json" json:"supervisionTypes"`               // 监管类型列表（JSON 数组）
	IndustryIds    string    `gorm:"type:json" json:"industryIds"`                    // 关联行业 ID 列表（JSON 数组）
	PublishOrg     string    `gorm:"size:100" json:"publishOrg"`                      // 发布机关
	PublishDate    *time.Time `gorm:"type:date" json:"publishDate"`                   // 发布日期
	EffectiveDate  *time.Time `gorm:"type:date" json:"effectiveDate"`                 // 生效日期
	Status         int       `gorm:"default:1" json:"status"`                         // 状态（0:废止/1:有效）
	Content        string    `gorm:"type:text" json:"content"`                        // 法规内容
	CreateBy       string    `gorm:"size:64;default:''" json:"createBy"`              // 创建者
	CreateTime     time.Time `gorm:"autoCreateTime" json:"createTime"`                // 创建时间
	UpdateBy       string    `gorm:"size:64;default:''" json:"updateBy"`              // 更新者
	UpdateTime     time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark         string    `gorm:"size:500" json:"remark"`                          // 备注

	// 关联字段
	Chapters []SysRegulationChapter `gorm:"-" json:"chapters"` // 章节列表
}

func (SysRegulation) TableName() string {
	return "law_regulation"
}

// SysRegulationChapter 法律法规章节表
type SysRegulationChapter struct {
	ChapterId    int64     `gorm:"primary_key;autoIncrement" json:"chapterId"`    // 章节 ID
	RegulationId int64     `gorm:"notNull" json:"regulationId"`                   // 法规 ID
	ChapterNo    int       `gorm:"default:0" json:"chapterNo"`                    // 章节序号
	ChapterTitle string    `gorm:"size:200" json:"chapterTitle"`                  // 章节标题
	ChapterType  string    `gorm:"size:20;default:'chapter'" json:"chapterType"`  // 章节类型
	ParentId     int64     `gorm:"default:0" json:"parentId"`                     // 父级章节 ID
	Level        int       `gorm:"default:1" json:"level"`                        // 层级
	Content      string    `gorm:"type:text" json:"content"`                      // 章节前言/概述内容
	SortOrder    int       `gorm:"default:0" json:"sortOrder"`                    // 排序
	CreateBy     string    `gorm:"size:64;default:''" json:"createBy"`            // 创建者
	CreateTime   time.Time `gorm:"autoCreateTime" json:"createTime"`              // 创建时间
	UpdateBy     string    `gorm:"size:64;default:''" json:"updateBy"`            // 更新者
	UpdateTime   time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"` // 更新时间
	Remark       string    `gorm:"size:500" json:"remark"`                        // 备注

	// 关联字段
	Articles []SysRegulationChapter `gorm:"-" json:"articles"` // 条款列表（复用结构）
}

func (SysRegulationChapter) TableName() string {
	return "law_regulation_chapter"
}

// SysRegulationArticle 法律法规条款表
type SysRegulationArticle struct {
	ArticleId        int64     `gorm:"primary_key;autoIncrement" json:"articleId"`        // 条款 ID
	RegulationId     int64     `gorm:"notNull" json:"regulationId"`                       // 法规 ID
	ChapterId        int64     `gorm:"default:0" json:"chapterId"`                        // 所属章节 ID
	ArticleNo        string    `gorm:"size:50" json:"articleNo"`                          // 条款编号
	ArticleNoSort    int       `gorm:"default:0" json:"articleNoSort"`                    // 条款序号
	Title            string    `gorm:"size:200" json:"title"`                             // 条款标题
	Content          string    `gorm:"type:text" json:"content"`                          // 条款内容
	PenaltyBasis     string    `gorm:"type:text" json:"penaltyBasis"`                     // 处罚依据
	PenaltyType      string    `gorm:"size:100" json:"penaltyType"`                       // 处罚种类
	DiscretionLevel  string    `gorm:"size:50" json:"discretionLevel"`                    // 裁量阶次
	ApplicableScenario string  `gorm:"type:text" json:"applicableScenario"`               // 适用情形
	PenaltyRange     string    `gorm:"size:200" json:"penaltyRange"`                      // 裁量幅度
	RemarkText       string    `gorm:"size:500" json:"remarkText"`                        // 备注
	SortOrder        int       `gorm:"default:0" json:"sortOrder"`                        // 排序
	CreateBy         string    `gorm:"size:64;default:''" json:"createBy"`                // 创建者
	CreateTime       time.Time `gorm:"autoCreateTime" json:"createTime"`                  // 创建时间
	UpdateBy         string    `gorm:"size:64;default:''" json:"updateBy"`                // 更新者
	UpdateTime       time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`   // 更新时间
	Remark           string    `gorm:"size:500" json:"remark"`                            // 备注
}

func (SysRegulationArticle) TableName() string {
	return "law_regulation_article"
}

// SysQualificationBasis 定性依据表
type SysQualificationBasis struct {
	BasisId            int64     `gorm:"primary_key;autoIncrement" json:"basisId"`            // 依据 ID
	RegulationId       int64     `gorm:"default:0" json:"regulationId"`                       // 关联法规 ID
	ArticleId          int64     `gorm:"default:0" json:"articleId"`                          // 关联条款 ID
	Title              string    `gorm:"size:200;notNull" json:"title"`                       // 依据标题
	Content            string    `gorm:"type:text" json:"content"`                            // 依据内容
	BasisType          string    `gorm:"size:50" json:"basisType"`                            // 依据类型
	LegalBasis         string    `gorm:"type:text" json:"legalBasis"`                         // 法律依据
	PenaltyBasis       string    `gorm:"type:text" json:"penaltyBasis"`                       // 处罚依据
	PenaltyType        string    `gorm:"size:100" json:"penaltyType"`                         // 处罚种类
	DiscretionLevel    string    `gorm:"size:50" json:"discretionLevel"`                      // 裁量阶次
	ApplicableScenario string    `gorm:"type:text" json:"applicableScenario"`                 // 适用情形
	PenaltyRange       string    `gorm:"size:200" json:"penaltyRange"`                        // 裁量幅度
	SortOrder          int       `gorm:"default:0" json:"sortOrder"`                          // 排序
	IsEnabled          int       `gorm:"default:1" json:"isEnabled"`                          // 是否启用
	CreateBy           string    `gorm:"size:64;default:''" json:"createBy"`                  // 创建者
	CreateTime         time.Time `gorm:"autoCreateTime" json:"createTime"`                    // 创建时间
	UpdateBy           string    `gorm:"size:64;default:''" json:"updateBy"`                  // 更新者
	UpdateTime         time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updateTime"`     // 更新时间
	Remark             string    `gorm:"size:500" json:"remark"`                              // 备注
}

func (SysQualificationBasis) TableName() string {
	return "law_qualification_basis"
}

// SysLegalTypeDict 法律类型字典表
type SysLegalTypeDict struct {
	TypeId    int64     `gorm:"primary_key;autoIncrement" json:"typeId"`       // 类型 ID
	TypeName  string    `gorm:"size:50;notNull" json:"typeName"`               // 类型名称
	TypeCode  string    `gorm:"size:50;notNull" json:"typeCode"`               // 类型代码
	SortOrder int       `gorm:"default:0" json:"sortOrder"`                    // 排序
	IsEnabled int       `gorm:"default:1" json:"isEnabled"`                    // 是否启用
	Remark    string    `gorm:"size:500" json:"remark"`                        // 备注
}

func (SysLegalTypeDict) TableName() string {
	return "law_legal_type_dict"
}

// SysSupervisionTypeDict 监管类型字典表
type SysSupervisionTypeDict struct {
	TypeId    int64     `gorm:"primary_key;autoIncrement" json:"typeId"`       // 类型 ID
	TypeName  string    `gorm:"size:100;notNull" json:"typeName"`              // 类型名称
	TypeCode  string    `gorm:"size:50;notNull" json:"typeCode"`               // 类型代码
	ParentId  int64     `gorm:"default:0" json:"parentId"`                     // 父级 ID
	Level     int       `gorm:"default:1" json:"level"`                        // 层级
	SortOrder int       `gorm:"default:0" json:"sortOrder"`                    // 排序
	IsEnabled int       `gorm:"default:1" json:"isEnabled"`                    // 是否启用
	Remark    string    `gorm:"size:500" json:"remark"`                        // 备注
}

func (SysSupervisionTypeDict) TableName() string {
	return "law_supervision_type_dict"
}

// ============================
// 法律法规数据库操作
// ============================

// SaveRegulation 保存法律法规
func SaveRegulation(regulation SysRegulation) string {
	if regulation.RegulationId == 0 {
		mysql.MysqlDb().Create(&regulation)
		return "添加成功"
	}
	mysql.MysqlDb().Save(&regulation)
	return "修改成功"
}

// DeleteRegulation 删除法律法规
func DeleteRegulation(ids []int64) string {
	mysql.MysqlDb().Delete(&SysRegulation{}, "regulation_id IN ?", ids)
	return "删除成功"
}

// FindRegulationById 根据 ID 查询法律法规
func FindRegulationById(id int64) SysRegulation {
	var regulation SysRegulation
	mysql.MysqlDb().Where("regulation_id = ?", id).First(&regulation)
	return regulation
}

// SelectRegulationList 查询法律法规列表
func SelectRegulationList(param SysRegulation, page int, pageSize int) ([]SysRegulation, int64) {
	db := mysql.MysqlDb().Model(&SysRegulation{})

	if param.Title != "" {
		db = db.Where("title LIKE ?", "%"+param.Title+"%")
	}
	if param.LegalType != "" {
		db = db.Where("legal_type = ?", param.LegalType)
	}
	if param.Status != -1 {
		db = db.Where("status = ?", param.Status)
	}

	var total int64
	db.Count(&total)

	var result []SysRegulation
	if page > 0 && pageSize > 0 {
		db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	db.Order("create_time DESC").Find(&result)

	return result, total
}

// SelectRegulationChapterList 查询章节列表
func SelectRegulationChapterList(regulationId int64) []SysRegulationChapter {
	var chapters []SysRegulationChapter
	mysql.MysqlDb().Where("regulation_id = ?", regulationId).Order("sort_order").Find(&chapters)
	return chapters
}

// SelectRegulationArticleList 查询条款列表
func SelectRegulationArticleList(regulationId int64, chapterId int64) []SysRegulationArticle {
	db := mysql.MysqlDb().Model(&SysRegulationArticle{}).Where("regulation_id = ?", regulationId)
	if chapterId > 0 {
		db = db.Where("chapter_id = ?", chapterId)
	}
	var articles []SysRegulationArticle
	db.Order("article_no_sort").Find(&articles)
	return articles
}

// SelectQualificationBasisList 查询定性依据列表
func SelectQualificationBasisList(regulationId int64) []SysQualificationBasis {
	var basisList []SysQualificationBasis
	mysql.MysqlDb().Where("regulation_id = ?", regulationId).Order("sort_order").Find(&basisList)
	return basisList
}

// SelectLegalTypeDictList 查询法律类型字典
func SelectLegalTypeDictList() []SysLegalTypeDict {
	var dictList []SysLegalTypeDict
	mysql.MysqlDb().Where("is_enabled = ?", 1).Order("sort_order").Find(&dictList)
	return dictList
}

// SelectSupervisionTypeDictList 查询监管类型字典
func SelectSupervisionTypeDictList() []SysSupervisionTypeDict {
	var dictList []SysSupervisionTypeDict
	mysql.MysqlDb().Where("is_enabled = ?", 1).Order("sort_order").Find(&dictList)
	return dictList
}
