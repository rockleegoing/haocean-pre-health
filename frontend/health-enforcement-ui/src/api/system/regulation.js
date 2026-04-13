import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询法律法规列表
export function listRegulation(query) {
  return request({
    url: '/api/admin/regulation/list',
    method: 'get',
    params: query
  })
}

// 查询法律法规详细
export function getRegulation(id) {
  return request({
    url: '/api/admin/regulation/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 新增法律法规
export function addRegulation(data) {
  return request({
    url: '/api/admin/regulation',
    method: 'post',
    data: data
  })
}

// 修改法律法规
export function updateRegulation(data) {
  return request({
    url: '/api/admin/regulation',
    method: 'put',
    data: data
  })
}

// 删除法律法规
export function deleteRegulation(ids) {
  return request({
    url: '/api/admin/regulation/' + ids,
    method: 'delete'
  })
}

// 获取法律法规首页数据
export function getRegulationHome() {
  return request({
    url: '/api/regulation/home',
    method: 'get'
  })
}

// 获取法律法规书本列表
export function getRegulationBookList(query) {
  return request({
    url: '/api/regulation/book-list',
    method: 'get',
    params: query
  })
}

// 获取法律法规书本详情
export function getRegulationBookDetail(id) {
  return request({
    url: '/api/regulation/book-detail/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 搜索法律法规
export function searchRegulation(query) {
  return request({
    url: '/api/regulation/search',
    method: 'get',
    params: query
  })
}

// 获取法律类型列表
export function getLegalTypeList() {
  return request({
    url: '/api/regulation/legal-type',
    method: 'get'
  })
}

// 获取监管类型列表
export function getSupervisionTypeList() {
  return request({
    url: '/api/regulation/supervision-type',
    method: 'get'
  })
}

// 获取定性依据列表
export function getBasisList(query) {
  return request({
    url: '/api/regulation/basis-list',
    method: 'get',
    params: query
  })
}

// 获取定性依据详情
export function getBasisDetail(id) {
  return request({
    url: '/api/regulation/basis-detail/' + parseStrEmpty(id),
    method: 'get'
  })
}
