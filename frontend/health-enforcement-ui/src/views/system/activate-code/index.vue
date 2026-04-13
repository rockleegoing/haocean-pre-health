<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="批次号" prop="batchNo">
        <el-input v-model="queryParams.batchNo" placeholder="请输入批次号" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="激活码状态" clearable size="small">
          <el-option label="未使用" :value="1" />
          <el-option label="已使用" :value="2" />
          <el-option label="已禁用" :value="3" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleGenerate" v-hasPermi="['system:activate-code:add']">生成</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:activate-code:remove']">删除</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="warning" plain icon="el-icon-download" size="mini" @click="handleExport" v-hasPermi="['system:activate-code:export']">导出</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="activateCodeList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="ID" prop="codeId" width="80" />
      <el-table-column label="激活码" prop="activateCode" :show-overflow-tooltip="true" />
      <el-table-column label="批次号" prop="batchNo" :show-overflow-tooltip="true" />
      <el-table-column label="执法人员" prop="officialName" :show-overflow-tooltip="true" />
      <el-table-column label="过期时间" align="center" prop="expireTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.expireTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="激活时间" align="center" prop="activateTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.activateTime) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 1" type="success">未使用</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="info">已使用</el-tag>
          <el-tag v-else-if="scope.row.status === 3" type="danger">已禁用</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="160">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-circle-check" @click="handleDisable(scope.row)" v-if="scope.row.status === 1" v-hasPermi="['system:activate-code:disable']">禁用</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:activate-code:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 生成激活码对话框 -->
    <el-dialog title="生成激活码" :visible.sync="openGenerate" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="批次号" prop="batchNo">
          <el-input v-model="form.batchNo" placeholder="请输入批次号（可选）" />
        </el-form-item>
        <el-form-item label="生成数量" prop="count">
          <el-input-number v-model="form.count" :min="1" :max="1000" :step="1" />
        </el-form-item>
        <el-form-item label="有效期（天）" prop="expireDay">
          <el-input-number v-model="form.expireDay" :min="1" :max="365" :step="1" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitGenerateForm">确 定</el-button>
        <el-button @click="cancelGenerate">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 激活码列表导出对话框 -->
    <el-dialog title="激活码列表" :visible.sync="openCodes" width="500px" append-to-body>
      <el-table :data="generatedCodes" height="400">
        <el-table-column label="序号" type="index" width="60" />
        <el-table-column label="激活码" prop="activateCode" />
        <el-table-column label="过期时间" prop="expireTime" width="180" />
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="openCodes = false">关 闭</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listActivateCode, getActivateCode, generateActivateCode, deleteActivateCode, disableActivateCode } from "@/api/system/activateCode"

export default {
  name: "ActivateCode",
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      multiple: true,
      total: 0,
      activateCodeList: [],
      openGenerate: false,
      openCodes: false,
      generatedCodes: [],
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        batchNo: undefined,
        status: undefined
      },
      rules: {
        count: [{ required: true, message: "生成数量不能为空", trigger: "blur" }],
        expireDay: [{ required: true, message: "有效期不能为空", trigger: "blur" }]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listActivateCode(this.queryParams).then(response => {
        this.activateCodeList = response.rows
        this.total = response.total
        this.loading = false
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
      this.ids = selection.map(item => item.codeId)
      this.multiple = !selection.length
    },
    handleGenerate() {
      this.reset()
      this.openGenerate = true
      this.title = "生成激活码"
    },
    submitGenerateForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          generateActivateCode(this.form).then(response => {
            this.msgSuccess("生成成功")
            this.openGenerate = false
            this.generatedCodes = response.data.codes
            this.openCodes = true
            this.getList()
          })
        }
      })
    },
    handleDisable(row) {
      this.$confirm('是否确认禁用激活码"' + row.activateCode + '"？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return disableActivateCode(row.codeId)
      }).then(() => {
        this.msgSuccess("禁用成功")
        this.getList()
      })
    },
    handleDelete(row) {
      const codeIds = row.codeId || this.ids
      this.$confirm('是否确认删除激活码编号为"' + codeIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteActivateCode(codeIds)
      }).then(() => {
        this.getList()
        this.msgSuccess("删除成功")
      })
    },
    handleExport() {
      this.$confirm('是否导出所有未使用的激活码？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // TODO: 实现导出功能
        this.msgSuccess("导出成功")
      })
    },
    cancelGenerate() {
      this.openGenerate = false
      this.reset()
    },
    reset() {
      this.form = {
        batchNo: undefined,
        count: 10,
        expireDay: 30
      }
      this.resetForm("form")
    }
  }
}
</script>
