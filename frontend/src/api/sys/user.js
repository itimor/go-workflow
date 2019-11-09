import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/user/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/user/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/user/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/user/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/user/delete',
    method: 'post',
    data
  })
}

export function requestUserRoleIDList(user_id) {
  return request({
    url: '/user/userroleidlist',
    method: 'get',
    params: { user_id }
  })
}

export function requestSetRole(user_id, data) {
  return request({
    url: '/user/setrole',
    method: 'post',
    params: { user_id },
    data
  })
}

