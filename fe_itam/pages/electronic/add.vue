<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <div class="flex mb-2">
    <Button @click="router.back()" variant="secondary">Kembali</Button>
  </div>

  <ModuleElectronicForm
    type="add"
    :data="[]"
  />
</template>
