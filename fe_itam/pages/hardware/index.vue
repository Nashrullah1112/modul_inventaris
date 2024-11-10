<script setup lang="ts">
import { Button } from '@/components/ui/button'

const route = useRoute()
</script>

<template>
  <div class="flex mb-2">
    <Button @click="navigateTo(route.path + '/add')">Tambah</Button>
  </div>

  <p>Tabel Data Hardware</p>
</template>