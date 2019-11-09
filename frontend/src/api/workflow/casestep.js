import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/workflow/casestep/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/workflow/casestep/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/workflow/casestep/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/workflow/casestep/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/workflow/casestep/delete',
    method: 'post',
    data
  })
}
