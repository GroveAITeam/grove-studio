import { createRouter, createWebHistory } from "vue-router";
import Home from "./components/Home.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: "/", name: "home", component: Home },
        { path: "/users", component: () => import("./components/Users.vue") },
        { path: "/setup", component: () => import("./components/Setup.vue") },
    ],
});

export default router;
