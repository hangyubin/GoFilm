import { ApiGet, ApiPost } from '../utils/request';

// 基础API配置
export const baseApi = {
  // 首页相关
  index: {
    getIndex: () => ApiGet('/index'),
    clearCache: () => ApiGet('/cache/del'),
    getSiteConfig: () => ApiGet('/config/basic'),
    getNavCategories: () => ApiGet('/navCategory'),
  },
  
  // 电影相关
  film: {
    getDetail: (params: any) => ApiGet('/filmDetail', params),
    getPlayInfo: (params: any) => ApiGet('/filmPlayInfo', params),
    searchFilm: (params: any) => ApiGet('/searchFilm', params),
    getClassify: (params: any) => ApiGet('/filmClassify', params),
    searchByClassify: (params: any) => ApiGet('/filmClassifySearch', params),
  },
  
  // 用户相关
  user: {
    login: (data: any) => ApiPost('/login', data),
    logout: () => ApiGet('/logout'),
    changePassword: (data: any) => ApiPost('/changePassword', data),
  },
};
