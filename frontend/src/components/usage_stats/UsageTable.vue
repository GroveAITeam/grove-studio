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

defineProps<{
  records: UsageRecord[];
}>();
</script>

<template>
  <div class="bg-base-100 rounded-lg p-6">
    <div class="overflow-x-auto">
      <table class="table table-zebra w-full">
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
          <tr v-for="record in records" :key="record.requestId">
            <td>{{ record.time }}</td>
            <td>
              <div class="flex items-center gap-2">
                <span class="text-xs px-2 py-1 rounded bg-base-200">
                  {{ record.provider }}
                </span>
                {{ record.model }}
              </div>
            </td>
            <td class="font-mono text-sm">{{ record.requestId }}</td>
            <td>
              <span
                :class="{
                  'text-success': record.status === 'success',
                  'text-error': record.status === 'error'
                }"
              >
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
  </div>
</template>
