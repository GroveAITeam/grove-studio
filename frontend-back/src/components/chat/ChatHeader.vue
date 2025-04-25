<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits } from 'vue';

const props = defineProps<{
  title: string
}>();

const emit = defineEmits<{
  'toggle-settings': []
}>();
</script>

<template>
  <div class="border-b border-base-200 p-2 flex justify-between items-center bg-base-100">
    <h2 class="text-base font-medium text-base-content">{{ title }}</h2>
    <div class="flex gap-2">
      <button @click="$emit('toggle-settings')" class="p-1.5 text-base-content/70 hover:bg-base-200 rounded-md transition-colors">
        <Icon icon="material-symbols:settings" class="text-lg" />
      </button>
    </div>
  </div>
</template>
