<script lang="ts" setup>
interface ModelSize {
  id: string
  label: string
  active: boolean
}

defineProps<{
  modelSizes: ModelSize[]
}>()

const emit = defineEmits<{
  switch: [size: string]
}>()
</script>

<template>
  <div class="flex gap-1.5">
    <button
      v-for="size in modelSizes"
      :key="size.id"
      class="bg-base-100/50 text-base-content hover:bg-base-200/50 px-2 py-1 rounded text-xs cursor-pointer" :class="[{ 'bg-base-100': size.active }]"
      @click="emit('switch', size.id)"
    >
      {{ size.label }}
    </button>
  </div>
</template>
