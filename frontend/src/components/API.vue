<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { Icon } from '@iconify/vue';
import {useRouter} from "vue-router";

const router = useRouter();

// 类型定义
interface ApiData {
  id: string;
  name: string;
  provider: string;
  apiKey: string;
  defaultModel: string;
  baseUrl: string;
}

interface UsageData {
  tokensUsed: number;
  costEstimate: number;
  quotaPercentage: number;
}

interface FormData {
  name: string;
  apiKey: string;
  provider: string;
  defaultModel: string;
  baseUrl: string;
}

interface ProviderInfo {
  description: string;
  helpText: string;
  models: string[];
}

interface ProviderHelp {
  [key: string]: ProviderInfo;
}

// 状态管理
const showForm = ref(false);
const showAdvanced = ref(false);
const showPassword = ref(false);
const showUsage = ref(false);
const selectedProvider = ref<string>('');
const currentApi = ref<ApiData | null>(null);
const editingId = ref<string | null>(null);
const apis = ref<ApiData[]>([]);

// 表单数据
const formData = ref<FormData>({
  name: '',
  apiKey: '',
  provider: '',
  defaultModel: '',
  baseUrl: ''
});

// 使用统计数据
const usageData = ref<UsageData>({
  tokensUsed: 0,
  costEstimate: 0,
  quotaPercentage: 0
});

// API提供商
const providers = [
  { id: 'openai', name: 'OpenAI', icon: '/assets/images/providers/openai.svg' },
  { id: 'anthropic', name: 'Anthropic', icon: '/assets/images/providers/anthropic.svg' },
  { id: 'google', name: 'Google AI', icon: '/assets/images/providers/google.svg' },
  { id: 'mistral', name: 'Mistral AI', icon: '/assets/images/providers/mistral.svg' },
  { id: 'cohere', name: 'Cohere', icon: '/assets/images/providers/cohere.svg' },
  { id: 'azure', name: 'Azure OpenAI', icon: '/assets/images/providers/azure.svg' }
];

// API提供商帮助信息
const providerHelp: ProviderHelp = {
  openai: {
    description: '使用OpenAI API访问GPT-3.5, GPT-4等模型。',
    helpText: `
      <ol>
        <li>访问 <a href="https://platform.openai.com/account/api-keys" target="_blank">OpenAI API密钥页面</a></li>
        <li>创建新的API密钥</li>
        <li>复制API密钥并粘贴到此处</li>
      </ol>
    `,
    models: ['gpt-4', 'gpt-4-turbo', 'gpt-3.5-turbo']
  },
  anthropic: {
    description: '使用Anthropic API访问Claude系列模型。',
    helpText: `
      <ol>
        <li>访问 <a href="https://console.anthropic.com/account/keys" target="_blank">Anthropic控制台</a></li>
        <li>创建新的API密钥</li>
        <li>复制API密钥并粘贴到此处</li>
      </ol>
    `,
    models: ['claude-3-opus', 'claude-3-sonnet', 'claude-3-haiku']
  },
  google: {
    description: '使用Google AI Studio API访问Gemini系列模型。',
    helpText: `
      <ol>
        <li>访问 <a href="https://makersuite.google.com/app/apikey" target="_blank">Google AI Studio</a></li>
        <li>创建新的API密钥</li>
        <li>复制API密钥并粘贴到此处</li>
      </ol>
    `,
    models: ['gemini-pro', 'gemini-ultra']
  },
  mistral: {
    description: '使用Mistral AI访问Mistral系列模型。',
    helpText: `
      <ol>
        <li>访问 <a href="https://console.mistral.ai/api-keys/" target="_blank">Mistral AI控制台</a></li>
        <li>创建新的API密钥</li>
        <li>复制API密钥并粘贴到此处</li>
      </ol>
    `,
    models: ['mistral-small', 'mistral-medium', 'mistral-large']
  },
  cohere: {
    description: '使用Cohere API访问Command系列模型。',
    helpText: `
      <ol>
        <li>访问 <a href="https://dashboard.cohere.ai/api-keys" target="_blank">Cohere控制台</a></li>
        <li>创建新的API密钥</li>
        <li>复制API密钥并粘贴到此处</li>
      </ol>
    `,
    models: ['command', 'command-light', 'command-r', 'command-r-plus']
  },
  azure: {
    description: '使用Azure OpenAI服务访问各种OpenAI模型。',
    helpText: `
      <ol>
        <li>登录到 <a href="https://portal.azure.com/" target="_blank">Azure门户</a></li>
        <li>找到您的Azure OpenAI资源</li>
        <li>在"密钥和终结点"部分获取密钥</li>
        <li>复制API密钥并粘贴到此处</li>
        <li>在高级选项中设置您的API基础URL</li>
      </ol>
    `,
    models: ['gpt-4', 'gpt-35-turbo', 'dall-e-3']
  }
};

