<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useRouter } from "vue-router";

interface Application {
  id: number;
  nama_aplikasi: string;
  tanggal_pembuatan: Date;
  tanggal_terima: Date;
  lokasi_server_penyimpanan: string;
  tipe_aplikasi: string;
  link_aplikasi: string;
  sertifikasi_aplikasi: string;
  tanggal_aktif: Date;
  tanggal_kadaluarsa: Date;
  asset_id: number;
  vendor_id: number;
}

// State for application data
const applications = ref<Application[]>([]);
const isLoading = ref(false);
const errorMessage = ref<string | null>(null);
const router = useRouter();

async function fetchApplications() {
  try {
    isLoading.value = true;
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Application[];
      error: any;
    }>("http://localhost:5000/api/asset-aplikasi", {
      method: "GET",
    });

    if (response.error) {
      throw new Error(response.error);
    }

    applications.value = response.data;
  } catch (error) {
    console.error("Error fetching applications:", error);
    errorMessage.value = "Failed to load application data. Please try again.";
  } finally {
    isLoading.value = false;
  }
}

onMounted(fetchApplications);

// Transform API data for table usage
const transformedApplications = computed(() =>
  applications.value.map((app) => ({
    id: app.id,
    namaAplikasi: app.nama_aplikasi,
    urlAplikasi: app.link_aplikasi,
    server: app.lokasi_server_penyimpanan,
  }))
);

// Define table columns
const columns = [
  {
    id: "select",
    header: ({ table }: { table: any }) =>
      h(Checkbox, {
        checked: table.getIsAllPageRowsSelected(),
        "onUpdate:checked": (value) => table.toggleAllPageRowsSelected(!!value),
        ariaLabel: "Select all",
      }),
    cell: ({ row }: { row: any }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        "onUpdate:checked": (value) => row.toggleSelected(!!value),
        ariaLabel: "Select row",
      }),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "namaAplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "urlAplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["URL Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "server",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Server Penyimpanan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    id: "actions",
    cell: ({ row }: { row: any }) => {
      return h("div", { class: "flex justify-end space-x-2" }, [
        h(
          Button,
          {
            variant: "outline",
            class: "bg-blue-500 hover:bg-blue-600 text-white w-[100px]",
            onClick: () => router.push(`/aplikasi/edit/${row.original.id}`),
          },
          () => "Update"
        ),
        h(
          Button,
          {
            variant: "destructive",
            class: "w-[100px]",
            onClick: () => deleteApplication(row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

// Delete action
async function deleteApplication(id: number) {
  try {
    await $fetch(`http://localhost:5000/api/asset-aplikasi/${id}`, {
      method: "DELETE",
    });
    applications.value = applications.value.filter((app) => app.id !== id);
    console.log("Application deleted:", id);
  } catch (error) {
    console.error("Error deleting application:", error);
    errorMessage.value = "Failed to delete the application. Please try again.";
  }
}
</script>

<template>
  <div class="p-8">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Aplikasi</h1>
      </div>
      <div class="p-6">
        <!-- Error Message -->
        <div v-if="errorMessage" class="text-red-500 mb-4">
          {{ errorMessage }}
        </div>

        <!-- Loading Spinner -->
        <div v-if="isLoading" class="flex justify-center items-center h-40">
          <span>Loading...</span>
        </div>

        <!-- Data Table -->
        <DataTable v-else :columns="columns" :data="transformedApplications" />
      </div>
    </div>
  </div>
</template>
