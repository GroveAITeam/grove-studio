<script lang="ts" setup>
import {ref, reactive, onMounted, computed, watch, nextTick} from 'vue'
import ConversationList from '../../components/chat/ConversationList.vue';
import ChatHeader from '../../components/chat/ChatHeader.vue';
import ChatMessages from '../../components/chat/ChatMessages.vue';
import ChatInput from '../../components/chat/ChatInput.vue';
import SettingsPanel from '../../components/chat/SettingsPanel.vue';
import OpenAI from "openai";
import {GetCloudLLMModels} from '../../../wailsjs/go/main/App';
import {useToast} from "../../utils/toast";
import {LLM_PROVIDERS} from '../../constants/LLMProviders';
import { EventsOn } from '../../../wailsjs/runtime';
import {StreamRequestMessage} from "../../../wailsjs/go/main/App";

const toast = useToast();
let client: OpenAI | null = null;
const SETTINGS_STORAGE_KEY = 'chat_settings';

// 加载可用的OpenAI API Key
const loadOpenAIApiKey = async () => {
  const result = await GetCloudLLMModels(1, 10);
  // 获取被启用的云端模型
  const enabledModel = result.items.find(model => model.enabled);

  if (enabledModel) {
    try {
      // 从localStorage加载设置
      loadSettingsFromLocalStorage();

      // 查找对应提供商的信息
      const providerInfo = LLM_PROVIDERS.find(p => p.id === enabledModel.provider);

      if (!providerInfo) {
        toast.error(`未找到提供商信息: ${enabledModel.provider}`);
        return false;
      }

      // 初始化OpenAI客户端
      client = new OpenAI({
        apiKey: enabledModel.api_key,
        baseURL: enabledModel.endpoint || providerInfo.endpoint,
        dangerouslyAllowBrowser: true,
      });

      // 如果模型为空，或者模型格式不是新格式 (configId:modelName)
      const isValidModelId = settings.model && settings.model.includes(':');

      if (!isValidModelId && providerInfo.models.length > 0) {
        // 使用新格式的模型ID: "configId:modelName"
        settings.model = `${enabledModel.id}:${providerInfo.models[0]}`;
      }

      return true;
    } catch (error) {
      console.error('加载设置失败:', error);
      toast.error('加载模型设置失败');
      return false;
    }
  } else {
    toast.error('未找到启用的云端模型配置');
    return false;
  }
};

// 类型定义
interface Conversation {
  id: number;
  title: string;
  group: string;
  active: boolean;
}

interface Message {
  type: 'user' | 'assistant';
  content: string;
  typing?: boolean;
}

// 聊天相关功能
const chatMessagesComponent = ref<any>(null)
const conversations = ref<Conversation[]>([
  {
    id: 1,
    title: '新对话',
    group: '今天',
    active: true
  }
])

const messages = ref<Message[]>([])

// 设置面板
const showSettings = ref(false)
const settings = reactive({
  model: '',
  temperature: 0.7,
  maxTokens: 2000,
  contextLength: 10
})

// 从localStorage加载设置时确保类型正确
const loadSettingsFromLocalStorage = () => {
  const savedSettings = localStorage.getItem(SETTINGS_STORAGE_KEY);
  if (savedSettings) {
    try {
      const parsedSettings = JSON.parse(savedSettings);

      // 确保数值类型正确
      if (parsedSettings.temperature) settings.temperature = Number(parsedSettings.temperature);
      if (parsedSettings.maxTokens) settings.maxTokens = Number(parsedSettings.maxTokens);
      if (parsedSettings.contextLength) settings.contextLength = Number(parsedSettings.contextLength);
      if (parsedSettings.model) settings.model = parsedSettings.model;
    } catch (error) {
      console.error('解析保存的设置失败:', error);
    }
  }
}

// 保存设置到localStorage
const saveSettings = () => {
  localStorage.setItem(SETTINGS_STORAGE_KEY, JSON.stringify(settings));
}

// 监听设置变化，自动保存
watch(settings, () => {
  saveSettings();
}, {deep: true});

const toggleSettings = () => {
  showSettings.value = !showSettings.value
}

// 点击空白处关闭设置面板
const handleClickOutside = (event: MouseEvent) => {
  // 获取设置面板和触发按钮的元素
  const settingsPanel = document.querySelector('.settings-panel');
  const settingsButton = document.querySelector('.settings-button');

  // 如果点击的不是设置面板内的元素，也不是触发按钮，则关闭面板
  if (settingsPanel &&
    !settingsPanel.contains(event.target as Node) &&
    settingsButton &&
    !settingsButton.contains(event.target as Node)) {
    showSettings.value = false;
  }
}

// 滚动到底部
const scrollToBottom = async () => {
  // 使用nextTick确保DOM已更新
  await nextTick();
  if (chatMessagesComponent.value?.messagesContainer) {
    chatMessagesComponent.value.messagesContainer.scrollTop =
      chatMessagesComponent.value.messagesContainer.scrollHeight;
  }
};