// 根据选择的提供商获取可用模型
const availableModels = computed(() => {
  if (!selectedProvider.value || !providerHelp[selectedProvider.value]) {
    return [];
  }
  return providerHelp[selectedProvider.value].models;
});

// 获取提供商图标
function getProviderIcon(providerId: string): string {
  const provider = providers.find(p => p.id === providerId);
  return provider ? provider.icon : '/assets/images/providers/default.svg';
}

// 显示添加API表单
function showAddApiForm(): void {
  showForm.value = true;
  resetForm();
}

// 隐藏添加API表单
function hideAddApiForm(): void {
  showForm.value = false;
  resetForm();
}

// 重置表单
function resetForm(): void {
  formData.value = {
    name: '',
    apiKey: '',
    provider: '',
    defaultModel: '',
    baseUrl: ''
  };
  selectedProvider.value = '';
  editingId.value = null;
  showAdvanced.value = false;
}

// 切换密码可见性
function togglePasswordVisibility(): void {
  showPassword.value = !showPassword.value;
}

// 切换高级选项显示
function toggleAdvancedOptions(): void {
  showAdvanced.value = !showAdvanced.value;
}

// 选择提供商
function selectProvider(providerId: string): void {
  selectedProvider.value = providerId;
  formData.value.provider = providerId;

  // 设置默认模型
  if (providerHelp[providerId] && providerHelp[providerId].models.length > 0) {
    formData.value.defaultModel = providerHelp[providerId].models[0];
  }
}

// 处理表单提交
function handleFormSubmit(): void {
  if (!selectedProvider.value) {
    alert('请选择API提供商');
    return;
  }

  const apiData: ApiData = {
    id: editingId.value || crypto.randomUUID(),
    name: formData.value.name,
    provider: selectedProvider.value,
    apiKey: formData.value.apiKey,
    defaultModel: formData.value.defaultModel,
    baseUrl: formData.value.baseUrl || ''
  };

  if (editingId.value) {
    // 更新现有API
    const index = apis.value.findIndex(api => api.id === editingId.value);
    if (index !== -1) {
      apis.value[index] = apiData;
    }
  } else {
    // 添加新API
    apis.value.push(apiData);
  }

  // 保存到本地存储
  saveApis();

  // 关闭表单
  hideAddApiForm();

  // 显示成功提示
  showToast('API密钥已保存');
}

// 删除API
function deleteApi(id: string): void {
  if (confirm('确定要删除这个API密钥吗？')) {
    apis.value = apis.value.filter(api => api.id !== id);
    saveApis();
    showToast('API密钥已删除');
  }
}

// 编辑API
function editApi(id: string): void {
  const api = apis.value.find(api => api.id === id);
  if (!api) return;

  editingId.value = id;
  selectedProvider.value = api.provider;
  formData.value = {
    name: api.name,
    apiKey: api.apiKey,
    provider: api.provider,
    defaultModel: api.defaultModel,
    baseUrl: api.baseUrl || ''
  };

  showForm.value = true;
}

