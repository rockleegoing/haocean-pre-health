import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询数据同步列表
export function listSync(query) {
  return request({
    url: '/system/sync/list',
    method: 'get',
    params: query
  })
}

// 查询数据同步详细
export function getSync(syncId) {
  return request({
    url: '/system/sync/' + parseStrEmpty(syncId),
    method: 'get'
  })
}

// 删除数据同步记录
export function delSync(syncId) {
  return request({
    url: '/system/sync/' + syncId,
    method: 'delete'
  })
}

// 重试同步任务
export function retrySync(data) {
  return request({
    url: '/system/sync/retry',
    method: 'post',
    data: data
  })
}
