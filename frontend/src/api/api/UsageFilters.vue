<script setup lang="ts">
import { defineProps } from 'vue';

interface Props {
  selectedTimeRange: string;
  selectedModel: string;
  selectedStatus: string;
}

const props = defineProps<Props>();

defineEmits<{
  (e: 'update:selectedTimeRange', value: string): void;
  (e: 'update:selectedModel', value: string): void;
  (e: 'update:selectedStatus', value: string): void;
  (e: 'filter'): void;
}>();
</script>

<template>
  <div class="usage-filters">
    <div class="filter-group">
      <span class="filter-label text-base-content/70">时间范围：</span>
      <select
        class="filter-select bg-base-200/80 text-base-content"
        :value="selectedTimeRange"
        @change="$emit('update:selectedTimeRange', ($event.target as HTMLSelectElement).value); $emit('filter')"
      >
        <option value="today">今天</option>
        <option value="yesterday">昨天</option>
        <option value="7days">近7天</option>
        <option value="30days">近30天</option>
        <option value="custom">自定义</option>
      </select>
    </div>
    <div class="filter-group">
      <span class="filter-label text-base-content/70">模型：</span>
      <select
        class="filter-select bg-base-200/80 text-base-content"
        :value="selectedModel"
        @change="$emit('update:selectedModel', ($event.target as HTMLSelectElement).value); $emit('filter')"
      >
        <option value="all">全部模型</option>
        <option value="gpt-4">GPT-4</option>
        <option value="gpt-3.5">GPT-3.5</option>
        <option value="claude-3">Claude 3</option>
        <option value="gemini-pro">Gemini Pro</option>
      </select>
    </div>
    <div class="filter-group">
      <span class="filter-label text-base-content/70">状态：</span>
      <select
        class="filter-select bg-base-200/80 text-base-content"
        :value="selectedStatus"
        @change="$emit('update:selectedStatus', ($event.target as HTMLSelectElement).value); $emit('filter')"
      >
        <option value="all">全部状态</option>
        <option value="success">成功</option>
        <option value="error">失败</option>
      </select>
    </div>
  </div>
</template>

<style scoped>
.usage-filters {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  flex-wrap: wrap;
  align-items: center;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 200px; /* 设置最小宽度确保对齐 */
}

.filter-label {
  font-size: 14px;
  min-width: 70px; /* 设置标签最小宽度确保对齐 */
  white-space: nowrap; /* 防止文字换行 */
}

.filter-select {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 14px;
  border: none;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
  flex: 1; /* 让选择框占据剩余空间 */
  min-width: 120px; /* 设置选择框最小宽度 */
}
</style>
