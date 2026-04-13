<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="检查类型" prop="checkType">
        <el-select v-model="queryParams.checkType" placeholder="请选择检查类型" clearable size="small">
          <el-option label="日常检查" value="日常检查" />
          <el-option label="专项检查" value="专项检查" />
          <el-option label="投诉举报" value="投诉举报" />
          <el-option label="双随机抽查" value="双随机抽查" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="记录状态" clearable size="small">
          <el-option label="待上报" :value="1" />
          <el-option label="已上报" :value="2" />
          <el-option label="已审核" :value="3" />
          <el-option label="已归档" :value="4" />
        </el-select>
      </el-form-item>
      <el-form-item label="检查日期">
        <el-date-picker v-model="dateRange" style="width: 240px" value-format="yyyy-MM-dd" type="daterange" range-separator="-" start-placeholder="开始日期" end-placeholder="结束日期"></el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:record:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:record:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="warning" plain icon="el-icon-upload" size="mini" :disabled="single" @click="handleSubmit" v-hasPermi="['system:record:submit']">上报</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:record:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="recordList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="记录 ID" prop="recordId" width="80" />
      <el-table-column label="记录编号" prop="recordNo" :show-overflow-tooltip="true" />
      <el-table-column label="监管单位" prop="subjectName" :show-overflow-tooltip="true" />
      <el-table-column label="行业分类" prop="industryName" :show-overflow-tooltip="true" />
      <el-table-column label="检查类型" prop="checkType" :show-overflow-tooltip="true" />
      <el-table-column label="检查日期" align="center" prop="checkDate" width="100">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.checkDate, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="证据数" align="center" prop="evidenceCount" width="80" />
      <el-table-column label="文书数" align="center" prop="documentCount" width="80" />
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="info">草稿</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="warning">待上报</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="success">已上报</el-tag>
          <el-tag v-else-if="scope.row.status === 3" type="primary">已审核</el-tag>
          <el-tag v-else-if="scope.row.status === 4" type="info">已归档</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="200">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-view" @click="handleView(scope.row)" v-hasPermi="['system:record:view']">查看</el-button>
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:record:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-upload" @click="handleSubmit(scope.row)" v-if="scope.row.status === 1" v-hasPermi="['system:record:submit']">上报</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:record:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改执法记录对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="120px">
        <el-form-item label="监管单位" prop="subjectId">
          <el-select v-model="form.subjectId" placeholder="请选择监管单位" filterable @change="handleSubjectChange">
            <el-option v-for="item in subjectOptions" :key="item.subjectId" :label="item.name" :value="item.subjectId" />
          </el-select>
        </el-form-item>
        <el-form-item label="检查日期" prop="checkDate">
          <el-date-picker v-model="form.checkDate" type="date" value-format="yyyy-MM-dd" placeholder="选择检查日期" />
        </el-form-item>
        <el-form-item label="检查类型" prop="checkType">
          <el-select v-model="form.checkType" placeholder="请选择检查类型">
            <el-option label="日常检查" value="日常检查" />
            <el-option label="专项检查" value="专项检查" />
            <el-option label="投诉举报" value="投诉举报" />
            <el-option label="双随机抽查" value="双随机抽查" />
          </el-select>
        </el-form-item>
        <el-form-item label="执法人员">
          <el-select v-model="form.officialIds" multiple placeholder="请选择执法人员">
            <el-option v-for="item in officialOptions" :key="item.officialId" :label="item.realname + ' (' + item.badgeNo + ')'" :value="item.officialId" />
          </el-select>
        </el-form-item>
        <el-form-item label="检查情况">
          <el-input v-model="form.checkSituation" type="textarea" :rows="4" placeholder="请输入检查情况" />
        </el-form-item>
        <el-form-item label="处理意见">
          <el-input v-model="form.processOpinion" type="textarea" :rows="3" placeholder="请输入处理意见" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">保 存</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>

    <!-- 查看执法记录详情对话框 -->
    <el-dialog title="执法记录详情" :visible.sync="openView" width="900px" append-to-body>
      <el-descriptions :column="2" border v-if="viewRecord">
        <el-descriptions-item label="记录编号">{{ viewRecord.recordNo }}</el-descriptions-item>
        <el-descriptions-item label="监管单位">{{ viewRecord.subjectName }}</el-descriptions-item>
        <el-descriptions-item label="行业分类">{{ viewRecord.industryName }}</el-descriptions-item>
        <el-descriptions-item label="检查日期">{{ parseTime(viewRecord.checkDate, '{y}-{m}-{d}') }}</el-descriptions-item>
        <el-descriptions-item label="检查类型">{{ viewRecord.checkType }}</el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag v-if="viewRecord.status === 0" size="small">草稿</el-tag>
          <el-tag v-else-if="viewRecord.status === 1" size="small" type="warning">待上报</el-tag>
          <el-tag v-else-if="viewRecord.status === 2" size="small" type="success">已上报</el-tag>
          <el-tag v-else-if="viewRecord.status === 3" size="small" type="primary">已审核</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="检查情况" :span="2">{{ viewRecord.checkSituation }}</el-descriptions-item>
        <el-descriptions-item label="处理意见" :span="2">{{ viewRecord.processOpinion }}</el-descriptions-item>
      </el-descriptions>

      <el-divider>证据材料</el-divider>
      <el-table :data="evidenceList" size="small">
        <el-table-column label="类型" prop="type" width="80" />
        <el-table-column label="标题" prop="title" />
        <el-table-column label="描述" prop="description" :show-overflow-tooltip="true" />
        <el-table-column label="文件大小" prop="fileSize" width="100" />
        <el-table-column label="操作" width="100">
          <template slot-scope="scope">
            <el-button size="mini" type="text" @click="handleViewEvidence(scope.row)">查看</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="openView = false">关 闭</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listRecord, getRecord, addRecord, updateRecord, deleteRecord, submitRecord } from "@/api/system/record"
