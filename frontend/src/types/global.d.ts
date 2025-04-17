import { ToastOptions } from '../plugins/toast'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $toast: {
      show: (options: ToastOptions) => void
      success: (message: string, duration?: number) => void
      error: (message: string, duration?: number) => void
      warning: (message: string, duration?: number) => void
      info: (message: string, duration?: number) => void
    }
  }
}
