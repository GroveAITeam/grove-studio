<script lang="ts" setup>
interface Spec {
  icon: string
  name: string
  compatible: boolean
}

interface Compatibility {
  status: string
  text: string
  specs: Spec[]
}

interface ModelInfo {
  size: string
  description: string
  compatibility: Compatibility
}

defineProps<{
  modelInfo: ModelInfo
}>()
</script>

<template>
  <div>
    <div class="bg-base-200/50 text-base-content inline-block px-2 py-1 rounded text-xs mb-1.5">{{ modelInfo.size }}</div>
    <p class="text-base-content/80 text-sm mb-3">{{ modelInfo.description }}</p>

    <!-- 设备兼容性 -->
    <div class="bg-base-100/50 rounded-md p-3">
      <div class="flex items-center justify-between">
        <span>您当前设备</span>
        <span
          class="px-2 py-1 rounded text-xs" :class="[
            modelInfo.compatibility.status === 'compatible' ? 'bg-success/20 text-success' : 'bg-error/20 text-error',
          ]"
        >
          {{ modelInfo.compatibility.text }}
        </span>
      </div>
      <div class="flex flex-col gap-1.5 mt-2">
        <div
          v-for="(spec, index) in modelInfo.compatibility.specs"
          :key="index"
          class="text-base-content flex items-center gap-2 text-sm"
        >
          <span>{{ spec.icon }}</span>
          <span>{{ spec.name }}</span>
          <span :class="[spec.compatible ? 'text-success' : 'text-error']">
            {{ spec.compatible ? '✓' : '✗' }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
