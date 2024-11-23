<template>
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
    <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
      <thead
        class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
      >
        <tr>
          <th scope="col" class="px-6 py-3">
            <input
              type="checkbox"
              v-model="selectAll"
              @change="selectAllAssets"
            />
          </th>
          <th scope="col" class="px-6 py-3">Nama Aset</th>
          <th scope="col" class="px-6 py-3">Tipe Aset</th>
          <th scope="col" class="px-6 py-3">Tanggal Masuk</th>
          <th scope="col" class="px-6 py-3">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(asset, index) in assets"
          :key="index"
          class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
        >
          <td class="px-6 py-4">
            <input type="checkbox" v-model="selectedAssets" :value="asset.id" />
          </td>
          <td class="px-6 py-4">{{ asset.name }}</td>
          <td class="px-6 py-4">{{ asset.type }}</td>
          <td class="px-6 py-4">{{ asset.entryDate }}</td>
          <td class="px-6 py-4">
            <button
              @click="viewDetail(asset)"
              class="text-blue-600 hover:text-blue-900"
            >
              Informasi Detail
            </button>
            <button
              @click="acceptAsset(asset.id)"
              class="text-green-600 hover:text-green-900 ml-2"
            >
              Terima
            </button>
            <button
              @click="rejectAsset(asset.id)"
              class="text-red-600 hover:text-red-900 ml-2"
            >
              Tolak
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

interface Asset {
  id: number;
  name: string;
  type: string;
  entryDate: string;
}

export default defineComponent({
  name: "ApprovalTable",
  props: {
    assets: {
      type: Array as () => Asset[],
      required: true,
    },
  },
  data() {
    return {
      selectedAssets: [] as number[],
      selectAll: false,
    };
  },
  methods: {
    selectAllAssets() {
      if (this.selectAll) {
        this.selectedAssets = this.assets.map((asset) => asset.id);
      } else {
        this.selectedAssets = [];
      }
    },
    viewDetail(asset: Asset) {
      // Tampilkan modal atau halaman detail aset
      alert(`Detail untuk aset: ${asset.name}`);
    },
    acceptAsset(id: number) {
      // Logika untuk menerima aset
      alert(`Aset ID ${id} diterima.`);
    },
    rejectAsset(id: number) {
      // Logika untuk menolak aset
      alert(`Aset ID ${id} ditolak.`);
    },
  },
});
</script>

<style scoped>
/* Tambahkan gaya CSS sesuai kebutuhan */
</style>
