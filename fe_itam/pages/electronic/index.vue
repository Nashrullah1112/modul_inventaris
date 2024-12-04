<script setup lang="ts">
import { h, ref, onMounted } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type { ColumnDef } from "@tanstack/vue-table";
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

const config = useRuntimeConfig();
const router = useRouter();
const { toast } = useToast()
const { showLoading, hideLoading } = useLoading();
const devices = ref<Device[]>([]);
const isDeleteDialogOpen = ref(false);
const deviceToDelete = ref<number | null>(null);

interface Device {
  id: number;
  lokasi_penerima: string;
  waktu_penerimaan: Date;
  tanda_terima: File;
  tipe_aset: string;
  waktu_aktivasi_aset: Date;
  hasil_pemeriksaan_aset: string;
  serial_number: string;
  model: string;
  masa_garansi_mulai: Date;
  nomor_kartu_garansi: string;
  prosesor: string;
  kapasitas_ram: number;
  kapasitas_rom: number;
  tipe_ram: string;
  tipe_penyimpnanan: string;
  status_aset: string;
  nilai_aset: number;
  nilai_depresiasi: number;
  jangka_masa_pakai: number;
  waktu_aset_keluar: Date;
  kondisi_aset_keluar: string;
  nota_pembelian: File;
  divisi_id: number;
  user_id: number;
  asset_id: number;
  vendor_id: number;
}

const columns: ColumnDef<Device>[] = [
  {
    id: "select",
    header: ({ table }) =>
      h(Checkbox, {
        checked: table.getIsAllPageRowsSelected(),
        "onUpdate:checked": (value: boolean) =>
          table.toggleAllPageRowsSelected(value),
        "aria-label": "Select all",
      }),
    cell: ({ row }) =>
      h(Checkbox, {
        checked: row.getIsSelected(),
        "onUpdate:checked": (value: boolean) => row.toggleSelected(value),
        "aria-label": "Select row",
      }),
    enableSorting: false,
    enableHiding: false,
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
  },
  {
    accessorKey: "tipe_aset",
    header: ({ column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe Asset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
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
        () => ["Serial Number", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }) =>
      h("div", { class: "capitalize" }, row.getValue("serial_number")),
  },
  {
    accessorKey: "asset_id",
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
    cell: ({ row }) => h("div", {}, row.getValue("asset_id")),
  },
  {
    accessorKey: "divisi_id",
    header: () => h("div", {}, "Divisi"),
  },
  {
    accessorKey: "hasil_pemeriksaan_aset",
    header: () => h("div", {}, "Hasil Pemeriksaan"),
  },
  {
    accessorKey: "status_aset",
    header: () => h("div", {}, "Status Asset"),
  },
  {
    id: "actions",
    header: () => h("div", {}, "Aksi"),
    cell: ({ row }) => {
      return h("div", { class: "flex gap-2" }, [
        h(
          ActionBtnEdit,
          {
            onClick: () => router.push(`/electronic/${row.original.id}/edit`),
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

const fetchDeviceData = async () => {
  try {
    showLoading();
    const response = await $fetch<{ data: Device[] }>(config.public.API_URL + '/asset-perangkat', {
      method: 'GET'
    });

    if (response && response.data) {
      hideLoading();
      devices.value = response.data;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Mengambil Data Perangkat Elektronik',
        description: 'Data perangkat elektronik gagal diambil dari server.',
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


onMounted(() => {
  fetchDeviceData();
})

async function deleteData(id: number) {
  try {
    showLoading();
    const { status } = await useFetch(config.public.API_URL + `/asset-perangkat/${id}`, {
      method: 'DELETE',
    });

    if (status.value == 'success') {
      hideLoading();
      toast({
        title: 'Data Elektronik Berhasil Dihapus',
        description: 'Data elektronik dengan ID ' + id + ' berhasil dihapus.',
      })
      devices.value = devices.value.filter(item => item.id !== id);
      await fetchDeviceData();
      isDeleteDialogOpen.value = false;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Menghapus Data Elektronik',
        description: `Data elektronik dengan ID ${id} gagal dihapus.`,
      })
    }
  } catch (error) {
    hideLoading();
    console.error('Error occured:', error);
  }
}

const openDeleteConfirmation = (id: number) => {
  deviceToDelete.value = id;
  isDeleteDialogOpen.value = true;
}

const confirmDeletion = () => {
  if (deviceToDelete.value !== null) {
    deleteData(deviceToDelete.value);
  }
}
</script>

<template>
  <div class="bg-white rounded-lg shadow-lg">
    <div class="px-6 py-2 border-b border-gray-200 flex justify-between">
      <h1 class="text-2xl font-bold text-gray-800">Data Aset Elektronik</h1>
      <Button @click="fetchDeviceData()" variant="secondary">Refresh</Button>
    </div>
    <div class="px-6 py-2">
      <DataTable :columns="columns" :data="devices || []" :dataStatus="status" />
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
