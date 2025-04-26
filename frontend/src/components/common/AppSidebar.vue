<script setup lang="ts">
import type {
  SidebarProps,
} from '@/components/ui/sidebar'
import Logo from '@/assets/images/appicon.png'
import NavMain from '@/components/common/NavMain.vue'
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenuButton,
  SidebarRail,
  useSidebar,
} from '@/components/ui/sidebar'
import { navigator } from '@/router/navigator'

import {
  Bot,
  MessageCircle,
  Settings2,
} from 'lucide-vue-next'

const props = withDefaults(defineProps<SidebarProps>(), {
  collapsible: 'icon',
})

const { setOpen } = useSidebar()
setOpen(false)

const content = [
  {
    title: '开启新对话',
    url: '/',
    icon: MessageCircle,
    isActive: true,
  },
  {
    title: '对话模型',
    url: '/LLM/index',
    icon: Bot,
  },
]
const footer = {
  title: '设置',
  url: '/setting',
  icon: Settings2,
}
const header = {
  title: 'Grove Studio',
  url: '/',
  icon: MessageCircle,
}

const handelPath = (path: string) => {
  navigator.navigate(path)
}
</script>

<template>
  <Sidebar v-bind="props">
    <SidebarHeader>
      <SidebarMenuButton class="cursor-pointer" :tooltip="header.title" @click="handelPath(header.url)">
        <img class="w-[16px] h-[16px]" :src="Logo">
        <span>{{ header.title }}</span>
      </SidebarMenuButton>
    </SidebarHeader>
    <SidebarContent>
      <NavMain :items="content" />
    </SidebarContent>
    <SidebarFooter>
      <SidebarMenuButton class="cursor-pointer" :tooltip="footer.title" @click="handelPath(footer.url)">
        <component :is="footer.icon" v-if="footer.icon" />
        <span>{{ footer.title }}</span>
      </SidebarMenuButton>
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
