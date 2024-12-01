<script setup lang="ts">
import { h, ref, onMounted } from "vue";
import DataTable from "@/components/DataTable.vue";
import { ArrowUpDown } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import type { ColumnDef } from "@tanstack/vue-table";


const config = useRuntimeConfig();

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

interface Electronic {
  id: number;
  serialNumber: string;
  assetNumber: string;
  deviceName: string;
  division: string;
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
    }>(config.public.API_URL + "/asset-perangkat", {
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

const columns: ColumnDef<Electronic>[] = [
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
    accessorKey: "serialNumber",
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
      h("div", { class: "capitalize" }, row.getValue("serialNumber")),
  },
  {
    accessorKey: "assetNumber",
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
    cell: ({ row }) => h("div", {}, row.getValue("assetNumber")),
  },
  {
    accessorKey: "deviceName",
    header: () => h("div", {}, "Data Perangkat"),
    cell: ({ row }) => h("div", {}, row.getValue("deviceName")),
  },
  {
    accessorKey: "division",
    header: () => h("div", {}, "Divisi"),
    cell: ({ row }) => h("div", {}, row.getValue("division")),
  },
  {
    id: "actions",
    header: () => h("div", {}, "Aksi"),
    cell: ({ row }) => {
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
</script>

<template>
  <div class="container mx-auto py-10">
    <div class="rounded-md border">
      <div class="p-6 border-b">
        <div class="flex items-center justify-between">
          <h1 class="text-lg font-medium">Data Aset Elektronik</h1>
          <Button variant="outline" @click="router.push('/electronic/add')">
            Tambah Data
          </Button>
        </div>
      </div>

      <div class="p-6">
        <DataTable :columns="columns" :data="electronics" />
      </div>
    </div>
  </div>
</template>
