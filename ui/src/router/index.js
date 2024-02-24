import { createWebHistory, createRouter } from "vue-router";
import HelloWorld from "@/views/HelloWorld.vue";
import NotFound from "@/views/NotFound.vue"

const routes = [
    {
        path: "/",
        alias: ['/index.html'],
        name: "home",
        component: HelloWorld 
    },
    {
        path: "/:pathMatch(.*)*",
        name: "NotFound",
        component: NotFound
    }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
