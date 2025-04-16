<script lang="ts" setup>
import APIIntegrationCard from './llm/APIIntegrationCard.vue';
import LocalModelCard from './llm/LocalModelCard.vue';
</script>

<template>
  <div class="bg-base-100 text-base-content flex flex-col gap-4">
    <!-- 顶部提示 -->
    <div class="bg-base-200/50 text-base-content border border-base-300/50 flex items-center gap-3 rounded-lg p-2">
      <div class="text-lg">💡</div>
      <div class="text-base-content/90 text-sm">API集成支持对接主流AI服务商，快速开始使用。本地模型即将推出，将提供完全的数据隐私保护和离线使用体验。</div>
    </div>

    <!-- 卡片区域 -->
    <div class="grid grid-cols-2 gap-4 flex-1 min-w-0">
      <APIIntegrationCard />
      <LocalModelCard />
    </div>
  </div>
</template>

<style scoped>
/* 所有样式已转换为Tailwind类，不再需要scoped CSS */
</style>
