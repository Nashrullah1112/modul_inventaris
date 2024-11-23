<template>
  <div
    class="relative overflow-x-auto rounded-lg border border-gray-200 shadow-sm"
  >
    <table class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-50">
        <tr>
          <th scope="col" class="px-4 py-3">
            <input
              type="checkbox"
              v-model="selectAll"
              @change="selectAllAssets"
              class="h-4 w-4 text-blue-600 border-gray-300 rounded"
            />
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Nama Aset
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Tipe Aset
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Tanggal Masuk
          </th>
          <th
            scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
          >
            Aksi
          </th>
        </tr>
      </thead>
      <tbody class="bg-white divide-y divide-gray-200">
        <tr
          v-for="(asset, index) in assets"
          :key="index"
          class="hover:bg-gray-50"
        >
          <td class="px-4 py-4 whitespace-nowrap">
            <input
              type="checkbox"
              v-model="selectedAssets"
              :value="asset.id"
              class="h-4 w-4 text-blue-600 border-gray-300 rounded"
            />
          </td>
          <td
            class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900"
          >
            {{ asset.name }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
            {{ asset.type }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
            {{ asset.entryDate }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
            <button
              @click="viewDetail(asset)"
              class="text-blue-600 hover:text-blue-900"
            >
              Informasi Detail
            </button>
            <button
              @click="acceptAsset(asset.id)"
              class="ml-2 text-green-600 hover:text-green-900"
            >
              Terima
            </button>
            <button
              @click="rejectAsset(asset.id)"
              class="ml-2 text-red-600 hover:text-red-900"
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
      alert(`Detail untuk aset: ${asset.name}`);
    },
    acceptAsset(id: number) {
      alert(`Aset ID ${id} diterima.`);
    },
    rejectAsset(id: number) {
      alert(`Aset ID ${id} ditolak.`);
    },
  },
});
</script>

<style scoped>
/* Gaya tambahan jika dibutuhkan */
</style>
