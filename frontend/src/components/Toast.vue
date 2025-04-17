<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface ToastProps {
  message: string
  type?: 'success' | 'error' | 'warning' | 'info'
  duration?: number
}

const toasts = ref<ToastProps[]>([])

const showToast = (options: ToastProps) => {
  const toast = {
    message: options.message,
    type: options.type || 'info',
    duration: options.duration || 3000
  }

  toasts.value.push(toast)
  setTimeout(() => {
    toasts.value.shift()
  }, toast.duration)
}

// 导出方法供其他组件使用
defineExpose({
  showToast
})
</script>

<template>
  <div class="toast toast-end toast-top z-[9999] gap-2">
    <transition-group name="toast">
      <div v-for="(toast, index) in toasts"
           :key="index"
           :class="[
             'alert shadow-lg min-w-[320px] animate-slide-in-right',
             {
               'alert-success': toast.type === 'success',
               'alert-error': toast.type === 'error',
               'alert-warning': toast.type === 'warning',
               'alert-info': toast.type === 'info'
             }
           ]"
      >
        <div class="flex items-center gap-2">
          <template v-if="toast.type === 'success'">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          </template>
          <template v-else-if="toast.type === 'error'">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          </template>
          <template v-else-if="toast.type === 'warning'">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
          </template>
          <template v-else>
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 stroke-current" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
          </template>
          <span class="text-sm font-medium">{{ toast.message }}</span>
        </div>
      </div>
    </transition-group>
  </div>
</template>

<style>
@keyframes slide-in-right {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.animate-slide-in-right {
  animation: slide-in-right 0.3s ease-out;
}

.toast-move,
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}
</style>
