<script setup lang="ts">
/**
 * Props untuk menerima tipe dan data dari parent component
 * @param {string} type - Tipe aset elektronik
 * @param {any[]} data - Data array yang diterima dari parent
 */
const props = defineProps<{
  type: string;
  data: any[];
}>();

// TODO: Tambahkan API endpoint untuk:
// - GET /api/suppliers untuk dropdown supplier
// - GET /api/asset-types untuk dropdown tipe asset
// - GET /api/employees untuk dropdown penanggung jawab
// - GET /api/divisions untuk dropdown divisi
// - POST /api/assets untuk submit form

/**
 * State untuk form inputs
 * Menyimpan semua nilai input form
 */
const namasupllier = ref(""); // Menyimpan nama supplier yang dipilih
const tipeasset = ref(""); // Menyimpan tipe aset yang dipilih
const notaNumber = ref(""); // Menyimpan nomor nota pembelian
const kondisiasset = ref(""); // Menyimpan kondisi aset (Baik/Cacat/Rusak)
const status = ref(""); // Menyimpan status aset (Sewa/Aset Perusahaan/dll)
const lokasipenerima = ref(""); // Menyimpan lokasi penerimaan aset
const nilaiasset = ref(""); // Menyimpan nilai/harga aset
const achievedDate = ref(""); // Menyimpan tanggal penerimaan aset
const tandaterima = ref(""); // Menyimpan file tanda terima
const hasilpemeriksaan = ref(""); // Menyimpan file hasil pemeriksaan
const tanggalaktivasi = ref(""); // Menyimpan tanggal aktivasi perangkat
const penanggungjawabasset = ref(""); // Menyimpan penanggung jawab aset
const divpengguna = ref(""); // Menyimpan divisi pengguna aset
const notapembelian = ref(""); // Menyimpan file nota pembelian
const lokasipenyimpanan = ref(""); // Menyimpan lokasi penyimpanan aset
const jangkamasapakaiaset = ref(""); // Menyimpan masa pakai aset
const model = ref(""); // Menyimpan model perangkat
const tanggalasetkeluar = ref(""); // Menyimpan tanggal aset keluar
const nomorgaransi = ref(""); // Menyimpan nomor kartu garansi
const merk = ref(""); // Menyimpan merek perangkat
const garansimulai = ref(""); // Menyimpan tanggal mulai garansi
const garansiakhir = ref(""); // Menyimpan tanggal berakhir garansi
const serialnumber = ref(""); // Menyimpan nomor seri perangkat
const prosesor = ref(""); // Menyimpan spesifikasi prosesor
const ram = ref(""); // Menyimpan kapasitas RAM
const tiperam = ref(""); // Menyimpan tipe RAM
const rom = ref(""); // Menyimpan kapasitas ROM/storage

/**
 * Method untuk handle form submission
 * Mengirim data form ke API endpoint
 */
