<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="法规标题" prop="title">
        <el-input v-model="queryParams.title" placeholder="请输入法规标题" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="法律类型" prop="legalType">
        <el-select v-model="queryParams.legalType" placeholder="请选择法律类型" clearable size="small">
          <el-option v-for="item in legalTypeOptions" :key="item.typeCode" :label="item.typeName" :value="item.typeCode" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="请选择状态" clearable size="small">
          <el-option label="有效" :value="1" />
          <el-option label="废止" :value="0" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['law:regulation:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['law:regulation:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['law:regulation:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="regulationList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="法规 ID" prop="regulationId" width="80" />
      <el-table-column label="法规标题" prop="title" :show-overflow-tooltip="true" />
      <el-table-column label="法律类型" prop="legalType" width="120">
        <template slot-scope="scope">
          <span>{{ formatLegalType(scope.row.legalType) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="发布机关" prop="publishOrg" width="150" :show-overflow-tooltip="true" />
      <el-table-column label="发布日期" prop="publishDate" width="100" align="center">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.publishDate, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">有效</el-tag>
          <el-tag v-else type="danger">废止</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="200" fixed="right">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['law:regulation:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-view" @click="handleView(scope.row)" v-hasPermi="['law:regulation:query']">查看</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['law:regulation:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改法律法规对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="法规标题" prop="title">
              <el-input v-model="form.title" placeholder="请输入法规标题" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="法律类型" prop="legalType">
              <el-select v-model="form.legalType" placeholder="请选择法律类型">
                <el-option v-for="item in legalTypeOptions" :key="item.typeCode" :label="item.typeName" :value="item.typeCode" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="发布机关" prop="publishOrg">
              <el-input v-model="form.publishOrg" placeholder="请输入发布机关" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="发布日期" prop="publishDate">
              <el-date-picker v-model="form.publishDate" type="date" value-format="yyyy-MM-dd" placeholder="选择发布日期" style="width: 100%" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="12">
            <el-form-item label="生效日期" prop="effectiveDate">
              <el-date-picker v-model="form.effectiveDate" type="date" value-format="yyyy-MM-dd" placeholder="选择生效日期" style="width: 100%" />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态" prop="status">
              <el-radio-group v-model="form.status">
                <el-radio :label="1">有效</el-radio>
                <el-radio :label="0">废止</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-form-item label="监管类型">
              <el-select v-model="form.supervisionTypes" multiple placeholder="请选择监管类型" style="width: 100%">
                <el-option v-for="item in supervisionTypeOptions" :key="item.typeCode" :label="item.typeName" :value="item.typeCode" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="24">
            <el-form-item label="法规内容" prop="content">
              <el-input v-model="form.content" type="textarea" :rows="10" placeholder="请输入法规内容" />
            </el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 查看法律法规详情对话框 -->
    <el-dialog title="法律法规详情" :visible.sync="viewOpen" width="800px" append-to-body>
      <el-descriptions :column="2" border>
        <el-descriptions-item label="法规标题">{{ viewForm.title }}</el-descriptions-item>
        <el-descriptions-item label="法律类型">{{ formatLegalType(viewForm.legalType) }}</el-descriptions-item>
        <el-descriptions-item label="发布机关">{{ viewForm.publishOrg }}</el-descriptions-item>
        <el-descriptions-item label="发布日期">{{ parseTime(viewForm.publishDate, '{y}-{m}-{d}') }}</el-descriptions-item>
        <el-descriptions-item label="生效日期">{{ parseTime(viewForm.effectiveDate, '{y}-{m}-{d}') }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="viewForm.status === 1" type="success" size="small">有效</el-tag>
          <el-tag v-else type="danger" size="small">废止</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="监管类型" :span="2">
          <el-tag v-for="(item, index) in viewForm.supervisionTypes" :key="index" size="small" style="margin-right: 5px">{{ item }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="法规内容" :span="2">
          <div style="white-space: pre-wrap; max-height: 300px; overflow-y: auto">{{ viewForm.content }}</div>
        </el-descriptions-item>
      </el-descriptions>
      <div slot="footer" class="dialog-footer">
        <el-button @click="viewOpen = false">关 闭</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listRegulation, getRegulation, addRegulation, updateRegulation, deleteRegulation, getLegalTypeList, getSupervisionTypeList } from "@/api/system/regulation"

export default {
  name: "Regulation",
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      regulationList: [],
      open: false,
      viewOpen: false,
      title: "",
      form: {},
      viewForm: {},
      legalTypeOptions: [],
      supervisionTypeOptions: [],
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        title: undefined,
        legalType: undefined,
        status: undefined
      },
      rules: {
        title: [{ required: true, message: "法规标题不能为空", trigger: "blur" }],
        legalType: [{ required: true, message: "法律类型不能为空", trigger: "change" }]
      }
    }
  },
  created() {
    this.getList()
    this.getDicts()
  },
  methods: {
    getList() {
      this.loading = true
      listRegulation(this.queryParams).then(response => {
        this.regulationList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    getDicts() {
      getLegalTypeList().then(response => {
        this.legalTypeOptions = response.data
      })
      getSupervisionTypeList().then(response => {
        this.supervisionTypeOptions = response.data
      })
    },
    formatLegalType(type) {
      const item = this.legalTypeOptions.find(x => x.typeCode === type)
      return item ? item.typeName : type
    },
    resetQuery() {
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.regulationId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加法律法规"
    },
    handleView(row) {
      getRegulation(row.regulationId).then(response => {
        this.viewForm = response.data
        this.viewOpen = true
      })
    },
    handleUpdate(row) {
      this.reset()
      const regulationId = row.regulationId || this.ids[0]
      getRegulation(regulationId).then(response => {
        this.form = response.data
        this.open = true
        this.title = "修改法律法规"
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.regulationId != undefined) {
            updateRegulation(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addRegulation(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const regulationIds = row.regulationId || this.ids
      this.$confirm('是否确认删除法规编号为"' + regulationIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteRegulation(regulationIds)
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
        regulationId: undefined,
        title: undefined,
        legalType: undefined,
        supervisionTypes: [],
        publishOrg: undefined,
        publishDate: undefined,
        effectiveDate: undefined,
        status: 1,
        content: undefined
      }
      this.resetForm("form")
    }
  }
}
</script>
