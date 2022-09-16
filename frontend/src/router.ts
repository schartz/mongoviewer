import { createRouter, createWebHashHistory } from "vue-router";
const routes = [
    { path: '/', name: 'welcome', component: () => import('/@src/components/HelloWorld.vue') },
    { path: '/dblist', name:'dblist', component: () => import('/@src/components/dblist.vue') },
]
export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes,
})