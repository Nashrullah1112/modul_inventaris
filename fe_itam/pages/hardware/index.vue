<script setup lang="ts">
<<<<<<< HEAD
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";

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

interface Hardware {
  id: number;
  assetNumber: string;
  serialNumber: string;
  assetType: string;
  division: string;
  currentDepreciation: string;
  assetDetail: string;
}

const columns = [
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
    accessorKey: "assetNumber",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Aset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", { class: "capitalize" }, row.getValue("assetNumber")),
  },
  {
    accessorKey: "serialNumber",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Serial Number", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("serialNumber")),
  },
  {
    accessorKey: "assetType",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe Aset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("assetType")),
  },
  {
    accessorKey: "division",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Divisi Pengguna", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("division")),
  },
  {
    accessorKey: "currentDepreciation",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nilai Depresiasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("currentDepreciation")),
  },
  {
    accessorKey: "assetDetail",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Detail Aset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("assetDetail")),
  },
  {
    id: "actions",
    header: () => h("div", {}, "Aksi"),
    cell: ({ row }: { row: TableRow }) => {
      return h("div", { class: "flex gap-2" }, [
        h(
          Button,
          {
            variant: "default",
            class: "bg-blue-500 hover:bg-blue-600",
            onClick: () => console.log("Update", row.original.id),
          },
          () => "Update"
        ),
        h(
          Button,
          {
            variant: "default",
            class: "bg-red-500 hover:bg-red-600",
            onClick: () => console.log("Delete", row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

const hardwareAssets = ref<Hardware[]>([
  {
    id: 1,
    assetNumber: "HW-001",
    serialNumber: "SN123456789",
    assetType: "SSD",
    division: "IT",
    currentDepreciation: "Rp 15.000.000",
    assetDetail: "Samsung 970 EVO Plus 1TB",
  },
  {
    id: 2,
    assetNumber: "HW-002",
    serialNumber: "SN987654321",
    assetType: "RAM",
    division: "Finance",
    currentDepreciation: "Rp 8.000.000",
    assetDetail: "Corsair Vengeance 32GB DDR4",
  },
]);

const isOpen = ref(false);
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Hardware</h1>
      </div>

      <div class="p-6">
        <DataTable :columns="columns" :data="hardwareAssets" />
      </div>
    </div>
  </div>
</template>
=======
import { Button } from '@/components/ui/button'

const route = useRoute()
</script>

<template>
  <div class="flex mb-2">
    <Button @click="navigateTo(route.path + '/add')">Tambah</Button>
  </div>

  <p>Tabel Data Hardware</p>
</template>
>>>>>>> 988390618515c82a10424c730a828e66c000c19b
