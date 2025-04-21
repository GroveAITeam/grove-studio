import { router } from '@/router'
import { createPinia } from 'pinia'
import { markRaw } from 'vue'

// 创建pinia实例
const pinia = createPinia()

// 添加router到pinia上下文中
pinia.use(({ store }: Store.Params) => {
  store.$router = markRaw(router)
})

export default pinia