// 发送消息到OpenAI并获取流式响应
const sendOpenAIRequest = async (userInput: string) => {
  if (!client) {
    toast.error("请先配置云端模型")
    return
  }

  // 检查是否已选择模型
  if (!settings.model) {
    toast.error("请先选择语言模型")
    return
  }

  // 解析模型ID: 格式为 "configId:modelName"
  const [configIdStr, modelName] = settings.model.split(':');

  if (!configIdStr || !modelName) {
    toast.error("模型ID格式错误");
    return;
  }

  // 获取云端模型配置
  const configId = Number(configIdStr);

  try {
    // 获取云端模型配置
    const result = await GetCloudLLMModels(1, 100);
    const modelConfig = result.items.find(item => item.id === configId);

    if (!modelConfig) {
      toast.error("未找到对应的模型配置");
      return;
    }

    // 添加助手消息（带有typing标记）
    messages.value.push({
      type: 'assistant',
      content: '',
      typing: true
    });

    // 确保消息添加后滚动到底部
    await scrollToBottom();

    EventsOn("stream-request-message", async (data) => {
      console.log(data)
      const lastMessage = messages.value[messages.value.length - 1];
      lastMessage.content += data.content;
      // 添加内容后定期滚动到底部
      if (lastMessage.content.length % 10 === 0) {
        await scrollToBottom();
      }
    })

    StreamRequestMessage({
      cloud_llm_id: modelConfig.id,
      conversation_id: 0,
      question: userInput,
      model_name: modelName,
      temperature: Number(settings.temperature),
      max_completion_tokens: settings.maxTokens,
      history_length: Number(settings.contextLength),
    }).then((totalToken) => {console.log("token消耗量为" + totalToken)})

    // 流式输出完成，移除typing标记
    const lastMessage = messages.value[messages.value.length - 1];
    lastMessage.typing = false;

    // 确保结束后再次滚动到底部
    await scrollToBottom();

    // 如果是第一条消息，更新对话标题
    if (messages.value.length === 3 && activeConversation.value.title === '新对话') {
      // 生成标题，通常使用用户消息的前几个字
      const userMessage = userInput.trim();
      const title = userMessage.length > 15
        ? userMessage.substring(0, 15) + '...'
        : userMessage;

      // 更新当前活动对话的标题
      const activeConvIndex = conversations.value.findIndex(c => c.active);
      if (activeConvIndex !== -1) {
        conversations.value[activeConvIndex].title = title;
      }
    }
  } catch (error) {
    toast.error('发送请求失败，请稍后重试');

    // 移除typing消息
    const typingIndex = messages.value.findIndex(msg => msg.typing);
    if (typingIndex !== -1) {
      messages.value.splice(typingIndex, 1);
    }
  }
}

// 发送消息
const sendMessage = async (text: string) => {
  if (!text.trim()) return;

  // 添加用户消息
  messages.value.push({
    type: 'user',
    content: text,
  });

  // 确保消息添加后滚动到底部
  await scrollToBottom();

  // 调用OpenAI API
  await sendOpenAIRequest(text);
};

// 新建对话
const createNewChat = () => {
  // 更新当前所有对话为非激活状态
  conversations.value.forEach(conv => conv.active = false);

  // 创建新对话
  conversations.value.unshift({
    id: Date.now(),
    title: '新对话',
    group: '今天',
    active: true
  });

  // 清空消息
  messages.value = [];
};

// 切换对话
const switchConversation = (conversation: Conversation) => {
  conversations.value.forEach(conv => {
    conv.active = (conv.id === conversation.id);
  });
  messages.value = [];
  // 滚动到底部
  scrollToBottom();
};

// 删除对话
const deleteConversation = (index: number, event: Event) => {
  // 阻止事件冒泡
  event.stopPropagation();

  const wasActive = conversations.value[index].active;
  conversations.value.splice(index, 1);

  // 如果删除的是当前活动对话，则激活第一个对话
  if (wasActive && conversations.value.length > 0) {
    conversations.value[0].active = true;

    // 更新消息
    switchConversation(conversations.value[0]);
  } else if (conversations.value.length === 0) {
    // 如果没有对话了，创建一个新对话
    createNewChat();
  }
};

// 获取当前活动对话
const activeConversation = computed(() => {
  return conversations.value.find(c => c.active) || {title: '新对话'};
});

// 监听消息变化自动滚动 - 使用更可靠的实现
watch(messages, async () => {
  await scrollToBottom();
}, {deep: true});

// 初始化
onMounted(async () => {
  await scrollToBottom();
  await loadOpenAIApiKey();

  // 添加窗口大小变化时的滚动处理
  window.addEventListener('resize', scrollToBottom);
});
</script>

<template>
  <div class="flex flex-col h-full w-full overflow-hidden" @click="showSettings && handleClickOutside($event)">
    <!-- 聊天界面 -->
    <div class="flex flex-1 overflow-hidden">
      <!-- 侧边栏 -->
      <ConversationList
        :conversations="conversations"
        @create-new="createNewChat"
        @switch="switchConversation"
        @delete="deleteConversation"
      />

      <!-- 主内容区 -->
      <main class="flex-1 flex flex-col bg-base-100">
        <!-- 聊天头部 -->
        <ChatHeader
          :title="activeConversation.title"
          @toggle-settings="toggleSettings"
          class="settings-button"
        />

        <!-- 设置面板 -->
        <SettingsPanel
          v-if="showSettings"
          :settings="settings"
          @close="toggleSettings"
          class="settings-panel"
        />

        <!-- 消息区域 -->
        <ChatMessages
          ref="chatMessagesComponent"
          :messages="messages"
        />

        <!-- 输入区域 -->
        <ChatInput @send="sendMessage"/>
      </main>
    </div>
  </div>
</template>

<style scoped>
/* 滚动条样式 */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  @apply bg-transparent;
}

::-webkit-scrollbar-thumb {
  @apply bg-base-content/20 hover:bg-base-content/30;
  border-radius: 10px;
}

@keyframes cursor-blink {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
}
</style>
