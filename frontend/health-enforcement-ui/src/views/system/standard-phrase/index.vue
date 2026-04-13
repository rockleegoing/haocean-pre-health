<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="监管类型" prop="supervisionTypeId">
        <el-select v-model="queryParams.supervisionTypeId" placeholder="请选择监管类型" clearable size="small" @change="handleSupervisionTypeChange">
          <el-option v-for="item in supervisionTypeOptions" :key="item.id" :label="item.name" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="类别名称" prop="categoryName">
        <el-input v-model="queryParams.categoryName" placeholder="请输入类别名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:standardPhrase:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:standardPhrase:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:standardPhrase:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="categoryList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="类别 ID" prop="id" width="80" />
      <el-table-column label="监管类型" prop="supervisionTypeName" width="150" />
      <el-table-column label="类别名称" prop="name" :show-overflow-tooltip="true" />
      <el-table-column label="类别代码" prop="code" width="120" />
      <el-table-column label="排序" prop="sortOrder" width="80" />
      <el-table-column label="启用状态" align="center" width="100">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.isEnabled === 1" type="success">启用</el-tag>
          <el-tag v-else type="info">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="200">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:standardPhrase:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:standardPhrase:remove']">删除</el-button>
          <el-button size="mini" type="text" icon="el-icon-list" @click="handleViewItems(scope.row)" v-hasPermi="['system:standardPhrase:view']">条目</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改规范类别对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="监管类型" prop="supervisionTypeId">
          <el-select v-model="form.supervisionTypeId" placeholder="请选择监管类型">
            <el-option v-for="item in supervisionTypeOptions" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="类别名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入类别名称" />
        </el-form-item>
        <el-form-item label="类别代码" prop="code">
          <el-input v-model="form.code" placeholder="请输入类别代码" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="form.sortOrder" :min="0" />
        </el-form-item>
        <el-form-item label="启用状态">
          <el-radio-group v-model="form.isEnabled">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 规范条目对话框 -->
    <el-dialog title="规范条目" :visible.sync="itemDialogOpen" width="800px" append-to-body>
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAddItem" v-hasPermi="['system:standardPhrase:add']">新增</el-button>
        </el-col>
      </el-row>
      <el-table :data="itemList" style="width: 100%">
        <el-table-column label="条目 ID" prop="id" width="80" />
        <el-table-column label="条目标题" prop="title" :show-overflow-tooltip="true" />
        <el-table-column label="适用场景" prop="scene" width="150" />
        <el-table-column label="排序" prop="sortOrder" width="80" />
        <el-table-column label="操作" align="center" width="160">
          <template slot-scope="scope">
            <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdateItem(scope.row)">修改</el-button>
            <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDeleteItem(scope.row)">删除</el-button>
            <el-button size="mini" type="text" icon="el-icon-document" @click="handleViewContent(scope.row)">内容</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 添加或修改规范条目对话框 -->
    <el-dialog :title="itemTitle" :visible.sync="itemOpen" width="600px" append-to-body>
      <el-form ref="itemForm" :model="itemForm" :rules="itemRules" label-width="100px">
        <el-form-item label="条目标题" prop="title">
          <el-input v-model="itemForm.title" placeholder="请输入条目标题" />
        </el-form-item>
        <el-form-item label="适用场景" prop="scene">
          <el-input v-model="itemForm.scene" placeholder="请输入适用场景" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="itemForm.sortOrder" :min="0" />
        </el-form-item>
        <el-form-item label="启用状态">
          <el-radio-group v-model="itemForm.isEnabled">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="0">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitItemForm">确 定</el-button>
        <el-button @click="cancelItem">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 规范内容对话框 -->
    <el-dialog title="规范内容" :visible.sync="contentDialogOpen" width="800px" append-to-body>
      <el-row :gutter="10" class="mb8">
        <el-col :span="1.5">
          <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAddContent" v-hasPermi="['system:standardPhrase:add']">新增</el-button>
        </el-col>
      </el-row>
      <el-table :data="contentList" style="width: 100%">
        <el-table-column label="内容 ID" prop="id" width="80" />
        <el-table-column label="规范内容" prop="content" :show-overflow-tooltip="true" show-overflow-tooltip />
        <el-table-column label="法律依据" prop="legalBasis" width="200" :show-overflow-tooltip="true" />
        <el-table-column label="操作" align="center" width="160">
          <template slot-scope="scope">
            <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdateContent(scope.row)">修改</el-button>
            <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDeleteContent(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 添加或修改规范内容对话框 -->
    <el-dialog :title="contentTitle" :visible.sync="contentOpen" width="700px" append-to-body>
      <el-form ref="contentForm" :model="contentForm" :rules="contentRules" label-width="100px">
        <el-form-item label="规范内容" prop="content">
          <el-input v-model="contentForm.content" type="textarea" :rows="6" placeholder="请输入规范内容" />
        </el-form-item>
        <el-form-item label="法律依据" prop="legalBasis">
          <el-input v-model="contentForm.legalBasis" type="textarea" :rows="3" placeholder="请输入法律依据" />
        </el-form-item>
        <el-form-item label="提示要点" prop="tips">
          <el-input v-model="contentForm.tips" type="textarea" :rows="3" placeholder="请输入提示要点" />
        </el-form-item>
        <el-form-item label="排序" prop="sortOrder">
          <el-input-number v-model="contentForm.sortOrder" :min="0" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitContentForm">确 定</el-button>
        <el-button @click="cancelContent">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  listSupervisionType, getSupervisionType, addSupervisionType, updateSupervisionType, deleteSupervisionType,
  listCategory, getCategory, addCategory, updateCategory, deleteCategory,
  listItem, getItem, addItem, updateItem, deleteItem,
  listContent, getContent, addContent, updateContent, deleteContent
} from "@/api/system/standardPhrase"

