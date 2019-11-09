import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/sys/role/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/sys/role/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/sys/role/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/sys/role/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/sys/role/delete',
    method: 'post',
    data
  })
}

export function requestRoleMenuIDList(roleid) {
  return request({
    url: '/sys/role/rolemenuidlist',
    method: 'get',
    params: { roleid }
  })
}

export function requestSetRole(roleid, data) {
  return request({
    url: '/sys/role/setrole',
    method: 'post',
    params: { roleid },
    data
  })
}

export function requestAll() {
  return request({
    url: '/sys/role/allrole',
    method: 'get'
  })
}
