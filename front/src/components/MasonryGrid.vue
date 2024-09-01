<template>
  <div class="masonry_list" @scroll="handleScroll">

    <Waterfall :list="list" rowKey="id" imgSelector="pic_url" :width="320" :breakpoints="{
      1200: { rowPerView: 3 },
      800: { rowPerView: 2 },
      500: { rowPerView: 2 }
    }" :gutter="10" :lazyload="true">
      <template #item="{ item}">
        <div class="card">
          <LazyImg :url="item.pic_url+'-pic_slimming3'" />
          <p class="text">{{item.name}} {{item.desc}}</p>
        </div>
      </template>
    </Waterfall>
    <button v-if="reachedBottom" @click="loadMoreData">加载更多</button>

  </div>
</template>

<script>
import { LazyImg, Waterfall } from 'vue-waterfall-plugin-next';

import 'vue-waterfall-plugin-next/dist/style.css';
import axios from 'axios';


export default {
  data() {
        return {
            list: [],
            loading: true,
            currentPage: 1, // 当前页码
            pageSize: 30, // 每页显示的条目数
        };
    },
    created() {
        this.fetchImages();
    },
    methods: {
        fetchImages() {
            this.loading = true;

            axios.get(`http://127.0.0.1:8080/api/pic_gallery/list?page=${this.currentPage}&size=${this.pageSize}`)
                .then(response => {
                    this.list = response.data.data.list;
                })
                .catch(error => {
                    console.error("Error fetching images:", error);
                })
                .finally(() => {
                    this.loading = false;
                });
        },
        handleScroll() {
          const container = this.$el.querySelector('.masonry_list');
          if (container.scrollHeight - container.scrollTop === container.clientHeight) {
            this.reachedBottom = true;
          } else {
            this.reachedBottom = false;
          }
      }
    },
  components: {
    LazyImg,
    Waterfall,
  },
}
</script>

<style>
.card {
  border: 1px solid #25131f;
  padding: 10px;
  background-color: #50d195;
  border-radius: 5px;
}

.text {
  margin-top: 10px;
  text-align: center;
}
</style>