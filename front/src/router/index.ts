import * as VueRouter from 'vue-router';
import { useProfileStore } from '../store/profile';

export const routes = [
    { path: '/', name: "Home", component: () => import('../pages/Home.vue'), meta: { requiresAuth: true } },
    { path: '/login', name: "Login", component: () => import('../pages/Login.vue'), },
    { path: '/trasanactions/:account_id', name: "Transaction", component: () => import('../pages/Transactions.vue'), meta: { requiresAuth: true } },
];



const router = VueRouter.createRouter({
    history: VueRouter.createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    if (to.meta.requiresAuth) {
        if (useProfileStore().profile.username) {
            console.log("You are logged in as: %c" + useProfileStore().profile.username, 'color:blue;');
            next();
        } else {
            next({ name: 'Login' });
        }
    }
    else if (to.name === 'Login' && useProfileStore().profile.username) {
        next({ name: 'Home' });
    }
    else {
        next();
    }
});

export default router;  