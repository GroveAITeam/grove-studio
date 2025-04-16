<script setup lang="ts">
import { defineProps } from 'vue';

interface MonthData {
  value: string;
  label: string;
}

interface Props {
  totalCalls: number;
  inputTokens: string;
  outputTokens: string;
  successRate: string;
  selectedMonth: string;
  months: MonthData[];
}

const props = defineProps<Props>();

defineEmits<{
  (e: 'update:selectedMonth', value: string): void;
}>();
</script>

<template>
  <div class="api-section bg-base-100 shadow-sm">
    <div class="usage-stats bg-base-200/50 rounded-lg">
      <div class="usage-header">
        <span class="usage-title text-base-content/70">使用概览</span>
        <div class="period-selector">
          <select
            class="filter-select text-base-content/70 hover:text-base-content"
            :value="selectedMonth"
            @change="$emit('update:selectedMonth', ($event.target as HTMLSelectElement).value)"
          >
            <option v-for="month in months" :key="month.value" :value="month.value">
              {{ month.label }}
            </option>
          </select>
        </div>
      </div>
      <div class="usage-metrics">
        <div class="usage-metric">
          <span class="metric-value text-base-content">{{ totalCalls }}</span>
          <span class="metric-label text-base-content/70">总调用次数</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value text-base-content">{{ inputTokens }}</span>
          <span class="metric-label text-base-content/70">输入Token</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value text-base-content">{{ outputTokens }}</span>
          <span class="metric-label text-base-content/70">输出Token</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value text-base-content">{{ successRate }}</span>
          <span class="metric-label text-base-content/70">成功率</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.usage-stats {
  padding: 12px;
}

.usage-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.usage-title {
  font-size: 13px;
}

.usage-metrics {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  gap: 12px;
}

.usage-metric {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.metric-value {
  font-size: 16px;
  font-weight: 500;
}

.metric-label {
  font-size: 12px;
}

.api-section {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  padding: 20px;
  margin-bottom: 20px;
}

.filter-select {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 14px;
  border: none;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
</style>
