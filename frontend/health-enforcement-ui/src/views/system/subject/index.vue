<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="单位名称" prop="name">
        <el-input v-model="queryParams.name" placeholder="请输入单位名称" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="行业分类" prop="industryId">
        <el-select v-model="queryParams.industryId" placeholder="请选择行业分类" clearable size="small">
          <el-option v-for="item in industryOptions" :key="item.industryId" :label="item.industryName" :value="item.industryId" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="单位状态" clearable size="small">
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
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:subject:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:subject:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:subject:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="subjectList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="单位 ID" prop="subjectId" width="80" />
      <el-table-column label="单位名称" prop="name" :show-overflow-tooltip="true" />
      <el-table-column label="行业分类" prop="industryName" :show-overflow-tooltip="true" />
      <el-table-column label="地址" prop="address" :show-overflow-tooltip="true" />
      <el-table-column label="联系人" prop="contactPerson" width="100" />
      <el-table-column label="联系电话" prop="contactPhone" width="120" />
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="danger">停业</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="success">正常</el-tag>
          <el-tag v-else type="info">未知</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center" prop="createTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.createTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="160">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:subject:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:subject:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改监管单位对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="单位名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入单位名称" />
        </el-form-item>
        <el-form-item label="行业分类" prop="industryId">
          <el-select v-model="form.industryId" placeholder="请选择行业分类" @change="handleIndustryChange">
            <el-option v-for="item in industryOptions" :key="item.industryId" :label="item.industryName" :value="item.industryId" />
          </el-select>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="form.address" placeholder="请输入地址" />
        </el-form-item>
        <el-form-item label="联系人" prop="contactPerson">
          <el-input v-model="form.contactPerson" placeholder="请输入联系人" />
        </el-form-item>
        <el-form-item label="联系电话" prop="contactPhone">
          <el-input v-model="form.contactPhone" placeholder="请输入联系电话" />
        </el-form-item>
        <el-form-item label="许可证号" prop="licenseNo">
          <el-input v-model="form.licenseNo" placeholder="请输入许可证号" />
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">停业</el-radio>
          </el-radio-group>
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
import { listSubject, getSubject, addSubject, updateSubject, deleteSubject } from "@/api/system/subject"
import { listIndustry } from "@/api/system/industry"

export default {
  name: "Subject",
  dicts: ['sys_normal_disable'],
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      subjectList: [],
      industryOptions: [],
      open: false,
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
        industryId: undefined,
        status: undefined
      },
      rules: {
        name: [{ required: true, message: "单位名称不能为空", trigger: "blur" }],
        industryId: [{ required: true, message: "行业分类不能为空", trigger: "change" }]
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
      listSubject(this.queryParams).then(response => {
        this.subjectList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    getIndustryList() {
      listIndustry({ isEnabled: 1 }).then(response => {
        this.industryOptions = response.rows
      })
    },
    handleIndustryChange(value) {
      const selected = this.industryOptions.find(item => item.industryId === value)
      if (selected) {
        this.form.industryName = selected.industryName
      }
    },
    resetQuery() {
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.subjectId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加监管单位"
    },
    handleUpdate(row) {
      this.reset()
      const subjectId = row.subjectId || this.ids[0]
      getSubject(subjectId).then(response => {
        this.form = response.data
        this.open = true
        this.title = "修改监管单位"
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.subjectId != undefined) {
            updateSubject(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addSubject(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const subjectIds = row.subjectId || this.ids
      this.$confirm('是否确认删除单位编号为"' + subjectIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteSubject(subjectIds)
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
        subjectId: undefined,
        name: undefined,
        industryId: undefined,
        industryName: undefined,
        address: undefined,
        contactPerson: undefined,
        contactPhone: undefined,
        licenseNo: undefined,
        status: 1
      }
      this.resetForm("form")
    }
  }
}
</script>
