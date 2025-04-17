<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import UsageOverview from '../../components/usage_stats/UsageOverview.vue';
import UsageFilters from '../../components/usage_stats/UsageFilters.vue';
import UsageChart from '../../components/usage_stats/UsageChart.vue';
import UsageTable from '../../components/usage_stats/UsageTable.vue';

// 类型定义
interface ApiData {
  id: string;
  name: string;
  provider: string;
  apiKey: string;
  defaultModel: string;
  baseUrl: string;
}

interface UsageRecord {
  time: string;
  model: string;
  provider: string;
  requestId: string;
  status: 'success' | 'error';
  responseTime: number;
  inputTokens: number;
  outputTokens: number;
}

interface MonthData {
  value: string;
  label: string;
}

// 响应式状态
const route = useRoute();
const router = useRouter();
const apiId = ref<string | null>(null);
const api = ref<ApiData | null>(null);
const selectedTimeRange = ref('30days');
const selectedModel = ref('all');
const selectedStatus = ref('all');
const selectedMonth = ref<string>('');
const months = ref<MonthData[]>([]);
const usageRecords = ref<UsageRecord[]>([]);

// 统计数据
const totalCalls = ref(8526);
const inputTokens = ref('1.25M');
const outputTokens = ref('3.76M');
const successRate = ref('98.6%');

// 图表数据
const chartDates = ref<string[]>([]);
const chartInputTokens = ref<number[]>([]);
const chartOutputTokens = ref<number[]>([]);

// 获取API信息
const getApiById = (id: string): ApiData | null => {
  const apis = JSON.parse(localStorage.getItem('grove_api_keys') || '[]');
  return apis.find((api: ApiData) => api.id === id) || null;
};

// 获取过去12个月
const getLast12Months = (): MonthData[] => {
  const monthList: MonthData[] = [];
  const now = new Date();

  for (let i = 0; i < 12; i++) {
    const date = new Date(now.getFullYear(), now.getMonth() - i, 1);
    monthList.push({
      value: date.toISOString().slice(0, 7),
      label: date.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long' })
    });
  }

  return monthList;
};

// 获取过去30天的日期
const getLast30Days = (): string[] => {
  const dates: string[] = [];
  for (let i = 29; i >= 0; i--) {
    const date = new Date();
    date.setDate(date.getDate() - i);
    dates.push(date.toLocaleDateString('zh-CN', { month: 'short', day: 'numeric' }));
  }
  return dates;
};

// 生成随机数据
const generateRandomData = (count: number, min: number, max: number): number[] => {
  return Array.from({ length: count }, () =>
    Math.floor(Math.random() * (max - min + 1) + min)
  );
};

// 生成模拟记录
const generateMockRecords = (count: number): UsageRecord[] => {
  const models = [
    { name: 'gpt-4', provider: 'openai' },
    { name: 'gpt-3.5-turbo', provider: 'openai' },
    { name: 'claude-3-opus', provider: 'anthropic' },
    { name: 'gemini-pro', provider: 'google' }
  ];

  return Array.from({ length: count }, () => {
    const model = models[Math.floor(Math.random() * models.length)];
    const inputTokens = Math.floor(Math.random() * 2000) + 100;
    const outputTokens = Math.floor(Math.random() * 3000) + 200;
    const status = Math.random() > 0.1 ? 'success' : 'error';

    return {
      time: new Date(Date.now() - Math.random() * 30 * 24 * 60 * 60 * 1000)
        .toLocaleString('zh-CN'),
      model: model.name,
      provider: model.provider,
      requestId: Math.random().toString(36).substring(2, 15),
      status: status,
      responseTime: Math.floor(Math.random() * 2000) + 500,
      inputTokens: inputTokens,
      outputTokens: outputTokens
    };
  });
};

// 筛选更新
const updateData = (): void => {
  // 重新生成模拟记录（实际应用中应该根据筛选条件请求API）
  usageRecords.value = generateMockRecords(20);

  // 更新图表数据
  chartDates.value = getLast30Days();
  chartInputTokens.value = generateRandomData(30, 1000, 5000);
  chartOutputTokens.value = generateRandomData(30, 2000, 8000);
};

// 初始化
onMounted(() => {
  // 从URL获取API ID
  apiId.value = route.query.id as string || null;

  if (!apiId.value) {
    router.push('/');
    return;
  }

  const apiData = getApiById(apiId.value);
  if (!apiData) {
    router.push('/');
    return;
  }

  api.value = apiData;

  // 初始化月份选择器
  months.value = getLast12Months();
  selectedMonth.value = months.value[0].value;

  // 生成模拟数据
  updateData();
});
</script>

<template>
  <!-- 使用统计概览 -->
  <UsageOverview
    v-model:selectedMonth="selectedMonth"
    :months="months"
    :totalCalls="totalCalls"
    :inputTokens="inputTokens"
    :outputTokens="outputTokens"
    :successRate="successRate"
  />

  <!-- 筛选器 -->
  <UsageFilters
    v-model:selectedTimeRange="selectedTimeRange"
    v-model:selectedModel="selectedModel"
    v-model:selectedStatus="selectedStatus"
    @filter="updateData"
  />

  <!-- 使用趋势图表 -->
  <UsageChart
    :dates="chartDates"
    :inputTokens="chartInputTokens"
    :outputTokens="chartOutputTokens"
  />

  <!-- 详细记录表格 -->
  <UsageTable :records="usageRecords" />
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
</style>
