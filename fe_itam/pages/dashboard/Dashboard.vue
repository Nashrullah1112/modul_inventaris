<script setup lang="ts">
import { ref, onMounted } from "vue";
import {
  BellIcon,
  MonitorIcon,
  BoxIcon,
  KeyIcon,
  UsersIcon,
  ClockIcon,
} from "lucide-vue-next";

// Menggunakan state dari sidebar
const isOpen = useState("is-sidebar-open", () => false);

const stats = ref({
  elektronik: 0,
  aplikasi: 0,
  lisensi: 0,
  pegawai: 0,
});

const notifications = ref([
  {
    id: 1,
    title: "Lisensi Windows 10 Pro",
    expiry: "2024-03-20",
    status: "warning",
    description: "Perlu perpanjangan dalam 30 hari",
  },
  {
    id: 2,
    title: "Lisensi Adobe Creative Cloud",
    expiry: "2024-03-25",
    status: "danger",
    description: "Perlu perpanjangan dalam 7 hari",
  },
]);

onMounted(() => {
  // Di sini nanti fetch data dari API
  stats.value = {
    elektronik: 150,
    aplikasi: 45,
    lisensi: 30,
    pegawai: 200,
  };
});
</script>

<template>
  <div>
    <div class="p-8 bg-white shadow-lg rounded-lg">
      <div class="flex items-center justify-between mb-8">
        <h1 class="text-3xl font-bold text-gray-800">Dashboard</h1>
        <div class="flex items-center gap-4">
          <span class="text-sm text-gray-500">
            <ClockIcon class="inline h-4 w-4 mr-1" />
            {{ new Date().toLocaleDateString("id-ID") }}
          </span>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-4 mb-8">
        <div
          class="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-xl shadow-sm hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm font-medium text-blue-800"
              >Total Aset Elektronik</span
            >
            <MonitorIcon class="h-5 w-5 text-blue-600" />
          </div>
          <div class="text-3xl font-bold text-blue-900">
            {{ stats.elektronik }}
          </div>
        </div>

        <div
          class="bg-gradient-to-br from-purple-50 to-purple-100 p-6 rounded-xl shadow-sm hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm font-medium text-purple-800"
              >Total Aplikasi</span
            >
            <BoxIcon class="h-5 w-5 text-purple-600" />
          </div>
          <div class="text-3xl font-bold text-purple-900">
            {{ stats.aplikasi }}
          </div>
        </div>

        <div
          class="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-xl shadow-sm hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm font-medium text-green-800"
              >Total Lisensi Software</span
            >
            <KeyIcon class="h-5 w-5 text-green-600" />
          </div>
          <div class="text-3xl font-bold text-green-900">
            {{ stats.lisensi }}
          </div>
        </div>

        <div
          class="bg-gradient-to-br from-orange-50 to-orange-100 p-6 rounded-xl shadow-sm hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between mb-4">
            <span class="text-sm font-medium text-orange-800"
              >Total Pegawai</span
            >
            <UsersIcon class="h-5 w-5 text-orange-600" />
          </div>
          <div class="text-3xl font-bold text-orange-900">
            {{ stats.pegawai }}
          </div>
        </div>
      </div>

      <!-- Notifications -->
      <div
        class="bg-gradient-to-br from-gray-50 to-gray-100 rounded-xl shadow-sm"
      >
        <div class="p-6 border-b border-gray-200">
          <h2 class="text-xl font-bold flex items-center gap-3 text-gray-800">
            <BellIcon class="h-6 w-6 text-gray-600" />
            Notifikasi Lisensi
          </h2>
        </div>
        <div class="p-6">
          <div class="space-y-4">
            <div
              v-for="notif in notifications"
              :key="notif.id"
              class="flex items-center justify-between p-5 border rounded-xl bg-white hover:shadow-md transition-shadow"
            >
              <div>
                <h3 class="font-semibold text-gray-800">{{ notif.title }}</h3>
                <p class="text-sm text-gray-600 mt-1">
                  <span class="font-medium">Berakhir pada:</span>
                  {{ new Date(notif.expiry).toLocaleDateString("id-ID") }}
                </p>
                <p class="text-sm text-gray-500 mt-1">
                  {{ notif.description }}
                </p>
              </div>
              <span
                :class="{
                  'px-4 py-2 text-sm font-medium rounded-full': true,
                  'bg-yellow-100 text-yellow-800 border border-yellow-200':
                    notif.status === 'warning',
                  'bg-red-100 text-red-800 border border-red-200':
                    notif.status === 'danger',
                }"
              >
                {{ notif.status === "warning" ? "Segera Berakhir" : "Kritis" }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
