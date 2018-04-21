import main from './main.vue'
import index from './template/index.vue'

const meta = {
  requiresAuth: true
}

export default [
  {
    path: 'log',
    component: main,
    redirect: {
      name: 'log.index'
    },
    children: [
      {
        path: 'index',
        name: 'log.index',
        component: index,
        meta
      }
    ]
  }
]
