import request from '@/utils/request'

// ============ 法律法规首页 API ============

// 获取首页数据（法律类型 + 监管类型 + 计数）
export function getHomeData() {
  return request({
    url: '/system/regulation/home',
    method: 'get'
  })
}

// ============ 法律类型 API ============

// 查询法律类型列表
export function listLegalType(query) {
  return request({
    url: '/system/regulation/legal-type/list',
    method: 'get',
    params: query
  })
}

// ============ 监管类型 API ============

// 查询监管类型列表
export function listSupervisionType(query) {
  return request({
    url: '/system/regulation/supervision-type/list',
    method: 'get',
    params: query
  })
}

// ============ 法律法规书本 API ============

// 查询法律法规书本列表
export function listBook(query) {
  return request({
    url: '/system/regulation/book/list',
    method: 'get',
    params: query
  })
}

// 获取法律法规书本详情
export function getBookDetail(id) {
  return request({
    url: '/system/regulation/book/' + id,
    method: 'get'
  })
}

// ============ 章节 API ============

// 获取章节详细内容
export function getChapterContent(chapterId) {
  return request({
    url: '/system/regulation/chapter/' + chapterId,
    method: 'get'
  })
}

// ============ 定性依据 API ============

// 获取定性依据列表
export function listBasis(query) {
  return request({
    url: '/system/regulation/basis/list',
    method: 'get',
    params: query
  })
}

// 获取定性依据详情
export function getBasisDetail(id) {
  return request({
    url: '/system/regulation/basis/' + id,
    method: 'get'
  })
}

// ============ 搜索 API ============

// 搜索法律法规
export function searchRegulation(query) {
  return request({
    url: '/system/regulation/search',
    method: 'get',
    params: query
  })
}

// ============ 收藏 API ============

// 添加收藏
export function addFavorite(data) {
  return request({
    url: '/system/regulation/favorite',
    method: 'post',
    data: data
  })
}

// 取消收藏
export function removeFavorite(id) {
  return request({
    url: '/system/regulation/favorite/' + id,
    method: 'delete'
  })
}

// 获取收藏列表
export function listFavorite(query) {
  return request({
    url: '/system/regulation/favorite/list',
    method: 'get',
    params: query
  })
}
