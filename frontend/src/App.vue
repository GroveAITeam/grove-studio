<script setup lang="ts">
import { Icon } from '@iconify/vue';
import {useDark, useToggle} from "@vueuse/core";
import { WindowMinimise, WindowMaximise, WindowUnmaximise, Quit } from "../wailsjs/runtime";
import { ref } from 'vue';
import { useRoute } from 'vue-router';

console.log('aa')
const menu = [
  { text: "首页", href: "/", icon: "material-symbols:home-app-logo" },
  { text: "对话模型", href: "/llm", icon: "material-symbols:supervisor-account-outline-rounded" },
]
const activeClass = 'text-primary'
const isNotMac = navigator.userAgent.toUpperCase().indexOf('MAC') < 0;
const isMaximised = ref(false);
let isDark = useDark()
const toggleDark = useToggle(isDark);
const route = useRoute();
</script>

<template>
  <main class="flex h-screen overflow-hidden" :data-theme="isDark ? 'grovedark' : 'grovechat'">
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
        <button @click="toggleDark()" class="text-2xl mt-1">
          <span v-if="isDark"><Icon icon="icon-park:dark-mode" /></span>
          <span v-else><Icon icon="icon-park:sun-one" /></span>
        </button>
        <router-link to="/setup" v-slot="{isActive}">
          <Icon icon="material-symbols:menu-rounded" :class="isActive && activeClass" />
        </router-link>
      </div>
    </nav>
    <!-- 内容面板 -->
    <div class="w-full bg-base-100 flex flex-col h-screen">
      <!-- windows 定制化窗口按钮 -->
      <div v-if="isNotMac" class="flex h-10 justify-between items-center flex-0 border-b border-base-200">
        <div class="px-4 text-base-content">{{ route.name }}</div>
        <div class="flex">
          <button class="btn btn-ghost btn-sm w-10 h-10" @click="WindowMinimise">
            <Icon icon="mdi:window-minimize" />
          </button>
          <button class="btn btn-ghost btn-sm w-10 h-10" @click="isMaximised ? (WindowUnmaximise(),isMaximised = false) : (WindowMaximise(), isMaximised = true)">
            <Icon icon="mdi:window-maximize" />
          </button>
          <button class="btn btn-error btn-sm w-10 h-10 hover:text-white" @click="Quit">
            <Icon icon="mdi:window-close" />
          </button>
        </div>
      </div>
      <div v-else class="h-8 flex items-center justify-center text-base-content border-b border-base-200" style="--wails-draggable:drag">
        {{ route.name }}
      </div>
      <!-- 页面内容 -->
      <div class="overflow-y-auto flex-1 py-4 px-6" style="height: calc(100vh - 40px);">
        <router-view />
      </div>
    </div>
  </main>
</template>

