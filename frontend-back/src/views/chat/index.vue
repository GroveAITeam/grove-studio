<script lang="ts" setup>
import {computed, nextTick, onMounted, onUnmounted, reactive, ref, watch} from 'vue'
import ConversationList from '../../components/chat/ConversationList.vue';
import ChatHeader from '../../components/chat/ChatHeader.vue';
import ChatMessages from '../../components/chat/ChatMessages.vue';
import ChatInput from '../../components/chat/ChatInput.vue';
import SettingsPanel from '../../components/chat/SettingsPanel.vue';
import OpenAI from "openai";
import {
  DestroyConversation,
  GetCloudLLMModels,
  GetConversationList,
  GetMessageList,
  StreamRequestMessage
} from '../../../wailsjs/go/main/App';
import {useToast} from "../../utils/toast";
import {LLM_PROVIDERS} from '../../constants/LLMProviders';
import {EventsOn} from '../../../wailsjs/runtime';

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
  created_at?: string;
  updated_at?: string;
  active: boolean; // 前端状态，非后端字段
  group: string;   // 前端展示用，非后端字段
}

interface Message {
  id?: number;
  conversation_id?: number;
  type: 'user' | 'assistant';
  content: string;
  role?: string;
  typing?: boolean;
  created_at?: string;
  updated_at?: string;
}

// 聊天相关功能
const chatMessagesComponent = ref<any>(null)
const conversations = ref<Conversation[]>([])
const messages = ref<Message[]>([])
const currentPage = ref(1)
const pageSize = ref(10)
const searchQuery = ref('')
const isLoading = ref(false)

// 加载会话列表
const loadConversations = async () => {
  try {
    isLoading.value = true
    const result = await GetConversationList(currentPage.value, pageSize.value, searchQuery.value)

    // 转换结果，添加前端需要的字段
    conversations.value = result.items.map(conv => ({
      ...conv,
      group: formatDateToGroup(conv.created_at),
      active: false
    }))

    // 如果有会话，默认选中第一个
    if (conversations.value.length > 0) {
      conversations.value[0].active = true
      await loadMessages(conversations.value[0].id)
    }
  } catch (error) {
    toast.error('加载会话列表失败')
    console.error('加载会话列表失败:', error)
  } finally {
    isLoading.value = false
  }
}

// 日期格式化为分组名称
const formatDateToGroup = (dateString: string | null | undefined): string => {
  if (!dateString) return '未分类'

  const date = new Date(dateString)
  const today = new Date()
  const yesterday = new Date(today)
  yesterday.setDate(yesterday.getDate() - 1)

  if (date.toDateString() === today.toDateString()) {
    return '今天'
  } else if (date.toDateString() === yesterday.toDateString()) {
    return '昨天'
  } else {
    // 返回月份和日期，例如：6月10日
    return `${date.getMonth() + 1}月${date.getDate()}日`
  }
}

// 加载消息列表
const loadMessages = async (conversationId: number, minId: number = 0, append: boolean = false) => {
  try {
    const size = 10

    const result = await GetMessageList(conversationId, minId, size)

    // 转换消息格式
    const loadedMessages = result.items.map(msg => ({
      id: msg.id,
      conversation_id: msg.conversation_id,
      type: msg.role as 'user' | 'assistant',
      content: msg.content,
      role: msg.role,
      created_at: msg.created_at,
      updated_at: msg.updated_at
    }))

    // 如果是追加模式
    if (append) {
      // 只有在有新消息时才追加
      if (loadedMessages.length > 0) {
        // 添加新加载的消息到现有列表的前面
        messages.value = [...loadedMessages, ...messages.value]
      } else {
        // 如果没有更多消息，可以设置一个标志表示已到达历史记录顶部
        console.log('没有更多历史消息')
      }
    } else {
      // 初始加载时，替换整个消息列表（只有在有消息时才替换）
      if (loadedMessages.length > 0) {
        messages.value = loadedMessages

        // 确保消息按创建时间正序排列
        messages.value.sort((a, b) => {
          if (!a.created_at || !b.created_at) return 0
          return new Date(a.created_at).getTime() - new Date(b.created_at).getTime()
        })
      }

      // 新加载整个会话后滚动到底部
      await scrollToBottom()
    }
  } catch (error) {
    toast.error('加载消息失败')
    console.error('加载消息失败:', error)
  }
}

