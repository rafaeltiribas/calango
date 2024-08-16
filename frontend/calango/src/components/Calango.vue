<script setup>
import { ref } from 'vue'
import Input from "./Input.vue";

defineProps({
  msg: String,
})

const url = ref("")
</script>

<template>
  <h1>{{ msg }}</h1>
  <p class="read-the-docs">Your favorite URL shorterner!</p>

  <Input />

  <p>
    Check out
    <a href="https://www.linkedin.com/in/rafaeltiribas/" target="_blank"
      >my LinkedIn</a
    > to see more projects I'm working on
  </p>
  <p class="read-the-docs">Click on the Calango logo to go to project's Github repository</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
