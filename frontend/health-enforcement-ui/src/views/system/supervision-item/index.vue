<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="事项名称" prop="itemName">
        <el-input v-model="queryParams.itemName" placeholder="请输入事项名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="监管类型" prop="supervisionType">
        <el-select v-model="queryParams.supervisionType" placeholder="请选择监管类型" clearable size="small">
          <el-option v-for="item in categoryOptions" :key="item.categoryCode" :label="item.categoryName" :value="item.categoryCode" />
        </el-select>
      </el-form-item>
      <el-form-item label="启用状态" prop="isEnabled">
        <el-select v-model="queryParams.isEnabled" placeholder="请选择启用状态" clearable size="small">
          <el-option v-for="dict in dict.type.sys_normal_disable" :key="dict.value" :label="dict.label" :value="dict.value" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:supervisionItem:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:supervisionItem:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:supervisionItem:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="itemList" row-key="itemId" :tree-props="{children: 'children', hasChildren: 'hasChildren'}" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="事项 ID" prop="itemId" width="80" />
      <el-table-column label="事项名称" prop="itemName" :show-overflow-tooltip="true" />
      <el-table-column label="监管类型" prop="supervisionType" width="120">
        <template slot-scope="scope">
          <span>{{ getSupervisionTypeName(scope.row.supervisionType) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="级别" prop="level" width="80">
        <template slot-scope="scope">
          <span v-if="scope.row.level === 1">一级</span>
          <span v-else-if="scope.row.level === 2">二级</span>
          <span v-else>三级</span>
        </template>
      </el-table-column>
      <el-table-column label="排序" prop="sortOrder" width="80" />
      <el-table-column label="启用状态" align="center" width="100">
        <template slot-scope="scope">
          <el-switch v-model="scope.row.isEnabled" active-value="1" inactive-value="0" @change="handleStatusChange(scope.row)"></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="160">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:supervisionItem:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:supervisionItem:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 添加或修改监管事项对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="120px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="上级事项" prop="parentId">
              <el-tree-select v-model="form.parentId" :data="itemTreeOptions" :normal-label="itemName" :expand-on-click-node="false" placeholder="选择上级事项" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="事项名称" prop="itemName">
              <el-input v-model="form.itemName" placeholder="请输入事项名称" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="监管类型" prop="supervisionType">
              <el-select v-model="form.supervisionType" placeholder="请选择监管类型">
                <el-option v-for="item in categoryOptions" :key="item.categoryCode" :label="item.categoryName" :value="item.categoryCode" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="级别" prop="level">
              <el-select v-model="form.level" placeholder="请选择级别">
                <el-option label="一级" :value="1" />
                <el-option label="二级" :value="2" />
                <el-option label="三级" :value="3" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="排序" prop="sortOrder">
              <el-input-number v-model="form.sortOrder" :min="0" placeholder="请输入排序" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="启用状态">
              <el-radio-group v-model="form.isEnabled">
                <el-radio :label="1">启用</el-radio>
                <el-radio :label="0">禁用</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="关联行业" prop="industryIds">
          <el-tree-select v-model="form.industryIds" :data="industryTreeOptions" :normal-label="industryName" :expand-on-click-node="false" placeholder="选择关联行业" multiple />
        </el-form-item>
        <el-form-item label="关联规范用语" prop="standardLanguageIds">
          <el-select v-model="form.standardLanguageIds" placeholder="请选择关联规范用语" multiple filterable>
            <el-option v-for="item in languageOptions" :key="item.id" :label="item.title" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="检查要点" prop="checkPoints">
          <el-input v-model="form.checkPoints" type="textarea" :rows="4" placeholder="请输入检查要点" />
        </el-form-item>
        <el-form-item label="法律依据" prop="legalBasis">
          <el-input v-model="form.legalBasis" type="textarea" :rows="4" placeholder="请输入法律依据" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listSupervisionItem, getSupervisionItem, addSupervisionItem, updateSupervisionItem, deleteSupervisionItem, getSupervisionTree, listSupervisionCategory } from "@/api/system/supervisionItem"
import { listIndustry } from "@/api/system/industry"
import Treeselect from "@riophae/vue-treeselect"
import "@riophae/vue-treeselect/dist/vue-treeselect.css"

export default {
  name: "SupervisionItem",
  dicts: ['sys_normal_disable'],
  components: { Treeselect },
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      itemList: [],
      itemTreeOptions: [],
      industryTreeOptions: [],
      categoryOptions: [],
      languageOptions: [],
      open: false,
      title: "",
      form: {},
      queryParams: {
        itemName: undefined,
        supervisionType: undefined,
        isEnabled: undefined
      },
      rules: {
        itemName: [{ required: true, message: "事项名称不能为空", trigger: "blur" }],
        level: [{ required: true, message: "级别不能为空", trigger: "blur" }]
      }
    }
  },
  created() {
    this.getList()
    this.getCategoryList()
  },
  methods: {
    getList() {
      this.loading = true
      listSupervisionItem(this.queryParams).then(response => {
        this.itemList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    getCategoryList() {
      listSupervisionCategory().then(response => {
        this.categoryOptions = response.data
      })
    },
    getItemTree() {
      getSupervisionTree().then(response => {
        this.itemTreeOptions = this.buildTree(response.data)
      })
    },
    getIndustryTree() {
      listIndustry({ isEnabled: 1 }).then(response => {
        this.industryTreeOptions = this.buildTree(response.rows)
      })
    },
    buildTree(data, parentId = 0) {
      const result = []
      data.forEach(item => {
        if (item.parentId === parentId) {
          const node = {
            id: item.itemId || item.industryId,
            label: item.itemName || item.industryName,
            children: this.buildTree(data, item.itemId || item.industryId)
          }
          result.push(node)
        }
      })
      return result
    },
    getSupervisionTypeName(code) {
      const category = this.categoryOptions.find(item => item.categoryCode === code)
      return category ? category.categoryName : code
    },
    resetQuery() {
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.itemId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.getItemTree()
      this.getIndustryTree()
      this.open = true
      this.title = "添加监管事项"
    },
    handleUpdate(row) {
      this.reset()
      this.getItemTree()
      this.getIndustryTree()
      const itemId = row.itemId || this.ids[0]
      getSupervisionItem(itemId).then(response => {
        this.form = response.data
        // 解析 JSON 字段
        if (this.form.industryIds && typeof this.form.industryIds === 'string') {
          this.form.industryIds = JSON.parse(this.form.industryIds)
        }
        if (this.form.standardLanguageIds && typeof this.form.standardLanguageIds === 'string') {
          this.form.standardLanguageIds = JSON.parse(this.form.standardLanguageIds)
        }
        this.open = true
        this.title = "修改监管事项"
      })
    },
    handleStatusChange(row) {
      let text = row.isEnabled === "1" ? "启用" : "禁用"
      this.$confirm('确认要"' + text + '""' + row.itemName + '"事项吗？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return updateSupervisionItem(row)
      }).then(() => {
        this.msgSuccess(text + "成功")
      }).catch(() => {
        row.isEnabled = row.isEnabled === "1" ? "0" : "1"
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          // 处理 JSON 字段
          if (this.form.industryIds && Array.isArray(this.form.industryIds)) {
            this.form.industryIds = JSON.stringify(this.form.industryIds)
          }
          if (this.form.standardLanguageIds && Array.isArray(this.form.standardLanguageIds)) {
            this.form.standardLanguageIds = JSON.stringify(this.form.standardLanguageIds)
          }
          if (this.form.itemId != undefined) {
            updateSupervisionItem(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addSupervisionItem(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const itemIds = row.itemId || this.ids
      this.$confirm('是否确认删除事项编号为"' + itemIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteSupervisionItem(itemIds)
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
        itemId: undefined,
        parentId: 0,
        itemName: undefined,
        supervisionType: undefined,
        level: 1,
        sortOrder: 0,
        isEnabled: 1,
        industryIds: [],
        standardLanguageIds: [],
        checkPoints: undefined,
        legalBasis: undefined
      }
      this.resetForm("form")
    }
  }
}
</script>
