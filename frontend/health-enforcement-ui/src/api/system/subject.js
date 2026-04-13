import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询监管单位列表
export function listSubject(query) {
  return request({
    url: '/system/subject/list',
    method: 'get',
    params: query
  })
}

// 查询监管单位详细
export function getSubject(subjectId) {
  return request({
    url: '/system/subject/' + parseStrEmpty(subjectId),
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

// 删除监管单位
export function deleteSubject(subjectIds) {
  return request({
    url: '/system/subject/' + subjectIds,
    method: 'delete'
  })
}

// 搜索监管单位
export function searchSubject(query) {
  return request({
    url: '/system/subject/search',
    method: 'get',
    params: query
  })
}
