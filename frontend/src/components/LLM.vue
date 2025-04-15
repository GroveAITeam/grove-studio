<script lang="ts" setup>
import { ref } from 'vue';
import OpenAIIcon from '../assets/icons/placeholder-openai.svg';
import ClaudeIcon from '../assets/icons/placeholder-claude.svg';
import GeminiIcon from '../assets/icons/placeholder-gemini.svg';

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
</script>

<template>
  <div class="llm-container">
    <!-- 顶部提示 -->
    <div class="top-alert">
      <div class="alert-icon">💡</div>
      <div class="alert-text">API集成支持对接主流AI服务商，快速开始使用。本地模型即将推出，将提供完全的数据隐私保护和离线使用体验。</div>
    </div>

    <!-- 卡片区域 -->
    <div class="cards-container">
      <!-- API集成卡片 -->
      <div class="card api-card">
        <div class="card-header">
          <h2>API集成</h2>
          <span class="recommend-badge">推荐</span>
        </div>
        <div class="card-content">
          <p class="description">快速集成主流AI服务商的API，立即开始使用</p>

          <!-- 特性标签 -->
          <div class="feature-tags">
            <div class="feature-tag">
              <span class="emoji">🚀</span>
              即刻使用
            </div>
            <div class="feature-tag">
              <span class="emoji">🔌</span>
              简单配置
            </div>
            <div class="feature-tag">
              <span class="emoji">🎯</span>
              功能完整
            </div>
          </div>

          <!-- API提供商图标 -->
          <div class="api-providers">
            <div v-for="provider in apiProviders" :key="provider.name" class="provider-item">
              <div class="provider-icon">
                <img :src="provider.icon" :alt="provider.name">
              </div>
              <span class="provider-name">{{ provider.name }}</span>
            </div>
            <div class="provider-item more">
              <div class="provider-icon more-icon">
                <span>+9</span>
              </div>
              <span class="provider-name">更多</span>
            </div>
          </div>

          <!-- API密钥提示 -->
          <div class="api-key-notice">
            <span class="notice-icon">ℹ️</span>
            <span>需要您自行提供API密钥</span>
          </div>

          <button class="primary-button">开始配置</button>
        </div>
      </div>

      <!-- 本地模型卡片 -->
      <div class="card local-card">
        <div class="coming-soon-mask">即将推出</div>
        <div class="card-header">
          <h2>本地模型</h2>
        </div>
        <div class="card-content">
          <p class="description">在您自己的设备上本地运行AI模型，完全保护数据隐私</p>

          <!-- 特性标签 -->
          <div class="feature-tags">
            <div class="feature-tag">
              <span class="emoji">🔒</span>
              数据完全私有
            </div>
            <div class="feature-tag">
              <span class="emoji">⚡️</span>
              低延迟响应
            </div>
            <div class="feature-tag">
              <span class="emoji">🔌</span>
              无需联网
            </div>
            <div class="feature-tag">
              <span class="emoji">🚀</span>
              一键安装
            </div>
            <div class="feature-tag">
              <span class="emoji">📚</span>
              多种开源模型
            </div>
          </div>

          <!-- 运行要求区域 -->
          <div class="requirements-section">
            <div class="requirements-header">
              <h3>运行要求</h3>
              <div class="size-tabs">
                <button
                  v-for="size in modelSizes"
                  :key="size.id"
                  :class="['size-tab', { active: size.active }]"
                  @click="switchModelSize(size.id)"
                >
                  {{ size.label }}
                </button>
              </div>
            </div>

            <div class="model-info">
              <div class="model-size">{{ currentModelInfo.size }}</div>
              <p class="model-desc">{{ currentModelInfo.description }}</p>

              <!-- 用例列表 -->
              <div class="use-cases">
                <div class="use-case">
                  <span class="use-case-icon">📝</span>
                  <span>文本整理与总结：整理会议记录、总结文档要点</span>
                </div>
                <div class="use-case">
                  <span class="use-case-icon">💬</span>
                  <span>日常对话：简单问答、基础知识咨询</span>
                </div>
                <div class="use-case">
                  <span class="use-case-icon">✍️</span>
                  <span>文案创作：编写简单文案、社交媒体内容</span>
                </div>
                <div class="use-case">
                  <span class="use-case-icon">🔍</span>
                  <span>知识库搜索：快速检索本地文档</span>
                </div>
              </div>

              <!-- 设备兼容性 -->
              <div class="compatibility">
                <div class="compatibility-header">
                  <span>您当前设备</span>
                  <span :class="['compatibility-status', currentModelInfo.compatibility.status]">
                    {{ currentModelInfo.compatibility.text }}
                  </span>
                </div>
                <div class="specs-list">
                  <div
                    v-for="(spec, index) in currentModelInfo.compatibility.specs"
                    :key="index"
                    class="spec-item"
                  >
                    <span class="spec-icon">{{ spec.icon }}</span>
                    <span class="spec-name">{{ spec.name }}</span>
                    <span :class="['spec-status', { compatible: spec.compatible }]">
                      {{ spec.compatible ? '✓' : '✗' }}
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <button class="secondary-button" disabled>敬请期待</button>
          <p class="coming-soon-note">我们正在开发本地模型功能，将第一时间通知您</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 容器基础样式 */
