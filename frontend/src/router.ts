import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/home/index.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", name: "Grove Studio", component: Home },
        { path: "/llm", name: '对话模型', component: () => import("./views/llm/LLM.vue") },
        { path: "/api", name: 'API集成', component: () => import("./views/api/index.vue") },
        { path: "/api-usage", name: 'API使用统计', component: () => import("./views/api/APIUsage.vue")},
        { path: "/setting", name: '设置', component: () => import("./views/settings/index.vue") },
    ],
});

export default router;
