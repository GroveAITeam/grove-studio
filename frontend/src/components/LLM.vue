<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import OpenAIIcon from '../assets/icons/placeholder-openai.svg';
import ClaudeIcon from '../assets/icons/placeholder-claude.svg';
import GeminiIcon from '../assets/icons/placeholder-gemini.svg';

const router = useRouter();

// 模型大小选项
const modelSizes = ref([
  { id: 'small', label: '轻量版', active: true },
  { id: 'medium', label: '标准版', active: false },
  { id: 'large', label: '增强版', active: false }
]);

// 当前选中的模型大小信息
const currentModelInfo = ref({
  size: '1-3B 参数',
  description: '适合日常对话和简单任务，对硬件要求较低',
  compatibility: {
    status: 'compatible',
    text: '兼容',
    specs: [
      { icon: '💻', name: 'CPU: i5 12代', compatible: true },
      { icon: '📊', name: '内存: 16GB', compatible: true },
      { icon: '🎮', name: '显卡: GTX 1060', compatible: true }
    ]
  }
});

// API提供商列表
const apiProviders = ref([
  { name: 'OpenAI', icon: OpenAIIcon },
  { name: 'Claude', icon: ClaudeIcon },
  { name: 'Gemini', icon: GeminiIcon }
]);

// 切换模型大小
const switchModelSize = (size: string) => {
  modelSizes.value.forEach(s => s.active = s.id === size);

  if (size === 'small') {
    currentModelInfo.value = {
      size: '1-3B 参数',
      description: '适合日常对话和简单任务，对硬件要求较低',
      compatibility: {
        status: 'compatible',
        text: '兼容',
        specs: [
          { icon: '💻', name: 'CPU: i5 12代', compatible: true },
          { icon: '📊', name: '内存: 16GB', compatible: true },
          { icon: '🎮', name: '显卡: GTX 1060', compatible: true }
        ]
      }
    };
  } else if (size === 'medium') {
    currentModelInfo.value = {
      size: '7-13B 参数',
      description: '适合更复杂的任务，能处理多种语言和领域知识',
      compatibility: {
        status: 'compatible',
        text: '兼容',
        specs: [
          { icon: '💻', name: 'CPU: i5 12代', compatible: true },
          { icon: '📊', name: '内存: 16GB', compatible: true },
          { icon: '🎮', name: '显卡: GTX 1060', compatible: true }
        ]
      }
    };
  } else {
    currentModelInfo.value = {
      size: '30B+ 参数',
      description: '适合高级任务，提供最高质量输出，但需要强大硬件',
      compatibility: {
        status: 'incompatible',
        text: '不兼容',
        specs: [
          { icon: '💻', name: 'CPU: i5 12代', compatible: true },
          { icon: '📊', name: '内存: 16GB', compatible: false },
          { icon: '🎮', name: '显卡: GTX 1060', compatible: false }
        ]
      }
    };
  }
};

// 跳转到API配置页面
const goToApiConfig = () => {
  router.push('/api');
};
</script>

<template>
  <div class="bg-base-100 text-base-content p-4 flex flex-col gap-4">
    <!-- 顶部提示 -->
    <div class="bg-base-200/50 text-base-content border border-base-300/50 flex items-center gap-3 rounded-lg p-2">
      <div class="text-lg">💡</div>
      <div class="text-base-content/90 text-sm">API集成支持对接主流AI服务商，快速开始使用。本地模型即将推出，将提供完全的数据隐私保护和离线使用体验。</div>
    </div>

    <!-- 卡片区域 -->
    <div class="grid grid-cols-2 gap-4 flex-1 min-w-0">
      <!-- API集成卡片 -->
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

      <!-- 本地模型卡片 -->
      <div class="bg-base-200/50 border-base-300/50 border min-w-0 rounded-xl p-4 relative backdrop-blur-md h-full flex flex-col">
        <div class="bg-base-300/90 text-primary absolute top-6 right-6 px-2 py-1 rounded text-sm font-semibold z-10 backdrop-blur">即将推出</div>
        <div class="flex justify-between items-center mb-3">
          <h2>本地模型</h2>
        </div>
        <div class="flex flex-col gap-4">
          <p class="text-base-content/80 text-sm">在您自己的设备上本地运行AI模型，完全保护数据隐私</p>

          <!-- 特性标签 -->
          <div class="flex flex-wrap gap-2">
            <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
              <span>🔒</span>
              数据完全私有
            </div>
            <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
              <span>⚡️</span>
              低延迟响应
            </div>
            <div class="bg-base-300/50 text-base-content flex items-center gap-1 px-2 py-1 rounded text-xs">
              <span>🔌</span>
              无需联网
            </div>
          </div>

          <!-- 简化的运行要求区域 -->
          <div class="bg-base-300/50 rounded-lg p-3">
            <div class="flex justify-between items-center mb-3">
              <h3>运行要求</h3>
              <div class="flex gap-1.5">
                <button
                  v-for="size in modelSizes"
                  :key="size.id"
                  :class="['bg-base-100/50 text-base-content hover:bg-base-200/50 px-2 py-1 rounded text-xs cursor-pointer', { 'bg-base-100': size.active }]"
                  @click="switchModelSize(size.id)"
                >
                  {{ size.label }}
                </button>
              </div>
            </div>

            <div>
              <div class="bg-base-200/50 text-base-content inline-block px-2 py-1 rounded text-xs mb-1.5">{{ currentModelInfo.size }}</div>
              <p class="text-base-content/80 text-sm mb-3">{{ currentModelInfo.description }}</p>

              <!-- 设备兼容性 -->
              <div class="bg-base-100/50 rounded-md p-3">
                <div class="flex items-center justify-between">
                  <span>您当前设备</span>
                  <span :class="[
                    currentModelInfo.compatibility.status === 'compatible' ? 'bg-success/20 text-success' : 'bg-error/20 text-error'
                  ]">
                    {{ currentModelInfo.compatibility.text }}
                  </span>
                </div>
                <div class="flex flex-col gap-1.5 mt-2">
                  <div
                    v-for="(spec, index) in currentModelInfo.compatibility.specs"
                    :key="index"
                    class="text-base-content flex items-center gap-2 text-sm"
                  >
                    <span>{{ spec.icon }}</span>
                    <span>{{ spec.name }}</span>
                    <span :class="[spec.compatible ? 'text-success' : 'text-error']">
                      {{ spec.compatible ? '✓' : '✗' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <button class="bg-base-300/50 text-base-content/50 w-full py-3 border-none rounded-lg text-sm cursor-not-allowed" disabled>敬请期待</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 所有样式已转换为Tailwind类，不再需要scoped CSS */
</style>
