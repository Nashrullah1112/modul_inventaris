<!-- <script setup lang="ts">
import UserAuthForm from "@/components/UserAuthForm.vue";
import { cn } from "@/lib/utils";
import { buttonVariants } from "@/components/ui/button";
</script> -->

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Elektronik</h1>
    <form @submit.prevent="registerAsset">
      <!-- Begi 2 sisi -->
      <div class="grid grid-cols-2 gap-x-4">
        <div class="mb-4">
          <label
            for="supplierName"
            class="block text-sm font-medium text-gray-700"
            >Nama Supplier</label
          >
          <input
            type="text"
            id="supplierName"
            v-model="supplierName"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label for="assetType" class="block text-sm font-medium text-gray-700"
            >Tipe Asset</label
          >
          <input
            type="text"
            id="assetType"
            v-model="assetType"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="notaNumber"
            class="block text-sm font-medium text-gray-700"
            >Nomor Nota</label
          >
          <input
            type="text"
            id="notaNumber"
            v-model="notaNumber"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="achievedDate"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Penerimaan</label
          >
          <input
            type="date"
            id="achievedDate"
            v-model="achievedDate"
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
      </div>

      <div class="flex justify-end mt-4">
        <button
          type="submit"
          class="py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
        >
          Daftarkan Aset
        </button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      supplierName: "",
      assetType: "",
      achievedDate: "",
      status: "",
    };
  },
  methods: {
    registerAsset() {
      // Logika untuk mendaftarkan aset
      console.log({
        supplierName: this.supplierName,
        assetType: this.assetType,
        achievedDate: this.achievedDate,
        status: this.status,
      });
    },
  },
};
</script>

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
