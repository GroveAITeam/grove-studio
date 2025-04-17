<script lang="ts" setup>
import {ref, reactive, onMounted, computed, watch} from 'vue'
import ConversationList from '../../components/chat/ConversationList.vue';
import ChatHeader from '../../components/chat/ChatHeader.vue';
import ChatMessages from '../../components/chat/ChatMessages.vue';
import ChatInput from '../../components/chat/ChatInput.vue';
import SettingsPanel from '../../components/chat/SettingsPanel.vue';

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
    title: '关于人工智能的基础知识',
    group: '今天',
    active: true
  },
  {
    id: 2,
    title: '深度学习模型的应用场景',
    group: '今天',
    active: false
  },
  {
    id: 3,
    title: 'Python编程入门指南',
    group: '昨天',
    active: false
  },
])

const messages = ref<Message[]>([
  {
    type: 'assistant',
    content: '你好！我是Grove AI助手，很高兴为你提供关于人工智能的基础知识。人工智能(AI)是计算机科学的一个分支，致力于开发能够执行通常需要人类智能的任务的系统。\n\n人工智能的主要类型包括：\n- **弱人工智能**：专注于执行特定任务的系统，如语音识别或自动驾驶。\n- **强人工智能**：理论上可以执行任何人类可以做的智力任务的系统。\n- **超级人工智能**：假设的系统，其能力远超人类智能。\n\n你想了解人工智能的哪个方面？',
    time: '10:25'
  },
  {
    type: 'user',
    content: '请详细介绍一下机器学习是如何工作的？特别是深度学习模型。',
    time: '10:27'
  },
  {
    type: 'assistant',
    content: '机器学习是人工智能的一个子领域，它专注于开发能够从数据中学习并做出预测或决策的算法，而无需被明确编程来执行特定任务。\n\n机器学习的基本工作原理是：\n\n1. **数据收集**：收集相关数据作为学习的基础。\n2. **数据预处理**：清洗、规范化和转换数据。\n3. **特征提取**：从原始数据中提取有用的特征。\n4. **模型训练**：利用算法在训练数据上学习模式。\n5. **模型评估**：使用测试数据评估模型性能。\n6. **模型部署**：将模型应用于实际问题。',
    time: '10:28'
  }
])

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

// 发送消息
const sendMessage = (text: string) => {
  if (!text.trim()) return

  // 获取当前时间
  const now = new Date()
  const timeString = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`

  // 添加用户消息
  messages.value.push({
    type: 'user',
    content: text,
    time: timeString
  })

  // 滚动到底部
  setTimeout(scrollToBottom, 50)

  // 模拟助手响应
  setTimeout(() => {
    const exampleResponses = [
      "我理解你的问题了。",
      "根据我的分析，你提到的内容涉及到几个关键点。",
      "首先，让我们考虑一下这个问题的核心。",
      "从技术角度来看，这是一个很好的问题。",
      "我可以提供一些有用的信息和建议。"
    ]

    // 随机选择一个响应
    const responseText = exampleResponses[Math.floor(Math.random() * exampleResponses.length)]

    // 添加助手消息（带有typing标记）
    messages.value.push({
      type: 'assistant',
      content: '',
      time: timeString,
      typing: true
    })

    // 滚动到底部
    setTimeout(scrollToBottom, 50)

    // 模拟流式输出
    let index = 0
    const interval = setInterval(() => {
      const lastMessage = messages.value[messages.value.length - 1]

      if (index < responseText.length) {
        lastMessage.content += responseText[index]
        index++
      } else {
        clearInterval(interval)
        lastMessage.typing = false
      }
    }, 50)
  }, 500)
}

// 新建对话
const createNewChat = () => {
  // 更新当前所有对话为非激活状态
  conversations.value.forEach(conv => conv.active = false)

  // 获取当前时间
  const now = new Date()
  const timeString = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`

  // 创建新对话
  conversations.value.unshift({
    id: Date.now(),
    title: '新对话',
    group: '今天',
    active: true
  })

  // 清空消息
  messages.value = [
    {
      type: 'assistant',
      content: '你好！我是Grove AI助手。有什么我可以帮助你的吗？',
      time: timeString
    }
  ]
}

// 切换对话
const switchConversation = (conversation: Conversation) => {
  conversations.value.forEach(conv => {
    conv.active = (conv.id === conversation.id)
  })

  // 这里可以加载对应的聊天历史
  // 示例中仅展示默认消息
  if (conversation.title === '关于人工智能的基础知识') {
    messages.value = [
      {
        type: 'assistant',
        content: '你好！我是Grove AI助手，很高兴为你提供关于人工智能的基础知识。人工智能(AI)是计算机科学的一个分支，致力于开发能够执行通常需要人类智能的任务的系统。\n\n人工智能的主要类型包括：\n- **弱人工智能**：专注于执行特定任务的系统，如语音识别或自动驾驶。\n- **强人工智能**：理论上可以执行任何人类可以做的智力任务的系统。\n- **超级人工智能**：假设的系统，其能力远超人类智能。\n\n你想了解人工智能的哪个方面？',
        time: '10:25'
      },
      {
        type: 'user',
        content: '请详细介绍一下机器学习是如何工作的？特别是深度学习模型。',
        time: '10:27'
      },
      {
        type: 'assistant',
        content: '机器学习是人工智能的一个子领域，它专注于开发能够从数据中学习并做出预测或决策的算法，而无需被明确编程来执行特定任务。\n\n机器学习的基本工作原理是：\n\n1. **数据收集**：收集相关数据作为学习的基础。\n2. **数据预处理**：清洗、规范化和转换数据。\n3. **特征提取**：从原始数据中提取有用的特征。\n4. **模型训练**：利用算法在训练数据上学习模式。\n5. **模型评估**：使用测试数据评估模型性能。\n6. **模型部署**：将模型应用于实际问题。',
        time: '10:28'
      }
    ]
  } else {
    // 其他对话只显示欢迎消息
    const now = new Date()
    const timeString = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`

    messages.value = [
      {
        type: 'assistant',
        content: `你好！我是Grove AI助手。这是关于"${conversation.title}"的对话。有什么我可以帮助你的吗？`,
        time: timeString
      }
    ]
  }

  // 滚动到底部
  setTimeout(scrollToBottom, 50)
}

// 删除对话
const deleteConversation = (index: number, event: Event) => {
  // 阻止事件冒泡
  event.stopPropagation()

  const wasActive = conversations.value[index].active
  conversations.value.splice(index, 1)

  // 如果删除的是当前活动对话，则激活第一个对话
  if (wasActive && conversations.value.length > 0) {
    conversations.value[0].active = true

    // 更新消息
    switchConversation(conversations.value[0])
  }
}

// 获取当前活动对话
const activeConversation = computed(() => {
  return conversations.value.find(c => c.active) || { title: '新对话' }
})

// 滚动到底部
const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

// 监视消息变化，自动滚动到底部
onMounted(() => {
  scrollToBottom()
})
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
