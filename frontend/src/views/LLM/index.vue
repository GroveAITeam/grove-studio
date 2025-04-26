<script setup lang="ts">
</script>

<template>
  <view class="flex flex-col justify-center items-center w-full h-full llm">
    <view>
      对话模型页面
    </view>
    <router-link to="/llm/cloud">
      <view class="text-blue-500">
        跳转到云端模型
      </view>
    </router-link>
  </view>
</template>

<style scoped>
.llm {
}
</style>
