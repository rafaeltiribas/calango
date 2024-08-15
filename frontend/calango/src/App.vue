<script setup>
import Calango from './components/Calango.vue'
</script>

<template>
  <div>
    <a href="https://github.com/rafaeltiribas/calango" target="_blank">
      <img src="./assets/calango.png" class="logo" alt="Calango logo" />
    </a>
  </div>
  <Calango msg="Calango" />
</template>

<style scoped>
.logo {
  height: 8em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em green);
}
</style>