export default {
  name: "StandardPhrase",
  dicts: ['sys_normal_disable'],
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      categoryList: [],
      supervisionTypeOptions: [],
      open: false,
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        supervisionTypeId: undefined,
        categoryName: undefined
      },
      rules: {
        supervisionTypeId: [{ required: true, message: "监管类型不能为空", trigger: "change" }],
        name: [{ required: true, message: "类别名称不能为空", trigger: "blur" }]
      },
      // 条目相关
      itemDialogOpen: false,
      itemOpen: false,
      itemTitle: "",
      itemList: [],
      itemForm: {},
      currentItem: null,
      itemRules: {
        title: [{ required: true, message: "条目标题不能为空", trigger: "blur" }]
      },
      // 内容相关
      contentDialogOpen: false,
      contentOpen: false,
      contentTitle: "",
      contentList: [],
      contentForm: {},
      currentContentItem: null,
      contentRules: {
        content: [{ required: true, message: "规范内容不能为空", trigger: "blur" }]
      }
    }
  },
  created() {
    this.getSupervisionTypeList()
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listCategory(this.queryParams).then(response => {
        this.categoryList = response.rows || response
        this.total = response.total || this.categoryList.length
        this.loading = false
      })
    },
    getSupervisionTypeList() {
      listSupervisionType({ isEnabled: 1 }).then(response => {
        this.supervisionTypeOptions = response.rows || response
      })
    },
    handleSupervisionTypeChange(value) {
      this.queryParams.supervisionTypeId = value
      this.getList()
    },
    resetQuery() {
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加规范类别"
    },
    handleUpdate(row) {
      this.reset()
      const id = row.id || this.ids[0]
      getCategory(id).then(response => {
        this.form = response.data
        this.open = true
        this.title = "修改规范类别"
      })
    },
    handleViewItems(row) {
      this.currentItem = row
      this.itemDialogOpen = true
      this.getItemList(row.id)
    },
    // 条目相关方法
    getItemList(categoryId) {
      listItem({ categoryId: categoryId }).then(response => {
        this.itemList = response.rows || response
      })
    },
    handleAddItem() {
      this.resetItem()
      this.itemForm.categoryId = this.currentItem.id
      this.itemOpen = true
      this.itemTitle = "添加规范条目"
    },
    handleUpdateItem(row) {
      this.resetItem()
      getItem(row.id).then(response => {
        this.itemForm = response.data
        this.itemOpen = true
        this.itemTitle = "修改规范条目"
      })
    },
    handleDeleteItem(row) {
      this.$confirm('是否确认删除条目标题为"' + row.title + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteItem(row.id)
      }).then(() => {
        this.getItemList(this.currentItem.id)
        this.msgSuccess("删除成功")
      })
    },
    handleViewContent(row) {
      this.currentContentItem = row
      this.contentDialogOpen = true
      this.getContentList(row.id)
    },
    // 内容相关方法
    getContentList(itemId) {
      listContent({ itemId: itemId }).then(response => {
        this.contentList = response.rows || response
      })
    },
    handleAddContent() {
      this.resetContent()
      this.contentForm.itemId = this.currentContentItem.id
      this.contentOpen = true
      this.contentTitle = "添加规范内容"
    },
    handleUpdateContent(row) {
      this.resetContent()
      getContent(row.id).then(response => {
        this.contentForm = response.data
        this.contentOpen = true
        this.contentTitle = "修改规范内容"
      })
    },
    handleDeleteContent(row) {
      this.$confirm('是否确认删除该规范内容？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteContent(row.id)
      }).then(() => {
        this.getContentList(this.currentContentItem.id)
        this.msgSuccess("删除成功")
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateCategory(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addCategory(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    submitItemForm() {
      this.$refs["itemForm"].validate(valid => {
        if (valid) {
          if (this.itemForm.id != undefined) {
            updateItem(this.itemForm).then(response => {
              this.msgSuccess("修改成功")
              this.itemOpen = false
              this.getItemList(this.currentItem.id)
            })
          } else {
            addItem(this.itemForm).then(response => {
              this.msgSuccess("新增成功")
              this.itemOpen = false
              this.getItemList(this.currentItem.id)
            })
          }
        }
      })
    },
    submitContentForm() {
      this.$refs["contentForm"].validate(valid => {
        if (valid) {
          if (this.contentForm.id != undefined) {
            updateContent(this.contentForm).then(response => {
              this.msgSuccess("修改成功")
              this.contentOpen = false
              this.getContentList(this.currentContentItem.id)
            })
          } else {
            addContent(this.contentForm).then(response => {
              this.msgSuccess("新增成功")
              this.contentOpen = false
              this.getContentList(this.currentContentItem.id)
            })
          }
        }
      })
    },
    handleDelete(row) {
      const ids = row.id || this.ids
      this.$confirm('是否确认删除类别名称为"' + row.name + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteCategory(ids)
      }).then(() => {
        this.getList()
        this.msgSuccess("删除成功")
      })
    },
    cancel() {
      this.open = false
      this.reset()
    },
    reset() {
      this.form = {
        id: undefined,
        supervisionTypeId: undefined,
        name: undefined,
        code: undefined,
        sortOrder: 0,
        isEnabled: 1
      }
      this.resetForm("form")
    },
    cancelItem() {
      this.itemOpen = false
      this.resetItem()
    },
    resetItem() {
      this.itemForm = {
        id: undefined,
        categoryId: undefined,
        title: undefined,
        scene: undefined,
        sortOrder: 0,
        isEnabled: 1
      }
      this.resetForm("itemForm")
    },
    cancelContent() {
      this.contentOpen = false
      this.resetContent()
    },
    resetContent() {
      this.contentForm = {
        id: undefined,
        itemId: undefined,
        content: undefined,
        legalBasis: undefined,
        tips: undefined,
        sortOrder: 0
      }
      this.resetForm("contentForm")
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
