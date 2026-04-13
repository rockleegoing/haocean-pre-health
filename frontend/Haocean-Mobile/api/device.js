import request from '@/utils/request'

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

// 更新最后登录
export function updateLastLogin(data) {
  return request({
    url: '/system/device/login',
    method: 'post',
    data: data
  })
}
