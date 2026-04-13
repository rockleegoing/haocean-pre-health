import request from '@/utils/request'

// 查询行业分类列表
export function listIndustry(query) {
  return request({
    url: '/system/industry/list',
    method: 'get',
    params: query
  })
}

// 获取行业分类树
export function getIndustryTree(query) {
  return request({
    url: '/system/industry/tree',
    method: 'get',
    params: query
  })
}
