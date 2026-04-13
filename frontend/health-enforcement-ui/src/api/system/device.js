import request from '@/utils/request'
import { parseStrEmpty } from "@/utils/ruoyi";

// 查询设备列表
export function listDevice(query) {
  return request({
    url: '/system/device/list',
    method: 'get',
    params: query
  })
}

// 查询设备详细
export function getDevice(deviceId) {
  return request({
    url: '/system/device/' + parseStrEmpty(deviceId),
    method: 'get'
  })
}

// 新增设备
export function addDevice(data) {
  return request({
    url: '/system/device',
    method: 'post',
    data: data
  })
}

// 修改设备
export function updateDevice(data) {
  return request({
    url: '/system/device',
    method: 'put',
    data: data
  })
}

// 删除设备
export function deleteDevice(deviceIds) {
  return request({
    url: '/system/device/' + deviceIds,
    method: 'delete'
  })
}

// 禁用设备
export function disableDevice(deviceId) {
  return request({
    url: '/system/device/disable',
    method: 'put',
    params: { deviceId }
  })
}

// 设备激活
export function activateDevice(data) {
  return request({
    url: '/system/device/activate',
    method: 'post',
    data: data
  })
}

// 获取设备信息
export function getDeviceInfo(deviceId) {
  return request({
    url: '/system/device/info',
    method: 'get',
    params: { deviceId }
  })
}
