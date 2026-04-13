import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询执法记录列表
export function listRecord(query) {
  return request({
    url: '/system/record/list',
    method: 'get',
    params: query
  })
}

// 查询执法记录详细
export function getRecord(recordId) {
  return request({
    url: '/system/record/' + parseStrEmpty(recordId),
    method: 'get'
  })
}

// 新增执法记录
export function addRecord(data) {
  return request({
    url: '/system/record',
    method: 'post',
    data: data
  })
}

// 修改执法记录
export function updateRecord(data) {
  return request({
    url: '/system/record',
    method: 'put',
    data: data
  })
}

// 删除执法记录
export function deleteRecord(recordIds) {
  return request({
    url: '/system/record/' + recordIds,
    method: 'delete'
  })
}

// 上报执法记录
export function submitRecord(recordId) {
  return request({
    url: '/system/record/submit/' + recordId,
    method: 'put'
  })
}

// 上传证据
export function uploadEvidence(data) {
  return request({
    url: '/system/evidence/upload',
    method: 'post',
    headers: { 'Content-Type': 'multipart/form-data' },
    data: data
  })
}

// 删除证据
export function deleteEvidence(evidenceId) {
  return request({
    url: '/system/evidence/' + evidenceId,
    method: 'delete'
  })
}
