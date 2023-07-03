import { createRouter, createWebHashHistory } from "vue-router";
import Home from "@/views/Home.vue";
import About from "@/views/About.vue";
import Login from "@/views/Login.vue";
import Grade from "@/views/Grade.vue";
import XK from "@/views/XK.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: Home,
    },
    {
      path: '/home',
      redirect: '/'
    },
    {
      path: "/about",
      name: "about",
      component: About,
    },
    {
      path: "/login",
      name: "login",
      component: Login,
    },
    {
      path: "/grade",
      name: "grade",
      component: Grade,
    },
    {
      path: "/xk",
      name: "xk",
      component: XK,
    },
  ],
});

export default router;
