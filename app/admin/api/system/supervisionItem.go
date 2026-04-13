package system

import (
	"github.com/gin-gonic/gin"
	"haocean/health-enforcement/app/admin/model/system"
	"haocean/health-enforcement/app/core/utils/R"
	"net/http"
	"strconv"
	"strings"
)

// ListSupervisionItem 查询监管事项列表
func ListSupervisionItem(c *gin.Context) {
	var param system.SearchSupervisionItemParam
	param.PageNum, _ = strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	param.PageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	param.ItemName = c.Query("itemName")
	param.SupervisionType = c.Query("supervisionType")
	param.IsEnabled, _ = strconv.Atoi(c.DefaultQuery("isEnabled", "-1"))
	param.ParentId, _ = strconv.ParseInt(c.DefaultQuery("parentId", "-1"), 10, 64)

	result := system.SelectSupervisionPageList(param)
	result.Code = http.StatusOK
	result.Msg = "查询成功"

	c.JSON(http.StatusOK, result)
}

// GetSupervisionItem 获取监管事项详情
func GetSupervisionItem(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	item := system.FindSupervisionItemById(id)
	if item.ItemId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("监管事项不存在"))
		return
	}

	// 查询子事项
	item.Children = system.FindSupervisionChildren(item.ItemId)

	c.JSON(http.StatusOK, R.ReturnSuccess(item))
}

// AddSupervisionItem 添加监管事项
func AddSupervisionItem(c *gin.Context) {
	var item system.SysSupervisionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.ItemName == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("事项名称不能为空"))
		return
	}

	// 设置默认值
	if item.Level == 0 {
		item.Level = 1
	}
	if item.IsEnabled == 0 {
		item.IsEnabled = 1
	}

	msg := system.SaveSupervisionItem(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// UpdateSupervisionItem 修改监管事项
func UpdateSupervisionItem(c *gin.Context) {
	var item system.SysSupervisionItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	if item.ItemId == 0 {
		c.JSON(http.StatusOK, R.ReturnFailMsg("事项 ID 不能为空"))
		return
	}

	msg := system.SaveSupervisionItem(item)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// DeleteSupervisionItem 删除监管事项
func DeleteSupervisionItem(c *gin.Context) {
	idsStr := c.Param("ids")
	if idsStr == "" {
		c.JSON(http.StatusOK, R.ReturnFailMsg("参数错误"))
		return
	}

	// 将逗号分隔的 ID 转换为数组
	idStrings := strings.Split(idsStr, ",")
	var ids []int64
	for _, idStr := range idStrings {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}

	// 检查是否有子事项
	for _, id := range ids {
		children := system.FindSupervisionChildren(id)
		if len(children) > 0 {
			c.JSON(http.StatusOK, R.ReturnFailMsg("请先删除子事项"))
			return
		}
	}

	msg := system.DeleteSupervisionItem(ids)
	c.JSON(http.StatusOK, R.ReturnSuccess(msg))
}

// GetSupervisionTree 获取监管事项树
func GetSupervisionTree(c *gin.Context) {
	var param system.SearchSupervisionItemParam
	param.ParentId = 0
	param.IsEnabled = 1
	param.Tree = true

	items := system.SelectSupervisionItemList(param)
	c.JSON(http.StatusOK, R.ReturnSuccess(items))
}

// GetSupervisionChildren 获取监管事项子项
func GetSupervisionChildren(c *gin.Context) {
	idStr := c.Param("parentId")
	parentId, _ := strconv.ParseInt(idStr, 10, 64)

	children := system.FindSupervisionChildren(parentId)
	c.JSON(http.StatusOK, R.ReturnSuccess(children))
}

// ListSupervisionCategory 查询监管事项分类列表
func ListSupervisionCategory(c *gin.Context) {
	// 监管类型列表（固定值）
	categories := []map[string]interface{}{
		{"categoryId": 1, "categoryName": "食品安全", "categoryCode": "SPAQ", "icon": "icon-food"},
		{"categoryId": 2, "categoryName": "公共场所", "categoryCode": "GGCS", "icon": "icon-public"},
		{"categoryId": 3, "categoryName": "消毒产品", "categoryCode": "XDCP", "icon": "icon-disinfect"},
		{"categoryId": 4, "categoryName": "生活饮用水", "categoryCode": "SHYSY", "icon": "icon-water"},
		{"categoryId": 5, "categoryName": "放射卫生", "categoryCode": "FSWS", "icon": "icon-radiation"},
		{"categoryId": 6, "categoryName": "职业卫生", "categoryCode": "ZYWS", "icon": "icon-occupation"},
		{"categoryId": 7, "categoryName": "医疗机构", "categoryCode": "YLJG", "icon": "icon-medical"},
		{"categoryId": 8, "categoryName": "学校卫生", "categoryCode": "XXWS", "icon": "icon-school"},
		{"categoryId": 9, "categoryName": "传染病防治", "categoryCode": "CRBFZ", "icon": "icon-infection"},
		{"categoryId": 10, "categoryName": "妇幼健康", "categoryCode": "FYJK", "icon": "icon-maternal"},
	}

	c.JSON(http.StatusOK, R.ReturnSuccess(categories))
}