// 处理加载更多消息
const handleLoadMore = async (oldestMessageId: number) => {
  // 获取当前活动会话
  const activeConv = activeConversation.value
  if (!activeConv || activeConv.id === 0) return

  // 确保有有效的最早消息ID
  if (!oldestMessageId || oldestMessageId === Infinity) {
    console.log('无法确定最早的消息ID，无法加载更多')
    return
  }

  // 加载更早的消息（附加到现有消息列表前面）
  await loadMessages(activeConv.id, oldestMessageId, true)
}

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

  // 获取当前活动会话ID
  const activeConv = activeConversation.value
  const conversationId = activeConv?.id || 0

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

    // 负责发起流式请求
    await StreamRequestMessage({
      cloud_llm_id: modelConfig.id,
      conversation_id: conversationId,
      question: userInput,
      model_name: modelName,
      temperature: Number(settings.temperature),
      max_completion_tokens: Number(settings.contextLength) <= 3500 ? 0 : Number(settings.contextLength),
      history_length: Number(settings.contextLength),
    }).then((totalToken) => {console.log("token消耗量为" + totalToken)})

    // 流式输出完成，移除typing标记
    const lastMessage = messages.value[messages.value.length - 1];
    lastMessage.typing = false;

    // 确保结束后再次滚动到底部
    await scrollToBottom();

    // 如果是新会话（ID为0），请求完成后需要刷新会话列表获取新创建的会话
    if (conversationId === 0) {
      await loadConversations()
    }
  } catch (error) {
    toast.error('发送请求失败，请稍后重试');
    console.error('发送请求失败:', error);

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

  // 清空消息
  messages.value = [];
};

// 切换对话
const switchConversation = async (conversation: Conversation) => {
  conversations.value.forEach(conv => {
    conv.active = (conv.id === conversation.id);
  });

  // 清空当前消息
  messages.value = [];

  // 加载所选会话的消息
  await loadMessages(conversation.id);

  // 滚动到底部
  await scrollToBottom();
};

// 删除对话
const deleteConversation = async (index: number, event: Event) => {
  // 阻止事件冒泡
  event.stopPropagation();

  const conversation = conversations.value[index];
  const wasActive = conversation.active;

  try {
    // 调用后端删除会话
    await DestroyConversation(conversation.id);

    // 从前端列表中移除
    conversations.value.splice(index, 1);

    // 如果删除的是当前活动对话，则激活第一个对话
    if (wasActive && conversations.value.length > 0) {
      conversations.value[0].active = true;

      // 加载新激活的会话消息
      await loadMessages(conversations.value[0].id);
    } else if (conversations.value.length === 0) {
      // 如果没有对话了，清空消息
      messages.value = [];
    }

    toast.success('删除会话成功');
  } catch (error) {
    toast.error('删除会话失败');
    console.error('删除会话失败:', error);
  }
};

// 获取当前活动对话
const activeConversation = computed(() => {
  const active = conversations.value.find(c => c.active);
  return active || { id: 0, title: '新对话' };
});

// 监听消息变化自动滚动 - 使用更可靠的实现
watch(messages, async () => {
  await scrollToBottom();
}, {deep: true});

// 用于缓存分片内容
const streamChunks = ref<Map<number, string>>(new Map());
const pendingDone = ref<boolean>(false);
const doneTimeout = ref<number | null>(null);

EventsOn("stream-request-message", (data) => {
  // data.index: 当前分片序号
  // data.content: 内容
  // data.done: 是否结束
  if (typeof data.index !== 'number') return;

  // 查找最后一个typing的assistant消息
  const lastTypingIndex = messages.value.slice().reverse().findIndex(msg => msg.type === 'assistant' && msg.typing);
  let lastMessage;
  if (lastTypingIndex !== -1) {
    lastMessage = messages.value[messages.value.length - 1 - lastTypingIndex];
  } else {
    // 兜底：找最后一个assistant消息
    lastMessage = messages.value.slice().reverse().find(msg => msg.type === 'assistant');
  }

  if (lastMessage) {
    try {
      // 如果收到done信号，设置延迟清理标记
      if (data.done) {
        pendingDone.value = true;
        // 清除之前的定时器
        if (doneTimeout.value) {
          clearTimeout(doneTimeout.value);
        }
        // 设置新的定时器，延迟200ms清理
        doneTimeout.value = setTimeout(() => {
          streamChunks.value.clear();
          pendingDone.value = false;
          doneTimeout.value = null;
        }, 200) as unknown as number;
        return;
      }

      // 处理内容
      if (data.content) {
        streamChunks.value.set(data.index, data.content);

        // 重新构建内容
        lastMessage.content = '';
        const sortedIndices = Array.from(streamChunks.value.keys()).sort((a, b) => a - b);
        for (const index of sortedIndices) {
          const chunk = streamChunks.value.get(index);
          if (chunk) {
            lastMessage.content += chunk;
          }
        }

        // 添加内容后定期滚动到底部
        if (lastMessage.content.length % 10 === 0) {
          scrollToBottom();
        }
      }
    } catch (error) {
      console.error('处理流式消息出错:', error);
    }
  }
});

// 在组件卸载时清理定时器
onUnmounted(() => {
  if (doneTimeout.value) {
    clearTimeout(doneTimeout.value);
  }
});

// 初始化
onMounted(async () => {
  await loadOpenAIApiKey();
  await loadConversations();

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
          @load-more="handleLoadMore"
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
