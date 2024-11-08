<script setup lang="ts">
import { mdiAccount, mdiLogout } from "@mdi/js";
import { Button } from "@/components/ui/button";
import {
  HamburgerMenuIcon,
  HomeIcon,
  TableIcon,
  StackIcon,
  PersonIcon,
  ArchiveIcon,
  GearIcon,
  LayersIcon,
} from "@radix-icons/vue";

const route = useRoute();

// State management
const isOpen = useState("is-sidebar-open", () => false);
const windowWidth = ref(0);

// Fungsi toggle sidebar
const toggleSidebar = () => (isOpen.value = !isOpen.value);

// Fungsi untuk update lebar window
const updateWidth = () => {
  windowWidth.value = window.innerWidth;
};

// Lifecycle hooks
onMounted(() => {
  updateWidth();
  window.addEventListener("resize", updateWidth);
});

onUnmounted(() => {
  window.removeEventListener("resize", updateWidth);
});

// Watch perubahan lebar window untuk responsive sidebar
watch(windowWidth, () => {
  isOpen.value = windowWidth.value >= 1280;
});

// Data struktur menu sidebar yang lebih dinamis
const sidebarItem = reactive({
  top: [
    {
      title: "Dashboard",
      url: "/",
      icon: HomeIcon,
      dropdown: false,
    },
    {
      title: "Registrasi Inventaris",
      icon: StackIcon,
      dropdown: true,
      isOpen: ref(false),
      children: [
        {
          title: "Perangkat Elektronik",
          url: "/electronic/",
        },
        {
          title: "Hardware",
          url: "/hardware/form",
        },
        {
          title: "Software Lisensi",
          url: "/software_license/form",
        },
        {
          title: "Software Aplikasi",
          url: "/aplikasi/form",
        },
      ],
    },
    {
      title: "Storage",
      icon: ArchiveIcon,
      dropdown: true,
      isOpen: ref(false),
      children: [
        {
          title: "Data Perangkat Elektronik",
          url: "/electronic/data",
        },
        {
          title: "Data Hardware",
          url: "/hardware/data",
        },
        {
          title: "Data Software Lisensi",
          url: "/software_license/data",
        },
        {
          title: "Data Software Aplikasi",
          url: "/software/data",
        },
        {
          title: "Tambah Pengguna",
          url: "/users/add",
        },
        {
          title: "Tambah Vendor",
          url: "/vendors/add",
        },
      ],
    },
    {
      title: "Approval",
      url: "/approval",
      icon: GearIcon,
      dropdown: false,
    },
    {
      title: "Pegawai",
      icon: PersonIcon,
      dropdown: true,
      isOpen: ref(false),
      children: [
        {
          title: "Tambah Pegawai",
          url: "/employee/add",
        },
        {
          title: "Data Pegawai",
          url: "/employee/data",
        },
      ],
    },
  ],
  bottom: [
    {
      title: "Profile",
      url: "/profile",
      icon: mdiAccount,
    },
    {
      title: "Logout",
      url: "/logout",
      icon: mdiLogout,
    },
  ],
});

// Helper functions
const isActive = (url: string): boolean => {
  return route.path === url;
};

const handleNavigation = (url: string) => {
  if (url) navigateTo(url);
};

const toggleDropdown = (item: any) => {
  if (item.dropdown) {
    item.isOpen = !item.isOpen;
  }
};
</script>

<template>
  <nav
    class="fixed top-0 left-0 right-0 h-14 border-b bg-background z-10"
  ></nav>

  <aside
    class="fixed flex flex-col top-0 left-0 h-screen w-64 bg-background border-r p-4 z-20 transition-transform duration-300"
    :class="{ '-translate-x-64': !isOpen }"
  >
    <!-- Toggle Button -->
    <Button
      variant="ghost"
      size="icon"
      class="absolute top-2.5 -right-14 hover:bg-primary hover:text-primary-foreground"
      @click="toggleSidebar"
    >
      <HamburgerMenuIcon class="h-4 w-4" />
    </Button>

    <!-- Logo -->
    <div class="flex items-center justify-center gap-2">
      <div class="w-3 h-3 bg-green-500 rounded-full animate-pulse"></div>
      <h1 class="text-xl font-semibold">IT Asset Management</h1>
    </div>

    <!-- Menu Items -->
    <div class="flex flex-col flex-1 gap-2 mt-10">
      <div class="flex flex-col flex-1">
        <template v-for="item in sidebarItem.top" :key="item.title">
          <Button
            v-if="!item.dropdown"
            :variant="isActive(item.url!) ? 'default' : 'ghost'"
            class="justify-start hover:bg-primary hover:text-primary-foreground"
            @click="handleNavigation(item.url!)"
          >
            <component :is="item.icon" class="h-4 w-4 mr-2" />
            {{ item.title }}
          </Button>

          <div v-else class="mb-2">
            <Button
              variant="ghost"
              class="justify-start w-full hover:bg-primary hover:text-primary-foreground"
              @click="toggleDropdown(item)"
            >
              <component :is="item.icon" class="h-4 w-4 mr-2" />
              {{ item.title }}
            </Button>

            <div v-show="item.isOpen" class="ml-6 flex flex-col gap-1">
              <Button
                v-for="child in item.children"
                :key="child.title"
                :variant="isActive(child.url) ? 'default' : 'ghost'"
                class="justify-start hover:bg-primary hover:text-primary-foreground"
                @click="handleNavigation(child.url)"
              >
                {{ child.title }}
              </Button>
            </div>
          </div>
        </template>
      </div>

      <!-- Bottom Menu -->
      <div class="flex flex-col">
        <Button
          v-for="item in sidebarItem.bottom"
          :key="item.title"
          :variant="isActive(item.url) ? 'default' : 'ghost'"
          class="justify-start hover:bg-primary hover:text-primary-foreground"
          @click="handleNavigation(item.url)"
        >
          <SvgIcon type="mdi" :path="item.icon" class="h-4 w-4 mr-2" />
          {{ item.title }}
        </Button>
      </div>
    </div>
  </aside>
</template>
