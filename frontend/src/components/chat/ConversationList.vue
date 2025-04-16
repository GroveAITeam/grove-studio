<script lang="ts" setup>
import { Icon } from '@iconify/vue';
import { defineProps, defineEmits, computed } from 'vue';

interface Conversation {
  id: number;
  title: string;
  group: string;
  active: boolean;
}

const props = defineProps<{
  conversations: Conversation[]
}>();

const emit = defineEmits<{
  'create-new': []
  'switch': [conversation: Conversation]
  'delete': [index: number, event: Event]
}>();

// 根据组名分组对话
const groupedConversations = computed(() => {
  return props.conversations.reduce((groups: Record<string, Conversation[]>, conv) => {
    (groups[conv.group] = groups[conv.group] || []).push(conv);
    return groups;
  }, {});
});

// 事件处理
const handleCreateNew = () => {
  emit('create-new');
};

const handleSwitch = (conversation: Conversation) => {
  emit('switch', conversation);
};

const handleDelete = (index: number, event: Event) => {
  emit('delete', index, event);
};
</script>

<template>
  <aside class="w-56 flex flex-col overflow-hidden bg-base-200/50 border-r border-base-300">
    <div class="p-2">
      <!-- 新对话按钮 -->
      <button @click="handleCreateNew" class="w-full flex items-center justify-center gap-2 bg-primary hover:bg-primary-focus text-primary-content font-medium rounded-md py-2 text-sm transition-colors">
        <Icon icon="material-symbols:add" />
        <span>新对话</span>
      </button>
    </div>

    <!-- 对话列表 -->
    <div class="overflow-y-auto flex-1 p-2">
      <template v-for="(group, groupName) in groupedConversations">
        <div class="mb-2">
          <div class="text-xs uppercase px-2 mb-1 font-medium text-base-content/60">{{ groupName }}</div>
          <div v-for="(conversation, index) in group" :key="conversation.id"
               :class="['flex items-center gap-1 px-2 py-1.5 rounded-md cursor-pointer transition-colors',
               conversation.active ? 'bg-primary/10 text-primary font-medium' : 'text-base-content hover:bg-base-300/30']"
               @click="handleSwitch(conversation)">
            <Icon icon="material-symbols:chat" class="text-base" :class="conversation.active ? 'text-primary' : 'text-base-content/70'" />
            <span class="flex-1 truncate text-sm">{{ conversation.title }}</span>
            <button @click="(event) => handleDelete(props.conversations.indexOf(conversation), event)"
                    class="opacity-0 group-hover:opacity-100 hover:bg-base-300 p-0.5 rounded-md transition-opacity">
              <Icon icon="material-symbols:close" class="text-xs text-base-content/70" />
            </button>
          </div>
        </div>
      </template>
    </div>
  </aside>
</template>
