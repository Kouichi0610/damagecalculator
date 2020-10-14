import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/current',
      name: 'Current',
      component: () => import('../views/Current.vue'),
    },
    {
      path: '/attacker',
      name: 'Attacker',
      component: () => import('../views/Attacker.vue'),
    },
    {
      path: '/defender',
      name: 'Defender',
      component: () => import('../views/Defender.vue'),
    },
    {
      path: '/speed',
      name: 'Speed',
      component: () => import('../views/Speed.vue'),
    },
    {
      path: '/sandboxts',
      name: 'SandBoxTs',
      component: () => import('../views/SandBoxTs.vue'),
    },
    {
      path: '/about',
      name: 'About',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    }
  ],
})

export default router
