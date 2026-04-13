<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" v-show="showSearch" label-width="68px">
      <el-form-item label="设备型号" prop="deviceModel">
        <el-input v-model="queryParams.deviceModel" placeholder="请输入设备型号" clearable size="small" @keyup.enter.native="handleQuery" />
      </el-form-item>
      <el-form-item label="操作系统" prop="osType">
        <el-select v-model="queryParams.osType" placeholder="请选择操作系统" clearable size="small">
          <el-option label="Android" value="Android" />
          <el-option label="iOS" value="iOS" />
          <el-option label="HarmonyOS" value="HarmonyOS" />
        </el-select>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="设备状态" clearable size="small">
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
        <el-button type="primary" plain icon="el-icon-plus" size="mini" @click="handleAdd" v-hasPermi="['system:device:add']">新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="success" plain icon="el-icon-edit" size="mini" :disabled="single" @click="handleUpdate" v-hasPermi="['system:device:edit']">修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button type="danger" plain icon="el-icon-delete" size="mini" :disabled="multiple" @click="handleDelete" v-hasPermi="['system:device:remove']">删除</el-button>
      </el-col>
      <right-toolbar :showSearch.sync="showSearch" @queryTable="getList"></right-toolbar>
    </el-row>

    <el-table v-loading="loading" :data="deviceList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="设备 ID" prop="deviceId" width="80" />
      <el-table-column label="设备型号" prop="deviceModel" :show-overflow-tooltip="true" />
      <el-table-column label="执法人员" prop="officialName" :show-overflow-tooltip="true" />
      <el-table-column label="执法证号" prop="badgeNo" :show-overflow-tooltip="true" />
      <el-table-column label="操作系统" prop="osType" width="100" />
      <el-table-column label="系统版本" prop="osVersion" width="100" />
      <el-table-column label="App 版本" prop="appVersion" width="100" />
      <el-table-column label="最后登录 IP" prop="lastLoginIp" width="140" />
      <el-table-column label="状态" align="center" width="80">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.status === 0" type="danger">禁用</el-tag>
          <el-tag v-else-if="scope.row.status === 1" type="success">正常</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width" width="160">
        <template slot-scope="scope">
          <el-button size="mini" type="text" icon="el-icon-edit" @click="handleUpdate(scope.row)" v-hasPermi="['system:device:edit']">修改</el-button>
          <el-button size="mini" type="text" icon="el-icon-delete" @click="handleDelete(scope.row)" v-hasPermi="['system:device:remove']">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total > 0" :total="total" :page.sync="queryParams.pageNum" :limit.sync="queryParams.pageSize" @pagination="getList" />

    <!-- 添加或修改设备对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="600px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="100px">
        <el-form-item label="设备型号" prop="deviceModel">
          <el-input v-model="form.deviceModel" placeholder="请输入设备型号" />
        </el-form-item>
        <el-form-item label="操作系统" prop="osType">
          <el-select v-model="form.osType" placeholder="请选择操作系统">
            <el-option label="Android" value="Android" />
            <el-option label="iOS" value="iOS" />
            <el-option label="HarmonyOS" value="HarmonyOS" />
          </el-select>
        </el-form-item>
        <el-form-item label="系统版本" prop="osVersion">
          <el-input v-model="form.osVersion" placeholder="请输入系统版本" />
        </el-form-item>
        <el-form-item label="App 版本" prop="appVersion">
          <el-input v-model="form.appVersion" placeholder="请输入 App 版本" />
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
import { listDevice, getDevice, addDevice, updateDevice, deleteDevice, disableDevice } from "@/api/system/device"

export default {
  name: "Device",
  dicts: ['sys_normal_disable'],
  data() {
    return {
      loading: true,
      showSearch: true,
      ids: [],
      single: true,
      multiple: true,
      total: 0,
      deviceList: [],
      open: false,
      title: "",
      form: {},
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        deviceModel: undefined,
        osType: undefined,
        status: undefined
      },
      rules: {
        deviceModel: [{ required: true, message: "设备型号不能为空", trigger: "blur" }]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      listDevice(this.queryParams).then(response => {
        this.deviceList = response.rows
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
      this.ids = selection.map(item => item.deviceId)
      this.single = selection.length !== 1
      this.multiple = !selection.length
    },
    handleAdd() {
      this.reset()
      this.open = true
      this.title = "添加设备"
    },
    handleUpdate(row) {
      this.reset()
      const deviceId = row.deviceId || this.ids[0]
      getDevice(deviceId).then(response => {
        this.form = response.data
        this.open = true
        this.title = "修改设备"
      })
    },
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.deviceId != undefined) {
            updateDevice(this.form).then(response => {
              this.msgSuccess("修改成功")
              this.open = false
              this.getList()
            })
          } else {
            addDevice(this.form).then(response => {
              this.msgSuccess("新增成功")
              this.open = false
              this.getList()
            })
          }
        }
      })
    },
    handleDelete(row) {
      const deviceIds = row.deviceId || this.ids
      this.$confirm('是否确认删除设备编号为"' + deviceIds + '"的数据项？', '警告', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        return deleteDevice(deviceIds)
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
        deviceId: undefined,
        deviceModel: undefined,
        osType: 'Android',
        osVersion: undefined,
        appVersion: undefined,
        status: 1
      }
      this.resetForm("form")
    }
  }
}
</script>
