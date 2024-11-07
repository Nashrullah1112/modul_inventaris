<script setup lang="ts">
import SvgIcon from "@jamescoyle/vue-icon";
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
      title: "Dashboard",
      url: "/",
      icon: HomeIcon,
    },
    {
      title: "Registrasi Inventaris",
      icon: StackIcon,
      dropdown: true,
      children: [
        {
          title: "Perangkat Elektronik",
          url: "/electronic/form",
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
          url: "/software/form",
        },
      ],
    },
    {
      title: "Storage",
      icon: ArchiveIcon,
      dropdown: true,
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
          title: "Approval",
          url: "/approval",
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
      title: "Pegawai",
      icon: PersonIcon,
      dropdown: true,
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
      url: "#",
      icon: mdiAccount,
    },
    {
      title: "Logout",
      url: "#",
      icon: mdiLogout,
    },
  ],
};

const isActive = (url: string) => {
  return route.path === url;
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
      <HamburgerMenuIcon class="size-4" />
    </Button>

    <div
      class="w-full text-xl text-center flex items-center justify-center gap-2"
    >
      <div class="w-3 h-3 bg-green-500 rounded-full"></div>
      <h1>IT Asset Management</h1>
    </div>

    <div class="flex flex-col flex-1 gap-2 mt-10">
      <div class="flex flex-col flex-1">
        <template v-for="item in sidebarItem.top">
          <Button
            v-if="!item.dropdown"
            :variant="isActive(item.url) ? 'default' : 'ghost'"
            class="justify-start"
            @click="navigateTo(item.url) || '#'"
          >
            <component :is="item.icon" class="size-4 mr-2" />
            {{ item.title }}
          </Button>

          <div v-else class="mb-2">
            <Button variant="ghost" class="justify-start w-full">
              <component :is="item.icon" class="size-4 mr-2" />
              {{ item.title }}
            </Button>

            <div class="ml-6 flex flex-col gap-1">
              <Button
                v-for="child in item.children"
                :variant="isActive(child.url) ? 'default' : 'ghost'"
                class="justify-start"
                @click="navigateTo(child.url)"
              >
                {{ child.title }}
              </Button>
            </div>
          </div>
        </template>
      </div>

      <div class="flex flex-col">
        <Button
          :variant="isActive(item.url) ? 'default' : 'ghost'"
          class="justify-start"
          v-for="item in sidebarItem.bottom"
          @click="navigateTo(item.url) || '#'"
        >
          <SvgIcon type="mdi" :path="item.icon" class="size-4 mr-2" />
          {{ item.title }}
        </Button>
      </div>
    </div>
  </div>
</template>
