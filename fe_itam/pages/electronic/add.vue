<!-- <script setup lang="ts">
const router = useRouter()
</script>

<template>
  <ModuleElectronicForm
    type="add"
    :data="[]"
  />
</template> -->
<template>
  <div>
    <ElectronicAdd />
  </div>
</template>

<script setup>
import ElectronicAdd from "~/components/module/electronic/form.vue";
</script>