import { listSubject } from "@/api/system/subject"
import { listOfficial } from "@/api/system/official"

export default {
  name: "Record",
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      recordList: [],
      subjectOptions: [],
      officialOptions: [],
      open: false,
      openView: false,
      viewRecord: null,
      evidenceList: [],
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        checkType: undefined,
        status: undefined,
        beginTime: undefined,
        endTime: undefined
      },
      dateRange: [],
      rules: {
        subjectId: [{ required: true, message: "监管单位不能为空", trigger: "change" }],
        checkDate: [{ required: true, message: "检查日期不能为空", trigger: "change" }],
        checkType: [{ required: true, message: "检查类型不能为空", trigger: "change" }]
      }
    }
  },
  created() {
    this.getList()
    this.getSubjectOptions()
    this.getOfficialOptions()
  },
  watch: {
    dateRange(val) {
      this.queryParams.beginTime = val ? val[0] : undefined
      this.queryParams.endTime = val ? val[1] : undefined
    }
  },
  methods: {
    getList() {
      this.loading = true
      listRecord(this.queryParams).then(response => {
        this.recordList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    getSubjectOptions() {
      listSubject({ pageSize: 1000 }).then(response => {
        this.subjectOptions = response.rows
      })
    },
    getOfficialOptions() {
      listOfficial({ pageSize: 1000, status: 1 }).then(response => {
        this.officialOptions = response.rows
      })
    },
    handleSubjectChange(value) {
      const selected = this.subjectOptions.find(item => item.subjectId === value)
      if (selected) {
        this.form.industryId = selected.industryId
        this.form.industryName = selected.industryName
      }
    },
    resetQuery() {
      this.dateRange = []
      this.resetForm("queryForm")
      this.getList()
    },
    handleQuery() {
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.recordId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加执法记录"
    },
    handleUpdate(row) {
      this.reset()
      const recordId = row.recordId || this.ids[0]
      getRecord(recordId).then(response => {
        this.form = response.data.record
        this.evidenceList = response.data.evidences || []
        this.open = true
        this.title = "修改执法记录"
      })
    },
    handleView(row) {
      const recordId = row.recordId
      getRecord(recordId).then(response => {
        this.viewRecord = response.data.record
        this.evidenceList = response.data.evidences || []
        this.openView = true
      })
    },
    handleSubmit(row) {
      const recordId = row.recordId || this.ids[0]
      this.$confirm('是否确认上报该执法记录？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return submitRecord(recordId)
      }).then(() => {
        this.msgSuccess("上报成功")
        this.getList()
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.recordId != undefined) {
            updateRecord(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addRecord(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const recordIds = row.recordId || this.ids
      this.$confirm('是否确认删除记录编号为"' + recordIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteRecord(recordIds)
      }).then(() => {
        this.getList()
        this.msgSuccess("删除成功")
      })
    },
    handleViewEvidence(row) {
      if (row.filePath) {
        window.open(row.filePath, '_blank')
      } else {
        this.msgInfo("暂无文件")
      }
    },
    cancel() {
      this.open = false
      this.reset()
    },
    reset() {
      this.form = {
        recordId: undefined,
        subjectId: undefined,
        subjectName: undefined,
        industryId: undefined,
        industryName: undefined,
        checkDate: undefined,
        checkType: undefined,
        officialIds: [],
        checkSituation: undefined,
        processOpinion: undefined
      }
      this.resetForm("form")
    }
  }
}
</script>