// 查看使用统计
function viewUsageStats(id: string): void {
  const api = apis.value.find(api => api.id === id);
  if (!api) return;

  // 跳转到使用统计页面，并传递API ID
  router.push(`/api-usage?id=${id}`);
}

// 隐藏使用统计
function hideUsageStats(): void {
  showUsage.value = false;
  currentApi.value = null;
}

// 获取进度条样式类
function getProgressBarClass(percentage: number): string {
  if (percentage > 90) return 'danger';
  if (percentage > 70) return 'warning';
  return '';
}

// 生成模拟使用统计数据
function generateMockUsageStats(): UsageData {
  const tokensUsed = Math.floor(Math.random() * 500000) + 10000;
  const costEstimate = tokensUsed * 0.000002;
  const quotaPercentage = Math.floor(Math.random() * 100);

  return {
    tokensUsed,
    costEstimate,
    quotaPercentage
  };
}

// 显示提示消息
function showToast(message: string): void {
  const toast = document.createElement('div');
  toast.className = 'toast';
  toast.textContent = message;
  document.body.appendChild(toast);

  setTimeout(() => {
    toast.classList.add('show');
  }, 10);

  setTimeout(() => {
    toast.classList.remove('show');
    setTimeout(() => toast.remove(), 300);
  }, 3000);
}

// 保存APIs到本地存储
function saveApis(): void {
  localStorage.setItem('grove_api_keys', JSON.stringify(apis.value));
}

// 从本地存储加载APIs
function loadApis(): void {
  const savedApis = localStorage.getItem('grove_api_keys');
  if (savedApis) {
    try {
      apis.value = JSON.parse(savedApis);
    } catch (e) {
      console.error('无法解析保存的API数据', e);
      apis.value = [];
    }
  }
}

// 组件挂载时加载APIs
onMounted(() => {
  loadApis();
});
</script>

