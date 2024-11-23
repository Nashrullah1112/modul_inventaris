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

const isOpen = useState("is-sidebar-open", () => false);

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
    accessorKey: "namaAplikasi",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "urlAplikasi",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["URL Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    accessorKey: "server",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Server Penyimpanan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
  },
  {
    id: "actions",
    cell: ({ row }: { row: TableRow }) => {
      return h("div", { class: "flex justify-end space-x-2" }, [
        h(
          Button,
          {
            variant: "outline",
            class: "bg-blue-500 hover:bg-blue-600 text-white w-[100px]",
            onClick: () => console.log("Update", row.original.id),
          },
          () => "Update"
        ),
        h(
          Button,
          {
            variant: "destructive",
            class: "w-[100px]",
            onClick: () => console.log("Delete", row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

const aplikasiList = ref([
  {
    id: 1,
    namaAplikasi: "SIMRS",
    urlAplikasi: "https://simrs.rsud.com",
    server: "Server 1 - 192.168.1.10",
  },
  {
    id: 2,
    namaAplikasi: "E-Office",
    urlAplikasi: "https://eoffice.rsud.com",
    server: "Server 2 - 192.168.1.20",
  },
  {
    id: 3,
    namaAplikasi: "Website RSUD",
    urlAplikasi: "https://rsud.com",
    server: "Server 3 - 192.168.1.30",
  },
]);
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Aplikasi</h1>
      </div>

      <div class="p-6">
        <DataTable :columns="columns" :data="aplikasiList" />
      </div>
    </div>
  </div>
</template>
