<script setup lang="ts">
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useToast } from "@/components/ui/toast/use-toast";
import ActionBtnEdit from "~/components/atoms/ActionBtnEdit.vue";
import ActionBtnDelete from "~/components/atoms/ActionBtnDelete.vue";

const router = useRouter();
const config = useRuntimeConfig();
const { toast } = useToast();

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
    id: "nama_aplikasi",
    accessorFn: (row) => row.asset?.merk || "-", // Ubah bagian ini
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
    cell: ({ row }: { row: TableRow }) => {
      const asset = row.original.asset;
      return h("div", {}, [
        h("div", { class: "font-medium" }, asset?.merk || "-"),
      ]);
    },
  },
  {
    accessorKey: "SN_perangkat_terpasang",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => [
          "SN Perangkat Terpasang",
          h(ArrowUpDown, { class: "ml-2 h-4 w-4" }),
        ]
      ),
  },
  {
    accessorKey: "kategori_lisensi",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Kategori", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "versi_lisensi",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Versi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "tipe_kepemilikan_aset",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe Kepemilikan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "waktu_aktivasi",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tanggal Aktivasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "tanggal_expired",
    header: ({ column }: { column: Column }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tanggal Expired", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    id: "actions",
    header: () => h("div", { class: "text-right" }, "Aksi"),
    cell: ({ row }: { row: TableRow }) => {
      return h("div", { class: "flex items-center justify-end gap-2" }, [
        h(
          ActionBtnEdit,
          {
            onClick: () =>
              router.push(`/software-license/${row.original.id}/edit`),
          },
          () => "Update"
        ),
        h(
          ActionBtnDelete,
          {
            onClick: () => deleteData(row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

const {
  data: result,
  status,
  refresh,
} = await useFetch(config.public.API_URL + "/asset-lisensi");
</script>

<template>
  <div class="bg-background rounded-lg border shadow-sm">
    <div class="flex items-center justify-between p-6 border-b">
      <h1 class="text-lg font-semibold">Data Lisensi Software</h1>
      <Button @click="refresh()" variant="outline" size="sm"> Refresh </Button>
    </div>
    <div class="p-6">
      <DataTable
        :columns="columns"
        :data="result?.data || []"
        :dataStatus="status"
      />
    </div>
  </div>
</template>
