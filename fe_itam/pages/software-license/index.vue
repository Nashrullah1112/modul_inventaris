<script setup lang="ts">
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useToast } from "@/components/ui/toast/use-toast";
import ActionBtnEdit from "~/components/atoms/ActionBtnEdit.vue";
import ActionBtnDelete from "~/components/atoms/ActionBtnDelete.vue";
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
import type { ColumnDef } from "@tanstack/vue-table";

const router = useRouter();
const config = useRuntimeConfig();
const { toast } = useToast();
const { showLoading, hideLoading } = useLoading();
const licenses = ref<License[]>([]);
const isDeleteDialogOpen = ref(false);
const licenseToDelete = ref<number | null>(null);

interface License {
  id: number;
  waktu_pembelian: string;
  SN_perangkat_terpasang: string;
  waktu_aktivasi: string;
  tanggal_expired: string;
  tipe_kepemilikan_aset: string;
  kategori_lisensi: string;
  versi_lisensi: string;
  maksimal_user_aplikasi: number;
  maksimal_perangkat_lisensi: number;
  tipe_lisensi: string;
  asset_id: number;
}

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

const columns: ColumnDef<License>[] = [
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
    header: ({ column }) =>
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
            onClick: () => openDeleteConfirmation(row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

const openDeleteConfirmation = (id: number) => {
  licenseToDelete.value = id;
  isDeleteDialogOpen.value = true;
}

const confirmDeletion = () => {
  if (licenseToDelete.value !== null) {
    deleteData(licenseToDelete.value);
  }
}

const fetchLicenseData = async () => {
  try {
    showLoading();
    const response = await $fetch<{ data: License[] }>(config.public.API_URL + '/asset-lisensi', {
      method: 'GET'
    });

    licenses.value = response.data;
    hideLoading();
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


async function deleteData(id: number) {
  try {
    showLoading
    const { status } = await useFetch(config.public.API_URL + `/asset-lisensi/${id}`, {
      method: 'DELETE',
    });

    if (status.value == 'success') {
      hideLoading();
      toast({
        title: 'Data Lisensi Berhasil Dihapus',
        description: 'Data dengan ID lisensi' + id + ' berhasil dihapus.',
      })
      licenses.value = licenses.value.filter(item => item.id !== id);
      await fetchLicenseData();
      isDeleteDialogOpen.value = false;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Menghapus Data Lisensi',
        description: `Data lisensi dengan ID ${id} gagal dihapus.`,
        variant: 'destructive',
      })
    }
  } catch (error) {
    hideLoading();
    console.error('Error occured:', error);
  }
}

onMounted(async () => {
  await fetchLicenseData();
});
</script>

<template>
  <div class="bg-background rounded-lg border shadow-sm">
    <div class="flex items-center justify-between p-6 border-b">
      <h1 class="text-lg font-semibold">Data Lisensi Software</h1>
      <Button @click="fetchLicenseData()" variant="outline" size="sm"> Refresh </Button>
    </div>
    <div class="p-6">
      <DataTable :columns="columns" :data="licenses || []" :dataStatus="status" />
    </div>
    <AlertDialog v-model:open="isDeleteDialogOpen">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Apakah Anda yakin ingin menghapus?</AlertDialogTitle>
          <AlertDialogDescription>
            Data yang Anda pilih akan dipindahkan ke disposal dan tidak akan dihapus secara permanen.
            Tindakan ini tidak dapat dibatalkan.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>Batal</AlertDialogCancel>
          <AlertDialogAction @click="confirmDeletion" class="bg-red-600 hover:bg-red-700 text-white">
            Hapus
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  </div>
</template>
