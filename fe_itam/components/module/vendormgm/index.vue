<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Tambah Vendor</h1>

    <Form @submit="onSubmit">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="namaPerusahaan" v-slot="{ field }">
          <FormItem>
            <FormLabel
              >Nama Perusahaan <span class="text-red-500">*</span></FormLabel
            >
            <FormControl>
              <Input v-model="values.namaPerusahaan" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="namaPIC" v-slot="{ field }">
          <FormItem>
            <FormLabel>Nama PIC <span class="text-red-500">*</span></FormLabel>
            <FormControl>
              <Input v-model="values.namaPIC" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="email" v-slot="{ field }">
          <FormItem>
            <FormLabel>Email <span class="text-red-500">*</span></FormLabel>
            <FormControl>
              <Input v-model="values.email" type="email" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorKontak" v-slot="{ field }">
          <FormItem>
            <FormLabel
              >Nomor Kontak <span class="text-red-500">*</span></FormLabel
            >
            <FormControl>
              <Input v-model="values.nomorKontak" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="lokasiPerusahaan" v-slot="{ field }">
          <FormItem>
            <FormLabel
              >Lokasi Perusahaan <span class="text-red-500">*</span></FormLabel
            >
            <FormControl>
              <Input v-model="values.lokasiPerusahaan" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorSIUP" v-slot="{ field }">
          <FormItem>
            <FormLabel
              >Nomor SIUP <span class="text-red-500">*</span></FormLabel
            >
            <FormControl>
              <Input v-model="values.nomorSIUP" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorNIB" v-slot="{ field }">
          <FormItem>
            <FormLabel>Nomor NIB <span class="text-red-500">*</span></FormLabel>
            <FormControl>
              <Input v-model="values.nomorNIB" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorNPWP" v-slot="{ field }">
          <FormItem>
            <FormLabel
              >Nomor NPWP <span class="text-red-500">*</span></FormLabel
            >
            <FormControl>
              <Input v-model="values.nomorNPWP" required />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>

      <div class="mt-6">
        <Button type="submit" :disabled="isLoading">
          {{ isLoading ? "Menyimpan..." : "Simpan" }}
        </Button>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";

const values = ref({
  namaPerusahaan: "",
  namaPIC: "",
  email: "",
  nomorKontak: "",
  lokasiPerusahaan: "",
  nomorSIUP: "",
  nomorNIB: "",
  nomorNPWP: "",
});

const isLoading = ref(false);
const errorMessage = ref(null);

const onSubmit = async () => {
  try {
    isLoading.value = true;
    const response = await $fetch("/api/vendors", {
      method: "POST",
      body: values.value,
    });

    console.log("Data vendor berhasil ditambahkan:", response);
    // Reset form setelah berhasil
    values.value = {
      namaPerusahaan: "",
      namaPIC: "",
      email: "",
      nomorKontak: "",
      lokasiPerusahaan: "",
      nomorSIUP: "",
      nomorNIB: "",
      nomorNPWP: "",
    };
  } catch (error) {
    console.error("Terjadi kesalahan saat menambahkan vendor:", error);
    errorMessage.value = "Gagal menambahkan vendor. Silakan coba lagi.";
  } finally {
    isLoading.value = false;
  }
};

// Fetch data vendor jika diperlukan
const fetchVendors = async () => {
  try {
    isLoading.value = true;
    const response = await $fetch("/api/vendors", {
      method: "GET",
    });
    console.log("Data vendor:", response);
  } catch (error) {
    console.error("Gagal mengambil data vendor:", error);
    errorMessage.value = "Gagal mengambil data vendor";
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchVendors();
});
</script>
