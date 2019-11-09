import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/sys/auth/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/sys/auth/info',
    method: 'get'
  })
}

export function getMenus(token) {
  return request({
    url: '/sys/auth/getmenus',
    method: 'post',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/sys/auth/logout',
    method: 'post'
  })
}

export function changepwd(data) {
  return request({
    url: '/sys/auth/changepwd',
    method: 'post',
    data
  })
}

