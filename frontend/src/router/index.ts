import Vue from 'vue'
import VueRouter from 'vue-router'
//import Home from '../views/Home.vue'
//import store from '../store/index'

Vue.use(VueRouter)

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'SelectPomkemon',
      component: () => import('../views/Select.vue'),
    },
    {
      path: '/attacker0',
      name: 'Attacker',
      component: () => import('../views/Attacker.vue'),
      props: {
        default: true,
        index: 0,
      },
    },
    {
      path: '/attacker1',
      name: 'Attacker',
      component: () => import('../views/Attacker.vue'),
      props: {
        default: true,
        index: 1,
      },
    },
    {
      path: '/attacker2',
      name: 'Attacker',
      component: () => import('../views/Attacker.vue'),
      props: {
        default: true,
        index: 2,
      },
    },
    {
      path: '/attacker3',
      name: 'Attacker',
      component: () => import('../views/Attacker.vue'),
      props: {
        default: true,
        index: 3,
      },
    },
    {
      path: '/defender',
      name: 'Defender',
      component: () => import('../views/Defender.vue'),
    },
    {
      path: '/speed',
      name: 'Speed',
      component: () => import ('../views/SpeedAdjuster.vue'),
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
