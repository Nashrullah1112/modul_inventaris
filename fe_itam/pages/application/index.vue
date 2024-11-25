<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useRouter } from "vue-router";

const router = useRouter();
const config = useRuntimeConfig();

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
const errorMessage = ref<string | null>(null);

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
    accessorKey: "nama_aplikasi",
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
    accessorKey: "tipe_aplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "link_aplikasi",
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
    accessorKey: "lokasi_server_penyimpanan",
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

console.log(config.public);

/* fetch data from api */
const {
  data: result,
  status,
  refresh,
} = await useFetch(config.public.API_URL + "/asset-aplikasi");

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
  <div class="bg-white rounded-lg shadow-lg">
    <div class="px-6 py-2 border-b border-gray-200 flex justify-between">
      <h1 class="text-2xl font-bold text-gray-800">Data Aplikasi</h1>
      <Button @click="refresh()" variant="secondary">Refresh</Button>
    </div>
    <div class="px-6 py-2">
      <DataTable :columns="columns" :data="result?.data || []" :dataStatus="status" />
    </div>
  </div>
</template>
