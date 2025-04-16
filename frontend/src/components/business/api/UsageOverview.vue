<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
  selectedMonth: string;
  months: Array<{ value: string; label: string }>;
  totalCalls: number;
  inputTokens: string;
  outputTokens: string;
  successRate: string;
}>();

const emit = defineEmits<{
  'update:selectedMonth': [value: string];
}>();

// 本地状态
const localMonth = ref(props.selectedMonth);

// 监听props变化更新本地状态
watch(() => props.selectedMonth, (newVal) => {
  localMonth.value = newVal;
});

const handleMonthChange = (event: Event) => {
  const target = event.target as HTMLSelectElement;
  if (target) {
    emit('update:selectedMonth', target.value);
  }
};
</script>

<template>
  <div class="bg-base-100 rounded-lg p-6 mb-6">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-xl font-semibold">使用统计</h2>
      <select
        :value="selectedMonth"
        class="select select-bordered min-h-[2.5rem] h-[2.5rem] text-sm px-4"
        @change="$emit('update:selectedMonth', ($event.target as HTMLSelectElement).value)"
      >
        <option v-for="month in months" :key="month.value" :value="month.value">
          {{ month.label }}
        </option>
      </select>
    </div>

    <div class="grid grid-cols-4 gap-4">
      <div class="stat bg-base-200/50 rounded-lg p-4">
        <div class="stat-title text-base-content/70">总调用次数</div>
        <div class="stat-value text-2xl">{{ totalCalls }}</div>
      </div>
      <div class="stat bg-base-200/50 rounded-lg p-4">
        <div class="stat-title text-base-content/70">输入 Token</div>
        <div class="stat-value text-2xl">{{ inputTokens }}</div>
      </div>
      <div class="stat bg-base-200/50 rounded-lg p-4">
        <div class="stat-title text-base-content/70">输出 Token</div>
        <div class="stat-value text-2xl">{{ outputTokens }}</div>
      </div>
      <div class="stat bg-base-200/50 rounded-lg p-4">
        <div class="stat-title text-base-content/70">成功率</div>
        <div class="stat-value text-2xl">{{ successRate }}</div>
      </div>
    </div>
  </div>
</template>
