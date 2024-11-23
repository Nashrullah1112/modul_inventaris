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
import { ArrowUpDown } from "lucide-vue-next";

const companies = ref([
  {
    id: 1,
    picName: "John Doe",
    email: "john.doe@company.com",
    contactNumber: "+62812345678",
    companyLocation: "Jakarta Selatan",
    siupNumber: "SIUP/2023/123456",
    nibNumber: "NIB/2023/789012",
    npwpNumber: "12.345.678.9-012.000",
  },
  {
    id: 2,
    picName: "Jane Smith",
    email: "jane.smith@company.com",
    contactNumber: "+62887654321",
    companyLocation: "Jakarta Pusat",
    siupNumber: "SIUP/2023/654321",
    nibNumber: "NIB/2023/210987",
    npwpNumber: "98.765.432.1-098.000",
  },
]);

const isOpen = ref(false);
const selectedRows = ref<number[]>([]);

const toggleSelectAll = (checked: boolean) => {
  if (checked) {
    selectedRows.value = companies.value.map((company) => company.id);
  } else {
    selectedRows.value = [];
  }
};

const toggleSelectRow = (id: number) => {
  const index = selectedRows.value.indexOf(id);
  if (index === -1) {
    selectedRows.value.push(id);
  } else {
    selectedRows.value.splice(index, 1);
  }
};

const handleUpdate = (company: any) => {
  console.log("Update company:", company);
  // Implement your update logic here
};

const handleDelete = (company: any) => {
  console.log("Delete company:", company);
  // Implement your delete logic here
};
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Perusahaan</h1>
      </div>

      <div class="p-6">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead class="w-12">
                <Checkbox
                  :checked="selectedRows.length === companies.length"
                  :indeterminate="
                    selectedRows.length > 0 &&
                    selectedRows.length < companies.length
                  "
                  @update:checked="toggleSelectAll"
                />
              </TableHead>
              <TableHead>
                <div class="flex items-center">
                  Nama PIC
                  <ArrowUpDown class="ml-2 h-4 w-4" />
                </div>
              </TableHead>
              <TableHead>
                <div class="flex items-center">
                  Email
                  <ArrowUpDown class="ml-2 h-4 w-4" />
                </div>
              </TableHead>
              <TableHead>Nomor Kontak</TableHead>
              <TableHead>Lokasi Perusahaan</TableHead>
              <TableHead>Nomor SIUP</TableHead>
              <TableHead>Nomor NIB</TableHead>
              <TableHead>Nomor NPWP</TableHead>
              <TableHead>Aksi</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow
              v-for="company in companies"
              :key="company.id"
              class="hover:bg-muted/50"
            >
              <TableCell class="w-12">
                <Checkbox
                  :checked="selectedRows.includes(company.id)"
                  @update:checked="() => toggleSelectRow(company.id)"
                />
              </TableCell>
              <TableCell>{{ company.picName }}</TableCell>
              <TableCell>{{ company.email }}</TableCell>
              <TableCell>{{ company.contactNumber }}</TableCell>
              <TableCell>{{ company.companyLocation }}</TableCell>
              <TableCell>{{ company.siupNumber }}</TableCell>
              <TableCell>{{ company.nibNumber }}</TableCell>
              <TableCell>{{ company.npwpNumber }}</TableCell>
              <TableCell>
                <div class="flex items-center gap-2">
                  <Button
                    size="sm"
                    variant="default"
                    class="bg-blue-500 hover:bg-blue-600 text-white"
                    @click="handleUpdate(company)"
                  >
                    Edit
                  </Button>
                  <Button
                    size="sm"
                    variant="default"
                    class="bg-red-500 hover:bg-red-600 text-white"
                    @click="handleDelete(company)"
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
</template>

<style scoped>
.button-container {
  display: flex;
  gap: 0.5rem;
}
</style>
