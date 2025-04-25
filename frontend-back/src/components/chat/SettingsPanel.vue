<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits, computed, inject, ref, onMounted } from 'vue';
import { LLM_PROVIDERS } from '../../constants/LLMProviders';
import { GetCloudLLMModels } from '../../../wailsjs/go/main/App';

interface Model {
  id: string;
  name: string;
  provider: string; // 提供商ID
  provider_display: string; // 提供商显示名称
  config_name: string; // 配置名称
}

interface CloudModel {
  id: number;
  name: string; // 配置名称
  provider: string; // 提供商ID
  endpoint: string;
  api_key: string;
  enabled: boolean;
}

// 云端模型列表
const cloudModels = ref<CloudModel[]>([]);

// 加载云端模型
const loadCloudModels = async () => {
  try {
    const result = await GetCloudLLMModels(1, 100);
    cloudModels.value = result.items;
  } catch (error) {
    console.error('加载云端模型失败:', error);
  }
};

// 获取所有可用模型，按照provider+name分组
const availableModels = computed(() => {
  const models: Model[] = [];

  // 根据云端返回的启用的模型进行过滤
  cloudModels.value.forEach(cloudModel => {
    if (cloudModel.enabled) {
      // 找到对应的提供商信息
      const providerInfo = LLM_PROVIDERS.find(p => p.id === cloudModel.provider);

      if (providerInfo) {
        // 添加该提供商下的所有模型
        providerInfo.models.forEach(modelName => {
          models.push({
            // 组合ID：云端配置ID + 模型名，用于唯一标识
            id: `${cloudModel.id}:${modelName}`,
            name: modelName,
            provider: cloudModel.provider, // 使用原始provider ID
            provider_display: providerInfo.name, // 用于显示的提供商名称
            config_name: cloudModel.name
          });
        });
      }
    }
  });

  return models;
});

// 按提供商分组
const modelGroups = computed(() => {
  const groups: {[key: string]: Model[]} = {};

  availableModels.value.forEach(model => {
    const groupKey = `${model.config_name}`;
    if (!groups[groupKey]) {
      groups[groupKey] = [];
    }
    groups[groupKey].push(model);
  });

  return groups;
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
  if (length <= 1167) return '简短';
  if (length <= 2334) return '适中';
  if (length <= 3500) return '详细';
  return '完整';
});

// 修正输入事件类型处理
const handleTemperatureChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  props.settings.temperature = Number(target.value);
};

const handleMaxTokensChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  props.settings.maxTokens = Number(target.value);
};

const handleContextLengthChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  props.settings.contextLength = Number(target.value);
};

// 初始化
onMounted(async () => {
  await loadCloudModels();
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
          <optgroup v-for="(models, groupName) in modelGroups" :key="String(groupName)" :label="String(groupName)">
            <option v-for="model in models" :key="model.id" :value="model.id">
              {{ model.name }} ({{ model.provider }})
            </option>
          </optgroup>
        </select>
      </div>

      <div>
        <label class="block text-sm font-medium text-base-content mb-1">
          回答灵活度
          <span class="text-xs text-base-content/60">（{{ temperatureFormatted }}）</span>
        </label>
        <input type="range"
               :value="Number(settings.temperature)"
               @input="handleTemperatureChange"
               min="0" max="1" step="0.1"
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
        <input type="range"
               :value="Number(settings.maxTokens)"
               @input="handleMaxTokensChange"
               :min="500" :max="4001" :step="1167"
               class="w-full h-2 bg-base-300 rounded-lg appearance-none cursor-pointer" />
        <div class="flex justify-between text-xs text-base-content/60 mt-1">
          <span>简短</span>
          <span>适中</span>
          <span>详细</span>
          <span>完整</span>
        </div>
        <p class="text-xs text-base-content/60 mt-1">调整AI回答的详细程度，从简短到完整的回答</p>
      </div>

      <div>
        <label class="block text-sm font-medium text-base-content mb-1">
          上下文长度
          <span class="text-xs text-base-content/60">（{{ settings.contextLength }}轮对话）</span>
        </label>
        <input type="range"
               :value="Number(settings.contextLength)"
               @input="handleContextLengthChange"
               :min="1" :max="20" step="1"
               class="w-full h-2 bg-base-300 rounded-lg appearance-none cursor-pointer" />
        <div class="flex justify-between text-xs text-base-content/60 mt-1">
          <span>1轮</span>
          <span>20轮</span>
        </div>
        <p class="text-xs text-base-content/60 mt-1">调整AI能记住的对话轮数，过长会影响响应速度</p>
      </div>
    </div>
  </div>
</template>
