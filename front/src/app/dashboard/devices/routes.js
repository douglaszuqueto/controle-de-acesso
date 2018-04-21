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
    path: 'devices',
    component: main,
    redirect: {
      name: 'devices.index'
    },
    children: [
      {
        path: 'index',
        name: 'devices.index',
        component: index,
        meta
      },
      {
        path: 'create',
        name: 'device.new',
        component: Create,
        meta
      },
      {
        path: 'update/:id',
        name: 'device.update',
        component: Update,
        meta
      },
      {
        path: ':id/tags',
        name: 'device.tags',
        component: Tags,
        meta
      }
    ]
  }
]
