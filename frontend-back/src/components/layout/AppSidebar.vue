<script setup lang="ts">
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits } from 'vue';

interface MenuItem {
  text: string;
  href: string;
  icon: string;
}

const props = defineProps<{
  menu: MenuItem[];
  activeClass: string;
  isDark: boolean;
}>();

const emit = defineEmits<{
  'toggle-dark': []
}>();
</script>

<template>
  <nav
    class="flex-none flex flex-col justify-between w-[70px] items-center text-center select-none z-20 bg-base-200"
    style="--wails-draggable: drag"
  >
    <div class="menu mt-10 my-4 flex flex-col gap-6 text-2xl text-base-content">
      <router-link v-for="item in menu" :key="item.text" :to="item.href" v-slot="{isActive}">
        <Icon :icon="item.icon" :class="isActive && activeClass" />
      </router-link>
    </div>
    <div class="menu my-4 flex flex-col gap-4 text-2xl text-base-content">
      <button @click="$emit('toggle-dark')" class="text-2xl mt-1">
        <span v-if="isDark"><Icon icon="material-symbols:dark-mode-rounded" /></span>
        <span v-else><Icon icon="material-symbols:light-mode-rounded" /></span>
      </button>
      <router-link to="/setting" v-slot="{isActive}">
        <Icon icon="material-symbols:settings-rounded" :class="isActive && activeClass" />
      </router-link>
    </div>
  </nav>
</template>
