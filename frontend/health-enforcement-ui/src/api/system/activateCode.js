import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询激活码列表
export function listActivateCode(query) {
  return request({
    url: '/system/activate-code/list',
    method: 'get',
    params: query
  })
}

// 查询激活码详细
export function getActivateCode(codeId) {
  return request({
    url: '/system/activate-code/' + parseStrEmpty(codeId),
    method: 'get'
  })
}

// 生成激活码
export function generateActivateCode(data) {
  return request({
    url: '/system/activate-code/generate',
    method: 'post',
    data: data
  })
}

// 删除激活码
export function deleteActivateCode(codeIds) {
  return request({
    url: '/system/activate-code/' + codeIds,
    method: 'delete'
  })
}

// 禁用激活码
export function disableActivateCode(codeId) {
  return request({
    url: '/system/activate-code/disable/' + codeId,
    method: 'put'
  })
}
