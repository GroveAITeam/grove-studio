<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

const props = defineProps<{
  show: boolean;
  title?: string;
  message: string;
  confirmText?: string;
  cancelText?: string;
}>();

const emit = defineEmits<{
  confirm: [];
  cancel: [];
}>();

const handleConfirm = () => {
  emit('confirm');
};

const handleCancel = () => {
  emit('cancel');
};
</script>

<template>
  <Teleport to="body">
    <!-- 遮罩层 -->
    <div v-if="show" class="fixed inset-0 bg-black/50 z-50" @click="handleCancel"></div>

    <!-- 对话框 -->
    <div v-if="show"
         class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 z-50 w-[90%] max-w-md bg-base-100 rounded-lg shadow-xl">
      <div class="p-6">
        <!-- 标题 -->
        <h3 v-if="title" class="text-lg font-bold mb-4">{{ title }}</h3>

        <!-- 消息内容 -->
        <p class="text-base-content/80">{{ message }}</p>

        <!-- 按钮组 -->
        <div class="flex justify-end gap-3 mt-6">
          <button class="btn btn-ghost" @click="handleCancel">
            {{ cancelText || '取消' }}
          </button>
          <button class="btn btn-primary" @click="handleConfirm">
            {{ confirmText || '确认' }}
          </button>
        </div>
      </div>
    </div>
  </Teleport>
</template>
