<template>
  <swiper
    :slidesPerView="3"
    :spaceBetween="30"
    :freeMode="true"
    :pagination="{
      clickable: true,
    }"
    :modules="modules"
    class="categoryList"
  >
    <swiper-slide v-for="(v, index) in images" :key="index">
      <div class="slide-content">
          <img :src="v" class="banner-image" />
          <div class="slide-caption">
            <h3 class="slide-title">这是标题</h3>
            <p class="slide-description">这是简介这是简介这是简介这是简介这是简介这是简介</p>
          </div>
        </div>
    </swiper-slide>
  </swiper>
</template>
<script>
  // Import Swiper Vue.js components
  import { Swiper, SwiperSlide } from 'swiper/vue';

  // Import Swiper styles
  import 'swiper/css';

  import 'swiper/css/free-mode';
  import 'swiper/css/pagination';

  // import required modules
  import { FreeMode, Pagination } from 'swiper/modules';

  export default {
    components: {
      Swiper,
      SwiperSlide,
    },
    setup() {
      return {
        modules: [FreeMode, Pagination],
        images: [
        'https://cdn.huelong.com/pic-go/2024/07/26/e55b093b82b6a64d9bd943d84157bbd1-07672f.jpg-pic_slimming',
        'https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg-pic_slimming',
        'https://cdn.huelong.com/pic-go/2024/07/26/9a55771884afafcd9aba1ac611c83663-90689d.png-pic_slimming',
        'https://cdn.huelong.com/pic-go/2024/07/26/e55b093b82b6a64d9bd943d84157bbd1-07672f.jpg-pic_slimming',
        'https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg-pic_slimming',
        'https://cdn.huelong.com/pic-go/2024/07/26/9a55771884afafcd9aba1ac611c83663-90689d.png-pic_slimming',
      ],
      };
    },
  };
</script>

<style>
.categoryList {
  width: 100%;
  height: 100%;
  margin: 20px auto;

}

.swiper-slide {
  text-align: center;
  font-size: 18px;
  background: #fff;

  /* Center slide text vertically */
  display: flex;
  justify-content: center;
  align-items: center;
}


.slide-content {
  position: relative;
  width: 100%;
  height: 100%;
}

.swiper-slide img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.slide-caption {
  position: absolute;
  bottom: 20px;
  left: 40px;
  background-color: rgba(0, 0, 0, 0.5);
  color: #fff;
  padding: 10px;
  border-radius: 5px;
}

.slide-title {
  margin: 0;
  font-size: 18px;
}

.slide-description {
  margin: 5px 0 0;
  font-size: 14px;
}

</style>