import { routes as Home } from './home'
import { routes as Device } from './devices'
import { routes as User } from './users'
import { routes as Tag } from './tags'
import { routes as Panel } from './panel'
import { routes as Log } from './log'
import App from './App'

export default [
  {
    path: '/dashboard',
    component: App,
    children: [
      ...Home,
      ...Device,
      ...User,
      ...Tag,
      ...Panel,
      ...Log
    ]
  }
]
