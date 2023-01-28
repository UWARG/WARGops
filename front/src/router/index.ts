
export const routes = [
    { path: '/', name: "Home", component: () => import('../pages/Home.vue') },
    { path: '/trasanactions/:account_id', name: "Transaction", component: () => import('../pages/Transactions.vue') },
];