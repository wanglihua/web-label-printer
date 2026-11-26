import {createRouter, createWebHistory} from "vue-router";

import AdminHome from "./admin/admin-home/admin-home.vue";
import AdminLayout from "./admin/admin-layout.vue";
import PageNotFound from "./core/page-not-found.vue";

const routes = [
    {
        path: "",
        component: AdminLayout,
        children: [
            {path: "", component: AdminHome},
            {path: "/", component: AdminHome},
            {path: "/admin", component: AdminHome},
        ],
    },
    {path: "/:pathMatch(.*)*", component: PageNotFound},
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;
