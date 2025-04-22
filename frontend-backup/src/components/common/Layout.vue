<script setup lang="ts">
import AppSidebar from '@/components/common/AppSidebar.vue'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar'
</script>

<template>
  <SidebarProvider>
    <AppSidebar />
    <SidebarInset>
      <header class="flex h-16 shrink-0 items-center gap-2 transition-[width,height] ease-linear group-has-[[data-collapsible=icon]]/sidebar-wrapper:h-12">
        <div class="flex gap-2 items-center px-4">
          <SidebarTrigger />
        </div>
        <slot name="header" />
      </header>
      <slot />
    </SidebarInset>
  </SidebarProvider>
</template>
