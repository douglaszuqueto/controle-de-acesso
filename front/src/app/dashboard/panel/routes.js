import component from './templates/index.vue'

const meta = {
  requiresAuth: true
}

export default [
  {
    path: 'panel',
    name: 'panel',
    component,
    meta
  }
]
