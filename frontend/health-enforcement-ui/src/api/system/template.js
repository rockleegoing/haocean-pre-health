import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询文书模板列表
export function listTemplate(query) {
  return request({
    url: '/system/template/list',
    method: 'get',
    params: query
  })
}

// 查询文书模板详细
export function getTemplate(templateId) {
  return request({
    url: '/system/template/' + parseStrEmpty(templateId),
    method: 'get'
  })
}

// 上传文书模板
export function uploadTemplate(data) {
  return request({
    url: '/system/template/upload',
    method: 'post',
    headers: { 'Content-Type': 'multipart/form-data' },
    data: data
  })
}

// 修改文书模板
export function updateTemplate(data) {
  return request({
    url: '/system/template',
    method: 'put',
    data: data
  })
}

// 删除文书模板
export function deleteTemplate(templateIds) {
  return request({
    url: '/system/template/' + templateIds,
    method: 'delete'
  })
}

// 预览文书模板
export function previewTemplate(templateId) {
  return request({
    url: '/system/template/preview/' + templateId,
    method: 'get',
    responseType: 'blob'
  })
}
