<script setup lang="ts">
import DataTable from "@/components/DataTable.vue";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type { ColumnDef } from "@tanstack/vue-table";
import { ArrowUpDown } from "lucide-vue-next";
import { h, onMounted, ref, computed } from "vue";

const config = useRuntimeConfig();
const employees = ref<Employee[]>([]);
const tableData = ref([]);
const { showLoading, hideLoading } = useLoading();
const errorMessage = ref<string | null>(null);
const isLoading = ref(false);

interface TableRow {
  getIsSelected: () => boolean;
  toggleSelected: (value: boolean) => void;
  getValue: (key: string) => any;
  original: any;
}

interface TableHeader {
  getIsAllPageRowsSelected: () => boolean;
  toggleAllPageRowsSelected: (value: boolean) => void;
}

interface Column {
  getIsSorted: () => string;
  toggleSorting: (asc: boolean) => void;
}

interface Position {
  id: number;
  nama: string;
}

interface Division {
  id: number;
  nama: string;
}

// Ini seperti model
interface Employee {
  id: number;
  nip: number;
  email: string;
  jabatan_id: number;
  jabatan: Position;
  divisi_id: number;
  divisi: Division;
  tanggal_bergabung: string;
  nama: string;
}

const isOpen = useState("is-sidebar-open", () => false);

const columns: ColumnDef<Employee>[] = [
  {
    id: "select",
    header: ({ table }: { table: TableHeader }) =>
      h(Checkbox, {
        checked: table.getIsAllPageRowsSelected(),
        "onUpdate:checked": (value) => table.toggleAllPageRowsSelected(!!value),
        ariaLabel: "Select all",
      }),
    cell: ({ row }: { row: TableRow }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        "onUpdate:checked": (value) => row.toggleSelected(!!value),
        ariaLabel: "Select row",
      }),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "employeeId",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["ID Pegawai", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "name",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "position",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Jabatan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "division",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Divisi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "joinDate", // The key for the date field in your data
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => [
          "Tanggal Bergabung",
          h(ArrowUpDown, { class: "ml-2 h-4 w-4" }),
        ]
      );
    },
    cell: ({ row }: { row: any }) => {
      // Get the raw date from the row and format it using the `parseDate` function
      const formattedDate = parseDate(row.getValue("joinDate"));
      return formattedDate;
    },
  },
];

// Compute the data status based on loading state and error
const dataStatus = computed(() => {
  if (isLoading.value) return 'loading';
  if (errorMessage.value) return 'error';
  return 'success';
});

async function fetchEmployees() {
  try {
    isLoading.value = true;
    showLoading();
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Employee[];
      error: any;
    }>(config.public.API_URL + "/user", {
      method: "GET",
    });

    if (response.error) {
      isLoading.value = false;
      hideLoading();
      throw new Error(response.error);
    }

    employees.value = response.data;

    // Transform the data for the table
    tableData.value = employees.value.map(employee => ({
      employeeId: employee.nip.toString(),
      name: employee.nama || 'Unnamed',
      position: employee.jabatan.nama, // Use actual position name
      division: employee.divisi.nama, // Use actual division name
      joinDate: new Date(employee.tanggal_bergabung).toISOString().split('T')[0]
    }));

    isLoading.value = false;
    hideLoading();
  } catch (error) {
    isLoading.value = false;
    hideLoading();
    console.error("Error fetching employees:", error);
    errorMessage.value = "Failed to load employee data. Please try again.";
  }
}

onMounted(() => {
  fetchEmployees();
});
</script>

<template>
  <div class="container mx-auto py-10">
    <div class="flex flex-col space-y-8">
      <div>
        <h2 class="text-2xl font-bold tracking-tight">Data Karyawan</h2>
        <p class="text-muted-foreground">
          Daftar seluruh karyawan yang terdaftar dalam sistem
        </p>
      </div>

      <DataTable :columns="columns" :data="tableData" :dataStatus="dataStatus" />

      <!-- Optional: Add error handling -->
      <div v-if="errorMessage" class="text-red-500">
        {{ errorMessage }}
      </div>
    </div>
  </div>
</template>