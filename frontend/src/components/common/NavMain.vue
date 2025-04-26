<script setup lang="ts">
import type { LucideIcon } from 'lucide-vue-next'
import {
  SidebarGroup,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '@/components/ui/sidebar'
import { navigator } from '@/router/navigator'

defineProps<{
  items: {
    title: string
    url: string
    icon?: LucideIcon
    isActive?: boolean
  }[]
}>()

const handelPath = (path: string) => {
  navigator.navigate(path)
}
</script>

<template>
  <SidebarGroup>
    <SidebarMenu>
      <SidebarMenuItem
        v-for="item in items"
        :key="item.title"
        class="flex items-center justify-center mt-3"
      >
        <SidebarMenuButton class="cursor-pointer" :tooltip="item.title" @click="handelPath(item.url)">
          <component style="width: 24px;height: 24px;" :is="item.icon" v-if="item.icon" />
          <span>{{ item.title }}</span>
        </SidebarMenuButton>
      </SidebarMenuItem>
    </SidebarMenu>
  </SidebarGroup>
</template>
