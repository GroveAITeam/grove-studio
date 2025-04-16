<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import * as echarts from 'echarts';

const props = defineProps<{
  dates: string[];
  inputTokens: number[];
  outputTokens: number[];
}>();

const chartContainer = ref<HTMLElement | null>(null);
let chart: echarts.ECharts | null = null;

const initChart = () => {
  if (!chartContainer.value) return;

  chart = echarts.init(chartContainer.value);
  updateChart();
};

const updateChart = () => {
  if (!chart) return;

  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      }
    },
    legend: {
      data: ['输入 Token', '输出 Token']
    },
    grid: {
      left: '3%',
      right: '4%',
      bottom: '3%',
      containLabel: true
    },
    xAxis: {
      type: 'category',
      data: props.dates,
      axisLabel: {
        rotate: 45
      }
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '输入 Token',
        type: 'line',
        smooth: true,
        data: props.inputTokens
      },
      {
        name: '输出 Token',
        type: 'line',
        smooth: true,
        data: props.outputTokens
      }
    ]
  };

  chart.setOption(option);
};

watch(
  () => [props.dates, props.inputTokens, props.outputTokens],
  () => updateChart(),
  { deep: true }
);

onMounted(() => {
  initChart();
  window.addEventListener('resize', () => chart?.resize());
});
</script>

<template>
  <div class="bg-base-100 rounded-lg p-6 mb-6">
    <div ref="chartContainer" style="width: 100%; height: 400px;"></div>
  </div>
</template>
