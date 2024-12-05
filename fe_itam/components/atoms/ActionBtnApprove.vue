<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { CheckIcon } from '@radix-icons/vue'
</script>

<template>
    <Button variant="default" size="icon" class="bg-green-500 hover:bg-green-600">
        <CheckIcon class="w-4 h-4" />
    </Button>
</template>