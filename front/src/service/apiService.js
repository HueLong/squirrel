import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://127.0.0.1:8080',
  headers: {
    'Content-Type': 'application/json',
  },
});

export default {
  getPicGalleryList() {
    return apiClient.get('/api/pic_gallery/list');
  },
  // 你还可以在这里定义其他的API调用方法
  
};
