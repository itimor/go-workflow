import request from '@/utils/request'

export function requestList(query) {
  return request({
    url: '/workflow/casetype/list',
    method: 'get',
    params: query
  })
}

export function requestDetail(id) {
  return request({
    url: '/workflow/casetype/detail',
    method: 'get',
    params: { id }
  })
}

export function requestUpdate(data) {
  return request({
    url: '/workflow/casetype/update',
    method: 'post',
    data
  })
}

export function requestCreate(data) {
  return request({
    url: '/workflow/casetype/create',
    method: 'post',
    data
  })
}

export function requestDelete(data) {
  return request({
    url: '/workflow/casetype/delete',
    method: 'post',
    data
  })
}

export function requestCreateSteps(data) {
  return request({
    url: '/workflow/casetypestep/createsteps',
    method: 'post',
    data
  })
}