<script setup lang="ts">
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import TusClient from "./components/TusClient.vue";
import logo from "./assets/logo.png";



const basePath = window['BasePath']

console.log(logo)
</script>

<template>
  <img alt="Vue logo" :src="`${basePath}${logo}`" />
  <TusClient />
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
