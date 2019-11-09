import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/casetypestep/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/casetypestep/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/casetypestep/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/casetypestep/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/casetypestep/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/casetypestep/delete',
    method: 'post',
    data
  })
}
