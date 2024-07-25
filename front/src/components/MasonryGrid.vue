<template>
    <masonry :cols="{ default: 3, 1100: 2, 700: 1 }" class="my-masonry-grid">
      <div v-for="item in items" :key="item.id" class="bg-white p-4 mb-4">
        <img :src="item.image" :alt="item.title" class="w-full h-auto" />
        <h2 class="text-xl font-semibold mt-2">{{ item.title }}</h2>
      </div>
    </masonry>
  </template>
  
  <script>
  import { ref, onMounted } from 'vue';
  import axios from 'axios';
  import Masonry from 'vue-masonry-css';
  
  export default {
    name: 'MasonryGrid',
    components: {
      Masonry,
    },
    setup() {
      const items = ref([]);
  
      onMounted(() => {
        axios.get('/api/works')
          .then(response => {
            items.value = response.data;
          })
          .catch(error => {
            console.error('Error fetching works:', error);
          });
      });
  
      return {
        items,
      };
    },
  };
  </script>
  
  <style scoped>
  .my-masonry-grid {
    display: flex;
    margin-left: -30px; /* gutter size offset */
    width: auto;
  }
  .my-masonry-grid_column {
    padding-left: 30px; /* gutter size */
    background-clip: padding-box;
  }
  .my-masonry-grid_column > div {
    background: grey;
    margin-bottom: 30px;
  }
  </style>
  