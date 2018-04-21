import Vue from 'vue'
import Router from 'vue-router'

import { routes as app } from '../app/index'

Vue.use(Router)

const defaultPage = {
  path: '*',
  redirect: '/dashboard'
}

const routes = [
  defaultPage,
  ...app
]

const router = new Router({
  linkActiveClass: 'active',
  // mode: 'history',
  routes
})

const needAuth = auth => auth === true

router.beforeEach((to, from, next) => {
  const auth = to.meta.requiresAuth

  if (!needAuth(auth)) {
    next()
    return
  }
  const u = JSON.parse(localStorage.getItem('user'))
  if (u === null) {
    next({ name: 'login' }) // redirect to login
    return
  }

  next()
})

export default router
