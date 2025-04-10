<script lang="ts" setup>
import { ref } from 'vue'
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

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
</script>

<template>
  <div class="space-y-6 mb-6">
    <p>首页</p>
    <div>
      <button type="button" @click="count++" class="btn btn-primary">计数器：{{ count }}</button>
    </div>
  </div>

  <div class="space-y-4">
    <div>{{ data.resultText }}</div>
    <div class="space-x-4">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button type="button" @click="greet" class="btn btn-secondary">欢迎</button>
    </div>
  </div>
</template>
