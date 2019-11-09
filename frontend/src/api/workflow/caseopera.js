import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/workflow/caseopera/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/workflow/caseopera/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/workflow/caseopera/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/workflow/caseopera/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/workflow/caseopera/delete',
    method: 'post',
    data
  })
}
