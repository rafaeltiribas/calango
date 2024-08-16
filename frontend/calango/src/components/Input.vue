<script setup>
import { ref } from 'vue'
import axios from 'axios'

const url = ref("")
const shorturl = ref("")

const handleGo = async () => {
  try {
    const response = await axios.post(
        'http://localhost:8080/calango',
        new URLSearchParams({url: url.value}).toString(),
        {
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          }
        })
    console.log(response.data)
    shorturl.value = response.data.shortUrl
  } catch (error) {
    console.error('Error making request:', error)
  }
}
</script>

<template>
  <div class="card">
    <input class="url" type="text" v-model="url" placeholder="Insert your URL here..."/>
    <button class="shorten-button" @click="handleGo">GO!</button>

    <p v-if="shorturl">http://localhost:8080/{{shorturl}}</p>
  </div>

</template>

<style scoped>
</style>
