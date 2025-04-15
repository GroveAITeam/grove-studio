<script lang="ts" setup>
import { ref, reactive, onMounted, computed, watch } from 'vue'
import { Icon } from '@iconify/vue';

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
const sendMessage = () => {
  if (!messageInput.value.trim()) return

  // 获取当前时间
  const now = new Date()
  const timeString = `${now.getHours().toString().padStart(2, '0')}:${now.getMinutes().toString().padStart(2, '0')}`

  // 添加用户消息
  messages.value.push({
    type: 'user',
    content: messageInput.value,
    time: timeString
  })

  // 清空输入框
  messageInput.value = ''

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

// 按Enter发送消息
const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault()
    sendMessage()
  }
}

// 自动调整输入框高度
const adjustTextareaHeight = (event: Event) => {
  const textarea = event.target as HTMLTextAreaElement
  textarea.style.height = 'auto'
  textarea.style.height = `${Math.min(textarea.scrollHeight, 120)}px`
}

// 格式化消息内容（支持Markdown语法）
const formatMessage = (content: string) => {
  if (!content) return ''

  // 转换粗体
  content = content.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')

  // 转换列表（简单实现）
  content = content.replace(/^- (.*?)$/gm, '<li>$1</li>')
  content = content.replace(/(<li>.*?<\/li>)/gs, '<ul>$1</ul>')

  // 转换有序列表
  content = content.replace(/^\d+\. (.*?)$/gm, '<li>$1</li>')
  content = content.replace(/(<li>.*?<\/li>)(?!\s*<\/ul>)/gs, '<ol>$1</ol>')

  // 转换段落
  content = content.split('\n\n').map((p: string) => {
    if (!p.startsWith('<ul>') && !p.startsWith('<ol>')) {
      return `<p>${p}</p>`
    }
    return p
  }).join('')

  return content
}

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
      <!-- 侧边栏 - 白色背景 -->
      <aside class="w-56 border-r border-gray-200 flex flex-col overflow-hidden">
        <div class="p-2 border-gray-200">
          <!-- 新对话按钮 -->
          <button @click="createNewChat" class="w-full flex items-center justify-center gap-2 bg-blue-500 hover:bg-blue-600 text-white font-medium rounded-md py-2 text-sm transition-colors">
            <Icon icon="material-symbols:add" />
            <span>新对话</span>
          </button>
        </div>

        <!-- 对话列表 -->
        <div class="overflow-y-auto flex-1 p-2">
          <template v-for="(group, groupName) in conversations.reduce((groups: Record<string, Conversation[]>, conv) => {
            (groups[conv.group] = groups[conv.group] || []).push(conv);
            return groups;
          }, {})">
            <div class="mb-2">
              <div class="text-xs text-gray-500 uppercase px-2 mb-1 font-medium">{{ groupName }}</div>
              <div v-for="(conversation, index) in group" :key="conversation.id"
                   :class="['flex items-center gap-1 px-2 py-1.5 rounded-md cursor-pointer transition-colors',
                   conversation.active ? 'bg-gray-100' : 'hover:bg-gray-50']"
                   @click="switchConversation(conversation)">
                <Icon icon="material-symbols:chat" class="text-base text-gray-600" />
                <span class="flex-1 truncate text-sm">{{ conversation.title }}</span>
                <button @click="(event) => deleteConversation(conversations.indexOf(conversation), event)"
                        class="opacity-0 hover:opacity-100 hover:bg-gray-200 p-0.5 rounded-md group-hover:opacity-100 transition-opacity">
                  <Icon icon="material-symbols:close" class="text-xs text-gray-500" />
                </button>
              </div>
            </div>
          </template>
        </div>
      </aside>

      <!-- 主内容区 -->
      <main class="flex-1 flex flex-col bg-white overflow-hidden">
        <!-- 聊天头部 -->
        <div class="border-b border-gray-200 p-2 flex justify-between items-center">
          <h2 class="text-base font-medium text-gray-800">{{ activeConversation.title }}</h2>
          <div class="flex gap-2">
            <button @click="toggleSettings" class="p-1.5 text-gray-500 hover:bg-gray-100 rounded-md transition-colors">
              <Icon icon="material-symbols:settings" class="text-lg" />
            </button>
          </div>
        </div>

        <!-- 设置面板 -->
        <div v-if="showSettings" class="absolute right-6 mt-10 w-72 bg-white border border-gray-200 shadow-lg rounded-md z-10 p-3">
          <div class="flex justify-between items-center mb-3">
            <h3 class="font-medium text-sm">对话设置</h3>
            <button @click="toggleSettings" class="text-gray-500 hover:bg-gray-100 p-1 rounded-md transition-colors">
              <Icon icon="material-symbols:close" class="text-base" />
            </button>
          </div>

          <div class="space-y-3">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">语言模型</label>
              <select v-model="settings.model" class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm">
                <option value="gpt-4">GPT-4（最强大）</option>
                <option value="gpt-3.5-turbo">GPT-3.5（均衡）</option>
                <option value="llama3">Llama 3（开源）</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">
                回答灵活度
                <span class="text-xs text-gray-500">（{{Math.round(settings.temperature * 10) / 10}}）</span>
              </label>
              <input type="range" v-model="settings.temperature" min="0" max="1" step="0.1"
                     class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer" />
              <div class="flex justify-between text-xs text-gray-500 mt-1">
                <span>精确</span>
                <span>创造性</span>
              </div>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1">上下文长度</label>
              <select v-model="settings.contextLength" class="w-full border border-gray-300 rounded-md px-2 py-1.5 text-sm">
                <option :value="5">短（5条消息）</option>
                <option :value="10">中（10条消息）</option>
                <option :value="20">长（20条消息）</option>
              </select>
            </div>
          </div>

          <div class="mt-3 pt-2 border-t border-gray-200">
            <button class="w-full bg-blue-500 hover:bg-blue-600 text-white rounded-md py-1.5 text-sm font-medium transition-colors">
              应用设置
            </button>
          </div>
        </div>

        <!-- 消息区域 -->
        <div ref="messagesContainer" class="flex-1 overflow-y-auto py-3 px-4 space-y-4">
          <div v-for="(message, index) in messages" :key="index"
               :class="['flex', message.type === 'user' ? 'justify-end' : '']">
            <div :class="['max-w-[80%] rounded-xl p-3',
                         message.type === 'user' ? 'bg-blue-500 text-white' : 'bg-gray-100 text-gray-800']">
              <div :class="['prose prose-sm', message.typing ? 'typing' : '']">
                <div v-if="message.typing && message.content === ''" class="flex items-center gap-1">
                  <span class="loading loading-dots loading-xs"></span>
                </div>
                <div v-else v-html="formatMessage(message.content)"></div>
              </div>
            </div>
          </div>
        </div>

        <!-- 输入区域 -->
        <div class="border-t border-gray-200 p-2">
          <div class="flex flex-col gap-2">
            <div class="relative">
              <textarea
                v-model="messageInput"
                @keydown="handleKeyDown"
                @input="adjustTextareaHeight"
                class="w-full border border-gray-300 rounded-lg pl-3 pr-12 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none text-sm min-h-[40px]"
                placeholder="给 Grove AI 发送消息..."
                rows="2"
              ></textarea>
              <!-- 发送按钮 -->
              <button
                @click="sendMessage"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 p-1.5 text-gray-500 hover:text-blue-500 hover:bg-gray-100 rounded-full transition-colors"
                :disabled="!messageInput.trim()"
              >
                <Icon icon="material-symbols:send" class="text-lg" />
              </button>
            </div>
            <div class="text-xs text-gray-500 flex items-center gap-1 justify-center">
              <Icon icon="material-symbols:info" class="text-sm" />
              <span>Grove AI 可能会产生不准确的信息，请自行判断内容的准确性。</span>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&display=swap');

/* 设置聊天内容样式 */
.prose {
  max-width: none;
  font-size: 0.875rem;
}

.prose ul {
  list-style-type: disc;
  padding-left: 1.5em;
  margin: 0.5em 0;
}

.prose ol {
  list-style-type: decimal;
  padding-left: 1.5em;
  margin: 0.5em 0;
}

.prose p {
  margin-bottom: 0.5em;
}

.prose p:last-child {
  margin-bottom: 0;
}

.typing:after {
  content: '|';
  animation: cursor-blink 1s infinite;
}

@keyframes cursor-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}

/* 自定义滚动条样式 */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, 0.1);
  border-radius: 10px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: rgba(0, 0, 0, 0.2);
}

/* 自定义range滑块样式 */
input[type=range]::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: #3b82f6;
  cursor: pointer;
}

/* 禁用按钮样式 */
button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>
