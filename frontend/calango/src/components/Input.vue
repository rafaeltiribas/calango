<script setup>
import { ref } from 'vue'

const url = ref("")
</script>

<template>
  <div class="card">
    <input class="url" type="text" v-model="url" placeholder="Insert your URL here..." />
    <button class="shorten-button">GO!</button>
  </div>
</template>

<style scoped>

</style>