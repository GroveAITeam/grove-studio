<script lang="ts" setup>
import { ref } from 'vue'
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'
import { Icon } from '@iconify/vue';

const count = ref(0)
const data = reactive({
  name: "",
  resultText: "请输入你的名字: ",
})

const greet = () => {
  Greet(data.name).then(result => {
    data.resultText = result
  })
}

const features = [
  { title: "简单易用", description: "一键安装，双击即可使用", icon: "material-symbols:touch-app-outline" },
  { title: "隐私安全", description: "所有数据留在你的设备上，支持离线使用", icon: "material-symbols:security" },
  { title: "多种模型", description: "支持deepseek、Llama、千问等主流模型", icon: "material-symbols:model-training" },
  { title: "基础服务", description: "知识库、联网搜索、规则引擎等常用功能", icon: "material-symbols:cloud-done" },
  { title: "丰富应用", description: "总结、翻译、创作、服务部署等实用应用", icon: "material-symbols:apps" },
]
</script>

<template>
  <div class="max-w-3xl mx-auto pb-4">
    <!-- 头部欢迎区域 -->
    <div class="card-simple p-5 mb-6">
      <h1 class="text-2xl font-medium mb-3">Grove Studio - 本地AI平台</h1>
      <p class="text-muted mb-5">简单、安全、高效的本地AI应用平台</p>

      <div class="flex items-center gap-3 mb-4">
        <input id="name" v-model="data.name" autocomplete="off" placeholder="请输入你的名字"
               class="input input-bordered w-full max-w-xs" type="text"/>
        <button type="button" @click="greet" class="btn-simple">
          欢迎
        </button>
      </div>

      <div v-if="data.resultText !== '请输入你的名字: '"
           class="bg-success/10 text-success p-3 rounded-md">
        <span>{{ data.resultText }}</span>
      </div>
    </div>

    <div class="divider-line"></div>

    <!-- 功能特点 -->
    <h2 class="text-xl font-medium mb-4">平台特点</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
      <div v-for="feature in features" :key="feature.title"
           class="border border-base-200 rounded-md p-3">
        <div class="flex items-start">
          <div class="bg-primary/10 rounded-md p-2 mr-3">
            <Icon :icon="feature.icon" class="text-primary text-lg" />
          </div>
          <div>
            <h3 class="font-medium">{{ feature.title }}</h3>
            <p class="text-sm text-muted">{{ feature.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <div class="divider-line"></div>

    <!-- 计数器示例 -->
    <div class="card-simple p-4 text-center">
      <h3 class="mb-3">计数器示例</h3>
      <button type="button" @click="count++" class="btn-simple">
        计数器：{{ count }}
      </button>
    </div>
  </div>
</template>
