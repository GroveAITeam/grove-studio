<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits, computed } from 'vue';
import { LLM_PROVIDERS } from '../../constants/LLMProviders';

// 获取所有可用模型
const availableModels = computed(() => {
  const models = [];

  // 从所有提供商中获取模型
  LLM_PROVIDERS.forEach(provider => {
    provider.models.forEach(model => {
      models.push({
        id: model,
        name: model,
        provider: provider.name
      });
    });
  });

  return models;
});

const props = defineProps<{
  settings: {
    model: string,
    temperature: number,
    maxTokens: number,
    contextLength: number
  }
}>();

const emit = defineEmits<{
  'close': []
}>();

const temperatureFormatted = computed(() => {
  return Math.round(props.settings.temperature * 10) / 10;
});

// 回答长度范围
const responseLength = computed(() => {
  const length = props.settings.maxTokens;
  if (length <= 500) return '简短';
  if (length <= 1000) return '适中';
  if (length <= 2000) return '详细';
  return '完整';
});
</script>

<template>
  <div class="absolute right-6 mt-10 w-72 bg-base-200 border border-base-300 shadow-lg rounded-md z-10 p-3">
    <div class="flex justify-between items-center mb-3">
      <h3 class="font-medium text-sm text-base-content">对话设置</h3>
      <button @click="$emit('close')" class="text-base-content/70 hover:bg-base-300 p-1 rounded-md transition-colors">
        <Icon icon="material-symbols:close" class="text-base" />
      </button>
    </div>

    <div class="space-y-3">
      <div>
        <label class="block text-sm font-medium text-base-content mb-1">语言模型</label>
        <select v-model="settings.model" class="w-full bg-base-100 border border-base-300 rounded-md px-2 py-1.5 text-sm text-base-content">
          <option v-for="model in availableModels" :key="model.id" :value="model.id">
            {{ model.name }} ({{ model.provider }})
          </option>
        </select>
      </div>

      <div>
        <label class="block text-sm font-medium text-base-content mb-1">
          回答灵活度
          <span class="text-xs text-base-content/60">（{{ temperatureFormatted }}）</span>
        </label>
        <input type="range" v-model="settings.temperature" min="0" max="1" step="0.1"
               class="w-full h-2 bg-base-300 rounded-lg appearance-none cursor-pointer" />
        <div class="flex justify-between text-xs text-base-content/60 mt-1">
          <span>精确</span>
          <span>创造性</span>
        </div>
      </div>

      <div>
        <label class="block text-sm font-medium text-base-content mb-1">
          回答长度
          <span class="text-xs text-base-content/60">（{{ responseLength }}）</span>
        </label>
        <input type="range" v-model="settings.maxTokens" :min="500" :max="4000" step="500"
               class="w-full h-2 bg-base-300 rounded-lg appearance-none cursor-pointer" />
        <div class="flex justify-between text-xs text-base-content/60 mt-1">
          <span>简短</span>
          <span>完整</span>
        </div>
        <p class="text-xs text-base-content/60 mt-1">调整AI回答的详细程度，从简短到完整的回答</p>
      </div>
    </div>
  </div>
</template>
