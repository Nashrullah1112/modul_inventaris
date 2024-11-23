<script setup lang="ts">
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";

const config = useRuntimeConfig();

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
    accessorKey: "asset_number",
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
      h("div", { class: "capitalize" }, row.getValue("asset_number")),
  },
  {
    accessorKey: "serial_number",
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
      h("div", {}, row.getValue("serial_number")),
  },
  {
    accessorKey: "asset_type",
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
      h("div", {}, row.getValue("asset_type")),
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
    accessorKey: "current_depreciation",
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
      h("div", {}, row.getValue("current_depreciation")),
  },
  {
    accessorKey: "asset_detail",
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
      h("div", {}, row.getValue("asset_detail")),
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

/* fetch data from api */
const { data: result, status, refresh } = await useFetch(config.public.API_URL + '/t_asset_hardware', {
  headers: { 
    apiKey: config.public.API_KEY,
  },
})
</script>

<template>
  <div>
    <div class="flex justify-end gap-2">
      <Button @click="navigateTo('/hardware/add')">Tambah</Button>
      <Button @click="refresh()" variant="secondary">Refresh</Button>
    </div>
    <div class="bg-white rounded-lg shadow-lg mt-2">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Hardware</h1>
      </div>
      
      <div class="p-6">
        <DataTable :columns="columns" :data="result ?? []" :dataStatus="status" />
      </div>
    </div>
  </div>
</template>
