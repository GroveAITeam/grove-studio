<script setup lang="ts">
import { ref, onMounted, watch } from 'vue';
import Chart from 'chart.js/auto';

interface Props {
  dates: string[];
  inputTokens: number[];
  outputTokens: number[];
}

const props = defineProps<Props>();
const chart = ref<Chart | null>(null);

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

  // 主题适配 - 获取当前主题的文本颜色
  const isDarkMode = document.documentElement.getAttribute('data-theme') === 'grovedark';
  const textColor = isDarkMode ? '#fff' : '#333';
  const gridColor = isDarkMode ? 'rgba(255, 255, 255, 0.1)' : 'rgba(0, 0, 0, 0.1)';

  chart.value = new Chart(ctx, {
    type: 'line',
    data: {
      labels: props.dates,
      datasets: [
        {
          label: '输入Token',
          data: props.inputTokens,
          borderColor: '#1677ff',
          backgroundColor: 'rgba(22, 119, 255, 0.1)',
          fill: true
        },
        {
          label: '输出Token',
          data: props.outputTokens,
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
            text: 'Token数量',
            color: textColor
          },
          grid: {
            color: gridColor
          },
          ticks: {
            color: textColor
          }
        },
        x: {
          grid: {
            color: gridColor
          },
          ticks: {
            color: textColor
          }
        }
      },
      plugins: {
        legend: {
          labels: {
            color: textColor
          }
        }
      }
    }
  });
};

// 监听主题变化
onMounted(() => {
  initChart();

  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.attributeName === 'data-theme') {
        initChart();
      }
    });
  });

  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ['data-theme']
  });
});

// 监听数据变化
watch(() => [props.dates, props.inputTokens, props.outputTokens], () => {
  initChart();
}, { deep: true });
</script>

<template>
  <div class="usage-chart bg-base-100 shadow-sm">
    <canvas id="usageChart"></canvas>
  </div>
</template>

<style scoped>
.usage-chart {
  border-radius: 12px;
  padding: 20px;
  margin-bottom: 24px;
  height: 300px;
}
</style>
