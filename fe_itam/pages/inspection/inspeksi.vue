<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type {
  ColumnDef
} from "@tanstack/vue-table";
import { ArrowUpDown } from "lucide-vue-next";
import { ref } from "vue";
import ActionBtnApprove from "~/components/atoms/ActionBtnApprove.vue";
import ActionBtnReject from "~/components/atoms/ActionBtnReject.vue";
import { useToast } from '@/components/ui/toast/use-toast'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'


const { showLoading, hideLoading } = useLoading();
const config = useRuntimeConfig();
const router = useRouter();
const { toast } = useToast()
const isApproveDialogOpen = ref(false);
const assetToApprove = ref<number | null>(null);

const assetList = ref<Data[]>([]);

interface Data {
  id: number;
  vendor_id: number;
  serial_number: string;
  merk: string;
  model: string;
  nomor_nota: string;
}

const fetchApprovalData = async () => {
  try {
    showLoading();
    const response = await $fetch<{ data: Data[] }>(config.public.API_URL + '/approval', {
      method: 'GET'
    });

    hideLoading();
    assetList.value = response.data;
  } catch (error) {
    hideLoading();
    console.error('Error occurred:', error);
    toast({
      title: 'Error',
      description: 'Terjadi kesalahan saat mengambil data.',
      variant: 'destructive',
    });
  } finally {
    hideLoading();
  }
}
onMounted(() => {
  fetchApprovalData();
});

const onAssetApproved = async (id: number) => {
  try {
    showLoading();
    const response = await $fetch<{ data: Data[] }>(config.public.API_URL + '/approval', {
      method: 'PATCH',
      body: {
        id: id,
      }
    });

    await fetchApprovalData();
    hideLoading();
    toast({
      title: 'Success',
      description: 'Data berhasil diapprove.',
    });
  } catch (error) {
    hideLoading();
    console.error('Error occurred:', error);
    toast({
      title: 'Error',
      description: 'Terjadi kesalahan saat mengapprove data.',
      variant: 'destructive',
    });
  } finally {
    hideLoading();
  }
}

const onAssetRejected = async (id: number) => {
  try {
    toast({
      title: 'Akan Datang!',
      description: 'Fitur ini akan segera hadir.',
    });
  } catch (error) {
    hideLoading();
    console.error('Error occurred:', error);
    toast({
      title: 'Error',
      description: 'Terjadi kesalahan saat mereject data.',
      variant: 'destructive',
    });
  } finally {
    hideLoading();
  }
}

const columns: ColumnDef<Data>[] = [
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
  },
  {
    accessorKey: "serial_number",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Seri", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) => h("div", {}, row.getValue("serial_number")),
  },

  {
    accessorKey: "merk",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Merek", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) => h("div", {}, row.getValue("merk")),
  },
  {
    accessorKey: "model",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Model", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) => h("div", {}, row.getValue("model")),
  },
  {
    accessorKey: "nomor_nota",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Nota", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) => h("div", {}, row.getValue("nomor_nota")),
  },
  {
    id: "detail",
    header: () => h("div", {}, ""),
    cell: ({ row }) => {
      return h("div", { class: "flex gap-2" }, [
        h(
          Button,
          {
            variant: "secondary",
            size: "sm",
            onClick: () => {
              console.log("Detail clicked for row:", row.original);
            },
          },
          () => "Rincian"
        ),
      ]);
    },
  },

  {
    id: "actions",
    header: () => h("div", {}, "Aksi"),
    cell: ({ row }) => {
      return h("div", { class: "flex gap-2" }, [
        h(
          ActionBtnApprove,
          {
            onClick: () => openApprovalConfirmation(row.original.id),
          },
          () => "Approve"
        ),
        h(
          ActionBtnReject,
          {
            onClick: () => onAssetRejected(row.original.id),
          },
          () => "Reject"
        ),
      ]);
    },
  }
];

const openApprovalConfirmation = (id: number) => {
  assetToApprove.value = id;
  isApproveDialogOpen.value = true;
}

const confirmApproval = () => {
  if (assetToApprove.value) {
    onAssetApproved(assetToApprove.value);
    isApproveDialogOpen.value = false;
  }
}
</script>

<template>
  <div class="p-8">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Inspeksi Aset</h1>
      </div>

      <div class="p-">
        <div class="px-6">
          <DataTable :columns="columns" :data="assetList || []" :dataStatus="status" />
        </div>

        <TableCell>
          <Button variant="secondary" size="sm">Rincian</Button>
        </TableCell>
      </div>
    </div>
    <AlertDialog v-model:open="isApproveDialogOpen">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Apakah Anda yakin ingin menyetujui?</AlertDialogTitle>
          <AlertDialogDescription>
            Data yang Anda pilih akan dipindahkan ke data aset.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Batal</AlertDialogCancel>
          <AlertDialogAction @click="confirmApproval" class="bg-green-600 hover:bg-green-700 text-white">
            Approve
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  </div>
</template>
