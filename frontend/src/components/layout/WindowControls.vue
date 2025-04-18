<script setup lang="ts">
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits } from 'vue';
import { RouteRecordName } from 'vue-router';

const props = defineProps<{
  isNotMac: boolean;
  isMaximised: boolean;
  routeName?: RouteRecordName | null;
}>();

const emit = defineEmits<{
  'minimise': [];
  'maximise': [];
  'unmaximise': [];
  'quit': [];
}>();
</script>

<template>
  <!-- windows 定制化窗口按钮 -->
  <div v-if="isNotMac" class="flex h-10 mr-6 justify-between justify-center items-center flex-0 border-b border-base-200" style="--wails-draggable:drag">
    <div class="px-4 text-base-content">{{ routeName }}</div>
    <div class="flex">
      <button class="btn btn-ghost btn-sm w-10 h-10" @click="$emit('minimise')">
        <Icon icon="mdi:window-minimize" />
      </button>
      <button class="btn btn-ghost btn-sm w-10 h-10" @click="isMaximised ? $emit('unmaximise') : $emit('maximise')">
        <Icon icon="mdi:window-maximize" />
      </button>
      <button class="btn btn-ghost btn-sm w-10 h-10" @click="$emit('quit')">
        <Icon icon="mdi:window-close" />
      </button>
    </div>
  </div>
  <!-- mac 标题栏 -->
  <div v-else class="h-8 flex items-center justify-center text-base-content border-b border-base-200" style="--wails-draggable:drag">
    {{ routeName }}
  </div>
</template>
