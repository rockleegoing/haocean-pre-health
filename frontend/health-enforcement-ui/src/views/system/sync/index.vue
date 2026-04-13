<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="同步类型" prop="syncType">
        <el-select v-model="queryParams.syncType" placeholder="请选择同步类型" clearable size="small">
          <el-option label="数据上报" value="数据上报" />
          <el-option label="数据下载" value="数据下载" />
          <el-option label="全量同步" value="全量同步" />
          <el-option label="增量同步" value="增量同步" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="同步状态" clearable size="small">
          <el-option label="成功" :value="0" />
          <el-option label="失败" :value="1" />
          <el-option label="进行中" :value="2" />
        </el-select>
      </el-form-item>
      <el-form-item label="同步时间">
        <el-date-picker v-model="dateRange" style="width: 240px" value-format="yyyy-MM-dd HH:mm:ss" type="datetimerange" range-separator="-" start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>

    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button type="primary" plain icon="el-icon-refresh" size="mini" @click="handleSync" v-hasPermi="['law:sync:retry']">立即同步</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['law:sync:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="syncList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="同步 ID" prop="syncId" width="80" />
      <el-table-column label="同步类型" prop="syncType" :show-overflow-tooltip="true" />
      <el-table-column label="数据模块" prop="moduleName" :show-overflow-tooltip="true" />
      <el-table-column label="记录数" align="center" prop="recordCount" width="80" />
      <el-table-column label="状态" align="center" width="100">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="success">成功</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="danger">失败</el-tag>
          <el-tag v-else-if="scope.row.status === 2" type="warning">进行中</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="消息" prop="message" :show-overflow-tooltip="true" />
      <el-table-column label="操作人" prop="operatorName" :show-overflow-tooltip="true" width="100" />
      <el-table-column label="同步时间" align="center" prop="syncTime" width="180">
        <template slot-scope="scope">
          <span>{{ parseTime(scope.row.syncTime, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="150">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-view" @click="handleView(scope.row)" v-hasPermi="['law:sync:query']">详情</el-button>
          <el-button size="mini" type="text" icon="el-icon-refresh" @click="handleRetry(scope.row)" v-hasPermi="['law:sync:retry']">重试</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 查看同步详情对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body>
      <el-form label-width="100px">
        <el-row>
          <el-col :span="12">
            <el-form-item label="同步 ID">{{ form.syncId }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="同步类型">{{ form.syncType }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数据模块">{{ form.moduleName }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="记录数">{{ form.recordCount }}</el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="状态">
              <el-tag v-if="form.status === 0" type="success">成功</el-tag>
              <el-tag v-else-if="form.status === 1" type="danger">失败</el-tag>
              <el-tag v-else-if="form.status === 2" type="warning">进行中</el-tag>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="操作人">{{ form.operatorName }}</el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="同步时间">{{ parseTime(form.syncTime, '{y}-{m}-{d} {h}:{i}:{s}') }}</el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="消息">{{ form.message }}</el-form-item>
          </el-col>
        </el-row>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="open = false">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { listSync, getSync, delSync, retrySync } from '@/api/system/sync'

export default {
  name: 'Sync',
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      syncList: [],
      open: false,
      title: '',
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        syncType: undefined,
        status: undefined
      },
      dateRange: []
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listSync(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
        this.syncList = response.rows
        this.total = response.total
        this.loading = false
      })
    },
    resetQuery() {
      this.dateRange = []
      this.resetForm('queryForm')
      this.getList()
    },
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.syncId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleView(row) {
      this.form = row
      this.title = '同步详情'
      this.open = true
    },
    handleSync() {
      this.$confirm('确认立即执行数据同步吗？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$modal.loading('正在同步数据，请稍候...')
        retrySync().then(response => {
          this.$modal.msgSuccess('同步请求已提交')
          this.getList()
        }).catch(() => {
          this.$modal.loading(false)
        })
      })
    },
    handleRetry(row) {
      this.$confirm('确认重新执行该同步任务吗？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        retrySync({ syncId: row.syncId }).then(response => {
          this.$modal.msgSuccess('重试请求已提交')
          this.getList()
        })
      })
    },
    handleDelete(row) {
      const syncIds = row.syncId || this.ids
      this.$modal.confirm('是否确认删除同步记录？').then(function() {
        return delSync(syncIds)
      }).then(() => {
        this.getList()
        this.$modal.msgSuccess('删除成功')
      }).catch(() => {})
    }
  }
}
</script>
