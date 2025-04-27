<script lang="ts" setup>
import { useRouter } from 'vue-router';
import { LLM_PROVIDERS } from '@/constants/LLMProviders';
import { Button } from '@/components/ui/button'

const router = useRouter();

// 跳转到云端模型配置页面
const goToCloudModel = () => {
  router.push('/llm/cloud');
};
</script>

<template>
  <div class="bg-base-200/50 border-base-300/50 border min-w-0 rounded-xl p-4 relative backdrop-blur-md h-full flex flex-col">
    <div class="flex justify-between items-center mb-3">
      <h2>云端模型</h2>
      <span class="bg-[rgb(239.8,248.9,235.3)] text-[#67c23a] px-2 py-1 rounded text-xs">推荐</span>
    </div>
    <div class="flex flex-col gap-4">
      <p class="text-base-content/80 text-sm">快速集成主流云端模型接口，立即开始使用</p>

      <!-- 特性标签 -->
      <div class="flex flex-wrap gap-2">
        <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
          <span>🚀</span>
          即刻使用
        </div>
        <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
          <span>🔌</span>
          简单配置
        </div>
        <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
          <span>🎯</span>
          功能完整
        </div>
      </div>

      <!-- API提供商图标 -->
      <div class="flex gap-4 mb-6 flex-wrap h-40 overflow-auto">
        <div v-for="provider in LLM_PROVIDERS" :key="provider.id" class="bg-base-300/50 text-base-content rounded-lg p-2 flex flex-col items-center min-w-[80px]">
          <div class="bg-base-100/50 w-10 h-10 flex items-center justify-center rounded-lg">
            <img :src="provider.icon" :alt="provider.name" class="w-6 h-6 opacity-90">
          </div>
          <span class="text-base-content/90 text-xs mt-1 text-center truncate w-full" :title="provider.name">{{ provider.name }}</span>
        </div>
      </div>

      <!-- API密钥提示 -->
      <div class="bg-info/20 text-info flex items-center gap-2 p-3 rounded-lg mb-6 text-sm">
        <span>ℹ️</span>
        <span>需要您自行提供API密钥</span>
      </div>

      <Button  @click="goToCloudModel">开始配置</Button>
    </div>
  </div>
</template>
