<template>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
      <div v-for="category in categories" :key="category.id" class="relative bg-cover bg-center h-48" :style="{ backgroundImage: `url(${category.image})` }">
        <div class="absolute bottom-0 left-0 right-0 bg-black bg-opacity-50 text-white p-2 text-center">
          {{ category.title }}
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CategoryList',
    data() {
      return {
        categories: [
          { id: 1, title: 'Category 1', image: '/path/to/category1.jpg' },
          { id: 2, title: 'Category 2', image: '/path/to/category2.jpg' },
          { id: 3, title: 'Category 3', image: '/path/to/category3.jpg' },
        ],
      };
    },
  };
  </script>
  