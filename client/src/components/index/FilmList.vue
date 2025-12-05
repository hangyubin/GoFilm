<template>
  <div class="c_content" v-if="d.list">
    <template v-if="d.list.length > 0">
      <div class="item film-card" v-for="item in d.list" :style="{width: `calc(${d.width-1}%)`}">
        <div v-if="item.id != -99 && global.isMobile" class="hidden-md-and-up">
          <a :href="`/filmDetail?link=${item.id}`" class="default_image link_content">
            <div class="tag_group">
              <span class="cus_tag ">{{ item.year ? item.year.slice(0, 4) : '未知' }}</span>
              <span class="cus_tag ">{{ item.cName }}</span>
              <span class="cus_tag ">{{ item.area.split(',')[0] }}</span>
            </div>
            <span class="cus_remark hidden-md-and-up">{{ item.remarks }}</span>
            <img :src="item.picture" :alt="item.name?.split('[')[0]" @error="handleImg">
          </a>
          <a :href="`/filmDetail?link=${item.id}`" class="content_text_tag">{{ item.name.split("[")[0] }}</a>
          <span class="cus_remark hidden-md-and-down">{{ item.remarks }}</span>
        </div>

        <div v-if="!global.isMobile"  class="film-card-inner">
          <div class="film-card-front">
            <a :href="`/filmDetail?link=${item.id}`" class="link_content">
              <div class="tag_group">
                <span class="cus_tag ">{{ item.year ? item.year.slice(0, 4) : '未知' }}</span>
                <span class="cus_tag ">{{ item.cName }}</span>
                <span class="cus_tag ">{{ item.area.split(',')[0] }}</span>
              </div>
              <span class="cus_remark hidden-md-and-up">{{ item.remarks }}</span>
              <img :src="item.picture" :alt="item.name?.split('[')[0]" @error="handleImg">
            </a>
          </div>
          <div class="film-card-back" @click="toDetail(item.id)">
            <p class="card-title" >{{item.name}}</p>
            <p v-show="item.blurb != ''" class="card-blurb">{{ item.blurb }}</p>
            <p v-show="item.blurb == ''" class="card-blurb"> 暂无简介 </p>
            <el-button class="card-detail" :icon="Discount" color="#626aef" plain round @click="toDetail(item.id)" >详情</el-button>
          </div>
        </div>
        <a v-if="!global.isMobile" :href="`/filmDetail?link=${item.id}`" class="content_text_tag hidden-sm-and-down">{{ item.name.split("[")[0] }}</a>

      </div>
    </template>
    <el-empty v-if="d.list.length <= 0" style="padding: 10px 0;margin: 0 auto" description="暂无相关数据"/>
  </div>
</template>

<script setup lang="ts">

import {defineProps, inject, reactive, watchEffect} from 'vue'
import {Discount} from "@element-plus/icons-vue";

const props = defineProps({
  list: Array,
  col: Number,
})
const d = reactive({
  col: 0,
  list: Array,
  width: 0,
})

const global = inject('global')
// 图片加载失败事件
const handleImg = (e: Event) => {
  e.target.style.display = "none"
}

const toDetail = (id:any) =>{
  location.href = `/filmDetail?link=${id}`
}

// 监听父组件传递的参数的变化
watchEffect(() => {
  // 首先获取当前设备类型
  const userAgent = navigator.userAgent.toLowerCase();
  let isMobile = /mobile|android|iphone|ipad|phone/i.test(userAgent)
  // 如果是PC, 为防止flex布局最后一行元素不足出现错位, 使用空元素补齐list
  let c = isMobile ? 3 : props.col ? props.col : 0
  let l: any = props.list
  let len = l.length
  d.width = isMobile ? 31 : Math.floor(100 / c)
  if (len % c != 0) {
    for (let i = 0; i < c - len % c; i++) {
      let temp: any = {...l[0] as any}
      temp.id = -99
      l.push(temp)
    }
  }
  d.list = l
})


</script>

<style scoped>
.default_image {
  background: url("/src/assets/image/404.png") no-repeat center center;
  background-size: cover;
}

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

/*全局卡片样式变量*/
:root {
  --card-border-radius: 8px;
  --card-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
  --card-hover-shadow: 0 12px 32px rgba(0, 0, 0, 0.4);
  --card-transition: all 0.3s ease;
  --card-bg: rgba(30, 30, 36, 0.85);
  --card-hover-bg: rgba(35, 35, 42, 0.95);
  --text-primary: rgb(240, 240, 240);
  --text-secondary: rgb(180, 180, 180);
  --accent-color: #79bbff;
  --tag-bg: rgba(121, 187, 255, 0.2);
  --tag-text: #79bbff;
}

