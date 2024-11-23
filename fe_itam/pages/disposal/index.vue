<script setup lang="ts">
import { ref } from "vue";
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

const disposalData = ref([
  {
    id: 1,
    assetNumber: "AST-2024-001",
    assetName: "Laptop Dell XPS 13",
    assetType: "Perangkat Elektronik",
    disposalDate: "2024-02-20",
    status: "Pending",
  },
  {
    id: 2,
    assetNumber: "AST-2024-002",
    assetName: "Printer HP LaserJet",
    assetType: "Perangkat Elektronik",
    disposalDate: "2024-02-21",
    status: "Approved",
  },
  {
    id: 3,
    assetNumber: "AST-2024-003",
    assetName: "Monitor LG 27inch",
    assetType: "Hardware",
    disposalDate: "2024-02-22",
    status: "Rejected",
  },
]);

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

const getStatusBadgeVariant = (status: string) => {
  switch (status) {
    case "Pending":
      return "warning";
    case "Approved":
      return "success";
    case "Rejected":
      return "destructive";
    default:
      return "secondary";
  }
};

const handleUpdate = (id: number) => {
  console.log("Update disposal:", id);
};

const handleDelete = (id: number) => {
  console.log("Delete disposal:", id);
};
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Disposal</h1>
      </div>

      <div class="p-6">
        <div class="rounded-md border">
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
                <TableHead class="font-semibold">Status</TableHead>
                <TableHead class="font-semibold">Aksi</TableHead>
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
                <TableCell>{{ item.assetNumber }}</TableCell>
                <TableCell>{{ item.assetName }}</TableCell>
                <TableCell>{{ item.assetType }}</TableCell>
                <TableCell>{{ item.disposalDate }}</TableCell>
                <TableCell>
                  <Badge :variant="getStatusBadgeVariant(item.status)">
                    {{ item.status }}
                  </Badge>
                </TableCell>
                <TableCell>
                  <div class="flex items-center gap-2">
                    <Button
                      variant="default"
                      size="sm"
                      class="bg-blue-500 hover:bg-blue-600 text-white"
                      @click="handleUpdate(item.id)"
                    >
                      Edit
                    </Button>
                    <Button
                      variant="destructive"
                      size="sm"
                      @click="handleDelete(item.id)"
                    >
                      Hapus
                    </Button>
                  </div>
                </TableCell>
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
