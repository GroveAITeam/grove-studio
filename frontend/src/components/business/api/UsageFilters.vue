<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
  selectedTimeRange: string;
  selectedModel: string;
  selectedStatus: string;
}>();

const emit = defineEmits<{
  'update:selectedTimeRange': [value: string];
  'update:selectedModel': [value: string];
  'update:selectedStatus': [value: string];
  'filter': [];
}>();

// 本地状态
const localTimeRange = ref(props.selectedTimeRange);
const localModel = ref(props.selectedModel);
const localStatus = ref(props.selectedStatus);

// 监听props变化更新本地状态
watch(() => props.selectedTimeRange, (newVal) => {
  localTimeRange.value = newVal;
});

watch(() => props.selectedModel, (newVal) => {
  localModel.value = newVal;
});

watch(() => props.selectedStatus, (newVal) => {
  localStatus.value = newVal;
});

const timeRanges = [
  { value: '7days', label: '最近7天' },
  { value: '30days', label: '最近30天' },
  { value: '90days', label: '最近90天' }
];

const models = [
  { value: 'all', label: '所有模型' },
  { value: 'gpt-4', label: 'GPT-4' },
  { value: 'gpt-3.5-turbo', label: 'GPT-3.5 Turbo' },
  { value: 'claude-3-opus', label: 'Claude 3 Opus' },
  { value: 'gemini-pro', label: 'Gemini Pro' }
];

const statuses = [
  { value: 'all', label: '所有状态' },
  { value: 'success', label: '成功' },
  { value: 'error', label: '失败' }
];

const handleTimeRangeChange = (event: Event) => {
  const target = event.target as HTMLSelectElement;
  if (target) {
    emit('update:selectedTimeRange', target.value);
  }
};

const handleModelChange = (event: Event) => {
  const target = event.target as HTMLSelectElement;
  if (target) {
    emit('update:selectedModel', target.value);
  }
};

const handleStatusChange = (event: Event) => {
  const target = event.target as HTMLSelectElement;
  if (target) {
    emit('update:selectedStatus', target.value);
  }
};
</script>

<template>
  <div class="bg-base-100 rounded-lg p-6 mb-6">
    <div class="flex gap-4">
      <select
        :value="selectedTimeRange"
        class="select select-bordered min-h-[2.5rem] h-[2.5rem] text-sm px-4"
        @change="$emit('update:selectedTimeRange', ($event.target as HTMLSelectElement).value)"
      >
        <option v-for="range in timeRanges" :key="range.value" :value="range.value">
          {{ range.label }}
        </option>
      </select>

      <select
        :value="selectedModel"
        class="select select-bordered min-h-[2.5rem] h-[2.5rem] text-sm px-4"
        @change="$emit('update:selectedModel', ($event.target as HTMLSelectElement).value)"
      >
        <option v-for="model in models" :key="model.value" :value="model.value">
          {{ model.label }}
        </option>
      </select>

      <select
        :value="selectedStatus"
        class="select select-bordered min-h-[2.5rem] h-[2.5rem] text-sm px-4"
        @change="$emit('update:selectedStatus', ($event.target as HTMLSelectElement).value)"
      >
        <option v-for="status in statuses" :key="status.value" :value="status.value">
          {{ status.label }}
        </option>
      </select>

      <button class="btn btn-primary min-h-[2.5rem] h-[2.5rem] px-4">
        应用筛选
      </button>
    </div>
  </div>
</template>
