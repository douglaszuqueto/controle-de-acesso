import main from './main.vue'
import index from './template/index.vue'
import Create from './template/create.vue'
import Update from './template/update.vue'
import Tags from './template/tags.vue'

const meta = {
  requiresAuth: true
}

export default [
  {
    path: 'users',
    component: main,
    redirect: {
      name: 'user.index'
    },
    children: [
      {
        path: 'index',
        name: 'user.index',
        component: index,
        meta
      },
      {
        path: 'create',
        name: 'user.new',
        component: Create,
        meta
      },
      {
        path: 'update/:id',
        name: 'user.update',
        component: Update,
        meta
      },
      {
        path: ':id/tags',
        name: 'user.tags',
        component: Tags,
        meta
      }
    ]
  }
]
