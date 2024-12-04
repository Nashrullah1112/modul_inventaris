<script setup lang="ts">
import { h, ref } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import ActionBtnEdit from "~/components/atoms/ActionBtnEdit.vue";
import ActionBtnDelete from "~/components/atoms/ActionBtnDelete.vue";
import { useToast } from "~/components/ui/toast";
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
const { showLoading, hideLoading } = useLoading();
const { toast } = useToast();
const router = useRouter();
const hardwareData = ref<Hardware[]>([]);
const errorMessage = ref<string | null>(null);
const isDeleteDialogOpen = ref(false);
const hardwareToDelete = ref<number | null>(null);



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

interface Hardware {
  id: number
  waktu_penerimaan: string
  bukti_penerimaan: string
  tipe_aset: string
  tanggal_aktivasi_perangkat: string
  hasil_pemeriksaan: string
  serial_number: string
  model: string
  waktu_garansi_mulai: string
  waktu_garansi_berakhir: string
  nomor_kartu_garansi: string
  spesifikasi_perangkat: string
  status_aset: string
  penanggungjawab_aset: string
  lokasi_penyimpanan_id: string
  jangka_masa_pakai: number
  waktu_aset_keluar: string
  kondisi_aset: string
  nota_pembelian: string
  divisi_id: number
  asset_id: number
  asset: HardwareAsset
}

interface HardwareAsset {
  id: number
  vendor_id: number
  serial_number: string
  merk: string
  model: string
  nomor_nota: string
}

const fetchHardware = async () => {
  try {
    showLoading();
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Hardware[];
      error: any;
    }>(config.public.API_URL + "/asset-hardware", {
      method: "GET",
    });

    if (response.error) {
      throw new Error(response.error);
    }
    hardwareData.value = response.data;
  } catch (error) {
    hideLoading();
    console.error("Error fetching devices:", error);
    errorMessage.value = "Failed to load device data. Please try again.";
  } finally {
    hideLoading();
  }
}

const deleteData = async (id: number) => {
  try {
    showLoading();
    errorMessage.value = null;

    const { status, data, error } = await useFetch<{
      message: string;
      data: Hardware[];
      error: any;
    }>(config.public.API_URL + "/asset-hardware/" + id, {
      method: "DELETE",
    });

    hardwareData.value = hardwareData.value.filter(item => item.id !== id);

    toast({
      title: "Deletion Successful",
      description: "The hardware item has been deleted.",
      variant: "default"
    });

  } catch (error) {
    hideLoading();
    console.error("Error deleting device:", error);
    errorMessage.value = "Failed to delete device. Please try again.";
    toast({
      title: "Deletion Failed",
      description: "Could not delete the hardware item." + error,
      variant: "destructive"
    });
  } finally {
    hideLoading();
  }
}

const openDeleteConfirmation = (id: number) => {
  hardwareToDelete.value = id;
  isDeleteDialogOpen.value = true;
}

const confirmDeletion = () => {
  if (hardwareToDelete.value !== null) {
    deleteData(hardwareToDelete.value);
  }
}

onMounted(() => {
  fetchHardware()
})

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
    accessorKey: "tipe_aset",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Tipe", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "model",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Model", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "serial_number",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Serial Number", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "status_aset",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Status", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "lokasi_penyimpanan_id",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Lokasi Penyimpanan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "kondisi_aset",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Kondisi", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "hasil_pemeriksaan",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Hasil Pemeriksaan", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      ),
  },
  {
    accessorKey: "penanggungjawab_aset",
    header: ({ column }: { column: any }) =>
      h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Penanggung Jawab", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
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

</script>

<template>
  <div class="bg-white rounded-lg shadow-lg">
    <div class="px-6 py-2 border-b border-gray-200 flex justify-between">
      <h1 class="text-2xl font-bold text-gray-800">Data Hardware</h1>
      <Button @click="fetchHardware()" variant="secondary">Refresh</Button>
    </div>
    <div class="px-6 py-2">
      <DataTable :columns="columns" :data="hardwareData || []" :dataStatus="status" />
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
