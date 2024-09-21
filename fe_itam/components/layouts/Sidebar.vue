<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { HamburgerMenuIcon } from "@radix-icons/vue";

const isOpen = useState("is-sidebar-open", () => false);
const windowWidth = ref(0);

const toggleSidebar = () => (isOpen.value = !isOpen.value);

const updateWidth = () => {
  windowWidth.value = window.innerWidth;
};

onMounted(() => {
  windowWidth.value = window.innerWidth;
  window.addEventListener("resize", updateWidth);
});

onUnmounted(() => {
  window.removeEventListener("resize", updateWidth);
});

watch(windowWidth, (newVal) => {
  if (windowWidth.value >= 1280) isOpen.value = true;
  else isOpen.value = false;
});

const sidebarItem = {
  top: [
    {
      title: "Home",
      url: "/",
      icon: "",
    },
    {
      title: "Asset",
      url: "/eg/asset",
      icon: "",
    },
    {
      title: "DataTable Example",
      url: "/eg/datatable",
      icon: "",
    },
  ],
  bottom: [
    {
      title: "Profile",
      url: "",
      icon: "",
    },
  ],
};
</script>

<template>
  <div class="fixed top-0 left-0 right-0 h-14 shadow z-10"></div>

  <div
    class="fixed flex flex-col top-0 left-0 h-screen w-64 bg-white p-4 shadow-xl z-20 transition-transform"
    :class="!isOpen ? '-translate-x-64' : ''"
  >
    <!-- sidebar toggle btn -->
    <Button
      variant="ghost"
      size="icon"
      class="absolute top-2.5 -right-14"
      @click="toggleSidebar"
    >
      <HamburgerMenuIcon class="w-4 h-4" />
    </Button>

    <div class="w-full text-xl text-center">
      <h1>Modul Inventaris</h1>
    </div>

    <div class="flex flex-col flex-1 gap-2 mt-10">
      <div class="flex flex-col flex-1">
        <Button
          variant="ghost"
          class="text-left"
          v-for="item in sidebarItem.top"
          @click="navigateTo(item.url)"
        >
          {{ item.title }}
        </Button>
      </div>
      <div class="flex flex-col">
        <Button
          variant="ghost"
          class="text-left"
          v-for="item in sidebarItem.bottom"
          @click="navigateTo(item.url)"
        >
          {{ item.title }}
        </Button>
      </div>
    </div>
  </div>
</template>
