import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询执法人员列表
export function listOfficial(query) {
  return request({
    url: '/system/official/list',
    method: 'get',
    params: query
  })
}

// 查询执法人员详细
export function getOfficial(officialId) {
  return request({
    url: '/system/official/' + parseStrEmpty(officialId),
    method: 'get'
  })
}

// 新增执法人员
export function addOfficial(data) {
  return request({
    url: '/system/official',
    method: 'post',
    data: data
  })
}

// 修改执法人员
export function updateOfficial(data) {
  return request({
    url: '/system/official',
    method: 'put',
    data: data
  })
}

// 删除执法人员
export function deleteOfficial(officialIds) {
  return request({
    url: '/system/official/' + officialIds,
    method: 'delete'
  })
}

// 绑定设备
export function bindDevice(data) {
  return request({
    url: '/system/official/bind-device',
    method: 'post',
    data: data
  })
}
