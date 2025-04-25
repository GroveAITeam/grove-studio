<script setup lang="ts">
import { Icon } from '@iconify/vue';
import {useDark, useToggle} from "@vueuse/core";
import { WindowMinimise, WindowMaximise, WindowUnmaximise, Quit } from "../wailsjs/runtime";
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import AppSidebar from './components/layout/AppSidebar.vue';
import WindowControls from './components/layout/WindowControls.vue';

const menu = [
  { text: "首页", href: "/", icon: "fluent:chat-24-filled" },
  { text: "对话模型", href: "/llm", icon: "fluent:cube-24-filled" },
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
    <!-- 侧边导航 -->
    <AppSidebar
      :menu="menu"
      :active-class="activeClass"
      :is-dark="isDark"
      @toggle-dark="toggleDark()"
    />

    <!-- 内容面板 -->
    <div class="w-full bg-base-100 flex flex-col h-screen">
      <!-- 窗口控制区域 -->
      <WindowControls
        :is-not-mac="isNotMac"
        :is-maximised="isMaximised"
        :route-name="route.name"
        @minimise="WindowMinimise"
        @maximise="isMaximised = true; WindowMaximise()"
        @unmaximise="isMaximised = false; WindowUnmaximise()"
        @quit="Quit"
      />

      <!-- 页面内容 -->
      <div class="overflow-y-auto flex-1 py-4 px-6" style="height: calc(100vh - 40px);">
        <router-view />
      </div>
    </div>
  </main>
</template>