/*wrap*/
@media (max-width: 768px) {
  /*展示区域*/
  .c_content {
    width: 100%;
    display: flex;
    flex-flow: wrap;
    justify-content: space-between;
    gap: 12px;
  }

  .c_content .item {
    flex-basis: calc(50% - 6px);
    margin: 0;
    box-sizing: border-box;
    overflow: hidden;
    transition: var(--card-transition);
  }

  .c_content .item:hover {
    transform: translateY(-2px);
  }

  .item .link_content {
    padding-top: 135%;
    position: relative;
    border-radius: var(--card-border-radius);
    display: flex;
    width: 100%;
    background-size: cover;
    overflow: hidden;
    box-shadow: var(--card-shadow);
    transition: var(--card-transition);
  }

  .item .link_content:hover {
    box-shadow: var(--card-hover-shadow);
  }

  img {
    position: absolute;
    top: 0;
    left: 0;
    border-radius: var(--card-border-radius);
    object-fit: cover;
    width: 100%;
    height: 100%;
    transition: transform 0.5s ease;
  }

  .item .link_content:hover img {
    transform: scale(1.05);
  }

  .tag_group {
    position: absolute;
    top: 8px;
    left: 8px;
    display: flex;
    flex-wrap: wrap;
    gap: 4px;
    z-index: 10;
  }

  .cus_tag {
    flex-shrink: 0;
    white-space: nowrap;
    color: var(--tag-text);
    padding: 2px 6px;
    background: var(--tag-bg);
    font-size: 10px;
    border-radius: 12px;
    backdrop-filter: blur(10px);
  }

  .content_text_tag {
    font-size: 12px !important;
    color: var(--text-primary);
    width: 100% !important;
    max-height: 44px;
    line-height: 22px;
    padding: 8px 0 0 0 !important;
    text-align: left;
    display: -webkit-box !important;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
    font-weight: 500;
    transition: color 0.3s ease;
  }

  .item:hover .content_text_tag {
    color: var(--accent-color);
  }

  .cus_remark {
    z-index: 10;
    position: absolute;
    bottom: 0;
    display: block;
    width: 100%;
    font-size: 11px;
    color: var(--text-secondary);
    text-align: center;
    background: linear-gradient(transparent, rgba(0, 0, 0, 0.8));
    padding: 20px 8px 4px 8px;
    border-radius: 0 0 var(--card-border-radius) var(--card-border-radius);
  }
}

/*pc*/
@media (min-width: 768px) {
  .c_content {
    width: 100%;
    display: flex;
    flex-flow: wrap;
    justify-content: space-between;
    gap: 16px;
  }

  .item {
    margin-bottom: 0;
    box-sizing: border-box;
    transition: var(--card-transition);
  }

  .item:hover {
    transform: translateY(-4px);
  }

  .link_content {
    padding-top: 135%;
    background-size: cover;
    width: 100%;
    display: block;
    position: relative;
    margin-bottom: 8px;
    overflow: hidden;
    border-radius: var(--card-border-radius);
    box-shadow: var(--card-shadow);
    transition: var(--card-transition);
  }

  .link_content:hover {
    box-shadow: var(--card-hover-shadow);
  }

  img {
    position: absolute;
    top: 0;
    left: 0;
    border-radius: var(--card-border-radius);
    object-fit: cover;
    width: 100%;
    height: 100%;
    transition: transform 0.5s ease;
  }

  .link_content:hover img {
    transform: scale(1.05);
  }

  .tag_group {
    position: absolute;
    top: 12px;
    left: 12px;
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    z-index: 10;
    line-height: 18px;
  }

  .cus_tag {
    flex-shrink: 0;
    white-space: nowrap;
    color: var(--tag-text);
    padding: 3px 8px;
    background: var(--tag-bg);
    font-size: 12px;
    border-radius: 14px;
    backdrop-filter: blur(10px);
    transition: var(--card-transition);
  }

  .link_content:hover .cus_tag {
    background: rgba(121, 187, 255, 0.3);
  }

  .content_text_tag {
    display: block;
    font-size: 14px !important;
    color: var(--text-primary);
    width: 100% !important;
    padding: 4px 0 0 0 !important;
    text-align: left;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
    font-weight: 500;
    transition: color 0.3s ease;
  }

  .item:hover .content_text_tag {
    color: var(--accent-color);
  }

  .cus_remark {
    display: block;
    width: 100%;
    padding: 4px 0 0 0;
    font-size: 12px;
    color: var(--text-secondary);
    text-align: left;
    transition: color 0.3s ease;
  }

  .item:hover .cus_remark {
    color: rgba(121, 187, 255, 0.8);
  }
}
</style>


<style scoped>
.film-card {
  background-color: transparent;
  width: 100%;
  perspective: 1200px;
  font-family: 'Inter', sans-serif;
}

.film-card-inner {
  padding-top: 135%;
  position: relative;
  width: 100%;
  text-align: center;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
  transform-style: preserve-3d;
  box-shadow: var(--card-shadow);
  border-radius: var(--card-border-radius);
}

.film-card:hover .film-card-inner {
  transform: rotateY(180deg);
  box-shadow: var(--card-hover-shadow);
}

.film-card-front, .film-card-back {
  border-radius: var(--card-border-radius);
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  width: 100%;
  height: 100%;
  -webkit-backface-visibility: hidden;
  backface-visibility: hidden;
  overflow: hidden;
}

.film-card-front {
  border: none;
  background: url("/src/assets/image/404.png") no-repeat center center;
  background-size: cover;
}

.film-card-back {
  cursor: pointer;
  transform: rotateY(180deg);
  padding: 16px;
  background: linear-gradient(135deg, rgba(30, 30, 36, 0.95) 0%, rgba(35, 35, 42, 0.9) 100%);
  border: 1px solid rgba(121, 187, 255, 0.2);
  backdrop-filter: blur(10px);
}

.card-title {
  max-width: 90%;
  margin: 0 auto 12px;
  font-size: 15px;
  font-weight: 600;
  color: var(--text-primary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.card-blurb {
  margin-bottom: 20px;
  display: -webkit-box;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical;
  overflow: hidden;
  font-size: 13px;
  line-height: 1.5;
  color: var(--text-secondary);
}

.card-detail {
  position: absolute;
  width: 70%;
  left: 15%;
  bottom: 12px;
  font-size: 12px;
  padding: 6px 0;
  transition: var(--card-transition);
}

.card-detail:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(121, 187, 255, 0.3);
}
</style>