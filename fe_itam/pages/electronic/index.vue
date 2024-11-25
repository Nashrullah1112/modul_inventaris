<script setup lang="ts">
import { h, ref, onMounted } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";

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

//state for application data
const devices = ref<Device[]>([]);
const isLoading = ref(false);
const errorMessage = ref<string | null>(null);

async function fetchDevices() {
  try {
    isLoading.value = true;
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Device[];
      error: any;
    }>("http://103.127.139.11:5000/api/asset-perangkat", {
      method: "GET",
    });

    if (response.error) {
      throw new Error(response.error);
    }
    // Extract the data array from the response
    devices.value = response.data;
  } catch (error) {
    console.error("Error fetching devices:", error);
    errorMessage.value = "Failed to load device data. Please try again.";
  } finally {
    isLoading.value = false;
  }
}

onMounted(fetchDevices);

// Transform API data for table usage
const electronics = computed(() =>
  devices.value.map((device) => ({
    id: device.id,
    serialNumber: device.serial_number,
    assetNumber: device.asset_id.toString(),
    deviceName: device.model,
    division: device.divisi_id.toString(),
  }))
);

//define table columns
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
    accessorKey: "serialNumber",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Serial Number", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", { class: "capitalize" }, row.getValue("serialNumber")),
  },
  {
    accessorKey: "assetNumber",
    header: ({ column }: { column: Column }) => {
      return h(
        Button,
        {
          variant: "ghost",
          onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
        },
        () => ["Nomor Asset", h(ArrowUpDown, { class: "ml-2 h-4 w-4" })]
      );
    },
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("assetNumber")),
  },
  {
    accessorKey: "deviceName",
    header: () => h("div", {}, "Data Perangkat"),
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("deviceName")),
  },
  {
    accessorKey: "division",
    header: () => h("div", {}, "Divisi"),
    cell: ({ row }: { row: TableRow }) =>
      h("div", {}, row.getValue("division")),
  },
  {
    id: "actions",
    header: () => h("div", {}, "Aksi"),
    cell: ({ row }: { row: TableRow }) => {
      return h("div", { class: "flex gap-2" }, [
        h(
          Button,
          {
            variant: "default",
            class: "bg-blue-500 hover:bg-blue-600",
            onClick: () => console.log("Update", row.original.id),
          },
          () => "Update"
        ),
        h(
          Button,
          {
            variant: "default",
            class: "bg-red-500 hover:bg-red-600",
            onClick: () => console.log("Delete", row.original.id),
          },
          () => "Delete"
        ),
      ]);
    },
  },
];

const isOpen = ref(false);
</script>

<template>
  <div class="p-8" :class="{ 'ml-64': isOpen, 'ml-20': !isOpen }">
    <div class="bg-white rounded-lg shadow-lg">
      <div class="p-6 border-b border-gray-200">
        <h1 class="text-2xl font-bold text-gray-800">Data Aset Elektronik</h1>
      </div>

      <div class="p-6">
        <DataTable :columns="columns" :data="electronics" />
      </div>
    </div>
  </div>
</template>
