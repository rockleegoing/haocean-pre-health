import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询行业分类列表
export function listIndustry(query) {
  return request({
    url: '/system/industry/list',
    method: 'get',
    params: query
  })
}

// 查询行业分类详细
export function getIndustry(industryId) {
  return request({
    url: '/system/industry/' + parseStrEmpty(industryId),
    method: 'get'
  })
}

// 新增行业分类
export function addIndustry(data) {
  return request({
    url: '/system/industry',
    method: 'post',
    data: data
  })
}

// 修改行业分类
export function updateIndustry(data) {
  return request({
    url: '/system/industry',
    method: 'put',
    data: data
  })
}

// 删除行业分类
export function deleteIndustry(industryId) {
  return request({
    url: '/system/industry/' + industryId,
    method: 'delete'
  })
}
