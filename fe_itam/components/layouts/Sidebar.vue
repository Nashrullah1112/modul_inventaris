<template>
  <div class="container">
    <!-- Navbar -->
    <nav
      class="fixed top-0 left-0 right-0 h-16 bg-white shadow-md z-30 transition-all duration-300 ease-in-out"
      :class="{ 'pl-64': isOpen }"
    >
      <div class="flex items-center h-full px-4">
        <button
          class="p-2 rounded-md hover:bg-gray-100 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-400"
          @click="toggleSidebar"
          aria-label="Toggle sidebar"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        </button>
      </div>
    </nav>

    <!-- Sidebar -->
    <aside
      class="fixed top-0 left-0 h-full w-64 bg-white shadow-md z-20 transition-transform duration-300 ease-in-out"
      :class="{ '-translate-x-full': !isOpen }"
    >
      <div class="h-16 bg-gray-100"></div>
      <div class="p-4 h-[calc(100%-4rem)] overflow-y-auto">
        <!-- Sidebar menu -->
        <div class="space-y-1">
          <template v-for="(item, index) in sidebarItems" :key="index">
            <button
              v-if="!item.isDropdown"
              class="w-full text-left p-2 rounded-md transition-all duration-200 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400"
              @click="navigateTo(item.url)"
            >
              {{ item.title }}
            </button>
            <div v-else class="space-y-1">
              <button
                class="w-full text-left p-2 rounded-md transition-all duration-200 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400 flex justify-between items-center"
                @click="toggleDropdown(item.title)"
              >
                <span>{{ item.title }}</span>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-4 w-4 transition-transform duration-200"
                  :class="{ 'rotate-180': isDropdownOpen === item.title }"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 9l-7 7-7-7"
                  />
                </svg>
              </button>
              <div
                v-if="isDropdownOpen === item.title"
                class="ml-4 space-y-1 transition-all duration-200 ease-in-out"
              >
                <button
                  v-for="dropdownItem in item.children"
                  :key="dropdownItem.url"
                  class="w-full text-left p-2 rounded-md text-sm transition-all duration-200 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400"
                  @click="navigateTo(dropdownItem.url)"
                >
                  {{ dropdownItem.title }}
                </button>
              </div>
            </div>
          </template>
        </div>

        <!-- Footer menu items -->
        <div class="mt-8 space-y-1">
          <button
            v-for="(item, index) in footerItems"
            :key="index"
            class="w-full text-left p-2 rounded-md transition-all duration-200 hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-400"
            @click="navigateTo(item.url)"
          >
            {{ item.title }}
          </button>
        </div>
      </div>
    </aside>

    <!-- Main content -->
    <main
      class="pt-16 transition-all duration-300 ease-in-out"
      :class="{ 'pl-64': isOpen }"
    >
      <!-- Your main content goes here -->
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const isOpen = ref(true);
const isDropdownOpen = ref<string | null>(null);
const router = useRouter();

const toggleSidebar = () => {
  isOpen.value = !isOpen.value;
};

const toggleDropdown = (title: string) => {
  isDropdownOpen.value = isDropdownOpen.value === title ? null : title;
};

const navigateTo = (url: string) => {
  if (url) {
    router.push(url);
  }
};

const sidebarItems = [
  { title: "Dashboard", url: "/" },
  { title: "Inspeksi", url: "/inspeksi" },
  {
    title: "Registrasi Inventaris",
    isDropdown: true,
    children: [
      { title: "New Asset", url: "/registrasi-inventaris/new-asset" },
      { title: "Upload File", url: "/registrasi-inventaris/upload" },
      { title: "Review Asset", url: "/registrasi-inventaris/review" },
      { title: "Archive Asset", url: "/registrasi-inventaris/archive" },
    ],
  },
  { title: "Storage", url: "/storage" },
  { title: "Tambah User", url: "/tambah-user" },
  { title: "Assignment", url: "/assignment" },
  { title: "Asset", url: "/eg/data_hardware" },
  { title: "DataTable Example", url: "/eg/datatable" },
];

const footerItems = [{ title: "Profile", url: "/profile" }];
</script>
