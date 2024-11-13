<<<<<<< HEAD
<template>
  <div>
    <ElectronicAdd />
  </div>
</template>

<script setup>
import ElectronicAdd from "@/components/module/electronic/add.vue";
</script>
=======
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
>>>>>>> 988390618515c82a10424c730a828e66c000c19b
