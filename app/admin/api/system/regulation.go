package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"haocean/health-enforcement/pkg/mysql"
	"net/http"
	"strconv"
	"strings"
)

// ListRegulation 查询法律法规列表
func ListRegulation(c *gin.Context) {
	var param system.SysRegulation
	param.Title = c.Query("title")
	param.LegalType = c.Query("legalType")
	statusStr := c.DefaultQuery("status", "-1")
	status, _ := strconv.Atoi(statusStr)
	param.Status = status

	page, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	regulations, total := system.SelectRegulationList(param, page, pageSize)
	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"rows": regulations,
		"total": total,
		"pageNum": page,
		"pageSize": pageSize,
	}))
}

// GetRegulation 获取法律法规详情
func GetRegulation(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	regulation := system.FindRegulationById(id)
	if regulation.RegulationId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("法律法规不存在"))
		return
	}

	// 查询关联的章节和条款
	regulation.Chapters = system.SelectRegulationChapterList(regulation.RegulationId)

	c.JSON(http.StatusOK, R.ReturnSuccess(regulation))
}

// AddRegulation 添加法律法规
func AddRegulation(c *gin.Context) {
	var regulation system.SysRegulation
	if err := c.ShouldBindJSON(&regulation); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if regulation.Title == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("法规标题不能为空"))
		return
	}

	msg := system.SaveRegulation(regulation)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateRegulation 修改法律法规
func UpdateRegulation(c *gin.Context) {
	var regulation system.SysRegulation
	if err := c.ShouldBindJSON(&regulation); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if regulation.RegulationId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("法规 ID 不能为空"))
		return
	}

	msg := system.SaveRegulation(regulation)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteRegulation 删除法律法规
func DeleteRegulation(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	idStrings := strings.Split(idsStr, ",")
	var ids []int64
	for _, idStr := range idStrings {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteRegulation(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// GetRegulationHome 获取法律法规首页数据（综合法律条例 + 监管类型）
func GetRegulationHome(c *gin.Context) {
	// 查询法律类型
	legalTypes := system.SelectLegalTypeDictList()

	// 查询监管类型
	supervisionTypes := system.SelectSupervisionTypeDictList()

	// 查询各类型的法律法规数量
	type CountInfo struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	var legalTypeCounts []CountInfo
	for _, lt := range legalTypes {
		var count int64
		mysql.MysqlDb().Model(&system.SysRegulation{}).Where("legal_type = ? AND status = 1", lt.TypeCode).Count(&count)
		legalTypeCounts = append(legalTypeCounts, CountInfo{
			Type:  lt.TypeCode,
			Name:  lt.TypeName,
			Count: int(count),
		})
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"legalTypes":       legalTypes,
		"supervisionTypes": supervisionTypes,
		"legalTypeCounts":  legalTypeCounts,
	}))
}

// GetRegulationBookList 获取法律法规书本列表（按法律类型筛选）
func GetRegulationBookList(c *gin.Context) {
	legalType := c.Query("legalType")

	var regulations []system.SysRegulation
	db := mysql.MysqlDb().Model(&system.SysRegulation{}).Where("status = 1")

	if legalType != "" {
		db = db.Where("legal_type = ?", legalType)
	}

	db.Order("create_time DESC").Find(&regulations)

	c.JSON(http.StatusOK, R.ReturnSuccess(regulations))
}

// GetRegulationBookDetail 获取法律法规书本详情（目录 + 章节内容）
func GetRegulationBookDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	regulation := system.FindRegulationById(id)
	if regulation.RegulationId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("法律法规不存在"))
		return
	}

	// 查询章节列表
	chapters := system.SelectRegulationChapterList(regulation.RegulationId)

	// 查询条款列表
	articles := system.SelectRegulationArticleList(regulation.RegulationId, 0)

	// 查询定性依据列表
	basisList := system.SelectQualificationBasisList(regulation.RegulationId)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"regulation": regulation,
		"chapters":   chapters,
		"articles":   articles,
		"basisList":  basisList,
	}))
}

// GetRegulationChapterContent 获取章节详细内容
func GetRegulationChapterContent(c *gin.Context) {
	chapterIdStr := c.Param("chapterId")
	chapterId, _ := strconv.ParseInt(chapterIdStr, 10, 64)

	var chapter system.SysRegulationChapter
	mysql.MysqlDb().Where("chapter_id = ?", chapterId).First(&chapter)

	if chapter.ChapterId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("章节不存在"))
		return
	}

	// 查询该章节下的条款
	articles := system.SelectRegulationArticleList(chapter.RegulationId, chapterId)

	c.JSON(http.StatusOK, R.ReturnSuccess(gin.H{
		"chapter":  chapter,
		"articles": articles,
	}))
}

// SearchRegulation 搜索法律法规
func SearchRegulation(c *gin.Context) {
	keyword := c.Query("keyword")
	legalType := c.Query("legalType")

	var regulations []system.SysRegulation
	db := mysql.MysqlDb().Model(&system.SysRegulation{}).Where("status = 1")

	if keyword != "" {
		db = db.Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if legalType != "" {
		db = db.Where("legal_type = ?", legalType)
	}

	db.Order("create_time DESC").Find(&regulations)

	c.JSON(http.StatusOK, R.ReturnSuccess(regulations))
}

// GetLegalTypeList 获取法律类型列表
func GetLegalTypeList(c *gin.Context) {
	legalTypes := system.SelectLegalTypeDictList()
	c.JSON(http.StatusOK, R.ReturnSuccess(legalTypes))
}

// GetSupervisionTypeList 获取监管类型列表
func GetSupervisionTypeList(c *gin.Context) {
	supervisionTypes := system.SelectSupervisionTypeDictList()
	c.JSON(http.StatusOK, R.ReturnSuccess(supervisionTypes))
}

// GetBasisList 获取定性依据列表
func GetBasisList(c *gin.Context) {
	regulationIdStr := c.Query("regulationId")
	regulationId, _ := strconv.ParseInt(regulationIdStr, 10, 64)

	basisList := system.SelectQualificationBasisList(regulationId)
	c.JSON(http.StatusOK, R.ReturnSuccess(basisList))
}

// GetBasisDetail 获取定性依据详情
func GetBasisDetail(c *gin.Context) {
	basisIdStr := c.Param("id")
	basisId, _ := strconv.ParseInt(basisIdStr, 10, 64)

	var basis system.SysQualificationBasis
	mysql.MysqlDb().Where("basis_id = ?", basisId).First(&basis)

	if basis.BasisId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("定性依据不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(basis))
}
