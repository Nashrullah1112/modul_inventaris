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

const router = useRouter();
const config = useRuntimeConfig();
const { toast } = useToast()
const { showLoading, hideLoading } = useLoading();
const applications = ref<Application[]>([]);
const isDeleteDialogOpen = ref(false);
const applicationToDelete = ref<number | null>(null);


interface Application {
  id: number;
  nama_aplikasi: string;
  tanggal_pembuatan: string;
  tanggal_terima: string;
  lokasi_server_penyimpanan: string;
  tipe_aplikasi: string;
  link_aplikasi: string;
  sertifikasi_aplikasi: string;
  tanggal_aktif: string;
  tanggal_kadaluarsa: string;
  asset_id: number;
}

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
    accessorKey: "nama_aplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nama Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "tipe_aplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "link_aplikasi",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["URL Aplikasi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "lokasi_server_penyimpanan",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Server Penyimpanan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    id: "actions",
    cell: ({ row }: { row: any }) => {
      return h("div", { class: "flex justify-end space-x-2" }, [
        h(
          ActionBtnEdit,
          {
            onClick: () => router.push(`/application/${row.original.id}/edit`),
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
  applicationToDelete.value = id;
  isDeleteDialogOpen.value = true;
}

const confirmDeletion = () => {
  if (applicationToDelete.value !== null) {
    deleteData(applicationToDelete.value);
  }
}

const fetchApplicationData = async () => {
  try {
    showLoading();
    const response = await $fetch<{ data: Application[] }>(config.public.API_URL + '/asset-aplikasi', {
      method: 'GET'
    });

    if (response && response.data) {
      hideLoading();
      applications.value = response.data;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Mengambil Data Aplikasi',
        description: 'Data aplikasi gagal diambil dari server.',
        variant: 'destructive',
      });
    }
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

onMounted(async () => {
  await fetchApplicationData();
});
// Delete action
async function deleteData(id: number) {
  try {
    showLoading
    const { status } = await useFetch(config.public.API_URL + `/asset-aplikasi/${id}`, {
      method: 'DELETE',
    });

    if (status.value == 'success') {
      hideLoading();
      toast({
        title: 'Data Aplikasi Berhasil Dihapus',
        description: 'Data dengan ID Aplikasi' + id + ' berhasil dihapus.',
      })
      applications.value = applications.value.filter(item => item.id !== id);
      await fetchApplicationData();
      isDeleteDialogOpen.value = false;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Menghapus Data Aplikasi',
        description: `Data aplikasi dengan ID ${id} gagal dihapus.`,
        variant: 'destructive',
      })
    }
  } catch (error) {
    hideLoading();
    console.error('Error occured:', error);
  }
}
</script>

<template>
  <div class="bg-white rounded-lg shadow-lg">
    <div class="px-6 py-2 border-b border-gray-200 flex justify-between">
      <h1 class="text-2xl font-bold text-gray-800">Data Aplikasi</h1>
      <Button @click="fetchApplicationData()" variant="secondary">Refresh</Button>
    </div>
    <div class="px-6 py-2">
      <DataTable :columns="columns" :data="applications || []" :dataStatus="status" />
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
