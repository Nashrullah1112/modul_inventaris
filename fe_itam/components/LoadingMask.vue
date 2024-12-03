<script setup lang="ts">
import { Vue3Lottie } from 'vue3-lottie'
import loadingAnimation from '~/assets/raw/loading-anim.json'
const { isVisible } = useLoading()
</script>

<template>
  <div v-if="isVisible" class="loading-mask">
    <Vue3Lottie 
      :animationData="loadingAnimation" 
      :height="300" 
      :width="300" 
      :loop="true"
      :speed="0.8"
    />
  </div>
</template>

<style scoped>
.loading-mask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
</style>