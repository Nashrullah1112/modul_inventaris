<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Inspeksi Aset</h1>
      </div>

      <div class="p-6">
        <div class="flex items-center justify-between py-4">
          <div class="flex items-center space-x-4">
            <Input class="max-w-sm" placeholder="Cari..." v-model="search" />
            <div class="flex items-center space-x-2">
              <Popover>
                <PopoverTrigger asChild>
                  <Button
                    variant="outline"
                    class="w-[240px] justify-start text-left font-normal"
                  >
                    <CalendarIcon class="mr-2 h-4 w-4" />
                    {{
                      tanggalMasuk
                        ? format(tanggalMasuk, "PPP")
                        : "Filter Tanggal Masuk"
                    }}
                  </Button>
                </PopoverTrigger>
                <PopoverContent class="w-auto p-0" align="start">
                  <Calendar mode="single" v-model="tanggalMasuk" initialFocus />
                </PopoverContent>
              </Popover>
            </div>
          </div>
          <div class="flex items-center space-x-2">
            <Select v-model="pageSize">
              <SelectTrigger class="w-[180px]">
                <SelectValue placeholder="Entri per halaman" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem :value="10">10</SelectItem>
                <SelectItem :value="20">20</SelectItem>
                <SelectItem :value="30">30</SelectItem>
                <SelectItem :value="40">40</SelectItem>
                <SelectItem :value="50">50</SelectItem>
              </SelectContent>
            </Select>
          </div>
        </div>

        <div class="rounded-md border">
          <Table>
            <TableHeader>
              <TableRow>
                <TableHead class="w-[50px]">
                  <Checkbox />
                </TableHead>
                <TableHead>Nama Aset</TableHead>
                <TableHead>Tipe Aset</TableHead>
                <TableHead>Tanggal Masuk</TableHead>
                <TableHead>Lihat</TableHead>
                <TableHead>Terima Barang</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow v-for="aset in filteredAsetList" :key="aset.id">
                <TableCell>
                  <Checkbox />
                </TableCell>
                <TableCell>{{ aset.namaAset }}</TableCell>
                <TableCell>{{ aset.tipeAset }}</TableCell>
                <TableCell>{{ aset.tanggalMasuk }}</TableCell>
                <TableCell>
                  <Button variant="secondary" size="sm">Rincian</Button>
                </TableCell>
                <TableCell>
                  <div class="flex space-x-2">
                    <Button
                      variant="outline"
                      size="icon"
                      class="bg-green-500 hover:bg-green-600 text-white"
                    >
                      <CheckIcon class="h-4 w-4" />
                    </Button>
                    <Button variant="destructive" size="icon">
                      <XIcon class="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>

        <div class="flex items-center justify-end space-x-2 py-4">
          <div
            class="flex w-[100px] items-center justify-center text-sm font-medium"
          >
            Halaman {{ halamanSaatIni }} dari {{ totalHalaman }}
          </div>
          <div class="flex items-center space-x-2">
            <Button
              variant="outline"
              size="sm"
              @click="halamanSebelumnya"
              :disabled="halamanSaatIni === 1"
            >
              <ChevronLeftIcon class="h-4 w-4" />
              Sebelumnya
            </Button>
            <Button
              variant="outline"
              size="sm"
              @click="halamanBerikutnya"
              :disabled="halamanSaatIni === totalHalaman"
            >
              Berikutnya
              <ChevronRightIcon class="h-4 w-4" />
            </Button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import {
  CheckIcon,
  XIcon,
  CalendarIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from "lucide-vue-next";
import { format } from "date-fns";

const search = ref("");
const pageSize = ref(10);
const halamanSaatIni = ref(1);
const isOpen = useState("is-sidebar-open", () => false);
const tanggalMasuk = ref(null);

const asetList = ref([
  {
    id: 1,
    namaAset: "Lenovo Ideapad",
    tipeAset: "Hardware",
    tanggalMasuk: "18 Mei 2024",
  },
  // ... tambahkan data lainnya
]);

const filteredAsetList = computed(() => {
  let filtered = [...asetList.value];

  if (search.value) {
    filtered = filtered.filter(
      (aset) =>
        aset.namaAset.toLowerCase().includes(search.value.toLowerCase()) ||
        aset.tipeAset.toLowerCase().includes(search.value.toLowerCase())
    );
  }

  if (tanggalMasuk.value) {
    const selectedDate = format(tanggalMasuk.value, "dd MMM yyyy");
    filtered = filtered.filter((aset) => aset.tanggalMasuk === selectedDate);
  }

  return filtered;
});

const totalHalaman = computed(() =>
  Math.ceil(filteredAsetList.value.length / pageSize.value)
);

const halamanSebelumnya = () => {
  if (halamanSaatIni.value > 1) halamanSaatIni.value--;
};

const halamanBerikutnya = () => {
  if (halamanSaatIni.value < totalHalaman.value) halamanSaatIni.value++;
};

const keHalaman = (halaman) => {
  halamanSaatIni.value = halaman;
};
</script>
