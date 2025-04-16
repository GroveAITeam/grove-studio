<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Chart from 'chart.js/auto';

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
const chart = ref<Chart | null>(null);

// 统计数据
const totalCalls = ref(8526);
const inputTokens = ref('1.25M');
const outputTokens = ref('3.76M');
const successRate = ref('98.6%');

// 获取API信息
const getApiById = (id: string): ApiData | null => {
  const apis = JSON.parse(localStorage.getItem('grove_api_keys') || '[]');
  return apis.find((api: ApiData) => api.id === id) || null;
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

  // 生成模拟记录
  usageRecords.value = generateMockRecords(20);

  // 初始化图表
  initChart();
});

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

// 初始化图表
const initChart = (): void => {
  // 获取Canvas元素
  const canvas = document.getElementById('usageChart') as HTMLCanvasElement;
  if (!canvas) return;

  const ctx = canvas.getContext('2d');
  if (!ctx) return;

  // 销毁旧图表
  if (chart.value) {
    chart.value.destroy();
  }

  // 生成示例数据
  const dates = getLast30Days();
  const inputTokens = generateRandomData(30, 1000, 5000);
  const outputTokens = generateRandomData(30, 2000, 8000);

  chart.value = new Chart(ctx, {
    type: 'line',
    data: {
      labels: dates,
      datasets: [
        {
          label: '输入Token',
          data: inputTokens,
          borderColor: '#1677ff',
          backgroundColor: 'rgba(22, 119, 255, 0.1)',
          fill: true
        },
        {
          label: '输出Token',
          data: outputTokens,
          borderColor: '#52c41a',
          backgroundColor: 'rgba(82, 196, 26, 0.1)',
          fill: true
        }
      ]
    },
    options: {
      responsive: true,
      interaction: {
        mode: 'index',
        intersect: false,
      },
      scales: {
        y: {
          type: 'linear',
          display: true,
          position: 'left',
          title: {
            display: true,
            text: 'Token数量'
          }
        }
      }
    }
  });
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
  // 更新图表
  initChart();
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
</script>

<template>
  <!-- 使用统计概览 -->
  <div class="api-section">
    <div class="usage-stats">
      <div class="usage-header">
        <span class="usage-title">使用概览</span>
        <div class="period-selector">
          <select class="filter-select" v-model="selectedMonth">
            <option v-for="month in months" :key="month.value" :value="month.value">
              {{ month.label }}
            </option>
          </select>
        </div>
      </div>
      <div class="usage-metrics">
        <div class="usage-metric">
          <span class="metric-value">{{ totalCalls }}</span>
          <span class="metric-label">总调用次数</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value">{{ inputTokens }}</span>
          <span class="metric-label">输入Token</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value">{{ outputTokens }}</span>
          <span class="metric-label">输出Token</span>
        </div>
        <div class="usage-metric">
          <span class="metric-value">{{ successRate }}</span>
          <span class="metric-label">成功率</span>
        </div>
      </div>
    </div>
  </div>

  <!-- 筛选器 -->
  <div class="usage-filters">
    <div class="filter-group">
      <span class="filter-label">时间范围：</span>
      <select class="filter-select" v-model="selectedTimeRange" @change="updateData">
        <option value="today">今天</option>
        <option value="yesterday">昨天</option>
        <option value="7days">近7天</option>
        <option value="30days">近30天</option>
        <option value="custom">自定义</option>
      </select>
    </div>
    <div class="filter-group">
      <span class="filter-label">模型：</span>
      <select class="filter-select" v-model="selectedModel" @change="updateData">
        <option value="all">全部模型</option>
        <option value="gpt-4">GPT-4</option>
        <option value="gpt-3.5">GPT-3.5</option>
        <option value="claude-3">Claude 3</option>
        <option value="gemini-pro">Gemini Pro</option>
      </select>
    </div>
    <div class="filter-group">
      <span class="filter-label">状态：</span>
      <select class="filter-select" v-model="selectedStatus" @change="updateData">
        <option value="all">全部状态</option>
        <option value="success">成功</option>
        <option value="error">失败</option>
      </select>
    </div>
  </div>

  <!-- 使用趋势图表 -->
  <div class="usage-chart">
    <canvas id="usageChart"></canvas>
  </div>

  <!-- 详细记录表格 -->
  <div class="usage-table">
    <div class="table-header">
      <h3 class="table-title">调用记录</h3>
    </div>
    <table>
      <thead>
      <tr>
        <th>时间</th>
        <th>模型</th>
        <th>请求ID</th>
        <th>状态</th>
        <th>响应时间</th>
        <th>输入Token</th>
        <th>输出Token</th>
      </tr>
      </thead>
      <tbody>
        <tr v-for="record in usageRecords" :key="record.requestId">
          <td>{{ record.time }}</td>
          <td>
            <div class="model-cell">
              <img :src="`/assets/images/providers/${record.provider}.svg`" class="model-icon" :alt="record.provider">
              {{ record.model }}
            </div>
          </td>
          <td>{{ record.requestId }}</td>
          <td>
            <span class="status-badge" :class="record.status === 'success' ? 'status-success' : 'status-error'">
              {{ record.status === 'success' ? '成功' : '失败' }}
            </span>
          </td>
          <td>{{ record.responseTime }}ms</td>
          <td>{{ record.inputTokens }}</td>
          <td>{{ record.outputTokens }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.usage-stats {
  background-color: #f8f9fa;
  border-radius: 6px;
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
  color: var(--text-secondary);
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
  color: var(--text-color);
}

.metric-label {
  font-size: 12px;
  color: var(--text-secondary);
}

.api-section {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border: 1px solid var(--border-color);
  padding: 20px;
  margin-bottom: 20px;
}

.usage-filters {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
  flex-wrap: wrap;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.filter-label {
  font-size: 14px;
  color: var(--text-secondary);
}

.filter-select {
  padding: 6px 12px;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 14px;
  color: var(--text-color);
  background-color: white;
}

.usage-chart {
  background: white;
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 24px;
  border: 1px solid var(--border-color);
  height: 300px;
}

.usage-table {
  background: white;
  border-radius: 12px;
  padding: 20px;
  border: 1px solid var(--border-color);
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.table-title {
  font-size: 16px;
  font-weight: 500;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid var(--border-color);
  font-size: 14px;
}

th {
  color: var(--text-secondary);
  font-weight: normal;
}

td {
  color: var(--text-color);
}

.status-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
}

.status-success {
  background-color: rgba(82, 196, 26, 0.1);
  color: #52c41a;
}

.status-error {
  background-color: rgba(255, 77, 79, 0.1);
  color: #ff4d4f;
}

.model-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.model-icon {
  width: 16px;
  height: 16px;
  opacity: 0.6;
}

.period-selector .filter-select {
  padding: 4px 8px;
  font-size: 14px;
  border: none;
  background: none;
  color: var(--text-secondary);
  cursor: pointer;
}

.period-selector .filter-select:hover {
  color: var(--text-color);
}

.period-selector .filter-select:focus {
  outline: none;
}

@media (prefers-color-scheme: dark) {
  .usage-chart,
  .usage-table {
    background-color: rgba(255, 255, 255, 0.05);
  }

  .filter-select {
    background-color: rgba(255, 255, 255, 0.05);
  }

  .usage-stats {
    background-color: rgba(255, 255, 255, 0.05);
  }

  .period-selector .filter-select {
    color: var(--text-secondary);
    background: none;
  }

  .period-selector .filter-select:hover {
    color: white;
  }
}
</style>
