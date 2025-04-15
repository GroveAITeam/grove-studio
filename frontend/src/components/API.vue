<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { Icon } from '@iconify/vue';

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

  currentApi.value = api;
  // 生成模拟使用数据
  usageData.value = generateMockUsageStats();
  showUsage.value = true;
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
  <div class="container">
    <header>
      <h1>API管理</h1>
      <p class="header-desc">管理您的API密钥</p>
    </header>

    <div class="quick-help">
      <div class="help-icon">💡</div>
      <div class="help-content">
        <p><span class="emphasis">自定义API</span>允许您使用第三方AI服务。请前往服务商官网获取API密钥，并在此页面进行设置。</p>
        <p>所有API密钥均存储在您的本地设备，不会上传至Grove服务器，确保您的账户安全。</p>
      </div>
    </div>

    <div class="privacy-alert">
      <span class="icon">🔒</span>
      <div class="privacy-message">
        <strong>隐私提示：</strong>第三方API服务可能会收集和存储您的数据。对于敏感数据，我们强烈推荐使用本地私有化模型，确保您的数据始终在本地处理，不经过任何外部服务器。
      </div>
    </div>

    <main>
      <!-- API列表和添加按钮容器 -->
      <div class="api-container">
        <!-- 添加API按钮 - 移到顶部 -->
        <button class="add-api-btn" @click="showAddApiForm">
          <span class="add-icon">+</span>
          <span>添加新API</span>
        </button>

        <!-- API列表区域 - 添加最大高度和滚动 -->
        <div class="api-list">
          <!-- 空状态显示 -->
          <div class="empty-state" v-if="apis.length === 0">
            <div class="empty-icon">🔑</div>
            <p>您还没有添加任何API密钥</p>
          </div>

          <!-- API列表 -->
          <div v-for="api in apis" :key="api.id" class="api-item">
            <div class="api-info">
              <img :src="getProviderIcon(api.provider)" class="api-logo" :alt="api.provider">
              <div class="api-details">
                <span class="api-name">{{ api.name }}</span>
                <span class="api-provider">{{ api.provider }}</span>
              </div>
            </div>
            <div class="api-actions">
              <button class="action-btn" @click="editApi(api.id)">
                <span>编辑</span>
              </button>
              <button class="action-btn delete-btn" @click="deleteApi(api.id)">
                <span>删除</span>
              </button>
              <button class="action-btn" @click="viewUsageStats(api.id)">
                <span>使用统计</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 添加/编辑API表单 弹窗 -->
      <Teleport to="body">
        <div class="form-overlay" v-if="showForm" @click.self="hideAddApiForm"></div>
        <div class="add-api-form" v-if="showForm">
          <form @submit.prevent="handleFormSubmit">
            <div class="form-group provider-options">
              <!-- API提供商选项 -->
              <div class="provider-grid">
                <div
                  v-for="provider in providers"
                  :key="provider.id"
                  class="provider-option"
                  :class="{ selected: selectedProvider === provider.id }"
                  :data-provider="provider.id"
                  @click="selectProvider(provider.id)"
                >
                  <img :src="provider.icon" :alt="provider.name">
                  <span>{{ provider.name }}</span>
                </div>
              </div>
            </div>

            <div class="form-group">
              <label for="apiName">名称</label>
              <input type="text" id="apiName" v-model="formData.name" required placeholder="为这个API起个名字">
            </div>

            <div class="form-group">
              <label for="apiKey">API密钥</label>
              <div class="input-group">
                <input :type="showPassword ? 'text' : 'password'" id="apiKey" v-model="formData.apiKey" required>
                <button type="button" class="toggle-password" @click="togglePasswordVisibility">
                  <span class="show-icon" :style="{ opacity: showPassword ? '1' : '0.5' }">👁️</span>
                </button>
              </div>
              <div class="api-help-text" v-if="selectedProvider && providerHelp[selectedProvider]">
                <div class="provider-help">
                  <div class="help-header">
                    <span class="help-icon">ℹ️</span>
                    <span class="help-title">如何获取API密钥</span>
                  </div>
                  <div class="help-content">
                    <p class="help-description">{{ providerHelp[selectedProvider].description }}</p>
                    <div class="help-steps" v-html="providerHelp[selectedProvider].helpText"></div>
                  </div>
                </div>
              </div>
            </div>

            <div class="form-group">
              <div class="advanced-toggle" @click="toggleAdvancedOptions">
                <span>{{ showAdvanced ? '隐藏高级选项' : '显示高级选项' }}</span>
                <span class="arrow-icon">▾</span>
              </div>
            </div>

            <div class="advanced-options" v-if="showAdvanced">
              <div class="form-group">
                <label for="defaultModel">默认模型</label>
                <select id="defaultModel" v-model="formData.defaultModel">
                  <option v-for="model in availableModels" :key="model" :value="model">{{ model }}</option>
                </select>
              </div>

              <div class="form-group">
                <label for="baseUrl">API基础URL</label>
                <input type="url" id="baseUrl" v-model="formData.baseUrl" placeholder="可选，用于自定义API端点">
              </div>
            </div>

            <div class="form-actions">
              <button type="button" class="secondary-btn" @click="hideAddApiForm">取消</button>
              <button type="submit" class="primary-btn">保存</button>
            </div>
          </form>
        </div>
      </Teleport>

      <!-- 使用统计弹窗 -->
      <Teleport to="body">
        <div class="form-overlay" v-if="showUsage" @click.self="hideUsageStats"></div>
        <div class="add-api-form" v-if="showUsage && currentApi">
          <div class="usage-stats">
            <div class="usage-header">
              <h2 class="usage-title">{{ currentApi.name }} 使用统计</h2>
              <div class="usage-period">过去30天</div>
            </div>

            <div class="usage-metrics">
              <div class="usage-metric">
                <div class="metric-value">{{ usageData.tokensUsed.toLocaleString() }}</div>
                <div class="metric-label">令牌使用量</div>
              </div>
              <div class="usage-metric">
                <div class="metric-value">${{ usageData.costEstimate.toFixed(2) }}</div>
                <div class="metric-label">估计费用</div>
              </div>
            </div>

            <div class="usage-progress">
              <div class="progress-label">配额使用：{{ usageData.quotaPercentage }}%</div>
              <div class="progress-bar-container">
                <div
                  class="progress-bar"
                  :class="getProgressBarClass(usageData.quotaPercentage)"
                  :style="{ width: usageData.quotaPercentage + '%' }"
                ></div>
              </div>
            </div>

            <div class="form-actions">
              <button type="button" class="secondary-btn" @click="hideUsageStats">关闭</button>
            </div>
          </div>
        </div>
      </Teleport>
    </main>
  </div>
