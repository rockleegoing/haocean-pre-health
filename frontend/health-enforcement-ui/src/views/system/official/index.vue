<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="姓名" prop="realname">
        <el-input v-model="queryParams.realname" placeholder="请输入姓名" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="执法证号" prop="badgeNo">
        <el-input v-model="queryParamss.badgeNo" placeholder="请输入执法证号" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="部门" prop="department">
        <el-input v-model="queryParams.department" placeholder="请输入部门" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="人员状态" clearable size="small">
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
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:official:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:official:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:official:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="officialList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="人员 ID" prop="officialId" width="80" />
      <el-table-column label="姓名" prop="realname" :show-overflow-tooltip="true" />
      <el-table-column label="执法证号" prop="badgeNo" :show-overflow-tooltip="true" />
      <el-table-column label="部门" prop="department" :show-overflow-tooltip="true" />
      <el-table-column label="职位" prop="position" :show-overflow-tooltip="true" />
      <el-table-column label="执法类型" prop="lawType" :show-overflow-tooltip="true" />
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="danger">禁用</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="success">正常</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="160">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:official:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:official:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改执法人员对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="姓名" prop="realname">
          <el-input v-model="form.realname" placeholder="请输入姓名" />
        </el-form-item>
        <el-form-item label="执法证号" prop="badgeNo">
          <el-input v-model="form.badgeNo" placeholder="请输入执法证号" />
        </el-form-item>
        <el-form-item label="部门" prop="department">
          <el-input v-model="form.department" placeholder="请输入部门" />
        </el-form-item>
        <el-form-item label="职位" prop="position">
          <el-input v-model="form.position" placeholder="请输入职位" />
        </el-form-item>
        <el-form-item label="执法类型" prop="lawType">
          <el-select v-model="form.lawType" placeholder="请选择执法类型">
            <el-option label="卫生执法" value="卫生执法" />
            <el-option label="交通执法" value="交通执法" />
            <el-option label="城管执法" value="城管执法" />
            <el-option label="市场监管" value="市场监管" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-radio-group v-model="form.status">
            <el-radio :label="1">正常</el-radio>
            <el-radio :label="0">禁用</el-radio>
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
import { listOfficial, getOfficial, addOfficial, updateOfficial, deleteOfficial } from "@/api/system/official"

export default {
  name: "Official",
  dicts: ['sys_normal_disable'],
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      officialList: [],
      open: false,
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        realname: undefined,
        badgeNo: undefined,
        department: undefined,
        status: undefined
      },
      rules: {
        realname: [{ required: true, message: "姓名不能为空", trigger: "blur" }],
        badgeNo: [{ required: true, message: "执法证号不能为空", trigger: "blur" }]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listOfficial(this.queryParams).then(response => {
        this.officialList = response.rows
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
      this.ids = selection.map(item => item.officialId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加执法人员"
    },
    handleUpdate(row) {
      this.reset()
      const officialId = row.officialId || this.ids[0]
      getOfficial(officialId).then(response => {
        this.form = response.data
        this.open = true
        this.title = "修改执法人员"
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.officialId != undefined) {
            updateOfficial(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addOfficial(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const officialIds = row.officialId || this.ids
      this.$confirm('是否确认删除人员编号为"' + officialIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteOfficial(officialIds)
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
        officialId: undefined,
        realname: undefined,
        badgeNo: undefined,
        department: undefined,
        position: undefined,
        lawType: undefined,
        status: 1
      }
      this.resetForm("form")
    }
  }
}
</script>
