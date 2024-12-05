<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Cross1Icon } from '@radix-icons/vue'
</script>

<template>
    <Button variant="default" size="icon" class="bg-red-500 hover:bg-red-600">
        <Cross1Icon class="w-4 h-4" />
    </Button>
</template>