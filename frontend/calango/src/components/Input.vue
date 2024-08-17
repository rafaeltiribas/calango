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
const copyToClipboard = () => {
  const fullShortUrl = `http://localhost:8080/${shorturl.value}`
  navigator.clipboard.writeText(fullShortUrl).then(() => {
    alert('URL copied to clipboard!')
  })
}
</script>

<template>
  <div class="card">
    <input class="url" type="text" v-model="url" placeholder="Insert your URL here..."/>
    <button
        class="shorten_button"
        @click="handleGo"
    >GO!
    </button>

    <div v-if="shorturl" class="url-box">
      <p class="short-url">http://localhost:8080/{{ shorturl }}</p>
      <button
          class=""
          @click="copyToClipboard"
      >
        <i class="fas fa-copy"></i>
      </button>
    </div>
  </div>

</template>

<style scoped>
.shorten_button {
}
.shorten_button:hover {
  filter: drop-shadow(0 0 1em green);
}
.url-box {
  display: flex;
  align-items: center;
  border: 2px solid rgb(60, 63, 68);
  padding: 2px;
  border-radius: 10px;
  background-color: rgb(31, 32, 35);
  margin-top: 20px;
}

.short-url {
  margin: 0;
  font-size: 16px;
  color: white;
  flex-grow: 1;
}

.copy-button {
  background-color: #007bff;
  border: none;
  color: white;
  padding: 8px 12px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 14px;
  border-radius: 5px;
  cursor: pointer;
  margin-left: 10px;
}

.copy-button:hover {
  background-color: #0056b3;
}
</style>
