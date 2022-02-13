import { createRouter, createWebHistory } from 'vue-router'
import Fib from '../views/Fib.vue'

const routes = [
  {
    path: '/',
    name: 'Fib',
    component: Fib
  },
  {
    path: '/other',
    name: 'Other',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "other" */ '../views/Other.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
