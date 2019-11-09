import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/workflow/caseform/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/workflow/caseform/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/workflow/caseform/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/workflow/caseform/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/workflow/caseform/delete',
    method: 'post',
    data
  })
}
