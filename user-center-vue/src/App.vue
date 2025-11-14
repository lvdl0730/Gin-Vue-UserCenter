<script setup>
import Header from "./components/Header.vue";
</script>

<template>
  <div class="basic-bg">
    <a-layout class="basic-Layout">
      <a-layout-header class="header">
        <Header />
      </a-layout-header>

      <a-layout-content class="content">
        <router-view />
      </a-layout-content>

      <a-layout-footer class="footer">
        这是底部栏
      </a-layout-footer>
    </a-layout>
  </div>
</template>

<style scoped>
:global(body){
  background-image: url("./assets/background.jpg");
  background-repeat: no-repeat;
  background-size: cover;
  background-position: center center;
}
.basic-Layout{
  background:transparent;
}
.content{
  padding: 20px;
  margin-bottom: 20px;
}
.header{
  background: #fff;
  color: unset;
  padding-inline-start: 20px;
}
.footer{
  background: linear-gradient(to left, #e86f18,white);
  text-align: center;
  position: fixed;
  border: black solid 1px;
  height: 20px;
  bottom: 0;
  left: 0;
  right: 0;
}

</style>