const registerAsset = async () => {
  try {
    // TODO: Implementasi POST request ke API endpoint
    // POST /api/assets dengan form data
  } catch (error) {
    console.error("Error submitting form:", error);
  }
};
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Elektronik</h1>
    <form @submit.prevent="registerAsset">
      <!-- Begi 2 sisi -->
      <div class="grid grid-cols-2 gap-x-4">
        <div class="mb-4">
          <label
            for="namasupplier"
            class="block text-sm font-medium text-gray-700"
            >Nama Supplier</label
          >
          <select
            id="namasupplier"
            v-model="namasupllier"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Supplier</option>
            <option value="">PT.Surya Kencana</option>
            <option value="">PT.DJAJA PUTRA</option>
            <option value="">PT INDO IT</option>
            <option value="">DB KLIK</option>
          </select>
        </div>
        <!-- </div> -->
        <div class="mb-4">
          <label
            for="namasupplier"
            class="block text-sm font-medium text-gray-700"
            >Tipe Asset</label
          >
          <select
            id="tipeasset"
            v-model="tipeasset"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Tipe Aset</option>
            <option value="">Laptop</option>
            <option value="">Mini PC</option>
            <option value="">Hanphone</option>
            <option value="">PC</option>
            <option value="">Printer</option>
            <option value="">Tablet</option>
            <option value="">Monitor</option>
            <option value="">Router</option>
            <option value="">Access Point</option>
            <option value="">Modem</option>
          </select>
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
            for="kondisi asset"
            class="block text-sm font-medium text-gray-700"
            >Kondisi Asset</label
          >
          <select
            id="kondisi asset"
            v-model="kondisiasset"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Kondisi Asset</option>
            <option value="Good">Baik</option>
            <option value="Defect">Cacat</option>
            <option value="Broke">Rusak</option>
          </select>
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
            <option value="available">Sewa</option>
            <option value="assigned">Aset Perusahaan</option>
            <option value="maintenance">Aset divisi</option>
            <option value="retired">Aset Pasif</option>
          </select>
        </div>
        <div class="mb-4">
          <label
            for="lokasipenerima"
            class="block text-sm font-medium text-gray-700"
            >Lokasi Penerimaan</label
          >
          <input
            type="text"
            id="lokasipenerima"
            v-model="lokasipenerima"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label for="nilaiaset" class="block text-sm font-medium text-gray-700"
            >Nilai Asset</label
          >
          <input
            type="number"
            id="nilaiasset"
            v-model="nilaiasset"
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
        <div class="grid w-full max-w-sm items-center gap-1.5">
          <Label for="tandaterima">Tanda Terima</Label>
          <Input
            id="tandaterima"
            type="file"
            v-model="tandaterima"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="grid w-full max-w-sm items-center gap-1.5">
          <Label for="tandaterima">Hasil Pemeriksaan</Label>
          <Input
            id="hasilpemeriksaan"
            type="file"
            v-model="hasilpemeriksaan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="achievedDate"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Aktivasi Perangkat</label
          >
          <input
            type="date"
            id="tanggalaktivasi"
            v-model="tanggalaktivasi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>

        <div class="mb-4">
          <label for="status" class="block text-sm font-medium text-gray-700"
            >Penangung Jawab Aset</label
          >
          <select
            id="penanggungjawabasset"
            v-model="penanggungjawabasset"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Pegawai</option>
            <option value="">Yunus</option>
            <option value="">Adi</option>
            <option value="">Yahya</option>
            <option value="">Lukman</option>
          </select>
        </div>

        <div class="mb-4">
          <label for="status" class="block text-sm font-medium text-gray-700"
            >Divisi Pengguna</label
          >
          <select
            id="divpengguna"
            v-model="divpengguna"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Divisi</option>
            <option value="">Purchasing</option>
            <option value="">IT</option>
            <option value="">Marketing</option>
            <option value="">Finance</option>
          </select>
        </div>
        <div>
          <Label for="tandaterima">Nota Pembelian</Label>
          <Input
            id="notapembelian"
            type="file"
            v-model="notapembelian"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="lokasipenyimpanan"
            class="block text-sm font-medium text-gray-700"
            >Lokasi Penyimpanan</label
          >
          <input
            type="text"
            id="lokasipenyimpanan"
            v-model="lokasipenyimpanan"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="jangkamasapakaiaset"
            class="block text-sm font-medium text-gray-700"
            >Jangka Masa Pakai Aset</label
          >
          <select
            id="jangkamasapakaiaset"
            v-model="jangkamasapakaiaset"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Pilih Masa Pakai</option>
            <option value="">1 Tahun</option>
            <option value="">2 Tahun</option>
            <option value="">3 Tahun</option>
            <option value="">4 Tahun</option>
            <option value="">5 Tahun</option>
            <option value="">6 Tahun</option>
            <option value="">7 Tahun</option>
            <option value="">8 Tahun</option>
          </select>
        </div>
        <div class="mb-4">
          <label for="model" class="block text-sm font-medium text-gray-700"
            >Model</label
          >
          <input
            type="text"
            id="model"
            v-model="model"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="tanggalasetkeluar"
            class="block text-sm font-medium text-gray-700"
            >Tanggal Aset Keluar</label
          >
          <input
            type="date"
            id="tanggalasetkeluar"
            v-model="tanggalasetkeluar"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="nomorgaransi"
            class="block text-sm font-medium text-gray-700"
            >Nomor Kartu Garansi</label
          >
          <input
            type="text"
            id="nomorgaransi"
            v-model="nomorgaransi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label for="merk" class="block text-sm font-medium text-gray-700"
            >Merek</label
          >
          <input
            type="text"
            id="merk"
            v-model="merk"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="garansimulai"
            class="block text-sm font-medium text-gray-700"
            >Garansi Mulai</label
          >
          <input
            type="date"
            id="garansimulai"
            v-model="garansimulai"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="garansiakhir"
            class="block text-sm font-medium text-gray-700"
            >Garansi Berakhir</label
          >
          <input
            type="date"
            id="garansiakhir"
            v-model="garansiakhir"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="nomorgaransi"
            class="block text-sm font-medium text-gray-700"
            >Nomor Kartu Garansi</label
          >
          <input
            type="text"
            id="nomorgaransi"
            v-model="nomorgaransi"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label
            for="serialnumber"
            class="block text-sm font-medium text-gray-700"
            >Serial Number</label
          >
          <input
            type="text"
            id="serialnumber"
            v-model="serialnumber"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          />
        </div>
        <div class="mb-4">
          <label for="prosesor" class="block text-sm font-medium text-gray-700"
            >Prosesor</label
          >
          <input
            type="text"
            id="prosesor"
            v-model="prosesor"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
        <div class="mb-4">
          <label for="ram" class="block text-sm font-medium text-gray-700"
            >RAM</label
          >
          <input
            type="number"
            id="ram"
            v-model="ram"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
        <div class="mb-4">
          <label for="tiperam" class="block text-sm font-medium text-gray-700"
            >Tipe RAM</label
          >
          <input
            type="text"
            id="tiperam"
            v-model="tiperam"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
        <div class="mb-4">
          <label for="rom" class="block text-sm font-medium text-gray-700"
            >ROM</label
          >
          <input
            type="number"
            id="rom"
            v-model="rom"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
      </div>
      <div class="flex justify-end mt-4 space-x-2">
        <button
          type="cancel"
          class="py-2 px-4 bg-white-600 text-black font-semibold rounded-md shadow-md hover:bg-red-300 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
        >
          Cancel
        </button>
        <button
          type="next"
          class="py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-opacity-75"
        >
          Next
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

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
