import component from './templates/index.vue'

const meta = {
  requiresAuth: true
}

export default [
  {
    path: '',
    name: 'dashboard',
    component,
    meta
  }
]
