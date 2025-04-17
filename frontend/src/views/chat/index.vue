<script lang="ts" setup>
import {ref, reactive, onMounted, computed, watch} from 'vue'
import ConversationList from '../../components/chat/ConversationList.vue';
import ChatHeader from '../../components/chat/ChatHeader.vue';
import ChatMessages from '../../components/chat/ChatMessages.vue';
import ChatInput from '../../components/chat/ChatInput.vue';
import SettingsPanel from '../../components/chat/SettingsPanel.vue';
import OpenAI from "openai";
import { GetCloudLLMModels } from '../../../wailsjs/go/main/App';

// OpenAI 客户端
let client: OpenAI | null = null;

// 加载可用的OpenAI API Key
const loadOpenAIApiKey = async () => {
  try {
    // 从Wails后端获取云端模型信息
    const result = await GetCloudLLMModels(1, 10);

    // 查找启用的OpenAI模型
    const openaiModel = result.items.find(model =>
      model.provider === 'openai' && model.enabled
    );

    if (openaiModel) {
      // 初始化OpenAI客户端
      client = new OpenAI({
        apiKey: openaiModel.api_key,
        dangerouslyAllowBrowser: true,
      });

      // 更新设置中的模型
      if (openaiModel.name) {
        settings.model = openaiModel.name;
      }

      console.log('已成功加载OpenAI API Key');
      return true;
    } else {
      console.error('找不到启用的OpenAI模型');
      return false;
    }
  } catch (error) {
    console.error('加载OpenAI API Key失败:', error);
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
  time: string;
  typing?: boolean;
}

// 聊天相关功能
const messageInput = ref('')
const messagesContainer = ref<HTMLElement | null>(null)
const conversations = ref<Conversation[]>([
  {
    id: 1,
    title: '新对话',
    group: '今天',
    active: true
  }
])

const messages = ref<Message[]>([])

// 获取当前时间格式化字符串
function getCurrentTime(): string {
  const now = new Date();
  return `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`;
}

// 设置面板
const showSettings = ref(false)
const settings = reactive({
  model: 'gpt-3.5-turbo',
  temperature: 0.7,
  maxTokens: 2000,
  contextLength: 10
})

const toggleSettings = () => {
  showSettings.value = !showSettings.value
}

// 监听消息变化自动滚动
watch(messages, () => {
  setTimeout(scrollToBottom, 50)
}, { deep: true })

// 发送消息到OpenAI并获取流式响应
const sendOpenAIRequest = async (userInput: string) => {
  // 确保客户端已初始化
  if (!client) {
    const success = await loadOpenAIApiKey();
    if (!success) {
      // 添加错误消息
      messages.value.push({
        type: 'assistant',
        content: '无法连接到OpenAI，请确保已在云端模型设置中配置有效的API密钥。',
        time: getCurrentTime()
      });
      return;
    }
  }

  try {
    // 收集历史消息作为上下文
    const history = messages.value.slice(-settings.contextLength).map(msg => ({
      role: msg.type,
      content: msg.content
    }));

    // 添加当前用户消息
    history.push({
      role: 'user',
      content: userInput
    });

    // 确保client不为null
    if (!client) {
      throw new Error('OpenAI客户端未初始化');
    }

    // 创建流式请求
    const response = await client.chat.completions.create({
      model: settings.model,
      messages: history,
      temperature: settings.temperature,
      max_tokens: settings.maxTokens,
      stream: true,
    });

    // 添加助手消息（带有typing标记）
    const timeString = getCurrentTime();
    messages.value.push({
      type: 'assistant',
      content: '',
      time: timeString,
      typing: true
    });

    // 滚动到底部
    setTimeout(scrollToBottom, 50);

    // 处理流式响应
    for await (const chunk of response) {
      const lastMessage = messages.value[messages.value.length - 1];
      const content = chunk.choices[0]?.delta?.content || '';

      if (content) {
        lastMessage.content += content;
      }
    }

    // 流式输出完成，移除typing标记
    const lastMessage = messages.value[messages.value.length - 1];
    lastMessage.typing = false;

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
    console.error('OpenAI API 调用失败:', error);

    // 添加错误消息
    const lastMessage = messages.value[messages.value.length - 1];
    if (lastMessage.type === 'assistant' && lastMessage.typing) {
      lastMessage.content += '\n\n[调用API时发生错误，请检查网络连接或API设置]';
      lastMessage.typing = false;
    } else {
      messages.value.push({
        type: 'assistant',
        content: '调用API时发生错误，请检查网络连接或API设置。',
        time: getCurrentTime(),
      });
    }
  }
};

// 发送消息
const sendMessage = async (text: string) => {
  if (!text.trim()) return;

  // 获取当前时间
  const timeString = getCurrentTime();

  // 添加用户消息
  messages.value.push({
    type: 'user',
    content: text,
    time: timeString
  });

  // 滚动到底部
  setTimeout(scrollToBottom, 50);

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

  // 这里只是切换前端对话，没有保存到数据库，所以每次切换回来都是初始状态
  // 未来可以实现对话历史的保存和加载
  const timeString = getCurrentTime();

  messages.value = [];

  // 滚动到底部
  setTimeout(scrollToBottom, 50);
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
  return conversations.value.find(c => c.active) || { title: '新对话' };
});

// 滚动到底部
const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
  }
};

// 监视消息变化，自动滚动到底部
onMounted(async () => {
  scrollToBottom();
  await loadOpenAIApiKey();
});
</script>

<template>
  <div class="flex flex-col h-full w-full overflow-hidden">
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
        />

        <!-- 设置面板 -->
        <SettingsPanel
          v-if="showSettings"
          :settings="settings"
          @close="toggleSettings"
        />

        <!-- 消息区域 -->
        <ChatMessages
          ref="messagesContainer"
          :messages="messages"
        />

        <!-- 输入区域 -->
        <ChatInput @send="sendMessage" />
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
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}
</style>
