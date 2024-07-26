<template>
  <div class="header">
    <img src="https://cdn.huelong.com/pic-go/2024/07/26/e55b093b82b6a64d9bd943d84157bbd1-07672f.jpg" alt="Logo"
      class="logo" />
    <div class="navbar">登陆</div>
  </div>
</template>

<script>
export default {
  name: 'Header',
};
</script>

<style>
.header {
  position: fixed;
  top: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #fff;
  border-bottom: 1px solid #ddd;
  padding: 10px 0px 20px 0 px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  z-index: 1000;
}

.logo {
  height: 40px;
  flex-shrink: 0;
  margin: 10px;
}

.navbar {
  flex-grow: 1;
  text-align: right;
  padding-left: 20px;
  font-size: 16px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin-right: 15px;
}
</style>