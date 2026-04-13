import request from '@/utils/request'

// ============ 监管类型 API ============

// 查询监管类型列表（移动端首页 -12 类网格）
export function listSupervisionType(query) {
  return request({
    url: '/system/standard-phrase/supervision-type/list',
    method: 'get',
    params: query
  })
}

// 获取监管类型详情
export function getSupervisionType(id) {
  return request({
    url: '/system/standard-phrase/supervision-type/' + id,
    method: 'get'
  })
}

// ============ 规范类别 API ============

// 查询规范类别列表
export function listCategory(query) {
  return request({
    url: '/system/standard-phrase/category/list',
    method: 'get',
    params: query
  })
}

// 获取规范类别详情
export function getCategory(id) {
  return request({
    url: '/system/standard-phrase/category/' + id,
    method: 'get'
  })
}

// ============ 规范条目 API ============

// 查询规范条目列表
export function listItem(query) {
  return request({
    url: '/system/standard-phrase/item/list',
    method: 'get',
    params: query
  })
}

// 获取规范条目详情
export function getItem(id) {
  return request({
    url: '/system/standard-phrase/item/' + id,
    method: 'get'
  })
}

// ============ 规范内容 API ============

// 查询规范内容列表
export function listContent(query) {
  return request({
    url: '/system/standard-phrase/content/list',
    method: 'get',
    params: query
  })
}

// 获取规范内容详情
export function getContent(id) {
  return request({
    url: '/system/standard-phrase/content/' + id,
    method: 'get'
  })
}

// ============ 搜索 API ============

// 搜索规范用语
export function searchStandardPhrase(query) {
  return request({
    url: '/system/standard-phrase/search',
    method: 'get',
    params: query
  })
}

// 获取完整树形结构
export function getFullTree(query) {
  return request({
    url: '/system/standard-phrase/tree',
    method: 'get',
    params: query
  })
}
