import main from './main.vue'
import index from './template/index.vue'
import Create from './template/create.vue'
import Update from './template/update.vue'

const meta = {
  requiresAuth: true
}

export default [
  {
    path: 'tags',
    component: main,
    redirect: {
      name: 'tag.index'
    },
    children: [
      {
        path: 'index',
        name: 'tag.index',
        component: index,
        meta
      },
      {
        path: 'create/:id_user?',
        name: 'tag.new',
        component: Create,
        meta
      },
      {
        path: 'update/:id',
        name: 'tag.update',
        component: Update,
        meta
      }
    ]
  }
]
