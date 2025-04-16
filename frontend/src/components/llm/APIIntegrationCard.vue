<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import OpenAIIcon from '../../assets/icons/placeholder-openai.svg';
import ClaudeIcon from '../../assets/icons/placeholder-claude.svg';
import GeminiIcon from '../../assets/icons/placeholder-gemini.svg';

const router = useRouter();

// API提供商列表
const apiProviders = ref([
  { name: 'OpenAI', icon: OpenAIIcon },
  { name: 'Claude', icon: ClaudeIcon },
  { name: 'Gemini', icon: GeminiIcon }
]);

// 跳转到API配置页面
const goToApiConfig = () => {
  router.push('/api');
};
</script>

<template>
  <div class="bg-base-200/50 border-base-300/50 border min-w-0 rounded-xl p-4 relative backdrop-blur-md h-full flex flex-col">
    <div class="flex justify-between items-center mb-3">
      <h2>API集成</h2>
      <span class="bg-primary/20 text-primary px-2 py-1 rounded text-xs">推荐</span>
    </div>
    <div class="flex flex-col gap-4">
      <p class="text-base-content/80 text-sm">快速集成主流AI服务商的API，立即开始使用</p>

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
      <div class="flex gap-4 mb-6">
        <div v-for="provider in apiProviders" :key="provider.name" class="bg-base-300/50 text-base-content hover:bg-base-100/50 rounded-lg p-2 transition-all duration-200">
          <div class="bg-base-100/50 w-10 h-10 flex items-center justify-center rounded-lg">
            <img :src="provider.icon" :alt="provider.name" class="w-6 h-6 opacity-90">
          </div>
          <span class="text-base-content/90 text-xs">{{ provider.name }}</span>
        </div>
        <div class="bg-base-300/50 text-base-content hover:bg-base-100/50 rounded-lg p-2 transition-all duration-200">
          <div class="bg-base-100/50 w-10 h-10 flex items-center justify-center rounded-lg text-base-content/70 text-sm">
            <span>+9</span>
          </div>
          <span class="text-base-content/90 text-xs">更多</span>
        </div>
      </div>

      <!-- API密钥提示 -->
      <div class="bg-info/20 text-info flex items-center gap-2 p-3 rounded-lg mb-6 text-sm">
        <span>ℹ️</span>
        <span>需要您自行提供API密钥</span>
      </div>

      <button class="bg-primary text-primary-content hover:bg-primary/90 w-full py-3 border-none rounded-lg text-sm cursor-pointer transition-all duration-200" @click="goToApiConfig">开始配置</button>
    </div>
  </div>
</template>
