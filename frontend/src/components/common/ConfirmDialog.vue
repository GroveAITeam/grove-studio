<script setup lang="ts">
import { ref, watch, defineProps, defineEmits } from 'vue';
import {
  Dialog,
  DialogContent,
  DialogFooter,
} from '@/components/ui/dialog'
import Button from '../ui/button/Button.vue';

const props = defineProps({
  modelValue: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    default: '确认操作'
  },
  message: {
    type: String,
    required: true
  },
  confirmButtonText: {
    type: String,
    default: '确认'
  },
  cancelButtonText: {
    type: String,
    default: '取消'
  }
});

const emit = defineEmits(['update:modelValue', 'confirm', 'cancel']);

const show = ref(props.modelValue);

watch(() => props.modelValue, (newValue) => {
  console.log(newValue,'----newValue')
  show.value = newValue;
});

watch(show, (newValue) => {
  emit('update:modelValue', newValue);
});

const confirm = () => {
  emit('confirm');
};

const cancel = () => {
  emit('cancel');
};
</script>

<template>
  <Dialog :open="show" :title="title">
    <DialogContent class="sm:max-w-[425px]">
      <p>{{ message }}</p>
      <DialogFooter>
        <Button class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600" @click="confirm">
        {{ confirmButtonText }}
      </Button>
      <Button class="bg-gray-300 text-black px-4 py-2 rounded ml-2 hover:bg-gray-400" @click="cancel">
        {{ cancelButtonText }}
      </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>