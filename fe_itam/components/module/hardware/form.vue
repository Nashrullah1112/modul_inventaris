<script setup lang="ts">

const props = defineProps<{
  type: string,
  data: any[],
}>()
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Hardware</h1>
    <form @submit.prevent="registerAsset">
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

        <div class="mb-4">
          <label for="tipeasset" class="block text-sm font-medium text-gray-700"
            >Tipe Asset</label
          >
          <select
            id="tipeasset"
            v-model="tipeasset"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          >
            <option value="" disabled>Tipe Aset</option>
            <option value="">Processor</option>
            <option value="">Motherboard</option>
            <option value="">RAM</option>
            <option value="">Storage</option>
            <option value="">Power Supply</option>
            <option value="">Casing</option>
            <option value="">Keyboard</option>
            <option value="">Mouse</option>
            <option value="">Cooling Fan</option>
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
          <label
            for="serialNumber"
            class="block text-sm font-medium text-gray-700"
            >Serial Number</label
          >
          <input
            type="text"
            id="serialNumber"
            v-model="serialNumber"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
            @input="$event.target.value = $event.target.value.toUpperCase()"
          />
        </div>

        <div class="mb-4">
          <label for="merk" class="block text-sm font-medium text-gray-700"
            >Merk</label
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
            for="spesifikasi"
            class="block text-sm font-medium text-gray-700"
            >Spesifikasi</label
          >
          <textarea
            id="spesifikasi"
            v-model="spesifikasi"
            rows="3"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            required
          ></textarea>
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

<!-- <script>
export default {
  data() {
    return {
      namasupllier: "",
      tipeasset: "",
      notaNumber: "",
      achievedDate: "",
      serialNumber: "",
      merk: "",
      model: "",
      spesifikasi: "",
      dokumentasi: null,
    };
  },
  methods: {
    handleFileUpload(event) {
      this.dokumentasi = event.target.files[0];
    },
    registerAsset() {
      console.log({
        namasupllier: this.namasupllier,
        tipeasset: this.tipeasset,
        notaNumber: this.notaNumber,
        achievedDate: this.achievedDate,
        serialNumber: this.serialNumber,
        merk: this.merk,
        model: this.model,
        spesifikasi: this.spesifikasi,
        dokumentasi: this.dokumentasi,
      });
    },
  },
};
</script> -->
