import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/caseform/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/caseform/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/caseform/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/caseform/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/caseform/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/caseform/delete',
    method: 'post',
    data
  })
}
