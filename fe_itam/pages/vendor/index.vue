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
const vendors = ref<Vendor[]>([]);
const isDeleteDialogOpen = ref(false);
const vendorToDelete = ref<number | null>(null);



interface Vendor {
  id: number
  nama_pic: string
  email: string
  nomor_kontak: string
  lokasi_perusahaan: string
  nomor_siup: string
  nomor_nib: string
  nomor_npwp: string
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
            onClick: () => router.push(`/vendor/${row.original.id}/edit`),
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
  vendorToDelete.value = id;
  isDeleteDialogOpen.value = true;
}

const confirmDeletion = () => {
  if (vendorToDelete.value !== null) {
    deleteData(vendorToDelete.value);
  }
}

/* fetch data from api */

const fetchVendoData = async () => {
  try {
    showLoading();
    const response = await $fetch<{ data: Vendor[] }>(config.public.API_URL + '/vendor', {
      method: 'GET'
    });

    if (response && response.data) {
      hideLoading();
      vendors.value = response.data;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Mengambil Data Vendor',
        description: 'Data vendor gagal diambil dari server.',
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
  fetchVendoData();
});

// Delete action
async function deleteData(id: number) {
  try {
    showLoading();
    const { status } = await useFetch(config.public.API_URL + `/vendor/${id}`, {
      method: 'DELETE',
    });

    if (status.value == 'success') {
      hideLoading();
      vendors.value = vendors.value.filter((vendor) => vendor.id !== id);
      toast({
        title: 'Berhasil Menghapus Data Vendor',
        description: `Data vendor dengan ID ${id} berhasil dihapus.`,
        variant: 'default',
      });
      await fetchVendoData();
      isDeleteDialogOpen.value = false;
    } else {
      hideLoading();
      toast({
        title: 'Gagal Menghapus Data Vendor',
        description: `Data vendor dengan ID ${id} gagal dihapus. Mungkin vendor terelasi dengan data lain.`,
        variant: 'destructive',
      })
    }
  } catch (error) {
    hideLoading();
    console.error('Error occured:', error);
  } finally {
    hideLoading();
  }
}
</script>

<template>
  <div class="bg-white rounded-lg shadow-lg">
    <div class="px-6 py-2 border-b border-gray-200 flex justify-between">
      <h1 class="text-2xl font-bold text-gray-800">Data Vendor</h1>
      <Button @click="fetchVendoData()" variant="secondary">Refresh</Button>
    </div>
    <div class="px-6 py-2">
      <DataTable :columns="columns" :data="vendors || []" :dataStatus="status" />
    </div>
    <AlertDialog v-model:open="isDeleteDialogOpen">
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Apakah Anda yakin ingin menghapus?</AlertDialogTitle>
          <AlertDialogDescription>
            Data vendor yang Anda pilih akan dihapus secara permanen. Tidakan ini tidak dapat dibatalkan.
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