.llm-container {
  @apply bg-base-100 text-base-content;
  min-height: 100%;
  padding: 24px;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

/* 顶部提示样式 */
.top-alert {
  @apply bg-base-200/50 text-base-content border border-base-300/50;
  display: flex;
  align-items: center;
  gap: 12px;
  border-radius: 8px;
  padding: 12px 16px;
}

.alert-icon {
  font-size: 20px;
}

.alert-text {
  @apply text-base-content/90;
  font-size: 14px;
}

.title-section {
  text-align: center;
  margin-bottom: 32px;
}

.title-section h1 {
  font-size: 24px;
  font-weight: 600;
  color: #111827;
  margin-bottom: 8px;
}

.title-section p {
  font-size: 14px;
  color: #6B7280;
}

.cards-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  flex: 1;
  min-width: 0; /* 防止溢出 */
}

/* 卡片基础样式 */
.card {
  @apply bg-base-200/50 border-base-300/50 border;
  min-width: 0;
  border-radius: 12px;
  padding: 24px;
  position: relative;
  backdrop-filter: blur(8px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  position: relative;
}

.card-header h2 {
  @apply text-base-content;
  font-size: 20px;
  font-weight: 600;
}

/* 推荐标签 */
.recommend-badge {
  @apply bg-primary/20 text-primary;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.description {
  @apply text-base-content/80;
  font-size: 14px;
  margin-bottom: 24px;
}

.feature-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 24px;
}

/* 特性标签 */
.feature-tag {
  @apply bg-base-300/50 text-base-content;
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.api-providers {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

/* API提供商图标 */
.provider-item {
  @apply bg-base-300/50 text-base-content hover:bg-base-100/50;
  border-radius: 8px;
  padding: 8px;
  transition: all 0.2s;
}

.provider-icon {
  @apply bg-base-100/50;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
}

.provider-icon img {
  width: 24px;
  height: 24px;
  @apply opacity-90;
}

.more-icon {
  @apply text-base-content/70;
  font-size: 14px;
}

.provider-name {
  @apply text-base-content/90;
  font-size: 12px;
}

/* API密钥提示 */
.api-key-notice {
  @apply bg-info/20 text-info;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px;
  border-radius: 8px;
  margin-bottom: 24px;
  font-size: 14px;
}

/* 按钮样式 */
.primary-button {
  @apply bg-primary text-primary-content hover:bg-primary/90;
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.2s;
}

.secondary-button {
  @apply bg-base-300/50 text-base-content/50;
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  cursor: not-allowed;
}

/* 即将推出遮罩 */
.coming-soon-mask {
  @apply bg-base-300/90 text-primary;
  position: absolute;
  top: 24px;
  right: 24px;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 14px;
  font-weight: 600;
  z-index: 1;
  backdrop-filter: blur(4px);
}

/* 运行要求区域 */
.requirements-section {
  @apply bg-base-300/50;
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 24px;
}

.requirements-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.requirements-header h3 {
  @apply text-base-content font-semibold;
  font-size: 16px;
}

.size-tabs {
  display: flex;
  gap: 8px;
}

/* 大小选择标签 */
.size-tab {
  @apply bg-base-100/50 text-base-content border-base-300/50 hover:bg-base-200/50;
  padding: 6px 12px;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s;
}

.size-tab.active {
  @apply bg-primary text-primary-content border-primary;
}

.model-size {
  @apply bg-base-200/50 text-base-content;
  display: inline-block;
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
  margin-bottom: 8px;
}

.model-desc {
  @apply text-base-content/80;
  font-size: 14px;
  margin-bottom: 16px;
}

.use-cases {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 24px;
}

/* 用例列表 */
.use-case {
  @apply bg-base-100/50 text-base-content;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 8px;
  border-radius: 4px;
  font-size: 14px;
}

.compatibility {
  @apply bg-base-100/50;
  border-radius: 8px;
  padding: 16px;
}

.compatibility-header {
  @apply text-base-content;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
  font-size: 14px;
}

/* 兼容性状态 */
.compatibility-status.compatible {
  @apply bg-success/20 text-success;
}

.compatibility-status.incompatible {
  @apply bg-error/20 text-error;
}

.specs-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

/* 规格列表 */
.spec-item {
  @apply text-base-content;
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
}

.spec-status {
  @apply text-base-content/50;
  margin-left: auto;
}

.spec-status.compatible {
  @apply text-success;
}

/* 即将推出提示 */
.coming-soon-note {
  @apply text-base-content/60;
  text-align: center;
  font-size: 14px;
  font-style: italic;
  margin-top: 12px;
}
</style>
