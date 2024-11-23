<script setup lang="ts">
const isOpen = useState("is-sidebar-open");
</script>

<template>
  <div class="w-full h-full">
    <!-- sidebar -->
    <LayoutsSidebar></LayoutsSidebar>

    <!-- main container -->
    <div class="pt-14 h-screen" :class="isOpen ? 'lg:ml-64' : ''">
      <div class="flex flex-col h-full overflow-y-auto bg-white">
        <!-- content -->
        <div class="px-6 py-4 flex-1">
          <slot />
        </div>

        <!-- footer -->
        <div
          v-if="false"
          class="flex justify-end py-2 px-3 w-full border-t text-slate-300"
        >
          <span class="text-xs">&copy; 2024</span>
        </div>
      </div>
    </div>
  </div>
</template>
