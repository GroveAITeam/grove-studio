<script setup lang="ts">
import { Button } from '@/components/ui/button'
import { navigator } from '@/router/navigator'

const back = () => {
  navigator.back()
}
</script>

<template>
  <div class="not-found-container">
    <h1 class="not-found-title">404</h1>
    <p class="not-found-message">抱歉，您访问的页面不存在</p>
    <Button @click="back">返回首页</Button>
  </div>
</template>

<style scoped>
.not-found-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
  padding: 0 1rem;
}

.not-found-title {
  font-size: 6rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 1rem;
}

.not-found-message {
  font-size: 1.5rem;
  color: #666;
  margin-bottom: 2rem;
}
</style>
