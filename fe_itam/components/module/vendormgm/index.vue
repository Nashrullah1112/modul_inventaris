<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Tambah Vendor</h1>

    <Form @submit="onSubmit">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="namaPIC">
          <FormItem>
            <FormLabel>Nama PIC</FormLabel>
            <FormControl>
              <Input v-model="values.namaPIC" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="email">
          <FormItem>
            <FormLabel>Email</FormLabel>
            <FormControl>
              <Input v-model="values.email" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorKontak">
          <FormItem>
            <FormLabel>Nomor Kontak</FormLabel>
            <FormControl>
              <Input v-model="values.nomorKontak" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="lokasiPerusahaan">
          <FormItem>
            <FormLabel>Lokasi Perusahaan</FormLabel>
            <FormControl>
              <Input v-model="values.lokasiPerusahaan" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorSIUP">
          <FormItem>
            <FormLabel>Nomor SIUP</FormLabel>
            <FormControl>
              <Input v-model="values.nomorSIUP" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorNIB">
          <FormItem>
            <FormLabel>Nomor NIB</FormLabel>
            <FormControl>
              <Input v-model="values.nomorNIB" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="nomorNPWP">
          <FormItem>
            <FormLabel>Nomor NPWP</FormLabel>
            <FormControl>
              <Input v-model="values.nomorNPWP" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { ref } from "vue";

const values = ref({
  namaPIC: "",
  email: "",
  nomorKontak: "",
  lokasiPerusahaan: "",
  nomorSIUP: "",
  nomorNIB: "",
  nomorNPWP: "",
});

const onSubmit = async () => {
  try {
    // Tambahkan logika untuk mengirimkan data ke backend
    console.log("Data vendor berhasil ditambahkan:", values.value);
  } catch (error) {
    console.error("Terjadi kesalahan saat menambahkan vendor:", error);
  }
};
</script>
