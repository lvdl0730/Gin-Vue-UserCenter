import { createRouter, createWebHistory } from "vue-router";

import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import UserManage from "../views/UserManage.vue";
import WebNav from "../views/WebNav.vue";

const routes = [
    {path: "/", component: Home},
    {path: "/login", component: Login},
    {path: "/register", component: Register},
    {path: "/manage", component: UserManage},
    {path:"/nav", component: WebNav},
];

const router = createRouter({
    history:createWebHistory(),
    routes,
})
export default router;