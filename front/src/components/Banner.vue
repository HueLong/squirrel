<template>
    <div class="w-full h-64 bg-cover bg-center" :style="{ backgroundImage: 'url(/path/to/banner.jpg)' }">
      <div class="bg-black bg-opacity-50 h-full flex items-center justify-center">
        <h1 class="text-white text-4xl font-bold">Banner Title</h1>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Banner',
  };
  </script>
  