<template>
  <div class="container mx-auto px-5 max-w-6xl">
    <header class="mb-6">
      <h1 class="text-2xl font-semibold mb-2 text-base-content">API管理</h1>
      <p class="text-base text-base-content/70">管理您的API密钥</p>
    </header>

    <div class="bg-primary/10 rounded-lg p-4 mb-6 flex gap-4">
      <div class="text-2xl">💡</div>
      <div class="flex flex-col gap-2">
        <p><span class="font-semibold">自定义API</span>允许您使用第三方AI服务。请前往服务商官网获取API密钥，并在此页面进行设置。</p>
        <p>所有API密钥均存储在您的本地设备，不会上传至Grove服务器，确保您的账户安全。</p>
      </div>
    </div>

    <div class="bg-warning/10 rounded-lg p-4 mb-6 flex gap-4">
      <span class="text-2xl">🔒</span>
      <div class="text-sm leading-relaxed">
        <strong>隐私提示：</strong>第三方API服务可能会收集和存储您的数据。对于敏感数据，我们强烈推荐使用本地私有化模型，确保您的数据始终在本地处理，不经过任何外部服务器。
      </div>
    </div>

    <main>
      <!-- API列表和添加按钮容器 -->
      <div class="flex flex-col gap-5 mb-6">
        <!-- 添加API按钮 - 移到顶部 -->
        <button class="btn btn-primary self-start flex items-center gap-2" @click="showAddApiForm">
          <span class="text-base font-bold">+</span>
          <span>添加新API</span>
        </button>

        <!-- API列表区域 - 添加最大高度和滚动 -->
        <div class="flex flex-col gap-3 max-h-[500px] overflow-y-auto pr-4">
          <!-- 空状态显示 -->
          <div class="flex flex-col items-center justify-center py-10 px-5 bg-base-200/30 rounded-lg text-center" v-if="apis.length === 0">
            <div class="text-3xl opacity-70 mb-4">🔑</div>
            <p class="text-base-content/70 mb-5">您还没有添加任何API密钥</p>
          </div>

          <!-- API列表 -->
          <div v-for="api in apis" :key="api.id" class="flex items-center justify-between p-4 bg-base-200/20 rounded-lg border border-base-300 transition-all hover:bg-base-200/40">
            <div class="flex items-center gap-3">
              <img :src="getProviderIcon(api.provider)" class="w-10 h-10 object-contain" :alt="api.provider">
              <div class="flex flex-col">
                <span class="font-medium text-base text-base-content">{{ api.name }}</span>
                <span class="text-sm text-base-content/70">{{ api.provider }}</span>
              </div>
            </div>
            <div class="flex gap-2">
              <button class="btn btn-ghost btn-sm text-base-content/70 hover:text-base-content hover:bg-base-300/50" @click="editApi(api.id)">
                <span>编辑</span>
              </button>
              <button class="btn btn-ghost btn-sm text-base-content/70 hover:text-error hover:bg-base-300/50" @click="deleteApi(api.id)">
                <span>删除</span>
              </button>
              <button class="btn btn-ghost btn-sm text-base-content/70 hover:text-base-content hover:bg-base-300/50" @click="viewUsageStats(api.id)">
                <span>使用统计</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 添加/编辑API表单 弹窗 -->
      <Teleport to="body">
        <div class="fixed inset-0 bg-black/50 z-50" v-if="showForm" @click.self="hideAddApiForm"></div>
        <div class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-xl shadow-lg w-[90%] max-w-[600px] max-h-[90vh] overflow-y-auto z-50" v-if="showForm">
          <form @submit.prevent="handleFormSubmit">
            <div class="mb-5">
              <!-- API提供商选项 -->
              <div class="grid grid-cols-3 gap-3 mb-4 md:grid-cols-3 sm:grid-cols-2">
                <div
                  v-for="provider in providers"
                  :key="provider.id"
                  class="flex flex-col items-center gap-2 p-4 border border-base-300 rounded-lg cursor-pointer transition-all hover:bg-base-200/20"
                  :class="{ 'border-primary bg-primary/5': selectedProvider === provider.id }"
                  :data-provider="provider.id"
                  @click="selectProvider(provider.id)"
                >
                  <img :src="provider.icon" class="w-8 h-8 object-contain" :alt="provider.name">
                  <span class="text-sm">{{ provider.name }}</span>
                </div>
              </div>
            </div>

            <div class="mb-5">
              <label for="apiName" class="block mb-2 font-medium text-base-content">名称</label>
              <input type="text" id="apiName" v-model="formData.name" required placeholder="为这个API起个名字" class="input input-bordered w-full">
            </div>

            <div class="mb-5">
              <label for="apiKey" class="block mb-2 font-medium text-base-content">API密钥</label>
              <div class="relative flex">
                <input :type="showPassword ? 'text' : 'password'" id="apiKey" v-model="formData.apiKey" required class="input input-bordered w-full pr-10">
                <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center justify-center" @click="togglePasswordVisibility">
                  <span class="text-base" :class="{ 'opacity-100': showPassword, 'opacity-50': !showPassword }">👁️</span>
                </button>
              </div>
              <div v-if="selectedProvider && providerHelp[selectedProvider]" class="mt-3">
                <div class="border border-base-300 rounded-lg overflow-hidden">
                  <div class="flex items-center gap-2 p-3 bg-base-200/20 border-b border-base-300">
                    <span class="text-base">ℹ️</span>
                    <span class="font-medium text-sm">如何获取API密钥</span>
                  </div>
                  <div class="p-3">
                    <p class="text-sm">{{ providerHelp[selectedProvider].description }}</p>
                    <div class="mt-3 text-sm" v-html="providerHelp[selectedProvider].helpText"></div>
                  </div>
                </div>
              </div>
            </div>

            <div class="mb-5">
              <div class="flex items-center justify-between py-2 cursor-pointer text-base-content/70" @click="toggleAdvancedOptions">
                <span>{{ showAdvanced ? '隐藏高级选项' : '显示高级选项' }}</span>
                <span class="transition-transform" :class="{ 'rotate-180': showAdvanced }">▾</span>
              </div>
            </div>

            <div v-if="showAdvanced" class="border-t border-base-300 pt-4 mt-4">
              <div class="mb-5">
                <label for="defaultModel" class="block mb-2 font-medium text-base-content">默认模型</label>
                <select id="defaultModel" v-model="formData.defaultModel" class="select select-bordered w-full">
                  <option v-for="model in availableModels" :key="model" :value="model">{{ model }}</option>
                </select>
              </div>

              <div class="mb-5">
                <label for="baseUrl" class="block mb-2 font-medium text-base-content">API基础URL</label>
                <input type="url" id="baseUrl" v-model="formData.baseUrl" placeholder="可选，用于自定义API端点" class="input input-bordered w-full">
              </div>
            </div>

            <div class="flex justify-end gap-3 mt-6 md:flex-row sm:flex-col">
              <button type="button" class="btn btn-outline" @click="hideAddApiForm">取消</button>
              <button type="submit" class="btn btn-primary">保存</button>
            </div>
          </form>
        </div>
      </Teleport>

      <!-- 使用统计弹窗 -->
      <Teleport to="body">
        <div class="fixed inset-0 bg-black/50 z-50" v-if="showUsage" @click.self="hideUsageStats"></div>
        <div class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-xl shadow-lg w-[90%] max-w-[600px] max-h-[90vh] overflow-y-auto z-50" v-if="showUsage && currentApi">
          <div class="p-4">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-lg font-semibold">{{ currentApi.name }} 使用统计</h2>
              <div class="text-sm text-base-content/70">过去30天</div>
            </div>

            <div class="flex gap-6 mb-6">
              <div class="flex-1 flex flex-col items-center bg-base-200/20 rounded-lg p-4">
                <div class="text-2xl font-semibold mb-1">{{ usageData.tokensUsed.toLocaleString() }}</div>
                <div class="text-sm text-base-content/70">令牌使用量</div>
              </div>
              <div class="flex-1 flex flex-col items-center bg-base-200/20 rounded-lg p-4">
                <div class="text-2xl font-semibold mb-1">${{ usageData.costEstimate.toFixed(2) }}</div>
                <div class="text-sm text-base-content/70">估计费用</div>
              </div>
            </div>

            <div class="mb-6">
              <div class="flex justify-between mb-2 text-sm">配额使用：{{ usageData.quotaPercentage }}%</div>
              <div class="h-2 bg-base-200/50 rounded-full overflow-hidden">
                <div
                  class="h-full rounded-full transition-all"
                  :class="{
                    'bg-primary': usageData.quotaPercentage <= 70,
                    'bg-warning': usageData.quotaPercentage > 70 && usageData.quotaPercentage <= 90,
                    'bg-error': usageData.quotaPercentage > 90
                  }"
                  :style="{ width: usageData.quotaPercentage + '%' }"
                ></div>
              </div>
            </div>

            <div class="flex justify-end">
              <button type="button" class="btn btn-outline" @click="hideUsageStats">关闭</button>
            </div>
          </div>
        </div>
      </Teleport>
    </main>
  </div>
</template>

<style scoped>
:global(.toast) {
  position: fixed;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%) translateY(100px);
  background-color: #333;
  color: white;
  padding: 12px 24px;
  border-radius: 8px;
  font-size: 14px;
  z-index: 1010;
  opacity: 0;
  transition: transform 0.3s, opacity 0.3s;
}

:global(.toast.show) {
  transform: translateX(-50%) translateY(0);
  opacity: 1;
}
</style>
