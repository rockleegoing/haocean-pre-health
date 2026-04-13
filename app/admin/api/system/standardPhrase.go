package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ============ 监管类型 API ============

// ListSupervisionType 查询监管类型列表
func ListSupervisionType(c *gin.Context) {
	var param system.SysStandardPhraseSupervisionType
	param.Name = c.Query("name")
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))

	result := system.SelectSupervisionTypeList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetSupervisionType 获取监管类型详情
func GetSupervisionType(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := system.FindSupervisionTypeById(id)
	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("监管类型不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(item))
}

// AddSupervisionType 添加监管类型
func AddSupervisionType(c *gin.Context) {
	var item system.SysStandardPhraseSupervisionType
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Name == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("监管类型名称不能为空"))
		return
	}

	// 设置默认值
	if item.SortOrder == 0 {
		item.SortOrder = 0
	}
	if item.IsEnabled == 0 {
		item.IsEnabled = 1
	}

	msg := system.SaveSupervisionType(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateSupervisionType 修改监管类型
func UpdateSupervisionType(c *gin.Context) {
	var item system.SysStandardPhraseSupervisionType
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("监管类型 ID 不能为空"))
		return
	}

	msg := system.SaveSupervisionType(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteSupervisionType 删除监管类型
func DeleteSupervisionType(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteSupervisionType(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// ============ 规范类别 API ============

// ListCategory 查询规范类别列表
func ListCategory(c *gin.Context) {
	var param system.SysStandardPhraseCategory
	param.SupervisionTypeId, _ = strconv.ParseInt(c.DefaultQuery("supervisionTypeId", "0"), 10, 64)
	param.Name = c.Query("name")
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))

	result := system.SelectCategoryList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetCategory 获取规范类别详情
func GetCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := system.FindCategoryById(id)
	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范类别不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(item))
}

// AddCategory 添加规范类别
func AddCategory(c *gin.Context) {
	var item system.SysStandardPhraseCategory
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Name == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范类别名称不能为空"))
		return
	}

	if item.SupervisionTypeId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("监管类型 ID 不能为空"))
		return
	}

	msg := system.SaveCategory(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateCategory 修改规范类别
func UpdateCategory(c *gin.Context) {
	var item system.SysStandardPhraseCategory
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范类别 ID 不能为空"))
		return
	}

	msg := system.SaveCategory(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteCategory 删除规范类别
func DeleteCategory(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteCategory(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// ============ 规范条目 API ============

// ListItem 查询规范条目列表
func ListItem(c *gin.Context) {
	var param system.SysStandardPhraseItem
	param.CategoryId, _ = strconv.ParseInt(c.DefaultQuery("categoryId", "0"), 10, 64)
	param.Title = c.Query("title")
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))

	result := system.SelectItemList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetItem 获取规范条目详情
func GetItem(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := system.FindItemById(id)
	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范条目不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(item))
}

// AddItem 添加规范条目
func AddItem(c *gin.Context) {
	var item system.SysStandardPhraseItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Title == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("条目标题不能为空"))
		return
	}

	if item.CategoryId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范类别 ID 不能为空"))
		return
	}

	msg := system.SaveItem(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateItem 修改规范条目
func UpdateItem(c *gin.Context) {
	var item system.SysStandardPhraseItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范条目 ID 不能为空"))
		return
	}

	msg := system.SaveItem(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteItem 删除规范条目
func DeleteItem(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteItem(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// ============ 规范内容 API ============

// ListContent 查询规范内容列表
func ListContent(c *gin.Context) {
	var param system.SysStandardPhraseContent
	param.ItemId, _ = strconv.ParseInt(c.DefaultQuery("itemId", "0"), 10, 64)

	result := system.SelectContentList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetContent 获取规范内容详情
func GetContent(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := system.FindContentById(id)
	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范内容不存在"))
		return
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(item))
}

// AddContent 添加规范内容
func AddContent(c *gin.Context) {
	var item system.SysStandardPhraseContent
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Content == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范内容不能为空"))
		return
	}

	if item.ItemId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范条目 ID 不能为空"))
		return
	}

	msg := system.SaveContent(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateContent 修改规范内容
func UpdateContent(c *gin.Context) {
	var item system.SysStandardPhraseContent
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.Id == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("规范内容 ID 不能为空"))
		return
	}

	msg := system.SaveContent(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteContent 删除规范内容
func DeleteContent(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	var ids []int64
	for _, idStr := range strings.Split(idsStr, ",") {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	msg := system.DeleteContent(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// ============ 搜索 API ============

// SearchStandardPhrase 搜索规范用语
func SearchStandardPhrase(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("搜索关键词不能为空"))
		return
	}

	result := system.SearchStandardPhrase(keyword)
	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}

// GetFullTree 获取完整树形结构（监管类型 -> 类别 -> 条目 -> 内容）
func GetFullTree(c *gin.Context) {
	supervisionTypeIdStr := c.DefaultQuery("supervisionTypeId", "0")
	supervisionTypeId, _ := strconv.ParseInt(supervisionTypeIdStr, 10, 64)

	var result []system.SysStandardPhraseSupervisionType
	if supervisionTypeId > 0 {
		// 查询指定监管类型
		item := system.FindSupervisionTypeById(supervisionTypeId)
		if item.Id != 0 {
			item.Categories = system.SelectCategoryListBySupervisionTypeId(item.Id)
			result = append(result, item)
		}
	} else {
		// 查询全部
		result = system.SelectSupervisionTypeList(system.SysStandardPhraseSupervisionType{IsEnabled: 1})
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(result))
}
