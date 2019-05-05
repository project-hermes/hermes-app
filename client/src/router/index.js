import Vue from 'vue';
import VueRouter from 'vue-router';
import {
    SignIn,
    NavView,
    UploadView,
    ThanksView,
    EmptyView,
    DivesListView,
    DiveDetailsView,
    DashboardView
} from '~/pages';
import store from '~/store';
Vue.use(VueRouter);
const readyPromise = store.getters['auth/readyPromise'];
const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/sign-in',
            name: 'signIn',
            component: SignIn,
            beforeEnter: (to, from, next) => {
                const isAuthorized = store.getters['auth/isAuthorized'];
                if (isAuthorized) {
                    return next('/dives');
                }
                next();
            }
        },
        {
            path: '/',
            component: NavView,
            meta: {
                requiresAuth: true
            },
            children: [
                {
                    path: '/dashboard',
                    name: 'dashboard',
                    component: DashboardView
                },
                {
                    path: '/dives',
                    name: 'dives',
                    component: DivesListView
                },
                {
                    path: '/dives/:diveId',
                    name: 'diveDetails',
                    component: DiveDetailsView
                },
                {
                    path: '/sensor',
                    name: 'sensor',
                    component: EmptyView
                },
                {
                    path: '/upload',
                    name: 'upload',
                    component: UploadView
                },
                {
                    path: '/upload/thank-you',
                    name: 'upload-thanks',
                    component: ThanksView
                },
                {
                    path: '/account-settings',
                    name: 'accountSettings',
                    component: EmptyView
                }
            ]
        }
    ]
});

router.beforeEach((to, from, next) => {
    readyPromise.then(() => {
        const isAuthorized = store.getters['auth/isAuthorized'];
        const requiresAuth = to.matched.some(
            record => record.meta.requiresAuth
        );
        if (requiresAuth && !isAuthorized) {
            next('/sign-in');
        } else {
            to.path === '/' ? next('dashboard') : next();
        }
    });
});

export default router;
