<script lang="ts" setup>
import { ref } from 'vue';
import { Icon } from '@iconify/vue';

const emit = defineEmits<{
  send: [text: string]
}>();

const messageInput = ref('');

// 按Enter发送消息
const handleKeyDown = (event: KeyboardEvent) => {
  if (event.key === 'Enter' && !event.shiftKey) {
    event.preventDefault();
    sendMessage();
  }
};

// 自动调整输入框高度
const adjustTextareaHeight = (event: Event) => {
  const textarea = event.target as HTMLTextAreaElement;
  textarea.style.height = 'auto';
  textarea.style.height = `${Math.min(textarea.scrollHeight, 120)}px`;
};

// 发送消息
const sendMessage = () => {
  if (!messageInput.value.trim()) return;

  emit('send', messageInput.value);

  // 清空输入框
  messageInput.value = '';
};
</script>

<template>
  <div class="border-t border-base-200 p-2 bg-base-100">
    <div class="flex flex-col gap-2">
      <div class="relative">
        <textarea
          v-model="messageInput"
          @keydown="handleKeyDown"
          @input="adjustTextareaHeight"
          class="w-full bg-base-200 border border-base-300 rounded-lg pl-3 pr-12 py-2 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary text-base-content placeholder:text-base-content/50 resize-none text-sm min-h-[40px]"
          placeholder="给 Grove AI 发送消息..."
          rows="2"
        ></textarea>
        <button
          @click="sendMessage"
          class="absolute right-3 top-1/2 transform -translate-y-1/2 p-1.5 text-base-content/70 hover:text-primary hover:bg-base-300/50 rounded-full transition-colors"
          :disabled="!messageInput.trim()"
        >
          <Icon icon="material-symbols:send" class="text-lg" />
        </button>
      </div>
      <div class="text-xs text-base-content/60 flex items-center gap-1 justify-center">
        <Icon icon="material-symbols:info" class="text-sm" />
        <span>Grove AI 可能会产生不准确的信息，请自行判断内容的准确性。</span>
      </div>
    </div>
  </div>
</template>
