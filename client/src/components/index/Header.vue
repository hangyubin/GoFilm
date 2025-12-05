<template>
  <div class="header">
    <!-- 左侧logo以及搜索 -->
    <div class="nav_left">
      <!--        <img class="logo" src="/src/assets/logo.png">-->
      <!--<el-avatar class="logo" :size="45" :src="data.site.logo" alt="GoFilm"/>-->
      <a href="/" class="site">{{ data.site.siteName }}</a>
      <div class="search_group">
        <input v-model="keyword" @keydown="(e)=>{e.keyCode == 13 && searchFilm()}" placeholder="搜索 动漫,剧集,电影 "
               class="search"/>
        <el-button @click="searchFilm" :icon="Search"/>
      </div>
    </div>
    <!--右侧顶级分类导航 -->
    <div class="nav_right">
      <div class="nav_link">
        <a href="/">首页</a>
        <template v-for="n in data.nav">
          <a :href="`/filmClassify?Pid=${n.id}`">{{ n.name }}</a>
        </template>
      </div>
      <div class="history-link hidden-md-and-down" v-on:mouseenter="handleHistory(true)"
           v-on:mouseleave="handleHistory(false)">
        <a :href="`/filmClassify?Pid=${nav.variety.id}`">
          <b style="font-size: 22px;" class="iconfont icon-history"/>
        </a>
        <Transition name="fade-slide" duration="300">
          <div v-if="data.historyFlag" class="dropdown-container">
            <div class="history-h">
              <b class="iconfont icon-record history-h-icon"/>
              <span class="history-h-title">历史观看记录</span>
              <a v-if="data.historyList.length > 0" class="iconfont icon-clear1 history-del" @click="clearHistory"/>
            </div>
            <div v-if="data.historyList.length > 0" class="history-c">
              <a :href="h.link" class="history-c-item" v-for="h in data.historyList">
                <span class="history-c-item-t">{{ h.name }}</span>
                <span class="history-c-item-e">{{ h.episode }}</span>
              </a>
            </div>
            <el-empty style="padding: 10px 0;" v-else description="暂无观看记录"/>
          </div>
        </Transition>
      </div>
      <!--        <span style="color:#777; font-weight: bold">|</span>-->
      <a href="/search" class="hidden-md-and-up">
        <el-icon style="font-size: 18px">
          <Search/>
        </el-icon>
      </a>


    </div>
    <!--弹窗模块,显示按钮对应信息-->
  </div>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref, watch} from "vue";
import {useRouter} from "vue-router";
import {Search, CircleClose} from '@element-plus/icons-vue'
import {ElMessage} from "element-plus";
import {ApiGet} from "../../utils/request";
import {cookieUtil, COOKIE_KEY_MAP} from "../../utils/cookie";

// 搜索关键字
const keyword = ref<string>('')
// 弹窗隐藏显示
const data = reactive({
  historyFlag: false,
  historyList: [{}],
  nav: Array,
  site: Object,
})


// 加载观看历史记录信息
const handleHistory = (flag: boolean) => {
  data.historyFlag = flag
  if (flag) {
    // 获取cookie中的filmHistory
    let historyMap = cookieUtil.getCookie(COOKIE_KEY_MAP.FILM_HISTORY) ? JSON.parse(cookieUtil.getCookie(COOKIE_KEY_MAP.FILM_HISTORY)) : null
    let arr = []
    if (historyMap) {
      for (let k in historyMap) {
        arr.push(historyMap[k])
      }
      arr.sort((item1, item2) => item2.timeStamp - item1.timeStamp)
    }
    data.historyList = arr
  }
}
const clearHistory = () => {
  cookieUtil.clearCookie(COOKIE_KEY_MAP.FILM_HISTORY)
  data.historyList = []
}

// 从父组件获取当前路由对象
const router = useRouter()
// 影片搜索
const searchFilm = () => {
  if (keyword.value.length <= 0) {
    ElMessage.error({message: "请先输入影片名称关键字再进行搜索", duration: 1500})
    return
  }
  location.href = `/search?search=${keyword.value}`
  // router.push({path: '/search', query:{search: keyword.value}, replace: true})
}

// 导航栏挂载完毕时发送一次请求拿到对应的分类id数据
const nav = reactive({
  cartoon: {},
  film: {},
  tv: {},
  variety: {},
})

// 获取站点信息
const getBasicInfo = () => {
  ApiGet(`/config/basic`).then((resp: any) => {
    if (resp.code === 0) {
      data.site = resp.data
    } else {
      ElMessage.error({message: resp.msg})
    }
  })
}
onMounted(() => {
  ApiGet('/navCategory').then((resp: any) => {
    if (resp.code === 0) {
      data.nav = resp.data
    } else {
      ElMessage.error({message: "导航分类信息获取失败", duration: 1000})
    }
  })
  getBasicInfo()
})


</script>


