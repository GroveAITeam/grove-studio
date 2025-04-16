<script setup lang="ts">
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

interface Props {
  records: UsageRecord[];
}

defineProps<Props>();
</script>

<template>
  <div class="usage-table bg-base-100 shadow-sm">
    <div class="table-header">
      <h3 class="table-title text-base-content">调用记录</h3>
    </div>
    <table>
      <thead>
        <tr>
          <th class="text-base-content/70">时间</th>
          <th class="text-base-content/70">模型</th>
          <th class="text-base-content/70">请求ID</th>
          <th class="text-base-content/70">状态</th>
          <th class="text-base-content/70">响应时间</th>
          <th class="text-base-content/70">输入Token</th>
          <th class="text-base-content/70">输出Token</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="record in records" :key="record.requestId">
          <td class="text-base-content">{{ record.time }}</td>
          <td class="text-base-content">
            <div class="model-cell">
              <img :src="`/assets/images/providers/${record.provider}.svg`" class="model-icon" :alt="record.provider">
              {{ record.model }}
            </div>
          </td>
          <td class="text-base-content">{{ record.requestId }}</td>
          <td>
            <span class="status-badge" :class="record.status === 'success' ? 'status-success' : 'status-error'">
              {{ record.status === 'success' ? '成功' : '失败' }}
            </span>
          </td>
          <td class="text-base-content">{{ record.responseTime }}ms</td>
          <td class="text-base-content">{{ record.inputTokens }}</td>
          <td class="text-base-content">{{ record.outputTokens }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.usage-table {
  border-radius: 12px;
  padding: 20px;
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
  border-bottom: 1px solid;
  border-color: theme('colors.base-300/20');
  font-size: 14px;
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
}
</style>
