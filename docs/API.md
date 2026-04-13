# 移动卫生执法系统 - API 接口文档

本文档描述移动卫生执法系统的所有 API 接口。

## 基本信息

- **Base URL**: `http://localhost:8080`
- **认证方式**: JWT Token
- **数据格式**: JSON

## 认证说明

大部分 API 需要在请求头中携带 JWT Token：

```
Authorization: Bearer <your_token>
```

---

## 1. 行业分类管理

### 1.1 查询行业分类列表

**接口**: `GET /system/industry/list`

**描述**: 获取行业分类列表，支持树形结构展示

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| industryName | string | 否 | 行业名称（模糊查询） |
| isEnabled | int | 否 | 启用状态 (-1:全部/0:禁用/1:启用) |

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "data": [
    {
      "industryId": 1,
      "industryCode": "PUBLIC",
      "industryName": "公共场所卫生",
      "parentId": 0,
      "level": 1,
      "isEnabled": 1,
      "orderNum": 1,
      "children": [
        {
          "industryId": 2,
          "industryCode": "PUBLIC_LODGE",
          "industryName": "住宿场所",
          "parentId": 1,
          "level": 2
        }
      ]
    }
  ]
}
```

### 1.2 获取行业分类详情

**接口**: `GET /system/industry/:id`

**描述**: 根据 ID 获取行业分类详情

**路径参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int64 | 是 | 行业分类 ID |

### 1.3 添加行业分类

**接口**: `POST /system/industry`

**描述**: 添加新的行业分类

**请求体**:
```json
{
  "industryCode": "NEW_CODE",
  "industryName": "新行业",
  "parentId": 0,
  "level": 1,
  "isEnabled": 1,
  "orderNum": 1
}
```

### 1.4 修改行业分类

**接口**: `PUT /system/industry`

**描述**: 修改行业分类信息

**请求体**:
```json
{
  "industryId": 1,
  "industryCode": "PUBLIC",
  "industryName": "公共场所卫生",
  "parentId": 0,
  "level": 1,
  "isEnabled": 1
}
```

### 1.5 删除行业分类

**接口**: `DELETE /system/industry/:ids`

**描述**: 批量删除行业分类

**路径参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| ids | string | 是 | 逗号分隔的 ID 列表 |

---

## 2. 监管单位管理

### 2.1 查询监管单位列表

**接口**: `GET /system/subject/list`

**描述**: 获取监管单位列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| subjectName | string | 否 | 单位名称（模糊查询） |
| industryId | int64 | 否 | 行业分类 ID |
| status | int | 否 | 状态 (-1:全部/0:停用/1:正常) |

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "rows": [
    {
      "subjectId": 1,
      "name": "某某酒店",
      "industryId": 2,
      "industryName": "住宿场所",
      "address": "某某街道 123 号",
      "contactPerson": "张三",
      "contactPhone": "13800138000",
      "status": 1
    }
  ],
  "total": 10
}
```

### 2.2 搜索监管单位

**接口**: `GET /system/subject/search`

**描述**: 根据关键词搜索监管单位

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| keyword | string | 否 | 搜索关键词 |

### 2.3 获取监管单位详情

**接口**: `GET /system/subject/:id`

**描述**: 根据 ID 获取监管单位详情

### 2.4 添加监管单位

**接口**: `POST /system/subject`

**描述**: 添加新的监管单位

**请求体**:
```json
{
  "name": "某某酒店",
  "industryId": 2,
  "address": "某某街道 123 号",
  "contactPerson": "张三",
  "contactPhone": "13800138000",
  "licenseNo": "SP123456",
  "status": 1
}
```

### 2.5 修改监管单位

**接口**: `PUT /system/subject`

**描述**: 修改监管单位信息

### 2.6 删除监管单位

**接口**: `DELETE /system/subject/:ids`

**描述**: 批量删除监管单位

---

## 3. 执法人员管理

### 3.1 查询执法人员列表

**接口**: `GET /system/official/list`

