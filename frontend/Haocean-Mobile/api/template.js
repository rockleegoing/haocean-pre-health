import request from '@/utils/request'

// 查询文书模板列表
export function listTemplate(query) {
  return request({
    url: '/system/template/list',
    method: 'get',
    params: query
  })
}

// 获取文书模板详情
export function getTemplate(templateId) {
  return request({
    url: '/system/template/' + templateId,
    method: 'get'
  })
}

// 生成文书
export function generateDocument(data) {
  return request({
    url: '/system/document/generate',
    method: 'post',
    data: data
  })
}
