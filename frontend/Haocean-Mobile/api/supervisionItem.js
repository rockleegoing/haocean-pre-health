import request from '@/utils/request'

// 查询监管事项分类列表
export function listSupervisionCategory() {
  return request({
    url: '/system/supervision/category/list',
    method: 'get'
  })
}

// 查询监管事项树
export function getSupervisionTree() {
  return request({
    url: '/system/supervision/tree',
    method: 'get'
  })
}

// 查询监管事项子项列表
export function getSupervisionChildren(parentId) {
  return request({
    url: '/system/supervision/children/' + parentId,
    method: 'get'
  })
}

// 查询监管事项详情
export function getSupervisionItem(itemId) {
  return request({
    url: '/system/supervision/' + itemId,
    method: 'get'
  })
}

// 查询监管事项列表（分页）
export function listSupervisionItem(query) {
  return request({
    url: '/system/supervision/list',
    method: 'get',
    params: query
  })
}
