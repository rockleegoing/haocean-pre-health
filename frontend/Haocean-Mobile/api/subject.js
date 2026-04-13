import request from '@/utils/request'

// 查询监管单位列表
export function listSubject(query) {
  return request({
    url: '/system/subject/list',
    method: 'get',
    params: query
  })
}

// 获取监管单位详情
export function getSubject(subjectId) {
  return request({
    url: '/system/subject/' + subjectId,
    method: 'get'
  })
}

// 新增监管单位
export function addSubject(data) {
  return request({
    url: '/system/subject',
    method: 'post',
    data: data
  })
}

// 修改监管单位
export function updateSubject(data) {
  return request({
    url: '/system/subject',
    method: 'put',
    data: data
  })
}

// 搜索监管单位
export function searchSubject(keyword) {
  return request({
    url: '/system/subject/search',
    method: 'get',
    params: { keyword }
  })
}
