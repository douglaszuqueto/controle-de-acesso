import axios from 'axios'
import swal from 'sweetalert2'

const u = JSON.parse(localStorage.getItem('user'))
let token = ''
if (u !== null) {
  token = u.Token
}

axios.defaults.headers.common['Authorization'] = `${token}`

export const http = axios.create({
  baseURL: process.env.CDA_API_HOST
})

http.interceptors.request.use((config) => {
  return config
}, (error) => {
  return Promise.reject(error)
})

// Add a response interceptor
http.interceptors.response.use((response) => {
  return response
}, (error) => {
  if (typeof error === 'object') {
    swal('Error', error.message, 'error')
  }
  return Promise.reject(error)
})