</template>

<style scoped>
/* API管理页面特定样式 */
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
}

header {
  margin-bottom: 24px;
}

h1 {
  margin-bottom: 8px;
  font-size: 28px;
  font-weight: 600;
  color: var(--text-color);
}

.header-desc {
  color: var(--text-secondary);
  font-size: 16px;
}

/* 快速帮助 */
.quick-help {
  display: flex;
  gap: 16px;
  padding: 16px;
  background-color: rgba(59, 130, 246, 0.1);
  border-radius: 8px;
  margin-bottom: 24px;
}

.help-icon {
  font-size: 24px;
}

.help-content {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.emphasis {
  font-weight: 600;
}

/* 隐私提示 */
.privacy-alert {
  display: flex;
  gap: 16px;
  padding: 16px;
  background-color: rgba(245, 158, 11, 0.1);
  border-radius: 8px;
  margin-bottom: 24px;
}

.icon {
  font-size: 24px;
}

.privacy-message {
  font-size: 14px;
  line-height: 1.5;
}

.privacy-message strong {
  font-weight: 600;
}

/* API列表区域 */
.api-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
  margin-bottom: 24px;
}

.add-api-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  background-color: var(--primary-color, #1a56db);
  color: white;
  border: none;
  border-radius: 6px;
  padding: 8px 16px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
  align-self: flex-start;
}

