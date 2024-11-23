<script setup lang="ts">
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
      h("div", { class: "capitalize" }, row.getValue("serialNumber")),
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
        () => ["Nomor Asset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("assetNumber")),
  },
  {
    accessorKey: "deviceName",
    header: () => h("div", {}, "Data Perangkat"),
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("deviceName")),
  },
  {
    accessorKey: "division",
    header: () => h("div", {}, "Divisi"),
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("division")),
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

const electronics = ref([
  {
    id: 1,
    serialNumber: "SN-2024-001",
    assetNumber: "AST-001",
    deviceName: "Laptop Dell XPS 13",
    division: "Divisi IT",
  },
  {
    id: 2,
    serialNumber: "SN-2024-002",
    assetNumber: "AST-002",
    deviceName: 'Monitor LG 27"',
    division: "Divisi Marketing",
  },
]);

const isOpen = ref(false);
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Aset Elektronik</h1>
      </div>

      <div class="p-6">
        <DataTable :columns="columns" :data="electronics" />
      </div>
    </div>
  </div>
</template>
