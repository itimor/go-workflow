import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/workflow/casetypestep/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/workflow/casetypestep/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/workflow/casetypestep/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/workflow/casetypestep/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/workflow/casetypestep/delete',
    method: 'post',
    data
  })
}
