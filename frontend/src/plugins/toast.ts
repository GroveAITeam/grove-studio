import { App, createApp } from 'vue'
import Toast from '../components/Toast.vue'

export interface ToastOptions {
  message: string
  type?: 'success' | 'error' | 'warning' | 'info'
  duration?: number
}

let toastInstance: any = null

export const toast = {
  install: (app: App) => {
    const toastContainer = document.createElement('div')
    toastContainer.id = 'toast-container'
    document.body.appendChild(toastContainer)

    const toastApp = createApp(Toast)
    toastInstance = toastApp.mount('#toast-container')

    app.config.globalProperties.$toast = {
      show: (options: ToastOptions) => {
        toastInstance.showToast(options)
      },
      success: (message: string, duration?: number) => {
        toastInstance.showToast({ message, type: 'success', duration })
      },
      error: (message: string, duration?: number) => {
        toastInstance.showToast({ message, type: 'error', duration })
      },
      warning: (message: string, duration?: number) => {
        toastInstance.showToast({ message, type: 'warning', duration })
      },
      info: (message: string, duration?: number) => {
        toastInstance.showToast({ message, type: 'info', duration })
      }
    }
  }
}
