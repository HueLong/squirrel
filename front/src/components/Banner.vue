<template>
  <swiper
    :cssMode="true"
    :navigation="true"
    :pagination="true"
    :mousewheel="true"
    :keyboard="true"
    :modules="modules"
    class="bannerSwiper"
  >
    <swiper-slide v-for="(image, index) in images" :key="index">
        <img :src="image" class="banner-image" />
      </swiper-slide>
  </swiper>
</template>
<script>
  // Import Swiper Vue.js components
  import { Swiper, SwiperSlide } from 'swiper/vue';

  // Import Swiper styles
  import 'swiper/css';

  import 'swiper/css/navigation';
  import 'swiper/css/pagination';

  // import required modules
  import { Navigation, Pagination, Mousewheel, Keyboard } from 'swiper/modules';

  export default {
    components: {
      Swiper,
      SwiperSlide,
    },
    setup() {
      return {
        modules: [Navigation, Pagination, Mousewheel, Keyboard],
        images: [
        'https://cdn.huelong.com/pic-go/2024/07/26/e55b093b82b6a64d9bd943d84157bbd1-07672f.jpg',
        'https://cdn.huelong.com/pic-go/2023/08/24/cf4769cb9f19bae83cfabacf5f44dfa4-1d65fe.jpg',
        'https://cdn.huelong.com/pic-go/2024/07/26/9a55771884afafcd9aba1ac611c83663-90689d.png',
      ],
      };
    },
  };
</script>
<style>
.bannerSwiper {
  margin-top: 75px;
  width: 100%;
  height: 400px;
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

.swiper-slide img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>