<script setup lang="ts">
import { h } from "vue";
import DropdownAction from "@/components/DataTableDropDown.vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown, ChevronDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";

const columns = [
  {
    id: "select",
    header: ({ table }) =>
      h(Checkbox, {
        checked: table.getIsAllPageRowsSelected(),
        "onUpdate:checked": (value) => table.toggleAllPageRowsSelected(!!value),
        ariaLabel: "Select all",
      }),
    cell: ({ row }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        "onUpdate:checked": (value) => row.toggleSelected(!!value),
        ariaLabel: "Select row",
      }),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "NomorAsset",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Asset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) =>
      h("div", { class: "capitalize" }, row.getValue("Nomor Asset")),
  },
  {
    accessorKey: "SerialNumber",
    header: () => h("div", {}, "Serial Number"),
    cell: ({ row }) => h("div", {}, row.getValue("Serial Number")),
  },
  {
    accessorKey: "Pengguna",
    header: () => h("div", {}, "Pengguna"),
    cell: ({ row }) => h("div", {}, row.getValue("Pengguna")),
  },
  {
    accessorKey: "CetakBarcode",
    header: () => h("div", {}, "Cetak Barcode"),
    cell: ({ row }) => h("div", {}, row.getValue("Cetak Barcode")),
  },
  {
    accessorKey: "year",
    header: () => h("div", {}, "year"),
    cell: ({ row }) => h("div", {}, row.getValue("year")),
  },
  {
    accessorKey: "status",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Status", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) =>
      h("div", { class: "capitalize" }, row.getValue("status")),
  },
  {
    accessorKey: "price",
    header: () => h("div", { class: "text-right" }, "Price"),
    cell: ({ row }) => {
      const price = row.getValue("price");
      const formatted = new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
      }).format(price);

      return h("div", { class: "text-right font-medium" }, formatted);
    },
  },

  {
    id: "actions",
    enableHiding: false,
    cell: ({ row }) => {
      const asset = row.original;

      return h(
        "div",
        { class: "relative" },
        h(DropdownAction, {
          asset,
          onExpand: row.toggleExpanded,
        })
      );
    },
  },
];

const data = ref<Payment[]>([]);

async function getData(): Promise<Payment[]> {
  // Fetch data from your API here.
  return [
    {
      NomorAsset: "Laptop X1",
      SerialNumber: "Electronics",
      Pengguna: "Pengguna A",
      CetakBarcode: "High-performance laptop for professionals.",
      year: 2021,
      status: "active",
      price: 12000000, // 12 juta
    },
    {
      NomorAsset: "Smartphone Y2",
      SerialNumber: "Electronics",
      Pengguna: "Pengguna B",
      CetakBarcode: "Latest smartphone with advanced features.",
      year: 2023,
      status: "active",
      price: 8000000, // 8 juta
    },
    {
      NomorAsset: "Office Chair Z3",
      SerialNumber: "Furniture",
      Pengguna: "Pengguna C",
      CetakBarcode: "Ergonomic office chair for comfort.",
      year: 2020,
      status: "active",
      price: 2500000, // 2.5 juta
    },
    {
      NomorAsset: "Projector Q4",
      SerialNumber: "Electronics",
      Pengguna: "Pengguna D",
      CetakBarcode: "Portable projector for presentations.",
      year: 2022,
      status: "active",
      price: 6000000, // 6 juta
    },
    {
      NomorAsset: "Desktop PC R5",
      SerialNumber: "Electronics",
      Pengguna: "Pengguna E",
      CetakBarcode: "Powerful desktop for gaming and work.",
      year: 2020,
      status: "active",
      price: 15000000, // 15 juta
    },
    {
      NomorAsset: "Printer S6",
      SerialNumber: "Electronics",
      Pengguna: "Pengguna F",
      CetakBarcode: "All-in-one printer with scanning capabilities.",
      year: 2019,
      status: "active",
      price: 3000000, // 3 juta
    },
    {
      NomorAsset: "Conference Table T7",
      SerialNumber: "Furniture",
      Pengguna: "Pengguna G",
      CetakBarcode: "Spacious table for meetings.",
      year: 2021,
      status: "active",
      price: 8000000, // 8 juta
    },
  ];
}

onMounted(async () => {
  data.value = await getData();
});
</script>

<template>
  <div>
    <DataTable :columns="columns" :data="data" />
  </div>
</template>
