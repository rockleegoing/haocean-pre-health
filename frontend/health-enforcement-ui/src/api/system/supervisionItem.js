import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询监管事项列表
export function listSupervisionItem(query) {
  return request({
    url: '/system/supervision/list',
    method: 'get',
    params: query
  })
}

// 查询监管事项详细
export function getSupervisionItem(itemId) {
  return request({
    url: '/system/supervision/' + parseStrEmpty(itemId),
    method: 'get'
  })
}

// 新增监管事项
export function addSupervisionItem(data) {
  return request({
    url: '/system/supervision',
    method: 'post',
    data: data
  })
}

// 修改监管事项
export function updateSupervisionItem(data) {
  return request({
    url: '/system/supervision',
    method: 'put',
    data: data
  })
}

// 删除监管事项
export function deleteSupervisionItem(itemId) {
  return request({
    url: '/system/supervision/' + itemId,
    method: 'delete'
  })
}

// 查询监管事项树
export function getSupervisionTree() {
  return request({
    url: '/system/supervision/tree',
    method: 'get'
  })
}

// 查询监管事项子项
export function getSupervisionChildren(parentId) {
  return request({
    url: '/system/supervision/children/' + parentId,
    method: 'get'
  })
}

// 查询监管事项分类列表
export function listSupervisionCategory() {
  return request({
    url: '/system/supervision/category/list',
    method: 'get'
  })
}
