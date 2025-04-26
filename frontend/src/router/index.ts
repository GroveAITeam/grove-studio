import { createRouter, createWebHistory } from 'vue-router'

// 路由配置
const routes = [
  {
    path: '/',
    name: 'Chat',
    component: () => import('../views/Chat/index.vue'),
    meta: {
      title: '对话首页',
      keepAlive: true,
    },
  },
  {
    path: '/llm/index',
    name: 'LLMIndex',
    component: () => import('../views/LLM/index.vue'),
    meta: {
      title: '对话模型',
    },
  },
  {
    path: '/llm/cloud',
    name: 'LLMCloud',
    component: () => import('../views/LLM/Cloud.vue'),
    meta: {
      title: '云端模型',
    },
  },
  {
    path: '/setting',
    name: 'setting',
    component: () => import('../views/Setting/index.vue'),
    meta: {
      title: '设置页',
    },
  },
  // 404页面
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound/index.vue'),
    meta: {
      title: '404',
      keepAlive: false,
    },
  },
]

// 创建路由实例
export const router = createRouter({
  history: createWebHistory(),
  routes,
})

// 路由前置守卫
router.beforeEach((to, from, next) => {
  // TODO
  // console.log(to, from)
  next()
})

export default router
