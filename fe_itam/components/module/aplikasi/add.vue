<template>
  <div
    class="p-8 bg-white shadow-lg rounded-lg"
    :class="isOpen ? 'lg:ml-64' : ''"
  >
    <h1 class="text-2xl font-bold mb-6">Registrasi Aplikasi</h1>
    <form @submit.prevent="registerAsset">
      <div class="grid grid-cols-2 gap-x-4">
        <div class="mb-4">
          <label for="vendor" class="block text-sm font-medium text-gray-700"
            >Nama Vendor</label
          >
          <select
            id="vendor"
            v-model="vendor"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Vendor</option>
            <option value="vendor1">Vendor 1</option>
            <option value="vendor2">Vendor 2</option>
            <option value="vendor3">Vendor 3</option>
          </select>
        </div>

        <div class="mb-4">
          <label
            for="tanggalPembuatan"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Pembuatan</label
          >

          <input
            type="date"
            id="tanggalPembuatan"
            v-model="tanggalPembuatan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="namaAplikasi"
            class="block text-sm font-medium text-gray-700"
            >Nama Aplikasi</label
          >
          <input
            type="text"
            id="namaAplikasi"
            v-model="namaAplikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalTerima"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Terima</label
          >
          <input
            type="date"
            id="tanggalTerima"
            v-model="tanggalTerima"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tipePlatform"
            class="block text-sm font-medium text-gray-700"
            >Tipe Platform</label
          >
          <select
            id="tipePlatform"
            v-model="tipePlatform"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Platform</option>
            <option value="web">Web</option>
            <option value="mobile">Aplikasi Mobile</option>
            <option value="desktop">Aplikasi Desktop</option>
            <option value="server">Aplikasi Server</option>
          </select>
        </div>

        <div class="mb-4">
          <label
            for="serverPenyimpanan"
            class="block text-sm font-medium text-gray-700"
            >Server Penyimpanan</label
          >
          <input
            type="number"
            id="serverPenyimpanan"
            v-model="serverPenyimpanan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="urlAplikasi"
            class="block text-sm font-medium text-gray-700"
            >URL Aplikasi</label
          >
          <input
            type="text"
            id="urlAplikasi"
            v-model="urlAplikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalAktifDomain"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Aktif Domain</label
          >
          <input
            type="date"
            id="tanggalAktifDomain"
            v-model="tanggalAktifDomain"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalExpiredDomain"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Expired Domain</label
          >
          <input
            type="date"
            id="tanggalExpiredDomain"
            v-model="tanggalExpiredDomain"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="sertifikatAplikasi"
            class="block text-sm font-medium text-gray-700"
            >Sertifikat Aplikasi</label
          >
          <input
            type="text"
            id="sertifikatAplikasi"
            v-model="sertifikatAplikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="dokumentasi"
            class="block text-sm font-medium text-gray-700"
            >Dokumentasi</label
          >
          <input
            type="file"
            id="dokumentasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
      </div>

      <div class="flex justify-end mt-4 space-x-2">
        <button
          type="button"
          class="py-2 px-4 bg-white-600 text-black font-semibold rounded-md shadow-md hover:bg-red-300 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
        >
          Cancel
        </button>
        <button
          type="submit"
          class="py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
        >
          Submit
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { useState } from "#app";

const isOpen = useState("is-sidebar-open");

const vendor = ref("");
const tanggalPembuatan = ref("");
const namaAplikasi = ref("");
const tanggalTerima = ref("");
const tipePlatform = ref("");
const serverPenyimpanan = ref(null);
const urlAplikasi = ref("");
const tanggalAktifDomain = ref("");
const tanggalExpiredDomain = ref("");
const sertifikatAplikasi = ref("");
const dokumentasi = ref(null);

const registerAsset = () => {
  console.log({
    vendor: vendor.value,
    tanggalPembuatan: tanggalPembuatan.value,
    namaAplikasi: namaAplikasi.value,
    tanggalTerima: tanggalTerima.value,
    tipePlatform: tipePlatform.value,
    serverPenyimpanan: serverPenyimpanan.value,
    urlAplikasi: urlAplikasi.value,
    tanggalAktifDomain: tanggalAktifDomain.value,
    tanggalExpiredDomain: tanggalExpiredDomain.value,
    sertifikatAplikasi: sertifikatAplikasi.value,
    dokumentasi: dokumentasi.value,
  });
};
</script>

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