**描述**: 获取执法人员列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| badgeNo | string | 否 | 执法证号（模糊查询） |
| department | string | 否 | 所属部门 |
| status | int | 否 | 状态 (-1:全部/0:禁用/1:启用) |

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "rows": [
    {
      "officialId": 1,
      "userId": 1,
      "badgeNo": "ZF2024001",
      "department": "卫生监督所",
      "position": "科员",
      "status": 1
    }
  ],
  "total": 5
}
```

### 3.2 获取执法人员详情

**接口**: `GET /system/official/:id`

### 3.3 添加执法人员

**接口**: `POST /system/official`

**请求体**:
```json
{
  "userId": 1,
  "badgeNo": "ZF2024001",
  "department": "卫生监督所",
  "position": "科员",
  "status": 1
}
```

### 3.4 修改执法人员

**接口**: `PUT /system/official`

### 3.5 删除执法人员

**接口**: `DELETE /system/official/:ids`

---

## 4. 设备管理

### 4.1 查询设备列表

**接口**: `GET /system/device/list`

**描述**: 获取设备列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| deviceModel | string | 否 | 设备型号 |
| osType | string | 否 | 操作系统类型 |
| status | int | 否 | 状态 |

### 4.2 获取设备详情

**接口**: `GET /system/device/:id`

### 4.3 添加设备

**接口**: `POST /system/device`

### 4.4 修改设备

**接口**: `PUT /system/device`

### 4.5 删除设备

**接口**: `DELETE /system/device/:ids`

---

## 5. 激活码管理

### 5.1 查询激活码列表

**接口**: `GET /system/activate-code/list`

**描述**: 获取激活码列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| batchNo | string | 否 | 批次号 |
| status | int | 否 | 状态 |

### 5.2 生成激活码

**接口**: `POST /system/activate-code/generate`

**描述**: 批量生成激活码

**请求体**:
```json
{
  "batchNo": "BATCH20240413",
  "count": 100,
  "expireDays": 365
}
```

### 5.3 删除激活码

**接口**: `DELETE /system/activate-code/:ids`

---

## 6. 文书模板管理

### 6.1 查询文书模板列表

**接口**: `GET /system/template/list`

**描述**: 获取文书模板列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| templateName | string | 否 | 模板名称 |
| categoryId | int64 | 否 | 分类 ID |
| industryId | int64 | 否 | 行业分类 ID |
| isEnabled | int | 否 | 启用状态 |

### 6.2 获取文书模板详情

**接口**: `GET /system/template/:id`

### 6.3 上传文书模板

**接口**: `POST /system/template/upload`

**描述**: 上传文书模板文件

**请求体** (FormData):
| 字段名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| templateName | string | 是 | 模板名称 |
| categoryId | int64 | 是 | 分类 ID |
| industryId | int64 | 是 | 行业分类 ID |
| file | file | 是 | 模板文件 |
| fields | string | 否 | 填空项定义 (JSON) |

### 6.4 删除文书模板

**接口**: `DELETE /system/template/:ids`

### 6.5 预览文书模板

**接口**: `GET /system/template/preview/:id`

---

## 7. 执法记录管理

### 7.1 查询执法记录列表

**接口**: `GET /system/record/list`

**描述**: 获取执法记录列表，支持分页和条件查询

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pageNum | int | 否 | 页码（默认 1） |
| pageSize | int | 否 | 每页数量（默认 10） |
| subjectId | int64 | 否 | 单位 ID |
| industryId | int64 | 否 | 行业分类 ID |
| status | int | 否 | 状态 |
| checkType | string | 否 | 检查类型 |
| beginTime | string | 否 | 开始时间 |
| endTime | string | 否 | 结束时间 |

### 7.2 获取执法记录详情

**接口**: `GET /system/record/:id`

**描述**: 获取执法记录详情，包含关联的证据列表

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "data": {
    "record": {
      "recordId": 1,
      "recordNo": "JL20240413120000",
      "subjectId": 1,
      "subjectName": "某某酒店",
      "checkDate": "2024-04-13",
      "checkType": "日常检查",
      "status": 1
    },
    "evidences": [
      {
        "evidenceId": 1,
        "type": "photo",
        "title": "现场照片",
        "filePath": "/static/evidence/photo1.jpg"
      }
    ]
  }
}
```

### 7.3 添加执法记录

**接口**: `POST /system/record`

### 7.4 修改执法记录

**接口**: `PUT /system/record`

### 7.5 删除执法记录

**接口**: `DELETE /system/record/:ids`

### 7.6 上报执法记录

**接口**: `POST /system/record/submit/:id`

### 7.7 上传证据

**接口**: `POST /system/evidence/upload`

**描述**: 上传证据材料（支持照片/音频/视频/文档）

**请求体** (FormData):
| 字段名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| recordId | int64 | 是 | 执法记录 ID |
| type | string | 是 | 证据类型 (photo/audio/video/document) |
| title | string | 否 | 证据标题 |
| description | string | 否 | 证据描述 |
| file | file | 是 | 证据文件 |

### 7.8 删除证据

**接口**: `DELETE /system/evidence/:id`

---

## 8. 数据同步

### 8.1 检查数据更新

**接口**: `GET /system/sync/check`

**描述**: 移动端检查是否有新的数据更新

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| deviceId | string | 是 | 设备 ID |
| lastSyncTime | string | 否 | 最后同步时间 |

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "data": {
    "hasUpdate": true,
    "data": [
      {
        "queueId": 1,
        "table_name": "law_enforcement_record",
        "record_id": 1,
        "action": "insert",
        "data": {}
      }
    ],
    "syncTime": "2024-04-13 12:00:00"
  }
}
```

### 8.2 同步行业分类

**接口**: `GET /system/sync/industries`

**描述**: 移动端同步行业分类数据

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "data": {
    "industries": [
      {
        "industryId": 1,
        "industryName": "公共场所卫生",
        "children": []
      }
    ],
    "updateTime": "2024-04-13 12:00:00"
  }
}
```

### 8.3 同步文书模板

**接口**: `GET /system/sync/templates`

**描述**: 移动端同步文书模板数据

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| industryId | int64 | 否 | 行业分类 ID |

### 8.4 上报执法记录

**接口**: `POST /system/sync/records`

**描述**: 移动端上报执法记录到服务器

**请求体**:
```json
[
  {
    "recordId": 1,
    "recordNo": "JL20240413120000",
    "subjectId": 1,
    "checkDate": "2024-04-13",
    "checkType": "日常检查"
  }
]
```

**响应示例**:
```json
{
  "code": 200,
  "msg": "查询成功",
  "data": {
    "successCount": 10,
    "failCount": 0
  }
}
```

### 8.5 上报单位变更

**接口**: `POST /system/sync/subjects`

**描述**: 移动端上报监管单位变更到服务器

### 8.6 获取同步状态

**接口**: `GET /system/sync/status`

**描述**: 获取当前同步状态

### 8.7 重试同步

**接口**: `POST /system/sync/retry`

**描述**: 重试失败的同步任务

**请求体**:
```json
{
  "queueId": 1,
  "recordId": 1
}
```

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 参数错误 |
| 401 | 未授权/Token 无效 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 统一响应格式

所有 API 响应遵循以下格式：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}
```

列表接口返回格式：

```json
{
  "code": 200,
  "msg": "查询成功",
  "rows": [],
  "total": 10
}
```
