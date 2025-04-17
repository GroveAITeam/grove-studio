<script lang="ts" setup>
import { defineProps, ref, watchEffect, nextTick } from 'vue';

interface Message {
  type: 'user' | 'assistant';
  content: string;
  typing?: boolean;
}

const props = defineProps<{
  messages: Message[]
}>();

const messagesContainer = ref<HTMLElement | null>(null);

// 滚动到底部
const scrollToBottom = async () => {
  await nextTick();
  if (messagesContainer.value) {
    const container = messagesContainer.value;
    container.scrollTop = container.scrollHeight;
  }
};

// 监听消息变化自动滚动
watchEffect(() => {
  if (props.messages.length) {
    // 使用两次延时确保在所有内容都渲染完成后滚动
    setTimeout(() => {
      scrollToBottom();
      // 再次滚动以确保捕获所有更新
      setTimeout(scrollToBottom, 100);
    }, 50);
  }
});
</script>

<template>
  <div class="flex-1 overflow-y-auto relative" ref="messagesContainer">
    <div class="px-4 py-4 space-y-6">
      <div v-for="(message, index) in messages" :key="index"
           :class="['flex', message.type === 'user' ? 'justify-end' : '']">
        <div :class="['max-w-[80%] rounded-xl p-3 relative',
                    message.type === 'user' ? 'bg-primary text-primary-content' : 'bg-base-200 text-base-content']">
          <div :class="['prose prose-sm', message.typing ? 'typing' : '']" v-html="message.content"></div>
        </div>
      </div>
    </div>
    <!-- 底部填充，确保最后一条消息不被输入框遮挡 -->
    <div class="h-20"></div>
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

/* 滚动容器样式 */
.overflow-y-auto {
  scrollbar-width: thin;
  scrollbar-color: rgba(0, 0, 0, 0.2) transparent;
}

.overflow-y-auto::-webkit-scrollbar {
  width: 4px;
}

.overflow-y-auto::-webkit-scrollbar-track {
  background: transparent;
}

.overflow-y-auto::-webkit-scrollbar-thumb {
  background-color: rgba(0, 0, 0, 0.2);
  border-radius: 4px;
}
</style>
