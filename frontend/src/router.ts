import { createRouter, createWebHistory } from "vue-router";
import Home from "./components/Home.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", name: "Grove Studio", component: Home },
        { path: "/llm", name: '对话模型', component: () => import("./components/LLM.vue") },
        { path: "/setup", name: '设置', component: () => import("./components/Setup.vue") },
    ],
});

export default router;