<!--全局样式变量-->
<style>
:root {
  --header-bg: rgba(22, 22, 26, 0.95);
  --header-border: 1px solid rgba(121, 187, 255, 0.1);
  --header-shadow: 0 4px 24px rgba(0, 0, 0, 0.3);
  --nav-link-color: rgba(240, 240, 240, 0.8);
  --nav-link-hover: #79bbff;
  --search-bg: rgba(30, 30, 36, 0.8);
  --search-border: 1px solid rgba(121, 187, 255, 0.2);
  --search-focus: rgba(121, 187, 255, 0.3);
  --dropdown-bg: rgba(22, 22, 26, 0.98);
  --dropdown-border: 1px solid rgba(121, 187, 255, 0.2);
  --dropdown-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  --accent-color: #79bbff;
}
</style>

<!--移动端适配-->
<style scoped>
/*小尺寸时隐藏状态栏*/
@media (max-width: 768px) {
  .header {
    width: 100% !important;
    height: 60px;
    background: var(--header-bg) !important;
    backdrop-filter: blur(10px);
    box-shadow: var(--header-shadow) !important;
    border-bottom: var(--header-border) !important;
    position: sticky !important;
    top: 0 !important;
    z-index: 100 !important;
  }

  .nav_right {
    display: flex;
    width: 100%;
    justify-content: space-between;
    height: 60px;
  }

  .nav_link {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    height: 60px;
    width: 90%;
    overflow-x: auto;
    overflow-y: hidden;
    scrollbar-width: none;
    -ms-overflow-style: none;
  }

  .nav_link::-webkit-scrollbar {
    display: none;
  }

  .nav_link a {
    white-space: nowrap;
    color: var(--nav-link-color);
    padding: 0 12px;
    line-height: 60px;
    font-size: 14px;
    font-weight: 500;
    transition: all 0.3s ease;
    position: relative;
    margin-right: 8px;
  }

  .nav_link a::after {
    content: '';
    position: absolute;
    bottom: 15px;
    left: 50%;
    width: 0;
    height: 2px;
    background: var(--nav-link-hover);
    transition: all 0.3s ease;
    transform: translateX(-50%);
  }

  .nav_link a:hover {
    color: var(--nav-link-hover);
  }

  .nav_link a:hover::after {
    width: 20px;
  }

  .nav_right .hidden-md-and-up {
    color: var(--nav-link-color);
    padding: 0 12px;
    line-height: 60px;
    font-size: 18px;
    transition: color 0.3s ease;
  }

  .nav_right .hidden-md-and-up:hover {
    color: var(--nav-link-hover);
  }

  .nav_left {
    display: none !important;
  }
}
</style>

