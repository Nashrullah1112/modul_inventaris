<script setup lang="ts">
import UserAuthForm from "@/components/UserAuthForm.vue";
import { cn } from "@/lib/utils";
import { buttonVariants } from "@/components/ui/button";
</script>

<template>
  <div class="max-w-lg mx-auto p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset IT</h1>
    <form @submit.prevent="registerAsset">
      <div class="mb-4">
        <label for="assetName" class="block text-sm font-medium text-gray-700"
          >Nama Aset</label
        >
        <input
          type="text"
          id="assetName"
          v-model="assetName"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          required
        />
      </div>
      <div class="mb-4">
        <label
          for="serialNumber"
          class="block text-sm font-medium text-gray-700"
          >Nomor Seri</label
        >
        <input
          type="text"
          id="serialNumber"
          v-model="serialNumber"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          required
        />
      </div>
      <div class="mb-4">
        <label
          for="purchaseDate"
          class="block text-sm font-medium text-gray-700"
          >Tanggal Pembelian</label
        >
        <input
          type="date"
          id="purchaseDate"
          v-model="purchaseDate"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          required
        />
      </div>
      <div class="mb-4">
        <label for="status" class="block text-sm font-medium text-gray-700"
          >Status</label
        >
        <select
          id="status"
          v-model="status"
          class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          required
        >
          <option value="" disabled>Pilih status</option>
          <option value="available">Tersedia</option>
          <option value="assigned">Dipinjam</option>
          <option value="maintenance">Dalam perbaikan</option>
          <option value="retired">Pensiun</option>
        </select>
      </div>
      <button
        type="submit"
        class="w-full py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
      >
        Daftarkan Aset
      </button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      assetName: "",
      serialNumber: "",
      purchaseDate: "",
      status: "",
    };
  },
  methods: {
    registerAsset() {
      // Logika untuk mendaftarkan aset
      console.log({
        assetName: this.assetName,
        serialNumber: this.serialNumber,
        purchaseDate: this.purchaseDate,
        status: this.status,
      });
    },
  },
};
</script>

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
