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
    accessorKey: "assetName",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Asset Name", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) =>
      h("div", { class: "capitalize" }, row.getValue("assetName")),
  },
  {
    accessorKey: "assetCategory",
    header: () => h("div", {}, "Asset Category"),
    cell: ({ row }) => h("div", {}, row.getValue("assetCategory")),
  },
  {
    accessorKey: "brand",
    header: () => h("div", {}, "Brand"),
    cell: ({ row }) => h("div", {}, row.getValue("brand")),
  },
  {
    accessorKey: "description",
    header: () => h("div", {}, "Description"),
    cell: ({ row }) => h("div", {}, row.getValue("description")),
  },
  {
    accessorKey: "year",
    header: () => h("div", {}, "Year"),
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
      assetName: "Laptop X1",
      assetCategory: "Electronics",
      brand: "Brand A",
      description: "High-performance laptop for professionals.",
      year: 2021,
      status: "active",
      price: 12000000, // 12 juta
    },
    {
      assetName: "Smartphone Y2",
      assetCategory: "Electronics",
      brand: "Brand B",
      description: "Latest smartphone with advanced features.",
      year: 2023,
      status: "active",
      price: 8000000, // 8 juta
    },
    {
      assetName: "Office Chair Z3",
      assetCategory: "Furniture",
      brand: "Brand C",
      description: "Ergonomic office chair for comfort.",
      year: 2020,
      status: "active",
      price: 2500000, // 2.5 juta
    },
    {
      assetName: "Projector Q4",
      assetCategory: "Electronics",
      brand: "Brand D",
      description: "Portable projector for presentations.",
      year: 2022,
      status: "non active",
      price: 6000000, // 6 juta
    },
    {
      assetName: "Desktop PC R5",
      assetCategory: "Electronics",
      brand: "Brand E",
      description: "Powerful desktop for gaming and work.",
      year: 2020,
      status: "active",
      price: 15000000, // 15 juta
    },
    {
      assetName: "Printer S6",
      assetCategory: "Electronics",
      brand: "Brand F",
      description: "All-in-one printer with scanning capabilities.",
      year: 2019,
      status: "non active",
      price: 3000000, // 3 juta
    },
    {
      assetName: "Conference Table T7",
      assetCategory: "Furniture",
      brand: "Brand G",
      description: "Spacious table for meetings.",
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
