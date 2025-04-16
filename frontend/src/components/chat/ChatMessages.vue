<script lang="ts" setup>
import { defineProps } from 'vue';

interface Message {
  type: 'user' | 'assistant';
  content: string;
  time: string;
  typing?: boolean;
}

const props = defineProps<{
  messages: Message[]
}>();

// 格式化消息内容（支持Markdown语法）
const formatMessage = (content: string) => {
  if (!content) return '';

  // 转换粗体
  content = content.replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>');

  // 转换列表（简单实现）
  content = content.replace(/^- (.*?)$/gm, '<li>$1</li>');
  content = content.replace(/(<li>.*?<\/li>)/gs, '<ul>$1</ul>');

  // 转换有序列表
  content = content.replace(/^\d+\. (.*?)$/gm, '<li>$1</li>');
  content = content.replace(/(<li>.*?<\/li>)(?!\s*<\/ul>)/gs, '<ol>$1</ol>');

  // 转换段落
  content = content.split('\n\n').map((p: string) => {
    if (!p.startsWith('<ul>') && !p.startsWith('<ol>')) {
      return `<p>${p}</p>`;
    }
    return p;
  }).join('');

  return content;
};
</script>

<template>
  <div class="flex-1 overflow-y-auto py-3 px-4 space-y-4">
    <div v-for="(message, index) in messages" :key="index"
         :class="['flex', message.type === 'user' ? 'justify-end' : '']">
      <div :class="['max-w-[80%] rounded-xl p-3',
                  message.type === 'user' ? 'bg-primary text-primary-content' : 'bg-base-200 text-base-content']">
        <div :class="['prose prose-sm', message.typing ? 'typing' : '']" v-html="formatMessage(message.content)"></div>
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
