<script lang="ts" setup>
import { defineProps, defineExpose, ref, onMounted, onUnmounted, watch } from 'vue';

interface Message {
  type: 'user' | 'assistant';
  content: string;
  typing?: boolean;
  id?: number;
}

const props = defineProps<{
  messages: Message[]
}>();

// 新增加载更多事件的emit
const emit = defineEmits(['load-more']);

const messagesContainer = ref<HTMLElement | null>(null);
const isLoadingMore = ref(false);
const noMoreMessages = ref(false);

// 处理滚动事件，当滚动到顶部时加载更多消息
const handleScroll = () => {
  if (!messagesContainer.value || isLoadingMore.value || noMoreMessages.value) return;

  // 当滚动到顶部20px范围内时触发加载更多
  if (messagesContainer.value.scrollTop <= 20) {
    loadMoreMessages();
  }
};

// 加载更多消息
const loadMoreMessages = async () => {
  if (isLoadingMore.value || props.messages.length === 0) return;

  // 获取所有带ID的消息
  const messagesWithId = props.messages.filter(m => m.id);
  if (messagesWithId.length === 0) {
    console.log('没有消息带有ID，无法加载更多');
    return;
  }

  isLoadingMore.value = true;

  try {
    // 获取当前消息列表中最早的消息ID作为minId参数
    const oldestMessageId = Math.min(...messagesWithId.map(m => m.id || Infinity));

    // 保存当前的消息数量
    const currentMessageCount = props.messages.length;

    // 触发加载更多事件，传递最早消息ID
    emit('load-more', oldestMessageId);

    // 保存当前滚动位置和高度
    const prevScrollHeight = messagesContainer.value?.scrollHeight || 0;

    // 等待消息加载完成（父组件会更新消息列表）
    await new Promise(resolve => setTimeout(resolve, 300));

    // 如果消息数量没有变化，意味着没有加载到新消息
    if (props.messages.length === currentMessageCount) {
      noMoreMessages.value = true;
      console.log('已加载全部历史消息');
    } else {
      // 调整滚动位置，保持原有内容的相对位置不变
      if (messagesContainer.value) {
        const newScrollHeight = messagesContainer.value.scrollHeight;
        messagesContainer.value.scrollTop = newScrollHeight - prevScrollHeight + 20;
      }
    }
  } finally {
    isLoadingMore.value = false;
  }
};

// 当消息列表改变时，重置noMoreMessages状态
// 这样在切换会话时可以重新加载历史消息
let prevMessagesLength = props.messages.length;
const checkMessagesChange = () => {
  // 如果消息总数减少了，可能是切换到了新的会话
  if (props.messages.length < prevMessagesLength) {
    noMoreMessages.value = false;
  }
  prevMessagesLength = props.messages.length;
};

// 监听props的消息变化
watch(() => props.messages, checkMessagesChange, { deep: true });

onMounted(() => {
  if (messagesContainer.value) {
    messagesContainer.value.addEventListener('scroll', handleScroll);
  }
});

onUnmounted(() => {
  if (messagesContainer.value) {
    messagesContainer.value.removeEventListener('scroll', handleScroll);
  }
});

// 暴露容器引用，使父组件可以操作
defineExpose({
  messagesContainer
});
</script>

<template>
  <div class="flex-1 overflow-y-auto py-3 px-4 space-y-4" ref="messagesContainer">
    <!-- 加载更多指示器 -->
    <div v-if="isLoadingMore" class="flex justify-center py-2">
      <div class="loading loading-spinner loading-sm"></div>
    </div>

    <!-- 所有历史消息已加载提示 -->
    <div v-if="noMoreMessages && !isLoadingMore && messages.length > 0" class="text-center text-xs text-base-content/50 py-1">
      已加载全部历史消息
    </div>

    <div v-for="(message, index) in messages" :key="index"
         :class="['flex', message.type === 'user' ? 'justify-end' : '']">
      <div :class="['max-w-[80%] rounded-xl p-3 relative',
                  message.type === 'user' ? 'bg-primary text-primary-content' : 'bg-base-200 text-base-content']">
        <div :class="['prose prose-sm', message.typing ? 'typing' : '']" v-html="message.content"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Markdown 样式覆盖 */
.prose {
  @apply max-w-none text-base-content;
}

.prose strong {
  @apply text-base-content font-semibold;
}

.prose ul {
  @apply text-base-content;
}

.prose ol {
  @apply text-base-content;
}

/* 打字动画 */
.typing:after {
  content: '|';
  animation: cursor-blink 1s infinite;
}

@keyframes cursor-blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}
</style>
