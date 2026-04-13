import request from '@/utils/request'

// 检查数据更新
export function checkSync(query) {
  return request({
    url: '/system/sync/check',
    method: 'get',
    params: query
  })
}

// 同步行业分类
export function syncIndustries() {
  return request({
    url: '/system/sync/industries',
    method: 'get'
  })
}

// 同步文书模板
export function syncTemplates(industryId) {
  return request({
    url: '/system/sync/templates',
    method: 'get',
    params: { industryId }
  })
}

// 上报执法记录
export function syncRecords(data) {
  return request({
    url: '/system/sync/records',
    method: 'post',
    data: data
  })
}

// 上报监管单位
export function syncSubjects(data) {
  return request({
    url: '/system/sync/subjects',
    method: 'post',
    data: data
  })
}

// 获取同步状态
export function getSyncStatus(query) {
  return request({
    url: '/system/sync/status',
    method: 'get',
    params: query
  })
}

// 重试同步
export function retrySync(data) {
  return request({
    url: '/system/sync/retry',
    method: 'post',
    data: data
  })
}
