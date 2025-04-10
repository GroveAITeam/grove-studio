<script setup lang="ts">
import { Icon } from '@iconify/vue';
import {useDark, useToggle} from "@vueuse/core";
import { WindowMinimise, WindowMaximise, WindowUnmaximise, Quit } from "../wailsjs/runtime";
import { ref } from 'vue';

const menu = [
  { text: "首页", href: "/", icon: "material-symbols:home-app-logo" },
  { text: "用户", href: "/users", icon: "material-symbols:supervisor-account-outline-rounded" },
]
const activeClass = 'text-primary'
const isNotMac = navigator.userAgent.toUpperCase().indexOf('MAC') < 0;
const isMaximised = ref(false);
let isDark = useDark()
const toggleDark = useToggle(isDark);
</script>

<template>
  <main class="flex h-screen" :data-theme="isDark ? 'dark' : 'light'">
    <nav
        class="flex-none flex flex-col justify-between w-[70px] items-center text-center select-none z-20 bg-base-200"
        style="--wails-draggable: drag"
    >
      <div class="menu mt-12 my-4 flex flex-col gap-6 text-2xl text-base-content">
        <router-link v-for="item in menu" :key="item.text" :to="item.href" v-slot="{isActive}">
          <Icon :icon="item.icon" :class="isActive && activeClass" />
        </router-link>
      </div>
      <div class="menu my-4 flex flex-col gap-4 text-2xl text-base-content">
        <button @click="toggleDark()" class="text-2xl mt-1 hover:rotate-6">
          <span v-if="isDark"><Icon icon="icon-park:dark-mode" /></span>
          <span v-else><Icon icon="icon-park:sun-one" /></span>
        </button>
        <router-link to="/setup" v-slot="{isActive}">
          <Icon icon="material-symbols:menu-rounded" :class="isActive && activeClass" />
        </router-link>
      </div>
    </nav>
    <!-- 内容面板 -->
    <div class="pl-2 w-full bg-base-100">
      <!-- windows 定制化窗口按钮 -->
      <div v-if="isNotMac" class="flex h-10 justify-end flex-0">
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
      <div v-else class="h-8"></div>
      <!-- 页面内容 -->
      <router-view class="overflow-y-auto h-screen p-4" style="--wails-draggable:none;" />
    </div>
  </main>
</template>

