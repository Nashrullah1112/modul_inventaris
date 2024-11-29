<script setup lang="ts">
import { ref, onMounted } from "vue";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import { useToast } from "@/components/ui/toast/use-toast";

const { toast } = useToast();

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

interface Disposal {
  id: number;
  asset_number: string;
  asset_name: string;
  asset_type: string;
  disposal_date: string;
  disposal_reason: string;
  disposal_status: string;
}

const disposalData = ref<Disposal[]>([]);
const isLoading = ref(false);
const errorMessage = ref<string | null>(null);

async function fetchDisposalData() {
  try {
    isLoading.value = true;
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Disposal[];
    }>("/api/disposal", {
      method: "GET",
    });

    disposalData.value = response.data;
  } catch (error) {
    console.error("Error fetching disposal data:", error);
    errorMessage.value = "Gagal mengambil data disposal";
  } finally {
    isLoading.value = false;
  }
}

onMounted(() => {
  fetchDisposalData();
});

const isOpen = ref(false);
const selectedRows = ref<number[]>([]);

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
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
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
                  <Checkbox
                    :checked="selectedRows.length === disposalData.length"
                    :indeterminate="
                      selectedRows.length > 0 &&
                      selectedRows.length < disposalData.length
                    "
                    @update:checked="handleSelectAll"
                  />
                </TableHead>
                <TableHead class="font-semibold">Nomor Aset</TableHead>
                <TableHead class="font-semibold">Nama Aset</TableHead>
                <TableHead class="font-semibold">Tipe Aset</TableHead>
                <TableHead class="font-semibold">Tanggal Disposal</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-for="item in disposalData" :key="item.id">
                <TableCell>
                  <Checkbox
                    :checked="selectedRows.includes(item.id)"
                    @update:checked="
                      (checked) => handleSelectRow(checked, item.id)
                    "
                  />
                </TableCell>
                <TableCell>{{ item.asset_number }}</TableCell>
                <TableCell>{{ item.asset_name }}</TableCell>
                <TableCell>{{ item.asset_type }}</TableCell>
                <TableCell>{{ item.disposal_date }}</TableCell>
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
