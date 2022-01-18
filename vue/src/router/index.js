import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Relayers from '@/views/Relayers.vue'
import BlockDetail from '@/views/BlockDetail.vue'

const routerHistory = createWebHistory()
const routes = [
  {
    path: '/',
    component: Index,
  },
  { path: '/blocks', component: Relayers },
  { path: '/block', component: BlockDetail },
  { path: '/relayers', component: Relayers },
]

const router = createRouter({
  history: routerHistory,
  routes,
})

export default router
