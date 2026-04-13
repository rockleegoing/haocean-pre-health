import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// ============ 监管类型 API ============

// 查询监管类型列表
export function listSupervisionType(query) {
  return request({
    url: '/system/standard-phrase/supervision-type/list',
    method: 'get',
    params: query
  })
}

// 查询监管类型详细
export function getSupervisionType(id) {
  return request({
    url: '/system/standard-phrase/supervision-type/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 新增监管类型
export function addSupervisionType(data) {
  return request({
    url: '/system/standard-phrase/supervision-type',
    method: 'post',
    data: data
  })
}

// 修改监管类型
export function updateSupervisionType(data) {
  return request({
    url: '/system/standard-phrase/supervision-type',
    method: 'put',
    data: data
  })
}

// 删除监管类型
export function deleteSupervisionType(ids) {
  return request({
    url: '/system/standard-phrase/supervision-type/' + ids,
    method: 'delete'
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

// 查询规范类别详细
export function getCategory(id) {
  return request({
    url: '/system/standard-phrase/category/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 新增规范类别
export function addCategory(data) {
  return request({
    url: '/system/standard-phrase/category',
    method: 'post',
    data: data
  })
}

// 修改规范类别
export function updateCategory(data) {
  return request({
    url: '/system/standard-phrase/category',
    method: 'put',
    data: data
  })
}

// 删除规范类别
export function deleteCategory(ids) {
  return request({
    url: '/system/standard-phrase/category/' + ids,
    method: 'delete'
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

// 查询规范条目详细
export function getItem(id) {
  return request({
    url: '/system/standard-phrase/item/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 新增规范条目
export function addItem(data) {
  return request({
    url: '/system/standard-phrase/item',
    method: 'post',
    data: data
  })
}

// 修改规范条目
export function updateItem(data) {
  return request({
    url: '/system/standard-phrase/item',
    method: 'put',
    data: data
  })
}

// 删除规范条目
export function deleteItem(ids) {
  return request({
    url: '/system/standard-phrase/item/' + ids,
    method: 'delete'
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

// 查询规范内容详细
export function getContent(id) {
  return request({
    url: '/system/standard-phrase/content/' + parseStrEmpty(id),
    method: 'get'
  })
}

// 新增规范内容
export function addContent(data) {
  return request({
    url: '/system/standard-phrase/content',
    method: 'post',
    data: data
  })
}

// 修改规范内容
export function updateContent(data) {
  return request({
    url: '/system/standard-phrase/content',
    method: 'put',
    data: data
  })
}

// 删除规范内容
export function deleteContent(ids) {
  return request({
    url: '/system/standard-phrase/content/' + ids,
    method: 'delete'
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
