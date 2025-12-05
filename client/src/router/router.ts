import {
    createRouter,
    createWebHistory,
} from "vue-router";

import {getToken} from "../utils/token";

// 2. 定义一个路由
const routes = [
    {
        path: '/',
        component: () => import("../views/IndexHome.vue"),
        redirect: '/index',
        children: [
            {path: 'index', component: () => import("../views/index/Home.vue")},
            {path: 'filmDetail', component: () => import("../views/index/FilmDetails.vue")},
            {path: 'play', component: () => import("../views/index/Play.vue")},
            {path: 'search', component: () => import("../views/index/SearchFilm.vue")},
            {path: 'filmClassify', component: () => import("../views/index/FilmClassify.vue")},
            {path: 'filmClassifySearch', component: () => import("../views/index/FilmClassifySearch.vue")},
            {path: '/custom/player', component: () => import("../views/index/CustomPlay.vue")},
            {path: '/history', component: () => import("../views/index/FilmHistory.vue")},
        ]
    },
    {path: '/login', component: () => import("../views/Login.vue")},
    {
        path: '/manage',
        component: () => import("../views/manage/ManageHome.vue"),
        redirect: '/manage/index',
        children: [
            {path: 'index', component: () => import("../views/manage/Index.vue")},
            {path: 'collect/index', component: () => import("../views/manage/collect/CollectManage.vue")},
            {path: 'collect/record', component: () => import("../views/manage/collect/FailureRecord.vue")},
            {path: 'system/webSite', component: () => import("../views/manage/system/SiteConfig.vue")},
            {path: 'system/banners', component: () => import("../views/manage/system/Banners.vue")},
            {path: 'cron/index', component: () => import("../views/manage/cron/CronManage.vue")},
            {path: 'file/upload', component: () => import("../views/manage/file/FileUpload.vue")},
            {path: 'file/gallery', component: () => import("../views/manage/file/Temp.vue")},
            {path: 'film', component: () => import("../views/manage/film/Film.vue")},
            {path: 'film/class', component: () => import("../views/manage/film/FilmClass.vue")},
            {path: 'film/add', component: () => import("../views/manage/film/FilmAdd.vue")},
            {path: 'film/detail', component: () => import("../views/manage/file/Temp.vue")},

        ]
    },
    {path: `/:pathMatch(.*)*`, component: () => import('../views/error/Error404.vue')},
]

// 创建路由实例并传递 routes配置
const router = createRouter({
    history: createWebHistory(),
    routes
})

// 添加全局前置守卫拦截未登录的跳转
router.beforeEach((to, from, next) =>{
    // 如果访问的是 /manage 下的路由, 且 token信息为空 则跳转到登录界面
    let matchPath = new RegExp(/^\/manage\//).test(to.path)
    let token = getToken()
    if ( matchPath && !token ) {
        next('/login')
    } else {
        next()
    }
})




export {router}

