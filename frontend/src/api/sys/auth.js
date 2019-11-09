import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/auth/info',
    method: 'get'
  })
}

export function getMenus(token) {
  return request({
    url: '/auth/getmenus',
    method: 'post',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}

export function changepwd(data) {
  return request({
    url: '/auth/changepwd',
    method: 'post',
    data
  })
}

