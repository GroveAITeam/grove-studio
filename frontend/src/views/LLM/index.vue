<script setup lang="ts">
import CloudModelCard from '@/views/LLM/components/CloudModelCard.vue'
import LocalModelCard from '@/views/LLM/components/LocalModelCard.vue'
</script>

<template>
  <div class="flex flex-col p-8 gap-4">
    <div class=" text-base-content border border-base-300/50 p-2 rounded-lg flex items-center justify-center ">模型配置</div>
    <!-- 顶部提示 -->
    <div class=" text-base-content border border-base-300/50 flex items-center gap-3 rounded-lg p-2">
      <div class="text-lg">💡</div>
      <div class="text-base-content/90 text-sm">云端模型支持对接主流服务商，快速开始使用。本地模型即将推出，将提供完全的数据隐私保护和离线使用体验。</div>
    </div>

    <!-- 卡片区域 -->
    <div class="grid grid-cols-2 gap-4 flex-1 min-w-0">
      <CloudModelCard />
      <LocalModelCard />
    </div>
  </div>
</template>

<style scoped>
.llm {
}
</style>
