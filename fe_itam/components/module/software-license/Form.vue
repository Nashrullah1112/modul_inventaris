<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Lisensi Software</h1>
    <form @submit.prevent="registerLicense">
      <div class="grid grid-cols-2 gap-x-4">
        <div class="mb-4">
          <label
            for="namavendor"
            class="block text-sm font-medium text-gray-700"
            >Nama Vendor</label
          >
          <input
            type="text"
            id="namavendor"
            v-model="namavendor"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalpembuatan"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Pembuatan</label
          >
          <input
            type="date"
            id="tanggalpembuatan"
            v-model="tanggalpembuatan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="namaaplikasi"
            class="block text-sm font-medium text-gray-700"
            >Nama Aplikasi</label
          >
          <input
            type="text"
            id="namaaplikasi"
            v-model="namaaplikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalpenerimaan"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Penerimaan</label
          >
          <input
            type="date"
            id="tanggalpenerimaan"
            v-model="tanggalpenerimaan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tipeplatform"
            class="block text-sm font-medium text-gray-700"
            >Tipe Platform</label
          >
          <select
            id="tipeplatform"
            v-model="tipeplatform"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Platform</option>
            <option value="web">Web</option>
            <option value="mobile">Mobile</option>
            <option value="desktop">Desktop</option>
          </select>
        </div>

        <div class="mb-4">
          <label
            for="lokasiserver"
            class="block text-sm font-medium text-gray-700"
            >Lokasi Server Penyimpanan</label
          >
          <input
            type="text"
            id="lokasiserver"
            v-model="lokasiserver"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="urlaplikasi"
            class="block text-sm font-medium text-gray-700"
            >URL Aplikasi</label
          >
          <input
            type="text"
            id="urlaplikasi"
            v-model="urlaplikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalaktifdomain"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Aktif Domain</label
          >
          <input
            type="date"
            id="tanggalaktifdomain"
            v-model="tanggalaktifdomain"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tipesertifikasi"
            class="block text-sm font-medium text-gray-700"
            >Tipe Sertifikasi Asset</label
          >
          <input
            type="text"
            id="tipesertifikasi"
            v-model="tipesertifikasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tanggalexpireddomain"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Expired Domain</label
          >
          <input
            type="date"
            id="tanggalexpireddomain"
            v-model="tanggalexpireddomain"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label for="serialkey" class="block text-sm font-medium text-gray-700"
            >Serial Key</label
          >
          <input
            type="text"
            id="serialkey"
            v-model="serialkey"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="serialnumberdevice"
            class="block text-sm font-medium text-gray-700"
            >Serial Number Device Terpasang</label
          >
          <select
            id="serialnumberdevice"
            v-model="serialnumberdevice"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Serial Number Device</option>
            <option value="device1">Device 1</option>
            <option value="device2">Device 2</option>
            <option value="device3">Device 3</option>
          </select>
        </div>

        <div class="mb-4">
          <label
            for="kategorilisensi"
            class="block text-sm font-medium text-gray-700"
            >Kategori Lisensi</label
          >
          <input
            type="text"
            id="kategorilisensi"
            v-model="kategorilisensi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="tipelisensi"
            class="block text-sm font-medium text-gray-700"
            >Tipe Lisensi</label
          >
          <select
            id="tipelisensi"
            v-model="tipelisensi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Tipe Lisensi</option>
            <option value="perpetual">Perpetual</option>
            <option value="subscription">Subscription</option>
            <option value="trial">Trial</option>
          </select>
        </div>

        <div class="mb-4">
          <label
            for="versilisensi"
            class="block text-sm font-medium text-gray-700"
            >Versi Lisensi</label
          >
          <input
            type="text"
            id="versilisensi"
            v-model="versilisensi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="estimasipemakaian"
            class="block text-sm font-medium text-gray-700"
            >Estimasi Pemakaian Software</label
          >
          <input
            type="text"
            id="estimasipemakaian"
            v-model="estimasipemakaian"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="maksimumpengguna"
            class="block text-sm font-medium text-gray-700"
            >Maksimum Pengguna Lisensi</label
          >
          <input
            type="number"
            id="maksimumpengguna"
            v-model="maksimumpengguna"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="limitlisensi"
            class="block text-sm font-medium text-gray-700"
            >Limit Lisensi</label
          >
          <input
            type="text"
            id="limitlisensi"
            v-model="limitlisensi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label
            for="maksimumperangkat"
            class="block text-sm font-medium text-gray-700"
            >Maksimum Perangkat Pengguna Lisensi</label
          >
          <input
            type="number"
            id="maksimumperangkat"
            v-model="maksimumperangkat"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4 grid w-full max-w-sm items-center gap-1.5">
          <label
            for="dokumentasi"
            class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
            >Dokumentasi</label
          >
          <input
            id="dokumentasi"
            type="file"
            @change="handleFileUpload"
            class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
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

<script>
export default {
  data() {
    return {
      namavendor: "",
      tanggalpembuatan: "",
      namaaplikasi: "",
      tanggalpenerimaan: "",
      tipeplatform: "",
      lokasiserver: "",
      urlaplikasi: "",
      tanggalaktifdomain: "",
      tipesertifikasi: "",
      tanggalexpireddomain: "",
      dokumentasi: null,
      serialkey: "",
      serialnumberdevice: "",
      kategorilisensi: "",
      tipelisensi: "",
      versilisensi: "",
      estimasipemakaian: "",
      maksimumpengguna: "",
      limitlisensi: "",
      maksimumperangkat: "",
    };
  },
  methods: {
    handleFileUpload(event) {
      this.dokumentasi = event.target.files[0];
    },
    registerLicense() {
      // Logic untuk mendaftarkan lisensi
      console.log({
        namavendor: this.namavendor,
        tanggalpembuatan: this.tanggalpembuatan,
        namaaplikasi: this.namaaplikasi,
        tanggalpenerimaan: this.tanggalpenerimaan,
        tipeplatform: this.tipeplatform,
        lokasiserver: this.lokasiserver,
        urlaplikasi: this.urlaplikasi,
        tanggalaktifdomain: this.tanggalaktifdomain,
        tipesertifikasi: this.tipesertifikasi,
        tanggalexpireddomain: this.tanggalexpireddomain,
        dokumentasi: this.dokumentasi,
        serialkey: this.serialkey,
        serialnumberdevice: this.serialnumberdevice,
        kategorilisensi: this.kategorilisensi,
        tipelisensi: this.tipelisensi,
        versilisensi: this.versilisensi,
        estimasipemakaian: this.estimasipemakaian,
        maksimumpengguna: this.maksimumpengguna,
        limitlisensi: this.limitlisensi,
        maksimumperangkat: this.maksimumperangkat,
      });
    },
  },
};
</script>
