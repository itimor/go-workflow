import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/casetype/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/casetype/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/casetype/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/casetype/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/casetype/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/casetype/delete',
    method: 'post',
    data
  })
}
