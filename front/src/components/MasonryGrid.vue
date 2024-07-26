<template>
  <div>
    <Waterfall :list="list" rowKey="id" imgSelector="src" :width="320" :breakpoints="{
      1200: { rowPerView: 3 },
      800: { rowPerView: 2 },
      500: { rowPerView: 2 }
    }" :gutter="10" :lazyload="true">
      <template #item="{ item, url, index }">
        <div class="card">
          <LazyImg :url="url" />
          <p class="text">这是具体内容</p>
        </div>
      </template>
    </Waterfall>
  </div>
</template>

<script>
import { LazyImg, Waterfall } from 'vue-waterfall-plugin-next';
import 'vue-waterfall-plugin-next/dist/style.css';

export default {
  data() {
    return {
      list: [
        {
          id: 1,
          src: "https://cdn.huelong.com/pic-go/2024/01/24/652a6b2fe2a9d078ba4bb7e03a9bcba5-78dbf7.jpg",
        },
        {
          id: 2,
          src: "https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg",
        },
        {
          id: 3,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/5181ce8853a0dddcafa3d9cf91949318-8e70da.png",
        },
        {
          id: 4,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/147f87539521953e92efa0e696a5f3c6-4482be.png",
        },
        {
          id: 5,
          src: "https://cdn.huelong.com/pic-go/2024/02/18/a23fc98995f693ab30cba87846e93ad6-d52645.jpeg",
        },
        {
          id: 1,
          src: "https://cdn.huelong.com/pic-go/2024/01/24/652a6b2fe2a9d078ba4bb7e03a9bcba5-78dbf7.jpg",
        },
        {
          id: 2,
          src: "https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg",
        },
        {
          id: 3,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/5181ce8853a0dddcafa3d9cf91949318-8e70da.png",
        },
        {
          id: 4,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/147f87539521953e92efa0e696a5f3c6-4482be.png",
        },
        {
          id: 5,
          src: "https://cdn.huelong.com/pic-go/2024/02/18/a23fc98995f693ab30cba87846e93ad6-d52645.jpeg",
        },
        {
          id: 1,
          src: "https://cdn.huelong.com/pic-go/2024/01/24/652a6b2fe2a9d078ba4bb7e03a9bcba5-78dbf7.jpg",
        },
        {
          id: 2,
          src: "https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg",
        },
        {
          id: 3,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/5181ce8853a0dddcafa3d9cf91949318-8e70da.png",
        },
        {
          id: 4,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/147f87539521953e92efa0e696a5f3c6-4482be.png",
        },
        {
          id: 5,
          src: "https://cdn.huelong.com/pic-go/2024/02/18/a23fc98995f693ab30cba87846e93ad6-d52645.jpeg",
        },
        {
          id: 1,
          src: "https://cdn.huelong.com/pic-go/2024/01/24/652a6b2fe2a9d078ba4bb7e03a9bcba5-78dbf7.jpg",
        },
        {
          id: 2,
          src: "https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg",
        },
        {
          id: 3,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/5181ce8853a0dddcafa3d9cf91949318-8e70da.png",
        },
        {
          id: 4,
          src: "https://cdn.huelong.com/pic-go/2023/09/12/147f87539521953e92efa0e696a5f3c6-4482be.png",
        },
        {
          id: 5,
          src: "https://cdn.huelong.com/pic-go/2024/02/18/a23fc98995f693ab30cba87846e93ad6-d52645.jpeg",
        },
      ],
    };
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