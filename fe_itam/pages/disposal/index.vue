<script setup lang="ts">
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useToast } from "@/components/ui/toast/use-toast";
import { onMounted, ref } from "vue";
const config = useRuntimeConfig();
const { showLoading, hideLoading } = useLoading();
const { toast } = useToast();
const isOpen = ref(false);
const selectedRows = ref<number[]>([]);
const disposalData = ref<Disposal[]>([]);
const errorMessage = ref<string | null>(null);

interface Disposal {
  id: number;
  vendor_id: number;
  serial_number: string;
  merk: string,
  model: string,
  nomor_nota: string,
}

async function handleDelete(id: number) {
  try {
    const response = await $fetch(`/api/disposal/${id}`, {
      method: "DELETE",
    });

    toast({
      title: "Berhasil",
      description: "Data disposal berhasil dihapus",
    });

    // Refresh data setelah delete
    await fetchDisposalData();
  } catch (error) {
    console.error("Error deleting disposal:", error);
    toast({
      title: "Gagal",
      description: "Gagal menghapus data disposal",
      variant: "destructive",
    });
  }
}

async function handleBulkDelete() {
  try {
    await $fetch("/api/disposal/bulk-delete", {
      method: "DELETE",
      body: {
        ids: selectedRows.value,
      },
    });

    toast({
      title: "Berhasil",
      description: "Data disposal berhasil dihapus",
    });

    selectedRows.value = [];
    await fetchDisposalData();
  } catch (error) {
    console.error("Error bulk deleting disposal:", error);
    toast({
      title: "Gagal",
      description: "Gagal menghapus data disposal",
      variant: "destructive",
    });
  }
}

async function fetchDisposalData() {
  try {
    showLoading();
    errorMessage.value = null;

    const { data, error } = await useFetch<{
      data: Disposal[];
    }>(config.public.API_URL + "/disposal", {
      method: "GET",
    });

    if (error.value) {
      throw error.value;
    }

    disposalData.value = data.value?.data || [];


  } catch (error) {
    console.error("Error fetching disposal data:", error);
    errorMessage.value = "Gagal mengambil data disposal";
  } finally {
    hideLoading();
  }
}

onMounted(() => {
  fetchDisposalData();
});

const handleSelectAll = (checked: boolean) => {
  if (checked) {
    selectedRows.value = disposalData.value.map((item) => item.id);
  } else {
    selectedRows.value = [];
  }
};

const handleSelectRow = (checked: boolean, id: number) => {
  if (checked) {
    selectedRows.value.push(id);
  } else {
    selectedRows.value = selectedRows.value.filter((rowId) => rowId !== id);
  }
};
</script>

<template>
  <div class="">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Disposal</h1>
      </div>

      <div class="p-6">
        <div v-if="isLoading" class="text-center py-4">Memuat data...</div>

        <div v-else-if="errorMessage" class="text-center text-red-500 py-4">
          {{ errorMessage }}
        </div>

        <div v-else class="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead class="w-12">
                  <Checkbox :checked="selectedRows.length === disposalData.length" :indeterminate="selectedRows.length > 0 &&
                    selectedRows.length < disposalData.length
                    " @update:checked="handleSelectAll" />
                </TableHead>
                <TableHead class="font-semibold">Id Aset</TableHead>
                <TableHead class="font-semibold">Nomor Serial</TableHead>
                <TableHead class="font-semibold">Merek Aset</TableHead>
                <TableHead class="font-semibold">Model Aset</TableHead>
                <TableHead class="font-semibold">Nomor Nota</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-for="item in disposalData" :key="item.id">
                <TableCell>
                  <Checkbox :checked="selectedRows.includes(item.id)" @update:checked="(checked) => handleSelectRow(checked, item.id)
                    " />
                </TableCell>
                <TableCell>{{ item.id }}</TableCell>
                <TableCell>{{ item.serial_number }}</TableCell>
                <TableCell>{{ item.merk }}</TableCell>
                <TableCell>{{ item.model }}</TableCell>
                <TableCell>{{ item.nomor_nota }}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.badge-pending {
  @apply bg-yellow-100 text-yellow-800;
}

.badge-approved {
  @apply bg-green-100 text-green-800;
}

.badge-rejected {
  @apply bg-red-100 text-red-800;
}
</style>
