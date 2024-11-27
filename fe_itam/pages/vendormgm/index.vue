<script setup lang="ts">
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useRouter } from "vue-router";
import { useToast } from '@/components/ui/toast/use-toast'
import ActionBtnEdit from "~/components/atoms/ActionBtnEdit.vue";
import ActionBtnDelete from "~/components/atoms/ActionBtnDelete.vue";

const router = useRouter();
const config = useRuntimeConfig();
const { toast } = useToast()

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
    accessorKey: "nama_pic",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama PIC", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "email",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Email", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "nomor_kontak",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Kontak", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "lokasi_perusahaan",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Lokasi Perusahaan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "nomor_siup",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor SIUP", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "nomor_nib",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor NIB", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "nomor_npwp",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor NPWP", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    id: "actions",
    cell: ({ row }: { row: any }) => {
      return h("div", { class: "flex justify-end space-x-2" }, [
        h(
          ActionBtnEdit,
          {
            onClick: () => router.push(`/vendormgm/${row.original.id}/edit`),
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

/* fetch data from api */
const {
  data: result,
  status,
  refresh,
} = await useFetch(config.public.API_URL + '/vendor');

// Delete action
async function deleteData(id: number) {
  try {
    const { status } = await useFetch(config.public.API_URL + `/vendor/${id}`, {
      method: 'DELETE',
    });

    if (status.value == 'success') {
      console.log('duarrr');
      toast({
        title: 'Success',
        description: 'Data deleted successfully.',
      })
      refresh()
    } else {
      toast({
        title: 'Failed',
        description: `Error when deleting data`,
      })
    }
  } catch (error) {
    console.error('Error occured:', error);
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