.add-api-btn:hover {
  background-color: var(--primary-hover, #0e4dbf);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.add-icon {
  font-size: 16px;
  font-weight: bold;
}

/* 空状态样式 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  background-color: rgba(0, 0, 0, 0.03);
  border-radius: 8px;
  text-align: center;
}

.empty-icon {
  font-size: 36px;
  margin-bottom: 16px;
  opacity: 0.7;
}

.empty-state p {
  color: var(--text-secondary, #6b7280);
  margin-bottom: 20px;
}

/* API列表样式 */
.api-list {
  max-height: 500px;
  overflow-y: auto;
  padding-right: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.api-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px;
  background-color: rgba(0, 0, 0, 0.02);
  border-radius: 8px;
  border: 1px solid var(--border-color, #e5e7eb);
  transition: all 0.3s ease;
}

.api-item:hover {
  background-color: rgba(0, 0, 0, 0.04);
}

.api-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.api-logo {
  width: 40px;
  height: 40px;
  object-fit: contain;
}

.api-details {
  display: flex;
  flex-direction: column;
}

.api-name {
  font-weight: 500;
  color: var(--text-color, #1f2937);
  font-size: 16px;
}

.api-provider {
  color: var(--text-secondary, #6b7280);
  font-size: 14px;
}

.api-actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  background: none;
  border: none;
  color: var(--text-secondary, #6b7280);
  cursor: pointer;
  padding: 6px;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.action-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
  color: var(--text-color, #1f2937);
}

.delete-btn:hover {
  color: var(--error-color, #dc2626);
}

/* 添加API表单 */
.add-api-form {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: white;
  padding: 24px;
  border-radius: 12px;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.1);
  width: 90%;
  max-width: 600px;
  max-height: 90vh;
  overflow-y: auto;
  z-index: 1000;
}

.form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  z-index: 999;
}

.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: var(--text-color, #1f2937);
}

input[type="text"],
input[type="password"],
input[type="url"],
select {
  width: 100%;
  padding: 10px;
  border: 1px solid var(--border-color, #e5e7eb);
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.3s;
}

input[type="text"]:focus,
input[type="password"]:focus,
input[type="url"]:focus,
select:focus {
  border-color: var(--primary-color, #1a56db);
  outline: none;
}

.input-group {
  position: relative;
  display: flex;
}

.toggle-password {
  background: none;
  border: none;
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}

.show-icon {
  font-size: 16px;
  opacity: 0.5;
  transition: opacity 0.2s;
}

/* 提供商选项 */
.provider-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 16px;
}

.provider-option {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  border: 1px solid var(--border-color, #e5e7eb);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
}

.provider-option img {
  width: 32px;
  height: 32px;
  object-fit: contain;
}

.provider-option span {
  font-size: 14px;
}

.provider-option:hover {
  background-color: rgba(0, 0, 0, 0.02);
}

.provider-option.selected {
  border-color: var(--primary-color, #1a56db);
  background-color: rgba(59, 130, 246, 0.05);
}

/* 高级选项 */
.advanced-toggle {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 0;
  cursor: pointer;
  color: var(--text-secondary, #6b7280);
}

.arrow-icon {
  transition: transform 0.3s;
}

.advanced-options {
  margin-top: 16px;
  padding-top: 16px;
  border-top: 1px solid var(--border-color, #e5e7eb);
}

/* 表单按钮 */
.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 24px;
}

.secondary-btn,
.primary-btn {
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s;
}

.secondary-btn {
  background: none;
  border: 1px solid var(--border-color, #e5e7eb);
  color: var(--text-color, #1f2937);
}

.primary-btn {
  background-color: var(--primary-color, #1a56db);
  border: none;
  color: white;
}

.secondary-btn:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.primary-btn:hover {
  background-color: var(--primary-hover, #0e4dbf);
}

/* 提供商帮助 */
.provider-help {
  margin-top: 12px;
  border-radius: 8px;
  border: 1px solid var(--border-color, #e5e7eb);
  overflow: hidden;
}

.help-header {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px;
  background-color: rgba(0, 0, 0, 0.02);
  border-bottom: 1px solid var(--border-color, #e5e7eb);
}

.help-icon {
  font-size: 16px;
}

.help-title {
  font-weight: 500;
  font-size: 14px;
}

.help-content {
  padding: 12px;
}

.help-steps {
  margin-top: 12px;
  font-size: 14px;
}

.help-steps ol {
  padding-left: 24px;
}

.help-steps li {
  margin-bottom: 8px;
}

.help-steps a {
  color: var(--primary-color, #1a56db);
  text-decoration: none;
}

/* 使用统计 */
.usage-stats {
  padding: 16px;
}

.usage-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.usage-title {
  font-size: 18px;
  font-weight: 600;
}

.usage-period {
  font-size: 14px;
  color: var(--text-secondary, #6b7280);
}

.usage-metrics {
  display: flex;
  gap: 24px;
  margin-bottom: 24px;
}

.usage-metric {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.02);
  border-radius: 8px;
  padding: 16px;
}

.metric-value {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 4px;
}

.metric-label {
  font-size: 14px;
  color: var(--text-secondary, #6b7280);
}

.usage-progress {
  margin-bottom: 24px;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  margin-bottom: 8px;
  font-size: 14px;
}

.progress-bar-container {
  height: 8px;
  background-color: rgba(0, 0, 0, 0.05);
  border-radius: 4px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background-color: var(--primary-color, #1a56db);
  border-radius: 4px;
  transition: width 0.3s ease;
}

.progress-bar.warning {
  background-color: #f59e0b;
}

.progress-bar.danger {
  background-color: #dc2626;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .provider-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .api-info {
    flex-direction: column;
    align-items: flex-start;
  }

  .form-actions {
    flex-direction: column;
  }

  .form-actions button {
    width: 100%;
  }
}

/* 提示消息 */
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
