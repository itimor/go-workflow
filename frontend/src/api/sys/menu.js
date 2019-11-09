import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/sys/menu/list',
    method: 'get',
    params: query
  })
}

export function requestAll() {
  return request({
    url: '/sys/menu/allmenu',
    method: 'get'
  })
}

export function requestDetail(id) {
  return request({
    url: '/sys/menu/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/sys/menu/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/sys/menu/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/sys/menu/delete',
    method: 'post',
    data
  })
}

export function requestMenuButton(menucode) {
  return request({
    url: '/sys/menu/menubuttonlist',
    method: 'get',
    params: { menucode }
  })
}

