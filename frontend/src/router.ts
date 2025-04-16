import { createRouter, createWebHistory } from "vue-router";
import Home from "./components/Home.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", name: "Grove Studio", component: Home },
        { path: "/llm", name: '对话模型', component: () => import("./components/LLM.vue") },
        { path: "/api", name: 'API集成', component: () => import("./components/API.vue") },
        { path: "/api-usage", name: 'API使用统计', component: () => import("./components/APIUsage.vue")},
        { path: "/setting", name: '设置', component: () => import("./components/Setting.vue") },
    ],
});

export default router;
