<script setup lang="ts">
import AppSidebar from '@/components/common/AppSidebar.vue'
import {SidebarProvider } from '@/components/ui/sidebar'
</script>

<template>
    <SidebarProvider>
      <AppSidebar />
      <slot />
    </SidebarProvider>
</template>

