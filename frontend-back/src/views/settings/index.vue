<script lang="ts" setup>
import { Icon } from '@iconify/vue'

// 保存设置
const saveSettings = () => {
  console.log('设置已保存')
}
</script>

<template>
  <section class="space-y-6">
    <!-- 保存按钮 -->
    <div class="flex justify-end mt-6">
      <button class="btn btn-primary" @click="saveSettings">
        <Icon icon="material-symbols:save" class="mr-2" />
        保存设置
      </button>
    </div>
  </section>
</template>
