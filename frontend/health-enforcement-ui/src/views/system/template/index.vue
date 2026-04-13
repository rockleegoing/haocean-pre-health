<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="模板名称" prop="templateName">
        <el-input v-model="queryParams.templateName" placeholder="请输入模板名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="行业分类" prop="industryId">
        <el-select v-model="queryParams.industryId" placeholder="请选择行业分类" clearable size="small">
          <el-option v-for="item in industryOptions" :key="item.industryId" :label="item.industryName" :value="item.industryId" />
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
        <el-button type="primary" plain icon="el-icon-upload" size="mini" @click="handleUpload" v-hasPermi="['system:template:upload']">上传</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:template:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:template:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="templateList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="模板 ID" prop="templateId" width="80" />
      <el-table-column label="模板名称" prop="templateName" :show-overflow-tooltip="true" />
      <el-table-column label="分类" prop="categoryName" :show-overflow-tooltip="true" />
      <el-table-column label="行业分类" prop="industryName" :show-overflow-tooltip="true" />
      <el-table-column label="版本" prop="version" width="80" />
      <el-table-column label="文件大小" prop="fileSize" width="100">
        <template slot-scope="scope">
          <span>{{ formatFileSize(scope.row.fileSize) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="启用状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.isEnabled === 1" type="success">启用</el-tag>
          <el-tag v-else type="info">禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="200">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-view" @click="handlePreview(scope.row)" v-hasPermi="['system:template:preview']">预览</el-button>
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:template:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:template:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 上传文书模板对话框 -->
    <el-dialog title="上传文书模板" :visible.sync="openUpload" width="600px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="模板名称" prop="templateName">
          <el-input v-model="form.templateName" placeholder="请输入模板名称" />
        </el-form-item>
        <el-form-item label="模板分类" prop="categoryId">
          <el-select v-model="form.categoryId" placeholder="请选择模板分类">
            <el-option label="现场检查笔录" :value="1" />
            <el-option label="询问笔录" :value="2" />
            <el-option label="责令改正通知书" :value="3" />
            <el-option label="行政处罚决定书" :value="4" />
            <el-option label="其他文书" :value="5" />
          </el-select>
        </el-form-item>
        <el-form-item label="行业分类" prop="industryId">
          <el-select v-model="form.industryId" placeholder="请选择行业分类">
            <el-option v-for="item in industryOptions" :key="item.industryId" :label="item.industryName" :value="item.industryId" />
          </el-select>
        </el-form-item>
        <el-form-item label="模板类型" prop="templateType">
          <el-select v-model="form.templateType" placeholder="请选择模板类型">
            <el-option label="Word 模板" value="word" />
            <el-option label="PDF 模板" value="pdf" />
          </el-select>
        </el-form-item>
        <el-form-item label="版本" prop="version">
          <el-input v-model="form.version" placeholder="请输入版本号" value="v1.0" />
        </el-form-item>
        <el-form-item label="字段定义" prop="fields">
          <el-input v-model="form.fieldsStr" type="textarea" :rows="4" placeholder='请输入 JSON 格式的字段定义，如：[{"name":"被检查单位","key":"subjectName"},{"name":"地址","key":"address"}]' />
        </el-form-item>
        <el-form-item label="模板文件" prop="file">
          <el-upload ref="upload" drag action="" :http-request="handleFileUpload" :limit="1" :auto-upload="false">
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
          </el-upload>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitUploadForm">确 定</el-button>
        <el-button @click="cancelUpload">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 预览文书模板对话框 -->
    <el-dialog title="文书预览" :visible.sync="openPreview" width="800px" append-to-body>
      <iframe :src="previewUrl" width="100%" height="600px" frameborder="0"></iframe>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="openPreview = false">关 闭</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listTemplate, getTemplate, uploadTemplate, updateTemplate, deleteTemplate } from "@/api/system/template"
import { listIndustry } from "@/api/system/industry"

export default {
  name: "Template",
  dicts: ['sys_normal_disable'],
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      templateList: [],
      industryOptions: [],
      openUpload: false,
      openPreview: false,
      previewUrl: '',
      title: "",
      form: {},
      selectedFile: null,
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        templateName: undefined,
        industryId: undefined,
        isEnabled: undefined
      },
      rules: {
        templateName: [{ required: true, message: "模板名称不能为空", trigger: "blur" }],
        categoryId: [{ required: true, message: "模板分类不能为空", trigger: "change" }],
        industryId: [{ required: true, message: "行业分类不能为空", trigger: "change" }],
        templateType: [{ required: true, message: "模板类型不能为空", trigger: "change" }]
      }
    }
  },
  created() {
    this.getList()
    this.getIndustryList()
  },
  methods: {
    getList() {
      this.loading = true
      listTemplate(this.queryParams).then(response => {
        this.templateList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    getIndustryList() {
      listIndustry({ isEnabled: 1 }).then(response => {
        this.industryOptions = response.rows
      })
    },
    resetQuery() {
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.templateId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleUpload() {
      this.reset()
      this.openUpload = true
      this.title = "上传文书模板"
    },
    handleFileUpload(file) {
      this.selectedFile = file.file
    },
    submitUploadForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (!this.selectedFile) {
            this.msgError("请选择模板文件")
            return
          }

          const formData = new FormData()
          formData.append('file', this.selectedFile)
          formData.append('templateName', this.form.templateName)
          formData.append('categoryId', this.form.categoryId)
          formData.append('categoryName', this.$refs.form.fields.find(f => f.prop === 'categoryId').label)
          formData.append('industryId', this.form.industryId)
          const selectedIndustry = this.industryOptions.find(item => item.industryId === this.form.industryId)
          if (selectedIndustry) {
            formData.append('industryName', selectedIndustry.industryName)
          }
          formData.append('templateType', this.form.templateType)
          formData.append('version', this.form.version || 'v1.0')
          if (this.form.fieldsStr) {
            formData.append('fields', this.form.fieldsStr)
          }

          uploadTemplate(formData).then(response => {
            this.msgSuccess("上传成功")
            this.openUpload = false
            this.getList()
          })
        }
      })
    },
    handleUpdate(row) {
      this.reset()
      const templateId = row.templateId || this.ids[0]
      getTemplate(templateId).then(response => {
        this.form = response.data
        this.openUpload = true
        this.title = "修改文书模板"
      })
    },
    handlePreview(row) {
      this.previewUrl = '/system/template/preview/' + row.templateId
      this.openPreview = true
    },
    handleDelete(row) {
      const templateIds = row.templateId || this.ids
      this.$confirm('是否确认删除模板编号为"' + templateIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteTemplate(templateIds)
      }).then(() => {
        this.getList()
        this.msgSuccess("删除成功")
      })
    },
    cancelUpload() {
      this.openUpload = false
      this.reset()
    },
    reset() {
      this.form = {
        templateId: undefined,
        templateName: undefined,
        categoryId: undefined,
        industryId: undefined,
        templateType: 'word',
        version: 'v1.0',
        fieldsStr: undefined
      }
      this.selectedFile = null
      this.resetForm("form")
    },
    formatFileSize(bytes) {
      if (!bytes) return '0 B'
      const sizes = ['B', 'KB', 'MB', 'GB']
      const i = Math.floor(Math.log(bytes) / Math.log(1024))
      return Math.round(bytes / Math.pow(1024, i) * 100) / 100 + ' ' + sizes[i]
    }
  }
}
</script>
