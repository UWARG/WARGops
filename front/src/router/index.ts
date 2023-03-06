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

router.beforeEach((to, _, next) => {
  // If the route does not require authentication, continue
  if (!to.meta.requiresAuth) {
    next();
    return;
  }
  
  console.log(useProfileStore().getLoggedIn);


  // If the user is not logged in, redirect to login
  if (to.name === "Login" && useProfileStore().getLoggedIn) {
    next({ name: "Home" });
    return;
  }

  // If the user is logged in, continue
  if (!useProfileStore().getLoggedIn) {
    next({ name: "Login" });
    return;
  }
  console.log(
    "You are logged in as: %c" + useProfileStore().profile.NickName,
    "color:blue;"
  );
  next();
});

export default router;
