import { getCurrentInstance } from 'vue'
import type { ToastOptions } from '../plugins/toast'

export const useToast = () => {
  const toast = getCurrentInstance()?.proxy?.$toast

  return {
    show: (options: ToastOptions) => toast?.show(options),
    success: (message: string, duration?: number) => toast?.success(message, duration),
    error: (message: string, duration?: number) => toast?.error(message, duration),
    warning: (message: string, duration?: number) => toast?.warning(message, duration),
    info: (message: string, duration?: number) => toast?.info(message, duration)
  }
}

// 为非组件环境提供一个全局访问点
let globalToast: ReturnType<typeof useToast> | null = null

export const initGlobalToast = () => {
  globalToast = useToast()
}

export const getGlobalToast = () => {
  if (!globalToast) {
    console.warn('Toast 未初始化，请确保在应用启动时调用 initGlobalToast()')
    return {
      show: () => {},
      success: () => {},
      error: () => {},
      warning: () => {},
      info: () => {}
    }
  }
  return globalToast
}
