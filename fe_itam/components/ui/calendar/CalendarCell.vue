<script lang="ts" setup>
import { cn } from '@/lib/utils'
import { CalendarCell, type CalendarCellProps, useForwardProps } from 'radix-vue'
import { computed, type HTMLAttributes } from 'vue'

const props = defineProps<CalendarCellProps & { class?: HTMLAttributes['class'] }>()

const delegatedProps = computed(() => {
  const { class: _, ...delegated } = props

  return delegated
})

const forwardedProps = useForwardProps(delegatedProps)
</script>

<template>
  <CalendarCell
<<<<<<< HEAD
    :class="cn('relative h-9 w-9 p-0 text-center text-sm focus-within:relative focus-within:z-20 [&:has([data-selected])]:rounded-md [&:has([data-selected])]:bg-slate-100 [&:has([data-selected][data-outside-view])]:bg-slate-100/50 dark:[&:has([data-selected])]:bg-slate-800 dark:[&:has([data-selected][data-outside-view])]:bg-slate-800/50', props.class)"
=======
    :class="cn('relative p-0 text-center text-sm focus-within:relative focus-within:z-20 [&:has([data-selected])]:rounded-md [&:has([data-selected])]:bg-slate-100 [&:has([data-selected][data-outside-view])]:bg-slate-100/50 dark:[&:has([data-selected])]:bg-slate-800 dark:[&:has([data-selected][data-outside-view])]:bg-slate-800/50', props.class)"
>>>>>>> 988390618515c82a10424c730a828e66c000c19b
    v-bind="forwardedProps"
  >
    <slot />
  </CalendarCell>
</template>
