import * as VueRouter from "vue-router";
import { useProfileStore } from "../store/profile";
import Home from "../pages/NewHome.vue";
import Login from "../pages/Login.vue";
import Transactions from "../pages/Transactions.vue";
import NotFound from "../pages/NotFound.vue";

export const routes = [
  { path: "/", name: "Home", component: Home, meta: { requiresAuth: true } },
  { path: "/login", name: "Login", component: Login },
  {
    path: "/trasanactions/:account_id",
    name: "Transaction",
    component: Transactions,
    meta: { requiresAuth: true },
  },
  { path: "/:pathMatch(.*)*", name: "NotFound", component: NotFound },
];

const router = VueRouter.createRouter({
  history: VueRouter.createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    if (useProfileStore().profile.username) {
      console.log(
        "You are logged in as: %c" + useProfileStore().profile.username,
        "color:blue;"
      );
      next();
    } else {
      next({ name: "Login" });
    }
  } else if (to.name === "Login" && useProfileStore().profile.username) {
    next({ name: "Home" });
  } else {
    next();
  }
});

export default router;

