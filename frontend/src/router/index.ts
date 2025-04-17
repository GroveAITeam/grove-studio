import { createRouter, createWebHashHistory } from "vue-router";
import Chat from "../views/chat/index.vue";

const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        { path: "/", name: "Grove Studio", component: Chat },
        { path: "/llm", name: '对话模型', component: () => import("../views/llm/LLM.vue") },
        { path: "/llm/local", name: '本地模型', component: () => import("../views/llm/Local.vue") },
        { path: "/llm/cloud", name: '云端模型', component: () => import("../views/llm/Cloud.vue") },
        { path: "/setting", name: '设置', component: () => import("../views/settings/index.vue") },
    ],
});

export default router;