<style scoped>
/*PC端样式*/
@media (min-width: 768px) {
  .header {
    width: 78%;
    z-index: 100;
    height: 70px;
    line-height: 70px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: var(--header-bg);
    backdrop-filter: blur(10px);
    border-bottom: var(--header-border);
    box-shadow: var(--header-shadow);
    border-radius: 0 0 12px 12px;
    padding: 0 20px;
    position: sticky;
    top: 0;
  }

  .nav_left {
    display: flex;
    align-items: center;
    width: 50%;
    gap: 20px;
  }

  /*site标志样式*/
  .site {
    font-weight: 700;
    font-style: italic;
    font-size: 28px;
    background: linear-gradient(118deg, #79bbff, #988cd7, #c965b3, #e91a90);
    background-size: 200% 200%;
    -webkit-background-clip: text;
    background-clip: text;
    color: transparent;
    animation: gradientShift 4s ease-in-out infinite;
    transition: all 0.3s ease;
  }

  @keyframes gradientShift {
    0% { background-position: 0% 50%; }
    50% { background-position: 100% 50%; }
    100% { background-position: 0% 50%; }
  }

  .site:hover {
    transform: scale(1.05);
  }

  /*搜索栏*/
  .search_group {
    width: 100%;
    display: flex;
    gap: 0;
    align-items: center;
  }

  .search {
    flex: 1;
    background-color: var(--search-bg) !important;
    border: var(--search-border) !important;
    height: 44px;
    border-radius: 22px 0 0 22px;
    padding-left: 20px;
    color: var(--nav-link-color);
    font-size: 14px;
    font-weight: 500;
    transition: all 0.3s ease;
    outline: none;
  }

  .search::placeholder {
    font-size: 14px;
    color: rgba(240, 240, 240, 0.5);
  }

  .search:focus {
    background-color: rgba(30, 30, 36, 1) !important;
    border-color: var(--search-focus) !important;
    box-shadow: 0 0 0 3px var(--search-focus) !important;
    color: #ffffff;
  }

  .search_group button {
    height: 44px;
    background-color: var(--accent-color);
    color: #ffffff;
    border: none !important;
    border-radius: 0 22px 22px 0;
    font-size: 18px;
    padding: 0 20px;
    transition: all 0.3s ease;
    cursor: pointer;
  }

  .search_group button:hover {
    background-color: rgba(121, 187, 255, 0.9);
    transform: translateY(-1px);
    box-shadow: 0 4px 12px rgba(121, 187, 255, 0.3);
  }

  .nav_right {
    display: flex;
    height: 70px;
    flex-direction: row;
    align-items: center;
    gap: 4px;
  }

  .nav_link {
    display: flex;
    align-items: center;
    gap: 4px;
  }

  .nav_link a, .history-link a {
    min-width: 80px;
    height: 40px;
    line-height: 40px;
    margin: 0 4px;
    font-size: 14px;
    text-align: center;
    font-weight: 500;
    color: var(--nav-link-color);
    border-radius: 20px;
    transition: all 0.3s ease;
    position: relative;
    overflow: hidden;
  }

  .nav_link a::before {
    content: '';
    position: absolute;
    top: 0;
    left: -100%;
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, rgba(121, 187, 255, 0.1), transparent);
    transition: left 0.5s ease;
  }

  .nav_link a:hover::before {
    left: 100%;
  }

  .nav_link a:hover, .history-link a:hover {
    color: var(--nav-link-hover);
    background: rgba(121, 187, 255, 0.1);
    transform: translateY(-2px);
  }

  .logo {
    height: 40px;
    margin-top: 0;
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }

  /*history preview*/
  /*element-plus empty state color style*/
  :deep(.el-empty) {
    --el-empty-fill-color-1: rgba(155, 73, 231, 0.72);
    --el-empty-fill-color-2: #67d9e891;
    --el-empty-fill-color-3: rgb(106 19 187 / 72%);
    --el-empty-fill-color-4: #67d9e8;
    --el-empty-fill-color-5: #5abcc9;
    --el-empty-fill-color-6: #9fb2d9;
    --el-empty-fill-color-7: #61989f;
    --el-empty-fill-color-8: #697dc5;
    --el-empty-fill-color-9: rgb(43 51 63 / 44%);
  }

  .dropdown-container {
    position: absolute;
    top: calc(100% + 8px);
    left: 50%;
    font-size: 14px;
    color: var(--nav-link-color);
    min-width: 320px;
    max-width: 380px;
    height: auto;
    z-index: 1000;
    border-radius: 12px;
    overflow: hidden;
    background: var(--dropdown-bg);
    transform: translate3d(-50%, 0, 0);
    border: var(--dropdown-border);
    box-shadow: var(--dropdown-shadow);
    backdrop-filter: blur(15px);
  }

  .history-link {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    min-width: 60px;
    height: 40px;
    margin: 0 4px;
    font-size: 14px;
    text-align: center;
    font-weight: 500;
    color: var(--nav-link-color);
    border-radius: 20px;
    transition: all 0.3s ease;
  }

  .history-link:hover {
    background: rgba(121, 187, 255, 0.1);
    transform: translateY(-2px);
  }

  .history-h {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px 20px;
    border-bottom: 1px solid rgba(121, 187, 255, 0.1);
    background: rgba(30, 30, 36, 0.8);
  }

  .history-h-icon {
    font-size: 20px;
    color: var(--accent-color);
    margin-right: 8px;
  }

  .history-h-title {
    font-size: 16px;
    font-weight: 600;
    color: #ffffff;
  }

  .history-del {
    font-size: 18px;
    color: rgba(240, 240, 240, 0.5);
    transition: all 0.3s ease;
    cursor: pointer;
  }

  .history-del:hover {
    color: var(--accent-color);
    transform: scale(1.1);
  }

  .history-c {
    max-height: 240px;
    overflow-y: auto;
    margin: 0;
    padding: 8px 0;
  }

  /* 自定义滚动条 */
  .history-c::-webkit-scrollbar {
    width: 6px;
  }

  .history-c::-webkit-scrollbar-track {
    background: rgba(30, 30, 36, 0.8);
  }

  .history-c::-webkit-scrollbar-thumb {
    background: rgba(121, 187, 255, 0.5);
    border-radius: 3px;
  }

  .history-c::-webkit-scrollbar-thumb:hover {
    background: rgba(121, 187, 255, 0.7);
  }

  .history-c .history-c-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 20px;
    margin: 0;
    transition: all 0.3s ease;
    position: relative;
    cursor: pointer;
  }

  .history-c-item:hover {
    background: rgba(121, 187, 255, 0.1);
  }

  .history-c-item-t {
    flex: 1;
    text-align: left;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    color: var(--nav-link-color);
    font-weight: 500;
    transition: color 0.3s ease;
  }

  .history-c-item:hover .history-c-item-t {
    color: var(--nav-link-hover);
  }

  .history-c-item-e {
    color: rgba(240, 240, 240, 0.5);
    font-size: 12px;
    margin-left: 16px;
    white-space: nowrap;
  }

  /* 移除时间线样式，简化设计 */
  .history-c-item::before,
  .history-c-item::after {
    display: none;
  }

  /* 淡入淡出动画 */
  .fade-slide-enter-active,
  .fade-slide-leave-active {
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  }

  .fade-slide-enter-from,
  .fade-slide-leave-to {
    opacity: 0;
    transform: translate3d(-50%, -10px, 0);
  }

  /* 首页链接高亮 */
  .nav_link a[href="/"] {
    color: var(--accent-color);
    background: rgba(121, 187, 255, 0.1);
  }
}
</style>