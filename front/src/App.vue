<template>
  <div id="app">
    <Header />
    <Banner />
    <CategoryList />
    <MasonryGrid />
  </div>
</template>

<script>
import Header from './components/Header.vue';
import Banner from './components/Banner.vue';
import CategoryList from './components/CategoryList.vue';
import MasonryGrid from './components/MasonryGrid.vue';

export default {
  name: 'App',
  components: {
    Header,
    Banner,
    CategoryList,
    MasonryGrid,
  },
};
</script>

<style>
/* Add any global styles here */
</style>
